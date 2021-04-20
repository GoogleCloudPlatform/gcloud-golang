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

// [START pubsub_generated_pubsub_Subscription_Update_pushConfigAuthenticationMethod]

package main

import (
	"context"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "project-id")
	if err != nil {
		// TODO: Handle error.
	}
	sub := client.Subscription("subName")
	subConfig, err := sub.Update(ctx, pubsub.SubscriptionConfigToUpdate{
		PushConfig: &pubsub.PushConfig{
			Endpoint: "https://example.com/push",
			AuthenticationMethod: &pubsub.OIDCToken{
				ServiceAccountEmail: "service-account-email",
				Audience:            "client-12345",
			},
		},
	})
	if err != nil {
		// TODO: Handle error.
	}
	_ = subConfig // TODO: Use SubscriptionConfig.
}

// [END pubsub_generated_pubsub_Subscription_Update_pushConfigAuthenticationMethod]
