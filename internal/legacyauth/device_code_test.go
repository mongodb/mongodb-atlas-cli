// Copyright 2026 MongoDB Inc
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

package legacyauth

import (
	"net/http"
	"os"
	"testing"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFlowWithConfig_ClientIDInference(t *testing.T) {
	tests := []struct {
		name          string
		service       string
		opsManagerURL string
		clientID      string
		envClientID   string
		wantClientID  string
	}{
		{
			name:          "local Bazel",
			service:       config.CloudService,
			opsManagerURL: "https://localhost:8080/",
			wantClientID:  devClientID,
		},
		{
			name:          "local Fern",
			service:       config.CloudService,
			opsManagerURL: "http://cloud-local.mmscloudteam.com/",
			wantClientID:  devClientID,
		},
		{
			name:          "cloud dev without trailing slash",
			service:       config.CloudService,
			opsManagerURL: "https://cloud-dev.mongodb.com",
			wantClientID:  devClientID,
		},
		{
			name:          "cloud QA",
			service:       config.CloudService,
			opsManagerURL: "https://cloud-qa.mongodb.com/",
			wantClientID:  devClientID,
		},
		{
			name:          "cloud internal",
			service:       config.CloudService,
			opsManagerURL: "https://cloud-internal.mongodb.com/",
			wantClientID:  internalClientID,
		},
		{
			name:          "cloud production keeps default",
			service:       config.CloudService,
			opsManagerURL: "https://cloud.mongodb.com/",
			wantClientID:  transport.ClientID,
		},
		{
			name:          "Gov dev",
			service:       config.CloudGovService,
			opsManagerURL: "https://cloud-dev.mongodbgov.com/",
			wantClientID:  govDevClientID,
		},
		{
			name:          "Gov production keeps default",
			service:       config.CloudGovService,
			opsManagerURL: "https://cloud.mongodbgov.com/",
			wantClientID:  transport.GovClientID,
		},
		{
			name:          "explicit MCLI_CLIENT_ID takes precedence",
			service:       config.CloudService,
			opsManagerURL: "https://cloud-dev.mongodb.com/",
			envClientID:   "explicit-client-id",
			wantClientID:  "explicit-client-id",
		},
		{
			name:          "profile client ID takes precedence",
			service:       config.CloudService,
			opsManagerURL: "https://cloud-dev.mongodb.com/",
			clientID:      "profile-client-id",
			wantClientID:  "profile-client-id",
		},
		{
			name:          "MCLI_CLIENT_ID overrides profile client ID",
			service:       config.CloudService,
			opsManagerURL: "https://cloud-dev.mongodb.com/",
			clientID:      "profile-client-id",
			envClientID:   "environment-client-id",
			wantClientID:  "environment-client-id",
		},
		{
			name:          "unknown Ops Manager URL keeps default",
			service:       config.CloudService,
			opsManagerURL: "https://cloud-stage.mongodb.com/",
			wantClientID:  transport.ClientID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("MCLI_CLIENT_ID", tt.envClientID)
			profile := testServiceConfig{
				service:       tt.service,
				opsManagerURL: tt.opsManagerURL,
				clientID:      tt.clientID,
			}

			flow, err := FlowWithConfig(profile, &http.Client{Transport: transport.Default()}, "test")
			require.NoError(t, err)
			assert.Equal(t, tt.wantClientID, flow.ClientID)
		})
	}
}

type testServiceConfig struct {
	service       string
	opsManagerURL string
	clientID      string
	accountURL    string
}

func (c testServiceConfig) Service() string       { return c.service }
func (c testServiceConfig) OpsManagerURL() string { return c.opsManagerURL }
func (c testServiceConfig) ClientID() string {
	if clientID := os.Getenv("MCLI_CLIENT_ID"); clientID != "" {
		return clientID
	}
	if c.clientID != "" {
		return c.clientID
	}
	return ""
}
func (c testServiceConfig) AccountURL() string { return c.accountURL }
