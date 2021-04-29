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

// [START spanner_v1_generated_Spanner_GetSession_sync]

package main

import (
	spanner "cloud.google.com/go/spanner/apiv1"
	"context"
	spannerpb "google.golang.org/genproto/googleapis/spanner/v1"
)

func main() {
	// import spannerpb "google.golang.org/genproto/googleapis/spanner/v1"

	ctx := context.Background()
	c, err := spanner.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &spannerpb.GetSessionRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetSession(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
// [END spanner_v1_generated_Spanner_GetSession_sync]
