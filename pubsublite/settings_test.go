// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and

package pubsublite

import (
	"testing"
	"time"
)

func TestValidatePublishSettings(t *testing.T) {
	for _, tc := range []struct {
		desc string
		// mutateSettings is passed a copy of DefaultPublishSettings to mutate.
		mutateSettings func(settings *PublishSettings)
		wantErr        bool
	}{
		{
			desc:           "valid: default",
			mutateSettings: func(settings *PublishSettings) {},
			wantErr:        false,
		},
		{
			desc: "valid: max",
			mutateSettings: func(settings *PublishSettings) {
				settings.CountThreshold = MaxPublishRequestCount
				settings.ByteThreshold = MaxPublishRequestBytes
				// These have no max bounds, check large values.
				settings.DelayThreshold = 24 * time.Hour
				settings.Timeout = 24 * time.Hour
				settings.BufferedByteLimit = 1e10
			},
			wantErr: false,
		},
		{
			desc: "invalid: zero CountThreshold",
			mutateSettings: func(settings *PublishSettings) {
				settings.CountThreshold = 0
			},
			wantErr: true,
		},
		{
			desc: "invalid: CountThreshold over MaxPublishRequestCount",
			mutateSettings: func(settings *PublishSettings) {
				settings.CountThreshold = MaxPublishRequestCount + 1
			},
			wantErr: true,
		},
		{
			desc: "invalid: ByteThreshold over MaxPublishRequestBytes",
			mutateSettings: func(settings *PublishSettings) {
				settings.ByteThreshold = MaxPublishRequestBytes + 1
			},
			wantErr: true,
		},
		{
			desc: "invalid: zero ByteThreshold",
			mutateSettings: func(settings *PublishSettings) {
				settings.ByteThreshold = 0
			},
			wantErr: true,
		},
		{
			desc: "invalid: zero DelayThreshold",
			mutateSettings: func(settings *PublishSettings) {
				settings.DelayThreshold = time.Duration(0)
			},
			wantErr: true,
		},
		{
			desc: "invalid: zero Timeout",
			mutateSettings: func(settings *PublishSettings) {
				settings.Timeout = time.Duration(0)
			},
			wantErr: true,
		},
		{
			desc: "invalid: zero BufferedByteLimit",
			mutateSettings: func(settings *PublishSettings) {
				settings.BufferedByteLimit = 0
			},
			wantErr: true,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			settings := DefaultPublishSettings
			tc.mutateSettings(&settings)

			gotErr := validatePublishSettings(settings)
			if (gotErr != nil) != tc.wantErr {
				t.Errorf("validatePublishSettings(%v) = %v, want err=%v", settings, gotErr, tc.wantErr)
			}
		})
	}
}
