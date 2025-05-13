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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

// PrivateEndpointsDeprecated encapsulates the logic to manage different cloud providers.
func (s *Store) PrivateEndpointsDeprecated(projectID string, opts *ListOptions) ([]atlas.PrivateEndpointConnectionDeprecated, error) {
	lst := &atlas.ListOptions{
		PageNum:      opts.PageNum,
		ItemsPerPage: opts.ItemsPerPage,
		IncludeCount: opts.IncludeCount,
	}
	result, _, err := s.client.PrivateEndpointsDeprecated.List(s.ctx, projectID, lst)
	return result, err
}

// PrivateEndpointDeprecated encapsulates the logic to manage different cloud providers.
func (s *Store) PrivateEndpointDeprecated(projectID, privateLinkID string) (*atlas.PrivateEndpointConnectionDeprecated, error) {
	result, _, err := s.client.PrivateEndpointsDeprecated.Get(s.ctx, projectID, privateLinkID)
	return result, err
}

// DeletePrivateEndpointDeprecated encapsulates the logic to manage different cloud providers.
func (s *Store) DeletePrivateEndpointDeprecated(projectID, privateLinkID string) error {
	_, err := s.client.PrivateEndpointsDeprecated.Delete(s.ctx, projectID, privateLinkID)
	return err
}

// CreateInterfaceEndpointDeprecated encapsulates the logic to manage different cloud providers.
func (s *Store) CreateInterfaceEndpointDeprecated(projectID, privateLinkID, interfaceEndpointID string) (*atlas.InterfaceEndpointConnectionDeprecated, error) {
	result, _, err := s.client.PrivateEndpointsDeprecated.AddOneInterfaceEndpoint(s.ctx, projectID, privateLinkID, interfaceEndpointID)
	return result, err
}

// CreatePrivateEndpointDeprecated encapsulates the logic to manage different cloud providers.
func (s *Store) CreatePrivateEndpointDeprecated(projectID string, r *atlas.PrivateEndpointConnectionDeprecated) (*atlas.PrivateEndpointConnectionDeprecated, error) {
	result, _, err := s.client.PrivateEndpointsDeprecated.Create(s.ctx, projectID, r)
	return result, err
}

// InterfaceEndpointDeprecated encapsulates the logic to manage different cloud providers.
func (s *Store) InterfaceEndpointDeprecated(projectID, privateLinkID, interfaceEndpointID string) (*atlas.InterfaceEndpointConnectionDeprecated, error) {
	result, _, err := s.client.PrivateEndpointsDeprecated.GetOneInterfaceEndpoint(s.ctx, projectID, privateLinkID, interfaceEndpointID)
	return result, err
}

// DeleteInterfaceEndpointDeprecated encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteInterfaceEndpointDeprecated(projectID, privateLinkID, interfaceEndpointID string) error {
	_, err := s.client.PrivateEndpointsDeprecated.DeleteOneInterfaceEndpoint(s.ctx, projectID, privateLinkID, interfaceEndpointID)
	return err
}
