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

// [START orgpolicy_v2_generated_OrgPolicy_CreatePolicy_sync]

package main

import (
	orgpolicy "cloud.google.com/go/orgpolicy/apiv2"
	"context"
	orgpolicypb "google.golang.org/genproto/googleapis/cloud/orgpolicy/v2"
)

func main() {
	// import orgpolicypb "google.golang.org/genproto/googleapis/cloud/orgpolicy/v2"

	ctx := context.Background()
	c, err := orgpolicy.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &orgpolicypb.CreatePolicyRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreatePolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
// [END orgpolicy_v2_generated_OrgPolicy_CreatePolicy_sync]
