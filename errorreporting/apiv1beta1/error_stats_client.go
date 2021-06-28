// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package errorreporting

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newErrorStatsClientHook clientHook

// ErrorStatsCallOptions contains the retry settings for each method of ErrorStatsClient.
type ErrorStatsCallOptions struct {
	ListGroupStats []gax.CallOption
	ListEvents     []gax.CallOption
	DeleteEvents   []gax.CallOption
}

func defaultErrorStatsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("clouderrorreporting.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("clouderrorreporting.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://clouderrorreporting.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultErrorStatsCallOptions() *ErrorStatsCallOptions {
	return &ErrorStatsCallOptions{
		ListGroupStats: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListEvents: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteEvents: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalErrorStatsClient is an interface that defines the methods availaible from Error Reporting API.
type internalErrorStatsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListGroupStats(context.Context, *clouderrorreportingpb.ListGroupStatsRequest, ...gax.CallOption) *ErrorGroupStatsIterator
	ListEvents(context.Context, *clouderrorreportingpb.ListEventsRequest, ...gax.CallOption) *ErrorEventIterator
	DeleteEvents(context.Context, *clouderrorreportingpb.DeleteEventsRequest, ...gax.CallOption) (*clouderrorreportingpb.DeleteEventsResponse, error)
}

// ErrorStatsClient is a client for interacting with Error Reporting API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// An API for retrieving and managing error statistics as well as data for
// individual events.
type ErrorStatsClient struct {
	// The internal transport-dependent client.
	internalClient internalErrorStatsClient

	// The call options for this service.
	CallOptions *ErrorStatsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ErrorStatsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ErrorStatsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ErrorStatsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListGroupStats lists the specified groups.
func (c *ErrorStatsClient) ListGroupStats(ctx context.Context, req *clouderrorreportingpb.ListGroupStatsRequest, opts ...gax.CallOption) *ErrorGroupStatsIterator {
	return c.internalClient.ListGroupStats(ctx, req, opts...)
}

// ListEvents lists the specified events.
func (c *ErrorStatsClient) ListEvents(ctx context.Context, req *clouderrorreportingpb.ListEventsRequest, opts ...gax.CallOption) *ErrorEventIterator {
	return c.internalClient.ListEvents(ctx, req, opts...)
}

// DeleteEvents deletes all error events of a given project.
func (c *ErrorStatsClient) DeleteEvents(ctx context.Context, req *clouderrorreportingpb.DeleteEventsRequest, opts ...gax.CallOption) (*clouderrorreportingpb.DeleteEventsResponse, error) {
	return c.internalClient.DeleteEvents(ctx, req, opts...)
}

// errorStatsGRPCClient is a client for interacting with Error Reporting API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type errorStatsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ErrorStatsClient
	CallOptions **ErrorStatsCallOptions

	// The gRPC API client.
	errorStatsClient clouderrorreportingpb.ErrorStatsServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewErrorStatsClient creates a new error stats service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// An API for retrieving and managing error statistics as well as data for
// individual events.
func NewErrorStatsClient(ctx context.Context, opts ...option.ClientOption) (*ErrorStatsClient, error) {
	clientOpts := defaultErrorStatsGRPCClientOptions()
	if newErrorStatsClientHook != nil {
		hookOpts, err := newErrorStatsClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := ErrorStatsClient{CallOptions: defaultErrorStatsCallOptions()}

	c := &errorStatsGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		errorStatsClient: clouderrorreportingpb.NewErrorStatsServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *errorStatsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *errorStatsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *errorStatsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *errorStatsGRPCClient) ListGroupStats(ctx context.Context, req *clouderrorreportingpb.ListGroupStatsRequest, opts ...gax.CallOption) *ErrorGroupStatsIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListGroupStats[0:len((*c.CallOptions).ListGroupStats):len((*c.CallOptions).ListGroupStats)], opts...)
	it := &ErrorGroupStatsIterator{}
	req = proto.Clone(req).(*clouderrorreportingpb.ListGroupStatsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*clouderrorreportingpb.ErrorGroupStats, string, error) {
		var resp *clouderrorreportingpb.ListGroupStatsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.errorStatsClient.ListGroupStats(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetErrorGroupStats(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

func (c *errorStatsGRPCClient) ListEvents(ctx context.Context, req *clouderrorreportingpb.ListEventsRequest, opts ...gax.CallOption) *ErrorEventIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListEvents[0:len((*c.CallOptions).ListEvents):len((*c.CallOptions).ListEvents)], opts...)
	it := &ErrorEventIterator{}
	req = proto.Clone(req).(*clouderrorreportingpb.ListEventsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*clouderrorreportingpb.ErrorEvent, string, error) {
		var resp *clouderrorreportingpb.ListEventsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.errorStatsClient.ListEvents(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetErrorEvents(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}
	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()
	return it
}

func (c *errorStatsGRPCClient) DeleteEvents(ctx context.Context, req *clouderrorreportingpb.DeleteEventsRequest, opts ...gax.CallOption) (*clouderrorreportingpb.DeleteEventsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "project_name", url.QueryEscape(req.GetProjectName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteEvents[0:len((*c.CallOptions).DeleteEvents):len((*c.CallOptions).DeleteEvents)], opts...)
	var resp *clouderrorreportingpb.DeleteEventsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.errorStatsClient.DeleteEvents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ErrorEventIterator manages a stream of *clouderrorreportingpb.ErrorEvent.
type ErrorEventIterator struct {
	items    []*clouderrorreportingpb.ErrorEvent
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*clouderrorreportingpb.ErrorEvent, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ErrorEventIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ErrorEventIterator) Next() (*clouderrorreportingpb.ErrorEvent, error) {
	var item *clouderrorreportingpb.ErrorEvent
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ErrorEventIterator) bufLen() int {
	return len(it.items)
}

func (it *ErrorEventIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// ErrorGroupStatsIterator manages a stream of *clouderrorreportingpb.ErrorGroupStats.
type ErrorGroupStatsIterator struct {
	items    []*clouderrorreportingpb.ErrorGroupStats
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*clouderrorreportingpb.ErrorGroupStats, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ErrorGroupStatsIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ErrorGroupStatsIterator) Next() (*clouderrorreportingpb.ErrorGroupStats, error) {
	var item *clouderrorreportingpb.ErrorGroupStats
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ErrorGroupStatsIterator) bufLen() int {
	return len(it.items)
}

func (it *ErrorGroupStatsIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
