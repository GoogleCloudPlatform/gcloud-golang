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

// [START jobs_generated_talent_apiv4_JobClient_ListJobs]

package main

import (
	"context"

	talent "cloud.google.com/go/talent/apiv4"
	"google.golang.org/api/iterator"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4"
)

func main() {
	// import talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.ListJobsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListJobs(ctx, req)
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

// [END jobs_generated_talent_apiv4_JobClient_ListJobs]
