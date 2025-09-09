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

package config

import (
	"testing"

	"github.com/mongodb/atlas-cli-core/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name     string
		isSecure bool
	}{
		{
			name:     "secure store redaction",
			isSecure: true,
		},
		{
			name:     "config file redaction",
			isSecure: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			profileName := "test"
			mockStore := mocks.NewMockStore(ctrl)

			// Set up mock expectations
			mockStore.EXPECT().GetProfileStringMap(profileName).Return(map[string]string{
				"project_id":     "proj123",
				"public_api_key": "public-key",
			}).Times(1)
			mockStore.EXPECT().GetProfileValue(profileName, "public_api_key").Return("pub-key").Times(1)
			mockStore.EXPECT().GetProfileValue(profileName, "private_api_key").Return("priv-key").Times(1)
			mockStore.EXPECT().GetProfileValue(profileName, "access_token").Return("").Times(1)
			mockStore.EXPECT().GetProfileValue(profileName, "refresh_token").Return("").Times(1)
			mockStore.EXPECT().GetProfileValue(profileName, "client_id").Return("").Times(1)
			mockStore.EXPECT().GetProfileValue(profileName, "client_secret").Return("").Times(1)

			mockStore.EXPECT().IsSecure().Return(tt.isSecure).Times(1)

			opts := &describeOpts{}
			result, err := opts.GetConfig(mockStore, profileName)
			require.NoError(t, err)
			assert.Equal(t, "proj123", result["project_id"])
			if tt.isSecure {
				assert.Contains(t, result["public_api_key"], redactedSecureText)
				assert.Contains(t, result["private_api_key"], redactedSecureText)
			} else {
				assert.Contains(t, result["public_api_key"], redactedConfigText)
				assert.Contains(t, result["private_api_key"], redactedConfigText)
			}
		})
	}
}
