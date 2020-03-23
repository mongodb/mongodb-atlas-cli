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

type AlertDescriber interface {
	Alert(string, string) (*atlas.Alert, error)
}

type AlertAcknowledge interface {
	Acknowledge(string, string, *atlas.AcknowledgeRequest) (*atlas.Alert, error)
}

type AlertsStore interface {
	AlertDescriber
	AlertAcknowledge
}
// Alert encapsulate the logic to manage different cloud providers
func (s *Store) Alert(projectID, alertID string) (*atlas.Alert, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Alerts.Get(context.Background(), projectID, alertID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Acknowledge encapsulate the logic to manage different cloud providers
func (s *Store) Acknowledge(projectID, alertID string, body *atlas.AcknowledgeRequest) (*atlas.Alert, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).Alerts.Acknowledge(context.Background(), projectID, alertID, body)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}