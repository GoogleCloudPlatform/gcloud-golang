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

// [START cloudshell_v1_generated_CloudShellService_AuthorizeEnvironment_sync]

package main

import (
	"context"

	shell "cloud.google.com/go/shell/apiv1"
	shellpb "google.golang.org/genproto/googleapis/cloud/shell/v1"
)

func main() {
	ctx := context.Background()
	c, err := shell.NewCloudShellClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &shellpb.AuthorizeEnvironmentRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.AuthorizeEnvironment(ctx, req)
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

// [END cloudshell_v1_generated_CloudShellService_AuthorizeEnvironment_sync]
