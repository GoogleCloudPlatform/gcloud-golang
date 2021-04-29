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

// [START vision_v1_generated_ProductSearch_PurgeProducts_sync]

package main

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"
)

func main() {
	// import visionpb "google.golang.org/genproto/googleapis/cloud/vision/v1"

	ctx := context.Background()
	c, err := vision.NewProductSearchClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &visionpb.PurgeProductsRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.PurgeProducts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
// [END vision_v1_generated_ProductSearch_PurgeProducts_sync]
