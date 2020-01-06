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
	UpdateAutomationConfig(string, *cloudmanager.AutomationConfig) (*cloudmanager.AutomationConfig, error)
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
func (s *Store) UpdateAutomationConfig(projectID string, automationConfig *cloudmanager.AutomationConfig) (*cloudmanager.AutomationConfig, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*cloudmanager.Client).AutomationConfig.Update(context.Background(), projectID, automationConfig)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
