// Copyright 2024 MongoDB Inc
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312008/admin"
)

// IdentityProviders encapsulate the logic to manage different cloud providers.
func (s *Store) IdentityProviders(opts *atlasv2.ListIdentityProvidersApiParams) (*atlasv2.PaginatedFederationIdentityProvider, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.ListIdentityProvidersWithParams(s.ctx, opts).Execute()
	return result, err
}

// IdentityProvider encapsulate the logic to manage different cloud providers.
func (s *Store) IdentityProvider(opts *atlasv2.GetIdentityProviderApiParams) (*atlasv2.FederationIdentityProvider, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetIdentityProviderWithParams(s.ctx, opts).Execute()
	return result, err
}

// CreateIdentityProvider encapsulate the logic to manage different cloud providers.
func (s *Store) CreateIdentityProvider(opts *atlasv2.CreateIdentityProviderApiParams) (*atlasv2.FederationOidcIdentityProvider, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.CreateIdentityProviderWithParams(s.ctx, opts).Execute()
	return result, err
}

// DeleteIdentityProvider encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteIdentityProvider(federationSettingsID string, identityProviderID string) error {
	_, err := s.clientv2.FederatedAuthenticationApi.DeleteIdentityProvider(s.ctx, federationSettingsID, identityProviderID).Execute()
	return err
}

// UpdateIdentityProvider encapsulate the logic to manage different cloud providers.
func (s *Store) UpdateIdentityProvider(opts *atlasv2.UpdateIdentityProviderApiParams) (*atlasv2.FederationIdentityProvider, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.UpdateIdentityProviderWithParams(s.ctx, opts).Execute()
	return result, err
}

// RevokeJwksFromIdentityProvider encapsulate the logic to manage different cloud providers.
func (s *Store) RevokeJwksFromIdentityProvider(federationSettingsID string, identityProviderID string) error {
	_, err := s.clientv2.FederatedAuthenticationApi.RevokeIdentityProviderJwks(s.ctx, federationSettingsID, identityProviderID).Execute()
	return err
}
