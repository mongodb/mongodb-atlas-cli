// Copyright 2023 MongoDB Inc
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

package apikeys

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestCreate_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProjectAPIKeyCreator(ctrl)

	createOpts := &CreateOpts{
		store:       mockStore,
		description: "desc",
		roles:       []string{},
	}
	buf := new(bytes.Buffer)
	require.NoError(t, createOpts.InitOutput(buf, createTemplate)())
	createOpts.ProjectID = "5a0a1e7e0f2912c554080adc"

	apiKey := &atlasv2.CreateAtlasProjectApiKey{
		Desc:  createOpts.description,
		Roles: []string{},
	}
	expected := &atlasv2.ApiKeyUserDetails{
		Id:         pointer.Get("id"),
		PublicKey:  pointer.Get("public"),
		PrivateKey: pointer.Get("private"),
	}

	mockStore.
		EXPECT().
		CreateProjectAPIKey(createOpts.ProjectID, apiKey).Return(expected, nil).
		Times(1)

	require.NoError(t, createOpts.Run())
	assert.Equal(t, `API Key 'id' created.
Public API Key public
Private API Key private
`, buf.String())
	t.Log("buf:", buf.String())
}

func TestCreateTemplate(t *testing.T) {
	test.VerifyOutputTemplate(t, createTemplate, atlasv2.ApiKeyUserDetails{})
}
