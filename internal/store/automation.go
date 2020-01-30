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

	"github.com/10gen/mcli/internal/config"
	"github.com/mongodb-labs/pcgc/cloudmanager"
)

type AutomationGetter interface {
	GetAutomationConfig(string) (*cloudmanager.AutomationConfig, error)
}

type AutomationUpdater interface {
	UpdateAutomationConfig(string, *cloudmanager.AutomationConfig) error
}

type AutomationStore interface {
	AutomationGetter
	AutomationUpdater
}

// GetAutomationConfig encapsulate the logic to manage different cloud providers
func (s *Store) GetAutomationConfig(projectID string) (*cloudmanager.AutomationConfig, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*cloudmanager.Client).AutomationConfig.Get(context.Background(), projectID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// UpdateAutomationConfig encapsulate the logic to manage different cloud providers
func (s *Store) UpdateAutomationConfig(projectID string, automationConfig *cloudmanager.AutomationConfig) error {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		_, err := s.client.(*cloudmanager.Client).AutomationConfig.Update(context.Background(), projectID, automationConfig)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
