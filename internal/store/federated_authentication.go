// Copyright 2023 MongoDB Inc
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

//go:generate mockgen -destination=../mocks/mock_federated_authentication.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store FederationAuthenticationStore

type FederationAuthenticationStore interface {
	AtlasFederatedAuthOrgConfig(opts *atlasv2.GetConnectedOrgConfigApiParams) (*atlasv2.ConnectedOrgConfig, error)
	AtlasIdentityProvider(ops *atlasv2.GetIdentityProviderApiParams) (*atlasv2.FederationIdentityProvider, error)
}

func (s *Store) AtlasFederatedAuthOrgConfig(opts *atlasv2.GetConnectedOrgConfigApiParams) (*atlasv2.ConnectedOrgConfig, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetConnectedOrgConfigWithParams(s.ctx, opts).Execute()
	return result, err
}

func (s *Store) AtlasIdentityProvider(opts *atlasv2.GetIdentityProviderApiParams) (*atlasv2.FederationIdentityProvider, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetIdentityProviderWithParams(s.ctx, opts).Execute()
	return result, err
}
