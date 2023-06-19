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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

//go:generate mockgen -destination=../mocks/mock_private_endpoints.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store PrivateEndpointLister,PrivateEndpointDescriber,PrivateEndpointCreator,PrivateEndpointDeleter,InterfaceEndpointDescriber,InterfaceEndpointCreator,InterfaceEndpointDeleter,RegionalizedPrivateEndpointSettingUpdater,RegionalizedPrivateEndpointSettingDescriber,DataLakePrivateEndpointLister,DataLakePrivateEndpointCreator,DataLakePrivateEndpointDeleter,DataLakePrivateEndpointDescriber

type PrivateEndpointLister interface {
	PrivateEndpoints(string, string) ([]atlasv2.EndpointService, error)
}

type DataLakePrivateEndpointLister interface {
	DataLakePrivateEndpoints(*atlasv2.ListDataFederationPrivateEndpointsApiParams) (*atlasv2.PaginatedPrivateNetworkEndpointIdEntry, error)
}

type PrivateEndpointDescriber interface {
	PrivateEndpoint(string, string, string) (interface{}, error)
}

type DataLakePrivateEndpointDescriber interface {
	DataLakePrivateEndpoint(string, string) (*atlasv2.PrivateNetworkEndpointIdEntry, error)
}

type PrivateEndpointCreator interface {
	CreatePrivateEndpoint(string, *atlasv2.CloudProviderEndpointServiceRequest) (interface{}, error)
}

type DataLakePrivateEndpointCreator interface {
	DataLakeCreatePrivateEndpoint(string, *atlasv2.PrivateNetworkEndpointIdEntry) (*atlasv2.PaginatedPrivateNetworkEndpointIdEntry, error)
}

type PrivateEndpointDeleter interface {
	DeletePrivateEndpoint(string, string, string) error
}

type DataLakePrivateEndpointDeleter interface {
	DataLakeDeletePrivateEndpoint(string, string) error
}

type InterfaceEndpointDescriber interface {
	InterfaceEndpoint(string, string, string, string) (interface{}, error)
}

type InterfaceEndpointCreator interface {
	CreateInterfaceEndpoint(string, string, string, *atlasv2.CreateEndpointRequest) (interface{}, error)
}

type InterfaceEndpointDeleter interface {
	DeleteInterfaceEndpoint(string, string, string, string) error
}

type RegionalizedPrivateEndpointSettingUpdater interface {
	UpdateRegionalizedPrivateEndpointSetting(string, bool) (*atlasv2.ProjectSettingItem, error)
}

type RegionalizedPrivateEndpointSettingDescriber interface {
	RegionalizedPrivateEndpointSetting(string) (*atlasv2.ProjectSettingItem, error)
}

// PrivateEndpoints encapsulates the logic to manage different cloud providers.
func (s *Store) PrivateEndpoints(projectID, provider string) ([]atlasv2.EndpointService, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.PrivateEndpointServicesApi.ListPrivateEndpointServices(s.ctx, projectID, provider).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DataLakePrivateEndpoints encapsulates the logic to manage different cloud providers.
func (s *Store) DataLakePrivateEndpoints(params *atlasv2.ListDataFederationPrivateEndpointsApiParams) (*atlasv2.PaginatedPrivateNetworkEndpointIdEntry, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.DataFederationApi.ListDataFederationPrivateEndpointsWithParams(s.ctx, params).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// PrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) PrivateEndpoint(projectID, provider, privateLinkID string) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.PrivateEndpointServicesApi.GetPrivateEndpointService(s.ctx, projectID, provider, privateLinkID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DataLakePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DataLakePrivateEndpoint(projectID, privateLinkID string) (*atlasv2.PrivateNetworkEndpointIdEntry, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.DataFederationApi.GetDataFederationPrivateEndpoint(s.ctx, projectID, privateLinkID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreatePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) CreatePrivateEndpoint(projectID string, r *atlasv2.CloudProviderEndpointServiceRequest) (interface{}, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.PrivateEndpointServicesApi.CreatePrivateEndpointService(s.ctx, projectID, r).
			Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DataLakeCreatePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DataLakeCreatePrivateEndpoint(projectID string, r *atlasv2.PrivateNetworkEndpointIdEntry) (*atlasv2.PaginatedPrivateNetworkEndpointIdEntry, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.DataFederationApi.CreateDataFederationPrivateEndpoint(s.ctx, projectID, r).
			Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeletePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DeletePrivateEndpoint(projectID, provider, privateLinkID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.PrivateEndpointServicesApi.DeletePrivateEndpointService(s.ctx, projectID, provider, privateLinkID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DataLakeDeletePrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DataLakeDeletePrivateEndpoint(projectID, endpointID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.DataFederationApi.DeleteDataFederationPrivateEndpoint(s.ctx, projectID, endpointID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateInterfaceEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) CreateInterfaceEndpoint(projectID, provider, endpointServiceID string, createRequest *atlasv2.CreateEndpointRequest) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.PrivateEndpointServicesApi.CreatePrivateEndpoint(s.ctx, projectID, provider,
			endpointServiceID, createRequest).Execute()
		return result.GetActualInstance(), err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// InterfaceEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) InterfaceEndpoint(projectID, cloudProvider, endpointServiceID, privateEndpointID string) (interface{}, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.PrivateEndpointServicesApi.GetPrivateEndpoint(s.ctx, projectID, cloudProvider, endpointServiceID, privateEndpointID).Execute()
		return result.GetActualInstance(), err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteInterfaceEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteInterfaceEndpoint(projectID, provider, endpointServiceID, privateEndpointID string) error {
	switch s.service {
	case config.CloudService:
		_, _, err := s.clientv2.PrivateEndpointServicesApi.DeletePrivateEndpoint(s.ctx, projectID, provider, privateEndpointID, endpointServiceID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// UpdateRegionalizedPrivateEndpointSetting encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateRegionalizedPrivateEndpointSetting(projectID string, enabled bool) (*atlasv2.ProjectSettingItem, error) {
	switch s.service {
	case config.CloudService:
		setting := atlasv2.ProjectSettingItem{
			Enabled: enabled,
		}
		result, _, err := s.clientv2.PrivateEndpointServicesApi.
			ToggleRegionalizedPrivateEndpointSetting(s.ctx, projectID, &setting).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// RegionalizedPrivateEndpointSetting encapsulates the logic to manage different cloud providers.
func (s *Store) RegionalizedPrivateEndpointSetting(projectID string) (*atlasv2.ProjectSettingItem, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.clientv2.PrivateEndpointServicesApi.GetRegionalizedPrivateEndpointSetting(s.ctx, projectID).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
