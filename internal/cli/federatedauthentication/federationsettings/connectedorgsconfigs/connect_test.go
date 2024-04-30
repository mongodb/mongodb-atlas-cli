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

//go:build unit

package connectedorgsconfigs

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115012/admin"
)

func TestConnect_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockConnectedOrgConfigsUpdater(ctrl)
	describeStore := mocks.NewMockConnectedOrgConfigsDescriber(ctrl)

	ConnectOpts := &ConnectOpts{
		store:                mockStore,
		federationSettingsID: "federationSettingsID",
		identityProviderID:   "id",
		protocol:             oidc,
		DescribeOrgConfigsOpts: &DescribeOrgConfigsOpts{
			describeStore: describeStore,
		},
	}

	mockStore.
		EXPECT().
		UpdateConnectedOrgConfig(gomock.Any()).
		Return(&atlasv2.ConnectedOrgConfig{}, nil).
		Times(1)

	describeStore.
		EXPECT().
		GetConnectedOrgConfig(gomock.Any()).
		Return(&atlasv2.ConnectedOrgConfig{}, nil).
		Times(1)

	if err := ConnectOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestConnectBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ConnectBuilder(),
		0,
		[]string{flag.Output, flag.FederationSettingsID, flag.IdentityProviderID, flag.Protocol, flag.OrgID},
	)
}
