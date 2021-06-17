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

// [START pubsublite_v1_generated_TopicStatsService_ComputeTimeCursor_sync]

package main

import (
	"context"

	pubsublite "cloud.google.com/go/pubsublite/apiv1"
	pubsublitepb "google.golang.org/genproto/googleapis/cloud/pubsublite/v1"
)

func main() {
	ctx := context.Background()
	c, err := pubsublite.NewTopicStatsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &pubsublitepb.ComputeTimeCursorRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.ComputeTimeCursor(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END pubsublite_v1_generated_TopicStatsService_ComputeTimeCursor_sync]
