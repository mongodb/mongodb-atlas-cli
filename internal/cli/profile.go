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
	"os"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
)

var errUnsupportedService = errors.New("unsupported service")

func InitProfile(profile string) error {
	if profile != "" {
		if err := config.SetName(profile); err != nil {
			return err
		}
	} else if profile = config.GetString(flag.Profile); profile != "" {
		if err := config.SetName(profile); err != nil {
			return err
		}
	} else if availableProfiles := config.List(); len(availableProfiles) == 1 {
		if err := config.SetName(availableProfiles[0]); err != nil {
			return err
		}
	}

	if !config.IsCloud() {
		return fmt.Errorf("%w: %s", errUnsupportedService, config.Service())
	}

	initAuthType()

	return nil
}

// initAuthType sets the AuthType if environment variables are being used.
// This will override authType if it already exists from a profile in config,
// which is desired behavior as environment variables take precedence.
func initAuthType() {
	if envVarAuthType := detectEnvVars(); envVarAuthType != "" {
		config.SetAuthType(envVarAuthType)
		return
	}
}

// detectEnvVars detects environment variables and returns the appropriate
// authentication mechanism.
func detectEnvVars() config.AuthMechanism {
	// Check for Service Account credentials
	if (os.Getenv("MONGODB_ATLAS_CLIENT_ID") != "" && os.Getenv("MONGODB_ATLAS_CLIENT_SECRET") != "") ||
		(os.Getenv("MCLI_CLIENT_ID") != "" && os.Getenv("MCLI_CLIENT_SECRET") != "") {
		return config.ServiceAccount
	}
	// Check for API Key credentials
	if (os.Getenv("MONGODB_ATLAS_PUBLIC_API_KEY") != "" && os.Getenv("MONGODB_ATLAS_PRIVATE_API_KEY") != "") ||
		(os.Getenv("MCLI_PUBLIC_API_KEY") != "" && os.Getenv("MCLI_PRIVATE_API_KEY") != "") {
		return config.APIKeys
	}
	// Check for User Account credentials
	if (os.Getenv("MONGODB_ATLAS_ACCESS_TOKEN") != "" && os.Getenv("MONGODB_ATLAS_REFRESH_TOKEN") != "") ||
		(os.Getenv("MCLI_ACCESS_TOKEN") != "" && os.Getenv("MCLI_REFRESH_TOKEN") != "") {
		return config.UserAccount
	}
	return ""
}
