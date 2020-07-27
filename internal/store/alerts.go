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

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_alerts.go -package=mocks github.com/mongodb/mongocli/internal/store AlertDescriber,AlertLister,AlertAcknowledger

type AlertDescriber interface {
	Alert(string, string) (*atlas.Alert, error)
}

type AlertLister interface {
	Alerts(string, *atlas.AlertsListOptions) (*atlas.AlertsResponse, error)
}

type AlertAcknowledger interface {
	AcknowledgeAlert(string, string, *atlas.AcknowledgeRequest) (*atlas.Alert, error)
}

// Alert encapsulate the logic to manage different cloud providers
func (s *Store) Alert(projectID, alertID string) (*atlas.Alert, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Alerts.Get(context.Background(), projectID, alertID)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Alerts.Get(context.Background(), projectID, alertID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Alerts encapsulate the logic to manage different cloud providers
func (s *Store) Alerts(projectID string, opts *atlas.AlertsListOptions) (*atlas.AlertsResponse, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Alerts.List(context.Background(), projectID, opts)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Alerts.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Acknowledge encapsulate the logic to manage different cloud providers
func (s *Store) AcknowledgeAlert(projectID, alertID string, body *atlas.AcknowledgeRequest) (*atlas.Alert, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Alerts.Acknowledge(context.Background(), projectID, alertID, body)
		return result, err
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Alerts.Acknowledge(context.Background(), projectID, alertID, body)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
