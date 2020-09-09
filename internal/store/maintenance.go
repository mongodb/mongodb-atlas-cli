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

	"go.mongodb.org/ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_maintenance.go -package=mocks github.com/mongodb/mongocli/internal/store MaintenanceWindowUpdater,MaintenanceWindowClearer,MaintenanceWindowDeferrer,MaintenanceWindowDescriber,MaintenanceWindowCreator

type MaintenanceWindowUpdater interface {
	UpdateMaintenanceWindow(string, *atlas.MaintenanceWindow) error
}

type MaintenanceWindowClearer interface {
	ClearMaintenanceWindow(string) error
}

type MaintenanceWindowDeferrer interface {
	DeferMaintenanceWindow(string) error
}

type MaintenanceWindowDescriber interface {
	MaintenanceWindow(string) (*atlas.MaintenanceWindow, error)
}

type MaintenanceWindowCreator interface {
	CreateMaintenanceWindow(string, *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error)
}

// UpdateMaintenanceWindow encapsulates the logic to manage different cloud providers
func (s *Store) UpdateMaintenanceWindow(projectID string, maintenanceWindow *atlas.MaintenanceWindow) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).MaintenanceWindows.Update(context.Background(), projectID, maintenanceWindow)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ClearMaintenanceWindow encapsulates the logic to manage different cloud providers
func (s *Store) ClearMaintenanceWindow(projectID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).MaintenanceWindows.Reset(context.Background(), projectID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeferMaintenanceWindow encapsulates the logic to manage different cloud providers
func (s *Store) DeferMaintenanceWindow(projectID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).MaintenanceWindows.Defer(context.Background(), projectID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// MaintenanceWindow encapsulates the logic to manage different cloud providers
func (s *Store) MaintenanceWindow(projectID string) (*atlas.MaintenanceWindow, error) {
	switch s.service {
	case config.CloudService:
		resp, _, err := s.client.(*atlas.Client).MaintenanceWindows.Get(context.Background(), projectID)
		return resp, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateMaintenanceWindow encapsulates the logic to manage different cloud providers
func (s *Store) CreateMaintenanceWindow(projectID string, maintenanceWindow *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		log, _, err := s.client.(*opsmngr.Client).MaintenanceWindows.Create(context.Background(), projectID, maintenanceWindow)
		return log, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}