// Copyright 2022 MongoDB Inc
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

package gcp

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockPrivateEndpointLister(ctrl)

	var expected []admin.EndpointService

	listOpts := &ListOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		PrivateEndpoints(listOpts.ProjectID, provider).
		Return(expected, nil).
		Times(1)

	err := listOpts.Run()
	require.NoError(t, err)

	test.VerifyOutputTemplate(t, listTemplate, expected)
}
