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

package cx

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	"github.com/golang/protobuf/proto"
	structpb "github.com/golang/protobuf/ptypes/struct"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newVersionsClientHook clientHook

// VersionsCallOptions contains the retry settings for each method of VersionsClient.
type VersionsCallOptions struct {
	ListVersions  []gax.CallOption
	GetVersion    []gax.CallOption
	CreateVersion []gax.CallOption
	UpdateVersion []gax.CallOption
	DeleteVersion []gax.CallOption
	LoadVersion   []gax.CallOption
}

func defaultVersionsGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dialogflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dialogflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dialogflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultVersionsCallOptions() *VersionsCallOptions {
	return &VersionsCallOptions{
		ListVersions: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetVersion: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateVersion: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdateVersion: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeleteVersion: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		LoadVersion: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalVersionsClient is an interface that defines the methods availaible from Dialogflow API.
type internalVersionsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListVersions(context.Context, *cxpb.ListVersionsRequest, ...gax.CallOption) *VersionIterator
	GetVersion(context.Context, *cxpb.GetVersionRequest, ...gax.CallOption) (*cxpb.Version, error)
	CreateVersion(context.Context, *cxpb.CreateVersionRequest, ...gax.CallOption) (*CreateVersionOperation, error)
	CreateVersionOperation(name string) *CreateVersionOperation
	UpdateVersion(context.Context, *cxpb.UpdateVersionRequest, ...gax.CallOption) (*cxpb.Version, error)
	DeleteVersion(context.Context, *cxpb.DeleteVersionRequest, ...gax.CallOption) error
	LoadVersion(context.Context, *cxpb.LoadVersionRequest, ...gax.CallOption) (*LoadVersionOperation, error)
	LoadVersionOperation(name string) *LoadVersionOperation
}

// VersionsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Versions.
type VersionsClient struct {
	// The internal transport-dependent client.
	internalClient internalVersionsClient

	// The call options for this service.
	CallOptions *VersionsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *VersionsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *VersionsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *VersionsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListVersions returns the list of all versions in the specified Flow.
func (c *VersionsClient) ListVersions(ctx context.Context, req *cxpb.ListVersionsRequest, opts ...gax.CallOption) *VersionIterator {
	return c.internalClient.ListVersions(ctx, req, opts...)
}

// GetVersion retrieves the specified Version.
func (c *VersionsClient) GetVersion(ctx context.Context, req *cxpb.GetVersionRequest, opts ...gax.CallOption) (*cxpb.Version, error) {
	return c.internalClient.GetVersion(ctx, req, opts...)
}

// CreateVersion creates a Version in the specified Flow.
func (c *VersionsClient) CreateVersion(ctx context.Context, req *cxpb.CreateVersionRequest, opts ...gax.CallOption) (*CreateVersionOperation, error) {
	return c.internalClient.CreateVersion(ctx, req, opts...)
}

// CreateVersionOperation returns a new CreateVersionOperation from a given name.
// The name must be that of a previously created CreateVersionOperation, possibly from a different process.
func (c *VersionsClient) CreateVersionOperation(name string) *CreateVersionOperation {
	return c.internalClient.CreateVersionOperation(name)
}

// UpdateVersion updates the specified Version.
func (c *VersionsClient) UpdateVersion(ctx context.Context, req *cxpb.UpdateVersionRequest, opts ...gax.CallOption) (*cxpb.Version, error) {
	return c.internalClient.UpdateVersion(ctx, req, opts...)
}

// DeleteVersion deletes the specified Version.
func (c *VersionsClient) DeleteVersion(ctx context.Context, req *cxpb.DeleteVersionRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteVersion(ctx, req, opts...)
}

// LoadVersion loads resources in the specified version to the draft flow.
func (c *VersionsClient) LoadVersion(ctx context.Context, req *cxpb.LoadVersionRequest, opts ...gax.CallOption) (*LoadVersionOperation, error) {
	return c.internalClient.LoadVersion(ctx, req, opts...)
}

// LoadVersionOperation returns a new LoadVersionOperation from a given name.
// The name must be that of a previously created LoadVersionOperation, possibly from a different process.
func (c *VersionsClient) LoadVersionOperation(name string) *LoadVersionOperation {
	return c.internalClient.LoadVersionOperation(name)
}

// versionsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type versionsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing VersionsClient
	CallOptions **VersionsCallOptions

	// The gRPC API client.
	versionsClient cxpb.VersionsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewVersionsClient creates a new versions client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Versions.
func NewVersionsClient(ctx context.Context, opts ...option.ClientOption) (*VersionsClient, error) {
	clientOpts := defaultVersionsGRPCClientOptions()
	if newVersionsClientHook != nil {
		hookOpts, err := newVersionsClientHook(ctx, clientHookParams{})
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
	client := VersionsClient{CallOptions: defaultVersionsCallOptions()}

	c := &versionsGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		versionsClient:   cxpb.NewVersionsClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *versionsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *versionsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *versionsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *versionsGRPCClient) ListVersions(ctx context.Context, req *cxpb.ListVersionsRequest, opts ...gax.CallOption) *VersionIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListVersions[0:len((*c.CallOptions).ListVersions):len((*c.CallOptions).ListVersions)], opts...)
	it := &VersionIterator{}
	req = proto.Clone(req).(*cxpb.ListVersionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.Version, string, error) {
		var resp *cxpb.ListVersionsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.versionsClient.ListVersions(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetVersions(), resp.GetNextPageToken(), nil
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

func (c *versionsGRPCClient) GetVersion(ctx context.Context, req *cxpb.GetVersionRequest, opts ...gax.CallOption) (*cxpb.Version, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetVersion[0:len((*c.CallOptions).GetVersion):len((*c.CallOptions).GetVersion)], opts...)
	var resp *cxpb.Version
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.GetVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *versionsGRPCClient) CreateVersion(ctx context.Context, req *cxpb.CreateVersionRequest, opts ...gax.CallOption) (*CreateVersionOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateVersion[0:len((*c.CallOptions).CreateVersion):len((*c.CallOptions).CreateVersion)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.CreateVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *versionsGRPCClient) UpdateVersion(ctx context.Context, req *cxpb.UpdateVersionRequest, opts ...gax.CallOption) (*cxpb.Version, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "version.name", url.QueryEscape(req.GetVersion().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateVersion[0:len((*c.CallOptions).UpdateVersion):len((*c.CallOptions).UpdateVersion)], opts...)
	var resp *cxpb.Version
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.UpdateVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *versionsGRPCClient) DeleteVersion(ctx context.Context, req *cxpb.DeleteVersionRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteVersion[0:len((*c.CallOptions).DeleteVersion):len((*c.CallOptions).DeleteVersion)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.versionsClient.DeleteVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *versionsGRPCClient) LoadVersion(ctx context.Context, req *cxpb.LoadVersionRequest, opts ...gax.CallOption) (*LoadVersionOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).LoadVersion[0:len((*c.CallOptions).LoadVersion):len((*c.CallOptions).LoadVersion)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.LoadVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &LoadVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateVersionOperation manages a long-running operation from CreateVersion.
type CreateVersionOperation struct {
	lro *longrunning.Operation
}

// CreateVersionOperation returns a new CreateVersionOperation from a given name.
// The name must be that of a previously created CreateVersionOperation, possibly from a different process.
func (c *versionsGRPCClient) CreateVersionOperation(name string) *CreateVersionOperation {
	return &CreateVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateVersionOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*cxpb.Version, error) {
	var resp cxpb.Version
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
func (op *CreateVersionOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*cxpb.Version, error) {
	var resp cxpb.Version
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
func (op *CreateVersionOperation) Metadata() (*cxpb.CreateVersionOperationMetadata, error) {
	var meta cxpb.CreateVersionOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateVersionOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateVersionOperation) Name() string {
	return op.lro.Name()
}

// LoadVersionOperation manages a long-running operation from LoadVersion.
type LoadVersionOperation struct {
	lro *longrunning.Operation
}

// LoadVersionOperation returns a new LoadVersionOperation from a given name.
// The name must be that of a previously created LoadVersionOperation, possibly from a different process.
func (c *versionsGRPCClient) LoadVersionOperation(name string) *LoadVersionOperation {
	return &LoadVersionOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *LoadVersionOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.WaitWithInterval(ctx, nil, time.Minute, opts...)
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
func (op *LoadVersionOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *LoadVersionOperation) Metadata() (*structpb.Struct, error) {
	var meta structpb.Struct
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *LoadVersionOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *LoadVersionOperation) Name() string {
	return op.lro.Name()
}

// VersionIterator manages a stream of *cxpb.Version.
type VersionIterator struct {
	items    []*cxpb.Version
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
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.Version, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *VersionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *VersionIterator) Next() (*cxpb.Version, error) {
	var item *cxpb.Version
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *VersionIterator) bufLen() int {
	return len(it.items)
}

func (it *VersionIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
