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
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

//go:generate mockgen -destination=../mocks/mock_ldap_configurations.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store LDAPConfigurationVerifier,LDAPConfigurationDescriber,LDAPConfigurationSaver,LDAPConfigurationDeleter,LDAPConfigurationGetter

type LDAPConfigurationVerifier interface {
	VerifyLDAPConfiguration(string, *atlasv2.LDAPVerifyConnectivityJobRequestParams) (*atlasv2.LDAPVerifyConnectivityJobRequest, error)
}

type LDAPConfigurationDescriber interface {
	GetStatusLDAPConfiguration(string, string) (*atlasv2.LDAPVerifyConnectivityJobRequest, error)
}

type LDAPConfigurationDeleter interface {
	DeleteLDAPConfiguration(string) error
}

type LDAPConfigurationSaver interface {
	SaveLDAPConfiguration(string, *atlasv2.UserSecurity) (*atlasv2.UserSecurity, error)
}

type LDAPConfigurationGetter interface {
	GetLDAPConfiguration(string) (*atlasv2.UserSecurity, error)
}

// VerifyLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) VerifyLDAPConfiguration(projectID string, ldap *atlasv2.LDAPVerifyConnectivityJobRequestParams) (*atlasv2.LDAPVerifyConnectivityJobRequest, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		resp, _, err := s.clientv2.LDAPConfigurationApi.VerifyLDAPConfiguration(s.ctx, projectID, ldap).
			Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// GetStatusLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) GetStatusLDAPConfiguration(projectID, requestID string) (*atlasv2.LDAPVerifyConnectivityJobRequest, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		resp, _, err := s.clientv2.LDAPConfigurationApi.GetLDAPConfigurationStatus(s.ctx, projectID, requestID).Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// SaveLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) SaveLDAPConfiguration(projectID string, ldap *atlasv2.UserSecurity) (*atlasv2.UserSecurity, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		resp, _, err := s.clientv2.LDAPConfigurationApi.SaveLDAPConfiguration(s.ctx, projectID, ldap).
			Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// DeleteLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) DeleteLDAPConfiguration(projectID string) error {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		_, _, err := s.clientv2.LDAPConfigurationApi.DeleteLDAPConfiguration(s.ctx, projectID).Execute()
		return err
	default:
		return fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

// GetLDAPConfiguration encapsulates the logic to manage different cloud providers.
func (s *Store) GetLDAPConfiguration(projectID string) (*atlasv2.UserSecurity, error) {
	switch s.service {
	case config.CloudService, config.CloudGovService:
		resp, _, err := s.clientv2.LDAPConfigurationApi.GetLDAPConfiguration(s.ctx, projectID).Execute()
		return resp, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}
