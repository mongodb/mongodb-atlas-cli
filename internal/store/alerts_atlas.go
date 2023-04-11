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
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_alerts.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store AlertDescriber,AlertLister,AlertAcknowledger

type AtlasAlertDescriber interface {
	AlertAtlas(string, string) (*atlas.Alert, error)
}

type AtlasAlertLister interface {
	AlertsAtlas(string, *atlas.AlertsListOptions) (*atlas.AlertsResponse, error)
}

type AtlasAlertAcknowledger interface {
	AcknowledgeAlertAtlas(string, string, *atlas.AcknowledgeRequest) (*atlas.Alert, error)
}

// Alert encapsulate the logic to manage different cloud providers.
func (s *Store) AlertAtlas(projectID, alertID string) (*atlas.Alert, error) {
	result, _, err := s.client.(*atlas.Client).Alerts.Get(s.ctx, projectID, alertID)
	return result, err
}

// Alerts encapsulate the logic to manage different cloud providers.
func (s *Store) AlertsAtlas(projectID string, opts *atlas.AlertsListOptions) (*atlas.AlertsResponse, error) {
	result, _, err := s.client.(*atlas.Client).Alerts.List(s.ctx, projectID, opts)
	return result, err
}

// Acknowledge encapsulate the logic to manage different cloud providers.
func (s *Store) AcknowledgeAlertAtlas(projectID, alertID string, body *atlas.AcknowledgeRequest) (*atlas.Alert, error) {
	result, _, err := s.client.(*atlas.Client).Alerts.Acknowledge(s.ctx, projectID, alertID, body)
	return result, err
}
