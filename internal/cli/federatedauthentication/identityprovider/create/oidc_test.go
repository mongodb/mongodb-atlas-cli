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

package create

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115010/admin"
)

func TestOidcBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		OIDCBuilder(),
		0,
		[]string{flag.FederationSettingsID, flag.IdpType, flag.Audience, flag.AuthorizationType, flag.ClientID, flag.Description, flag.GroupsClaim, flag.UserClaim, flag.IssuerURI, flag.AssociatedDomains, flag.RequestedScopes, flag.Output},
	)
}
func TestOidcCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIdentityProviderCreator(ctrl)

	expected := &atlasv2.FederationOidcIdentityProvider{}

	createOpts := &OidcOpts{
		AssociatedDomains:    []string{"test"},
		FederationSettingsID: "id",
		Audience:             "audience",
		ClientID:             "cliendId",
		AuthorizationType:    "auth",
		Description:          "desc",
		DisplayName:          "name",
		IdpType:              "type",
		IssuerURI:            "uri",
		Protocol:             "oidc",
		GroupsClaim:          "groups",
		UserClaim:            "user",
		RequestedScopes:      []string{"scope"},
		store:                mockStore,
	}

	mockStore.
		EXPECT().
		CreateIdentityProvider(createOpts.newIdentityProvider()).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}
