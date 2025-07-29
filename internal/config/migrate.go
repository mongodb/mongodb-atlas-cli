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

const configVersion2 = 2

// MigrateVersions migrates the profile to the latest version.
// This function can be expanded to support future migrations.
func MigrateVersions() error { return Default().MigrateVersions() }
func (p *Profile) MigrateVersions() error {
	if p.GetVersion() >= configVersion2 {
		return nil
	}

	p.SetAuthTypes()

	// TODO: Remaining migration steps to move credentials to secure storage will be done as a part of CLOUDP-329802
	p.SetVersion(configVersion2)

	return p.Save()
}

// SetAuthTypes sets the auth type for each profile based on the credentials available.
// "not_logged_in" is used when no or incomplete credentials are present.
func (p *Profile) SetAuthTypes() {
	profileNames := p.configStore.GetProfileNames()
	for _, name := range profileNames {
		profile := NewProfile(name, p.configStore)
		switch {
		case profile.PublicAPIKey() != "" && profile.PrivateAPIKey() != "":
			profile.SetAuthType(APIKeys)
		case profile.AccessToken() != "" && profile.RefreshToken() != "":
			profile.SetAuthType(UserAccount)
		case profile.ClientID() != "" && profile.ClientSecret() != "":
			profile.SetAuthType(ServiceAccount)
		}
	}
}
