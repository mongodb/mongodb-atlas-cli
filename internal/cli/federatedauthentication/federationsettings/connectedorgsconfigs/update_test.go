// Copyright 2024 MongoDB Inc
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

package connectedorgsconfigs

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312008/admin"
	"go.uber.org/mock/gomock"
)

func TestUpdate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockConnectedOrgConfigsUpdater(ctrl)

	updateOpts := &UpdateOpts{
		store:                mockStore,
		federationSettingsID: "federationSettingsID",
		file:                 "config.json",
		fs:                   afero.NewMemMapFs(),
		OrgOpts: cli.OrgOpts{
			OrgID: "6627f3ee0c9eba75f37240b3",
		},
	}

	fileContents := `
	{
		"domainAllowList": ["test.com"],
		"domainRestrictionEnabled": false,
		"orgId": "6627f3ee0c9eba75f37240b3",
		"roleMappings": []
	}`
	require.NoError(t, afero.WriteFile(updateOpts.fs, updateOpts.file, []byte(fileContents), 0600))

	domains := []string{"test.com"}
	expected := &atlasv2.ConnectedOrgConfig{
		OrgId:                    "6627f3ee0c9eba75f37240b3",
		DomainAllowList:          &domains,
		DomainRestrictionEnabled: false,
	}

	mockStore.
		EXPECT().
		UpdateConnectedOrgConfig(gomock.Any()).
		Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, updateTemplate, expected)
}
