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
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../mocks/mock_access_logs.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AccessLogsLister

type AccessLogsLister interface {
	AccessLogsByHostname(string, string) (atlasv2.AccessTrackingApiListAccessLogsByHostnameRequest, error)
	AccessLogsByClusterName(string, string) (atlasv2.AccessTrackingApiListAccessLogsByClusterNameRequest, error)
}

// AccessLogsByHostname encapsulates the logic to manage different cloud providers.
func (s *Store) AccessLogsByHostname(groupID, hostname string) (atlasv2.AccessTrackingApiListAccessLogsByHostnameRequest, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		return s.clientv2.AccessTrackingApi.ListAccessLogsByHostname(s.ctx, groupID, hostname), nil
	default:
		return atlasv2.AccessTrackingApiListAccessLogsByHostnameRequest{}, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AccessLogsByClusterName encapsulates the logic to manage different cloud providers.
func (s *Store) AccessLogsByClusterName(groupID, clusterName string) (atlasv2.AccessTrackingApiListAccessLogsByClusterNameRequest, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		return s.clientv2.AccessTrackingApi.ListAccessLogsByClusterName(s.ctx, groupID, clusterName), nil

	default:
		return atlasv2.AccessTrackingApiListAccessLogsByClusterNameRequest{}, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
