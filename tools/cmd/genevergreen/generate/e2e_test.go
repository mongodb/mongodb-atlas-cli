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

func TestE2EPackages(t *testing.T) {
	packages, err := e2ePackages(e2eWaterfallPkgPrefix)
	require.NoError(t, err)
	assert.NotEmpty(t, packages)
	for _, pkg := range packages {
		assert.True(t, strings.HasPrefix(pkg, e2eWaterfallPkgPrefix), pkg)
		assert.NotContains(t, pkg, "test/internal")
	}

	packages, err = e2ePackages(e2ePRPkgPrefix)
	require.NoError(t, err)
	assert.NotEmpty(t, packages)
	for _, pkg := range packages {
		assert.True(t, strings.HasPrefix(pkg, e2ePRPkgPrefix), pkg)
		assert.NotContains(t, pkg, "test/internal")
	}

	_, err = e2ePackages("test/internal/")
	assert.Error(t, err)
}

func TestE2EWaterfallTasks(t *testing.T) {
	testE2ETasks(t, E2EWaterfallTasks, e2eWaterfallVariant, e2eWaterfallTaskPrefix, e2eWaterfallPkgPrefix)
}

func TestE2EPRTasks(t *testing.T) {
	testE2ETasks(t, E2EPRTasks, e2ePRVariant, e2ePRTaskPrefix, e2ePRPkgPrefix)
}

func testE2ETasks(t *testing.T, generateTasks func(*shrub.Configuration) error, variantName, taskPrefix, pkgPrefix string) {
	t.Helper()

	c := &shrub.Configuration{}
	require.NoError(t, generateTasks(c))

	assert.Len(t, c.Variants, 1)
	v := c.Variants[0]
	assert.Equal(t, variantName, v.BuildName)
	assert.NotEmpty(t, v.DistroRunOn)
	assert.Len(t, v.TaskSpecs, len(c.Tasks))

	assert.NotEmpty(t, c.Tasks)
	for _, task := range c.Tasks {
		assert.True(t, strings.HasPrefix(task.Name, taskPrefix), task.Name)
		assert.Contains(t, task.Tags, "e2e")
		assert.NotNil(t, task.MustHaveResults)
		assert.True(t, *task.MustHaveResults)
		assert.Equal(t, e2eExecTimeoutSecs, task.ExecTimeoutSecs)
		assert.Len(t, task.Dependencies, 1)
		assert.Equal(t, "compile", task.Dependencies[0].Name)
		assert.Equal(t, "code_health", task.Dependencies[0].Variant)
		assert.NotNil(t, task.Dependencies[0].PatchOptional)
		assert.True(t, *task.Dependencies[0].PatchOptional)

		functions := make([]string, 0, len(task.Commands))
		e2eTestFound := false
		for _, command := range task.Commands {
			functions = append(functions, command.FunctionName)
			if command.FunctionName != "e2e test" {
				continue
			}
			e2eTestFound = true
			assert.True(t, strings.HasPrefix(command.Vars["E2E_TEST_PACKAGES"], "./"+pkgPrefix), command.Vars["E2E_TEST_PACKAGES"])
			assert.Equal(t, "${e2e_cloud_role_id}", command.Vars["E2E_CLOUD_ROLE_ID"])
			assert.Equal(t, "${e2e_test_bucket}", command.Vars["E2E_TEST_BUCKET"])
			assert.Equal(t, "3h", command.Vars["E2E_TIMEOUT"])
		}
		assert.True(t, e2eTestFound, task.Name)
		assert.Equal(t,
			[]string{"install mongodb database tools", "increase inotify limits", "authenticate to ecr", "restore go modules", "e2e test"},
			functions)
	}
}
