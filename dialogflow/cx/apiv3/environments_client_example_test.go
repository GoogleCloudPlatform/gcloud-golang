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

package cx_test

import (
	"context"

	cx "cloud.google.com/go/dialogflow/cx/apiv3"
	"google.golang.org/api/iterator"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"
)

func ExampleNewEnvironmentsClient() {
	ctx := context.Background()
	c, err := cx.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleEnvironmentsClient_ListEnvironments() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := cx.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.ListEnvironmentsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListEnvironments(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleEnvironmentsClient_GetEnvironment() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"

	ctx := context.Background()
	c, err := cx.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.GetEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEnvironmentsClient_CreateEnvironment() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"

	ctx := context.Background()
	c, err := cx.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.CreateEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEnvironmentsClient_UpdateEnvironment() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"

	ctx := context.Background()
	c, err := cx.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.UpdateEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UpdateEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEnvironmentsClient_DeleteEnvironment() {
	ctx := context.Background()
	c, err := cx.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.DeleteEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleEnvironmentsClient_LookupEnvironmentHistory() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := cx.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.LookupEnvironmentHistoryRequest{
		// TODO: Fill request struct fields.
	}
	it := c.LookupEnvironmentHistory(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
