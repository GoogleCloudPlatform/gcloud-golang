// Copyright 2020 Google LLC
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

package talent

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newJobClientHook clientHook

// JobCallOptions contains the retry settings for each method of JobClient.
type JobCallOptions struct {
	CreateJob          []gax.CallOption
	BatchCreateJobs    []gax.CallOption
	GetJob             []gax.CallOption
	UpdateJob          []gax.CallOption
	BatchUpdateJobs    []gax.CallOption
	DeleteJob          []gax.CallOption
	BatchDeleteJobs    []gax.CallOption
	ListJobs           []gax.CallOption
	SearchJobs         []gax.CallOption
	SearchJobsForAlert []gax.CallOption
}

func defaultJobClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("jobs.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("jobs.mtls.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultJobCallOptions() *JobCallOptions {
	return &JobCallOptions{
		CreateJob:       []gax.CallOption{},
		BatchCreateJobs: []gax.CallOption{},
		GetJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateJob:       []gax.CallOption{},
		BatchUpdateJobs: []gax.CallOption{},
		DeleteJob: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchDeleteJobs: []gax.CallOption{},
		ListJobs: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SearchJobs:         []gax.CallOption{},
		SearchJobsForAlert: []gax.CallOption{},
	}
}

// JobClient is a client for interacting with Cloud Talent Solution API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type JobClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	jobClient talentpb.JobServiceClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *JobCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewJobClient creates a new job service client.
//
// A service handles job management, including job CRUD, enumeration and search.
func NewJobClient(ctx context.Context, opts ...option.ClientOption) (*JobClient, error) {
	clientOpts := defaultJobClientOptions()

	if newJobClientHook != nil {
		hookOpts, err := newJobClientHook(ctx, clientHookParams{})
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
	c := &JobClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultJobCallOptions(),

		jobClient: talentpb.NewJobServiceClient(connPool),
	}
	c.setGoogleClientInfo()

	c.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *JobClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *JobClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *JobClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// CreateJob creates a new job.
//
// Typically, the job becomes searchable within 10 seconds, but it may take
// up to 5 minutes.
func (c *JobClient) CreateJob(ctx context.Context, req *talentpb.CreateJobRequest, opts ...gax.CallOption) (*talentpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateJob[0:len(c.CallOptions.CreateJob):len(c.CallOptions.CreateJob)], opts...)
	var resp *talentpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.CreateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchCreateJobs begins executing a batch create jobs operation.
func (c *JobClient) BatchCreateJobs(ctx context.Context, req *talentpb.BatchCreateJobsRequest, opts ...gax.CallOption) (*BatchCreateJobsOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchCreateJobs[0:len(c.CallOptions.BatchCreateJobs):len(c.CallOptions.BatchCreateJobs)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.BatchCreateJobs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchCreateJobsOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// GetJob retrieves the specified job, whose status is OPEN or recently EXPIRED
// within the last 90 days.
func (c *JobClient) GetJob(ctx context.Context, req *talentpb.GetJobRequest, opts ...gax.CallOption) (*talentpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetJob[0:len(c.CallOptions.GetJob):len(c.CallOptions.GetJob)], opts...)
	var resp *talentpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.GetJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateJob updates specified job.
//
// Typically, updated contents become visible in search results within 10
// seconds, but it may take up to 5 minutes.
func (c *JobClient) UpdateJob(ctx context.Context, req *talentpb.UpdateJobRequest, opts ...gax.CallOption) (*talentpb.Job, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "job.name", url.QueryEscape(req.GetJob().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateJob[0:len(c.CallOptions.UpdateJob):len(c.CallOptions.UpdateJob)], opts...)
	var resp *talentpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.UpdateJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchUpdateJobs begins executing a batch update jobs operation.
func (c *JobClient) BatchUpdateJobs(ctx context.Context, req *talentpb.BatchUpdateJobsRequest, opts ...gax.CallOption) (*BatchUpdateJobsOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchUpdateJobs[0:len(c.CallOptions.BatchUpdateJobs):len(c.CallOptions.BatchUpdateJobs)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.BatchUpdateJobs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchUpdateJobsOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// DeleteJob deletes the specified job.
//
// Typically, the job becomes unsearchable within 10 seconds, but it may take
// up to 5 minutes.
func (c *JobClient) DeleteJob(ctx context.Context, req *talentpb.DeleteJobRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteJob[0:len(c.CallOptions.DeleteJob):len(c.CallOptions.DeleteJob)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.jobClient.DeleteJob(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// BatchDeleteJobs deletes a list of Jobs by filter.
func (c *JobClient) BatchDeleteJobs(ctx context.Context, req *talentpb.BatchDeleteJobsRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchDeleteJobs[0:len(c.CallOptions.BatchDeleteJobs):len(c.CallOptions.BatchDeleteJobs)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.jobClient.BatchDeleteJobs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// ListJobs lists jobs by filter.
func (c *JobClient) ListJobs(ctx context.Context, req *talentpb.ListJobsRequest, opts ...gax.CallOption) *JobIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListJobs[0:len(c.CallOptions.ListJobs):len(c.CallOptions.ListJobs)], opts...)
	it := &JobIterator{}
	req = proto.Clone(req).(*talentpb.ListJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.Job, string, error) {
		var resp *talentpb.ListJobsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.jobClient.ListJobs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetJobs(), resp.GetNextPageToken(), nil
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

// SearchJobs searches for jobs using the provided SearchJobsRequest.
//
// This call constrains the visibility of jobs
// present in the database, and only returns jobs that the caller has
// permission to search against.
func (c *JobClient) SearchJobs(ctx context.Context, req *talentpb.SearchJobsRequest, opts ...gax.CallOption) (*talentpb.SearchJobsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 30000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SearchJobs[0:len(c.CallOptions.SearchJobs):len(c.CallOptions.SearchJobs)], opts...)
	var resp *talentpb.SearchJobsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.jobClient.SearchJobs(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchJobsForAlert searches for jobs using the provided SearchJobsRequest.
//
// This API call is intended for the use case of targeting passive job
// seekers (for example, job seekers who have signed up to receive email
// alerts about potential job opportunities), and has different algorithmic
// adjustments that are targeted to passive job seekers.
//
// This call constrains the visibility of jobs
// present in the database, and only returns jobs the caller has
// permission to search against.
func (c *JobClient) SearchJobsForAlert(ctx context.Context, req *talentpb.SearchJobsRequest, opts ...gax.CallOption) *SearchJobsResponse_MatchingJobIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.SearchJobsForAlert[0:len(c.CallOptions.SearchJobsForAlert):len(c.CallOptions.SearchJobsForAlert)], opts...)
	it := &SearchJobsResponse_MatchingJobIterator{}
	req = proto.Clone(req).(*talentpb.SearchJobsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*talentpb.SearchJobsResponse_MatchingJob, string, error) {
		var resp *talentpb.SearchJobsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.jobClient.SearchJobsForAlert(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetMatchingJobs(), resp.GetNextPageToken(), nil
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

// BatchCreateJobsOperation manages a long-running operation from BatchCreateJobs.
type BatchCreateJobsOperation struct {
	lro *longrunning.Operation
}

// BatchCreateJobsOperation returns a new BatchCreateJobsOperation from a given name.
// The name must be that of a previously created BatchCreateJobsOperation, possibly from a different process.
func (c *JobClient) BatchCreateJobsOperation(name string) *BatchCreateJobsOperation {
	return &BatchCreateJobsOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchCreateJobsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*talentpb.JobOperationResult, error) {
	var resp talentpb.JobOperationResult
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *BatchCreateJobsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*talentpb.JobOperationResult, error) {
	var resp talentpb.JobOperationResult
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *BatchCreateJobsOperation) Metadata() (*talentpb.BatchOperationMetadata, error) {
	var meta talentpb.BatchOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchCreateJobsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchCreateJobsOperation) Name() string {
	return op.lro.Name()
}

// BatchUpdateJobsOperation manages a long-running operation from BatchUpdateJobs.
type BatchUpdateJobsOperation struct {
	lro *longrunning.Operation
}

// BatchUpdateJobsOperation returns a new BatchUpdateJobsOperation from a given name.
// The name must be that of a previously created BatchUpdateJobsOperation, possibly from a different process.
func (c *JobClient) BatchUpdateJobsOperation(name string) *BatchUpdateJobsOperation {
	return &BatchUpdateJobsOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchUpdateJobsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*talentpb.JobOperationResult, error) {
	var resp talentpb.JobOperationResult
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *BatchUpdateJobsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*talentpb.JobOperationResult, error) {
	var resp talentpb.JobOperationResult
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *BatchUpdateJobsOperation) Metadata() (*talentpb.BatchOperationMetadata, error) {
	var meta talentpb.BatchOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchUpdateJobsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchUpdateJobsOperation) Name() string {
	return op.lro.Name()
}

// JobIterator manages a stream of *talentpb.Job.
type JobIterator struct {
	items    []*talentpb.Job
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
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.Job, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *JobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *JobIterator) Next() (*talentpb.Job, error) {
	var item *talentpb.Job
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *JobIterator) bufLen() int {
	return len(it.items)
}

func (it *JobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// SearchJobsResponse_MatchingJobIterator manages a stream of *talentpb.SearchJobsResponse_MatchingJob.
type SearchJobsResponse_MatchingJobIterator struct {
	items    []*talentpb.SearchJobsResponse_MatchingJob
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
	InternalFetch func(pageSize int, pageToken string) (results []*talentpb.SearchJobsResponse_MatchingJob, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *SearchJobsResponse_MatchingJobIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *SearchJobsResponse_MatchingJobIterator) Next() (*talentpb.SearchJobsResponse_MatchingJob, error) {
	var item *talentpb.SearchJobsResponse_MatchingJob
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *SearchJobsResponse_MatchingJobIterator) bufLen() int {
	return len(it.items)
}

func (it *SearchJobsResponse_MatchingJobIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
