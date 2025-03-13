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
	_ "embed"
	"regexp"
	"strings"

	pluginCmd "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/internal/metadatatypes"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

//go:embed api_docs_long_text.txt
var atlasAPIDocsAdditionalLongText string

var additionalLongTexts = map[string]string{
	"atlas api": atlasAPIDocsAdditionalLongText,
}

func addAdditionalLongText(cmd *cobra.Command) {
	commandPath := cmd.CommandPath()
	if additionalLongText, found := additionalLongTexts[commandPath]; found && additionalLongText != "" {
		cmd.Long += "\n\n"
		cmd.Long += additionalLongText
	}
}

func isAPICommand(cmd *cobra.Command) bool {
	return regexp.MustCompile("^atlas api( |$)").MatchString(cmd.CommandPath())
}

func setDisableAutoGenTag(cmd *cobra.Command) {
	cmd.DisableAutoGenTag = true
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

func removePluginCommands(cmd *cobra.Command) {
	if plugin.IsPluginCmd(cmd) && !isFirstClassPlugin(cmd) {
		cmd.Parent().RemoveCommand(cmd)
	}
}

func replaceFlagUsage(cmd *cobra.Command, f *pflag.Flag) {
	operationID := cmd.Annotations["operationId"]
	if operationID == "" {
		return
	}

	cmdMetadata, ok := metadata[operationID]
	if !ok {
		return
	}

	paramMetadata, ok := cmdMetadata.Parameters[f.Name]
	if !ok {
		return
	}

	f.Usage = paramMetadata.Usage
}

func buildExamples(cmd *cobra.Command, examples map[string][]metadatatypes.Example) string {
	if len(examples) == 0 {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(`Examples
-----------------

.. tabs::
`)

	for version, exList := range examples {
		for _, ex := range exList {
			sb.WriteString("   .. tab:: ")
			if ex.Name == "" {
				sb.WriteString("Example")
			} else {
				sb.WriteString(ex.Name)
			}
			sb.WriteString("\n      :tabid: ")
			sb.WriteString(version)
			sb.WriteString("_")
			if ex.Source == "-" {
				sb.WriteString("default")
			} else {
				sb.WriteString(strings.ToLower(strings.ReplaceAll(ex.Source, " ", "_")))
			}
			sb.WriteString("\n")
			if ex.Description != "" {
				sb.WriteString("      " + ex.Description + "\n")
			}
			sb.WriteString("\n      .. code-block::\n\n")
			if ex.Value != "" {
				sb.WriteString("         cat <<EOF > payload.json\n")
				lines := strings.Split(ex.Value, "\n")
				for _, line := range lines {
					sb.WriteString("            " + line + "\n")
				}
				sb.WriteString("         EOF\n")
			}
			sb.WriteString("         " + cmd.CommandPath())
			sb.WriteString(" --version " + version)
			if ex.Value != "" {
				sb.WriteString(" --file payload.json")
			}
			for flagName, flagValue := range ex.Flags {
				sb.WriteString(" --" + flagName + " " + flagValue)
			}
			sb.WriteString("\n\n      .. Code end marker, please don't delete this comment\n\n")
		}
	}

	return sb.String()
}

func updateExamples(cmd *cobra.Command) error {
	operationID := cmd.Annotations["operationId"]
	if operationID == "" {
		return nil
	}

	cmdMetadata, ok := metadata[operationID]
	if !ok || cmdMetadata.Examples == nil {
		return nil
	}

	cmd.Example = buildExamples(cmd, cmdMetadata.Examples)

	return nil
}

func applyTransformations(cmd *cobra.Command) error {
	setDisableAutoGenTag(cmd)
	removePluginCommands(cmd)
	addAdditionalLongText(cmd)

	if isAPICommand(cmd) {
		markExperimenalToAPICommands(cmd)
		updateAPICommandDescription(cmd)
		if err := updateExamples(cmd); err != nil {
			return err
		}
	}

	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		replaceFlagUsage(cmd, f)
	})

	for _, subCmd := range cmd.Commands() {
		if err := applyTransformations(subCmd); err != nil {
			return err
		}
	}

	return nil
}
