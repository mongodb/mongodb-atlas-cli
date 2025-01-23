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

	"github.com/mongodb-labs/cobra2snooty"
	pluginCmd "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/root"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/spf13/cobra"
)

func setDisableAutoGenTag(cmd *cobra.Command) {
	cmd.DisableAutoGenTag = true
	for _, cmd := range cmd.Commands() {
		setDisableAutoGenTag(cmd)
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
		if plugin.IsPluginCmd(cmd) && !isFCP(cmd) {
			atlasBuilder.RemoveCommand(cmd)
		}
	}

	atlasBuilder.InitDefaultCompletionCmd()

	setDisableAutoGenTag(atlasBuilder)

	if err := cobra2snooty.GenTreeDocs(atlasBuilder, "./docs/command"); err != nil {
		log.Fatal(err)
	}

	firstClassPaths := []string{
		"./docs/command/atlas-kubernetes.txt",
	}

	for _, filePath := range firstClassPaths {
		err := os.Remove(filePath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func isFCP(command *cobra.Command) bool {
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
