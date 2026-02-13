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
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312014/admin"
)

// AlertConfigurations encapsulate the logic to manage different cloud providers.
func (s *Store) AlertConfigurations(params *atlasv2.ListAlertConfigsApiParams) (*atlasv2.PaginatedAlertConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListAlertConfigsWithParams(s.ctx, params).Execute()
	return result, err
}

// CreateAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAlertConfiguration(alertConfig *atlasv2.GroupAlertsConfig) (*atlasv2.GroupAlertsConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.CreateAlertConfig(s.ctx, alertConfig.GetGroupId(), alertConfig).Execute()
	return result, err
}

// DeleteAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteAlertConfiguration(projectID, id string) error {
	_, err := s.clientv2.AlertConfigurationsApi.DeleteAlertConfig(s.ctx, projectID, id).Execute()
	return err
}

// MatcherFields encapsulate the logic to manage different cloud providers.
func (s *Store) MatcherFields() ([]string, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListMatcherFieldNames(s.ctx).Execute()
	return result, err
}

func (s *Store) UpdateAlertConfiguration(alertConfig *atlasv2.GroupAlertsConfig) (*atlasv2.GroupAlertsConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.UpdateAlertConfig(s.ctx, alertConfig.GetGroupId(), alertConfig.GetId(), alertConfig).Execute()
	return result, err
}

// EnableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) EnableAlertConfiguration(projectID, id string) (*atlasv2.GroupAlertsConfig, error) {
	toggle := atlasv2.AlertsToggle{
		Enabled: pointer.Get(true),
	}
	result, _, err := s.clientv2.AlertConfigurationsApi.ToggleAlertConfig(s.ctx, projectID, id, &toggle).Execute()
	return result, err
}

// DisableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DisableAlertConfiguration(projectID, id string) (*atlasv2.GroupAlertsConfig, error) {
	toggle := atlasv2.AlertsToggle{
		Enabled: pointer.Get(false),
	}
	result, _, err := s.clientv2.AlertConfigurationsApi.ToggleAlertConfig(s.ctx, projectID, id, &toggle).Execute()
	return result, err
}

func (s *Store) AlertConfiguration(projectID, id string) (*atlasv2.GroupAlertsConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.GetAlertConfig(s.ctx, projectID, id).Execute()
	return result, err
}
