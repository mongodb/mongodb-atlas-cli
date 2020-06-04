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

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
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
	opts.store.Set(opts.prop, opts.val)
	if err := opts.store.Save(); err != nil {
		return err
	}
	fmt.Printf("Updated property '%s'\n", opts.prop)
	return nil
}

func SetBuilder() *cobra.Command {
	opts := &SetOpts{
		store: config.Default(),
	}
	cmd := &cobra.Command{
		Use:   "set <property> <value>",
		Short: description.ConfigSetDescription,
		Long:  fmt.Sprintf(description.ConfigSetLong, config.Properties()),
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("accepts %d arg(s), received %d", 2, len(args))
			}
			if !search.StringInSlice(cmd.ValidArgs, args[0]) {
				return fmt.Errorf("invalid property: %q", args[0])
			}
			return nil
		},
		ValidArgs: config.Properties(),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.prop = args[0]
			opts.val = args[1]
			return opts.Run()
		},
	}

	return cmd
}
