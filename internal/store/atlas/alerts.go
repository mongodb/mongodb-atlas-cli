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

package atlasv2

import (
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlasv2/mock_alerts.go -package=atlasv2 github.com/mongodb/mongodb-atlasv2-cli/internal/store/atlasv2 AlertDescriber,AlertLister,AlertAcknowledger

type AlertDescriber interface {
	Alert(string, string) (*atlasv2.AlertViewForNdsGroup, error)
}

type AlertLister interface {
	Alerts(projectID string, status string, opts atlas.ListOptions) (*atlasv2.PaginatedAlert, error)
}

type AlertAcknowledger interface {
	// Issue: This should return altas.Alert instead
	AcknowledgeAlert(string, string, *atlasv2.AlertViewForNdsGroup) (*atlasv2.AlertViewForNdsGroup, error)
}

// Alert encapsulate the logic to manage different cloud providers.
func (s *Store) Alert(projectID, alertID string) (*atlasv2.AlertViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertsApi.GetAlert(s.ctx, projectID, alertID).Execute()
	return result, err
}

// Alerts encapsulate the logic to manage different cloud providers.
func (s *Store) Alerts(projectID string, status string, opts ListOptions) (*atlasv2.PaginatedAlert, error) {
	result, _, err := s.clientv2.AlertsApi.ListAlerts(s.ctx, projectID).
		PageNum(int32(opts.PageNum)).
		Status(status).
		ItemsPerPage(int32(opts.ItemsPerPage)).Execute()
	return result, err
}

// Acknowledge encapsulate the logic to manage different cloud providers.
func (s *Store) AcknowledgeAlert(projectID, alertID string, body *atlasv2.AlertViewForNdsGroup) (*atlasv2.AlertViewForNdsGroup, error) {
	// Issue: AlertViewForNdsGroup contains whole object where originally we only had AcknowledgedUntil field
	result, _, err := s.clientv2.AlertsApi.AcknowledgeAlert(s.ctx, projectID, alertID).AlertViewForNdsGroup(*body).Execute()
	return result, err
}
