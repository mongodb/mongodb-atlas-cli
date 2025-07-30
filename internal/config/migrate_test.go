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

//go:build unit

package config

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func Test_SetAuthTypes(t *testing.T) {
	tests := []struct {
		name             string
		setupExpect      func(mockStore *MockStore)
		setupProfile     func(p *Profile)
		expectedAuthType AuthMechanism
	}{
		{
			name: "API Keys",
			setupExpect: func(mockStore *MockStore) {
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
			setupProfile: func(p *Profile) {
				p.SetPublicAPIKey("public")
				p.SetPrivateAPIKey("private")
			},
			expectedAuthType: APIKeys,
		},
		{
			name: "User Account",
			setupExpect: func(mockStore *MockStore) {
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
			setupProfile: func(p *Profile) {
				p.SetAccessToken("token")
				p.SetRefreshToken("token")
			},
			expectedAuthType: UserAccount,
		},
		{
			name: "Service Account",
			setupExpect: func(mockStore *MockStore) {
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
			setupProfile: func(p *Profile) {
				p.SetClientID("id")
				p.SetClientSecret("secret")
			},
			expectedAuthType: ServiceAccount,
		},
		{
			name: "Empty Profile",
			setupExpect: func(mockStore *MockStore) {
				mockStore.EXPECT().
					GetProfileNames().
					Return([]string{"test"}).
					Times(1)
				mockStore.EXPECT().
					GetHierarchicalValue("test", "auth_type").
					Return("").
					AnyTimes()
			},
			setupProfile:     func(*Profile) {},
			expectedAuthType: AuthMechanism(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := NewMockStore(ctrl)
			tt.setupExpect(mockStore)

			p := NewProfile("test", mockStore)
			tt.setupProfile(p)
			setAuthTypes(mockStore, func(*Profile) AuthMechanism {
				return tt.expectedAuthType
			})
			require.Equal(t, tt.expectedAuthType, p.AuthType())
		})
	}
}

func Test_GetAuthType(t *testing.T) {
	tests := []struct {
		name             string
		setup            func(p *Profile)
		expectedAuthType AuthMechanism
	}{
		{
			name: "API Keys",
			setup: func(p *Profile) {
				p.SetPublicAPIKey("public")
				p.SetPrivateAPIKey("private")
			},
			expectedAuthType: APIKeys,
		},
		{
			name: "User Account",
			setup: func(p *Profile) {
				p.SetAccessToken("token")
				p.SetRefreshToken("refresh")
			},
			expectedAuthType: UserAccount,
		},
		{
			name: "Service Account",
			setup: func(p *Profile) {
				p.SetClientID("id")
				p.SetClientSecret("secret")
			},
			expectedAuthType: ServiceAccount,
		},
		{
			name:             "Empty Profile",
			setup:            func(*Profile) {},
			expectedAuthType: AuthMechanism(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := NewInMemoryStore()
			p := NewProfile("test", store)
			tt.setup(p)
			require.Equal(t, tt.expectedAuthType, getAuthType(p))
		})
	}
}
