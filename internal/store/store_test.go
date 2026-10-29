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
	"testing"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/stretchr/testify/require"
	atlasauth "go.mongodb.org/atlas/auth"
)

type auth struct {
	username     string
	password     string
	refreshToken string
	clientID     string
	clientSecret string
	accessToken  *atlasauth.Token
}

func (a auth) Token() (*atlasauth.Token, error) {
	return a.accessToken, nil
}

func (a auth) RefreshToken() string {
	return a.refreshToken
}

func (a auth) PublicAPIKey() string {
	return a.username
}

func (a auth) PrivateAPIKey() string {
	return a.password
}

func (a auth) ClientID() string {
	return a.clientID
}

func (a auth) ClientSecret() string {
	return a.clientSecret
}

func (a auth) AuthType() config.AuthMechanism {
	if a.username != "" {
		return config.APIKeys
	}
	if a.accessToken != nil {
		return config.UserAccount
	}
	if a.clientID != "" {
		return config.ServiceAccount
	}
	return ""
}

var _ CredentialsGetter = &auth{}

func TestService(t *testing.T) {
	c, err := New(Service(config.CloudService))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.service != config.CloudService {
		t.Errorf("New() service = %s; expected %s", c.service, "cloud")
	}
}

func TestWithBaseURL(t *testing.T) {
	c, err := New(Service(config.CloudService), WithBaseURL("http://test"))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.baseURL != "http://test" {
		t.Errorf("New() baseURL = %s; expected %s", c.baseURL, "http://test")
	}
}

type testConfig struct {
	url string
	auth
}

func (testConfig) OpsManagerCACertificate() string {
	return ""
}

func (testConfig) OpsManagerSkipVerify() string {
	return "false"
}

func (testConfig) Service() string {
	return config.CloudService
}

func (c testConfig) OpsManagerURL() string {
	return c.url
}

var _ AuthenticatedConfig = &testConfig{}

func TestWithAuthentication(t *testing.T) {
	tests := []struct {
		name           string
		setTestProfile func(p *config.Profile)
	}{
		{
			name: "api keys",
			setTestProfile: func(p *config.Profile) {
				p.SetAuthType(config.APIKeys)
				p.SetPublicAPIKey("test-key")
				p.SetPrivateAPIKey("test-secret")
			},
		},
		{
			name: "service account",
			setTestProfile: func(p *config.Profile) {
				p.SetAuthType(config.ServiceAccount)
				p.SetClientID("id")
				p.SetClientSecret("secret")
			},
		},
		{
			name: "user account",
			setTestProfile: func(p *config.Profile) {
				p.SetAuthType(config.UserAccount)
				p.SetAccessToken("token")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up profile for testing
			profile := config.NewProfile("test", config.NewInMemoryStore())
			tt.setTestProfile(profile)
			config.SetDefaultProfile(profile)

			c, err := New(Service("cloud"), WithAuthentication())
			require.NoError(t, err)
			require.NotNil(t, c.httpClient)
			require.NotNil(t, c.httpClient.Transport)
			require.NotEqual(t, c.transport(), c.httpClient.Transport) // Check transport is not default
		})
	}
}

func TestWithContext(t *testing.T) {
	c, err := New(Service(config.CloudService))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.ctx != context.Background() {
		t.Errorf("New() got %v; expected %v", c.ctx, t.Context())
	}

	type myCustomType string
	var k, v myCustomType = "custom key", "custom value"

	ctx := context.WithValue(t.Context(), k, v)

	c, err = New(Service(config.CloudService), WithContext(ctx))
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if c.ctx != ctx {
		t.Errorf("New() got %v; expected %v", c.ctx, ctx)
	}
}
