// Copyright 2021 MongoDB Inc
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

package aws

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestDelete_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockPrivateEndpointDeleter(ctrl)

	deleteOpts := &DeleteOpts{
		DeleteOpts: &cli.DeleteOpts{
			Entry:   "to_delete",
			Confirm: true,
		},
		store: mockStore,
	}

	mockStore.
		EXPECT().
		DeletePrivateEndpoint(deleteOpts.ProjectID, provider, deleteOpts.Entry).
		Return(nil).
		Times(1)

	err := deleteOpts.Run()
	require.NoError(t, err)
}
