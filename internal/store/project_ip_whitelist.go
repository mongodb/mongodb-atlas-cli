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

	"github.com/10gen/mcli/internal/config"
	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

type ProjectIPWhitelistCreator interface {
	CreateProjectIPWhitelist(*atlas.ProjectIPWhitelist) ([]atlas.ProjectIPWhitelist, error)
}

type ProjectIPWhitelistStore interface {
	ProjectIPWhitelistCreator
}

// CreateProjectIPWhitelist encapsulate the logic to manage different cloud providers
func (s *Store) CreateProjectIPWhitelist(whitelist *atlas.ProjectIPWhitelist) ([]atlas.ProjectIPWhitelist, error) {
	switch s.service {
	case config.CloudService:
		result, _, err := s.client.(*atlas.Client).ProjectIPWhitelist.Create(context.Background(), whitelist.GroupID, []*atlas.ProjectIPWhitelist{whitelist})
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
