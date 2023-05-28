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

package atlas

import (
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_alerts.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas AlertDescriber,AlertLister,AlertAcknowledger

type AlertDescriber interface {
	Alert(string, string) (*atlasv2.AlertViewForNdsGroup, error)
}

type AlertLister interface {
	Alerts(projectID string, status string) (*atlasv2.PaginatedAlert, error)
}

type AlertAcknowledger interface {
	AcknowledgeAlert(string, string, *atlasv2.AlertViewForNdsGroup) (*atlasv2.AlertViewForNdsGroup, error)
}

// Alert encapsulate the logic to manage different cloud providers.
func (s *Store) Alert(projectID, alertID string) (*atlasv2.AlertViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertsApi.GetAlert(s.ctx, projectID, alertID).Execute()
	return result, err
}

// Alerts encapsulate the logic to manage different cloud providers.
func (s *Store) Alerts(projectID string, status string) (*atlasv2.PaginatedAlert, error) {
	request := s.clientv2.AlertsApi.ListAlerts(s.ctx, projectID)
	if status != "" {
		request = request.Status(status)
	}
	result, _, err := request.Execute()
	return result, err
}

// Acknowledge encapsulate the logic to manage different cloud providers.
func (s *Store) AcknowledgeAlert(projectID, alertID string, body *atlasv2.AlertViewForNdsGroup) (*atlasv2.AlertViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertsApi.AcknowledgeAlert(s.ctx, projectID, alertID, body).Execute()
	return result, err
}
