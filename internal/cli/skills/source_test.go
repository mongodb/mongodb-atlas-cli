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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/google/go-github/v61/github"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testCommitSHA = "0123456789abcdef0123456789abcdef01234567"
const githubBranchEndpoint = "/repos/mongodb/agent-skills/branches/main"

func TestGitHubSourceSnapshot(t *testing.T) {
	files := map[string]string{
		"/mongodb/agent-skills/" + testCommitSHA + "/skills/alpha/SKILL.md": `---
name: alpha
description: Alpha skill description.
---

# Alpha
`,
		"/mongodb/agent-skills/" + testCommitSHA + "/skills/beta/SKILL.md": `---
description: |
  Beta skill description
  continued on another line.
---

# Beta
`,
	}
	var gotRecursiveTree bool
	var rawRequests int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case githubBranchEndpoint:
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":   "main",
				"commit": map[string]any{"sha": testCommitSHA},
			})
		case "/repos/mongodb/agent-skills/git/trees/" + testCommitSHA:
			gotRecursiveTree = r.URL.Query().Get("recursive") == "1"
			_ = json.NewEncoder(w).Encode(map[string]any{
				"sha":       "tree-sha",
				"truncated": false,
				"tree": []map[string]string{
					{"path": "skills/OWNERS.yaml", "mode": "100644", "type": "blob"},
					{"path": "skills/alpha", "mode": "040000", "type": "tree"},
					{"path": "skills/alpha/SKILL.md", "mode": "100644", "type": "blob"},
					{"path": "skills/alpha/references/guide.md", "mode": "100644", "type": "blob"},
					{"path": "skills/alpha/scripts/run.py", "mode": "100755", "type": "blob"},
					{"path": "skills/beta/SKILL.md", "mode": "100644", "type": "blob"},
					{"path": "testing/ignored/SKILL.md", "mode": "100644", "type": "blob"},
				},
			})
		default:
			if content, ok := files[r.URL.Path]; ok {
				rawRequests++
				_, _ = w.Write([]byte(content))
				return
			}
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)

	source := githubSourceForTest(t, server)
	snapshot, err := source.Snapshot(context.Background())
	require.NoError(t, err)
	assert.Equal(t, testCommitSHA, snapshot.commitSHA)
	assert.True(t, gotRecursiveTree)
	assert.Zero(t, rawRequests, "snapshot enumeration must not read skill metadata")
	assert.Equal(t, []Skill{
		{Name: "alpha"},
		{Name: "beta"},
	}, snapshot.skills)

	alphaFiles, err := snapshot.filesForSkill("alpha")
	require.NoError(t, err)
	assert.Equal(t, []skillFile{
		{repositoryPath: "skills/alpha/SKILL.md", relativePath: "SKILL.md", mode: 0o644},
		{repositoryPath: "skills/alpha/references/guide.md", relativePath: "references/guide.md", mode: 0o644},
		{repositoryPath: "skills/alpha/scripts/run.py", relativePath: "scripts/run.py", mode: 0o755},
	}, alphaFiles)
}

func TestGitHubSourceRejectsTruncatedTree(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == githubBranchEndpoint {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":   "main",
				"commit": map[string]any{"sha": testCommitSHA},
			})
			return
		}
		if r.URL.Path == "/repos/mongodb/agent-skills/git/trees/"+testCommitSHA {
			_ = json.NewEncoder(w).Encode(map[string]any{"truncated": true, "tree": []map[string]string{}})
			return
		}
		http.NotFound(w, r)
	}))
	t.Cleanup(server.Close)

	_, err := githubSourceForTest(t, server).Snapshot(context.Background())
	require.ErrorContains(t, err, "truncated file tree")
}

func TestGitHubSourceEmptyCatalogIsEmptySlice(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == githubBranchEndpoint {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":   "main",
				"commit": map[string]any{"sha": testCommitSHA},
			})
			return
		}
		if r.URL.Path == "/repos/mongodb/agent-skills/git/trees/"+testCommitSHA {
			_ = json.NewEncoder(w).Encode(map[string]any{"truncated": false, "tree": []map[string]string{}})
			return
		}
		http.NotFound(w, r)
	}))
	t.Cleanup(server.Close)

	snapshot, err := githubSourceForTest(t, server).Snapshot(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, snapshot.skills)
	assert.Empty(t, snapshot.skills)
}

func TestGitHubSourceRejectsUnsafeSkillPaths(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == githubBranchEndpoint {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"name":   "main",
				"commit": map[string]any{"sha": testCommitSHA},
			})
			return
		}
		if r.URL.Path == "/repos/mongodb/agent-skills/git/trees/"+testCommitSHA {
			_ = json.NewEncoder(w).Encode(map[string]any{
				"truncated": false,
				"tree":      []map[string]string{{"path": "skills/alpha/../outside.md", "mode": "100644", "type": "blob"}},
			})
			return
		}
		http.NotFound(w, r)
	}))
	t.Cleanup(server.Close)

	_, err := githubSourceForTest(t, server).Snapshot(context.Background())
	require.ErrorContains(t, err, "unsafe file path")
}

func TestGitHubSourceRejectsControlCharactersInSkillNames(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{name: "newline", value: "bad\nname"},
		{name: "escape", value: "bad\x1bname"},
		{name: "bidi override", value: "bad\u202ename"},
		{name: "line separator", value: "bad\u2028name"},
		{name: "paragraph separator", value: "bad\u2029name"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.False(t, isSafePathComponent(test.value))
		})
	}
}

func TestIsSafeRelativePathRejectsWindowsDotSegments(t *testing.T) {
	for _, filePath := range []string{
		".. ", "../outside", "folder/.. ", "folder/.. /outside", `folder\..\outside`,
		"folder/NUL.txt", "folder/data:stream.md", "folder/CON",
	} {
		assert.False(t, isSafeRelativePath(filePath), "%q must not be a safe path", filePath)
	}
}

func TestSkillDescriptionIgnoresMarkdownLineThatStartsLikeDelimiter(t *testing.T) {
	content := []byte("---\ndescription: |\n  first line\n  --- example\n  last line\n---\n\n# Skill\n")
	description, err := skillDescription(content)
	require.NoError(t, err)
	assert.Equal(t, "first line\n--- example\nlast line", description)
}

func TestIsFrontmatterDelimiterRequiresWholeLine(t *testing.T) {
	assert.True(t, isFrontmatterDelimiter([]byte("---")))
	assert.True(t, isFrontmatterDelimiter([]byte("---\r")))
	assert.False(t, isFrontmatterDelimiter([]byte("--- example")))
	assert.False(t, isFrontmatterDelimiter([]byte("  ---")))
}

func TestSkillDescriptionSupportsCRLF(t *testing.T) {
	content := []byte("---\r\ndescription: Example skill.\r\n---\r\n\r\n# Skill\r\n")
	description, err := skillDescription(content)
	require.NoError(t, err)
	assert.Equal(t, "Example skill.", description)
}

func githubSourceForTest(t *testing.T, server *httptest.Server) *githubSource {
	t.Helper()
	apiURL, err := url.Parse(server.URL + "/")
	require.NoError(t, err)
	client := github.NewClient(server.Client())
	client.BaseURL = apiURL
	client.UploadURL = apiURL
	return &githubSource{apiClient: client, httpClient: server.Client(), rawBaseURL: server.URL}
}
