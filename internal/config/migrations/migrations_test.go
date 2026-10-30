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
	"errors"
	"testing"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/mocks"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// The migrated version must match what atlas-cli-core accepts, or every command fails to load config.
func Test_DefaultMigrator_matchesCoreMaxSupportedVersion(t *testing.T) {
	require.Equal(t, config.MaxSupportedVersion, len(NewDefaultMigrator().migrations)+1)
}

type failingSaveStore struct {
	*config.InMemoryStore
}

func (failingSaveStore) Save() error { return errors.New("save failed") }

func newTestMigrator(t *testing.T, store config.Store) *Migrator {
	t.Helper()
	ctrl := gomock.NewController(t)
	secureStore := mocks.NewMockSecureStore(ctrl)
	secureStore.EXPECT().Available().Return(false).AnyTimes()
	return NewMigrator(MigrationDependencies{
		GetInsecureStore: func() (config.Store, error) { return store, nil },
		GetSecureStore:   func() (config.SecureStore, error) { return secureStore, nil },
	}, []MigrationFunc{NewMigrateToVersion2(), NewMigrateToVersion3()})
}

func Test_Migrate_userAccountToLegacy(t *testing.T) {
	tests := []struct {
		name    string
		version any
		setup   func(p *config.Profile)
	}{
		{
			name:    "version 1 with device-flow tokens",
			version: nil,
			setup: func(p *config.Profile) {
				p.SetAccessToken("token")
				p.SetRefreshToken("refresh")
			},
		},
		{
			name:    "version 2 user_account",
			version: int64(2),
			setup: func(p *config.Profile) {
				p.SetAuthType(config.UserAccount)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := config.NewInMemoryStore()
			if tt.version != nil {
				store.SetGlobalValue(VersionKey, tt.version)
			}
			p := config.NewProfile("test", store)
			tt.setup(p)

			require.NoError(t, newTestMigrator(t, store).Migrate())

			require.Equal(t, string(config.UserAccountLegacy), store.GetProfileValue("test", config.AuthTypeField))
			require.Equal(t, 3, store.GetGlobalValue(VersionKey))
		})
	}
}

func Test_Migrate_version3Untouched(t *testing.T) {
	store := config.NewInMemoryStore()
	store.SetGlobalValue(VersionKey, int64(3))
	config.NewProfile("test", store).SetAuthType(config.UserAccount)

	require.NoError(t, newTestMigrator(t, store).Migrate())

	require.Equal(t, string(config.UserAccount), store.GetProfileValue("test", config.AuthTypeField))
}

func Test_Migrate_saveFailureKeepsVersion(t *testing.T) {
	store := failingSaveStore{config.NewInMemoryStore()}
	store.SetGlobalValue(VersionKey, int64(2))
	config.NewProfile("test", store).SetAuthType(config.UserAccount)

	require.Error(t, newTestMigrator(t, store).Migrate())

	require.Equal(t, int64(2), store.GetGlobalValue(VersionKey))
}
