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
	"net/http"

	atlas "go.mongodb.org/atlas/mongodbatlas"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
)

//go:generate mockgen -destination=../mocks/mock_telemetry.go -package=mocks github.com/mongodb/mongodb-atlas-cli/internal/store EventsSender

const urlPath = "api/private/v1.0/telemetry/events"

type EventsSender interface {
	SendEvents(body interface{}) error
}

func (s *Store) SendEvents(body interface{}) error {
	switch s.service {
	case config.CloudService:
		client := s.client.(*atlas.Client)
		request, err := client.NewRequest(s.ctx, http.MethodPost, urlPath, body)
		if err != nil {
			return err
		}
		_, err = client.Do(s.ctx, request, nil)
		return err
	default:
		return nil
	}
}
