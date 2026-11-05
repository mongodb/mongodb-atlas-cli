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
	"strings"
	"testing"

	"github.com/evergreen-ci/shrub"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadE2EPackages(t *testing.T) {
	packages, err := loade2ePackages()
	require.NoError(t, err)
	assert.NotEmpty(t, packages)

	var flags, pr, gov int
	for _, pkg := range packages {
		assert.True(t, strings.HasPrefix(pkg.path, "test/e2e/"), pkg.path)
		if pkg.path == "test/e2e/atlas/clusters/flags/clustersflags" {
			flags++
			assert.True(t, pkg.hasGroup("gov"))
		}
		if pkg.hasGroup("pr") {
			pr++
		}
		if pkg.hasGroup("gov") {
			gov++
		}
	}
	assert.Equal(t, 1, flags)
	assert.Positive(t, pr)
	assert.Positive(t, gov)
}

func TestValidateE2EManifest(t *testing.T) {
	valid := e2eManifest{
		Groups:   []string{"pr", "gov", "noapi"},
		Fixtures: []string{"docker", "aws_role", "flex_mtm"},
		Needs:    []string{"test/e2e/atlas/clusters/flags/clustersflags"},
		Timeout:  "3h",
	}
	require.NoError(t, validateE2EManifest(valid))

	require.Error(t, validateE2EManifest(e2eManifest{Groups: []string{"wat"}}))
	require.Error(t, validateE2EManifest(e2eManifest{Fixtures: []string{"wat"}}))
	require.Error(t, validateE2EManifest(e2eManifest{Fixtures: []string{"cluster"}}))
	require.Error(t, validateE2EManifest(e2eManifest{Needs: []string{"wat"}}))
	require.Error(t, validateE2EManifest(e2eManifest{Timeout: "forever"}))
}

func TestValidateE2ENeeds(t *testing.T) {
	a := e2ePackage{path: "test/e2e/a", manifest: e2eManifest{Needs: []string{"test/e2e/b"}}}
	b := e2ePackage{path: "test/e2e/b"}
	require.NoError(t, validateE2ENeeds(map[string]e2ePackage{"test/e2e/a": a, "test/e2e/b": b}, "v"))

	missing := e2ePackage{path: "test/e2e/a", manifest: e2eManifest{Needs: []string{"test/e2e/gone"}}}
	require.ErrorContains(t, validateE2ENeeds(map[string]e2ePackage{"test/e2e/a": missing}, "v"), "which is not generated")

	self := e2ePackage{path: "test/e2e/a", manifest: e2eManifest{Needs: []string{"test/e2e/a"}}}
	require.ErrorContains(t, validateE2ENeeds(map[string]e2ePackage{"test/e2e/a": self}, "v"), "cannot need itself")

	c1 := e2ePackage{path: "test/e2e/a", manifest: e2eManifest{Needs: []string{"test/e2e/b"}}}
	c2 := e2ePackage{path: "test/e2e/b", manifest: e2eManifest{Needs: []string{"test/e2e/a"}}}
	require.ErrorContains(t, validateE2ENeeds(map[string]e2ePackage{"test/e2e/a": c1, "test/e2e/b": c2}, "v"), "circular")
}

func TestE2EWaterfallTasks(t *testing.T) {
	testE2ETasks(t, E2EOptions{
		Variant:       e2eWaterfallVariant,
		TaskPrefix:    e2eWaterfallTaskPrefix,
		ExcludeGroups: []string{"pr"},
	}, func(pkg e2ePackage) bool {
		return !pkg.hasGroup("pr")
	}, "")
}

func TestE2EPRTasks(t *testing.T) {
	testE2ETasks(t, E2EOptions{
		Variant:          e2ePRVariant,
		TaskPrefix:       e2ePRTaskPrefix,
		IncludeAnyGroups: []string{"pr"},
	}, func(pkg e2ePackage) bool {
		return pkg.hasGroup("pr")
	}, "")
}

func TestE2EGovTasks(t *testing.T) {
	testE2ETasks(t, E2EOptions{
		Variant:          e2eGovVariant,
		TaskPrefix:       e2eGovTaskPrefix,
		Profile:          e2eGovProfile,
		IncludeAnyGroups: []string{"gov"},
	}, func(pkg e2ePackage) bool {
		return pkg.hasGroup("gov")
	}, e2eGovProfile)
}

func testE2ETasks(t *testing.T, opts E2EOptions, selected func(e2ePackage) bool, profile string) {
	t.Helper()

	packages, err := loade2ePackages()
	require.NoError(t, err)

	expected := make(map[string]e2ePackage)
	nameByPath := make(map[string]string, len(packages))
	for _, pkg := range packages {
		if !selected(pkg) {
			continue
		}
		name, err := e2eTaskName(pkg, opts.TaskPrefix)
		require.NoError(t, err)
		expected[name] = pkg
		nameByPath[pkg.path] = name
	}
	require.NotEmpty(t, expected)

	c := &shrub.Configuration{}
	require.NoError(t, E2ETasks(c, opts))

	variantName := opts.Variant

	assert.Len(t, c.Variants, 1)
	v := c.Variants[0]
	assert.Equal(t, variantName, v.BuildName)
	assert.Empty(t, v.BuildDisplayName)
	assert.Empty(t, v.DistroRunOn)
	assert.Empty(t, v.Expansions)
	assert.Len(t, v.TaskSpecs, len(c.Tasks))

	generated := make(map[string]*shrub.Task, len(c.Tasks))
	for _, task := range c.Tasks {
		generated[task.Name] = task
	}
	assert.Len(t, generated, len(expected))

	for name, pkg := range expected {
		task, ok := generated[name]
		if !assert.True(t, ok, name) {
			continue
		}

		assert.Contains(t, task.Tags, "e2e", name)
		assert.NotNil(t, task.MustHaveResults, name)
		assert.True(t, *task.MustHaveResults, name)

		expectedDeps := map[string]bool{"compile": false}
		for _, needed := range pkg.manifest.Needs {
			expectedDeps[nameByPath[needed]] = false
		}
		require.Len(t, task.Dependencies, len(expectedDeps), name)
		for _, dep := range task.Dependencies {
			if dep.Name == "compile" {
				require.Equal(t, "code_health", dep.Variant, name)
			} else {
				require.Equal(t, variantName, dep.Variant, name)
			}
			require.NotNil(t, dep.PatchOptional, name)
			require.True(t, *dep.PatchOptional, name)
			_, ok := expectedDeps[dep.Name]
			require.True(t, ok, "%s: unexpected dependency %s", name, dep.Name)
			expectedDeps[dep.Name] = true
		}
		for depName, found := range expectedDeps {
			require.True(t, found, "%s: missing dependency %s", name, depName)
		}

		var functions []string
		var e2eTest *shrub.CommandDefinition
		for _, command := range task.Commands {
			functions = append(functions, command.FunctionName)
			if command.FunctionName == "e2e test" {
				e2eTest = command
			}
		}
		require.NotNil(t, e2eTest, name)
		dockerIdx := -1
		for i, fn := range functions {
			if fn == "install mongodb database tools" {
				dockerIdx = i
			}
		}
		assert.Equal(t, pkg.hasFixture("docker"), dockerIdx != -1, name)
		if dockerIdx != -1 {
			assert.Equal(t, e2eDockerFunctions, functions[dockerIdx:dockerIdx+len(e2eDockerFunctions)], name)
		}
		assert.Equal(t, "restore go modules", functions[len(functions)-2], name)
		assert.Equal(t, "e2e test", functions[len(functions)-1], name)

		vars := e2eTest.Vars
		assert.Equal(t, "./"+pkg.path+"/...", vars["E2E_TEST_PACKAGES"], name)
		if profile != "" {
			assert.Equal(t, profile, vars["E2E_PROFILE_NAME"], name)
		} else {
			assert.NotContains(t, vars, "E2E_PROFILE_NAME", name)
		}
		assert.Equal(t, pkg.manifest.Timeout != "", vars["E2E_TIMEOUT"] != "", name)
		if pkg.manifest.Timeout != "" {
			assert.Equal(t, pkg.manifest.Timeout, vars["E2E_TIMEOUT"], name)
			assert.Equal(t, e2eExecTimeoutSecs, task.ExecTimeoutSecs, name)
		} else {
			assert.Zero(t, task.ExecTimeoutSecs, name)
		}
		for _, fixture := range pkg.manifest.Fixtures {
			for key, value := range e2eFixtureVars[fixture] {
				assert.Equal(t, value, vars[key], "%s (%s)", name, fixture)
			}
		}
	}
}

func TestE2EExclusions(t *testing.T) {
	packages, err := loade2ePackages()
	require.NoError(t, err)

	prCount := 0
	federationCount := 0
	for _, pkg := range packages {
		if pkg.hasGroup("pr") {
			prCount++
			if pkg.hasFixture("federation_settings") {
				federationCount++
			}
		}
	}
	require.Positive(t, federationCount)

	c := &shrub.Configuration{}
	require.NoError(t, E2ETasks(c, E2EOptions{
		Variant:          e2ePRVariant,
		TaskPrefix:       e2ePRTaskPrefix,
		IncludeAnyGroups: []string{"pr"},
		ExcludeFixtures:  []string{"federation_settings"},
	}))
	assert.Len(t, c.Tasks, prCount-federationCount)
	for _, task := range c.Tasks {
		assert.NotContains(t, task.Name, "federationsettings", task.Name)
	}

	c = &shrub.Configuration{}
	require.NoError(t, E2ETasks(c, E2EOptions{
		Variant:         e2eWaterfallVariant,
		TaskPrefix:      e2eWaterfallTaskPrefix,
		ExcludeGroups:   []string{"pr"},
		ExcludeFixtures: []string{"docker"},
	}))
	for _, task := range c.Tasks {
		assert.NotContains(t, task.Name, "deployments_local", task.Name)
	}

	err = E2ETasks(&shrub.Configuration{}, E2EOptions{
		Variant:         e2eWaterfallVariant,
		TaskPrefix:      e2eWaterfallTaskPrefix,
		ExcludeGroups:   []string{"pr"},
		ExcludeFixtures: []string{"wat"},
	})
	require.ErrorContains(t, err, "unknown e2e fixture to exclude: wat")

	err = E2ETasks(&shrub.Configuration{}, E2EOptions{
		Variant:          e2ePRVariant,
		TaskPrefix:       e2ePRTaskPrefix,
		IncludeAnyGroups: []string{"wat"},
	})
	require.ErrorContains(t, err, "unknown e2e group: wat")

	err = E2ETasks(&shrub.Configuration{}, E2EOptions{
		Variant:          e2ePRVariant,
		TaskPrefix:       e2ePRTaskPrefix,
		IncludeAnyGroups: []string{"pr"},
		ExcludeGroups:    []string{"pr"},
	})
	assert.ErrorContains(t, err, "no e2e packages selected")
}
