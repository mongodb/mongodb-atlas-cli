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

package atlas

import (
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115001/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_cloud_provider_backup.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas ScheduleDescriber

type ScheduleDescriber interface {
	DescribeSchedule(string, string) (*atlasv2.DiskBackupSnapshotSchedule, error)
}

// DescribeSchedule encapsulates the logic to manage different cloud providers.
func (s *Store) DescribeSchedule(projectID, clusterName string) (*atlasv2.DiskBackupSnapshotSchedule, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.CloudBackupsApi.GetBackupSchedule(s.ctx, projectID, clusterName).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
