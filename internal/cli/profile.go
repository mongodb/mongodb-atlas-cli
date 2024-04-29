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
