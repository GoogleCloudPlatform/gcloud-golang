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

package compute_test

import (
	"context"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

func ExampleNewInstanceGroupManagersRESTClient() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleInstanceGroupManagersClient_AbandonInstances() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AbandonInstancesInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AbandonInstances(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_AggregatedList() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.AggregatedListInstanceGroupManagersRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.AggregatedList(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_ApplyUpdatesToInstances() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ApplyUpdatesToInstancesInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ApplyUpdatesToInstances(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_CreateInstances() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.CreateInstancesInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateInstances(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_Delete() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DeleteInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Delete(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_DeleteInstances() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DeleteInstancesInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.DeleteInstances(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_DeletePerInstanceConfigs() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.DeletePerInstanceConfigsInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.DeletePerInstanceConfigs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_Get() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.GetInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Get(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_Insert() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.InsertInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Insert(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_List() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListInstanceGroupManagersRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.List(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_ListErrors() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListErrorsInstanceGroupManagersRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListErrors(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_ListManagedInstances() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListManagedInstancesInstanceGroupManagersRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListManagedInstances(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_ListPerInstanceConfigs() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ListPerInstanceConfigsInstanceGroupManagersRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ListPerInstanceConfigs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_Patch() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.PatchInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Patch(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_PatchPerInstanceConfigs() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.PatchPerInstanceConfigsInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.PatchPerInstanceConfigs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_RecreateInstances() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.RecreateInstancesInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.RecreateInstances(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_Resize() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.ResizeInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Resize(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_SetInstanceTemplate() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.SetInstanceTemplateInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetInstanceTemplate(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_SetTargetPools() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.SetTargetPoolsInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.SetTargetPools(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleInstanceGroupManagersClient_UpdatePerInstanceConfigs() {
	ctx := context.Background()
	c, err := compute.NewInstanceGroupManagersRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &computepb.UpdatePerInstanceConfigsInstanceGroupManagerRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdatePerInstanceConfigs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
