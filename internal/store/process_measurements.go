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

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

//go:generate mockgen -destination=../mocks/mock_process_measurements.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ProcessMeasurementLister

type ProcessMeasurementLister interface {
	ProcessMeasurements(*atlasv2.GetHostMeasurementsApiParams) (*atlasv2.ApiMeasurementsGeneralViewAtlas, error)
}

// ProcessMeasurements encapsulate the logic to manage different cloud providers.
func (s *Store) ProcessMeasurements(params *atlasv2.GetHostMeasurementsApiParams) (*atlasv2.ApiMeasurementsGeneralViewAtlas, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		result, _, err := s.clientv2.MonitoringAndLogsApi.GetHostMeasurementsWithParams(s.ctx, params).Execute()
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
