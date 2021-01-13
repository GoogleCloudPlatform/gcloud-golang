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

package pubsub_test

import (
	"context"

	pubsub "cloud.google.com/go/pubsub/apiv1"
	"google.golang.org/api/iterator"
	pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
)

func ExampleNewSchemaClient() {
	ctx := context.Background()
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleSchemaClient_CreateSchema() {
	// import pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"

	ctx := context.Background()
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &pubsubpb.CreateSchemaRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateSchema(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSchemaClient_GetSchema() {
	// import pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"

	ctx := context.Background()
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &pubsubpb.GetSchemaRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetSchema(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSchemaClient_ListSchemas() {
	// import pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &pubsubpb.ListSchemasRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListSchemas(ctx, req)
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

func ExampleSchemaClient_DeleteSchema() {
	ctx := context.Background()
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &pubsubpb.DeleteSchemaRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteSchema(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleSchemaClient_ValidateSchema() {
	// import pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"

	ctx := context.Background()
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &pubsubpb.ValidateSchemaRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ValidateSchema(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSchemaClient_ValidateMessage() {
	// import pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"

	ctx := context.Background()
	c, err := pubsub.NewSchemaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &pubsubpb.ValidateMessageRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ValidateMessage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
