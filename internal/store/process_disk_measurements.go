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
	"context"
	"fmt"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/config"
)

//go:generate mockgen  -destination=../mocks/mock_process_disk_measurements.go -package=mocks github.com/mongodb/mongocli/internal/store ProcessDiskMeasurementsLister

type ProcessDiskMeasurementsLister interface {
	ProcessDiskMeasurements(string, string, int, string, *atlas.ProcessMeasurementListOptions) (*atlas.ProcessDiskMeasurements, error)
}

// ProcessDiskMeasurements encapsulate the logic to manage different cloud providers
func (s *Store) ProcessDiskMeasurements(groupID, host string, port int, partitionName string, opts *atlas.ProcessMeasurementListOptions) (*atlas.ProcessDiskMeasurements, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProcessDiskMeasurements.List(context.Background(), groupID, host, port, partitionName, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
