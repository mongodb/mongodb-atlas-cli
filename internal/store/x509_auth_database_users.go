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

	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

//go:generate mockgen -destination=../mocks/mock_x509_certificate_store.go -package=mocks github.com/mongodb/mongocli/internal/store X509CertificateDescriber,X509CertificateSaver,X509CertificateStore

type X509CertificateDescriber interface {
	X509Configuration(string) (*atlas.CustomerX509, error)
}

type X509CertificateSaver interface {
	SaveX509Configuration(string, string) (*atlas.CustomerX509, error)
}

type X509CertificateDisabler interface {
	DisableX509Configuration(string) error
}

type UserCertificateDescriber interface {
	GetUserCertificates(string, string) ([]atlas.UserCertificate, error)
}

type X509CertificateStore interface {
	UserCertificateDescriber
	X509CertificateDescriber
	X509CertificateSaver
	X509CertificateDisabler
}

// X509Configuration retrieves the current user managed certificates for a database user
func (s *Store) X509Configuration(projectID string) (*atlas.CustomerX509, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).X509AuthDBUsers.GetCurrentX509Conf(context.Background(), projectID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// SaveX509Configuration saves a customer-managed X.509 configuration for an Atlas project.
func (s *Store) SaveX509Configuration(projectID, certificate string) (*atlas.CustomerX509, error) {
	switch s.service {
	case config.CloudService:
		userCertificate := &atlas.CustomerX509{Cas: certificate}
		result, _, err := s.client.(*atlas.Client).X509AuthDBUsers.SaveConfiguration(context.Background(), projectID, userCertificate)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// GetUserCertificates retrieves the current user managed certificates for a database user
func (s *Store) GetUserCertificates(projectID, username string) ([]atlas.UserCertificate, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).X509AuthDBUsers.GetUserCertificates(context.Background(), projectID, username)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// DisableX509Configuration disables customer-managed X.509 configuration for an Atlas project.
func (s *Store) DisableX509Configuration(projectID string) error {
	switch s.service {
	case config.CloudService:
		_, err := s.client.(*atlas.Client).X509AuthDBUsers.DisableCustomerX509(context.Background(), projectID)
		return err
	default:
		return fmt.Errorf("unsupported service: %s", s.service)
	}
}
