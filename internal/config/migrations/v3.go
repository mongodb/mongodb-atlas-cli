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
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
)

// NewMigrateToVersion3 renames profiles using the pre-Authorization Code
// user_account mechanism to user_account_legacy, in lockstep with
// atlas-cli-core's config version 3. Idempotent: already renamed profiles
// are left untouched.
func NewMigrateToVersion3() MigrationFunc {
	return func(dependencies MigrationDependencies) error {
		insecureStore, err := dependencies.GetInsecureStore()
		if err != nil {
			return fmt.Errorf("failed to get store: %w", err)
		}

		renameUserAccountProfiles(insecureStore)

		if err := insecureStore.Save(); err != nil {
			return fmt.Errorf("failed to save store: %w", err)
		}

		return nil
	}
}

func renameUserAccountProfiles(store config.Store) {
	for _, name := range store.GetProfileNames() {
		profile := config.NewProfile(name, store)
		if profile.AuthType() == config.UserAccount {
			profile.SetAuthType(config.UserAccountLegacy)
		}
	}
}
