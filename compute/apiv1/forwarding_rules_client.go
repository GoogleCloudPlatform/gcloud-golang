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

var newForwardingRulesClientHook clientHook

// ForwardingRulesCallOptions contains the retry settings for each method of ForwardingRulesClient.
type ForwardingRulesCallOptions struct {
	AggregatedList []gax.CallOption
	Delete         []gax.CallOption
	Get            []gax.CallOption
	Insert         []gax.CallOption
	List           []gax.CallOption
	Patch          []gax.CallOption
	SetLabels      []gax.CallOption
	SetTarget      []gax.CallOption
}

// internalForwardingRulesClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalForwardingRulesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AggregatedList(context.Context, *computepb.AggregatedListForwardingRulesRequest, ...gax.CallOption) (*computepb.ForwardingRuleAggregatedList, error)
	Delete(context.Context, *computepb.DeleteForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetForwardingRuleRequest, ...gax.CallOption) (*computepb.ForwardingRule, error)
	Insert(context.Context, *computepb.InsertForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListForwardingRulesRequest, ...gax.CallOption) (*computepb.ForwardingRuleList, error)
	Patch(context.Context, *computepb.PatchForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetLabels(context.Context, *computepb.SetLabelsForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetTarget(context.Context, *computepb.SetTargetForwardingRuleRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// ForwardingRulesClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The ForwardingRules API.
type ForwardingRulesClient struct {
	// The internal transport-dependent client.
	internalClient internalForwardingRulesClient

	// The call options for this service.
	CallOptions *ForwardingRulesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ForwardingRulesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ForwardingRulesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ForwardingRulesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AggregatedList retrieves an aggregated list of forwarding rules.
func (c *ForwardingRulesClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListForwardingRulesRequest, opts ...gax.CallOption) (*computepb.ForwardingRuleAggregatedList, error) {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified ForwardingRule resource.
func (c *ForwardingRulesClient) Delete(ctx context.Context, req *computepb.DeleteForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified ForwardingRule resource.
func (c *ForwardingRulesClient) Get(ctx context.Context, req *computepb.GetForwardingRuleRequest, opts ...gax.CallOption) (*computepb.ForwardingRule, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Insert creates a ForwardingRule resource in the specified project and region using the data included in the request.
func (c *ForwardingRulesClient) Insert(ctx context.Context, req *computepb.InsertForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of ForwardingRule resources available to the specified project and region.
func (c *ForwardingRulesClient) List(ctx context.Context, req *computepb.ListForwardingRulesRequest, opts ...gax.CallOption) (*computepb.ForwardingRuleList, error) {
	return c.internalClient.List(ctx, req, opts...)
}

// Patch updates the specified forwarding rule with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules. Currently, you can only patch the network_tier field.
func (c *ForwardingRulesClient) Patch(ctx context.Context, req *computepb.PatchForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Patch(ctx, req, opts...)
}

// SetLabels sets the labels on the specified resource. To learn more about labels, read the Labeling Resources documentation.
func (c *ForwardingRulesClient) SetLabels(ctx context.Context, req *computepb.SetLabelsForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetLabels(ctx, req, opts...)
}

// SetTarget changes target URL for forwarding rule. The new target should be of the same type as the old target.
func (c *ForwardingRulesClient) SetTarget(ctx context.Context, req *computepb.SetTargetForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetTarget(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type forwardingRulesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewForwardingRulesRESTClient creates a new forwarding rules rest client.
//
// The ForwardingRules API.
func NewForwardingRulesRESTClient(ctx context.Context, opts ...option.ClientOption) (*ForwardingRulesClient, error) {
	clientOpts := append(defaultForwardingRulesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &forwardingRulesRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &ForwardingRulesClient{internalClient: c, CallOptions: &ForwardingRulesCallOptions{}}, nil
}

func defaultForwardingRulesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *forwardingRulesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *forwardingRulesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *forwardingRulesRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AggregatedList retrieves an aggregated list of forwarding rules.
func (c *forwardingRulesRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListForwardingRulesRequest, opts ...gax.CallOption) (*computepb.ForwardingRuleAggregatedList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/forwardingRules", req.GetProject())

	params := url.Values{}
	if req != nil && req.Filter != nil {
		params.Add("filter", fmt.Sprintf("%v", req.GetFilter()))
	}
	if req != nil && req.IncludeAllScopes != nil {
		params.Add("includeAllScopes", fmt.Sprintf("%v", req.GetIncludeAllScopes()))
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

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
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
	rsp := &computepb.ForwardingRuleAggregatedList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Delete deletes the specified ForwardingRule resource.
func (c *forwardingRulesRESTClient) Delete(ctx context.Context, req *computepb.DeleteForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/forwardingRules/%v", req.GetProject(), req.GetRegion(), req.GetForwardingRule())

	params := url.Values{}
	if req != nil && req.RequestId != nil {
		params.Add("requestId", fmt.Sprintf("%v", req.GetRequestId()))
	}

	baseUrl.RawQuery = params.Encode()

	httpReq, err := http.NewRequest("DELETE", baseUrl.String(), bytes.NewReader(jsonReq))
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

// Get returns the specified ForwardingRule resource.
func (c *forwardingRulesRESTClient) Get(ctx context.Context, req *computepb.GetForwardingRuleRequest, opts ...gax.CallOption) (*computepb.ForwardingRule, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/forwardingRules/%v", req.GetProject(), req.GetRegion(), req.GetForwardingRule())

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
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
	rsp := &computepb.ForwardingRule{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a ForwardingRule resource in the specified project and region using the data included in the request.
func (c *forwardingRulesRESTClient) Insert(ctx context.Context, req *computepb.InsertForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetForwardingRuleResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/forwardingRules", req.GetProject(), req.GetRegion())

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

// List retrieves a list of ForwardingRule resources available to the specified project and region.
func (c *forwardingRulesRESTClient) List(ctx context.Context, req *computepb.ListForwardingRulesRequest, opts ...gax.CallOption) (*computepb.ForwardingRuleList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/forwardingRules", req.GetProject(), req.GetRegion())

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

	httpReq, err := http.NewRequest("GET", baseUrl.String(), bytes.NewReader(jsonReq))
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
	rsp := &computepb.ForwardingRuleList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Patch updates the specified forwarding rule with the data included in the request. This method supports PATCH semantics and uses the JSON merge patch format and processing rules. Currently, you can only patch the network_tier field.
func (c *forwardingRulesRESTClient) Patch(ctx context.Context, req *computepb.PatchForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetForwardingRuleResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/forwardingRules/%v", req.GetProject(), req.GetRegion(), req.GetForwardingRule())

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

// SetLabels sets the labels on the specified resource. To learn more about labels, read the Labeling Resources documentation.
func (c *forwardingRulesRESTClient) SetLabels(ctx context.Context, req *computepb.SetLabelsForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetRegionSetLabelsRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/forwardingRules/%v/setLabels", req.GetProject(), req.GetRegion(), req.GetResource())

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

// SetTarget changes target URL for forwarding rule. The new target should be of the same type as the old target.
func (c *forwardingRulesRESTClient) SetTarget(ctx context.Context, req *computepb.SetTargetForwardingRuleRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTargetReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/forwardingRules/%v/setTarget", req.GetProject(), req.GetRegion(), req.GetForwardingRule())

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
