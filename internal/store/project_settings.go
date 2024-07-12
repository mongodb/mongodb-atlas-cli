// Copyright 2022 MongoDB Inc
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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_project_settings.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store ProjectSettingsDescriber,ProjectSettingsUpdater

type ProjectSettingsDescriber interface {
	ProjectSettings(string) (*atlasv2.GroupSettings, error)
}

type ProjectSettingsUpdater interface {
	UpdateProjectSettings(string, *atlasv2.GroupSettings) (*atlasv2.GroupSettings, error)
}

// ProjectSettings encapsulates the logic of getting settings of a particular project.
func (s *Store) ProjectSettings(projectID string) (*atlasv2.GroupSettings, error) {
	result, _, err := s.clientv2.ProjectsApi.GetProjectSettings(s.ctx, projectID).Execute()
	return result, err
}

// UpdateProjectSettings encapsulates the logic of updating settings of a particular project.
func (s *Store) UpdateProjectSettings(projectID string, projectSettings *atlasv2.GroupSettings) (*atlasv2.GroupSettings, error) {
	result, _, err := s.clientv2.ProjectsApi.UpdateProjectSettings(s.ctx, projectID, projectSettings).Execute()
	return result, err
}
