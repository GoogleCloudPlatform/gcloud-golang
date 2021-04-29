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

// [START assuredworkloads_v1beta1_generated_AssuredWorkloadsService_ListWorkloads_sync]

package main

import (
	assuredworkloads "cloud.google.com/go/assuredworkloads/apiv1beta1"
	"context"
	"google.golang.org/api/iterator"
	assuredworkloadspb "google.golang.org/genproto/googleapis/cloud/assuredworkloads/v1beta1"
)

func main() {
	// import assuredworkloadspb "google.golang.org/genproto/googleapis/cloud/assuredworkloads/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := assuredworkloads.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &assuredworkloadspb.ListWorkloadsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListWorkloads(ctx, req)
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
// [END assuredworkloads_v1beta1_generated_AssuredWorkloadsService_ListWorkloads_sync]
