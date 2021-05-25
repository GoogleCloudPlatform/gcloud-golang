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

// [START pubsub_v1_generated_Subscriber_Pull_sync]

package main

import (
	"context"

	pubsub "cloud.google.com/go/pubsub/apiv1"
	pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
)

func main() {
	ctx := context.Background()
	c, err := pubsub.NewSubscriberClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &pubsubpb.PullRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.Pull(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

// [END pubsub_v1_generated_Subscriber_Pull_sync]
