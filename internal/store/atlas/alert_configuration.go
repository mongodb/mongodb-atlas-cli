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
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlasv2/mock_alert_configuration.go -package=atlasv2 github.com/mongodb/mongodb-atlasv2-cli/internal/store/atlasv2 AlertConfigurationLister,AlertConfigurationCreator,AlertConfigurationDeleter,AlertConfigurationUpdater,MatcherFieldsLister,AlertConfigurationEnabler,AlertConfigurationDisabler

type AlertConfigurationLister interface {
	AlertConfigurations(string, *atlas.ListOptions) (*atlasv2.PaginatedAlertConfig, error)
}

type AlertConfigurationCreator interface {
	CreateAlertConfiguration(*atlasv2.AlertConfigViewForNdsGroup) (*atlasv2.AlertConfigViewForNdsGroup, error)
}

type AlertConfigurationDeleter interface {
	DeleteAlertConfiguration(string, string) error
}

type AlertConfigurationUpdater interface {
	UpdateAlertConfiguration(*atlasv2.AlertConfigViewForNdsGroup) (*atlasv2.AlertConfigViewForNdsGroup, error)
}

type MatcherFieldsLister interface {
	MatcherFields() ([]string, error)
}

type AlertConfigurationEnabler interface {
	EnableAlertConfiguration(string, string) (*atlasv2.AlertConfigViewForNdsGroup, error)
}

type AlertConfigurationDisabler interface {
	DisableAlertConfiguration(string, string) (*atlasv2.AlertConfigViewForNdsGroup, error)
}

// AlertConfigurations encapsulate the logic to manage different cloud providers.
func (s *Store) AlertConfigurations(projectID string, opts *atlas.ListOptions) (*atlasv2.PaginatedAlertConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListAlertConfigurations(s.ctx, projectID).
		IncludeCount(opts.IncludeCount).
		PageNum(int32(opts.PageNum)).
		ItemsPerPage(int32(opts.ItemsPerPage)).Execute()
	return result, err
}

// CreateAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAlertConfiguration(alertConfig *atlasv2.AlertConfigViewForNdsGroup) (*atlasv2.AlertConfigViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.
		CreateAlertConfiguration(s.ctx, alertConfig.GetGroupId()).AlertConfigViewForNdsGroup(*alertConfig).Execute()
	return result, err
}

// DeleteAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DeleteAlertConfiguration(projectID, id string) error {
	_, err := s.client.AlertConfigurations.Delete(s.ctx, projectID, id)
	return err
}

// MatcherFields encapsulate the logic to manage different cloud providers.
func (s *Store) MatcherFields() ([]string, error) {
	result, _, err := s.client.AlertConfigurations.ListMatcherFields(s.ctx)
	return result, err
}

func (s *Store) UpdateAlertConfiguration(alertConfig *atlasv2.AlertConfigViewForNdsGroup) (*atlasv2.AlertConfigViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.UpdateAlertConfiguration(s.ctx, alertConfig.GetGroupId(), alertConfig.GetId()).
		AlertConfigViewForNdsGroup(*alertConfig).Execute()
	return result, err
}

// EnableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) EnableAlertConfiguration(projectID, id string) (*atlasv2.AlertConfigViewForNdsGroup, error) {
	enabled := true
	result, _, err := s.clientv2.AlertConfigurationsApi.
		ToggleAlertConfiguration(s.ctx, projectID, id).Toggle(atlasv2.Toggle{Enabled: &enabled}).Execute()
	return result, err
}

// DisableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DisableAlertConfiguration(projectID, id string) (*atlasv2.AlertConfigViewForNdsGroup, error) {
	enabled := false
	result, _, err := s.clientv2.AlertConfigurationsApi.
		ToggleAlertConfiguration(s.ctx, projectID, id).Toggle(atlasv2.Toggle{Enabled: &enabled}).Execute()
	return result, err
}
