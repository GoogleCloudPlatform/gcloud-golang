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

// [START recaptchaenterprise_v1beta1_generated_RecaptchaEnterpriseServiceV1Beta1_ListKeys_sync]

package main

import (
	recaptchaenterprise "cloud.google.com/go/recaptchaenterprise/apiv1beta1"
	"context"
	"google.golang.org/api/iterator"
	recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1beta1"
)

func main() {
	// import recaptchaenterprisepb "google.golang.org/genproto/googleapis/cloud/recaptchaenterprise/v1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := recaptchaenterprise.NewRecaptchaEnterpriseServiceV1Beta1Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &recaptchaenterprisepb.ListKeysRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListKeys(ctx, req)
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
// [END recaptchaenterprise_v1beta1_generated_RecaptchaEnterpriseServiceV1Beta1_ListKeys_sync]
