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

// [START clouderrorreporting_v1beta1_generated_ErrorStatsService_DeleteEvents_sync]

package main

import (
	errorreporting "cloud.google.com/go/errorreporting/apiv1beta1"
	"context"
	clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"
)

func main() {
	// import clouderrorreportingpb "google.golang.org/genproto/googleapis/devtools/clouderrorreporting/v1beta1"

	ctx := context.Background()
	c, err := errorreporting.NewErrorStatsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &clouderrorreportingpb.DeleteEventsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.DeleteEvents(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
// [END clouderrorreporting_v1beta1_generated_ErrorStatsService_DeleteEvents_sync]
