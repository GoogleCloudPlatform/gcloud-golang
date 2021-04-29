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

// [START domains_v1beta1_generated_Domains_DeleteRegistration_sync]

package main

import (
	domains "cloud.google.com/go/domains/apiv1beta1"
	"context"
	domainspb "google.golang.org/genproto/googleapis/cloud/domains/v1beta1"
)

func main() {
	// import domainspb "google.golang.org/genproto/googleapis/cloud/domains/v1beta1"

	ctx := context.Background()
	c, err := domains.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &domainspb.DeleteRegistrationRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteRegistration(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}
// [END domains_v1beta1_generated_Domains_DeleteRegistration_sync]
