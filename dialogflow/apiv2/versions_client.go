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

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
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
	}
}

// internalVersionsClient is an interface that defines the methods availaible from Dialogflow API.
type internalVersionsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListVersions(context.Context, *dialogflowpb.ListVersionsRequest, ...gax.CallOption) *VersionIterator
	GetVersion(context.Context, *dialogflowpb.GetVersionRequest, ...gax.CallOption) (*dialogflowpb.Version, error)
	CreateVersion(context.Context, *dialogflowpb.CreateVersionRequest, ...gax.CallOption) (*dialogflowpb.Version, error)
	UpdateVersion(context.Context, *dialogflowpb.UpdateVersionRequest, ...gax.CallOption) (*dialogflowpb.Version, error)
	DeleteVersion(context.Context, *dialogflowpb.DeleteVersionRequest, ...gax.CallOption) error
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
func (c *VersionsClient) setGoogleClientInfo(...string) {
	c.internalClient.setGoogleClientInfo()
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *VersionsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListVersions returns the list of all versions of the specified agent.
func (c *VersionsClient) ListVersions(ctx context.Context, req *dialogflowpb.ListVersionsRequest, opts ...gax.CallOption) *VersionIterator {
	return c.internalClient.ListVersions(ctx, req, opts...)
}

// GetVersion retrieves the specified agent version.
func (c *VersionsClient) GetVersion(ctx context.Context, req *dialogflowpb.GetVersionRequest, opts ...gax.CallOption) (*dialogflowpb.Version, error) {
	return c.internalClient.GetVersion(ctx, req, opts...)
}

// CreateVersion creates an agent version.
//
// The new version points to the agent instance in the “default” environment.
func (c *VersionsClient) CreateVersion(ctx context.Context, req *dialogflowpb.CreateVersionRequest, opts ...gax.CallOption) (*dialogflowpb.Version, error) {
	return c.internalClient.CreateVersion(ctx, req, opts...)
}

// UpdateVersion updates the specified agent version.
//
// Note that this method does not allow you to update the state of the agent
// the given version points to. It allows you to update only mutable
// properties of the version resource.
func (c *VersionsClient) UpdateVersion(ctx context.Context, req *dialogflowpb.UpdateVersionRequest, opts ...gax.CallOption) (*dialogflowpb.Version, error) {
	return c.internalClient.UpdateVersion(ctx, req, opts...)
}

// DeleteVersion delete the specified agent version.
func (c *VersionsClient) DeleteVersion(ctx context.Context, req *dialogflowpb.DeleteVersionRequest, opts ...gax.CallOption) error {
	return c.internalClient.DeleteVersion(ctx, req, opts...)
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
	versionsClient dialogflowpb.VersionsClient

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
		versionsClient:   dialogflowpb.NewVersionsClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

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

func (c *versionsGRPCClient) ListVersions(ctx context.Context, req *dialogflowpb.ListVersionsRequest, opts ...gax.CallOption) *VersionIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListVersions[0:len((*c.CallOptions).ListVersions):len((*c.CallOptions).ListVersions)], opts...)
	it := &VersionIterator{}
	req = proto.Clone(req).(*dialogflowpb.ListVersionsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*dialogflowpb.Version, string, error) {
		var resp *dialogflowpb.ListVersionsResponse
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

func (c *versionsGRPCClient) GetVersion(ctx context.Context, req *dialogflowpb.GetVersionRequest, opts ...gax.CallOption) (*dialogflowpb.Version, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetVersion[0:len((*c.CallOptions).GetVersion):len((*c.CallOptions).GetVersion)], opts...)
	var resp *dialogflowpb.Version
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

func (c *versionsGRPCClient) CreateVersion(ctx context.Context, req *dialogflowpb.CreateVersionRequest, opts ...gax.CallOption) (*dialogflowpb.Version, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateVersion[0:len((*c.CallOptions).CreateVersion):len((*c.CallOptions).CreateVersion)], opts...)
	var resp *dialogflowpb.Version
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.versionsClient.CreateVersion(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *versionsGRPCClient) UpdateVersion(ctx context.Context, req *dialogflowpb.UpdateVersionRequest, opts ...gax.CallOption) (*dialogflowpb.Version, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "version.name", url.QueryEscape(req.GetVersion().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateVersion[0:len((*c.CallOptions).UpdateVersion):len((*c.CallOptions).UpdateVersion)], opts...)
	var resp *dialogflowpb.Version
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

func (c *versionsGRPCClient) DeleteVersion(ctx context.Context, req *dialogflowpb.DeleteVersionRequest, opts ...gax.CallOption) error {
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

// VersionIterator manages a stream of *dialogflowpb.Version.
type VersionIterator struct {
	items    []*dialogflowpb.Version
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
	InternalFetch func(pageSize int, pageToken string) (results []*dialogflowpb.Version, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *VersionIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *VersionIterator) Next() (*dialogflowpb.Version, error) {
	var item *dialogflowpb.Version
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
