// Copyright 2022 MongoDB Inc
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
		if plugin.IsPluginCmd(cmd) || pluginCmd.IsFirstClassPluginCmd(cmd) {
			atlasBuilder.RemoveCommand(cmd)
		}
	}

	atlasBuilder.InitDefaultCompletionCmd()

	setDisableAutoGenTag(atlasBuilder)

	if err := cobra2snooty.GenTreeDocs(atlasBuilder, "./docs/command"); err != nil {
		log.Fatal(err)
	}
}
