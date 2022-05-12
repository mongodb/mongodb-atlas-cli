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
	"net/http"

	"github.com/mongodb/mongodb-atlas-cli/internal/config"
)

const urlPath = "api/private/v1.0/telemetry/events"

func SendEvents(ctx context.Context, body interface{}) error {
	if config.Service() != config.CloudService {
		// Only send events to Atlas - not to AtlasGov or OpsManager or CloudManager
		return nil
	}
	s, err := New(AuthenticatedPreset(config.Default()), WithContext(ctx), Telemetry())
	if err != nil {
		return err
	}
	client, err := s.GetAtlasClient()
	if err != nil {
		return err
	}
	request, err := client.NewRequest(ctx, http.MethodPost, urlPath, body)
	if err != nil {
		return err
	}
	_, err = client.Do(ctx, request, nil)
	return err
}
