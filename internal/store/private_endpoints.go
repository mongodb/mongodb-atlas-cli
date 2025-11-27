// Copyright 2021 MongoDB Inc
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312010/admin"
)

// PrivateEndpoints encapsulates the logic to manage different cloud providers.
func (s *Store) PrivateEndpoints(projectID, provider string) ([]atlasv2.EndpointService, error) {
	result, _, err := s.clientv2.PrivateEndpointServicesApi.ListPrivateEndpointService(s.ctx, projectID, provider).Execute()
	return result, err
}

// DataLakePrivateEndpoints encapsulates the logic to manage different cloud providers.
func (s *Store) DataLakePrivateEndpoints(params *atlasv2.ListPrivateEndpointIdsApiParams) (*atlasv2.PaginatedPrivateNetworkEndpointIdEntry, error) {
	result, _, err := s.clientv2.DataFederationApi.ListPrivateEndpointIdsWithParams(s.ctx, params).Execute()
	return result, err
}

// PrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) PrivateEndpoint(projectID, provider, privateLinkID string) (*atlasv2.EndpointService, error) {
	result, _, err := s.clientv2.PrivateEndpointServicesApi.GetPrivateEndpointService(s.ctx, projectID, provider, privateLinkID).Execute()
	return result, err
}

// DataLakePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DataLakePrivateEndpoint(projectID, privateLinkID string) (*atlasv2.PrivateNetworkEndpointIdEntry, error) {
	result, _, err := s.clientv2.DataFederationApi.GetPrivateEndpointId(s.ctx, projectID, privateLinkID).Execute()
	return result, err
}

// CreatePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) CreatePrivateEndpoint(projectID string, r *atlasv2.CloudProviderEndpointServiceRequest) (*atlasv2.EndpointService, error) {
	result, _, err := s.clientv2.PrivateEndpointServicesApi.CreatePrivateEndpointService(s.ctx, projectID, r).
		Execute()
	return result, err
}

// DataLakeCreatePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DataLakeCreatePrivateEndpoint(projectID string, r *atlasv2.PrivateNetworkEndpointIdEntry) (*atlasv2.PaginatedPrivateNetworkEndpointIdEntry, error) {
	result, _, err := s.clientv2.DataFederationApi.CreatePrivateEndpointId(s.ctx, projectID, r).
		Execute()
	return result, err
}

// DeletePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DeletePrivateEndpoint(projectID, provider, privateLinkID string) error {
	_, err := s.clientv2.PrivateEndpointServicesApi.DeletePrivateEndpointService(s.ctx, projectID, provider, privateLinkID).Execute()
	return err
}

// DataLakeDeletePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DataLakeDeletePrivateEndpoint(projectID, endpointID string) error {
	_, err := s.clientv2.DataFederationApi.DeletePrivateEndpointId(s.ctx, projectID, endpointID).Execute()
	return err
}

// CreateInterfaceEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) CreateInterfaceEndpoint(projectID, provider, endpointServiceID string, createRequest *atlasv2.CreateEndpointRequest) (*atlasv2.PrivateLinkEndpoint, error) {
	result, _, err := s.clientv2.PrivateEndpointServicesApi.CreatePrivateEndpoint(s.ctx, projectID, provider,
		endpointServiceID, createRequest).Execute()
	return result, err
}

// InterfaceEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) InterfaceEndpoint(projectID, cloudProvider, privateEndpointID, endpointServiceID string) (*atlasv2.PrivateLinkEndpoint, error) {
	result, _, err := s.clientv2.PrivateEndpointServicesApi.GetPrivateEndpoint(s.ctx, projectID, cloudProvider, privateEndpointID, endpointServiceID).Execute()
	return result, err
}

// DeleteInterfaceEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteInterfaceEndpoint(projectID, provider, endpointServiceID, privateEndpointID string) error {
	_, err := s.clientv2.PrivateEndpointServicesApi.DeletePrivateEndpoint(s.ctx, projectID, provider, privateEndpointID, endpointServiceID).Execute()
	return err
}

// UpdateRegionalizedPrivateEndpointSetting encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateRegionalizedPrivateEndpointSetting(projectID string, enabled bool) (*atlasv2.ProjectSettingItem, error) {
	setting := atlasv2.ProjectSettingItem{
		Enabled: enabled,
	}
	result, _, err := s.clientv2.PrivateEndpointServicesApi.
		ToggleRegionalEndpointMode(s.ctx, projectID, &setting).Execute()
	return result, err
}

// RegionalizedPrivateEndpointSetting encapsulates the logic to manage different cloud providers.
func (s *Store) RegionalizedPrivateEndpointSetting(projectID string) (*atlasv2.ProjectSettingItem, error) {
	result, _, err := s.clientv2.PrivateEndpointServicesApi.GetRegionalEndpointMode(s.ctx, projectID).Execute()
	return result, err
}
