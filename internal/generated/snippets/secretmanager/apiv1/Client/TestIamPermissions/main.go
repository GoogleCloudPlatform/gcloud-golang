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

// [START secretmanager_v1_generated_SecretManagerService_TestIamPermissions_sync]

package main

import (
	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"context"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func main() {
	// import iampb "google.golang.org/genproto/googleapis/iam/v1"

	ctx := context.Background()
	c, err := secretmanager.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
// [END secretmanager_v1_generated_SecretManagerService_TestIamPermissions_sync]
