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
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_version_manifest.go -package=mocks github.com/mongodb/mongocli/internal/store VersionManifestUpdater

type VersionManifestUpdater interface {
	UpdateVersionManifest(string) (*opsmngr.VersionManifest, error)
}

// UpdateVersionManifest encapsulates the logic to manage different cloud providers
func (s *Store) UpdateVersionManifest(version string) (*opsmngr.VersionManifest, error) {
	switch s.service {
	case config.OpsManagerService:
		versionManifestStruct, _, err := s.client.(*opsmngr.Client).VersionManifest.Get(context.Background(), version)
		if err != nil {
			return nil, err
		}
		result, _, err := s.client.(*opsmngr.Client).VersionManifest.Update(context.Background(), versionManifestStruct)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
