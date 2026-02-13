// Copyright 2020 MongoDB Inc
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

package store

import (
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

// VerifyLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) VerifyLDAPConfiguration(projectID string, ldap *atlasv2.LDAPVerifyConnectivityJobRequestParams) (*atlasv2.LDAPVerifyConnectivityJobRequest, error) {
	resp, _, err := s.clientv2.LDAPConfigurationApi.VerifyUserSecurityLdap(s.ctx, projectID, ldap).
		Execute()
	return resp, err
}

// GetStatusLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) GetStatusLDAPConfiguration(projectID, requestID string) (*atlasv2.LDAPVerifyConnectivityJobRequest, error) {
	resp, _, err := s.clientv2.LDAPConfigurationApi.GetUserSecurityVerify(s.ctx, projectID, requestID).Execute()
	return resp, err
}

// SaveLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) SaveLDAPConfiguration(projectID string, ldap *atlasv2.UserSecurity) (*atlasv2.UserSecurity, error) {
	resp, _, err := s.clientv2.LDAPConfigurationApi.UpdateUserSecurity(s.ctx, projectID, ldap).
		Execute()
	return resp, err
}

// DeleteLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteLDAPConfiguration(projectID string) error {
	_, _, err := s.clientv2.LDAPConfigurationApi.DeleteLdapUserMapping(s.ctx, projectID).Execute()
	return err
}

// GetLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) GetLDAPConfiguration(projectID string) (*atlasv2.UserSecurity, error) {
	resp, _, err := s.clientv2.LDAPConfigurationApi.GetUserSecurity(s.ctx, projectID).Execute()
	return resp, err
}
