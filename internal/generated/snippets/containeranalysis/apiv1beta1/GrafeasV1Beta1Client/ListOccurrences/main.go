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

// [START containeranalysis_v1beta1_generated_GrafeasV1Beta1_ListOccurrences_sync]

package main

import (
	"context"

	containeranalysis "cloud.google.com/go/containeranalysis/apiv1beta1"
	"google.golang.org/api/iterator"
	grafeaspb "google.golang.org/genproto/googleapis/devtools/containeranalysis/v1beta1/grafeas"
)

func main() {
	ctx := context.Background()
	c, err := containeranalysis.NewGrafeasV1Beta1Client(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &grafeaspb.ListOccurrencesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListOccurrences(ctx, req)
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

// [END containeranalysis_v1beta1_generated_GrafeasV1Beta1_ListOccurrences_sync]
