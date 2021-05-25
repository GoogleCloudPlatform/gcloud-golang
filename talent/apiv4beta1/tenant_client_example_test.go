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

package talent_test

import (
	"context"

	talent "cloud.google.com/go/talent/apiv4beta1"
	"google.golang.org/api/iterator"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
)

func ExampleNewTenantClient() {
	ctx := context.Background()
	c, err := talent.NewTenantClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleTenantClient_CreateTenant() {
	ctx := context.Background()
	c, err := talent.NewTenantClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.CreateTenantRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateTenant(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTenantClient_GetTenant() {
	ctx := context.Background()
	c, err := talent.NewTenantClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.GetTenantRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetTenant(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTenantClient_UpdateTenant() {
	ctx := context.Background()
	c, err := talent.NewTenantClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.UpdateTenantRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateTenant(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleTenantClient_DeleteTenant() {
	ctx := context.Background()
	c, err := talent.NewTenantClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.DeleteTenantRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteTenant(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleTenantClient_ListTenants() {
	ctx := context.Background()
	c, err := talent.NewTenantClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &talentpb.ListTenantsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListTenants(ctx, req)
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
