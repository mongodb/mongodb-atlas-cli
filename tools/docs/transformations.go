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
	"regexp"

	"github.com/spf13/cobra"
)

func setDisableAutoGenTag(cmd *cobra.Command) {
	cmd.DisableAutoGenTag = true
}

func isApiCommand(cmd *cobra.Command) bool {
	return regexp.MustCompile(" api( |$)").MatchString(cmd.CommandPath())
}

func removeFirstClassPluginCommands(cmd *cobra.Command) {
	if isFirstClassPlugin(cmd) {
		cmd.RemoveCommand(cmd)
	}
}

func applyTransformations(cmd *cobra.Command) {
	setDisableAutoGenTag(cmd)
	removeFirstClassPluginCommands(cmd)
	applyAutogenTransformations(cmd)

	for _, subCmd := range cmd.Commands() {
		applyTransformations(subCmd)
	}
}
