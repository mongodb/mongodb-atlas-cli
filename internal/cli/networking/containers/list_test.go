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

package containers

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
	"go.uber.org/mock/gomock"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockLister(ctrl)

	t.Run("no provider", func(t *testing.T) {
		buf := new(bytes.Buffer)
		listOpts := &ListOpts{
			store: mockStore,
			OutputOpts: cli.OutputOpts{
				Template:  listTemplate,
				OutWriter: buf,
			},
		}

		expected := []atlasv2.CloudProviderContainer{{
			Id:             pointer.Get("1234567890"),
			ProviderName:   pointer.Get("AWS"),
			Region:         pointer.Get("US_EAST_1"),
			AtlasCidrBlock: pointer.Get("Test"),
			Provisioned:    pointer.Get(false),
		}}

		mockStore.
			EXPECT().
			AllContainers(listOpts.ProjectID, listOpts.NewAtlasListOptions()).
			Return(expected, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		test.VerifyOutputTemplate(t, listTemplate, expected)
	})
	t.Run("with provider", func(t *testing.T) {
		listOpts := &ListOpts{
			store:    mockStore,
			provider: "test",
		}

		var expected []atlasv2.CloudProviderContainer
		mockStore.
			EXPECT().
			ContainersByProvider(listOpts.ProjectID, listOpts.newContainerListOptions()).
			Return(expected, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
	})
}
