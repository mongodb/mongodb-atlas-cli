// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package organizations

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestCreateAtlasBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		CreateAtlasBuilder(),
		0,
		[]string{
			flag.OwnerID,
			flag.APIKeyDescription,
			flag.APIKeyRole,
			flag.Output,
		},
	)
}

func TestCreateAtlasOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockOrganizationCreator(ctrl)

	federationSettings := "federationId"
	expected := &atlasv2.CreateOrganizationRequest{
		ApiKey:               nil,
		Name:                 "Org 0",
		OrgOwnerId:           nil,
		FederationSettingsId: &federationSettings,
	}
	resp := &atlasv2.CreateOrganizationResponse{}
	mockStore.
		EXPECT().
		CreateAtlasOrganization(expected).Return(resp, nil).
		Times(1)

	createOpts := &CreateAtlasOpts{
		store:                mockStore,
		name:                 "Org 0",
		federationSettingsID: federationSettings,
	}
	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
