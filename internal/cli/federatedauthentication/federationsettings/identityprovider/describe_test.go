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

package identityprovider

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250219001/admin"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIdentityProviderDescriber(ctrl)
	buf := new(bytes.Buffer)

	describeOpts := &DescribeOpts{
		store:                mockStore,
		FederationSettingsID: "federationSettingsID",
		IdentityProviderID:   "id",
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	params := &atlasv2.GetIdentityProviderApiParams{
		FederationSettingsId: describeOpts.FederationSettingsID,
		IdentityProviderId:   describeOpts.IdentityProviderID,
	}

	displayName := "displayName"
	issuerURI := "issuerUri"
	clientID := "clientId"
	idpType := "WORKFORCE"

	expected := &atlasv2.FederationIdentityProvider{
		Id:          describeOpts.IdentityProviderID,
		DisplayName: &displayName,
		IssuerUri:   &issuerURI,
		ClientId:    &clientID,
		IdpType:     &idpType,
	}

	mockStore.
		EXPECT().
		IdentityProvider(params).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplate, *expected)
}
