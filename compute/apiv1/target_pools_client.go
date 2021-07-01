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

var newTargetPoolsClientHook clientHook

// TargetPoolsCallOptions contains the retry settings for each method of TargetPoolsClient.
type TargetPoolsCallOptions struct {
	AddHealthCheck    []gax.CallOption
	AddInstance       []gax.CallOption
	AggregatedList    []gax.CallOption
	Delete            []gax.CallOption
	Get               []gax.CallOption
	GetHealth         []gax.CallOption
	Insert            []gax.CallOption
	List              []gax.CallOption
	RemoveHealthCheck []gax.CallOption
	RemoveInstance    []gax.CallOption
	SetBackup         []gax.CallOption
}

// internalTargetPoolsClient is an interface that defines the methods availaible from Google Compute Engine API.
type internalTargetPoolsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	AddHealthCheck(context.Context, *computepb.AddHealthCheckTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	AddInstance(context.Context, *computepb.AddInstanceTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	AggregatedList(context.Context, *computepb.AggregatedListTargetPoolsRequest, ...gax.CallOption) (*computepb.TargetPoolAggregatedList, error)
	Delete(context.Context, *computepb.DeleteTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	Get(context.Context, *computepb.GetTargetPoolRequest, ...gax.CallOption) (*computepb.TargetPool, error)
	GetHealth(context.Context, *computepb.GetHealthTargetPoolRequest, ...gax.CallOption) (*computepb.TargetPoolInstanceHealth, error)
	Insert(context.Context, *computepb.InsertTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	List(context.Context, *computepb.ListTargetPoolsRequest, ...gax.CallOption) (*computepb.TargetPoolList, error)
	RemoveHealthCheck(context.Context, *computepb.RemoveHealthCheckTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	RemoveInstance(context.Context, *computepb.RemoveInstanceTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
	SetBackup(context.Context, *computepb.SetBackupTargetPoolRequest, ...gax.CallOption) (*computepb.Operation, error)
}

// TargetPoolsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The TargetPools API.
type TargetPoolsClient struct {
	// The internal transport-dependent client.
	internalClient internalTargetPoolsClient

	// The call options for this service.
	CallOptions *TargetPoolsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TargetPoolsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TargetPoolsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *TargetPoolsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// AddHealthCheck adds health check URLs to a target pool.
func (c *TargetPoolsClient) AddHealthCheck(ctx context.Context, req *computepb.AddHealthCheckTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddHealthCheck(ctx, req, opts...)
}

// AddInstance adds an instance to a target pool.
func (c *TargetPoolsClient) AddInstance(ctx context.Context, req *computepb.AddInstanceTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.AddInstance(ctx, req, opts...)
}

// AggregatedList retrieves an aggregated list of target pools.
func (c *TargetPoolsClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetPoolsRequest, opts ...gax.CallOption) (*computepb.TargetPoolAggregatedList, error) {
	return c.internalClient.AggregatedList(ctx, req, opts...)
}

// Delete deletes the specified target pool.
func (c *TargetPoolsClient) Delete(ctx context.Context, req *computepb.DeleteTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Delete(ctx, req, opts...)
}

// Get returns the specified target pool. Gets a list of available target pools by making a list() request.
func (c *TargetPoolsClient) Get(ctx context.Context, req *computepb.GetTargetPoolRequest, opts ...gax.CallOption) (*computepb.TargetPool, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// GetHealth gets the most recent health check results for each IP for the instance that is referenced by the given target pool.
func (c *TargetPoolsClient) GetHealth(ctx context.Context, req *computepb.GetHealthTargetPoolRequest, opts ...gax.CallOption) (*computepb.TargetPoolInstanceHealth, error) {
	return c.internalClient.GetHealth(ctx, req, opts...)
}

// Insert creates a target pool in the specified project and region using the data included in the request.
func (c *TargetPoolsClient) Insert(ctx context.Context, req *computepb.InsertTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.Insert(ctx, req, opts...)
}

// List retrieves a list of target pools available to the specified project and region.
func (c *TargetPoolsClient) List(ctx context.Context, req *computepb.ListTargetPoolsRequest, opts ...gax.CallOption) (*computepb.TargetPoolList, error) {
	return c.internalClient.List(ctx, req, opts...)
}

// RemoveHealthCheck removes health check URL from a target pool.
func (c *TargetPoolsClient) RemoveHealthCheck(ctx context.Context, req *computepb.RemoveHealthCheckTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveHealthCheck(ctx, req, opts...)
}

// RemoveInstance removes instance URL from a target pool.
func (c *TargetPoolsClient) RemoveInstance(ctx context.Context, req *computepb.RemoveInstanceTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.RemoveInstance(ctx, req, opts...)
}

// SetBackup changes a backup target pool’s configurations.
func (c *TargetPoolsClient) SetBackup(ctx context.Context, req *computepb.SetBackupTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	return c.internalClient.SetBackup(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type targetPoolsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTargetPoolsRESTClient creates a new target pools rest client.
//
// The TargetPools API.
func NewTargetPoolsRESTClient(ctx context.Context, opts ...option.ClientOption) (*TargetPoolsClient, error) {
	clientOpts := append(defaultTargetPoolsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	c := &targetPoolsRESTClient{
		endpoint:   endpoint,
		httpClient: httpClient,
	}
	c.setGoogleClientInfo()

	return &TargetPoolsClient{internalClient: c, CallOptions: &TargetPoolsCallOptions{}}, nil
}

func defaultTargetPoolsRESTClientOptions() []option.ClientOption {
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
func (c *targetPoolsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *targetPoolsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *targetPoolsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// AddHealthCheck adds health check URLs to a target pool.
func (c *targetPoolsRESTClient) AddHealthCheck(ctx context.Context, req *computepb.AddHealthCheckTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTargetPoolsAddHealthCheckRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools/%v/addHealthCheck", req.GetProject(), req.GetRegion(), req.GetTargetPool())

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

// AddInstance adds an instance to a target pool.
func (c *targetPoolsRESTClient) AddInstance(ctx context.Context, req *computepb.AddInstanceTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTargetPoolsAddInstanceRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools/%v/addInstance", req.GetProject(), req.GetRegion(), req.GetTargetPool())

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

// AggregatedList retrieves an aggregated list of target pools.
func (c *targetPoolsRESTClient) AggregatedList(ctx context.Context, req *computepb.AggregatedListTargetPoolsRequest, opts ...gax.CallOption) (*computepb.TargetPoolAggregatedList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/aggregated/targetPools", req.GetProject())

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
	rsp := &computepb.TargetPoolAggregatedList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Delete deletes the specified target pool.
func (c *targetPoolsRESTClient) Delete(ctx context.Context, req *computepb.DeleteTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools/%v", req.GetProject(), req.GetRegion(), req.GetTargetPool())

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

// Get returns the specified target pool. Gets a list of available target pools by making a list() request.
func (c *targetPoolsRESTClient) Get(ctx context.Context, req *computepb.GetTargetPoolRequest, opts ...gax.CallOption) (*computepb.TargetPool, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools/%v", req.GetProject(), req.GetRegion(), req.GetTargetPool())

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
	rsp := &computepb.TargetPool{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// GetHealth gets the most recent health check results for each IP for the instance that is referenced by the given target pool.
func (c *targetPoolsRESTClient) GetHealth(ctx context.Context, req *computepb.GetHealthTargetPoolRequest, opts ...gax.CallOption) (*computepb.TargetPoolInstanceHealth, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetInstanceReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools/%v/getHealth", req.GetProject(), req.GetRegion(), req.GetTargetPool())

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
	rsp := &computepb.TargetPoolInstanceHealth{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// Insert creates a target pool in the specified project and region using the data included in the request.
func (c *targetPoolsRESTClient) Insert(ctx context.Context, req *computepb.InsertTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTargetPoolResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools", req.GetProject(), req.GetRegion())

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

// List retrieves a list of target pools available to the specified project and region.
func (c *targetPoolsRESTClient) List(ctx context.Context, req *computepb.ListTargetPoolsRequest, opts ...gax.CallOption) (*computepb.TargetPoolList, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools", req.GetProject(), req.GetRegion())

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
	rsp := &computepb.TargetPoolList{}

	return rsp, unm.Unmarshal(buf, rsp)
}

// RemoveHealthCheck removes health check URL from a target pool.
func (c *targetPoolsRESTClient) RemoveHealthCheck(ctx context.Context, req *computepb.RemoveHealthCheckTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTargetPoolsRemoveHealthCheckRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools/%v/removeHealthCheck", req.GetProject(), req.GetRegion(), req.GetTargetPool())

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

// RemoveInstance removes instance URL from a target pool.
func (c *targetPoolsRESTClient) RemoveInstance(ctx context.Context, req *computepb.RemoveInstanceTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTargetPoolsRemoveInstanceRequestResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools/%v/removeInstance", req.GetProject(), req.GetRegion(), req.GetTargetPool())

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

// SetBackup changes a backup target pool’s configurations.
func (c *targetPoolsRESTClient) SetBackup(ctx context.Context, req *computepb.SetBackupTargetPoolRequest, opts ...gax.CallOption) (*computepb.Operation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, EmitUnpopulated: true}
	body := req.GetTargetReferenceResource()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, _ := url.Parse(c.endpoint)
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/regions/%v/targetPools/%v/setBackup", req.GetProject(), req.GetRegion(), req.GetTargetPool())

	params := url.Values{}
	if req != nil && req.FailoverRatio != nil {
		params.Add("failoverRatio", fmt.Sprintf("%v", req.GetFailoverRatio()))
	}
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
