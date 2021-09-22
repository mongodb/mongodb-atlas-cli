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
	"strings"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/cli/require"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/search"
	"github.com/mongodb/mongocli/internal/validate"
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
	var value interface{}
	value = opts.val
	if search.StringInSlice(config.BooleanProperties(), opts.prop) {
		value = config.IsTrue(opts.val)
	}
	if search.StringInSlice(config.GlobalProperties(), opts.prop) {
		opts.store.SetGlobal(opts.prop, value)
	} else {
		opts.store.Set(opts.prop, value)
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
		Long: `You can set the following properties in the specified profile:
project_id - Unique identifier of a project.
org_id - Unique identifier of an organization.
service - MongoDB service type. Valid values: cloud|cloudgov|cloud-manager|ops-manager
public_api_key - Public API key for programmatic access.
private_api_key - Private API key for programmatic access.
output - Output fields and format. Valid values: json|json-path|go-template|go-template-file
ops_manager_url - Ops Manager only. Base URL for API calls. The URL must end with a forward slash (/).
base_url - Base URL for API calls. The URL must end with a forward slash (/).
ops_manager_ca-certificate - Ops Manager only. Path on your local system to the PEM-encoded Certificate Authority (CA) certificate used to sign the client and Ops Manager TLS certificates.
ops_manager_skip_verify - Ops Manager only. When set to yes, the Ops Manager CA TLS certificate is not verified. This prevents your connections from being rejected due to an invalid certificate. This is insecure and not recommended in production environments. Valid values: yes|no
mongosh_path - Path to the MongoDB shell (mongosh) on your system. Default value: /usr/local/bin/mongosh`,
		Example: `
  Set Ops Manager Base URL in the profile myProfile:
  $ mongocli config set ops_manager_url http://localhost:30700/ -P myProfile

  Set Organization ID in the default profile:
  $ mognocli config set org_id 5dd5aaef7a3e5a6c5bd12de4

  Set path for the MongoDB Shell in the default profile:
  $ mongocli config set mongosh_path /usr/local/bin/mongosh`,
		Annotations: map[string]string{
			"args":             "propertyName,value",
			"requiredArgs":     "propertyName,value",
			"propertyNameDesc": "Property to set in the profile.",
			"valueDesc":        "Value for the property to set in the profile.",
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if err := require.ExactArgs(argsN)(cmd, args); err != nil {
				return err
			}
			if !search.StringInSlice(cmd.ValidArgs, args[0]) {
				return fmt.Errorf("invalid property: %q", args[0])
			}
			return nil
		},
		ValidArgs: config.Properties(),
		RunE: func(cmd *cobra.Command, args []string) error {
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
