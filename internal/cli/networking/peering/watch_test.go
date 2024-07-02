// Copyright 2020 MongoDB Inc
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

//go:build unit

package peering

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestWatchBuilder(t *testing.T) {
	test.CmdValidator(t, WatchBuilder(), 0, []string{flag.ProjectID, flag.Output})
}

func TestWatchOpts_Run(t *testing.T) {
	tests := []struct {
		name     string
		expected *atlasv2.BaseNetworkPeeringConnectionSettings
	}{
		{
			name: "AWS",
			expected: &atlasv2.BaseNetworkPeeringConnectionSettings{
				ProviderName: pointer.Get("AWS"),
				StatusName:   pointer.Get("PENDING_ACCEPTANCE"),
			},
		},
		{
			name: "AZURE",
			expected: &atlasv2.BaseNetworkPeeringConnectionSettings{
				ProviderName: pointer.Get("AZURE"),
				Status:       pointer.Get("AVAILABLE")},
		},
		{
			name: "GCP",
			expected: &atlasv2.BaseNetworkPeeringConnectionSettings{
				ProviderName: pointer.Get("GCP"),
				Status:       pointer.Get("WAITING_FOR_USER")},
		},
	}
	for _, tt := range tests {
		expected := tt.expected
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			mockStore := mocks.NewMockPeeringConnectionDescriber(ctrl)

			describeOpts := &WatchOpts{
				id:    "test",
				store: mockStore,
			}
			mockStore.
				EXPECT().
				PeeringConnection(describeOpts.ProjectID, describeOpts.id).
				Return(expected, nil).
				Times(1)

			if err := describeOpts.Run(); err != nil {
				t.Fatalf("Run() unexpected error: %v", err)
			}
		})
	}
}
