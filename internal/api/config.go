// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
)

type AuthenticatedConfigWrapper struct {
	authenticatedConfig store.AuthenticatedConfig
}

func NewAuthenticatedConfigWrapper(authenticatedConfig store.AuthenticatedConfig) *AuthenticatedConfigWrapper {
	return &AuthenticatedConfigWrapper{
		authenticatedConfig: authenticatedConfig,
	}
}

func (c *AuthenticatedConfigWrapper) GetAccessToken() (string, error) {
	token, err := c.authenticatedConfig.Token()
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func (c *AuthenticatedConfigWrapper) GetBaseURL() (string, error) {
	// If the profile has overwritten the URL, use that one
	if configURL := c.authenticatedConfig.OpsManagerURL(); configURL != "" {
		return configURL, nil
	}

	// If the service is cloud gov, use the cloud gov base url
	if c.authenticatedConfig.Service() == config.CloudGovService {
		return store.CloudGovServiceURL, nil
	}

	// By default, return the default base URL
	return store.CloudServiceURL, nil
}
