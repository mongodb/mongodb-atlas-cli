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
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

//go:generate mockgen -destination=../mocks/mock_integrations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store IntegrationCreator,IntegrationLister,IntegrationDeleter,IntegrationDescriber

type IntegrationCreator interface {
	CreateIntegration(string, string, *atlasv2.ThirdPartyIntegration) (*atlasv2.PaginatedIntegration, error)
}

type IntegrationLister interface {
	Integrations(string) (*atlasv2.PaginatedIntegration, error)
}

type IntegrationDeleter interface {
	DeleteIntegration(string, string) error
}

type IntegrationDescriber interface {
	Integration(string, string) (*atlasv2.ThirdPartyIntegration, error)
}

// CreateIntegration encapsulates the logic to manage different cloud providers.
func (s *Store) CreateIntegration(projectID, integrationType string, integration *atlasv2.ThirdPartyIntegration) (*atlasv2.PaginatedIntegration, error) {
	resp, _, err := s.clientv2.ThirdPartyIntegrationsApi.CreateThirdPartyIntegration(s.ctx,
		integrationType, projectID, integration).Execute()
	return resp, err
}

// Integrations encapsulates the logic to manage different cloud providers.
func (s *Store) Integrations(projectID string) (*atlasv2.PaginatedIntegration, error) {
	resp, _, err := s.clientv2.ThirdPartyIntegrationsApi.ListThirdPartyIntegrations(s.ctx, projectID).Execute()
	return resp, err
}

// DeleteIntegration encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteIntegration(projectID, integrationType string) error {
	_, _, err := s.clientv2.ThirdPartyIntegrationsApi.DeleteThirdPartyIntegration(s.ctx, integrationType, projectID).Execute()
	return err
}

// Integration encapsulates the logic to manage different cloud providers.
func (s *Store) Integration(projectID, integrationType string) (*atlasv2.ThirdPartyIntegration, error) {
	resp, _, err := s.clientv2.ThirdPartyIntegrationsApi.GetThirdPartyIntegration(s.ctx, projectID, integrationType).Execute()
	return resp, err
}
