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
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubCatalog struct {
	snapshot *skillSnapshot
	err      error
	files    map[string][]byte
	readErr  error
}

func (catalog stubCatalog) Snapshot(context.Context) (*skillSnapshot, error) {
	return catalog.snapshot, catalog.err
}

func (catalog stubCatalog) ReadFile(_ context.Context, _, path string) ([]byte, error) {
	if catalog.readErr != nil {
		return nil, catalog.readErr
	}
	content, ok := catalog.files[path]
	if !ok {
		return nil, fmt.Errorf("file %s not found in test source", path)
	}
	return content, nil
}

func TestListBuilderJSON(t *testing.T) {
	cmd := buildListCommand(stubCatalog{
		snapshot: &skillSnapshot{
			commitSHA: "commit",
			skills:    []Skill{{Name: "mongodb-schema-design"}},
			filesByName: map[string][]skillFile{
				"mongodb-schema-design": {{repositoryPath: "skills/mongodb-schema-design/SKILL.md", relativePath: "SKILL.md"}},
			},
		},
		files: map[string][]byte{
			"skills/mongodb-schema-design/SKILL.md": []byte("---\ndescription: Design schemas.\n---\n"),
		},
	})
	output, errOutput := new(bytes.Buffer), new(bytes.Buffer)
	cmd.SetOut(output)
	cmd.SetErr(errOutput)
	cmd.SetArgs([]string{"--output", jsonOutputFormat})

	require.NoError(t, cmd.Execute())
	assert.JSONEq(t, `[{"name":"mongodb-atlas","type":"plugin","description":"Install with atlas skills install --plugin mongodb-atlas --agent claude or atlas skills install --plugin mongodb-atlas --agent codex."},{"name":"mongodb-schema-design","type":"skill","description":"Design schemas."}]`, output.String())
}

func TestListBuilderPlainTextFlattensDescriptions(t *testing.T) {
	cmd := buildListCommand(stubCatalog{
		snapshot: &skillSnapshot{
			commitSHA: "commit",
			skills:    []Skill{{Name: "mongodb-mcp-setup"}},
			filesByName: map[string][]skillFile{
				"mongodb-mcp-setup": {{repositoryPath: "skills/mongodb-mcp-setup/SKILL.md", relativePath: "SKILL.md"}},
			},
		},
		files: map[string][]byte{
			"skills/mongodb-mcp-setup/SKILL.md": []byte("---\ndescription: |\n  Line one\n  Line two with a tab.\n---\n"),
		},
	})
	output := new(bytes.Buffer)
	cmd.SetOut(output)
	cmd.SetArgs([]string{"--output", "plaintext"})

	require.NoError(t, cmd.Execute())
	assert.Contains(t, output.String(), "mongodb-atlas")
	assert.Contains(t, output.String(), "plugin")
	assert.Contains(t, output.String(), "--plugin")
	assert.Contains(t, output.String(), "mongodb-mcp-setup")
	assert.Contains(t, output.String(), "Line one Line two with a tab.")
}

func TestListBuilderPlainTextWrapsDescriptionsUnderDescriptionColumn(t *testing.T) {
	const skillName = "mongodb-mcp-setup"
	description := strings.TrimSpace(strings.Repeat("word ", 20))
	output := renderSkillTable([]skillListEntry{{Name: skillName, Type: "skill", Description: description}}, defaultSkillsListWidth)
	lines := strings.Split(strings.TrimSuffix(output, "\n"), "\n")
	require.Len(t, lines, 3)
	assert.Equal(t, "NAME"+strings.Repeat(" ", 16)+"TYPE"+strings.Repeat(" ", 4)+"DESCRIPTION", lines[0])
	assert.Equal(t, skillName+"   skill   "+strings.TrimSpace(strings.Repeat("word ", 14)), lines[1])
	assert.Equal(t, strings.Repeat(" ", 28)+strings.TrimSpace(strings.Repeat("word ", 6)), lines[2])
}

func TestSingleLineRemovesControlAndFormatSeparators(t *testing.T) {
	assert.Equal(t, "Line one Line two", singleLine("Line one\nLine two\u202e\u2028"))
}

func TestListOptsRunReturnsSourceError(t *testing.T) {
	wantErr := errors.New("GitHub unavailable")
	opts := &skillListOpts{source: stubCatalog{err: wantErr}}

	assert.ErrorIs(t, opts.Run(context.Background()), wantErr)
}

func TestListOptsRunReturnsSkillMetadataError(t *testing.T) {
	wantErr := errors.New("unrelated skill metadata malformed")
	opts := &skillListOpts{source: stubCatalog{
		snapshot: &skillSnapshot{
			commitSHA: "commit",
			skills:    []Skill{{Name: "mongodb-schema-design"}},
			filesByName: map[string][]skillFile{
				"mongodb-schema-design": {{repositoryPath: "skills/mongodb-schema-design/SKILL.md", relativePath: "SKILL.md"}},
			},
		},
		readErr: wantErr,
	}}

	assert.ErrorIs(t, opts.Run(context.Background()), wantErr)
}

func TestListBuilderRejectsArguments(t *testing.T) {
	cmd := buildListCommand(stubCatalog{snapshot: &skillSnapshot{}})
	cmd.SilenceErrors = true
	cmd.SilenceUsage = true
	cmd.SetArgs([]string{"not-a-flag"})

	require.Error(t, cmd.Execute())
}
