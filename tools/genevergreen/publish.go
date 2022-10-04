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

package main

import (
	"fmt"

	"github.com/evergreen-ci/shrub"
)

const (
	arm64  = "arm64"
	x86_64 = "x86_64"
	deb    = "deb"
	rpm    = "rpm"
)

var // if updating this list verify build/ci/repo_config.yaml matches.
distros = map[string]Platform{
	"amazon2": {
		extension:     rpm,
		architectures: []string{x86_64, arm64},
	},
	"rhel70": {
		extension:     rpm,
		architectures: []string{x86_64},
	},
	"rhel80": {
		extension:     rpm,
		architectures: []string{x86_64, arm64},
	},
	"rhel90": {
		extension:     rpm,
		architectures: []string{x86_64, arm64},
	},
	"debian10": {
		extension:     rpm,
		architectures: []string{x86_64},
	},
	"debian11": {
		extension:     deb,
		architectures: []string{x86_64},
	},
	"ubuntu18.04": {
		extension:     deb,
		architectures: []string{x86_64, arm64},
	},
	"ubuntu20.04": {
		extension:     deb,
		architectures: []string{x86_64, arm64},
	},
	"ubuntu22.04": {
		extension:     deb,
		architectures: []string{x86_64, arm64},
	},
}

type Platform struct {
	extension     string
	architectures []string
}

func generatePublishStableTasks(c *shrub.Configuration, toolName string) {
	dependency := []shrub.TaskDependency{
		{
			Name:    fmt.Sprintf("compile_%s", toolName),
			Variant: "code_health",
		},
		{
			Name:    fmt.Sprintf("release_%s", toolName),
			Variant: fmt.Sprintf("release_%s_github", toolName),
		},
	}
	for _, sv := range serverVersions {
		v := &shrub.Variant{
			BuildName:        fmt.Sprintf("generated_release_%s_publish_%s", toolName, sv),
			BuildDisplayName: fmt.Sprintf("Publish %s yum/apt %s", toolName, sv),
			DistroRunOn:      []string{"rhel80-small"},
		}
		publishVariant(
			c,
			v,
			toolName,
			sv,
			"_stable",
			dependency,
			true,
		)
	}
}

func generatePublishSnapshotTasks(c *shrub.Configuration, toolName string) {
	dependency := []shrub.TaskDependency{
		{
			Name:    fmt.Sprintf("compile_%s", toolName),
			Variant: "code_health",
		},
		{
			Name:    "package_goreleaser",
			Variant: fmt.Sprintf("goreleaser_%s_snapshot", toolName),
		},
	}
	publishVariant(
		c,
		c.Variant(fmt.Sprintf("publish_%s_snapshot", toolName)),
		toolName,
		"4.4",
		"",
		dependency,
		false,
	)
}

func publishVariant(c *shrub.Configuration, v *shrub.Variant, toolName, sv, stableSuffix string, dependency []shrub.TaskDependency, stable bool) {
	taskServerVersion := fmt.Sprintf("%s.0", sv)
	notaryKey := fmt.Sprintf("server-%s", sv)
	taskSv := "_" + sv
	if !stable {
		taskServerVersion = "4.4.0-rc3"
		notaryKey = "server-4.0"
		taskSv = ""
	}
	for k, d := range distros {
		for _, r := range repos {
			for _, a := range d.architectures {
				taskName := fmt.Sprintf("push_%s_%s_%s_%s%s%s", toolName, k, r, a, taskSv, stableSuffix)
				t := newPublishTask(taskName, toolName, d.extension, r, k, taskServerVersion, notaryKey, a, stable, dependency)
				c.Tasks = append(c.Tasks, t)
				v.AddTasks(t.Name)
			}
		}
	}
	c.Variants = append(c.Variants, v)
}

func newPublishTask(taskName, toolName, extension, edition, distro, taskServerVersion, notaryKey, arch string, stable bool, dependency []shrub.TaskDependency) *shrub.Task {
	t := &shrub.Task{
		Name: taskName,
	}
	t.Stepback(false).
		GitTagOnly(stable).
		Dependency(dependency...).
		Function("clone").
		Function("install curator").
		Patchable(false).
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
