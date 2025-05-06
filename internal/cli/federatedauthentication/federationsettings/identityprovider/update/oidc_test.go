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

package update

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312002/admin"
	"go.uber.org/mock/gomock"
)

func TestOidcUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockUpdater(ctrl)

	expected := &atlasv2.FederationIdentityProvider{}

	buf := new(bytes.Buffer)
	updateOpts := &OidcOpts{
		associatedDomains:    []string{"test"},
		federationSettingsID: "id",
		audience:             "audience",
		clientID:             "cliendId",
		authorizationType:    "auth",
		description:          "desc",
		displayName:          "name",
		idpType:              "type",
		issuerURI:            "uri",
		protocol:             "oidc",
		groupsClaim:          "groups",
		userClaim:            "user",
		requestedScopes:      []string{"scope"},
		store:                mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  updateTemplate,
			OutWriter: buf,
		}}

	mockStore.
		EXPECT().
		UpdateIdentityProvider(gomock.Any()).Return(expected, nil).
		Times(1)

	if err := updateOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, updateTemplate, *expected)
}
