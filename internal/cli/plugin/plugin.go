// Copyright 2024 MongoDB Inc
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

package plugin

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

func RegisterCommands(rootCmd *cobra.Command) {
	plugins := plugin.GetAllValidPlugins(rootCmd.Commands())

	rootCmd.AddCommand(Builder(plugins))

	for _, p := range plugins {
		rootCmd.AddCommand(p.GetCobraCommands()...)
	}
}

func Builder(plugins []*plugin.Plugin) *cobra.Command {
	const use = "plugin"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage plugins for the AtlasCLI.",
	}

	cmd.AddCommand(
		ListBuilder(plugins),
		InstallBuilder(plugins))

	return cmd
}
