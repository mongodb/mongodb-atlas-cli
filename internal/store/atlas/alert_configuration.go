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

package atlas

import (
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_alert_configuration.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas AlertConfigurationLister,AlertConfigurationCreator,AlertConfigurationDeleter,AlertConfigurationUpdater,MatcherFieldsLister,AlertConfigurationEnabler,AlertConfigurationDisabler

type AlertConfigurationLister interface {
	AlertConfigurations(string, *atlas.ListOptions) ([]atlas.AlertConfiguration, error)
}

type AlertConfigurationCreator interface {
	CreateAlertConfiguration(*atlas.AlertConfiguration) (*atlas.AlertConfiguration, error)
}

type AlertConfigurationDeleter interface {
	DeleteAlertConfiguration(string, string) error
}

type AlertConfigurationUpdater interface {
	UpdateAlertConfiguration(*atlas.AlertConfiguration) (*atlas.AlertConfiguration, error)
}

type MatcherFieldsLister interface {
	MatcherFields() ([]string, error)
}

type AlertConfigurationEnabler interface {
	EnableAlertConfiguration(string, string) (*atlas.AlertConfiguration, error)
}

type AlertConfigurationDisabler interface {
	DisableAlertConfiguration(string, string) (*atlas.AlertConfiguration, error)
}

// AlertConfigurations encapsulate the logic to manage different cloud providers.
func (s *Store) AlertConfigurations(projectID string, opts *atlas.ListOptions) ([]atlas.AlertConfiguration, error) {
	result, _, err := s.client.(*atlas.Client).AlertConfigurations.List(s.ctx, projectID, opts)
	return result, err
}

// CreateAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAlertConfiguration(alertConfig *atlas.AlertConfiguration) (*atlas.AlertConfiguration, error) {
	result, _, err := s.client.(*atlas.Client).AlertConfigurations.Create(s.ctx, alertConfig.GroupID, alertConfig)
	return result, err
}

// DeleteAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteAlertConfiguration(projectID, id string) error {
	_, err := s.client.(*atlas.Client).AlertConfigurations.Delete(s.ctx, projectID, id)
	return err
}

// MatcherFields encapsulate the logic to manage different cloud providers.
func (s *Store) MatcherFields() ([]string, error) {
	result, _, err := s.client.(*atlas.Client).AlertConfigurations.ListMatcherFields(s.ctx)
	return result, err
}

func (s *Store) UpdateAlertConfiguration(alertConfig *atlas.AlertConfiguration) (*atlas.AlertConfiguration, error) {
	result, _, err := s.client.(*atlas.Client).AlertConfigurations.Update(s.ctx, alertConfig.GroupID, alertConfig.ID, alertConfig)
	return result, err

}

// EnableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) EnableAlertConfiguration(projectID, id string) (*atlas.AlertConfiguration, error) {
	result, _, err := s.client.(*atlas.Client).AlertConfigurations.EnableAnAlertConfig(s.ctx, projectID, id, pointer.Get(true))
	return result, err
}

// DisableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DisableAlertConfiguration(projectID, id string) (*atlas.AlertConfiguration, error) {
	result, _, err := s.client.(*atlas.Client).AlertConfigurations.EnableAnAlertConfig(s.ctx, projectID, id, pointer.Get(false))
	return result, err
}
