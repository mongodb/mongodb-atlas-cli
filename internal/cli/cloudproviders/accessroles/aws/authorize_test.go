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

package aws

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.uber.org/mock/gomock"
)

func TestAuthorizeTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, authorizeTemplate, atlasv2.CloudProviderAccessRole{})
}

func TestAuthorizeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockCloudProviderAccessRoleAuthorizer(ctrl)

	expected := &atlasv2.CloudProviderAccessRole{}

	opts := &AuthorizeOpts{
		store: mockStore,
	}

	mockStore.
		EXPECT().
		AuthorizeCloudProviderAccessRole(opts.ProjectID, opts.roleID, opts.newCloudProviderAuthorizationRequest()).
		Return(expected, nil).
		Times(1)
	require.NoError(t, opts.Run())
}
