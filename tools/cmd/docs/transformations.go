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
	"cmp"
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	pluginCmd "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/plugin"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/internal/metadatatypes"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Snooty official Style documentation: https://www.mongodb.com/docs/meta/style-guide
//
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
	// Snooty does not support the string "|---|---|---|---|" that we use in some API field description to generate a table.
	// Snooty error: ERROR(): Substitution reference could not be replaced: "|---|"
	usage := strings.ReplaceAll(paramMetadata.Usage, "|---|---|---|---|", "")

	// Snooty error: ERROR() Malformed external link Did you mean: ` <database>.<collection>`
	usage = strings.ReplaceAll(usage, "`<database>.<collection>`", "``<database>.<collection>``")
	f.Usage = usage
}

func sortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

func countExamples(examples map[string][]metadatatypes.Example) int {
	count := 0
	for _, exs := range examples {
		count += len(exs)
	}
	return count
}

func buildExamples(cmd *cobra.Command, examples map[string][]metadatatypes.Example) (string, error) { //nolint:gocyclo // cyclomatic complexity is high, but the code is readable and only used for documentation generation
	if len(examples) == 0 {
		return "", nil
	}

	var sb strings.Builder
	sb.WriteString(`Examples
--------

`)

	tabs := countExamples(examples) != 1
	if tabs {
		sb.WriteString(`.. tabs::

`)
	}

	exampleIdx := 0
	for _, version := range sortedKeys(examples) {
		for _, ex := range examples[version] {
			source := strings.ToLower(strings.ReplaceAll(ex.Source, " ", "_"))
			if ex.Source == "-" {
				source = "default"
			}

			if tabs {
				sb.WriteString("   .. tab:: ")
				if ex.Name == "" {
					sb.WriteString("Example")
					if exampleIdx > 0 {
						sb.WriteString(" ")
						sb.WriteString(strconv.Itoa(exampleIdx))
					}
					exampleIdx++
				} else {
					sb.WriteString(ex.Name)
				}
				sb.WriteString("\n      :tabid: ")
				sb.WriteString(version)
				sb.WriteString("_")
				sb.WriteString(source)
				sb.WriteString("\n\n")
			}

			if ex.Value != "" {
				if tabs {
					sb.WriteString("      ")
				}
				sb.WriteString(ex.Description + "\n\n")

				if tabs {
					sb.WriteString("      ")
				}
				sb.WriteString("Create the file below and save it as ``payload.json``\n\n")

				includePayloadFileName := fmt.Sprintf("%s-%s-%s-payload.json", strings.ReplaceAll(cmd.CommandPath(), " ", "-"), version, source)
				const permissions = 0600
				err := os.WriteFile(fmt.Sprintf("%s/%s", includesLocalPath, includePayloadFileName), []byte(ex.Value), permissions)
				if err != nil {
					return "", fmt.Errorf("failed to write payload file %s: %w", includePayloadFileName, err)
				}

				if tabs {
					sb.WriteString("      ")
				}
				sb.WriteString(fmt.Sprintf("   .. literalinclude:: %s/%s\n", includesImportPath, includePayloadFileName))
				if tabs {
					sb.WriteString("      ")
				}
				sb.WriteString("      :language: shell\n")
				if tabs {
					sb.WriteString("      ")
				}
				sb.WriteString("After creating ``payload.json``, run the command below in the same directory.\n\n")
			} else if tabs {
				sb.WriteString("      ")
			}

			var cmdSB strings.Builder

			cmdSB.WriteString(cmd.CommandPath())
			cmdSB.WriteString(" --version " + version)
			for _, flagName := range sortedKeys(ex.Flags) {
				cmdSB.WriteString(" --" + flagName + " " + ex.Flags[flagName])
			}
			cmdSB.WriteString("\n")

			includeCommandFileName := fmt.Sprintf("%s-%s-%s.sh", strings.ReplaceAll(cmd.CommandPath(), " ", "-"), version, source)
			const permissions = 0600
			err := os.WriteFile(fmt.Sprintf("%s/%s", includesLocalPath, includeCommandFileName), []byte(cmdSB.String()), permissions)
			if err != nil {
				return "", fmt.Errorf("failed to write file %s: %w", includeCommandFileName, err)
			}

			if tabs {
				sb.WriteString("      ")
			}
			sb.WriteString(fmt.Sprintf(".. literalinclude:: %s/%s\n", includesImportPath, includeCommandFileName))
			if tabs {
				sb.WriteString("      ")
			}
			sb.WriteString("   :language: shell\n")
		}
	}

	return sb.String(), nil
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

	var err error
	cmd.Example, err = buildExamples(cmd, cmdMetadata.Examples)

	return err
}

func removeCommandsWithOnlyPrivatePreview(cmd *cobra.Command) {
	operationID := cmd.Annotations["operationId"]
	if operationID == "" {
		return
	}

	cmdMetadata, ok := metadata[operationID]
	if ok && cmdMetadata.OnlyPrivatePreview {
		cmd.Parent().RemoveCommand(cmd)
	}
}

func applyTransformations(cmd *cobra.Command) error {
	setDisableAutoGenTag(cmd)
	removePluginCommands(cmd)
	addAdditionalLongText(cmd)

	if isAPICommand(cmd) {
		removeCommandsWithOnlyPrivatePreview(cmd)
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
