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

// [START monitoring_v3_generated_ServiceMonitoringService_ListServices_sync]

package main

import (
	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"context"
	"google.golang.org/api/iterator"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

func main() {
	// import monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := monitoring.NewServiceMonitoringClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &monitoringpb.ListServicesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListServices(ctx, req)
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
// [END monitoring_v3_generated_ServiceMonitoringService_ListServices_sync]
