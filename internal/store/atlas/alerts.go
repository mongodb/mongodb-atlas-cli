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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_alerts.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas AlertDescriber,AlertLister,AlertAcknowledger

type AlertDescriber interface {
	Alert(string, string) (*atlas.Alert, error)
}

type AlertLister interface {
	Alerts(string, *atlas.AlertsListOptions) (*atlas.AlertsResponse, error)
}

type AlertAcknowledger interface {
	AcknowledgeAlert(string, string, *atlas.AcknowledgeRequest) (*atlas.Alert, error)
}

// Alert encapsulate the logic to manage different cloud providers.
func (s *Store) Alert(projectID, alertID string) (*atlas.Alert, error) {
	result, _, err := s.client.(*atlas.Client).Alerts.Get(s.ctx, projectID, alertID)
	return result, err
}

// Alerts encapsulate the logic to manage different cloud providers.
func (s *Store) Alerts(projectID string, opts *atlas.AlertsListOptions) (*atlas.AlertsResponse, error) {
	result, _, err := s.client.(*atlas.Client).Alerts.List(s.ctx, projectID, opts)
	return result, err
}

// Acknowledge encapsulate the logic to manage different cloud providers.
func (s *Store) AcknowledgeAlert(projectID, alertID string, body *atlas.AcknowledgeRequest) (*atlas.Alert, error) {
	result, _, err := s.client.(*atlas.Client).Alerts.Acknowledge(s.ctx, projectID, alertID, body)
	return result, err
}
