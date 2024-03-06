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
	atlas "go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

//go:generate mockgen -destination=../mocks/mock_service_version.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store ServiceVersionDescriber

type ServiceVersionDescriber interface {
	ServiceVersion() (*opsmngr.ServiceVersion, error)
}

// ServiceVersion encapsulates the logic to manage different cloud providers.
func (s *Store) ServiceVersion() (*atlas.ServiceVersion, error) {
	result, _, err := s.client.ServiceVersion.Get(s.ctx)
	return result, err
}
