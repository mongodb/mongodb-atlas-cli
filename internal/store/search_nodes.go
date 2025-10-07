// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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

// SearchNodes encapsulate the logic to manage different cloud providers.
func (s *Store) SearchNodes(projectID, clusterName string) (*atlasv2.ApiSearchDeploymentResponse, error) {
	result, _, err := s.clientv2.AtlasSearchApi.GetClusterSearchDeployment(s.ctx, projectID, clusterName).Execute()
	return result, err
}
func (s *Store) CreateSearchNodes(projectID, clusterName string, spec *atlasv2.ApiSearchDeploymentRequest) (*atlasv2.ApiSearchDeploymentResponse, error) {
	result, _, err := s.clientv2.AtlasSearchApi.CreateClusterSearchDeployment(s.ctx, projectID, clusterName, spec).Execute()
	return result, err
}

func (s *Store) UpdateSearchNodes(projectID, clusterName string, spec *atlasv2.ApiSearchDeploymentRequest) (*atlasv2.ApiSearchDeploymentResponse, error) {
	result, _, err := s.clientv2.AtlasSearchApi.UpdateClusterSearchDeployment(s.ctx, projectID, clusterName, spec).Execute()
	return result, err
}

func (s *Store) DeleteSearchNodes(projectID, clusterName string) error {
	_, err := s.clientv2.AtlasSearchApi.DeleteClusterSearchDeployment(s.ctx, projectID, clusterName).Execute()
	return err
}
