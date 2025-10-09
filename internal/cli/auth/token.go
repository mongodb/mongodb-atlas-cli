// Copyright 2025 MongoDB Inc
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

package auth

import (
	"fmt"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/spf13/cobra"
)

type tokenOpts struct {
	name string
	cli.OutputOpts
}

var tokenTemplate = `{{.access_token}}`

func (*tokenOpts) GetConfig(configStore config.Store, profileName string) (map[string]string, error) {
	profileMap := configStore.GetProfileStringMap(profileName)

	if v := configStore.GetProfileValue(profileName, "access_token"); v != nil && v != "" {
		profileMap["access_token"] = v.(string)
	} else {
		return nil, fmt.Errorf("no access token found for profile %s", profileName)
	}

	return profileMap, nil
}

func (opts *tokenOpts) Run() error {
	// Create a new config proxy store
	configStore, err := config.NewStoreWithEnvOption(false)
	if err != nil {
		return fmt.Errorf("could not create config store: %w", err)
	}

	// get curren tprofile
	profile := config.Default()
	opts.name = profile.Name()

	mapConfig, err := opts.GetConfig(configStore, opts.name)
	if err != nil {
		return err
	}

	return opts.Print(mapConfig)
}

func TokenBuilder() *cobra.Command {
	opts := &tokenOpts{}
	opts.Template = tokenTemplate
	cmd := &cobra.Command{
		Use:    "token",
		Hidden: true,
		Short:  "Return the token for the current profile.",
		Example: `  # Return the token for the current profile:
  atlas auth token

  # Return the token for the current profile and save it to a file:
  atlas auth token > token.txt

  # Return the token for a specific profile:
  atlas auth token --profile <profile_name>
  `,
		Args: require.NoArgs,
		PreRun: func(cmd *cobra.Command, _ []string) {
			opts.OutWriter = cmd.OutOrStdout()
		},
		RunE: func(_ *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	opts.AddOutputOptFlags(cmd)

	return cmd
}
