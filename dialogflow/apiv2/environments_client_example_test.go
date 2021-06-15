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

package dialogflow_test

import (
	"context"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"google.golang.org/api/iterator"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
)

func ExampleNewEnvironmentsClient() {
	ctx := context.Background()
	c, err := dialogflow.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleEnvironmentsClient_ListEnvironments() {
	ctx := context.Background()
	c, err := dialogflow.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.ListEnvironmentsRequest{
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
	ctx := context.Background()
	c, err := dialogflow.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.GetEnvironmentRequest{
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
	ctx := context.Background()
	c, err := dialogflow.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.CreateEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEnvironmentsClient_UpdateEnvironment() {
	ctx := context.Background()
	c, err := dialogflow.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.UpdateEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleEnvironmentsClient_DeleteEnvironment() {
	ctx := context.Background()
	c, err := dialogflow.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.DeleteEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteEnvironment(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleEnvironmentsClient_GetEnvironmentHistory() {
	ctx := context.Background()
	c, err := dialogflow.NewEnvironmentsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &dialogflowpb.GetEnvironmentHistoryRequest{
		// TODO: Fill request struct fields.
	}
	it := c.GetEnvironmentHistory(ctx, req)
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
