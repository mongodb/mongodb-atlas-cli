// Copyright 2022 MongoDB Inc
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240805003/admin"
)

//go:generate mockgen -destination=../mocks/mock_serverless_private_endpoints.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ServerlessPrivateEndpointsLister

type ServerlessPrivateEndpointsLister interface {
	ServerlessPrivateEndpoints(string, string) ([]atlasv2.ServerlessTenantEndpoint, error)
}

func (s *Store) ServerlessPrivateEndpoints(projectID, instanceName string) ([]atlasv2.ServerlessTenantEndpoint, error) {
	result, _, err := s.clientv2.ServerlessPrivateEndpointsApi.ListServerlessPrivateEndpoints(s.ctx, projectID, instanceName).Execute()
	return result, err
}
