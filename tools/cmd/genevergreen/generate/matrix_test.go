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

package generate

import (
	"cmp"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"slices"
	"testing"

	"github.com/evergreen-ci/shrub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var update = flag.Bool("update", false, "rewrite golden files")

func TestParseServerMatrix(t *testing.T) {
	m, err := ParseServerMatrix([]byte(`{"9.0": ["rhel80"], "10.0": ["rhel80"], "8.2": ["debian12", "rhel80"]}`))
	require.NoError(t, err)
	assert.Equal(t, []string{"8.2", "9.0", "10.0"}, m.Versions())
	assert.True(t, m.Supports("8.2", "debian12"))
	assert.False(t, m.Supports("9.0", "debian12"))
	assert.False(t, m.Supports("7.0", "rhel80"))
}

func TestParseServerMatrixRejects(t *testing.T) {
	tests := map[string]string{
		"empty":            `{}`,
		"not an object":    `["8.0"]`,
		"patch version":    `{"8.0.1": ["rhel80"]}`,
		"major only":       `{"8": ["rhel80"]}`,
		"no distros":       `{"8.0": []}`,
		"unknown distro":   `{"8.0": ["rhel70"]}`,
		"duplicate distro": `{"8.0": ["rhel80", "rhel80"]}`,
		"trailing data":    `{"8.0": ["rhel80"]} {}`,
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := ParseServerMatrix([]byte(input))
			assert.Error(t, err)
		})
	}
}

func TestPublishSnapshotTasksRequiresSnapshotVersion(t *testing.T) {
	err := PublishSnapshotTasks(&shrub.Configuration{}, ServerMatrix{"99.0": {"rhel80"}})
	assert.Error(t, err)
}

func TestPublishDistros(t *testing.T) {
	d := PublishDistros()
	assert.Len(t, d, len(distros))
	assert.True(t, slices.IsSorted(d))
}

// TestGolden checks generated tasks for a fixture matrix with a mock 99.0. Run with -update to rewrite.
func TestGolden(t *testing.T) {
	m, err := LoadServerMatrix(filepath.Join("testdata", "server-matrix.json"))
	require.NoError(t, err)

	tests := map[string]func(*testing.T, *shrub.Configuration){
		"repo": func(_ *testing.T, c *shrub.Configuration) {
			RepoTasks(c, m)
		},
		"postpkg-meta": func(_ *testing.T, c *shrub.Configuration) {
			PostPkgMetaTasks(c, m)
		},
		"snapshot": func(t *testing.T, c *shrub.Configuration) {
			t.Helper()
			require.NoError(t, PublishSnapshotTasks(c, m))
		},
		"publish": func(_ *testing.T, c *shrub.Configuration) {
			PublishStableTasks(c, m)
		},
	}
	for name, gen := range tests {
		t.Run(name, func(t *testing.T) {
			c := &shrub.Configuration{}
			gen(t, c)
			got := normalizedJSON(t, c)

			path := filepath.Join("testdata", "golden", name+".json")
			if *update {
				require.NoError(t, os.WriteFile(path, got, 0o600))
			}
			want, err := os.ReadFile(path)
			require.NoError(t, err)
			assert.Equal(t, string(want), string(got))
		})
	}
}

// normalizedJSON sorts tasks by name, since publish tasks follow map order.
func normalizedJSON(t *testing.T, c *shrub.Configuration) []byte {
	t.Helper()
	byName := func(a, b *shrub.Task) int { return cmp.Compare(a.Name, b.Name) }
	slices.SortFunc(c.Tasks, byName)
	for _, v := range c.Variants {
		slices.SortFunc(v.TaskSpecs, func(a, b shrub.TaskSpec) int { return cmp.Compare(a.Name, b.Name) })
	}
	b, err := json.MarshalIndent(c, "", "\t")
	require.NoError(t, err)
	return append(b, '\n')
}
