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

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongocli/internal/config"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_service_version.go -package=mocks github.com/mongodb/mongocli/internal/store ServiceVersionGetter

type ServiceVersionGetter interface {
	GetServiceVersion() (*opsmngr.ServiceVersion, error)
}

// GetCurrentVersionManifest encapsulates the logic to manage different cloud providers.
func (s *Store) GetServiceVersion() (*atlas.ServiceVersion, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).ServiceVersion.Get(context.Background())
		return result, err
	case config.CloudService, config.CloudGovService:
		result, _, err := s.client.(*atlas.Client).ServiceVersion.Get(context.Background())
		return result, err
	default:
		return nil, fmt.Errorf("%w: %s", errUnsupportedService, s.service)
	}
}

func omAtLeastFive(version *atlas.ServiceVersion) (bool, error) {
	sv, err := semver.NewVersion(version.Version)
	if err != nil {
		return false, err
	}

	constrain, _ := semver.NewConstraint(">= 5.0")
	return constrain.Check(sv), nil
}

func checkOMIsAtLeastFive(s *Store) (bool, error) {
	version, err := s.GetServiceVersion()
	if err != nil {
		return false, err
	}
	omIsAtLeastFive, err := omAtLeastFive(version)
	if err != nil {
		return false, err
	}
	return omIsAtLeastFive, nil
}
