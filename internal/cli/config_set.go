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

package cli

import (
	"fmt"

	"github.com/mongodb/mcli/internal/config"
	"github.com/mongodb/mcli/internal/search"
	"github.com/spf13/cobra"
)

type configSetOpts struct {
	*globalOpts
	prop string
	val  string
}

func (opts *configSetOpts) Run() error {
	config.Set(opts.prop, opts.val)
	if err := config.Save(); err != nil {
		return err
	}
	fmt.Printf("Updated prop '%s'\n", opts.prop)
	return nil
}

func ConfigSetBuilder() *cobra.Command {
	opts := &configSetOpts{
		globalOpts: newGlobalOpts(),
	}
	cmd := &cobra.Command{
		Use:   "set [prop] [val]",
		Short: "Configure the tool.",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return fmt.Errorf("accepts %d arg(s), received %d", 2, len(args))
			}
			if !search.StringInSlice(cmd.ValidArgs, args[0]) {
				return fmt.Errorf("invalid prop %q", args[0])
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
