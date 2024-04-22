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

// This code was autogenerated at 2023-06-22T17:46:21+01:00. Note: Manual updates are allowed, but may be overwritten.

package store

import (
	"go.mongodb.org/atlas-sdk/v20231115010/admin"
)

//go:generate mockgen -destination=../mocks/mock_data_federation_private_endpoint.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store DataFederationPrivateEndpointLister,DataFederationPrivateEndpointDescriber,DataFederationPrivateEndpointCreator,DataFederationPrivateEndpointDeleter

type DataFederationPrivateEndpointLister interface {
	DataFederationPrivateEndpoints(string) (*admin.PaginatedPrivateNetworkEndpointIdEntry, error)
}

type DataFederationPrivateEndpointCreator interface {
	CreateDataFederationPrivateEndpoint(string, *admin.PrivateNetworkEndpointIdEntry) (*admin.PaginatedPrivateNetworkEndpointIdEntry, error)
}

type DataFederationPrivateEndpointDeleter interface {
	DeleteDataFederationPrivateEndpoint(string, string) error
}

type DataFederationPrivateEndpointDescriber interface {
	DataFederationPrivateEndpoint(string, string) (*admin.PrivateNetworkEndpointIdEntry, error)
}

// DataFederationPrivateEndpoints encapsulates the logic to manage different cloud providers.
func (s *Store) DataFederationPrivateEndpoints(projectID string) (*admin.PaginatedPrivateNetworkEndpointIdEntry, error) {
	result, _, err := s.clientv2.DataFederationApi.ListDataFederationPrivateEndpoints(s.ctx, projectID).Execute()
	return result, err
}

// DataFederationPrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DataFederationPrivateEndpoint(projectID, id string) (*admin.PrivateNetworkEndpointIdEntry, error) {
	result, _, err := s.clientv2.DataFederationApi.GetDataFederationPrivateEndpoint(s.ctx, projectID, id).Execute()
	return result, err
}

// CreateDataFederationPrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) CreateDataFederationPrivateEndpoint(projectID string, opts *admin.PrivateNetworkEndpointIdEntry) (*admin.PaginatedPrivateNetworkEndpointIdEntry, error) {
	result, _, err := s.clientv2.DataFederationApi.CreateDataFederationPrivateEndpoint(s.ctx, projectID, opts).Execute()
	return result, err
}

// DeleteDataFederationPrivateEndpoint encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteDataFederationPrivateEndpoint(projectID, id string) error {
	_, _, err := s.clientv2.DataFederationApi.DeleteDataFederationPrivateEndpoint(s.ctx, projectID, id).Execute()
	return err
}
