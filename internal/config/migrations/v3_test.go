// Copyright 2026 MongoDB Inc
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

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/stretchr/testify/require"
)

func Test_MigrateToVersion3(t *testing.T) {
	tests := []struct {
		name             string
		setupProfile     func(p *config.Profile)
		expectedAuthType config.AuthMechanism
	}{
		{
			name: "User Account",
			setupProfile: func(p *config.Profile) {
				p.SetAuthType(config.UserAccount)
			},
			expectedAuthType: config.UserAccountLegacy,
		},
		{
			name: "already migrated",
			setupProfile: func(p *config.Profile) {
				p.SetAuthType(config.UserAccountLegacy)
			},
			expectedAuthType: config.UserAccountLegacy,
		},
		{
			name: "Service Account untouched",
			setupProfile: func(p *config.Profile) {
				p.SetAuthType(config.ServiceAccount)
			},
			expectedAuthType: config.ServiceAccount,
		},
		{
			name: "API Keys untouched",
			setupProfile: func(p *config.Profile) {
				p.SetAuthType(config.APIKeys)
			},
			expectedAuthType: config.APIKeys,
		},
		{
			name:             "empty profile untouched",
			setupProfile:     func(*config.Profile) {},
			expectedAuthType: config.AuthMechanism(""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := config.NewInMemoryStore()
			p := config.NewProfile("test", store)
			tt.setupProfile(p)

			renameUserAccountProfiles(store)

			require.Equal(t, tt.expectedAuthType, p.AuthType())
		})
	}
}

func Test_MigrateToVersion3_multipleProfiles(t *testing.T) {
	store := config.NewInMemoryStore()
	userAccount := config.NewProfile("userAccount", store)
	userAccount.SetAuthType(config.UserAccount)
	legacy := config.NewProfile("legacy", store)
	legacy.SetAuthType(config.UserAccountLegacy)
	apiKeys := config.NewProfile("apiKeys", store)
	apiKeys.SetAuthType(config.APIKeys)

	renameUserAccountProfiles(store)

	require.Equal(t, config.UserAccountLegacy, userAccount.AuthType())
	require.Equal(t, config.UserAccountLegacy, legacy.AuthType())
	require.Equal(t, config.APIKeys, apiKeys.AuthType())
}

func Test_MigrateToVersion3_idempotent(t *testing.T) {
	store := config.NewInMemoryStore()
	p := config.NewProfile("test", store)
	p.SetAuthType(config.UserAccount)

	renameUserAccountProfiles(store)
	renameUserAccountProfiles(store)

	require.Equal(t, config.UserAccountLegacy, p.AuthType())
}
