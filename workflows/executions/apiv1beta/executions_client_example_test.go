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

package executions_test

import (
	"context"

	executions "cloud.google.com/go/workflows/executions/apiv1beta"
	"google.golang.org/api/iterator"
	executionspb "google.golang.org/genproto/googleapis/cloud/workflows/executions/v1beta"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := executions.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_ListExecutions() {
	// import executionspb "google.golang.org/genproto/googleapis/cloud/workflows/executions/v1beta"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := executions.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &executionspb.ListExecutionsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListExecutions(ctx, req)
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

func ExampleClient_CreateExecution() {
	// import executionspb "google.golang.org/genproto/googleapis/cloud/workflows/executions/v1beta"

	ctx := context.Background()
	c, err := executions.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &executionspb.CreateExecutionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateExecution(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetExecution() {
	// import executionspb "google.golang.org/genproto/googleapis/cloud/workflows/executions/v1beta"

	ctx := context.Background()
	c, err := executions.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &executionspb.GetExecutionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetExecution(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CancelExecution() {
	// import executionspb "google.golang.org/genproto/googleapis/cloud/workflows/executions/v1beta"

	ctx := context.Background()
	c, err := executions.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &executionspb.CancelExecutionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CancelExecution(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
