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

// [START cloudscheduler_v1_generated_CloudScheduler_ListJobs_sync]

package main

import (
	"context"

	scheduler "cloud.google.com/go/scheduler/apiv1"
	"google.golang.org/api/iterator"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1"
)

func main() {
	ctx := context.Background()
	c, err := scheduler.NewCloudSchedulerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &schedulerpb.ListJobsRequest{
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

// [END cloudscheduler_v1_generated_CloudScheduler_ListJobs_sync]
