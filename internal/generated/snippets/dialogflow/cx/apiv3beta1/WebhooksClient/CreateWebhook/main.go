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

// [START dialogflow_v3beta1_generated_Webhooks_CreateWebhook_sync]

package main

import (
	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"context"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
)

func main() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewWebhooksClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.CreateWebhookRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateWebhook(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
// [END dialogflow_v3beta1_generated_Webhooks_CreateWebhook_sync]
