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

// [START gameservices_v1_generated_GameServerClustersService_ListGameServerClusters_sync]

package main

import (
	gaming "cloud.google.com/go/gaming/apiv1"
	"context"
	"google.golang.org/api/iterator"
	gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"
)

func main() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := gaming.NewGameServerClustersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.ListGameServerClustersRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListGameServerClusters(ctx, req)
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
// [END gameservices_v1_generated_GameServerClustersService_ListGameServerClusters_sync]
