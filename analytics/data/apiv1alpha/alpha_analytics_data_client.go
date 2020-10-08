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

package data

import (
	"context"
	"math"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	gtransport "google.golang.org/api/transport/grpc"
	datapb "google.golang.org/genproto/googleapis/analytics/data/v1alpha"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newAlphaAnalyticsDataClientHook clientHook

// AlphaAnalyticsDataCallOptions contains the retry settings for each method of AlphaAnalyticsDataClient.
type AlphaAnalyticsDataCallOptions struct {
	RunReport            []gax.CallOption
	RunPivotReport       []gax.CallOption
	BatchRunReports      []gax.CallOption
	BatchRunPivotReports []gax.CallOption
	GetUniversalMetadata []gax.CallOption
}

func defaultAlphaAnalyticsDataClientOptions() []option.ClientOption {
	return []option.ClientOption{
		option.WithEndpoint("analyticsdata.googleapis.com:443"),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultAlphaAnalyticsDataCallOptions() *AlphaAnalyticsDataCallOptions {
	return &AlphaAnalyticsDataCallOptions{
		RunReport:            []gax.CallOption{},
		RunPivotReport:       []gax.CallOption{},
		BatchRunReports:      []gax.CallOption{},
		BatchRunPivotReports: []gax.CallOption{},
		GetUniversalMetadata: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unknown,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// AlphaAnalyticsDataClient is a client for interacting with .
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type AlphaAnalyticsDataClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	alphaAnalyticsDataClient datapb.AlphaAnalyticsDataClient

	// The call options for this service.
	CallOptions *AlphaAnalyticsDataCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewAlphaAnalyticsDataClient creates a new alpha analytics data client.
//
// Google Analytics reporting data service.
func NewAlphaAnalyticsDataClient(ctx context.Context, opts ...option.ClientOption) (*AlphaAnalyticsDataClient, error) {
	clientOpts := defaultAlphaAnalyticsDataClientOptions()

	if newAlphaAnalyticsDataClientHook != nil {
		hookOpts, err := newAlphaAnalyticsDataClientHook(ctx, clientHookParams{})
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
	c := &AlphaAnalyticsDataClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultAlphaAnalyticsDataCallOptions(),

		alphaAnalyticsDataClient: datapb.NewAlphaAnalyticsDataClient(connPool),
	}
	c.setGoogleClientInfo()

	return c, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *AlphaAnalyticsDataClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *AlphaAnalyticsDataClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *AlphaAnalyticsDataClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// RunReport returns a customized report of your Google Analytics event data. Reports
// contain statistics derived from data collected by the Google Analytics
// tracking code. The data returned from the API is as a table with columns
// for the requested dimensions and metrics. Metrics are individual
// measurements of user activity on your property, such as active users or
// event count. Dimensions break down metrics across some common criteria,
// such as country or event name.
func (c *AlphaAnalyticsDataClient) RunReport(ctx context.Context, req *datapb.RunReportRequest, opts ...gax.CallOption) (*datapb.RunReportResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.RunReport[0:len(c.CallOptions.RunReport):len(c.CallOptions.RunReport)], opts...)
	var resp *datapb.RunReportResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.alphaAnalyticsDataClient.RunReport(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RunPivotReport returns a customized pivot report of your Google Analytics event data.
// Pivot reports are more advanced and expressive formats than regular
// reports. In a pivot report, dimensions are only visible if they are
// included in a pivot. Multiple pivots can be specified to further dissect
// your data.
func (c *AlphaAnalyticsDataClient) RunPivotReport(ctx context.Context, req *datapb.RunPivotReportRequest, opts ...gax.CallOption) (*datapb.RunPivotReportResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.RunPivotReport[0:len(c.CallOptions.RunPivotReport):len(c.CallOptions.RunPivotReport)], opts...)
	var resp *datapb.RunPivotReportResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.alphaAnalyticsDataClient.RunPivotReport(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchRunReports returns multiple reports in a batch. All reports must be for the same
// Entity.
func (c *AlphaAnalyticsDataClient) BatchRunReports(ctx context.Context, req *datapb.BatchRunReportsRequest, opts ...gax.CallOption) (*datapb.BatchRunReportsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.BatchRunReports[0:len(c.CallOptions.BatchRunReports):len(c.CallOptions.BatchRunReports)], opts...)
	var resp *datapb.BatchRunReportsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.alphaAnalyticsDataClient.BatchRunReports(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// BatchRunPivotReports returns multiple pivot reports in a batch. All reports must be for the same
// Entity.
func (c *AlphaAnalyticsDataClient) BatchRunPivotReports(ctx context.Context, req *datapb.BatchRunPivotReportsRequest, opts ...gax.CallOption) (*datapb.BatchRunPivotReportsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.BatchRunPivotReports[0:len(c.CallOptions.BatchRunPivotReports):len(c.CallOptions.BatchRunPivotReports)], opts...)
	var resp *datapb.BatchRunPivotReportsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.alphaAnalyticsDataClient.BatchRunPivotReports(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetUniversalMetadata returns metadata for dimensions and metrics available in reporting methods.
// Used to explore the dimensions and metrics. Dimensions and metrics will be
// mostly added over time, but renames and deletions may occur.
//
// This method returns Universal Metadata. Universal Metadata are dimensions
// and metrics applicable to any property such as country and totalUsers.
func (c *AlphaAnalyticsDataClient) GetUniversalMetadata(ctx context.Context, req *datapb.GetUniversalMetadataRequest, opts ...gax.CallOption) (*datapb.UniversalMetadata, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.GetUniversalMetadata[0:len(c.CallOptions.GetUniversalMetadata):len(c.CallOptions.GetUniversalMetadata)], opts...)
	var resp *datapb.UniversalMetadata
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.alphaAnalyticsDataClient.GetUniversalMetadata(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
