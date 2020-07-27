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

	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/spf13/cobra"
)

type ListOpts struct {
	name string
}

func (opts *ListOpts) Run() error {
	if !config.Exists(opts.name) {
		return fmt.Errorf("no profile with name '%s'", opts.name)
	}
	config.SetName(opts.name)
	c := config.Map()
	for _, k := range config.SortedKeys() {
		fmt.Printf("%s = %s\n", k, c[k])
	}

	return nil
}

func DescribeBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:   "describe <name>",
		Short: description.ConfigDescribe,
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run()
		},
	}

	return cmd
}
