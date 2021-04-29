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

// [START analyticsadmin_v1alpha_generated_AnalyticsAdminService_DeleteWebDataStream_sync]

package main

import (
	admin "cloud.google.com/go/analytics/admin/apiv1alpha"
	"context"
	adminpb "google.golang.org/genproto/googleapis/analytics/admin/v1alpha"
)

func main() {
	ctx := context.Background()
	c, err := admin.NewAnalyticsAdminClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &adminpb.DeleteWebDataStreamRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteWebDataStream(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
// [END analyticsadmin_v1alpha_generated_AnalyticsAdminService_DeleteWebDataStream_sync]
