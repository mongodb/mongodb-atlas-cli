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
	atlasOld "go.mongodb.org/atlas/mongodbatlas"
	atlas "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../../mocks/atlas/mock_alert_configuration.go -package=atlas github.com/mongodb/mongodb-atlas-cli/internal/store/atlas AlertConfigurationLister,AlertConfigurationCreator,AlertConfigurationDeleter,AlertConfigurationUpdater,MatcherFieldsLister,AlertConfigurationEnabler,AlertConfigurationDisabler

type AlertConfigurationLister interface {
	AlertConfigurations(string, *atlasOld.ListOptions) ([]atlasOld.AlertConfiguration, error)
}

type AlertConfigurationCreator interface {
	CreateAlertConfiguration(*atlas.AlertConfigViewForNdsGroup) (*atlas.AlertConfigViewForNdsGroup, error)
}

type AlertConfigurationDeleter interface {
	DeleteAlertConfiguration(string, string) error
}

type AlertConfigurationUpdater interface {
	UpdateAlertConfiguration(*atlas.AlertConfigViewForNdsGroup) (*atlas.AlertConfigViewForNdsGroup, error)
}

type MatcherFieldsLister interface {
	MatcherFields() ([]string, error)
}

type AlertConfigurationEnabler interface {
	EnableAlertConfiguration(string, string) (*atlas.AlertConfigViewForNdsGroup, error)
}

type AlertConfigurationDisabler interface {
	DisableAlertConfiguration(string, string) (*atlas.AlertConfigViewForNdsGroup, error)
}

// AlertConfigurations encapsulate the logic to manage different cloud providers.
func (s *Store) AlertConfigurations(projectID string, opts ListOptions) (*atlas.PaginatedAlertConfig, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.ListAlertConfigurations(s.ctx, projectID).
		IncludeCount(opts.IncludeCount).
		PageNum(int32(opts.PageNum)).
		ItemsPerPage(int32(opts.ItemsPerPage)).Execute()
	return result, err
}

// CreateAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) CreateAlertConfiguration(alertConfig *atlas.AlertConfigViewForNdsGroup) (*atlas.AlertConfigViewForNdsGroup, error) {
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

func (s *Store) UpdateAlertConfiguration(alertConfig atlas.AlertConfigViewForNdsGroup) (*atlas.AlertConfigViewForNdsGroup, error) {
	result, _, err := s.clientv2.AlertConfigurationsApi.UpdateAlertConfiguration(s.ctx, alertConfig.GetGroupId(), alertConfig.GetId()).AlertConfigViewForNdsGroup(alertConfig).Execute()
	return result, err
}

// EnableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) EnableAlertConfiguration(projectID, id string) (*atlas.AlertConfigViewForNdsGroup, error) {
	enabled := true
	result, _, err := s.clientv2.AlertConfigurationsApi.
		ToggleAlertConfiguration(s.ctx, projectID, id).Toggle(atlas.Toggle{Enabled: &enabled}).Execute()
	return result, err
}

// DisableAlertConfiguration encapsulate the logic to manage different cloud providers.
func (s *Store) DisableAlertConfiguration(projectID, id string) (*atlas.AlertConfigViewForNdsGroup, error) {
	enabled := false
	result, _, err := s.clientv2.AlertConfigurationsApi.
		ToggleAlertConfiguration(s.ctx, projectID, id).Toggle(atlas.Toggle{Enabled: &enabled}).Execute()
	return result, err
}
