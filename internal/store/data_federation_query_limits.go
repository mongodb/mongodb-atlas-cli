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

// This code was autogenerated at 2023-06-23T15:50:53+01:00. Note: Manual updates are allowed, but may be overwritten.

package store

import (
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

// DataFederationQueryLimits encapsulates the logic to manage different cloud providers.
func (s *Store) DataFederationQueryLimits(projectID, tenantName string) ([]atlasv2.DataFederationTenantQueryLimit, error) {
	result, _, err := s.clientv2.DataFederationApi.ReturnFederatedDatabaseQueryLimits(s.ctx, projectID, tenantName).Execute()
	return result, err
}

// DataFederationQueryLimit encapsulates the logic to manage different cloud providers.
func (s *Store) DataFederationQueryLimit(projectID, tenantName, limitName string) (*atlasv2.DataFederationTenantQueryLimit, error) {
	result, _, err := s.clientv2.DataFederationApi.ReturnFederatedDatabaseQueryLimit(s.ctx, projectID, tenantName, limitName).Execute()
	return result, err
}

// CreateDataFederationQueryLimit encapsulates the logic to manage different cloud providers.
func (s *Store) CreateDataFederationQueryLimit(projectID, tenantName, limitName string, opts *atlasv2.DataFederationTenantQueryLimit) (*atlasv2.DataFederationTenantQueryLimit, error) {
	result, _, err := s.clientv2.DataFederationApi.CreateOneDataFederationQueryLimit(s.ctx, projectID, tenantName, limitName, opts).Execute()
	return result, err
}

// DeleteDataFederationQueryLimit encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteDataFederationQueryLimit(projectID, tenantName, limitName string) error {
	_, err := s.clientv2.DataFederationApi.DeleteOneDataFederationInstanceQueryLimit(s.ctx, projectID, tenantName, limitName).Execute()
	return err
}
