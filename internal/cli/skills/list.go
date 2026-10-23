// Copyright 2026 MongoDB Inc
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

package skills

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const listTemplate = `NAME	TYPE	DESCRIPTION{{range valueOrEmptySlice .}}
{{.Name}}	{{.Type}}	{{.Description}}{{end}}
`

const (
	defaultSkillsListWidth     = 100
	skillsListColumnGap        = 3
	atlasPluginListDescription = "Install with atlas skills install --plugin mongodb-atlas --agent claude or atlas skills install --plugin mongodb-atlas --agent codex."
)

// Skill describes an installable MongoDB agent skill.
type Skill struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type skillListEntry struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description,omitempty"`
}

type skillListOpts struct {
	cli.OutputOpts
	source skillSource
}

func (opts *skillListOpts) Run(ctx context.Context) error {
	snapshot, err := opts.source.Snapshot(ctx)
	if err != nil {
		return err
	}
	entries := make([]skillListEntry, 0, len(snapshot.skills)+1)
	entries = append(entries, skillListEntry{
		Name:        atlasPluginName,
		Type:        "plugin",
		Description: atlasPluginListDescription,
	})
	for _, available := range snapshot.skills {
		files, err := snapshot.filesForSkill(available.Name)
		if err != nil {
			return err
		}
		skillFileIndex := -1
		for j, file := range files {
			if file.relativePath == skillDefinitionFile {
				skillFileIndex = j
				break
			}
		}
		if skillFileIndex < 0 {
			return fmt.Errorf("upstream skill %q does not contain SKILL.md", available.Name)
		}
		content, err := opts.source.ReadFile(ctx, snapshot.commitSHA, files[skillFileIndex].repositoryPath)
		if err != nil {
			return fmt.Errorf("could not read %s: %w", files[skillFileIndex].repositoryPath, err)
		}
		description, err := skillDescription(content)
		if err != nil {
			return fmt.Errorf("could not read metadata from %s: %w", files[skillFileIndex].repositoryPath, err)
		}
		entries = append(entries, skillListEntry{
			Name:        available.Name,
			Type:        "skill",
			Description: description,
		})
	}

	if !opts.IsPlainOutput() {
		return opts.Print(entries)
	}

	for i := range entries {
		entries[i].Description = singleLine(entries[i].Description)
	}

	writer := opts.ConfigWriter()
	_, err = io.WriteString(writer, renderSkillTable(entries, skillsListOutputWidth(writer)))
	return err
}

func skillsListOutputWidth(writer io.Writer) int {
	if file, ok := writer.(*os.File); ok {
		fd := file.Fd()
		maxInt := (uint64(1) << (strconv.IntSize - 1)) - 1
		if uint64(fd) > maxInt {
			return defaultSkillsListWidth
		}
		width, _, err := term.GetSize(int(fd))
		if err == nil && width > 0 {
			return width
		}
	}
	return defaultSkillsListWidth
}

func renderSkillTable(skills []skillListEntry, width int) string {
	if width <= 0 {
		width = defaultSkillsListWidth
	}

	nameWidth := utf8.RuneCountInString("NAME")
	typeWidth := utf8.RuneCountInString("TYPE")
	for _, skill := range skills {
		if skillWidth := utf8.RuneCountInString(skill.Name); skillWidth > nameWidth {
			nameWidth = skillWidth
		}
		if skillTypeWidth := utf8.RuneCountInString(skill.Type); skillTypeWidth > typeWidth {
			typeWidth = skillTypeWidth
		}
	}

	descriptionWidth := width - nameWidth - typeWidth - skillsListColumnGap - skillsListColumnGap
	if descriptionWidth < 1 {
		descriptionWidth = 1
	}
	columnGap := strings.Repeat(" ", skillsListColumnGap)
	continuationIndent := strings.Repeat(" ", nameWidth+typeWidth+skillsListColumnGap+skillsListColumnGap)

	var output strings.Builder
	output.WriteString(padRight("NAME", nameWidth))
	output.WriteString(columnGap)
	output.WriteString(padRight("TYPE", typeWidth))
	output.WriteString(columnGap)
	output.WriteString("DESCRIPTION\n")
	for _, skill := range skills {
		lines := wrapDescription(skill.Description, descriptionWidth)
		output.WriteString(padRight(skill.Name, nameWidth))
		output.WriteString(columnGap)
		output.WriteString(padRight(skill.Type, typeWidth))
		output.WriteString(columnGap)
		output.WriteString(lines[0])
		output.WriteByte('\n')
		for _, line := range lines[1:] {
			output.WriteString(continuationIndent)
			output.WriteString(line)
			output.WriteByte('\n')
		}
	}
	return output.String()
}

func padRight(value string, width int) string {
	if padding := width - utf8.RuneCountInString(value); padding > 0 {
		return value + strings.Repeat(" ", padding)
	}
	return value
}

func wrapDescription(description string, width int) []string {
	if width < 1 {
		width = 1
	}

	words := strings.Fields(description)
	if len(words) == 0 {
		return []string{""}
	}

	lines := make([]string, 0, len(words))
	line := ""
	for _, word := range words {
		wordRunes := []rune(word)
		if len(wordRunes) > width {
			if line != "" {
				lines = append(lines, line)
				line = ""
			}
			for len(wordRunes) > width {
				lines = append(lines, string(wordRunes[:width]))
				wordRunes = wordRunes[width:]
			}
			word = string(wordRunes)
		}

		wordWidth := utf8.RuneCountInString(word)
		if line == "" {
			line = word
			continue
		}
		if utf8.RuneCountInString(line)+1+wordWidth > width {
			lines = append(lines, line)
			line = word
		} else {
			line += " " + word
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}

func singleLine(value string) string {
	value = strings.Map(func(char rune) rune {
		if unicode.IsControl(char) || unicode.Is(unicode.Cf, char) || unicode.Is(unicode.Zl, char) || unicode.Is(unicode.Zp, char) {
			return ' '
		}
		return char
	}, value)
	return strings.Join(strings.Fields(value), " ")
}

// ListBuilder builds the skills list command.
func ListBuilder() *cobra.Command {
	return buildListCommand(newGitHubSource())
}

func buildListCommand(source skillSource) *cobra.Command {
	opts := &skillListOpts{source: source}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List MongoDB agent skills and plugins available upstream.",
		Long:  "List the standalone MongoDB agent skills and bundled mongodb-atlas plugin available for Claude Code or Codex.",
		Example: `  # List available MongoDB agent skills and plugins:
  atlas skills list

  # List skills as JSON:
  atlas skills list --output json`,
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.InitOutput(cmd.OutOrStdout(), listTemplate)()
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.Run(cmd.Context())
		},
	}
	hideAtlasProfileFlag(cmd)

	opts.AddOutputOptFlags(cmd)
	return cmd
}
