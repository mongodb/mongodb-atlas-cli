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

const (
	runOn    = "ubuntu2004-small"
	atlascli = "atlascli"
	mongocli = "mongocli"
)

var (
	serverVersions = []string{
		"4.4",
		"5.0",
		"6.0",
		"7.0",
	}
	oses = []string{
		"amazonlinux2",
		"centos7",
		"centos8",
		"rhel9",
		"debian10",
		"debian11",
		"debian12",
		"ubuntu20.04",
		"ubuntu22.04",
	}
	repos      = []string{"org", "enterprise"}
	postPkgImg = map[string]string{
		"centos7":      "centos7-rpm",
		"centos8":      "centos8-rpm",
		"rhel9":        "rhel9-rpm",
		"amazonlinux2": "amazonlinux2-rpm",
		"ubuntu20.04":  "ubuntu20.04-deb",
		"ubuntu22.04":  "ubuntu22.04-deb",
		"debian10":     "debian10-deb",
		"debian11":     "debian11-deb",
		"debian12":     "debian12-deb",
	}
	newOs = map[string]string{
		"centos7":      "rhel70",
		"centos8":      "rhel80",
		"rhel9":        "rhel90",
		"amazonlinux2": "amazon2",
		"ubuntu20.04":  "ubuntu2004",
		"ubuntu22.04":  "ubuntu2204",
		"debian10":     "debian10",
		"debian11":     "debian11",
		"debian12":     "debian12",
	}
)

func newDependency(os, serverVersion, repo string) shrub.TaskDependency {
	return shrub.TaskDependency{
		Name:    fmt.Sprintf("push_mongocli_%s_%s_%s_%s_stable", newOs[os], repo, x86_64, strings.ReplaceAll(serverVersion, ".", "")),
		Variant: fmt.Sprintf("generated_release_mongocli_publish_%s", strings.ReplaceAll(serverVersion, ".", "")),
	}
}

func RepoTasks(c *shrub.Configuration) {
	for _, serverVersion := range serverVersions {
		v := &shrub.Variant{
			BuildName:        fmt.Sprintf("test_repo_mongocli_%v", serverVersion),
			BuildDisplayName: fmt.Sprintf("Test mongocli on repo %v", serverVersion),
			DistroRunOn:      []string{runOn},
		}

		for _, os := range oses {
			for _, repo := range repos {
				mongoRepo := "https://repo.mongodb.com"
				if repo == "org" {
					mongoRepo = "https://repo.mongodb.org"
				}

				t := &shrub.Task{
					Name: fmt.Sprintf("test_repo_mongocli_%v_%v_%v", os, repo, serverVersion),
				}
				t = t.Stepback(false).
					GitTagOnly(true).
					Dependency(newDependency(os, serverVersion, repo)).
					Function("clone").
					FunctionWithVars("docker build repo", map[string]string{
						"server_version": serverVersion,
						"package":        "mongocli",
						"entrypoint":     "mongocli",
						"image":          os,
						"mongo_package":  fmt.Sprintf("mongodb-%v", repo),
						"mongo_repo":     mongoRepo,
					})
				c.Tasks = append(c.Tasks, t)
				v.AddTasks(t.Name)
			}
		}

		c.Variants = append(c.Variants, v)
	}
}

func PostPkgTasks(c *shrub.Configuration) {
	v := &shrub.Variant{
		BuildName:        "pkg_smoke_tests_docker_mongocli_generated",
		BuildDisplayName: "Generated post packaging smoke tests (Docker / mongocli)",
		DistroRunOn:      []string{runOn},
	}

	for _, os := range oses {
		t := &shrub.Task{
			Name: fmt.Sprintf("pkg_test_mongocli_docker_%v", os),
		}
		t = t.Dependency(shrub.TaskDependency{
			Name:    "package_goreleaser",
			Variant: "goreleaser_mongocli_snapshot",
		}).Function("clone").FunctionWithVars("docker build", map[string]string{
			"image": postPkgImg[os],
		})
		c.Tasks = append(c.Tasks, t)
		v.AddTasks(t.Name)
	}

	c.Variants = append(c.Variants, v)
}

func PublishStableTasks(c *shrub.Configuration) {
	dependency := []shrub.TaskDependency{
		{
			Name:    "compile",
			Variant: "code_health",
		},
		{
			Name:    "package_goreleaser",
			Variant: "release_mongocli_github",
		},
	}
	for _, sv := range serverVersions {
		v := &shrub.Variant{
			BuildName:        fmt.Sprintf("generated_release_mongocli_publish_%s", strings.ReplaceAll(sv, ".", "")),
			BuildDisplayName: fmt.Sprintf("Publish mongocli yum/apt %s", sv),
			DistroRunOn:      []string{"rhel80-small"},
		}
		publishVariant(
			c,
			v,
			sv,
			"_stable",
			dependency,
			true,
		)
	}
}

func PublishSnapshotTasks(c *shrub.Configuration) {
	dependency := []shrub.TaskDependency{
		{
			Name:    "compile",
			Variant: "code_health",
		},
		{
			Name:    "package_goreleaser",
			Variant: "goreleaser_mongocli_snapshot",
		},
	}
	v := c.Variant("publish_mongocli_snapshot")
	publishVariant(
		c,
		v,
		"5.0",
		"",
		dependency,
		false,
	)
}
