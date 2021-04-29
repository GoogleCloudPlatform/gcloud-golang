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

// [START firestore_v1_generated_Firestore_DeleteDocument_sync]

package main

import (
	firestore "cloud.google.com/go/firestore/apiv1"
	"context"
	firestorepb "google.golang.org/genproto/googleapis/firestore/v1"
)

func main() {
	ctx := context.Background()
	c, err := firestore.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &firestorepb.DeleteDocumentRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteDocument(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
// [END firestore_v1_generated_Firestore_DeleteDocument_sync]
