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
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115002/admin"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_maintenance.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store MaintenanceWindowUpdater,MaintenanceWindowClearer,MaintenanceWindowDeferrer,MaintenanceWindowDescriber,OpsManagerMaintenanceWindowCreator,OpsManagerMaintenanceWindowLister,OpsManagerMaintenanceWindowDeleter,OpsManagerMaintenanceWindowDescriber,OpsManagerMaintenanceWindowUpdater

type MaintenanceWindowUpdater interface {
	UpdateMaintenanceWindow(string, *atlasv2.GroupMaintenanceWindow) error
}

type MaintenanceWindowClearer interface {
	ClearMaintenanceWindow(string) error
}

type MaintenanceWindowDeferrer interface {
	DeferMaintenanceWindow(string) error
}

type MaintenanceWindowDescriber interface {
	MaintenanceWindow(string) (*atlasv2.GroupMaintenanceWindow, error)
}

type OpsManagerMaintenanceWindowCreator interface {
	CreateOpsManagerMaintenanceWindow(string, *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error)
}

type OpsManagerMaintenanceWindowLister interface {
	OpsManagerMaintenanceWindows(string) (*opsmngr.MaintenanceWindows, error)
}

type OpsManagerMaintenanceWindowDeleter interface {
	DeleteOpsManagerMaintenanceWindow(string, string) error
}

type OpsManagerMaintenanceWindowDescriber interface {
	OpsManagerMaintenanceWindow(string, string) (*opsmngr.MaintenanceWindow, error)
}

type OpsManagerMaintenanceWindowUpdater interface {
	UpdateOpsManagerMaintenanceWindow(string, *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error)
}

// UpdateMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateMaintenanceWindow(projectID string, maintenanceWindow *atlasv2.GroupMaintenanceWindow) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.MaintenanceWindowsApi.UpdateMaintenanceWindow(s.ctx, projectID, maintenanceWindow).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// ClearMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) ClearMaintenanceWindow(projectID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.clientv2.MaintenanceWindowsApi.ResetMaintenanceWindow(s.ctx, projectID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeferMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) DeferMaintenanceWindow(projectID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.clientv2.MaintenanceWindowsApi.DeferMaintenanceWindow(s.ctx, projectID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// MaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) MaintenanceWindow(projectID string) (*atlasv2.GroupMaintenanceWindow, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		resp, _, err := s.clientv2.MaintenanceWindowsApi.GetMaintenanceWindow(s.ctx, projectID).Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// CreateOpsManagerMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) CreateOpsManagerMaintenanceWindow(projectID string, maintenanceWindow *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		log, _, err := s.client.(*opsmngr.Client).MaintenanceWindows.Create(s.ctx, projectID, maintenanceWindow)
		return log, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OpsManagerMaintenanceWindows encapsulates the logic to manage different cloud providers.
func (s *Store) OpsManagerMaintenanceWindows(projectID string) (*opsmngr.MaintenanceWindows, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		log, _, err := s.client.(*opsmngr.Client).MaintenanceWindows.List(s.ctx, projectID)
		return log, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteOpsManagerMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteOpsManagerMaintenanceWindow(projectID, maintenanceWindowID string) error {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		_, err := s.client.(*opsmngr.Client).MaintenanceWindows.Delete(s.ctx, projectID, maintenanceWindowID)
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// OpsManagerMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) OpsManagerMaintenanceWindow(projectID, maintenanceWindowID string) (*opsmngr.MaintenanceWindow, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		log, _, err := s.client.(*opsmngr.Client).MaintenanceWindows.Get(s.ctx, projectID, maintenanceWindowID)
		return log, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// UpdateOpsManagerMaintenanceWindow encapsulates the logic to manage different cloud providers.
func (s *Store) UpdateOpsManagerMaintenanceWindow(projectID string, maintenanceWindow *opsmngr.MaintenanceWindow) (*opsmngr.MaintenanceWindow, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		log, _, err := s.client.(*opsmngr.Client).MaintenanceWindows.Update(s.ctx, projectID, maintenanceWindow.ID, maintenanceWindow)
		return log, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
