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
	"fmt"
	"strings"

	"github.com/evergreen-ci/shrub"
)

type Platform struct {
	extension     string
	architectures []string
}

const (
	arm64   = "arm64"
	aarch64 = "aarch64"
	x86_64  = "x86_64"
	deb     = "deb"
	rpm     = "rpm"
)

// distros - if updating this list verify build/ci/repo_config.yaml matches.
var distros = map[string]Platform{
	"amazon2": {
		extension:     rpm,
		architectures: []string{x86_64, aarch64},
	},
	"rhel70": {
		extension:     rpm,
		architectures: []string{x86_64},
	},
	"rhel80": {
		extension:     rpm,
		architectures: []string{x86_64, aarch64},
	},
	"rhel90": {
		extension:     rpm,
		architectures: []string{x86_64, aarch64},
	},
	"debian10": {
		extension:     deb,
		architectures: []string{x86_64},
	},
	"debian11": {
		extension:     deb,
		architectures: []string{x86_64},
	},
	"ubuntu1804": {
		extension:     deb,
		architectures: []string{x86_64, arm64},
	},
	"ubuntu2004": {
		extension:     deb,
		architectures: []string{x86_64, arm64},
	},
	"ubuntu2204": {
		extension:     deb,
		architectures: []string{x86_64, arm64},
	},
}

func newPublishTask(taskName, toolName, extension, edition, distro, taskServerVersion, notaryKey, arch string, stable bool, dependency []shrub.TaskDependency) *shrub.Task {
	t := &shrub.Task{
		Name: taskName,
	}
	t.Stepback(false).
		Patchable(true). // don't publish patches
		GitTagOnly(stable).
		Dependency(dependency...).
		Function("clone").
		Function("install curator").
		FunctionWithVars("push", map[string]string{
			"tool_name":       toolName,
			"distro":          distro,
			"ext":             extension,
			"server_version":  taskServerVersion,
			"notary_key_name": notaryKey,
			"arch":            arch,
			"edition":         edition,
		})
	return t
}

func publishVariant(c *shrub.Configuration, v *shrub.Variant, toolName, sv, stableSuffix string, dependency []shrub.TaskDependency, stable bool) {
	taskServerVersion := fmt.Sprintf("%s.0", sv)
	notaryKey := fmt.Sprintf("server-%s", sv)
	taskSv := "_" + sv
	if !stable {
		taskServerVersion += "-rc1"
		taskSv = ""
	}
	for _, r := range repos {
		for k, d := range distros {
			for _, a := range d.architectures {
				taskName := fmt.Sprintf("push_%s_%s_%s_%s%s%s", toolName, k, r, a, strings.ReplaceAll(taskSv, ".", ""), stableSuffix)
				t := newPublishTask(taskName, toolName, d.extension, r, k, taskServerVersion, notaryKey, a, stable, dependency)
				c.Tasks = append(c.Tasks, t)
				v.AddTasks(t.Name)
			}
		}
	}
	c.Variants = append(c.Variants, v)
}
