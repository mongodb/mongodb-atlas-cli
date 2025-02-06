// Copyright 2025 MongoDB Inc
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

package privatelink

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

const fileName = "test-privateLink.json"

func TestCreateOpts_Run(t *testing.T) {
	validPrivateLinkConfigFileContents := `
		{
			"provider": "Azure",
			"region": "US_EAST_2",
			"serviceEndpointId": "/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace",
			"dnsDomain": "test-namespace.servicebus.windows.net"
		}
	`

	t.Run("should fail if no file is passed in", func(t *testing.T) {
		createOpts := &CreateOpts{
			fs: afero.NewMemMapFs(),
		}

		err := createOpts.Run()
		assert.Error(t, err)
	})

	t.Run("should fail if the file does not contain a provider", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		fileContents := `
			{
				"region": "US_EAST_2",
				"serviceEndpointId": "/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace",
				"dnsDomain": "test-namespace.servicebus.windows.net"
			}
		`
		require.NoError(t, afero.WriteFile(fs, fileName, []byte(fileContents), 0600))

		createOpts := &CreateOpts{
			fs:       fs,
			filename: fileName,
		}

		err := createOpts.Run()
		assert.ErrorContains(t, err, "provider missing")
	})

	t.Run("should call the store create privateLink method with the correct parameters", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs, fileName, []byte(validPrivateLinkConfigFileContents), 0600))

		ctrl := gomock.NewController(t)
		mockStore := mocks.NewMockPrivateLinkCreator(ctrl)

		createOpts := &CreateOpts{
			store:    mockStore,
			fs:       fs,
			filename: fileName,
		}

		expectedProvider := "Azure"
		expectedRegion := "US_EAST_2"
		expectedServiceEndpointID := "/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace"
		expectedDNSDomain := "test-namespace.servicebus.windows.net"

		expected := atlasv2.NewStreamsPrivateLinkConnection()
		expected.Provider = &expectedProvider
		expected.Region = &expectedRegion
		expected.ServiceEndpointId = &expectedServiceEndpointID
		expected.DnsDomain = &expectedDNSDomain

		mockStore.
			EXPECT().
			CreatePrivateLinkEndpoint(gomock.Eq(createOpts.ConfigProjectID()), gomock.Eq(expected)).
			Times(1)

		require.NoError(t, createOpts.Run())
	})

	t.Run("should print the result", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs, fileName, []byte(validPrivateLinkConfigFileContents), 0600))

		ctrl := gomock.NewController(t)
		mockStore := mocks.NewMockPrivateLinkCreator(ctrl)

		buf := new(bytes.Buffer)
		createOpts := &CreateOpts{
			store:    mockStore,
			fs:       fs,
			filename: fileName,
			OutputOpts: cli.OutputOpts{
				Template:  createTemplate,
				OutWriter: buf,
			},
		}

		expectedInterfaceEndpointID := "vpc-1234567890abcdef0"

		expected := atlasv2.NewStreamsPrivateLinkConnection()
		expected.InterfaceEndpointId = &expectedInterfaceEndpointID

		mockStore.
			EXPECT().
			// This test does not assert the parameters passed to the store method
			CreatePrivateLinkEndpoint(gomock.Any(), gomock.Any()).
			Return(expected, nil).
			Times(1)

		require.NoError(t, createOpts.Run())
		assert.Equal(t, "Atlas Stream Processing PrivateLink endpoint "+expectedInterfaceEndpointID+" created.\n", buf.String())
	})
}
