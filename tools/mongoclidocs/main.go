// Copyright 2021 MongoDB Inc
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
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/root/mongocli"
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

	var profile string
	const docsPermissions = 0766
	if err := os.MkdirAll("./docs/command", docsPermissions); err != nil {
		log.Fatal(err)
	}

	mongocliBuilder := mongocli.Builder(&profile, []string{})
	mongocliBuilder.InitDefaultCompletionCmd()
	removeDeprecateStringAtlasCommand(mongocliBuilder)
	setDisableAutoGenTag(mongocliBuilder)

	if err := cobra2snooty.GenTreeDocs(mongocliBuilder, "./docs/command"); err != nil {
		log.Fatal(err)
	}
}

func removeDeprecateStringAtlasCommand(cmd *cobra.Command) {
	for _, c := range cmd.Commands() {
		if c.Use == "atlas" {
			c.Long = ""
			return
		}
	}
}
