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

// [START gameservices_v1_generated_RealmsService_DeleteRealm_sync]

package main

import (
	gaming "cloud.google.com/go/gaming/apiv1"
	"context"
	gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"
)

func main() {
	// import gamingpb "google.golang.org/genproto/googleapis/cloud/gaming/v1"

	ctx := context.Background()
	c, err := gaming.NewRealmsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &gamingpb.DeleteRealmRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteRealm(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
// [END gameservices_v1_generated_RealmsService_DeleteRealm_sync]
