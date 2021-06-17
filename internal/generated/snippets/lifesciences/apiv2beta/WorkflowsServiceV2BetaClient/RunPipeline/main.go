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

// [START lifesciences_v2beta_generated_WorkflowsServiceV2Beta_RunPipeline_sync]

package main

import (
	"context"

	lifesciences "cloud.google.com/go/lifesciences/apiv2beta"
	lifesciencespb "google.golang.org/genproto/googleapis/cloud/lifesciences/v2beta"
)

func main() {
	ctx := context.Background()
	c, err := lifesciences.NewWorkflowsServiceV2BetaClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &lifesciencespb.RunPipelineRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.RunPipeline(ctx, req)
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

// [END lifesciences_v2beta_generated_WorkflowsServiceV2Beta_RunPipeline_sync]
