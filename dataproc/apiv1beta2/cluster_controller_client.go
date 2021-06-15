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

package dataproc

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
	dataprocpb "google.golang.org/genproto/googleapis/cloud/dataproc/v1beta2"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newClusterControllerClientHook clientHook

// ClusterControllerCallOptions contains the retry settings for each method of ClusterControllerClient.
type ClusterControllerCallOptions struct {
	CreateCluster   []gax.CallOption
	UpdateCluster   []gax.CallOption
	DeleteCluster   []gax.CallOption
	GetCluster      []gax.CallOption
	ListClusters    []gax.CallOption
	DiagnoseCluster []gax.CallOption
}

func defaultClusterControllerGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataproc.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataproc.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataproc.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultClusterControllerCallOptions() *ClusterControllerCallOptions {
	return &ClusterControllerCallOptions{
		CreateCluster: []gax.CallOption{
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
		UpdateCluster: []gax.CallOption{
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
		DeleteCluster: []gax.CallOption{
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
		GetCluster: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Internal,
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListClusters: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Internal,
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DiagnoseCluster: []gax.CallOption{
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

// internalClusterControllerClient is an interface that defines the methods availaible from Cloud Dataproc API.
type internalClusterControllerClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateCluster(context.Context, *dataprocpb.CreateClusterRequest, ...gax.CallOption) (*CreateClusterOperation, error)
	CreateClusterOperation(name string) *CreateClusterOperation
	UpdateCluster(context.Context, *dataprocpb.UpdateClusterRequest, ...gax.CallOption) (*UpdateClusterOperation, error)
	UpdateClusterOperation(name string) *UpdateClusterOperation
	DeleteCluster(context.Context, *dataprocpb.DeleteClusterRequest, ...gax.CallOption) (*DeleteClusterOperation, error)
	DeleteClusterOperation(name string) *DeleteClusterOperation
	GetCluster(context.Context, *dataprocpb.GetClusterRequest, ...gax.CallOption) (*dataprocpb.Cluster, error)
	ListClusters(context.Context, *dataprocpb.ListClustersRequest, ...gax.CallOption) *ClusterIterator
	DiagnoseCluster(context.Context, *dataprocpb.DiagnoseClusterRequest, ...gax.CallOption) (*DiagnoseClusterOperation, error)
	DiagnoseClusterOperation(name string) *DiagnoseClusterOperation
}

// ClusterControllerClient is a client for interacting with Cloud Dataproc API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The ClusterControllerService provides methods to manage clusters
// of Compute Engine instances.
type ClusterControllerClient struct {
	// The internal transport-dependent client.
	internalClient internalClusterControllerClient

	// The call options for this service.
	CallOptions *ClusterControllerCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ClusterControllerClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ClusterControllerClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ClusterControllerClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateCluster creates a cluster in a project. The returned
// Operation.metadata will be
// ClusterOperationMetadata (at https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#clusteroperationmetadata).
func (c *ClusterControllerClient) CreateCluster(ctx context.Context, req *dataprocpb.CreateClusterRequest, opts ...gax.CallOption) (*CreateClusterOperation, error) {
	return c.internalClient.CreateCluster(ctx, req, opts...)
}

// CreateClusterOperation returns a new CreateClusterOperation from a given name.
// The name must be that of a previously created CreateClusterOperation, possibly from a different process.
func (c *ClusterControllerClient) CreateClusterOperation(name string) *CreateClusterOperation {
	return c.internalClient.CreateClusterOperation(name)
}

// UpdateCluster updates a cluster in a project. The returned
// Operation.metadata will be
// ClusterOperationMetadata (at https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#clusteroperationmetadata).
func (c *ClusterControllerClient) UpdateCluster(ctx context.Context, req *dataprocpb.UpdateClusterRequest, opts ...gax.CallOption) (*UpdateClusterOperation, error) {
	return c.internalClient.UpdateCluster(ctx, req, opts...)
}

// UpdateClusterOperation returns a new UpdateClusterOperation from a given name.
// The name must be that of a previously created UpdateClusterOperation, possibly from a different process.
func (c *ClusterControllerClient) UpdateClusterOperation(name string) *UpdateClusterOperation {
	return c.internalClient.UpdateClusterOperation(name)
}

// DeleteCluster deletes a cluster in a project. The returned
// Operation.metadata will be
// ClusterOperationMetadata (at https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#clusteroperationmetadata).
func (c *ClusterControllerClient) DeleteCluster(ctx context.Context, req *dataprocpb.DeleteClusterRequest, opts ...gax.CallOption) (*DeleteClusterOperation, error) {
	return c.internalClient.DeleteCluster(ctx, req, opts...)
}

// DeleteClusterOperation returns a new DeleteClusterOperation from a given name.
// The name must be that of a previously created DeleteClusterOperation, possibly from a different process.
func (c *ClusterControllerClient) DeleteClusterOperation(name string) *DeleteClusterOperation {
	return c.internalClient.DeleteClusterOperation(name)
}

// GetCluster gets the resource representation for a cluster in a project.
func (c *ClusterControllerClient) GetCluster(ctx context.Context, req *dataprocpb.GetClusterRequest, opts ...gax.CallOption) (*dataprocpb.Cluster, error) {
	return c.internalClient.GetCluster(ctx, req, opts...)
}

// ListClusters lists all regions/{region}/clusters in a project alphabetically.
func (c *ClusterControllerClient) ListClusters(ctx context.Context, req *dataprocpb.ListClustersRequest, opts ...gax.CallOption) *ClusterIterator {
	return c.internalClient.ListClusters(ctx, req, opts...)
}

// DiagnoseCluster gets cluster diagnostic information. The returned
// Operation.metadata will be
// ClusterOperationMetadata (at https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#clusteroperationmetadata).
// After the operation completes,
// Operation.response
// contains
// Empty.
func (c *ClusterControllerClient) DiagnoseCluster(ctx context.Context, req *dataprocpb.DiagnoseClusterRequest, opts ...gax.CallOption) (*DiagnoseClusterOperation, error) {
	return c.internalClient.DiagnoseCluster(ctx, req, opts...)
}

// DiagnoseClusterOperation returns a new DiagnoseClusterOperation from a given name.
// The name must be that of a previously created DiagnoseClusterOperation, possibly from a different process.
func (c *ClusterControllerClient) DiagnoseClusterOperation(name string) *DiagnoseClusterOperation {
	return c.internalClient.DiagnoseClusterOperation(name)
}

// clusterControllerGRPCClient is a client for interacting with Cloud Dataproc API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type clusterControllerGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ClusterControllerClient
	CallOptions **ClusterControllerCallOptions

	// The gRPC API client.
	clusterControllerClient dataprocpb.ClusterControllerClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClusterControllerClient creates a new cluster controller client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The ClusterControllerService provides methods to manage clusters
// of Compute Engine instances.
func NewClusterControllerClient(ctx context.Context, opts ...option.ClientOption) (*ClusterControllerClient, error) {
	clientOpts := defaultClusterControllerGRPCClientOptions()
	if newClusterControllerClientHook != nil {
		hookOpts, err := newClusterControllerClientHook(ctx, clientHookParams{})
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
	client := ClusterControllerClient{CallOptions: defaultClusterControllerCallOptions()}

	c := &clusterControllerGRPCClient{
		connPool:                connPool,
		disableDeadlines:        disableDeadlines,
		clusterControllerClient: dataprocpb.NewClusterControllerClient(connPool),
		CallOptions:             &client.CallOptions,
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
func (c *clusterControllerGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *clusterControllerGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *clusterControllerGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *clusterControllerGRPCClient) CreateCluster(ctx context.Context, req *dataprocpb.CreateClusterRequest, opts ...gax.CallOption) (*CreateClusterOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 300000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateCluster[0:len((*c.CallOptions).CreateCluster):len((*c.CallOptions).CreateCluster)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.clusterControllerClient.CreateCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateClusterOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *clusterControllerGRPCClient) UpdateCluster(ctx context.Context, req *dataprocpb.UpdateClusterRequest, opts ...gax.CallOption) (*UpdateClusterOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 300000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion()), "cluster_name", url.QueryEscape(req.GetClusterName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateCluster[0:len((*c.CallOptions).UpdateCluster):len((*c.CallOptions).UpdateCluster)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.clusterControllerClient.UpdateCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateClusterOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *clusterControllerGRPCClient) DeleteCluster(ctx context.Context, req *dataprocpb.DeleteClusterRequest, opts ...gax.CallOption) (*DeleteClusterOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 300000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion()), "cluster_name", url.QueryEscape(req.GetClusterName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteCluster[0:len((*c.CallOptions).DeleteCluster):len((*c.CallOptions).DeleteCluster)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.clusterControllerClient.DeleteCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteClusterOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *clusterControllerGRPCClient) GetCluster(ctx context.Context, req *dataprocpb.GetClusterRequest, opts ...gax.CallOption) (*dataprocpb.Cluster, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 300000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion()), "cluster_name", url.QueryEscape(req.GetClusterName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCluster[0:len((*c.CallOptions).GetCluster):len((*c.CallOptions).GetCluster)], opts...)
	var resp *dataprocpb.Cluster
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.clusterControllerClient.GetCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *clusterControllerGRPCClient) ListClusters(ctx context.Context, req *dataprocpb.ListClustersRequest, opts ...gax.CallOption) *ClusterIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListClusters[0:len((*c.CallOptions).ListClusters):len((*c.CallOptions).ListClusters)], opts...)
	it := &ClusterIterator{}
	req = proto.Clone(req).(*dataprocpb.ListClustersRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dataprocpb.Cluster, string, error) {
		var resp *dataprocpb.ListClustersResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.clusterControllerClient.ListClusters(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetClusters(), resp.GetNextPageToken(), nil
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

func (c *clusterControllerGRPCClient) DiagnoseCluster(ctx context.Context, req *dataprocpb.DiagnoseClusterRequest, opts ...gax.CallOption) (*DiagnoseClusterOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 300000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "region", url.QueryEscape(req.GetRegion()), "cluster_name", url.QueryEscape(req.GetClusterName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DiagnoseCluster[0:len((*c.CallOptions).DiagnoseCluster):len((*c.CallOptions).DiagnoseCluster)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.clusterControllerClient.DiagnoseCluster(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DiagnoseClusterOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// CreateClusterOperation manages a long-running operation from CreateCluster.
type CreateClusterOperation struct {
	lro *longrunning.Operation
}

// CreateClusterOperation returns a new CreateClusterOperation from a given name.
// The name must be that of a previously created CreateClusterOperation, possibly from a different process.
func (c *clusterControllerGRPCClient) CreateClusterOperation(name string) *CreateClusterOperation {
	return &CreateClusterOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*dataprocpb.Cluster, error) {
	var resp dataprocpb.Cluster
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
func (op *CreateClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*dataprocpb.Cluster, error) {
	var resp dataprocpb.Cluster
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
func (op *CreateClusterOperation) Metadata() (*dataprocpb.ClusterOperationMetadata, error) {
	var meta dataprocpb.ClusterOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateClusterOperation) Name() string {
	return op.lro.Name()
}

// DeleteClusterOperation manages a long-running operation from DeleteCluster.
type DeleteClusterOperation struct {
	lro *longrunning.Operation
}

// DeleteClusterOperation returns a new DeleteClusterOperation from a given name.
// The name must be that of a previously created DeleteClusterOperation, possibly from a different process.
func (c *clusterControllerGRPCClient) DeleteClusterOperation(name string) *DeleteClusterOperation {
	return &DeleteClusterOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
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
func (op *DeleteClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteClusterOperation) Metadata() (*dataprocpb.ClusterOperationMetadata, error) {
	var meta dataprocpb.ClusterOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteClusterOperation) Name() string {
	return op.lro.Name()
}

// DiagnoseClusterOperation manages a long-running operation from DiagnoseCluster.
type DiagnoseClusterOperation struct {
	lro *longrunning.Operation
}

// DiagnoseClusterOperation returns a new DiagnoseClusterOperation from a given name.
// The name must be that of a previously created DiagnoseClusterOperation, possibly from a different process.
func (c *clusterControllerGRPCClient) DiagnoseClusterOperation(name string) *DiagnoseClusterOperation {
	return &DiagnoseClusterOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DiagnoseClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) error {
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
func (op *DiagnoseClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) error {
	return op.lro.Poll(ctx, nil, opts...)
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DiagnoseClusterOperation) Metadata() (*dataprocpb.ClusterOperationMetadata, error) {
	var meta dataprocpb.ClusterOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DiagnoseClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DiagnoseClusterOperation) Name() string {
	return op.lro.Name()
}

// UpdateClusterOperation manages a long-running operation from UpdateCluster.
type UpdateClusterOperation struct {
	lro *longrunning.Operation
}

// UpdateClusterOperation returns a new UpdateClusterOperation from a given name.
// The name must be that of a previously created UpdateClusterOperation, possibly from a different process.
func (c *clusterControllerGRPCClient) UpdateClusterOperation(name string) *UpdateClusterOperation {
	return &UpdateClusterOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateClusterOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*dataprocpb.Cluster, error) {
	var resp dataprocpb.Cluster
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
func (op *UpdateClusterOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*dataprocpb.Cluster, error) {
	var resp dataprocpb.Cluster
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
func (op *UpdateClusterOperation) Metadata() (*dataprocpb.ClusterOperationMetadata, error) {
	var meta dataprocpb.ClusterOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateClusterOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateClusterOperation) Name() string {
	return op.lro.Name()
}

// ClusterIterator manages a stream of *dataprocpb.Cluster.
type ClusterIterator struct {
	items    []*dataprocpb.Cluster
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
	InternalFetch func(pageSize int, pageToken string) (results []*dataprocpb.Cluster, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *ClusterIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *ClusterIterator) Next() (*dataprocpb.Cluster, error) {
	var item *dataprocpb.Cluster
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *ClusterIterator) bufLen() int {
	return len(it.items)
}

func (it *ClusterIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
