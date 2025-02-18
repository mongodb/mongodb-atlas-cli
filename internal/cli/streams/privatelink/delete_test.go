// Copyright 2025 MongoDB Inc
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

package privatelink

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/stretchr/testify/require"
)

func TestDeleteOpts_Run(t *testing.T) {
	t.Run("should call the store delete privateLink method with the correct parameters", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := mocks.NewMockPrivateLinkDeleter(ctrl)

		const projectID = "a-project-id"
		const connectionID = "the-connection-id"

		deleteOpts := &DeleteOpts{
			store: mockStore,
			ProjectOpts: cli.ProjectOpts{
				ProjectID: projectID,
			},
			DeleteOpts: &cli.DeleteOpts{
				Confirm: true,
				Entry:   connectionID,
			},
		}

		mockStore.
			EXPECT().
			DeletePrivateLinkEndpoint(gomock.Eq(projectID), gomock.Eq(connectionID)).
			Times(1)

		require.NoError(t, deleteOpts.Run())
	})

	t.Run("should delete without error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := mocks.NewMockPrivateLinkDeleter(ctrl)

		deleteOpts := &DeleteOpts{
			DeleteOpts: &cli.DeleteOpts{
				Entry:   "some-connection-id",
				Confirm: true,
			},
			store: mockStore,
		}

		mockStore.
			EXPECT().
			DeletePrivateLinkEndpoint(gomock.Any(), gomock.Any()).
			Return(nil).
			Times(1)

		err := deleteOpts.Run()

		require.NoError(t, err)
	})
}
