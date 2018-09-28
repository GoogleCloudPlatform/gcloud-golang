// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pubsub

import (
	"io"
	"log"
	"sync"
	"time"

	gax "github.com/googleapis/gax-go"
	"golang.org/x/net/context"
	pb "google.golang.org/genproto/googleapis/pubsub/v1"
	"google.golang.org/grpc"
)

// A pullStream supports the methods of a StreamingPullClient, but re-opens
// the stream on a retryable error.
type pullStream struct {
	ctx  context.Context
	open func() (pb.Subscriber_StreamingPullClient, error)

	mu  sync.Mutex
	spc *pb.Subscriber_StreamingPullClient
	err error // permanent error
}

// for testing
type streamingPullFunc func(context.Context, ...gax.CallOption) (pb.Subscriber_StreamingPullClient, error)

func newPullStream(ctx context.Context, streamingPull streamingPullFunc, subName string) *pullStream {
	ctx = withSubscriptionKey(ctx, subName)
	return &pullStream{
		ctx: ctx,
		open: func() (pb.Subscriber_StreamingPullClient, error) {
			spc, err := streamingPull(ctx, gax.WithGRPCOptions(grpc.MaxCallRecvMsgSize(maxSendRecvBytes)))
			if err == nil {
				recordStat(ctx, StreamRequestCount, 1)
				log.Print("starting open")
				err = spc.Send(&pb.StreamingPullRequest{
					Subscription: subName,
					// We modack messages when we receive them, so this value doesn't matter too much.
					StreamAckDeadlineSeconds: 60,
				})
				log.Printf("open returned %v", err)
			}
			if err != nil {
				return nil, err
			}
			return spc, nil
		},
	}
}

// get returns either a valid *StreamingPullClient (SPC), or a permanent error.
// If the argument is nil, this is the first call for an RPC, and the current
// SPC will be returned (or a new one will be opened). Otherwise, this call is a
// request to re-open the stream because of a retryable error, and the argument
// is a pointer to the SPC that returned the error.
func (s *pullStream) get(spc *pb.Subscriber_StreamingPullClient) (*pb.Subscriber_StreamingPullClient, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// A stored error is permanent.
	if s.err != nil {
		return nil, s.err
	}
	// If the context is done, so are we.
	select {
	case <-s.ctx.Done():
		s.err = s.ctx.Err()
		return nil, s.err
	default:
	}
	// TODO(jba): We can use the following instead of the above after we drop support for 1.8:
	// s.err = s.ctx.Err()
	// if s.err != nil {
	// 	return nil, s.err
	// }

	// If the current and argument SPCs differ, return the current one. This subsumes two cases:
	// 1. We have an SPC and the caller is getting the stream for the first time.
	// 2. The caller wants to retry, but they have an older SPC; we've already retried.
	if spc != s.spc {
		return s.spc, nil
	}
	// Either this is the very first call on this stream (s.spc == nil), or we have a valid
	// retry request. Either way, open a new stream.
	// The lock is held here for a long time, but it doesn't matter because no callers could get
	// anything done anyway.
	s.spc = new(pb.Subscriber_StreamingPullClient)
	*s.spc, s.err = s.openWithRetry() // Any error from openWithRetry is permanent.
	return s.spc, s.err
}

func (s *pullStream) openWithRetry() (pb.Subscriber_StreamingPullClient, error) {
	var bo gax.Backoff
	for {
		recordStat(s.ctx, StreamOpenCount, 1)
		spc, err := s.open()
		if err != nil && isRetryable(err) {
			recordStat(s.ctx, StreamRetryCount, 1)
			if err := gax.Sleep(s.ctx, bo.Pause()); err != nil {
				return nil, err
			}
			continue
		}
		return spc, err
	}
}

func (s *pullStream) call(f func(pb.Subscriber_StreamingPullClient) error) error {
	var (
		spc *pb.Subscriber_StreamingPullClient
		err error
		bo  gax.Backoff
	)
	for {
		spc, err = s.get(spc)
		if err != nil {
			return err
		}
		start := time.Now()
		err = f(*spc)
		if err != nil {
			if isRetryable(err) {
				recordStat(s.ctx, StreamRetryCount, 1)
				if time.Since(start) < 30*time.Second { // don't sleep if we've been blocked for a while
					if err := gax.Sleep(s.ctx, bo.Pause()); err != nil {
						return err
					}
				}
				continue
			}
			s.mu.Lock()
			s.err = err
			s.mu.Unlock()
		}
		return err
	}
}

func (s *pullStream) Send(req *pb.StreamingPullRequest) error {
	return s.call(func(spc pb.Subscriber_StreamingPullClient) error {
		recordStat(s.ctx, AckCount, int64(len(req.AckIds)))
		zeroes := 0
		for _, mds := range req.ModifyDeadlineSeconds {
			if mds == 0 {
				zeroes++
			}
		}
		recordStat(s.ctx, NackCount, int64(zeroes))
		recordStat(s.ctx, ModAckCount, int64(len(req.ModifyDeadlineSeconds)-zeroes))
		recordStat(s.ctx, StreamRequestCount, 1)
		log.Print("starting Send")
		err := spc.Send(req)
		log.Printf("Send returned %v", err)
		return err
	})
}

func (s *pullStream) Recv() (*pb.StreamingPullResponse, error) {
	var res *pb.StreamingPullResponse
	err := s.call(func(spc pb.Subscriber_StreamingPullClient) error {
		var err error
		recordStat(s.ctx, StreamResponseCount, 1)
		log.Print("starting Recv")
		res, err = spc.Recv()
		log.Printf("Recv returned %v", err)
		if err == nil {
			recordStat(s.ctx, PullCount, int64(len(res.ReceivedMessages)))
		}
		return err
	})
	return res, err
}

func (s *pullStream) CloseSend() error {
	err := s.call(func(spc pb.Subscriber_StreamingPullClient) error {
		return spc.CloseSend()
	})
	s.mu.Lock()
	s.err = io.EOF // should not be retried
	s.mu.Unlock()
	return err
}
