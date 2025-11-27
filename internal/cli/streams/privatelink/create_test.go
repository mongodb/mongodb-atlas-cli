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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312010/admin"
	"go.uber.org/mock/gomock"
)

const fileName = "test-privateLink.json"

func TestCreateOpts_Run(t *testing.T) {
	testCases := []struct {
		name         string
		fileContents string
		wantErr      require.ErrorAssertionFunc
	}{
		{
			name:         "no file passed in",
			fileContents: "",
			wantErr:      require.Error,
		},
		{
			name: "file does not contain a provider",
			fileContents: `
				{
					"region": "US_EAST_2",
					"serviceEndpointId": "/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace",
					"dnsDomain": "test-namespace.servicebus.windows.net"
				}
			`,
			wantErr: func(tt require.TestingT, err error, _ ...any) {
				require.ErrorContains(tt, err, "provider missing")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			fs := afero.NewMemMapFs()

			if tc.fileContents != "" {
				require.NoError(t, afero.WriteFile(fs, fileName, []byte(tc.fileContents), 0600))
			}

			createOpts := &CreateOpts{
				fs:       fs,
				filename: fileName,
			}

			tc.wantErr(t, createOpts.Run())
		})
	}

	validPrivateLinkConfigFileContents := `
		{
			"provider": "Azure",
			"region": "US_EAST_2",
			"serviceEndpointId": "/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace",
			"dnsDomain": "test-namespace.servicebus.windows.net"
		}
	`

	t.Run("should call the store create privateLink method with the correct parameters", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		require.NoError(t, afero.WriteFile(fs, fileName, []byte(validPrivateLinkConfigFileContents), 0600))

		ctrl := gomock.NewController(t)
		mockStore := NewMockCreator(ctrl)

		createOpts := &CreateOpts{
			store:    mockStore,
			fs:       fs,
			filename: fileName,
			provider: "AZURE",
		}

		expected := atlasv2.NewStreamsPrivateLinkConnection(createOpts.provider)
		expected.SetProvider("Azure")
		expected.SetRegion("US_EAST_2")
		expected.SetServiceEndpointId("/subscriptions/fd01adff-b37e-4693-8497-83ecf183a145/resourceGroups/test-rg/providers/Microsoft.EventHub/namespaces/test-namespace")
		expected.SetDnsDomain("test-namespace.servicebus.windows.net")

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
		mockStore := NewMockCreator(ctrl)

		buf := new(bytes.Buffer)
		createOpts := &CreateOpts{
			store:    mockStore,
			fs:       fs,
			filename: fileName,
			OutputOpts: cli.OutputOpts{
				Template:  createTemplate,
				OutWriter: buf,
			},
			provider: "AWS",
		}

		expectedInterfaceEndpointID := "vpc-1234567890abcdef0"

		expected := atlasv2.NewStreamsPrivateLinkConnection(createOpts.provider)
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
