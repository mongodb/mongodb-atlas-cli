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
	"errors"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/set"
	"github.com/spf13/cobra"
)

const (
	sourcePluginName = "sourcePluginName"
)

type Opts struct {
	plugins          []*plugin.Plugin
	existingCommands []*cobra.Command
}

// find a plugin given the input argument of a plugin command
// the input arg can be the plugin name, the github values (<repo-owner>/<repo-name>) or the entire github URL of the plugin.
func (opts *Opts) findPluginWithArg(arg string) (*plugin.Plugin, error) {
	// try to parse input to github values
	// if parsing fails it will be assumed that the input is the plugin name
	githubValues, err := parseGithubReleaseValues(arg)
	var pluginToUninstall *plugin.Plugin

	if err == nil {
		pluginToUninstall, err = opts.findPluginWithGithubValues(githubValues.owner, githubValues.name)
	} else {
		pluginToUninstall, err = opts.findPluginWithName(arg)
	}

	if err != nil {
		return nil, err
	}

	return pluginToUninstall, nil
}

func createExistingCommandsSet(existingCommands []*cobra.Command) set.Set[string] {
	existingCommandsSet := set.NewSet[string]()
	for _, cmd := range existingCommands {
		existingCommandsSet.Add(cmd.Name())
		for _, alias := range cmd.Aliases {
			existingCommandsSet.Add(alias)
		}
	}

	return existingCommandsSet
}

func (opts *Opts) findPluginWithGithubValues(owner string, name string) (*plugin.Plugin, error) {
	for _, p := range opts.plugins {
		if p.Github != nil && p.Github.Equals(owner, name) {
			return p, nil
		}
	}
	return nil, fmt.Errorf(`could not find plugin with github values %s/%s`, owner, name)
}

func (opts *Opts) findPluginWithName(name string) (*plugin.Plugin, error) {
	for _, p := range opts.plugins {
		if p.Name == name {
			return p, nil
		}
	}
	return nil, fmt.Errorf(`could not find plugin with name %s`, name)
}

func RegisterCommands(rootCmd *cobra.Command) {
	plugins := plugin.GetAllValidPlugins(createExistingCommandsSet(rootCmd.Commands()))

	for _, p := range plugins {
		rootCmd.AddCommand(p.GetCobraCommands()...)
	}

	rootCmd.AddCommand(getFirstClassPluginCommands(plugins)...)

	rootCmd.AddCommand(Builder(plugins, rootCmd.Commands()))
}

func validateManifest(manifest *plugin.Manifest) error {
	if valid, errorList := manifest.IsValid(); !valid {
		var manifestErrorLog strings.Builder
		manifestErrorLog.WriteString(fmt.Sprintf("plugin in directory \"%s\" could not be loaded due to the following error(s) in the manifest.yaml:\n", manifest.PluginDirectoryPath))
		for _, err := range errorList {
			manifestErrorLog.WriteString(fmt.Sprintf("\t- %s\n", err.Error()))
		}

		return errors.New(manifestErrorLog.String())
	}
	return nil
}

func Builder(plugins []*plugin.Plugin, existingCommands []*cobra.Command) *cobra.Command {
	const use = "plugin"
	cmd := &cobra.Command{
		Use:     use,
		Aliases: cli.GenerateAliases(use),
		Short:   "Manage plugins for the AtlasCLI.",
	}

	pluginOpts := &Opts{
		plugins:          plugins,
		existingCommands: existingCommands,
	}

	cmd.AddCommand(
		ListBuilder(pluginOpts),
		InstallBuilder(pluginOpts),
		UninstallBuilder(pluginOpts),
		UpdateBuilder(pluginOpts),
	)

	return cmd
}
