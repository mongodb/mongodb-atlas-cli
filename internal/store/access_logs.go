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
	"strconv"

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312011/admin"
)

type AccessLogOptions struct {
	Start      string
	End        string
	NLogs      int
	IPAddress  string
	AuthResult *bool
}

// AccessLogsByHostname encapsulates the logic to manage different cloud providers.
func (s *Store) AccessLogsByHostname(groupID, hostname string, opts *AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error) {
	result := s.clientv2.AccessTrackingApi.GetAccessHistoryProcess(s.ctx, groupID, hostname)

	if opts != nil {
		if opts.Start != "" {
			startTime, _ := strconv.ParseInt(opts.Start, 10, 64)
			result = result.Start(startTime)
		}
		if opts.End != "" {
			endTime, _ := strconv.ParseInt(opts.End, 10, 64)
			result = result.End(endTime)
		}

		if opts.NLogs > 0 {
			result = result.NLogs(opts.NLogs)
		}

		if opts.IPAddress != "" {
			result = result.IpAddress(opts.IPAddress)
		}

		if opts.AuthResult != nil {
			result = result.AuthResult(*opts.AuthResult)
		}
	}

	res, _, err := result.Execute()
	return res, err
}

// AccessLogsByClusterName encapsulates the logic to manage different cloud providers.
func (s *Store) AccessLogsByClusterName(groupID, clusterName string, opts *AccessLogOptions) (*atlasv2.MongoDBAccessLogsList, error) {
	result := s.clientv2.AccessTrackingApi.GetAccessHistoryCluster(s.ctx, groupID, clusterName)

	if opts != nil {
		if opts.Start != "" {
			startTime, _ := strconv.ParseInt(opts.Start, 10, 64)
			result = result.Start(startTime)
		}
		if opts.End != "" {
			endTime, _ := strconv.ParseInt(opts.End, 10, 64)
			result = result.End(endTime)
		}

		if opts.NLogs > 0 {
			result = result.NLogs(opts.NLogs)
		}

		if opts.IPAddress != "" {
			result = result.IpAddress(opts.IPAddress)
		}

		if opts.AuthResult != nil {
			result = result.AuthResult(*opts.AuthResult)
		}
	}
	res, _, err := result.Execute()

	return res, err
}
