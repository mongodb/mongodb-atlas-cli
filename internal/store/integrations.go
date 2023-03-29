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
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	atlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

//go:generate mockgen -destination=../mocks/mock_integrations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store IntegrationCreator,IntegrationLister,IntegrationDeleter,IntegrationDescriber

type IntegrationCreator interface {
	CreateIntegration(string, string, *atlas.ThirdPartyIntegration) (*atlasv2.GroupPaginatedIntegration, error)
}

type IntegrationLister interface {
	Integrations(string) (*atlasv2.GroupPaginatedIntegration, error)
}

type IntegrationDeleter interface {
	DeleteIntegration(string, string) error
}

type IntegrationDescriber interface {
	Integration(string, string) (*atlasv2.IntegrationViewForNdsGroup, error)
}

// CreateIntegration encapsulates the logic to manage different cloud providers.
func (s *Store) CreateIntegration(projectID, integrationType string, integration *atlas.ThirdPartyIntegration) (*atlasv2.GroupPaginatedIntegration, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		thirdPartyIntegration := getThirdPartyIntegration(integration)
		resp, _, err := s.clientv2.ThirdPartyIntegrationsApi.CreateThirdPartyIntegration(s.ctx, integrationType, projectID).IntegrationViewForNdsGroup(thirdPartyIntegration).Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

func getThirdPartyIntegration(integration *atlas.ThirdPartyIntegration) atlasv2.IntegrationViewForNdsGroup {
	var result atlasv2.IntegrationViewForNdsGroup

	switch integration.Type {
	case "PAGER_DUTY":
		result = atlasv2.IntegrationViewForNdsGroup{
			PagerDuty: &atlasv2.PagerDuty{
				Region:     &integration.Region,
				ServiceKey: integration.ServiceKey,
				Type:       &integration.Type,
			},
		}

	case "SLACK":
		result = atlasv2.IntegrationViewForNdsGroup{
			Slack: &atlasv2.Slack{
				ApiToken: integration.APIToken,
				// ChannelName: integration.ChannelName,
				TeamName: &integration.TeamName,
				Type:     &integration.Type,
			},
		}

	case "DATADOG":
		result = atlasv2.IntegrationViewForNdsGroup{
			Datadog: &atlasv2.Datadog{
				ApiKey: integration.APIKey,
				Region: &integration.Region,
				Type:   &integration.Type,
			},
		}
	case "OPS_GENIE":
		result = atlasv2.IntegrationViewForNdsGroup{
			OpsGenie: &atlasv2.OpsGenie{
				ApiKey: integration.APIKey,
				Region: &integration.Region,
				Type:   &integration.Type,
			},
		}

	case "WEBHOOK":
		result = atlasv2.IntegrationViewForNdsGroup{
			Webhook: &atlasv2.Webhook{
				Secret: &integration.Secret,
				Type:   &integration.Type,
				Url:    integration.URL,
			},
		}

	case "MICROSOFT_TEAMS":
		result = atlasv2.IntegrationViewForNdsGroup{
			MicrosoftTeams: &atlasv2.MicrosoftTeams{
				MicrosoftTeamsWebhookUrl: integration.MicrosoftTeamsWebhookURL,
				Type:                     &integration.Type,
			},
		}

	case "PROMETHEUS":
		result = atlasv2.IntegrationViewForNdsGroup{
			Prometheus: &atlasv2.Prometheus{
				Enabled: integration.Enabled,
				//ListenAddress: ,
				Password: &integration.Password,
				//RateLimitInterval: ,
				Scheme:           integration.Scheme,
				ServiceDiscovery: integration.ServiceDiscovery,
				//TlsPemPath: ,
				Username: integration.UserName,
				Type:     &integration.Type,
			},
		}

	case "VICTOR_OPS":
		result = atlasv2.IntegrationViewForNdsGroup{
			VictorOps: &atlasv2.VictorOps{
				ApiKey:     integration.APIKey,
				RoutingKey: &integration.RoutingKey,
				Type:       &integration.Type,
			},
		}

	case "NEW_RELIC":
		result = atlasv2.IntegrationViewForNdsGroup{
			NewRelic: &atlasv2.NewRelic{
				AccountId:  integration.AccountID,
				LicenseKey: integration.LicenseKey,
				ReadToken:  integration.ReadToken,
				WriteToken: integration.WriteToken,
				Type:       &integration.Type,
			},
		}
	}
	return result
}

// Integrations encapsulates the logic to manage different cloud providers.
func (s *Store) Integrations(projectID string) (*atlasv2.GroupPaginatedIntegration, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		resp, _, err := s.clientv2.ThirdPartyIntegrationsApi.ListThirdPartyIntegrations(s.ctx, projectID).Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteIntegration encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteIntegration(projectID, integrationType string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, err := s.clientv2.ThirdPartyIntegrationsApi.DeleteThirdPartyIntegration(s.ctx, integrationType, projectID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// Integration encapsulates the logic to manage different cloud providers.
func (s *Store) Integration(projectID, integrationType string) (*atlasv2.IntegrationViewForNdsGroup, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		resp, _, err := s.clientv2.ThirdPartyIntegrationsApi.GetThirdPartyIntegration(s.ctx, projectID, integrationType).Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
