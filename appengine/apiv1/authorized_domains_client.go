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

package appengine

import (
	"context"
	"fmt"
	"math"
	"net/url"

	"github.com/golang/protobuf/proto"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newAuthorizedDomainsClientHook clientHook

// AuthorizedDomainsCallOptions contains the retry settings for each method of AuthorizedDomainsClient.
type AuthorizedDomainsCallOptions struct {
	ListAuthorizedDomains []gax.CallOption
}

func defaultAuthorizedDomainsClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("appengine.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("appengine.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://appengine.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAuthorizedDomainsCallOptions() *AuthorizedDomainsCallOptions {
	return &AuthorizedDomainsCallOptions{
		ListAuthorizedDomains: []gax.CallOption{},
	}
}

// AuthorizedDomainsClient is a client for interacting with App Engine Audit Data.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type AuthorizedDomainsClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	authorizedDomainsClient appenginepb.AuthorizedDomainsClient

	// The call options for this service.
	CallOptions *AuthorizedDomainsCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAuthorizedDomainsClient creates a new authorized domains client.
//
// Manages domains a user is authorized to administer. To authorize use of a
// domain, verify ownership via
// Webmaster Central (at https://www.google.com/webmasters/verification/home).
func NewAuthorizedDomainsClient(ctx context.Context, opts ...option.ClientOption) (*AuthorizedDomainsClient, error) {
	clientOpts := defaultAuthorizedDomainsClientOptions()

	if newAuthorizedDomainsClientHook != nil {
		hookOpts, err := newAuthorizedDomainsClientHook(ctx, clientHookParams{})
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
	c := &AuthorizedDomainsClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultAuthorizedDomainsCallOptions(),

		authorizedDomainsClient: appenginepb.NewAuthorizedDomainsClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AuthorizedDomainsClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AuthorizedDomainsClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AuthorizedDomainsClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListAuthorizedDomains lists all domains the user is authorized to administer.
func (c *AuthorizedDomainsClient) ListAuthorizedDomains(ctx context.Context, req *appenginepb.ListAuthorizedDomainsRequest, opts ...gax.CallOption) *AuthorizedDomainIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListAuthorizedDomains[0:len(c.CallOptions.ListAuthorizedDomains):len(c.CallOptions.ListAuthorizedDomains)], opts...)
	it := &AuthorizedDomainIterator{}
	req = proto.Clone(req).(*appenginepb.ListAuthorizedDomainsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*appenginepb.AuthorizedDomain, string, error) {
		var resp *appenginepb.ListAuthorizedDomainsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.authorizedDomainsClient.ListAuthorizedDomains(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetDomains(), resp.GetNextPageToken(), nil
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

// AuthorizedDomainIterator manages a stream of *appenginepb.AuthorizedDomain.
type AuthorizedDomainIterator struct {
	items    []*appenginepb.AuthorizedDomain
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
	InternalFetch func(pageSize int, pageToken string) (results []*appenginepb.AuthorizedDomain, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *AuthorizedDomainIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *AuthorizedDomainIterator) Next() (*appenginepb.AuthorizedDomain, error) {
	var item *appenginepb.AuthorizedDomain
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *AuthorizedDomainIterator) bufLen() int {
	return len(it.items)
}

func (it *AuthorizedDomainIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
