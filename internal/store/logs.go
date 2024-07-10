// Copyright 2020 MongoDB Inc
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
	"errors"
	"io"

	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_logs.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store LogsDownloader

type LogsDownloader interface {
	DownloadLog(*atlasv2.GetHostLogsApiParams) (io.ReadCloser, error)
}

// DownloadLog encapsulates the logic to manage different cloud providers.
func (s *Store) DownloadLog(params *atlasv2.GetHostLogsApiParams) (io.ReadCloser, error) {
	result, _, err := s.clientv2.MonitoringAndLogsApi.GetHostLogsWithParams(s.ctx, params).Execute()
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, errors.New("returned file is empty")
	}
	return result, nil
}
