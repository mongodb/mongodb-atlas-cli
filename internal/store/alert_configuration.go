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
	"github.com/mongodb/mcli/internal/config"
)

type AlertConfigurationLister interface {
	AlertConfigurations(string, *atlas.ListOptions) ([]atlas.AlertConfiguration, error)
}

type AlertConfigurationCreator interface {
	CreateAlertConfiguration(*atlas.AlertConfiguration) (*atlas.AlertConfiguration, error)
}

type AlertConfigurationDeleter interface {
	DeleteAlertConfiguration(string, string) error
}

type AlertConfigurationStore interface {
	AlertConfigurationLister
	AlertConfigurationCreator
	AlertConfigurationDeleter
}

// AlertConfigurations encapsulate the logic to manage different cloud providers
func (s *Store) AlertConfigurations(projectID string, opts *atlas.ListOptions) ([]atlas.AlertConfiguration, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).AlertConfigurations.List(context.Background(), projectID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateAlertConfiguration encapsulate the logic to manage different cloud providers
func (s *Store) CreateAlertConfiguration(alertConfig *atlas.AlertConfiguration) (*atlas.AlertConfiguration, error) {
	switch s.service {
	case config.CloudService:
		// TODO: fix me https://github.com/mongodb/go-client-mongodb-atlas/pull/56
		project := alertConfig.GroupID
		alertConfig.GroupID = ""
		result, _, err := s.client.(*atlas.Client).AlertConfigurations.Create(context.Background(), project, alertConfig)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DeleteAlertConfiguration encapsulate the logic to manage different cloud providers
func (s *Store) DeleteAlertConfiguration(projectID, id string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).AlertConfigurations.Delete(context.Background(), projectID, id)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
