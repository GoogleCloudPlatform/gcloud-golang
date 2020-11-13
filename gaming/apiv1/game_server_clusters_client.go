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
	gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newGameServerClustersClientHook clientHook

// GameServerClustersCallOptions contains the retry settings for each method of GameServerClustersClient.
type GameServerClustersCallOptions struct {
	ListGameServerClusters         []gax.CallOption
	GetGameServerCluster           []gax.CallOption
	CreateGameServerCluster        []gax.CallOption
	PreviewCreateGameServerCluster []gax.CallOption
	DeleteGameServerCluster        []gax.CallOption
	PreviewDeleteGameServerCluster []gax.CallOption
	UpdateGameServerCluster        []gax.CallOption
	PreviewUpdateGameServerCluster []gax.CallOption
}

func defaultGameServerClustersClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("gameservices.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("gameservices.mtls.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultGameServerClustersCallOptions() *GameServerClustersCallOptions {
	return &GameServerClustersCallOptions{
		ListGameServerClusters: []gax.CallOption{
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
		GetGameServerCluster: []gax.CallOption{
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
		CreateGameServerCluster: []gax.CallOption{},
		PreviewCreateGameServerCluster: []gax.CallOption{
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
		DeleteGameServerCluster: []gax.CallOption{},
		PreviewDeleteGameServerCluster: []gax.CallOption{
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
		UpdateGameServerCluster: []gax.CallOption{},
		PreviewUpdateGameServerCluster: []gax.CallOption{
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
	}
}

// GameServerClustersClient is a client for interacting with .
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type GameServerClustersClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	gameServerClustersClient gamingpb.GameServerClustersServiceClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *GameServerClustersCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewGameServerClustersClient creates a new game server clusters service client.
//
// The game server cluster maps to Kubernetes clusters running Agones and is
// used to manage fleets within clusters.
func NewGameServerClustersClient(ctx context.Context, opts ...option.ClientOption) (*GameServerClustersClient, error) {
	clientOpts := defaultGameServerClustersClientOptions()

	if newGameServerClustersClientHook != nil {
		hookOpts, err := newGameServerClustersClientHook(ctx, clientHookParams{})
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
	c := &GameServerClustersClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultGameServerClustersCallOptions(),

		gameServerClustersClient: gamingpb.NewGameServerClustersServiceClient(connPool),
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
func (c *GameServerClustersClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GameServerClustersClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GameServerClustersClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListGameServerClusters lists game server clusters in a given project and location.
func (c *GameServerClustersClient) ListGameServerClusters(ctx context.Context, req *gamingpb.ListGameServerClustersRequest, opts ...gax.CallOption) *GameServerClusterIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListGameServerClusters[0:len(c.CallOptions.ListGameServerClusters):len(c.CallOptions.ListGameServerClusters)], opts...)
	it := &GameServerClusterIterator{}
	req = proto.Clone(req).(*gamingpb.ListGameServerClustersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*gamingpb.GameServerCluster, string, error) {
		var resp *gamingpb.ListGameServerClustersResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.gameServerClustersClient.ListGameServerClusters(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetGameServerClusters(), resp.GetNextPageToken(), nil
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

// GetGameServerCluster gets details of a single game server cluster.
func (c *GameServerClustersClient) GetGameServerCluster(ctx context.Context, req *gamingpb.GetGameServerClusterRequest, opts ...gax.CallOption) (*gamingpb.GameServerCluster, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetGameServerCluster[0:len(c.CallOptions.GetGameServerCluster):len(c.CallOptions.GetGameServerCluster)], opts...)
	var resp *gamingpb.GameServerCluster
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerClustersClient.GetGameServerCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateGameServerCluster creates a new game server cluster in a given project and location.
func (c *GameServerClustersClient) CreateGameServerCluster(ctx context.Context, req *gamingpb.CreateGameServerClusterRequest, opts ...gax.CallOption) (*CreateGameServerClusterOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 120000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateGameServerCluster[0:len(c.CallOptions.CreateGameServerCluster):len(c.CallOptions.CreateGameServerCluster)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerClustersClient.CreateGameServerCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateGameServerClusterOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// PreviewCreateGameServerCluster previews creation of a new game server cluster in a given project and
// location.
func (c *GameServerClustersClient) PreviewCreateGameServerCluster(ctx context.Context, req *gamingpb.PreviewCreateGameServerClusterRequest, opts ...gax.CallOption) (*gamingpb.PreviewCreateGameServerClusterResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.PreviewCreateGameServerCluster[0:len(c.CallOptions.PreviewCreateGameServerCluster):len(c.CallOptions.PreviewCreateGameServerCluster)], opts...)
	var resp *gamingpb.PreviewCreateGameServerClusterResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerClustersClient.PreviewCreateGameServerCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteGameServerCluster deletes a single game server cluster.
func (c *GameServerClustersClient) DeleteGameServerCluster(ctx context.Context, req *gamingpb.DeleteGameServerClusterRequest, opts ...gax.CallOption) (*DeleteGameServerClusterOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.DeleteGameServerCluster[0:len(c.CallOptions.DeleteGameServerCluster):len(c.CallOptions.DeleteGameServerCluster)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerClustersClient.DeleteGameServerCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteGameServerClusterOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// PreviewDeleteGameServerCluster previews deletion of a single game server cluster.
func (c *GameServerClustersClient) PreviewDeleteGameServerCluster(ctx context.Context, req *gamingpb.PreviewDeleteGameServerClusterRequest, opts ...gax.CallOption) (*gamingpb.PreviewDeleteGameServerClusterResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.PreviewDeleteGameServerCluster[0:len(c.CallOptions.PreviewDeleteGameServerCluster):len(c.CallOptions.PreviewDeleteGameServerCluster)], opts...)
	var resp *gamingpb.PreviewDeleteGameServerClusterResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerClustersClient.PreviewDeleteGameServerCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateGameServerCluster patches a single game server cluster.
func (c *GameServerClustersClient) UpdateGameServerCluster(ctx context.Context, req *gamingpb.UpdateGameServerClusterRequest, opts ...gax.CallOption) (*UpdateGameServerClusterOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "game_server_cluster.name", url.QueryEscape(req.GetGameServerCluster().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateGameServerCluster[0:len(c.CallOptions.UpdateGameServerCluster):len(c.CallOptions.UpdateGameServerCluster)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerClustersClient.UpdateGameServerCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateGameServerClusterOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// PreviewUpdateGameServerCluster previews updating a GameServerCluster.
func (c *GameServerClustersClient) PreviewUpdateGameServerCluster(ctx context.Context, req *gamingpb.PreviewUpdateGameServerClusterRequest, opts ...gax.CallOption) (*gamingpb.PreviewUpdateGameServerClusterResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "game_server_cluster.name", url.QueryEscape(req.GetGameServerCluster().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.PreviewUpdateGameServerCluster[0:len(c.CallOptions.PreviewUpdateGameServerCluster):len(c.CallOptions.PreviewUpdateGameServerCluster)], opts...)
	var resp *gamingpb.PreviewUpdateGameServerClusterResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.gameServerClustersClient.PreviewUpdateGameServerCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateGameServerClusterOperation manages a long-running operation from CreateGameServerCluster.
type CreateGameServerClusterOperation struct {
	lro *longrunning.Operation
}

// CreateGameServerClusterOperation returns a new CreateGameServerClusterOperation from a given name.
// The name must be that of a previously created CreateGameServerClusterOperation, possibly from a different process.
func (c *GameServerClustersClient) CreateGameServerClusterOperation(name string) *CreateGameServerClusterOperation {
	return &CreateGameServerClusterOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateGameServerClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*gamingpb.GameServerCluster, error) {
	var resp gamingpb.GameServerCluster
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
func (op *CreateGameServerClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*gamingpb.GameServerCluster, error) {
	var resp gamingpb.GameServerCluster
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
func (op *CreateGameServerClusterOperation) Metadata() (*gamingpb.OperationMetadata, error) {
	var meta gamingpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateGameServerClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateGameServerClusterOperation) Name() string {
	return op.lro.Name()
}

// DeleteGameServerClusterOperation manages a long-running operation from DeleteGameServerCluster.
type DeleteGameServerClusterOperation struct {
	lro *longrunning.Operation
}

// DeleteGameServerClusterOperation returns a new DeleteGameServerClusterOperation from a given name.
// The name must be that of a previously created DeleteGameServerClusterOperation, possibly from a different process.
func (c *GameServerClustersClient) DeleteGameServerClusterOperation(name string) *DeleteGameServerClusterOperation {
	return &DeleteGameServerClusterOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteGameServerClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
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
func (op *DeleteGameServerClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteGameServerClusterOperation) Metadata() (*gamingpb.OperationMetadata, error) {
	var meta gamingpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteGameServerClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteGameServerClusterOperation) Name() string {
	return op.lro.Name()
}

// UpdateGameServerClusterOperation manages a long-running operation from UpdateGameServerCluster.
type UpdateGameServerClusterOperation struct {
	lro *longrunning.Operation
}

// UpdateGameServerClusterOperation returns a new UpdateGameServerClusterOperation from a given name.
// The name must be that of a previously created UpdateGameServerClusterOperation, possibly from a different process.
func (c *GameServerClustersClient) UpdateGameServerClusterOperation(name string) *UpdateGameServerClusterOperation {
	return &UpdateGameServerClusterOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateGameServerClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*gamingpb.GameServerCluster, error) {
	var resp gamingpb.GameServerCluster
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
func (op *UpdateGameServerClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*gamingpb.GameServerCluster, error) {
	var resp gamingpb.GameServerCluster
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
func (op *UpdateGameServerClusterOperation) Metadata() (*gamingpb.OperationMetadata, error) {
	var meta gamingpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateGameServerClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateGameServerClusterOperation) Name() string {
	return op.lro.Name()
}

// GameServerClusterIterator manages a stream of *gamingpb.GameServerCluster.
type GameServerClusterIterator struct {
	items    []*gamingpb.GameServerCluster
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
	InternalFetch func(pageSize int, pageToken string) (results []*gamingpb.GameServerCluster, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *GameServerClusterIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *GameServerClusterIterator) Next() (*gamingpb.GameServerCluster, error) {
	var item *gamingpb.GameServerCluster
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *GameServerClusterIterator) bufLen() int {
	return len(it.items)
}

func (it *GameServerClusterIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
