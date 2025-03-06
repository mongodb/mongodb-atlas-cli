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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func applyAutogenTransformationFlags(cmd *cobra.Command, flag *pflag.Flag) {
	if cmd.CommandPath() == "atlas api clusters createCluster" {
		flag.Usage = "[DEBUG] " + flag.Usage
	}
}

func markExperimenalToAPICommands(cmd *cobra.Command) {
	cmd.Short = "`experimental <https://www.mongodb.com/docs/atlas/cli/current/command/atlas-api/>`_: " + cmd.Short
}

func updateAPICommandDescription(cmd *cobra.Command) {
	if len(cmd.Commands()) > 0 {
		return
	}

	lines := strings.Split(cmd.Long, "\n")
	// Replace last line if it contains the extected text: "For more information and examples, see: <AtlasCLI docs url>"
	if strings.HasPrefix(lines[len(lines)-1], "For more information and examples, see: https://www.mongodb.com/docs/atlas/cli/current/command/") {
		lines = lines[:len(lines)-1]
		newLine := "For more information and examples, see the referenced API documentation linked above."
		lines = append(lines, newLine)
	}

	cmd.Long = strings.Join(lines, "\n")
}

func applyAutogenTransformations(cmd *cobra.Command) {
	if !isApiCommand(cmd) {
		return
	}

	markExperimenalToAPICommands(cmd)
	updateAPICommandDescription(cmd)
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		applyAutogenTransformationFlags(cmd, flag)
	})
}
