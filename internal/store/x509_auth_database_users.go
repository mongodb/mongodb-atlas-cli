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

//go:generate mockgen -destination=../mocks/mock_x509_certificate_store.go -package=mocks github.com/mongodb/mongocli/internal/store X509CertificateLister,X509CertificateStore

type X509CertificateLister interface {
	// TODO :: can we make this use the Result object
	X509Certificates(string, string) ([]atlas.UserCertificate, error)
}

type X509CertificateStore interface {
	X509CertificateLister
}

// X509Certificates retrieves all unexpired certificates for a database user
func (s *Store) X509Certificates(projectID, username string) ([]atlas.UserCertificate, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).X509AuthDBUsers.GetUserCertificates(context.Background(), projectID, username)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
