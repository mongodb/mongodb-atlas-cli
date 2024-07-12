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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIdentityProviderLister(ctrl)
	buf := new(bytes.Buffer)

	ListOpts := &ListOpts{
		store:                mockStore,
		federationSettingsID: "federationSettingsID",
		idpType:              "workforce",
		protocol:             oidc,
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
		ListOpts: &cli.ListOpts{},
	}

	displayName := "listName"
	displayName2 := "listName2"
	issuerURI := "uri"
	clientID := "id"
	idpType := "WORKLOAD"

	expected := &atlasv2.PaginatedFederationIdentityProvider{
		Links: &[]atlasv2.Link{
			{
				Rel:  pointer.Get("test"),
				Href: pointer.Get("test"),
			},
		},
		Results: &[]atlasv2.FederationIdentityProvider{
			{
				DisplayName: &displayName,
				IssuerUri:   &issuerURI,
				ClientId:    &clientID,
				IdpType:     &idpType,
			},
			{
				DisplayName: &displayName2,
				IssuerUri:   &issuerURI,
				ClientId:    &clientID,
				IdpType:     &idpType,
			},
		},
		TotalCount: pointer.Get(2),
	}

	mockStore.
		EXPECT().
		IdentityProviders(gomock.Any()).
		Return(expected, nil).
		Times(1)

	if err := ListOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, listTemplate, *expected)
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.Output, flag.FederationSettingsID, flag.IdpType, flag.Page, flag.Limit, flag.Protocol},
	)
}
