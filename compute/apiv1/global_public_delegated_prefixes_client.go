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

package compute

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newGlobalPublicDelegatedPrefixesClientHook clientHook

// GlobalPublicDelegatedPrefixesCallOptions contains the retry settings for each method of GlobalPublicDelegatedPrefixesClient.
type GlobalPublicDelegatedPrefixesCallOptions struct {
	Delete []gax.CallOption
	Get    []gax.CallOption
	Insert []gax.CallOption
	List   []gax.CallOption
	Patch  []gax.CallOption
}

// internalGlobalPublicDelegatedPrefixesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalGlobalPublicDelegatedPrefixesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Delete(context.Context, *computepb.DeleteGlobalPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetGlobalPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error)
	Insert(context.Context, *computepb.InsertGlobalPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListGlobalPublicDelegatedPrefixesRequest, ...gax.CallOption) (*computepb.PublicDelegatedPrefixList, error)
	Patch(context.Context, *computepb.PatchGlobalPublicDelegatedPrefixeRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// GlobalPublicDelegatedPrefixesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The GlobalPublicDelegatedPrefixes API.
type GlobalPublicDelegatedPrefixesClient struct {
	// The internal transport-dependent client.
	internalClient internalGlobalPublicDelegatedPrefixesClient

	// The call options for this service.
	CallOptions *GlobalPublicDelegatedPrefixesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *GlobalPublicDelegatedPrefixesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *GlobalPublicDelegatedPrefixesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *GlobalPublicDelegatedPrefixesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Delete deletes the specified global PublicDelegatedPrefix.
func (c *GlobalPublicDelegatedPrefixesClient) Delete(ctx context.Context, req *computepb.DeleteGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified global PublicDelegatedPrefix resource.
func (c *GlobalPublicDelegatedPrefixesClient) Get(ctx context.Context, req *computepb.GetGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a global PublicDelegatedPrefix in the specified project using the parameters that are included in the request.
func (c *GlobalPublicDelegatedPrefixesClient) Insert(ctx context.Context, req *computepb.InsertGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List lists the global PublicDelegatedPrefixes for a project.
func (c *GlobalPublicDelegatedPrefixesClient) List(ctx context.Context, req *computepb.ListGlobalPublicDelegatedPrefixesRequest, opts ...gax.CallOption) (*computepb.PublicDelegatedPrefixList, error) {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch patches the specified global PublicDelegatedPrefix resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *GlobalPublicDelegatedPrefixesClient) Patch(ctx context.Context, req *computepb.PatchGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type globalPublicDelegatedPrefixesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewGlobalPublicDelegatedPrefixesRESTClient creates a new global public delegated prefixes rest client.
//
// The GlobalPublicDelegatedPrefixes API.
func NewGlobalPublicDelegatedPrefixesRESTClient(ctx context.Context, opts ...option.ClientOption) (*GlobalPublicDelegatedPrefixesClient, error) {
	clientOpts := append(defaultGlobalPublicDelegatedPrefixesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &globalPublicDelegatedPrefixesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &GlobalPublicDelegatedPrefixesClient{internalClient: c, CallOptions: &GlobalPublicDelegatedPrefixesCallOptions{}}, nil
}

func defaultGlobalPublicDelegatedPrefixesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *globalPublicDelegatedPrefixesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *globalPublicDelegatedPrefixesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *globalPublicDelegatedPrefixesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Delete deletes the specified global PublicDelegatedPrefix.
func (c *globalPublicDelegatedPrefixesRESTClient) Delete(ctx context.Context, req *computepb.DeleteGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicDelegatedPrefixes/%v", req.GetProject(), req.GetPublicDelegatedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Get returns the specified global PublicDelegatedPrefix resource.
func (c *globalPublicDelegatedPrefixesRESTClient) Get(ctx context.Context, req *computepb.GetGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.PublicDelegatedPrefix, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicDelegatedPrefixes/%v", req.GetProject(), req.GetPublicDelegatedPrefix())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.PublicDelegatedPrefix{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a global PublicDelegatedPrefix in the specified project using the parameters that are included in the request.
func (c *globalPublicDelegatedPrefixesRESTClient) Insert(ctx context.Context, req *computepb.InsertGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetPublicDelegatedPrefixResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicDelegatedPrefixes", req.GetProject())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// List lists the global PublicDelegatedPrefixes for a project.
func (c *globalPublicDelegatedPrefixesRESTClient) List(ctx context.Context, req *computepb.ListGlobalPublicDelegatedPrefixesRequest, opts ...gax.CallOption) (*computepb.PublicDelegatedPrefixList, error) {
	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicDelegatedPrefixes", req.GetProject())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.MaxResults != nil {
		params.Add("maxResults", fmt.Sprintf("%v", req.GetMaxResults()))
	}
	if req != nil && req.OrderBy != nil {
		params.Add("orderBy", fmt.Sprintf("%v", req.GetOrderBy()))
	}
	if req != nil && req.PageToken != nil {
		params.Add("pageToken", fmt.Sprintf("%v", req.GetPageToken()))
	}
	if req != nil && req.ReturnPartialSuccess != nil {
		params.Add("returnPartialSuccess", fmt.Sprintf("%v", req.GetReturnPartialSuccess()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.PublicDelegatedPrefixList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Patch patches the specified global PublicDelegatedPrefix resource with the data included in the request. This method supports PATCH semantics and uses JSON merge patch format and processing rules.
func (c *globalPublicDelegatedPrefixesRESTClient) Patch(ctx context.Context, req *computepb.PatchGlobalPublicDelegatedPrefixeRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true}
	body := req.GetPublicDelegatedPrefixResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/global/publicDelegatedPrefixes/%v", req.GetProject(), req.GetPublicDelegatedPrefix())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("PATCH", baseUrl.String(), bytes.NewReader(jsonReq))
	if err != nil {
		return nil, err
	}
	httpReq = httpReq.WithContext(ctx)
	// Set the headers
	for k, v := range c.xGoogMetadata {
		httpReq.Header[k] = v
	}
	httpReq.Header["Content-Type"] = []string{"application/json"}

	httpRsp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer httpRsp.Body.Close()

	if httpRsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(httpRsp.Status)
	}

	buf, err := ioutil.ReadAll(httpRsp.Body)
	if err != nil {
		return nil, err
	}

	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	rsp := &computepb.Operation{}

	return rsp, unm.Unmarshal(buf, rsp)
}
