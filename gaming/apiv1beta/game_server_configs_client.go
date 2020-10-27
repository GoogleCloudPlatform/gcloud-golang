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

package gaming

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
	gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1beta"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newGameServerConfigsClientHook clientHook

// GameServerConfigsCallOptions contains the retry settings for each method of GameServerConfigsClient.
type GameServerConfigsCallOptions struct {
	ListGameServerConfigs  []gax.CallOption
	GetGameServerConfig    []gax.CallOption
	CreateGameServerConfig []gax.CallOption
	DeleteGameServerConfig []gax.CallOption
}

func defaultGameServerConfigsClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("gameservices.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("gameservices.mtls.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultGameServerConfigsCallOptions() *GameServerConfigsCallOptions {
	return &GameServerConfigsCallOptions{
		ListGameServerConfigs: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetGameServerConfig: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateGameServerConfig: []gax.CallOption{},
		DeleteGameServerConfig: []gax.CallOption{},
	}
}

// GameServerConfigsClient is a client for interacting with .
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type GameServerConfigsClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	gameServerConfigsClient gamingpb.GameServerConfigsServiceClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *GameServerConfigsCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewGameServerConfigsClient creates a new game server configs service client.
//
// The game server config configures the game servers in an Agones fleet.
func NewGameServerConfigsClient(ctx context.Context, opts ...option.ClientOption) (*GameServerConfigsClient, error) {
	clientOpts := defaultGameServerConfigsClientOptions()

	if newGameServerConfigsClientHook != nil {
		hookOpts, err := newGameServerConfigsClientHook(ctx, clientHookParams{})
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
	c := &GameServerConfigsClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultGameServerConfigsCallOptions(),

		gameServerConfigsClient: gamingpb.NewGameServerConfigsServiceClient(connPool),
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
func (c *GameServerConfigsClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GameServerConfigsClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GameServerConfigsClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListGameServerConfigs lists game server configs in a given project, location, and game server
// deployment.
func (c *GameServerConfigsClient) ListGameServerConfigs(ctx context.Context, req *gamingpb.ListGameServerConfigsRequest, opts ...gax.CallOption) *GameServerConfigIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListGameServerConfigs[0:len(c.CallOptions.ListGameServerConfigs):len(c.CallOptions.ListGameServerConfigs)], opts...)
	it := &GameServerConfigIterator{}
	req = proto.Clone(req).(*gamingpb.ListGameServerConfigsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*gamingpb.GameServerConfig, string, error) {
		var resp *gamingpb.ListGameServerConfigsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.gameServerConfigsClient.ListGameServerConfigs(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetGameServerConfigs(), resp.GetNextPageToken(), nil
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

// GetGameServerConfig gets details of a single game server config.
func (c *GameServerConfigsClient) GetGameServerConfig(ctx context.Context, req *gamingpb.GetGameServerConfigRequest, opts ...gax.CallOption) (*gamingpb.GameServerConfig, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetGameServerConfig[0:len(c.CallOptions.GetGameServerConfig):len(c.CallOptions.GetGameServerConfig)], opts...)
	var resp *gamingpb.GameServerConfig
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerConfigsClient.GetGameServerConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateGameServerConfig creates a new game server config in a given project, location, and game
// server deployment. Game server configs are immutable, and are not applied
// until referenced in the game server deployment rollout resource.
func (c *GameServerConfigsClient) CreateGameServerConfig(ctx context.Context, req *gamingpb.CreateGameServerConfigRequest, opts ...gax.CallOption) (*CreateGameServerConfigOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateGameServerConfig[0:len(c.CallOptions.CreateGameServerConfig):len(c.CallOptions.CreateGameServerConfig)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerConfigsClient.CreateGameServerConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateGameServerConfigOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// DeleteGameServerConfig deletes a single game server config. The deletion will fail if the game
// server config is referenced in a game server deployment rollout.
func (c *GameServerConfigsClient) DeleteGameServerConfig(ctx context.Context, req *gamingpb.DeleteGameServerConfigRequest, opts ...gax.CallOption) (*DeleteGameServerConfigOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteGameServerConfig[0:len(c.CallOptions.DeleteGameServerConfig):len(c.CallOptions.DeleteGameServerConfig)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerConfigsClient.DeleteGameServerConfig(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteGameServerConfigOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// CreateGameServerConfigOperation manages a long-running operation from CreateGameServerConfig.
type CreateGameServerConfigOperation struct {
	lro *longrunning.Operation
}

// CreateGameServerConfigOperation returns a new CreateGameServerConfigOperation from a given name.
// The name must be that of a previously created CreateGameServerConfigOperation, possibly from a different process.
func (c *GameServerConfigsClient) CreateGameServerConfigOperation(name string) *CreateGameServerConfigOperation {
	return &CreateGameServerConfigOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateGameServerConfigOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*gamingpb.GameServerConfig, error) {
	var resp gamingpb.GameServerConfig
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
func (op *CreateGameServerConfigOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*gamingpb.GameServerConfig, error) {
	var resp gamingpb.GameServerConfig
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
func (op *CreateGameServerConfigOperation) Metadata() (*gamingpb.OperationMetadata, error) {
	var meta gamingpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateGameServerConfigOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateGameServerConfigOperation) Name() string {
	return op.lro.Name()
}

// DeleteGameServerConfigOperation manages a long-running operation from DeleteGameServerConfig.
type DeleteGameServerConfigOperation struct {
	lro *longrunning.Operation
}

// DeleteGameServerConfigOperation returns a new DeleteGameServerConfigOperation from a given name.
// The name must be that of a previously created DeleteGameServerConfigOperation, possibly from a different process.
func (c *GameServerConfigsClient) DeleteGameServerConfigOperation(name string) *DeleteGameServerConfigOperation {
	return &DeleteGameServerConfigOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteGameServerConfigOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
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
func (op *DeleteGameServerConfigOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteGameServerConfigOperation) Metadata() (*gamingpb.OperationMetadata, error) {
	var meta gamingpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteGameServerConfigOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteGameServerConfigOperation) Name() string {
	return op.lro.Name()
}

// GameServerConfigIterator manages a stream of *gamingpb.GameServerConfig.
type GameServerConfigIterator struct {
	items    []*gamingpb.GameServerConfig
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
	InternalFetch func(pageSize int, pageToken string) (results []*gamingpb.GameServerConfig, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *GameServerConfigIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *GameServerConfigIterator) Next() (*gamingpb.GameServerConfig, error) {
	var item *gamingpb.GameServerConfig
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *GameServerConfigIterator) bufLen() int {
	return len(it.items)
}

func (it *GameServerConfigIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
