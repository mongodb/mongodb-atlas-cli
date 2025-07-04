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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

// X509Configuration retrieves the current user-managed certificates for a database user.
func (s *Store) X509Configuration(projectID string) (*atlasv2.UserSecurity, error) {
	result, _, err := s.clientv2.LDAPConfigurationApi.GetLdapConfiguration(s.ctx, projectID).Execute()
	return result, err
}

// SaveX509Configuration saves a customer-managed X.509 configuration for an Atlas project.
func (s *Store) SaveX509Configuration(projectID, certificate string) (*atlasv2.UserSecurity, error) {
	userCertificate := atlasv2.UserSecurity{
		CustomerX509: &atlasv2.DBUserTLSX509Settings{
			Cas: &certificate,
		},
	}
	result, _, err := s.clientv2.LDAPConfigurationApi.SaveLdapConfiguration(s.ctx, projectID, &userCertificate).Execute()
	return result, err
}

// DisableX509Configuration disables customer-managed X.509 configuration for an Atlas project.
func (s *Store) DisableX509Configuration(projectID string) error {
	_, _, err := s.clientv2.X509AuthenticationApi.DisableCustomerManagedX509(s.ctx, projectID).Execute()
	return err
}
