// Copyright 2025 MongoDB Inc
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

package migrations

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_MigrateToVersion2(t *testing.T) {
	tests := []struct {
		name             string
		setupExpect      func(mockStore *config.MockStore)
		setupProfile     func(p *config.Profile)
		expectedAuthType config.AuthMechanism
	}{
		{
			name: "API Keys",
			setupExpect: func(mockStore *config.MockStore) {
				mockStore.EXPECT().
					GetProfileNames().
					Return([]string{"test"}).
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "public_api_key", "public").
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "private_api_key", "private").
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "auth_type", "api_keys").
					Times(1)
				mockStore.EXPECT().
					GetHierarchicalValue("test", "auth_type").
					Return("api_keys").
					AnyTimes()
			},
			setupProfile: func(p *config.Profile) {
				p.SetPublicAPIKey("public")
				p.SetPrivateAPIKey("private")
			},
			expectedAuthType: config.APIKeys,
		},
		{
			name: "User Account",
			setupExpect: func(mockStore *config.MockStore) {
				mockStore.EXPECT().
					GetProfileNames().
					Return([]string{"test"}).
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "access_token", "token").
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "refresh_token", "token").
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "auth_type", "user_account").
					Times(1)
				mockStore.EXPECT().
					GetHierarchicalValue("test", "auth_type").
					Return("user_account").
					AnyTimes()
			},
			setupProfile: func(p *config.Profile) {
				p.SetAccessToken("token")
				p.SetRefreshToken("token")
			},
			expectedAuthType: config.UserAccount,
		},
		{
			name: "Service Account",
			setupExpect: func(mockStore *config.MockStore) {
				mockStore.EXPECT().
					GetProfileNames().
					Return([]string{"test"}).
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "client_id", "id").
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "client_secret", "secret").
					Times(1)
				mockStore.EXPECT().
					SetProfileValue("test", "auth_type", "service_account").
					Times(1)
				mockStore.EXPECT().
					GetHierarchicalValue("test", "auth_type").
					Return("service_account").
					AnyTimes()
			},
			setupProfile: func(p *config.Profile) {
				p.SetClientID("id")
				p.SetClientSecret("secret")
			},
			expectedAuthType: config.ServiceAccount,
		},
		{
			name: "Empty Profile",
			setupExpect: func(mockStore *config.MockStore) {
				mockStore.EXPECT().
					GetProfileNames().
					Return([]string{"test"}).
					Times(1)
				mockStore.EXPECT().
					GetHierarchicalValue("test", "auth_type").
					Return("").
					AnyTimes()
			},
			setupProfile:     func(*config.Profile) {},
			expectedAuthType: config.AuthMechanism(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := config.NewMockStore(ctrl)
			tt.setupExpect(mockStore)

			p := config.NewProfile("test", mockStore)
			tt.setupProfile(p)
			setAuthTypes(mockStore, func(*config.Profile) config.AuthMechanism {
				return tt.expectedAuthType
			})
			require.Equal(t, tt.expectedAuthType, p.AuthType())
		})
	}
}

func Test_GetAuthType(t *testing.T) {
	tests := []struct {
		name             string
		setup            func(p *config.Profile)
		expectedAuthType config.AuthMechanism
	}{
		{
			name: "API Keys",
			setup: func(p *config.Profile) {
				p.SetPublicAPIKey("public")
				p.SetPrivateAPIKey("private")
			},
			expectedAuthType: config.APIKeys,
		},
		{
			name: "User Account",
			setup: func(p *config.Profile) {
				p.SetAccessToken("token")
				p.SetRefreshToken("refresh")
			},
			expectedAuthType: config.UserAccount,
		},
		{
			name: "Service Account",
			setup: func(p *config.Profile) {
				p.SetClientID("id")
				p.SetClientSecret("secret")
			},
			expectedAuthType: config.ServiceAccount,
		},
		{
			name:             "Empty Profile",
			setup:            func(*config.Profile) {},
			expectedAuthType: config.AuthMechanism(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := config.NewInMemoryStore()
			p := config.NewProfile("test", store)
			tt.setup(p)
			require.Equal(t, tt.expectedAuthType, getAuthType(p))
		})
	}
}

func Test_MigrateSecrets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock stores
	mockInsecureStore := config.NewMockStore(ctrl)
	mockSecureStore := config.NewMockSecureStore(ctrl)

	// Define test profiles
	profileNames := []string{"profile1", "profile2"}

	// Setup expectations for GetProfileNames
	mockInsecureStore.EXPECT().GetProfileNames().Return(profileNames)

	// Note: This test verifies that migrateSecrets only processes properties in config.SecureProperties.
	// If profiles had dummy properties, they would be ignored completely.
	// We do NOT set up mock expectations for such properties because the function should never
	// attempt to access them. If it did, the test would fail with "unexpected call" errors.

	// Setup mock expectations for profile1 - all secure properties
	// public_api_key
	mockInsecureStore.EXPECT().GetProfileValue("profile1", "public_api_key").Return("public1")
	mockSecureStore.EXPECT().Set("profile1", "public_api_key", "public1")
	mockInsecureStore.EXPECT().SetProfileValue("profile1", "public_api_key", nil)

	// private_api_key
	mockInsecureStore.EXPECT().GetProfileValue("profile1", "private_api_key").Return("private1")
	mockSecureStore.EXPECT().Set("profile1", "private_api_key", "private1")
	mockInsecureStore.EXPECT().SetProfileValue("profile1", "private_api_key", nil)

	// access_token
	mockInsecureStore.EXPECT().GetProfileValue("profile1", "access_token").Return("access1")
	mockSecureStore.EXPECT().Set("profile1", "access_token", "access1")
	mockInsecureStore.EXPECT().SetProfileValue("profile1", "access_token", nil)

	// refresh_token
	mockInsecureStore.EXPECT().GetProfileValue("profile1", "refresh_token").Return("refresh1")
	mockSecureStore.EXPECT().Set("profile1", "refresh_token", "refresh1")
	mockInsecureStore.EXPECT().SetProfileValue("profile1", "refresh_token", nil)

	// client_id
	mockInsecureStore.EXPECT().GetProfileValue("profile1", "client_id").Return("client1")
	mockSecureStore.EXPECT().Set("profile1", "client_id", "client1")
	mockInsecureStore.EXPECT().SetProfileValue("profile1", "client_id", nil)

	// client_secret
	mockInsecureStore.EXPECT().GetProfileValue("profile1", "client_secret").Return("secret1")
	mockSecureStore.EXPECT().Set("profile1", "client_secret", "secret1")
	mockInsecureStore.EXPECT().SetProfileValue("profile1", "client_secret", nil)

	// Setup mock expectations for profile2 - all secure properties
	// public_api_key
	mockInsecureStore.EXPECT().GetProfileValue("profile2", "public_api_key").Return("public2")
	mockSecureStore.EXPECT().Set("profile2", "public_api_key", "public2")
	mockInsecureStore.EXPECT().SetProfileValue("profile2", "public_api_key", nil)

	// private_api_key
	mockInsecureStore.EXPECT().GetProfileValue("profile2", "private_api_key").Return("private2")
	mockSecureStore.EXPECT().Set("profile2", "private_api_key", "private2")
	mockInsecureStore.EXPECT().SetProfileValue("profile2", "private_api_key", nil)

	// access_token
	mockInsecureStore.EXPECT().GetProfileValue("profile2", "access_token").Return("access2")
	mockSecureStore.EXPECT().Set("profile2", "access_token", "access2")
	mockInsecureStore.EXPECT().SetProfileValue("profile2", "access_token", nil)

	// refresh_token
	mockInsecureStore.EXPECT().GetProfileValue("profile2", "refresh_token").Return("refresh2")
	mockSecureStore.EXPECT().Set("profile2", "refresh_token", "refresh2")
	mockInsecureStore.EXPECT().SetProfileValue("profile2", "refresh_token", nil)

	// client_id
	mockInsecureStore.EXPECT().GetProfileValue("profile2", "client_id").Return("client2")
	mockSecureStore.EXPECT().Set("profile2", "client_id", "client2")
	mockInsecureStore.EXPECT().SetProfileValue("profile2", "client_id", nil)

	// client_secret
	mockInsecureStore.EXPECT().GetProfileValue("profile2", "client_secret").Return("secret2")
	mockSecureStore.EXPECT().Set("profile2", "client_secret", "secret2")
	mockInsecureStore.EXPECT().SetProfileValue("profile2", "client_secret", nil)

	// Call the function under test
	migrateSecrets(mockInsecureStore, mockSecureStore)
}
