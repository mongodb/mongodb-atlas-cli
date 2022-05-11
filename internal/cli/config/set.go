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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/search"
	"github.com/mongodb/mongodb-atlas-cli/internal/validate"
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
		Long: fmt.Sprintf(`Configure specific properties of the profile.
Available properties include: %v.`, config.Properties()),
		Args: func(cmd *cobra.Command, args []string) error {
			if err := require.ExactArgs(argsN)(cmd, args); err != nil {
				return err
			}
			if !search.StringInSlice(cmd.ValidArgs, args[0]) {
				return fmt.Errorf("invalid property: %q", args[0])
			}
			return nil
		},
		Example: fmt.Sprintf(`
  Set Ops Manager Base URL in the profile myProfile:
  $ %[1]s config set ops_manager_url http://localhost:30700/ -P myProfile
  Set Organization ID in the default profile:
  $ %[1]s config set org_id 5dd5aaef7a3e5a6c5bd12de4
  Set path for the MongoDB Shell in the default profile:
  $ %[1]s config set mongosh_path /usr/local/bin/mongosh`, config.BinName()),
		Annotations: map[string]string{
			"args":             "propertyName,value",
			"requiredArgs":     "propertyName,value",
			"propertyNameDesc": "Property to set in the profile.",
			"valueDesc":        "Value for the property to set in the profile.",
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
