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

// [START dialogflow_v3beta1_generated_SecuritySettingsService_ListSecuritySettings_sync]

package main

import (
	"context"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"google.golang.org/api/iterator"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
)

func main() {
	ctx := context.Background()
	c, err := cx.NewSecuritySettingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &cxpb.ListSecuritySettingsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListSecuritySettings(ctx, req)
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

// [END dialogflow_v3beta1_generated_SecuritySettingsService_ListSecuritySettings_sync]
