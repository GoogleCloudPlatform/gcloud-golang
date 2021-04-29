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

// [START cloudresourcemanager_v2_generated_Folders_UndeleteFolder_sync]

package main

import (
	resourcemanager "cloud.google.com/go/resourcemanager/apiv2"
	"context"
	resourcemanagerpb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v2"
)

func main() {
	// import resourcemanagerpb "google.golang.org/genproto/googleapis/cloud/resourcemanager/v2"

	ctx := context.Background()
	c, err := resourcemanager.NewFoldersClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &resourcemanagerpb.UndeleteFolderRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UndeleteFolder(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
// [END cloudresourcemanager_v2_generated_Folders_UndeleteFolder_sync]
