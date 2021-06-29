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

package dialogflow

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	structpb "github.com/golang/protobuf/ptypes/struct"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newIntentsClientHook clientHook

// IntentsCallOptions contains the retry settings for each method of IntentsClient.
type IntentsCallOptions struct {
	ListIntents        []gax.CallOption
	GetIntent          []gax.CallOption
	CreateIntent       []gax.CallOption
	UpdateIntent       []gax.CallOption
	DeleteIntent       []gax.CallOption
	BatchUpdateIntents []gax.CallOption
	BatchDeleteIntents []gax.CallOption
}

func defaultIntentsGRPCClientOptions() []option.ClientOption {
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

func defaultIntentsCallOptions() *IntentsCallOptions {
	return &IntentsCallOptions{
		ListIntents: []gax.CallOption{
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
		GetIntent: []gax.CallOption{
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
		CreateIntent: []gax.CallOption{
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
		UpdateIntent: []gax.CallOption{
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
		DeleteIntent: []gax.CallOption{
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
		BatchUpdateIntents: []gax.CallOption{
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
		BatchDeleteIntents: []gax.CallOption{
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

// internalIntentsClient is an interface that defines the methods availaible from Dialogflow API.
type internalIntentsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListIntents(context.Context, *dialogflowpb.ListIntentsRequest, ...gax.CallOption) *IntentIterator
	GetIntent(context.Context, *dialogflowpb.GetIntentRequest, ...gax.CallOption) (*dialogflowpb.Intent, error)
	CreateIntent(context.Context, *dialogflowpb.CreateIntentRequest, ...gax.CallOption) (*dialogflowpb.Intent, error)
	UpdateIntent(context.Context, *dialogflowpb.UpdateIntentRequest, ...gax.CallOption) (*dialogflowpb.Intent, error)
	DeleteIntent(context.Context, *dialogflowpb.DeleteIntentRequest, ...gax.CallOption) error
	BatchUpdateIntents(context.Context, *dialogflowpb.BatchUpdateIntentsRequest, ...gax.CallOption) (*BatchUpdateIntentsOperation, error)
	BatchUpdateIntentsOperation(name string) *BatchUpdateIntentsOperation
	BatchDeleteIntents(context.Context, *dialogflowpb.BatchDeleteIntentsRequest, ...gax.CallOption) (*BatchDeleteIntentsOperation, error)
	BatchDeleteIntentsOperation(name string) *BatchDeleteIntentsOperation
}

// IntentsClient is a client for interacting with Dialogflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service for managing Intents.
type IntentsClient struct {
	// The internal transport-dependent client.
	internalClient internalIntentsClient

	// The call options for this service.
	CallOptions *IntentsCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *IntentsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *IntentsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *IntentsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListIntents returns the list of all intents in the specified agent.
func (c *IntentsClient) ListIntents(ctx context.Context, req *dialogflowpb.ListIntentsRequest, opts ...gax.CallOption) *IntentIterator {
	return c.internalClient.ListIntents(ctx, req, opts...)
}

// GetIntent retrieves the specified intent.
func (c *IntentsClient) GetIntent(ctx context.Context, req *dialogflowpb.GetIntentRequest, opts ...gax.CallOption) (*dialogflowpb.Intent, error) {
	return c.internalClient.GetIntent(ctx, req, opts...)
}

// CreateIntent creates an intent in the specified agent.
//
// Note: You should always train an agent prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/es/docs/training).
func (c *IntentsClient) CreateIntent(ctx context.Context, req *dialogflowpb.CreateIntentRequest, opts ...gax.CallOption) (*dialogflowpb.Intent, error) {
	return c.internalClient.CreateIntent(ctx, req, opts...)
}

// UpdateIntent updates the specified intent.
//
// Note: You should always train an agent prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/es/docs/training).
func (c *IntentsClient) UpdateIntent(ctx context.Context, req *dialogflowpb.UpdateIntentRequest, opts ...gax.CallOption) (*dialogflowpb.Intent, error) {
	return c.internalClient.UpdateIntent(ctx, req, opts...)
}

// DeleteIntent deletes the specified intent and its direct or indirect followup intents.
//
// Note: You should always train an agent prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/es/docs/training).
func (c *IntentsClient) DeleteIntent(ctx context.Context, req *dialogflowpb.DeleteIntentRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteIntent(ctx, req, opts...)
}

// BatchUpdateIntents updates/Creates multiple intents in the specified agent.
//
// Note: You should always train an agent prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/es/docs/training).
func (c *IntentsClient) BatchUpdateIntents(ctx context.Context, req *dialogflowpb.BatchUpdateIntentsRequest, opts ...gax.CallOption) (*BatchUpdateIntentsOperation, error) {
	return c.internalClient.BatchUpdateIntents(ctx, req, opts...)
}

// BatchUpdateIntentsOperation returns a new BatchUpdateIntentsOperation from a given name.
// The name must be that of a previously created BatchUpdateIntentsOperation, possibly from a different process.
func (c *IntentsClient) BatchUpdateIntentsOperation(name string) *BatchUpdateIntentsOperation {
	return c.internalClient.BatchUpdateIntentsOperation(name)
}

// BatchDeleteIntents deletes intents in the specified agent.
//
// Note: You should always train an agent prior to sending it queries. See the
// training
// documentation (at https://cloud.google.com/dialogflow/es/docs/training).
func (c *IntentsClient) BatchDeleteIntents(ctx context.Context, req *dialogflowpb.BatchDeleteIntentsRequest, opts ...gax.CallOption) (*BatchDeleteIntentsOperation, error) {
	return c.internalClient.BatchDeleteIntents(ctx, req, opts...)
}

// BatchDeleteIntentsOperation returns a new BatchDeleteIntentsOperation from a given name.
// The name must be that of a previously created BatchDeleteIntentsOperation, possibly from a different process.
func (c *IntentsClient) BatchDeleteIntentsOperation(name string) *BatchDeleteIntentsOperation {
	return c.internalClient.BatchDeleteIntentsOperation(name)
}

// intentsGRPCClient is a client for interacting with Dialogflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type intentsGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing IntentsClient
	CallOptions **IntentsCallOptions

	// The gRPC API client.
	intentsClient dialogflowpb.IntentsClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewIntentsClient creates a new intents client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service for managing Intents.
func NewIntentsClient(ctx context.Context, opts ...option.ClientOption) (*IntentsClient, error) {
	clientOpts := defaultIntentsGRPCClientOptions()
	if newIntentsClientHook != nil {
		hookOpts, err := newIntentsClientHook(ctx, clientHookParams{})
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
	client := IntentsClient{CallOptions: defaultIntentsCallOptions()}

	c := &intentsGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		intentsClient:    dialogflowpb.NewIntentsClient(connPool),
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
func (c *intentsGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *intentsGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *intentsGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *intentsGRPCClient) ListIntents(ctx context.Context, req *dialogflowpb.ListIntentsRequest, opts ...gax.CallOption) *IntentIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListIntents[0:len((*c.CallOptions).ListIntents):len((*c.CallOptions).ListIntents)], opts...)
	it := &IntentIterator{}
	req = proto.Clone(req).(*dialogflowpb.ListIntentsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dialogflowpb.Intent, string, error) {
		var resp *dialogflowpb.ListIntentsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.intentsClient.ListIntents(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetIntents(), resp.GetNextPageToken(), nil
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

func (c *intentsGRPCClient) GetIntent(ctx context.Context, req *dialogflowpb.GetIntentRequest, opts ...gax.CallOption) (*dialogflowpb.Intent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIntent[0:len((*c.CallOptions).GetIntent):len((*c.CallOptions).GetIntent)], opts...)
	var resp *dialogflowpb.Intent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.intentsClient.GetIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *intentsGRPCClient) CreateIntent(ctx context.Context, req *dialogflowpb.CreateIntentRequest, opts ...gax.CallOption) (*dialogflowpb.Intent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateIntent[0:len((*c.CallOptions).CreateIntent):len((*c.CallOptions).CreateIntent)], opts...)
	var resp *dialogflowpb.Intent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.intentsClient.CreateIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *intentsGRPCClient) UpdateIntent(ctx context.Context, req *dialogflowpb.UpdateIntentRequest, opts ...gax.CallOption) (*dialogflowpb.Intent, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "intent.name", url.QueryEscape(req.GetIntent().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateIntent[0:len((*c.CallOptions).UpdateIntent):len((*c.CallOptions).UpdateIntent)], opts...)
	var resp *dialogflowpb.Intent
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.intentsClient.UpdateIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *intentsGRPCClient) DeleteIntent(ctx context.Context, req *dialogflowpb.DeleteIntentRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteIntent[0:len((*c.CallOptions).DeleteIntent):len((*c.CallOptions).DeleteIntent)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.intentsClient.DeleteIntent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *intentsGRPCClient) BatchUpdateIntents(ctx context.Context, req *dialogflowpb.BatchUpdateIntentsRequest, opts ...gax.CallOption) (*BatchUpdateIntentsOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchUpdateIntents[0:len((*c.CallOptions).BatchUpdateIntents):len((*c.CallOptions).BatchUpdateIntents)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.intentsClient.BatchUpdateIntents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchUpdateIntentsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *intentsGRPCClient) BatchDeleteIntents(ctx context.Context, req *dialogflowpb.BatchDeleteIntentsRequest, opts ...gax.CallOption) (*BatchDeleteIntentsOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchDeleteIntents[0:len((*c.CallOptions).BatchDeleteIntents):len((*c.CallOptions).BatchDeleteIntents)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.intentsClient.BatchDeleteIntents(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchDeleteIntentsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// BatchDeleteIntentsOperation manages a long-running operation from BatchDeleteIntents.
type BatchDeleteIntentsOperation struct {
	lro *longrunning.Operation
}

// BatchDeleteIntentsOperation returns a new BatchDeleteIntentsOperation from a given name.
// The name must be that of a previously created BatchDeleteIntentsOperation, possibly from a different process.
func (c *intentsGRPCClient) BatchDeleteIntentsOperation(name string) *BatchDeleteIntentsOperation {
	return &BatchDeleteIntentsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchDeleteIntentsOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
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
func (op *BatchDeleteIntentsOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *BatchDeleteIntentsOperation) Metadata() (*structpb.Struct, error) {
	var meta structpb.Struct
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchDeleteIntentsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchDeleteIntentsOperation) Name() string {
	return op.lro.Name()
}

// BatchUpdateIntentsOperation manages a long-running operation from BatchUpdateIntents.
type BatchUpdateIntentsOperation struct {
	lro *longrunning.Operation
}

// BatchUpdateIntentsOperation returns a new BatchUpdateIntentsOperation from a given name.
// The name must be that of a previously created BatchUpdateIntentsOperation, possibly from a different process.
func (c *intentsGRPCClient) BatchUpdateIntentsOperation(name string) *BatchUpdateIntentsOperation {
	return &BatchUpdateIntentsOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchUpdateIntentsOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*dialogflowpb.BatchUpdateIntentsResponse, error) {
	var resp dialogflowpb.BatchUpdateIntentsResponse
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
func (op *BatchUpdateIntentsOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*dialogflowpb.BatchUpdateIntentsResponse, error) {
	var resp dialogflowpb.BatchUpdateIntentsResponse
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
func (op *BatchUpdateIntentsOperation) Metadata() (*structpb.Struct, error) {
	var meta structpb.Struct
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchUpdateIntentsOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchUpdateIntentsOperation) Name() string {
	return op.lro.Name()
}

// IntentIterator manages a stream of *dialogflowpb.Intent.
type IntentIterator struct {
	items    []*dialogflowpb.Intent
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
	InternalFetch func(pageSize int, pageToken string) (results []*dialogflowpb.Intent, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *IntentIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *IntentIterator) Next() (*dialogflowpb.Intent, error) {
	var item *dialogflowpb.Intent
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *IntentIterator) bufLen() int {
	return len(it.items)
}

func (it *IntentIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
