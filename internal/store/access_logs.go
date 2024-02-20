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
	"strconv"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115007/admin"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_access_logs.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AccessLogsListerByClusterName,AccessLogsListerByHostname,AccessLogsLister

type AccessLogsListerByClusterName interface {
	AccessLogsByClusterName(string, string, *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error)
}

type AccessLogsListerByHostname interface {
	AccessLogsByHostname(string, string, *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error)
}

type AccessLogsLister interface {
	AccessLogsByHostname(string, string, *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error)
	AccessLogsByClusterName(string, string, *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error)
}

// AccessLogsByHostname encapsulates the logic to manage different cloud providers.
func (s *Store) AccessLogsByHostname(groupID, hostname string, opts *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		params := &atlasv2.ListAccessLogsByHostnameApiParams{
			GroupId:  groupID,
			Hostname: hostname,
		}
		if opts != nil {
			if opts.Start != "" {
				startTime, _ := strconv.ParseInt(opts.Start, 10, 64)
				params.Start = &startTime
			}
			if opts.End != "" {
				endTime, _ := strconv.ParseInt(opts.End, 10, 64)
				params.End = &endTime
			}

			if opts.NLogs > 0 {
				params.NLogs = &opts.NLogs
			}

			if opts.IPAddress != "" {
				params.IpAddress = &opts.IPAddress
			}

			if opts.AuthResult != nil {
				params.AuthResult = opts.AuthResult
			}
		}
		res, _, err := s.clientv2.AccessTrackingApi.ListAccessLogsByHostnameWithParams(s.ctx, params).Execute()
		return res, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// AccessLogsByClusterName encapsulates the logic to manage different cloud providers.
func (s *Store) AccessLogsByClusterName(groupID, clusterName string, opts *atlas.AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		params := &atlasv2.ListAccessLogsByClusterNameApiParams{
			GroupId:     groupID,
			ClusterName: clusterName,
		}
		if opts != nil {
			if opts.Start != "" {
				startTime, _ := strconv.ParseInt(opts.Start, 10, 64)
				params.Start = &startTime
			}
			if opts.End != "" {
				endTime, _ := strconv.ParseInt(opts.End, 10, 64)
				params.End = &endTime
			}

			if opts.NLogs > 0 {
				params.NLogs = &opts.NLogs
			}

			if opts.IPAddress != "" {
				params.IpAddress = &opts.IPAddress
			}

			if opts.AuthResult != nil {
				params.AuthResult = opts.AuthResult
			}
		}
		res, _, err := s.clientv2.AccessTrackingApi.ListAccessLogsByClusterNameWithParams(s.ctx, params).Execute()

		return res, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
