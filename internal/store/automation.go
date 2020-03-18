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

	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
)

type AutomationGetter interface {
	GetAutomationConfig(string) (*om.AutomationConfig, error)
}

type AutomationUpdater interface {
	UpdateAutomationConfig(string, *om.AutomationConfig) error
}

type AutomationStatusGetter interface {
	GetAutomationStatus(string) (*om.AutomationStatus, error)
}

type AllClusterLister interface {
	ListAllClustersProjects() (*om.AllClustersProjects, error)
}

type AutomationStore interface {
	AutomationGetter
	AutomationUpdater
	AutomationStatusGetter
}

type CloudManagerClustersLister interface {
	AutomationGetter
	AllClusterLister
}

func (s *Store) GetAutomationStatus(projectID string) (*om.AutomationStatus, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*om.Client).AutomationStatus.Get(context.Background(), projectID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// GetAutomationConfig encapsulate the logic to manage different cloud providers
func (s *Store) GetAutomationConfig(projectID string) (*om.AutomationConfig, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*om.Client).AutomationConfig.Get(context.Background(), projectID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// UpdateAutomationConfig encapsulate the logic to manage different cloud providers
func (s *Store) UpdateAutomationConfig(projectID string, automationConfig *om.AutomationConfig) error {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*om.Client).AutomationConfig.Update(context.Background(), projectID, automationConfig)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}

// ListAllClustersProjects encapsulate the logic to manage different cloud providers
func (s *Store) ListAllClustersProjects() (*om.AllClustersProjects, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*om.Client).AllClusters.List(context.Background())
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
