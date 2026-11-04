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
	"errors"
	"net/http"
	"strings"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	atlasauth "go.mongodb.org/atlas/auth"
)

const (
	devClientID      = "0oadn4hoajpzxeSEy357"
	internalClientID = "0oayl6fukpaL0ilPP297"
	govDevClientID   = "0oafm2iv10KRawp5U357"
)

// ServiceConfig provides the profile settings used to configure the legacy device-code flow.
type ServiceConfig interface {
	Service() string
	OpsManagerURL() string
	ClientID() string
}

type legacyDeviceCodeServiceGetter struct {
	profile ServiceConfig
}

func (c legacyDeviceCodeServiceGetter) Service() string       { return c.profile.Service() }
func (c legacyDeviceCodeServiceGetter) OpsManagerURL() string { return c.profile.OpsManagerURL() }
func (c legacyDeviceCodeServiceGetter) AccountURL() string {
	if profile, ok := c.profile.(interface{ AccountURL() string }); ok {
		return profile.AccountURL()
	}
	return ""
}

func (c legacyDeviceCodeServiceGetter) ClientID() string {
	if clientID := c.profile.ClientID(); clientID != "" {
		return clientID
	}
	return clientIDForOpsManagerURL(c.profile.OpsManagerURL())
}

// FlowWithConfig constructs the legacy device-code flow using the supplied profile.
func FlowWithConfig(profile ServiceConfig, client *http.Client, version string) (*atlasauth.Config, error) {
	return transport.FlowWithConfig(legacyDeviceCodeServiceGetter{profile: profile}, client, version)
}

// HTTPClientFromProfile creates an HTTP client using the profile's authentication settings.
func HTTPClientFromProfile(profile transport.ProfileProvider, version string, base http.RoundTripper) (*http.Client, error) {
	if profile.AuthType() != config.UserAccountLegacy {
		return transport.HTTPClientFromProfile(profile, version, base)
	}

	token, err := profile.Token()
	if err != nil {
		return nil, err
	}
	if token == nil {
		return &http.Client{Transport: base}, nil
	}

	accessTokenTransport, err := newLegacyAccessTokenTransport(profile, token, base, version, func(t *atlasauth.Token) error {
		profile.SetAccessToken(t.AccessToken)
		profile.SetRefreshToken(t.RefreshToken)
		return profile.Save()
	})
	if err != nil {
		return nil, err
	}
	return &http.Client{Transport: accessTokenTransport}, nil
}

func newLegacyAccessTokenTransport(profile ServiceConfig, token *atlasauth.Token, base http.RoundTripper, version string, saveToken func(*atlasauth.Token) error) (http.RoundTripper, error) {
	if token == nil {
		return nil, errors.New("token is nil")
	}
	client := &http.Client{Transport: transport.Default()}
	flow, err := FlowWithConfig(profile, client, version)
	if err != nil {
		return nil, err
	}
	return transport.NewAccessTokenTransportWithConfig(token, base, flow, saveToken)
}

func clientIDForOpsManagerURL(opsManagerURL string) string {
	switch strings.TrimRight(opsManagerURL, "/") {
	case "https://localhost:8080", "http://cloud-local.mmscloudteam.com", "https://cloud-dev.mongodb.com", "https://cloud-qa.mongodb.com":
		return devClientID
	case "https://cloud-internal.mongodb.com":
		return internalClientID
	case "https://cloud-dev.mongodbgov.com":
		return govDevClientID
	default:
		// Leave the client ID unset so transport.FlowWithConfig uses its production default.
		return ""
	}
}
