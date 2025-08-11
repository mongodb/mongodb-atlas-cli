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

const ConfigVersion2 = 2

// MigrateVersions migrates the profile to the latest version.
// This function can be expanded to support future migrations.
func MigrateVersions(store Store) error {
	if GetVersion() >= ConfigVersion2 {
		return nil
	}

	setAuthTypes(store, getAuthType)

	// TODO: Remaining migration steps to move credentials to secure storage will be done as a part of CLOUDP-329802
	SetVersion(ConfigVersion2)

	return Save()
}

// setAuthTypes sets the auth type for each profile based on the credentials available.
// Sets NoAuth for profiles without authentication.
func setAuthTypes(store Store, getAuthType func(*Profile) AuthMechanism) {
	profileNames := store.GetProfileNames()
	for _, name := range profileNames {
		profile := NewProfile(name, store)
		authType := getAuthType(profile)
		// Always set the auth type, including NoAuth for profiles without authentication
		profile.SetAuthType(authType)
	}
}

func getAuthType(profile *Profile) AuthMechanism {
	switch {
	case profile.PublicAPIKey() != "" && profile.PrivateAPIKey() != "":
		return APIKeys
	case profile.AccessToken() != "" && profile.RefreshToken() != "":
		return UserAccount
	case profile.ClientID() != "" && profile.ClientSecret() != "":
		return ServiceAccount
	}
	return NoAuth // Profile has no authentication configured
}
