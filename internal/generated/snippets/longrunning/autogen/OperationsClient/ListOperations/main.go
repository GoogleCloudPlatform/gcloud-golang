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

// [START longrunning_generated_longrunning_autogen_OperationsClient_ListOperations]

package main

import (
	"context"

	longrunning "cloud.google.com/go/longrunning/autogen"
	"google.golang.org/api/iterator"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func main() {
	// import longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := longrunning.NewOperationsClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &longrunningpb.ListOperationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListOperations(ctx, req)
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

// [END longrunning_generated_longrunning_autogen_OperationsClient_ListOperations]
