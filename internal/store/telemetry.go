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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
)

//go:generate mockgen -destination=../mocks/mock_telemetry.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store EventsSender,UnauthEventsSender

type EventsSender interface {
	SendEvents(body any) error
}

type UnauthEventsSender interface {
	SendUnauthEvents(body any) error
}

func (s *Store) SendEvents(body any) error {
	switch s.service {
	case config.CloudService:
		client := s.client
		request, err := client.NewRequest(s.ctx, http.MethodPost, "api/private/v1.0/telemetry/events", body)
		if err != nil {
			return err
		}
		_, err = client.Do(s.ctx, request, nil)
		return err
	default:
		return nil
	}
}

func (s *Store) SendUnauthEvents(body any) error {
	switch s.service {
	case config.CloudService:
		client := s.client
		request, err := client.NewRequest(s.ctx, http.MethodPost, "api/private/unauth/telemetry/events", body)
		if err != nil {
			return err
		}
		_, err = client.Do(s.ctx, request, nil)
		return err
	default:
		return nil
	}
}
