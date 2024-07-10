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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_alert_configuration.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store AlertConfigurationLister,AlertConfigurationCreator,AlertConfigurationDeleter,AlertConfigurationUpdater,MatcherFieldsLister,AlertConfigurationEnabler,AlertConfigurationDisabler,AlertConfigurationDescriber

type AlertConfigurationLister interface {
	AlertConfigurations(*admin.ListAlertConfigurationsApiParams) (*admin.PaginatedAlertConfig, error)
}

type AlertConfigurationCreator interface {
	CreateAlertConfiguration(*admin.GroupAlertsConfig) (*admin.GroupAlertsConfig, error)
}

type AlertConfigurationDeleter interface {
	DeleteAlertConfiguration(string, string) error
}

type AlertConfigurationUpdater interface {
	UpdateAlertConfiguration(*admin.GroupAlertsConfig) (*admin.GroupAlertsConfig, error)
}

type MatcherFieldsLister interface {
	MatcherFields() ([]string, error)
}

type AlertConfigurationEnabler interface {
	EnableAlertConfiguration(string, string) (*admin.GroupAlertsConfig, error)
}

type AlertConfigurationDisabler interface {
	DisableAlertConfiguration(string, string) (*admin.GroupAlertsConfig, error)
}

type AlertConfigurationDescriber interface {
	AlertConfiguration(string, string) (*admin.GroupAlertsConfig, error)
}

// AlertConfigurations encapsulate the logic to manage different cloud providers.
func (s *Store) AlertConfigurations(params *admin.ListAlertConfigurationsApiParams) (*admin.PaginatedAlertConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListAlertConfigurationsWithParams(s.ctx, params).Execute()
	return result, err
}

// CreateAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAlertConfiguration(alertConfig *admin.GroupAlertsConfig) (*admin.GroupAlertsConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.CreateAlertConfiguration(s.ctx, alertConfig.GetGroupId(), alertConfig).Execute()
	return result, err
}

// DeleteAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteAlertConfiguration(projectID, id string) error {
	_, err := s.clientv2.AlertConfigurationsApi.DeleteAlertConfiguration(s.ctx, projectID, id).Execute()
	return err
}

// MatcherFields encapsulate the logic to manage different cloud providers.
func (s *Store) MatcherFields() ([]string, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListAlertConfigurationMatchersFieldNames(s.ctx).Execute()
	return result, err
}

func (s *Store) UpdateAlertConfiguration(alertConfig *admin.GroupAlertsConfig) (*admin.GroupAlertsConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.UpdateAlertConfiguration(s.ctx, alertConfig.GetGroupId(), alertConfig.GetId(), alertConfig).Execute()
	return result, err
}

// EnableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) EnableAlertConfiguration(projectID, id string) (*admin.GroupAlertsConfig, error) {
	toggle := admin.AlertsToggle{
		Enabled: pointer.Get(true),
	}
	result, _, err := s.clientv2.AlertConfigurationsApi.ToggleAlertConfiguration(s.ctx, projectID, id, &toggle).Execute()
	return result, err
}

// DisableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DisableAlertConfiguration(projectID, id string) (*admin.GroupAlertsConfig, error) {
	toggle := admin.AlertsToggle{
		Enabled: pointer.Get(false),
	}
	result, _, err := s.clientv2.AlertConfigurationsApi.ToggleAlertConfiguration(s.ctx, projectID, id, &toggle).Execute()
	return result, err
}

func (s *Store) AlertConfiguration(projectID, id string) (*admin.GroupAlertsConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.GetAlertConfiguration(s.ctx, projectID, id).Execute()
	return result, err
}
