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

package config

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/cobra"
)

type SetOpts struct {
	cli.GlobalOpts
	prop  string
	val   string
	store config.SetSaver
}

func (opts *SetOpts) Run() error {
	if strings.HasSuffix(opts.prop, "_url") {
		if err := validate.URL(opts.val); err != nil {
			return err
		}
	} else if strings.HasSuffix(opts.prop, "_id") {
		if err := validate.ObjectID(opts.val); err != nil {
			return err
		}
	}
	var value any
	value = opts.val
	if slices.Contains(config.BooleanProperties(), opts.prop) {
		value = config.IsTrue(opts.val)
	}
	if slices.Contains(config.GlobalProperties(), opts.prop) {
		opts.store.SetGlobal(opts.prop, value)
	} else {
		opts.store.Set(opts.prop, value)
	}

	if opts.prop == config.TelemetryEnabledProperty && mongosh.Detect() {
		err := mongosh.SetTelemetry(value.(bool))
		if err != nil {
			return fmt.Errorf("error enabling telemetry on mongosh: %w", err)
		}
	}

	if err := opts.store.Save(); err != nil {
		return err
	}
	fmt.Printf("Updated property '%s'\n", opts.prop)
	return nil
}

func SetBuilder() *cobra.Command {
	const argsN = 2
	cmd := &cobra.Command{
		Use:   "set <propertyName> <value>",
		Short: "Configure specific properties of a profile.",
		Args: func(cmd *cobra.Command, args []string) error {
			if err := require.ExactArgs(argsN)(cmd, args); err != nil {
				return err
			}
			if !slices.Contains(cmd.ValidArgs, args[0]) {
				return fmt.Errorf("invalid property: %q", args[0])
			}
			return nil
		},
		Example: `
  Set the organization ID in the default profile to 5dd5aaef7a3e5a6c5bd12de4:
  atlas config set org_id 5dd5aaef7a3e5a6c5bd12de4`,
		Annotations: map[string]string{
			"propertyNameDesc": "Property to set in the profile. Valid values for Atlas CLI are project_id, org_id, service, public_api_key, private_api_key, output, mongosh_path, skip_update_check, telemetry_enabled, access_token, and refresh_token.",
			"valueDesc":        "Value for the property to set in the profile.",
		},
		ValidArgs: config.Properties(),
		RunE: func(_ *cobra.Command, args []string) error {
			opts := &SetOpts{
				store: config.Default(),
				prop:  args[0],
				val:   args[1],
			}
			return opts.Run()
		},
	}

	return cmd
}
