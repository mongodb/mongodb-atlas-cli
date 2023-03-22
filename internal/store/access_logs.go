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
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas/api/v1alpha"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_access_logs.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AccessLogsListerByClusterName,AccessLogsListerByHostname,AccessLogsLister

type AccessLogsListerByClusterName interface {
	AccessLogsByClusterName(string, string, *atlas.AccessLogOptions) (*atlas.AccessLogSettings, error)
}

type AccessLogsListerByHostname interface {
	AccessLogsByHostname(string, string, *atlas.AccessLogOptions) (*atlas.AccessLogSettings, error)
}

type AccessLogsLister interface {
	AccessLogsByHostname(string, string, *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error)
	AccessLogsByClusterName(string, string, *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error)
}

// AccessLogsByHostname encapsulates the logic to manage different cloud providers.
func (s *Store) AccessLogsByHostname(groupID, hostname string, opts *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		startTime, _ := time.Parse(time.RFC3339, opts.Start)
		endTime, _ := time.Parse(time.RFC3339, opts.End)
		result, _, err := s.clientv2.AccessTrackingApi.ListAccessLogsByHostname(s.ctx, groupID, hostname).Start(startTime).End(endTime).NLogs(int32(opts.NLogs)).IpAddress(opts.IPAddress).AuthResult(*opts.AuthResult).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AccessLogsByClusterName encapsulates the logic to manage different cloud providers.
func (s *Store) AccessLogsByClusterName(groupID, clusterName string, opts *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		startTime, _ := time.Parse(time.RFC3339, opts.Start)
		result, _, err := s.clientv2.AccessTrackingApi.ListAccessLogsByClusterName(s.ctx, groupID, clusterName).Start(startTime).End(opts.End).NLogs(int64(opts.NLogs)).IpAddress(opts.IPAddress).AuthResult(*opts.AuthResult).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
