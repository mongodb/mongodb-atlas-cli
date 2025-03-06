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

package main

import (
	"log"
	"os"
	"strings"

	"github.com/mongodb-labs/cobra2snooty"
	pluginCmd "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/root"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

const apiCommandName = "api"

func setDisableAutoGenTag(cmd *cobra.Command) {
	cmd.DisableAutoGenTag = true
	for _, cmd := range cmd.Commands() {
		setDisableAutoGenTag(cmd)
	}
}

func addExperimenalToAPICommands(cmd *cobra.Command) {
	var apiCommand *cobra.Command
	for _, subCommand := range cmd.Commands() {
		if subCommand.Use == apiCommandName {
			apiCommand = subCommand
		}
	}

	if apiCommand == nil {
		panic("api command not found!")
	}

	markExperimentalRecursively(apiCommand)
}

func markExperimentalRecursively(cmd *cobra.Command) {
	cmd.Short = "`experimental <https://www.mongodb.com/docs/atlas/cli/current/command/atlas-api/>`_: " + cmd.Short

	for _, subCommand := range cmd.Commands() {
		markExperimentalRecursively(subCommand)
	}
}

func updateAPICommandDescription(cmd *cobra.Command) {
	var apiCommand *cobra.Command
	for _, subCommand := range cmd.Commands() {
		if subCommand.Use == apiCommandName {
			apiCommand = subCommand
		}
	}

	if apiCommand == nil {
		panic("api command not found!")
	}

	updateLeafDescriptions(apiCommand)
}

func updateLeafDescriptions(cmd *cobra.Command) {
	if len(cmd.Commands()) == 0 {
		lines := strings.Split(cmd.Long, "\n")
		// Replace last line if it contains the extected text: "For more information and examples, see: <AtlasCLI docs url>"
		if strings.HasPrefix(lines[len(lines)-1], "For more information and examples, see: https://www.mongodb.com/docs/atlas/cli/current/command/") {
			lines = lines[:len(lines)-1]
			newLine := "For more information and examples, see the referenced API documentation linked above."
			lines = append(lines, newLine)
		}

		cmd.Long = strings.Join(lines, "\n")
	}

	for _, subCommand := range cmd.Commands() {
		updateLeafDescriptions(subCommand)
	}
}

func addAdditionalLongText(cmd *cobra.Command) {
	if additionalLongText, found := cmd.Annotations["DocsAdditionalLongText"]; found && additionalLongText != "" {
		cmd.Long += "\n\n"
		cmd.Long += additionalLongText
	}

	for _, cmd := range cmd.Commands() {
		addAdditionalLongText(cmd)
	}
}

func main() {
	if err := os.RemoveAll("./docs/command"); err != nil {
		log.Fatal(err)
	}

	const docsPermissions = 0766
	if err := os.MkdirAll("./docs/command", docsPermissions); err != nil {
		log.Fatal(err)
	}

	atlasBuilder := root.Builder()

	for _, cmd := range atlasBuilder.Commands() {
		if plugin.IsPluginCmd(cmd) && !isFirstClassPlugin(cmd) {
			atlasBuilder.RemoveCommand(cmd)
		}
	}

	atlasBuilder.InitDefaultCompletionCmd()

	setDisableAutoGenTag(atlasBuilder)
	addExperimenalToAPICommands(atlasBuilder)
	updateAPICommandDescription(atlasBuilder)
	addAdditionalLongText(atlasBuilder)

	if err := cobra2snooty.GenTreeDocs(atlasBuilder, "./docs/command"); err != nil {
		log.Fatal(err)
	}

	firstClassPaths := make([]string, 0, len(pluginCmd.FirstClassPlugins))
	for _, fcp := range pluginCmd.FirstClassPlugins {
		cmd := fcp.Commands
		for _, c := range cmd {
			filePath := "./docs/command/atlas-" + c.Name + ".txt"
			firstClassPaths = append(firstClassPaths, filePath)
		}
	}

	for _, filePath := range firstClassPaths {
		err := os.Remove(filePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func isFirstClassPlugin(command *cobra.Command) bool {
	for _, fcp := range pluginCmd.FirstClassPlugins {
		cmd := fcp.Commands
		for _, c := range cmd {
			if command.Name() == c.Name {
				return true
			}
		}
	}
	return false
}
