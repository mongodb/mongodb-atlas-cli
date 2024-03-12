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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_measurements.go -package=mocks github.com/mongodb/mongodb-atlas-cli/mongocli/v2/internal/store HostMeasurementLister,HostDiskMeasurementsLister,HostDatabaseMeasurementsLister

type HostMeasurementLister interface {
	HostMeasurements(string, string, *opsmngr.ProcessMeasurementListOptions) (*opsmngr.ProcessMeasurements, error)
}

type HostDiskMeasurementsLister interface {
	HostDiskMeasurements(string, string, string, *opsmngr.ProcessMeasurementListOptions) (*opsmngr.ProcessDiskMeasurements, error)
}

type HostDatabaseMeasurementsLister interface {
	HostDatabaseMeasurements(string, string, string, *opsmngr.ProcessMeasurementListOptions) (*opsmngr.ProcessDatabaseMeasurements, error)
}

// HostMeasurements encapsulate the logic to manage different cloud providers.
func (s *Store) HostMeasurements(groupID, host string, opts *opsmngr.ProcessMeasurementListOptions) (*opsmngr.ProcessMeasurements, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Measurements.Host(s.ctx, groupID, host, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// HostDiskMeasurements encapsulates the logic to manage different cloud providers.
func (s *Store) HostDiskMeasurements(groupID, hostID, partitionName string, opts *opsmngr.ProcessMeasurementListOptions) (*opsmngr.ProcessDiskMeasurements, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Measurements.Disk(s.ctx, groupID, hostID, partitionName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// HostDatabaseMeasurements encapsulate the logic to manage different cloud providers.
func (s *Store) HostDatabaseMeasurements(groupID, hostID, databaseName string, opts *opsmngr.ProcessMeasurementListOptions) (*opsmngr.ProcessDatabaseMeasurements, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.Measurements.Database(s.ctx, groupID, hostID, databaseName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
