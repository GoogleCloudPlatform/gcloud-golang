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

package pubsublite

import (
	"context"
	"time"

	vkit "cloud.google.com/go/pubsublite/apiv1"
	pb "google.golang.org/genproto/googleapis/cloud/pubsublite/v1"
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

// BacklogLocation refers to a location with respect to the message backlog.
type BacklogLocation int

const (
	// End refers to the location past all currently published messages. End
	// skips the entire message backlog.
	End BacklogLocation = iota + 1

	// Beginning refers to the location of the oldest retained message.
	Beginning
)

// SeekTarget is the target location to seek a subscription to. Implemented by
// BacklogLocationSeekTarget, PublishTimeSeekTarget, EventTimeSeekTarget.
type SeekTarget interface {
	setRequest(req *pb.SeekSubscriptionRequest)
}

// BacklogLocationSeekTarget is a seek target to a location with respect to the
// message backlog.
type BacklogLocationSeekTarget struct {
	Location BacklogLocation
}

func (b BacklogLocationSeekTarget) setRequest(req *pb.SeekSubscriptionRequest) {
	target := pb.SeekSubscriptionRequest_TAIL
	if b.Location == End {
		target = pb.SeekSubscriptionRequest_HEAD
	}
	req.Target = &pb.SeekSubscriptionRequest_NamedTarget_{
		NamedTarget: target,
	}
}

// PublishTimeSeekTarget is a seek target to a message publish timestamp.
type PublishTimeSeekTarget struct {
	Time time.Time
}

func (p PublishTimeSeekTarget) setRequest(req *pb.SeekSubscriptionRequest) {
	req.Target = &pb.SeekSubscriptionRequest_TimeTarget{
		TimeTarget: &pb.TimeTarget{
			Time: &pb.TimeTarget_PublishTime{tspb.New(p.Time)},
		},
	}
}

// EventTimeSeekTarget is a seek target to a message event timestamp.
type EventTimeSeekTarget struct {
	Time time.Time
}

func (p EventTimeSeekTarget) setRequest(req *pb.SeekSubscriptionRequest) {
	req.Target = &pb.SeekSubscriptionRequest_TimeTarget{
		TimeTarget: &pb.TimeTarget{
			Time: &pb.TimeTarget_EventTime{tspb.New(p.Time)},
		},
	}
}

// SeekSubscriptionOption is reserved for future options.
type SeekSubscriptionOption interface{}

// Metadata for long-running operations.
type OperationMetadata struct {
	// The target of the operation. For example, targets of seeks are
	// subscriptions, structured like:
	// "projects/PROJECT_ID/locations/ZONE/subscriptions/SUBSCRIPTION_ID"
	Target string

	// The verb describing the kind of operation.
	Verb string

	// The time the operation was created.
	CreateTime time.Time

	// The time the operation finished running. Is zero if the operation has not
	// completed.
	EndTime time.Time
}

func protoToOperationMetadata(o *pb.OperationMetadata) (*OperationMetadata, error) {
	if err := o.GetCreateTime().CheckValid(); err != nil {
		return nil, err
	}
	metadata := &OperationMetadata{
		Target:     o.Target,
		Verb:       o.Verb,
		CreateTime: o.GetCreateTime().AsTime(),
	}
	if o.GetEndTime() != nil {
		if err := o.GetEndTime().CheckValid(); err != nil {
			return nil, err
		}
		metadata.EndTime = o.GetEndTime().AsTime()
	}
	return metadata, nil
}

// SeekSubscriptionOperation manages a long-running seek operation from
// AdminClient.SeekSubscription.
type SeekSubscriptionOperation struct {
	op *vkit.SeekSubscriptionOperation
}

// Name returns the path of the seek operation, in the format:
// "projects/PROJECT_ID/locations/LOCATION/operations/OPERATION_ID".
func (s *SeekSubscriptionOperation) Name() string {
	return s.op.Name()
}

// Done returns whether the seek operation has completed.
func (s *SeekSubscriptionOperation) Done() bool {
	return s.op.Done()
}

// Metadata returns metadata associated with the seek operation. To get the
// latest metadata, call this method after a successful call to Wait.
func (s *SeekSubscriptionOperation) Metadata() (*OperationMetadata, error) {
	m, err := s.op.Metadata()
	if err != nil {
		return nil, err
	}
	return protoToOperationMetadata(m)
}

// Wait polls until the seek operation is complete and returns one of the
// following:
//  - nil if the operation is complete and succeeded.
//  - failure reason if the operation is complete and failed.
//  - error if polling the operation status failed due to a non-retryable error.
func (s *SeekSubscriptionOperation) Wait(ctx context.Context) error {
	_, err := s.op.Wait(ctx)
	return err
}
