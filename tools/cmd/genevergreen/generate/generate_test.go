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

package generate

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/evergreen-ci/shrub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	push = "push"
)

// loadMatrix loads the committed build/ci/server-matrix.json.
func loadMatrix(t *testing.T) ServerMatrix {
	t.Helper()
	root, err := e2eRepoRoot()
	require.NoError(t, err)
	m, err := LoadServerMatrix(filepath.Join(root, "build", "ci", "server-matrix.json"))
	require.NoError(t, err)
	return m
}

func taskNames(c *shrub.Configuration) []string {
	names := make([]string, 0, len(c.Tasks))
	for _, t := range c.Tasks {
		names = append(names, t.Name)
	}
	return names
}

func variantTaskNames(v *shrub.Variant) []string {
	names := make([]string, 0, len(v.TaskSpecs))
	for _, s := range v.TaskSpecs {
		names = append(names, s.Name)
	}
	return names
}

// commandVars returns the vars of the task's only call to fn.
func commandVars(t *testing.T, task *shrub.Task, fn string) map[string]string {
	t.Helper()
	var vars []map[string]string
	for _, c := range task.Commands {
		if c.FunctionName == fn {
			vars = append(vars, c.Vars)
		}
	}
	if !assert.Len(t, vars, 1, "%s should call %q once", task.Name, fn) {
		return nil
	}
	return vars[0]
}

// expectedPublishTasks lists one push task per repo, supported distro and architecture.
func expectedPublishTasks(m ServerMatrix, sv, suffix string) []string {
	var names []string
	for _, r := range repos {
		for d, p := range distros {
			if !m.Supports(sv, d) {
				continue
			}
			for _, a := range p.architectures {
				names = append(names, "push_atlascli_"+d+"_"+r+"_"+a+suffix)
			}
		}
	}
	return names
}

func TestGeneratorInputs(t *testing.T) {
	for _, os := range oses {
		assert.Contains(t, distros, newOs[os], "os %s has no publish distro", os)
		assert.Contains(t, postPkgImg, os)
	}
}

func TestPublishSnapshotTasks(t *testing.T) {
	m := loadMatrix(t)
	c := &shrub.Configuration{}
	require.NoError(t, PublishSnapshotTasks(c, m))

	expected := expectedPublishTasks(m, snapshotServerVersion, "")
	assert.ElementsMatch(t, expected, taskNames(c))
	assert.NotEmpty(t, c.Variants)
	for _, v := range c.Variants {
		assert.Equal(t, "publish_atlascli_snapshot", v.BuildName)
		assert.ElementsMatch(t, expected, variantTaskNames(v))
	}

	commandFound := false
	for _, task := range c.Tasks {
		for _, c := range task.Commands {
			if c.FunctionName == push {
				commandFound = true
			}
		}
	}
	assert.True(t, commandFound, "expected to find a push command")
}

func TestPublishStableTasks(t *testing.T) {
	m := loadMatrix(t)
	c := &shrub.Configuration{}
	PublishStableTasks(c, m)

	expected := make([]string, 0, len(c.Tasks))
	require.Len(t, c.Variants, len(m))
	for i, sv := range m.Versions() {
		svTasks := expectedPublishTasks(m, sv, "_"+strings.ReplaceAll(sv, ".", "")+"_stable")
		expected = append(expected, svTasks...)
		assert.Equal(t, buildNamePrefix+strings.ReplaceAll(sv, ".", ""), c.Variants[i].BuildName)
		assert.ElementsMatch(t, svTasks, variantTaskNames(c.Variants[i]))
	}
	assert.ElementsMatch(t, expected, taskNames(c))

	commandFound := false
	for _, task := range c.Tasks {
		for _, c := range task.Commands {
			if c.FunctionName != push {
				continue
			}
			commandFound = true
			assert.NotEmpty(t, c.Vars["pgp_server_version"])
		}
	}
	assert.True(t, commandFound, "expected to find a push command")
}

func TestPostPkgMetaTasks(t *testing.T) {
	m := loadMatrix(t)
	c := &shrub.Configuration{}
	PostPkgMetaTasks(c, m)

	var expected []string
	for _, os := range oses {
		for _, sv := range m.Versions() {
			if m.Supports(sv, newOs[os]) {
				expected = append(expected, "pkg_test_atlascli_meta_docker_"+sv+"_"+os)
			}
		}
	}
	assert.ElementsMatch(t, expected, taskNames(c))
	require.Len(t, c.Variants, 1)
	assert.ElementsMatch(t, expected, variantTaskNames(c.Variants[0]))

	for _, task := range c.Tasks {
		assert.NotEmpty(t, commandVars(t, task, "docker build meta")["pgp_server_version"], task.Name)
	}
}

func TestRepoTasks(t *testing.T) {
	m := loadMatrix(t)
	c := &shrub.Configuration{}
	RepoTasks(c, m)

	expected := make([]string, 0, len(c.Tasks))
	expectedDeps := map[string]shrub.TaskDependency{}
	require.Len(t, c.Variants, len(m))
	for i, sv := range m.Versions() {
		svNoDot := strings.ReplaceAll(sv, ".", "")
		var svTasks []string
		for _, os := range oses {
			if !m.Supports(sv, newOs[os]) {
				continue
			}
			for _, r := range repos {
				name := "test_repo_atlascli_" + os + "_" + r + "_" + sv
				svTasks = append(svTasks, name)
				expectedDeps[name] = shrub.TaskDependency{
					Name:    "push_atlascli_" + newOs[os] + "_" + r + "_" + x86_64 + "_" + svNoDot + "_stable",
					Variant: buildNamePrefix + svNoDot,
				}
			}
		}
		expected = append(expected, svTasks...)
		assert.Equal(t, "test_repo_atlascli_"+sv, c.Variants[i].BuildName)
		assert.ElementsMatch(t, svTasks, variantTaskNames(c.Variants[i]))
	}
	assert.ElementsMatch(t, expected, taskNames(c))

	stable := &shrub.Configuration{}
	PublishStableTasks(stable, m)
	stableByVariant := map[string][]string{}
	for _, v := range stable.Variants {
		stableByVariant[v.BuildName] = variantTaskNames(v)
	}

	for _, task := range c.Tasks {
		assert.Equal(t, []shrub.TaskDependency{expectedDeps[task.Name]}, task.Dependencies, task.Name)
		for _, d := range task.Dependencies {
			assert.Contains(t, stableByVariant[d.Variant], d.Name, "%s depends on a task missing from %s", task.Name, d.Variant)
		}
		assert.NotEmpty(t, commandVars(t, task, "docker build repo")["pgp_server_version"], task.Name)
	}
}

func TestGetGpgServerVersion(t *testing.T) {
	assert.Equal(t, "8.0", getGpgServerVersion("8.3"))
	assert.Equal(t, "8.0", getGpgServerVersion("8.2"))
	assert.Equal(t, "8.0", getGpgServerVersion("8.0"))
	assert.Equal(t, "7.0", getGpgServerVersion("7.0"))
	assert.Equal(t, "6.0", getGpgServerVersion("6.0"))
	assert.Equal(t, "9", getGpgServerVersion("9.0"))
	assert.Equal(t, "9", getGpgServerVersion("9.1"))
}
