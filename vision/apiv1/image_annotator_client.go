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

package vision

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newImageAnnotatorClientHook clientHook

// ImageAnnotatorCallOptions contains the retry settings for each method of ImageAnnotatorClient.
type ImageAnnotatorCallOptions struct {
	BatchAnnotateImages      []gax.CallOption
	BatchAnnotateFiles       []gax.CallOption
	AsyncBatchAnnotateImages []gax.CallOption
	AsyncBatchAnnotateFiles  []gax.CallOption
}

func defaultImageAnnotatorGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("vision.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("vision.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://vision.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultImageAnnotatorCallOptions() *ImageAnnotatorCallOptions {
	return &ImageAnnotatorCallOptions{
		BatchAnnotateImages: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchAnnotateFiles: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		AsyncBatchAnnotateImages: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		AsyncBatchAnnotateFiles: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
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

// internalImageAnnotatorClient is an interface that defines the methods availaible from Cloud Vision API.
type internalImageAnnotatorClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	BatchAnnotateImages(context.Context, *visionpb.BatchAnnotateImagesRequest, ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error)
	BatchAnnotateFiles(context.Context, *visionpb.BatchAnnotateFilesRequest, ...gax.CallOption) (*visionpb.BatchAnnotateFilesResponse, error)
	AsyncBatchAnnotateImages(context.Context, *visionpb.AsyncBatchAnnotateImagesRequest, ...gax.CallOption) (*AsyncBatchAnnotateImagesOperation, error)
	AsyncBatchAnnotateImagesOperation(name string) *AsyncBatchAnnotateImagesOperation
	AsyncBatchAnnotateFiles(context.Context, *visionpb.AsyncBatchAnnotateFilesRequest, ...gax.CallOption) (*AsyncBatchAnnotateFilesOperation, error)
	AsyncBatchAnnotateFilesOperation(name string) *AsyncBatchAnnotateFilesOperation
}

// ImageAnnotatorClient is a client for interacting with Cloud Vision API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service that performs Google Cloud Vision API detection tasks over client
// images, such as face, landmark, logo, label, and text detection. The
// ImageAnnotator service returns detected entities from the images.
type ImageAnnotatorClient struct {
	// The internal transport-dependent client.
	internalClient internalImageAnnotatorClient

	// The call options for this service.
	CallOptions *ImageAnnotatorCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ImageAnnotatorClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ImageAnnotatorClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ImageAnnotatorClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// BatchAnnotateImages run image detection and annotation for a batch of images.
func (c *ImageAnnotatorClient) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error) {
	return c.internalClient.BatchAnnotateImages(ctx, req, opts...)
}

// BatchAnnotateFiles service that performs image detection and annotation for a batch of files.
// Now only “application/pdf”, “image/tiff” and “image/gif” are supported.
//
// This service will extract at most 5 (customers can specify which 5 in
// AnnotateFileRequest.pages) frames (gif) or pages (pdf or tiff) from each
// file provided and perform detection and annotation for each image
// extracted.
func (c *ImageAnnotatorClient) BatchAnnotateFiles(ctx context.Context, req *visionpb.BatchAnnotateFilesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateFilesResponse, error) {
	return c.internalClient.BatchAnnotateFiles(ctx, req, opts...)
}

// AsyncBatchAnnotateImages run asynchronous image detection and annotation for a list of images.
//
// Progress and results can be retrieved through the
// google.longrunning.Operations interface.
// Operation.metadata contains OperationMetadata (metadata).
// Operation.response contains AsyncBatchAnnotateImagesResponse (results).
//
// This service will write image annotation outputs to json files in customer
// GCS bucket, each json file containing BatchAnnotateImagesResponse proto.
func (c *ImageAnnotatorClient) AsyncBatchAnnotateImages(ctx context.Context, req *visionpb.AsyncBatchAnnotateImagesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateImagesOperation, error) {
	return c.internalClient.AsyncBatchAnnotateImages(ctx, req, opts...)
}

// AsyncBatchAnnotateImagesOperation returns a new AsyncBatchAnnotateImagesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateImagesOperation, possibly from a different process.
func (c *ImageAnnotatorClient) AsyncBatchAnnotateImagesOperation(name string) *AsyncBatchAnnotateImagesOperation {
	return c.internalClient.AsyncBatchAnnotateImagesOperation(name)
}

// AsyncBatchAnnotateFiles run asynchronous image detection and annotation for a list of generic
// files, such as PDF files, which may contain multiple pages and multiple
// images per page. Progress and results can be retrieved through the
// google.longrunning.Operations interface.
// Operation.metadata contains OperationMetadata (metadata).
// Operation.response contains AsyncBatchAnnotateFilesResponse (results).
func (c *ImageAnnotatorClient) AsyncBatchAnnotateFiles(ctx context.Context, req *visionpb.AsyncBatchAnnotateFilesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateFilesOperation, error) {
	return c.internalClient.AsyncBatchAnnotateFiles(ctx, req, opts...)
}

// AsyncBatchAnnotateFilesOperation returns a new AsyncBatchAnnotateFilesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateFilesOperation, possibly from a different process.
func (c *ImageAnnotatorClient) AsyncBatchAnnotateFilesOperation(name string) *AsyncBatchAnnotateFilesOperation {
	return c.internalClient.AsyncBatchAnnotateFilesOperation(name)
}

// imageAnnotatorGRPCClient is a client for interacting with Cloud Vision API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type imageAnnotatorGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ImageAnnotatorClient
	CallOptions **ImageAnnotatorCallOptions

	// The gRPC API client.
	imageAnnotatorClient visionpb.ImageAnnotatorClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewImageAnnotatorClient creates a new image annotator client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service that performs Google Cloud Vision API detection tasks over client
// images, such as face, landmark, logo, label, and text detection. The
// ImageAnnotator service returns detected entities from the images.
func NewImageAnnotatorClient(ctx context.Context, opts ...option.ClientOption) (*ImageAnnotatorClient, error) {
	clientOpts := defaultImageAnnotatorGRPCClientOptions()
	if newImageAnnotatorClientHook != nil {
		hookOpts, err := newImageAnnotatorClientHook(ctx, clientHookParams{})
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
	client := ImageAnnotatorClient{CallOptions: defaultImageAnnotatorCallOptions()}

	c := &imageAnnotatorGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		imageAnnotatorClient: visionpb.NewImageAnnotatorClient(connPool),
		CallOptions:          &client.CallOptions,
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
func (c *imageAnnotatorGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *imageAnnotatorGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *imageAnnotatorGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *imageAnnotatorGRPCClient) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchAnnotateImages[0:len((*c.CallOptions).BatchAnnotateImages):len((*c.CallOptions).BatchAnnotateImages)], opts...)
	var resp *visionpb.BatchAnnotateImagesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.BatchAnnotateImages(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) BatchAnnotateFiles(ctx context.Context, req *visionpb.BatchAnnotateFilesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateFilesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchAnnotateFiles[0:len((*c.CallOptions).BatchAnnotateFiles):len((*c.CallOptions).BatchAnnotateFiles)], opts...)
	var resp *visionpb.BatchAnnotateFilesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.BatchAnnotateFiles(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) AsyncBatchAnnotateImages(ctx context.Context, req *visionpb.AsyncBatchAnnotateImagesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateImagesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AsyncBatchAnnotateImages[0:len((*c.CallOptions).AsyncBatchAnnotateImages):len((*c.CallOptions).AsyncBatchAnnotateImages)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.AsyncBatchAnnotateImages(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &AsyncBatchAnnotateImagesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *imageAnnotatorGRPCClient) AsyncBatchAnnotateFiles(ctx context.Context, req *visionpb.AsyncBatchAnnotateFilesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateFilesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AsyncBatchAnnotateFiles[0:len((*c.CallOptions).AsyncBatchAnnotateFiles):len((*c.CallOptions).AsyncBatchAnnotateFiles)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.AsyncBatchAnnotateFiles(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &AsyncBatchAnnotateFilesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// AsyncBatchAnnotateFilesOperation manages a long-running operation from AsyncBatchAnnotateFiles.
type AsyncBatchAnnotateFilesOperation struct {
	lro *longrunning.Operation
}

// AsyncBatchAnnotateFilesOperation returns a new AsyncBatchAnnotateFilesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateFilesOperation, possibly from a different process.
func (c *imageAnnotatorGRPCClient) AsyncBatchAnnotateFilesOperation(name string) *AsyncBatchAnnotateFilesOperation {
	return &AsyncBatchAnnotateFilesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *AsyncBatchAnnotateFilesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateFilesResponse, error) {
	var resp visionpb.AsyncBatchAnnotateFilesResponse
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
func (op *AsyncBatchAnnotateFilesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateFilesResponse, error) {
	var resp visionpb.AsyncBatchAnnotateFilesResponse
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
func (op *AsyncBatchAnnotateFilesOperation) Metadata() (*visionpb.OperationMetadata, error) {
	var meta visionpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *AsyncBatchAnnotateFilesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *AsyncBatchAnnotateFilesOperation) Name() string {
	return op.lro.Name()
}

// AsyncBatchAnnotateImagesOperation manages a long-running operation from AsyncBatchAnnotateImages.
type AsyncBatchAnnotateImagesOperation struct {
	lro *longrunning.Operation
}

// AsyncBatchAnnotateImagesOperation returns a new AsyncBatchAnnotateImagesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateImagesOperation, possibly from a different process.
func (c *imageAnnotatorGRPCClient) AsyncBatchAnnotateImagesOperation(name string) *AsyncBatchAnnotateImagesOperation {
	return &AsyncBatchAnnotateImagesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *AsyncBatchAnnotateImagesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateImagesResponse, error) {
	var resp visionpb.AsyncBatchAnnotateImagesResponse
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
func (op *AsyncBatchAnnotateImagesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateImagesResponse, error) {
	var resp visionpb.AsyncBatchAnnotateImagesResponse
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
func (op *AsyncBatchAnnotateImagesOperation) Metadata() (*visionpb.OperationMetadata, error) {
	var meta visionpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *AsyncBatchAnnotateImagesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *AsyncBatchAnnotateImagesOperation) Name() string {
	return op.lro.Name()
}
