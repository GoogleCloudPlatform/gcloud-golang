// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START networkconnectivity_v1alpha1_generated_HubService_CreateSpoke_sync]

package main

import (
	networkconnectivity "cloud.google.com/go/networkconnectivity/apiv1alpha1"
	"context"
	networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"
)

func main() {
	// import networkconnectivitypb "google.golang.org/genproto/googleapis/cloud/networkconnectivity/v1alpha1"

	ctx := context.Background()
	c, err := networkconnectivity.NewHubClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &networkconnectivitypb.CreateSpokeRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateSpoke(ctx, req)
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
// [END networkconnectivity_v1alpha1_generated_HubService_CreateSpoke_sync]
