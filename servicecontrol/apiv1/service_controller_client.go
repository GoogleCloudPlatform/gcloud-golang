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

package servicecontrol

import (
	"context"
	"fmt"
	"math"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	servicecontrolpb "google.golang.org/genproto/googleapis/api/servicecontrol/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newServiceControllerClientHook clientHook

// ServiceControllerCallOptions contains the retry settings for each method of ServiceControllerClient.
type ServiceControllerCallOptions struct {
	Check  []gax.CallOption
	Report []gax.CallOption
}

func defaultServiceControllerClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("servicecontrol.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("servicecontrol.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://servicecontrol.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultServiceControllerCallOptions() *ServiceControllerCallOptions {
	return &ServiceControllerCallOptions{
		Check:  []gax.CallOption{},
		Report: []gax.CallOption{},
	}
}

// ServiceControllerClient is a client for interacting with .
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type ServiceControllerClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	serviceControllerClient servicecontrolpb.ServiceControllerClient

	// The call options for this service.
	CallOptions *ServiceControllerCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewServiceControllerClient creates a new service controller client.
//
// Google Service Control API (at https://cloud.google.com/service-control/overview)
//
// Lets clients check and report operations against a managed
// service (at https://cloud.google.com/service-management/reference/rpc/google.api/servicemanagement.v1#google.api.servicemanagement.v1.ManagedService).
func NewServiceControllerClient(ctx context.Context, opts ...option.ClientOption) (*ServiceControllerClient, error) {
	clientOpts := defaultServiceControllerClientOptions()

	if newServiceControllerClientHook != nil {
		hookOpts, err := newServiceControllerClientHook(ctx, clientHookParams{})
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
	c := &ServiceControllerClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultServiceControllerCallOptions(),

		serviceControllerClient: servicecontrolpb.NewServiceControllerClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ServiceControllerClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ServiceControllerClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ServiceControllerClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Check checks whether an operation on a service should be allowed to proceed
// based on the configuration of the service and related policies. It must be
// called before the operation is executed.
//
// If feasible, the client should cache the check results and reuse them for
// 60 seconds. In case of any server errors, the client should rely on the
// cached results for much longer time to avoid outage.
// WARNING: There is general 60s delay for the configuration and policy
// propagation, therefore callers MUST NOT depend on the Check method having
// the latest policy information.
//
// NOTE: the CheckRequest has the size limit of 64KB.
//
// This method requires the servicemanagement.services.check permission
// on the specified service. For more information, see
// Cloud IAM (at https://cloud.google.com/iam).
func (c *ServiceControllerClient) Check(ctx context.Context, req *servicecontrolpb.CheckRequest, opts ...gax.CallOption) (*servicecontrolpb.CheckResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_name", url.QueryEscape(req.GetServiceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.Check[0:len(c.CallOptions.Check):len(c.CallOptions.Check)], opts...)
	var resp *servicecontrolpb.CheckResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceControllerClient.Check(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Report reports operation results to Google Service Control, such as logs and
// metrics. It should be called after an operation is completed.
//
// If feasible, the client should aggregate reporting data for up to 5
// seconds to reduce API traffic. Limiting aggregation to 5 seconds is to
// reduce data loss during client crashes. Clients should carefully choose
// the aggregation time window to avoid data loss risk more than 0.01%
// for business and compliance reasons.
//
// NOTE: the ReportRequest has the size limit (wire-format byte size) of
// 1MB.
//
// This method requires the servicemanagement.services.report permission
// on the specified service. For more information, see
// Google Cloud IAM (at https://cloud.google.com/iam).
func (c *ServiceControllerClient) Report(ctx context.Context, req *servicecontrolpb.ReportRequest, opts ...gax.CallOption) (*servicecontrolpb.ReportResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "service_name", url.QueryEscape(req.GetServiceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.Report[0:len(c.CallOptions.Report):len(c.CallOptions.Report)], opts...)
	var resp *servicecontrolpb.ReportResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.serviceControllerClient.Report(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
