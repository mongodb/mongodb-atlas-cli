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

package cli

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
)

var errUnsupportedService = errors.New("unsupported service")

func InitProfile(profile string) error {
	initAuthType()
	if profile != "" {
		return config.SetName(profile)
	} else if profile = config.GetString(flag.Profile); profile != "" {
		return config.SetName(profile)
	} else if availableProfiles := config.List(); len(availableProfiles) == 1 {
		return config.SetName(availableProfiles[0])
	}

	if !config.IsCloud() {
		return fmt.Errorf("%w: %s", errUnsupportedService, config.Service())
	}

	return nil
}

// initAuthType initializes the authentication type based on the current configuration.
// If the user has set credentials via environment variables and has not set
// 'MONGODB_ATLAS_AUTH_TYPE', it will set the auth type accordingly.
func initAuthType() {
	// If the auth type is already set, we don't need to do anything.
	authType := config.AuthType()
	if authType != "" {
		return
	}
	// If the auth type is not set, we try to determine it based on the available credentials.
	if config.PrivateAPIKey() != "" && config.PublicAPIKey() != "" {
		config.SetAuthType(config.APIKeys)
	}
	if config.AccessToken() != "" && config.RefreshToken() != "" {
		config.SetAuthType(config.UserAccount)
	}
	if config.ClientID() != "" && config.ClientSecret() != "" {
		config.SetAuthType(config.ServiceAccount)
	}
}
