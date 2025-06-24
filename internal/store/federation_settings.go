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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312004/admin"
)

// FederationSetting encapsulate the logic to manage different cloud providers.
func (s *Store) FederationSetting(opts *atlasv2.GetFederationSettingsApiParams) (*atlasv2.OrgFederationSettings, error) {
	result, _, err := s.clientv2.FederatedAuthenticationApi.GetFederationSettingsWithParams(s.ctx, opts).Execute()
	return result, err
}
