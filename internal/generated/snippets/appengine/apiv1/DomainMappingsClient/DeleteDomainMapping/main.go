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

// [START appengine_v1_generated_DomainMappings_DeleteDomainMapping_sync]

package main

import (
	"context"

	appengine "cloud.google.com/go/appengine/apiv1"
	appenginepb "google.golang.org/genproto/googleapis/appengine/v1"
)

func main() {
	ctx := context.Background()
	c, err := appengine.NewDomainMappingsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &appenginepb.DeleteDomainMappingRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteDomainMapping(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

// [END appengine_v1_generated_DomainMappings_DeleteDomainMapping_sync]
