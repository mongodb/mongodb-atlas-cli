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
	"strings"

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
	return ClientIDForOpsManagerURL(c.profile.OpsManagerURL())
}

// FlowWithConfig constructs the legacy device-code flow using the supplied profile.
func FlowWithConfig(profile ServiceConfig, client *http.Client, version string) (*atlasauth.Config, error) {
	return transport.FlowWithConfig(legacyDeviceCodeServiceGetter{profile: profile}, client, version)
}

// ClientIDForOpsManagerURL returns the legacy device-code client ID for a known environment.
// An empty result means the core transport's production or Gov default should be used.
func ClientIDForOpsManagerURL(opsManagerURL string) string {
	switch strings.TrimRight(opsManagerURL, "/") {
	case "https://localhost:8080", "http://cloud-local.mmscloudteam.com", "https://cloud-dev.mongodb.com", "https://cloud-qa.mongodb.com":
		return devClientID
	case "https://cloud-internal.mongodb.com":
		return internalClientID
	case "https://cloud-dev.mongodbgov.com":
		return govDevClientID
	case "https://cloud.mongodb.com":
		return transport.ClientID
	case "https://cloud.mongodbgov.com":
		return transport.GovClientID
	default:
		// Leave the client ID unset so transport.FlowWithConfig uses its production default.
		return ""
	}
}
