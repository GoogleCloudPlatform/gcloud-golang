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

package cx

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
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newTestCasesClientHook clientHook

// TestCasesCallOptions contains the retry settings for each method of TestCasesClient.
type TestCasesCallOptions struct {
	ListTestCases        []gax.CallOption
	BatchDeleteTestCases []gax.CallOption
	GetTestCase          []gax.CallOption
	CreateTestCase       []gax.CallOption
	UpdateTestCase       []gax.CallOption
	RunTestCase          []gax.CallOption
	BatchRunTestCases    []gax.CallOption
	CalculateCoverage    []gax.CallOption
	ImportTestCases      []gax.CallOption
	ExportTestCases      []gax.CallOption
	ListTestCaseResults  []gax.CallOption
}

func defaultTestCasesClientOptions() []option.ClientOption {
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

func defaultTestCasesCallOptions() *TestCasesCallOptions {
	return &TestCasesCallOptions{
		ListTestCases: []gax.CallOption{
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
		BatchDeleteTestCases: []gax.CallOption{
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
		GetTestCase: []gax.CallOption{
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
		CreateTestCase: []gax.CallOption{
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
		UpdateTestCase: []gax.CallOption{
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
		RunTestCase: []gax.CallOption{
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
		BatchRunTestCases: []gax.CallOption{
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
		CalculateCoverage: []gax.CallOption{
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
		ImportTestCases: []gax.CallOption{
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
		ExportTestCases: []gax.CallOption{
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
		ListTestCaseResults: []gax.CallOption{
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

// TestCasesClient is a client for interacting with Dialogflow API.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type TestCasesClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// The gRPC API client.
	testCasesClient cxpb.TestCasesClient

	// LROClient is used internally to handle longrunning operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient

	// The call options for this service.
	CallOptions *TestCasesCallOptions

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTestCasesClient creates a new test cases client.
//
// Service for managing [Test Cases][google.cloud.dialogflow.cx.v3beta1.TestCase] and
// [Test Case Results][google.cloud.dialogflow.cx.v3beta1.TestCaseResult].
func NewTestCasesClient(ctx context.Context, opts ...option.ClientOption) (*TestCasesClient, error) {
	clientOpts := defaultTestCasesClientOptions()

	if newTestCasesClientHook != nil {
		hookOpts, err := newTestCasesClientHook(ctx, clientHookParams{})
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
	c := &TestCasesClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		CallOptions:      defaultTestCasesCallOptions(),

		testCasesClient: cxpb.NewTestCasesClient(connPool),
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
func (c *TestCasesClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TestCasesClient) Close() error {
	return c.connPool.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TestCasesClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// ListTestCases fetches a list of test cases for a given agent.
func (c *TestCasesClient) ListTestCases(ctx context.Context, req *cxpb.ListTestCasesRequest, opts ...gax.CallOption) *TestCaseIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListTestCases[0:len(c.CallOptions.ListTestCases):len(c.CallOptions.ListTestCases)], opts...)
	it := &TestCaseIterator{}
	req = proto.Clone(req).(*cxpb.ListTestCasesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.TestCase, string, error) {
		var resp *cxpb.ListTestCasesResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.testCasesClient.ListTestCases(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTestCases(), resp.GetNextPageToken(), nil
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

// BatchDeleteTestCases batch deletes test cases.
func (c *TestCasesClient) BatchDeleteTestCases(ctx context.Context, req *cxpb.BatchDeleteTestCasesRequest, opts ...gax.CallOption) error {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchDeleteTestCases[0:len(c.CallOptions.BatchDeleteTestCases):len(c.CallOptions.BatchDeleteTestCases)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.testCasesClient.BatchDeleteTestCases(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

// GetTestCase gets a test case.
func (c *TestCasesClient) GetTestCase(ctx context.Context, req *cxpb.GetTestCaseRequest, opts ...gax.CallOption) (*cxpb.TestCase, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.GetTestCase[0:len(c.CallOptions.GetTestCase):len(c.CallOptions.GetTestCase)], opts...)
	var resp *cxpb.TestCase
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testCasesClient.GetTestCase(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateTestCase creates a test case for the given agent.
func (c *TestCasesClient) CreateTestCase(ctx context.Context, req *cxpb.CreateTestCaseRequest, opts ...gax.CallOption) (*cxpb.TestCase, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CreateTestCase[0:len(c.CallOptions.CreateTestCase):len(c.CallOptions.CreateTestCase)], opts...)
	var resp *cxpb.TestCase
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testCasesClient.CreateTestCase(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateTestCase updates the specified test case.
func (c *TestCasesClient) UpdateTestCase(ctx context.Context, req *cxpb.UpdateTestCaseRequest, opts ...gax.CallOption) (*cxpb.TestCase, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "test_case.name", url.QueryEscape(req.GetTestCase().GetName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.UpdateTestCase[0:len(c.CallOptions.UpdateTestCase):len(c.CallOptions.UpdateTestCase)], opts...)
	var resp *cxpb.TestCase
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testCasesClient.UpdateTestCase(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RunTestCase kicks off a test case run.
func (c *TestCasesClient) RunTestCase(ctx context.Context, req *cxpb.RunTestCaseRequest, opts ...gax.CallOption) (*RunTestCaseOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append(c.CallOptions.RunTestCase[0:len(c.CallOptions.RunTestCase):len(c.CallOptions.RunTestCase)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testCasesClient.RunTestCase(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &RunTestCaseOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// BatchRunTestCases kicks off a batch run of test cases.
func (c *TestCasesClient) BatchRunTestCases(ctx context.Context, req *cxpb.BatchRunTestCasesRequest, opts ...gax.CallOption) (*BatchRunTestCasesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.BatchRunTestCases[0:len(c.CallOptions.BatchRunTestCases):len(c.CallOptions.BatchRunTestCases)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testCasesClient.BatchRunTestCases(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &BatchRunTestCasesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// CalculateCoverage calculates the test coverage for an agent.
func (c *TestCasesClient) CalculateCoverage(ctx context.Context, req *cxpb.CalculateCoverageRequest, opts ...gax.CallOption) (*cxpb.CalculateCoverageResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "agent", url.QueryEscape(req.GetAgent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.CalculateCoverage[0:len(c.CallOptions.CalculateCoverage):len(c.CallOptions.CalculateCoverage)], opts...)
	var resp *cxpb.CalculateCoverageResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testCasesClient.CalculateCoverage(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ImportTestCases imports the test cases from a Cloud Storage bucket or a local file. It
// always creates new test cases and won’t overwite any existing ones. The
// provided ID in the imported test case is neglected.
func (c *TestCasesClient) ImportTestCases(ctx context.Context, req *cxpb.ImportTestCasesRequest, opts ...gax.CallOption) (*ImportTestCasesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ImportTestCases[0:len(c.CallOptions.ImportTestCases):len(c.CallOptions.ImportTestCases)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testCasesClient.ImportTestCases(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ImportTestCasesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// ExportTestCases exports the test cases under the agent to a Cloud Storage bucket or a local
// file. Filter can be applied to export a subset of test cases.
func (c *TestCasesClient) ExportTestCases(ctx context.Context, req *cxpb.ExportTestCasesRequest, opts ...gax.CallOption) (*ExportTestCasesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ExportTestCases[0:len(c.CallOptions.ExportTestCases):len(c.CallOptions.ExportTestCases)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.testCasesClient.ExportTestCases(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &ExportTestCasesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, resp),
	}, nil
}

// ListTestCaseResults fetches a list of results for a given test case.
func (c *TestCasesClient) ListTestCaseResults(ctx context.Context, req *cxpb.ListTestCaseResultsRequest, opts ...gax.CallOption) *TestCaseResultIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append(c.CallOptions.ListTestCaseResults[0:len(c.CallOptions.ListTestCaseResults):len(c.CallOptions.ListTestCaseResults)], opts...)
	it := &TestCaseResultIterator{}
	req = proto.Clone(req).(*cxpb.ListTestCaseResultsRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*cxpb.TestCaseResult, string, error) {
		var resp *cxpb.ListTestCaseResultsResponse
		req.PageToken = pageToken
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.testCasesClient.ListTestCaseResults(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTestCaseResults(), resp.GetNextPageToken(), nil
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

// BatchRunTestCasesOperation manages a long-running operation from BatchRunTestCases.
type BatchRunTestCasesOperation struct {
	lro *longrunning.Operation
}

// BatchRunTestCasesOperation returns a new BatchRunTestCasesOperation from a given name.
// The name must be that of a previously created BatchRunTestCasesOperation, possibly from a different process.
func (c *TestCasesClient) BatchRunTestCasesOperation(name string) *BatchRunTestCasesOperation {
	return &BatchRunTestCasesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *BatchRunTestCasesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*cxpb.BatchRunTestCasesResponse, error) {
	var resp cxpb.BatchRunTestCasesResponse
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
func (op *BatchRunTestCasesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*cxpb.BatchRunTestCasesResponse, error) {
	var resp cxpb.BatchRunTestCasesResponse
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
func (op *BatchRunTestCasesOperation) Metadata() (*cxpb.BatchRunTestCasesMetadata, error) {
	var meta cxpb.BatchRunTestCasesMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *BatchRunTestCasesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *BatchRunTestCasesOperation) Name() string {
	return op.lro.Name()
}

// ExportTestCasesOperation manages a long-running operation from ExportTestCases.
type ExportTestCasesOperation struct {
	lro *longrunning.Operation
}

// ExportTestCasesOperation returns a new ExportTestCasesOperation from a given name.
// The name must be that of a previously created ExportTestCasesOperation, possibly from a different process.
func (c *TestCasesClient) ExportTestCasesOperation(name string) *ExportTestCasesOperation {
	return &ExportTestCasesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ExportTestCasesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*cxpb.ExportTestCasesResponse, error) {
	var resp cxpb.ExportTestCasesResponse
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
func (op *ExportTestCasesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*cxpb.ExportTestCasesResponse, error) {
	var resp cxpb.ExportTestCasesResponse
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
func (op *ExportTestCasesOperation) Metadata() (*cxpb.ExportTestCasesMetadata, error) {
	var meta cxpb.ExportTestCasesMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ExportTestCasesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ExportTestCasesOperation) Name() string {
	return op.lro.Name()
}

// ImportTestCasesOperation manages a long-running operation from ImportTestCases.
type ImportTestCasesOperation struct {
	lro *longrunning.Operation
}

// ImportTestCasesOperation returns a new ImportTestCasesOperation from a given name.
// The name must be that of a previously created ImportTestCasesOperation, possibly from a different process.
func (c *TestCasesClient) ImportTestCasesOperation(name string) *ImportTestCasesOperation {
	return &ImportTestCasesOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *ImportTestCasesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*cxpb.ImportTestCasesResponse, error) {
	var resp cxpb.ImportTestCasesResponse
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
func (op *ImportTestCasesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*cxpb.ImportTestCasesResponse, error) {
	var resp cxpb.ImportTestCasesResponse
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
func (op *ImportTestCasesOperation) Metadata() (*cxpb.ImportTestCasesMetadata, error) {
	var meta cxpb.ImportTestCasesMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *ImportTestCasesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *ImportTestCasesOperation) Name() string {
	return op.lro.Name()
}

// RunTestCaseOperation manages a long-running operation from RunTestCase.
type RunTestCaseOperation struct {
	lro *longrunning.Operation
}

// RunTestCaseOperation returns a new RunTestCaseOperation from a given name.
// The name must be that of a previously created RunTestCaseOperation, possibly from a different process.
func (c *TestCasesClient) RunTestCaseOperation(name string) *RunTestCaseOperation {
	return &RunTestCaseOperation{
		lro: longrunning.InternalNewOperation(c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *RunTestCaseOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*cxpb.RunTestCaseResponse, error) {
	var resp cxpb.RunTestCaseResponse
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
func (op *RunTestCaseOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*cxpb.RunTestCaseResponse, error) {
	var resp cxpb.RunTestCaseResponse
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
func (op *RunTestCaseOperation) Metadata() (*cxpb.RunTestCaseMetadata, error) {
	var meta cxpb.RunTestCaseMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *RunTestCaseOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *RunTestCaseOperation) Name() string {
	return op.lro.Name()
}

// TestCaseIterator manages a stream of *cxpb.TestCase.
type TestCaseIterator struct {
	items    []*cxpb.TestCase
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
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.TestCase, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TestCaseIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TestCaseIterator) Next() (*cxpb.TestCase, error) {
	var item *cxpb.TestCase
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TestCaseIterator) bufLen() int {
	return len(it.items)
}

func (it *TestCaseIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// TestCaseResultIterator manages a stream of *cxpb.TestCaseResult.
type TestCaseResultIterator struct {
	items    []*cxpb.TestCaseResult
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
	InternalFetch func(pageSize int, pageToken string) (results []*cxpb.TestCaseResult, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TestCaseResultIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TestCaseResultIterator) Next() (*cxpb.TestCaseResult, error) {
	var item *cxpb.TestCaseResult
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TestCaseResultIterator) bufLen() int {
	return len(it.items)
}

func (it *TestCaseResultIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
