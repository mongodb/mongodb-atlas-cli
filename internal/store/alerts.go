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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
)

// Alert encapsulates the logic to manage different cloud providers.
func (s *Store) Alert(params *atlasv2.GetAlertApiParams) (*atlasv2.AlertViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertsApi.GetAlertWithParams(s.ctx, params).Execute()
	return result, err
}

// Alerts encapsulate the logic to manage different cloud providers.
func (s *Store) Alerts(params *atlasv2.ListAlertsApiParams) (*atlasv2.PaginatedAlert, error) {
	result, _, err := s.clientv2.AlertsApi.ListAlertsWithParams(s.ctx, params).Execute()
	return result, err
}

// AcknowledgeAlert encapsulate the logic to manage different cloud providers.
func (s *Store) AcknowledgeAlert(params *atlasv2.AcknowledgeAlertApiParams) (*atlasv2.AlertViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertsApi.AcknowledgeAlertWithParams(s.ctx, params).Execute()
	return result, err
}
