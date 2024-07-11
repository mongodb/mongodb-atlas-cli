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
	"slices"
	"strings"

	"github.com/evergreen-ci/shrub"
)

const (
	runOn = "ubuntu2004-small"
)

var (
	serverVersions = []string{
		"5.0",
		"6.0",
		"7.0",
		"8.0",
	}

	unsupportedNewOsByVersion = map[string][]string{
		"8.0": {"debian11"},
		"7.0": {"ubuntu2404"},
		"6.0": {"ubuntu2404"},
		"5.0": {"ubuntu2404"},
	}

	oses = []string{
		"centos8",
		"rhel9",
		"debian11",
		"debian12",
		"ubuntu20.04",
		"ubuntu22.04",
		"ubuntu24.04",
	}
	repos      = []string{"org", "enterprise"}
	postPkgImg = map[string]string{
		"centos8":     "centos8-rpm",
		"rhel9":       "rhel9-rpm",
		"ubuntu20.04": "ubuntu20.04-deb",
		"ubuntu22.04": "ubuntu22.04-deb",
		"ubuntu24.04": "ubuntu24.04-deb",
		"debian11":    "debian11-deb",
		"debian12":    "debian12-deb",
	}
	newOs = map[string]string{
		"centos8":         "rhel80",
		"rhel9":           "rhel90",
		"ubuntu20.04":     "ubuntu2004",
		"ubuntu22.04":     "ubuntu2204",
		"ubuntu24.04":     "ubuntu2404",
		"debian11":        "debian11",
		"debian12":        "debian12",
		"amazonlinux2023": "amazon2023",
	}
)

func newDependency(os, serverVersion, repo string) shrub.TaskDependency {
	return shrub.TaskDependency{
		Name:    fmt.Sprintf("push_atlascli_%s_%s_%s_%s_stable", newOs[os], repo, x86_64, strings.ReplaceAll(serverVersion, ".", "")),
		Variant: buildNamePrefix + strings.ReplaceAll(serverVersion, ".", ""),
	}
}

func RepoTasks(c *shrub.Configuration) {
	for _, serverVersion := range serverVersions {
		v := &shrub.Variant{
			BuildName:        fmt.Sprintf("test_repo_atlascli_%v", serverVersion),
			BuildDisplayName: fmt.Sprintf("Test atlascli on repo %v", serverVersion),
			DistroRunOn:      []string{runOn},
		}

		pkg := "mongodb-atlas-cli"
		entrypoint := "atlas"

		for _, os := range oses {
			for _, repo := range repos {
				mongoRepo := "https://repo.mongodb.com"
				if repo == "org" {
					mongoRepo = "https://repo.mongodb.org"
				}

				t := &shrub.Task{
					Name: fmt.Sprintf("test_repo_atlascli_%v_%v_%v", os, repo, serverVersion),
				}

				if slices.Contains(unsupportedNewOsByVersion[serverVersion], newOs[os]) {
					continue
				}

				t = t.Stepback(false).
					GitTagOnly(true).
					Dependency(newDependency(os, serverVersion, repo)).
					Function("clone").
					FunctionWithVars("docker build repo", map[string]string{
						"server_version": serverVersion,
						"package":        pkg,
						"entrypoint":     entrypoint,
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
		BuildName:        "pkg_smoke_tests_docker_atlascli_generated",
		BuildDisplayName: "Generated post packaging smoke tests (Docker / atlascli)",
		DistroRunOn:      []string{runOn},
	}

	for _, os := range oses {
		t := &shrub.Task{
			Name: fmt.Sprintf("pkg_test_atlascli_docker_%v", os),
		}
		t = t.Dependency(shrub.TaskDependency{
			Name:    "package_goreleaser",
			Variant: "goreleaser_atlascli_snapshot",
		}).Function("clone").FunctionWithVars("docker build", map[string]string{
			"image": postPkgImg[os],
		})
		c.Tasks = append(c.Tasks, t)
		v.AddTasks(t.Name)
	}

	c.Variants = append(c.Variants, v)
}

func PostPkgMetaTasks(c *shrub.Configuration) {
	v := &shrub.Variant{
		BuildName:        "pkg_smoke_tests_docker_meta_atlascli_generated",
		BuildDisplayName: "Generated post packaging smoke tests (Meta / atlascli)",
		DistroRunOn:      []string{runOn},
	}

	for _, os := range oses {
		for _, sv := range serverVersions {
			if slices.Contains(unsupportedNewOsByVersion[sv], newOs[os]) {
				continue
			}

			t := &shrub.Task{
				Name: "pkg_test_atlascli_meta_docker_" + sv + "_" + os,
			}
			t = t.Dependency(shrub.TaskDependency{
				Name:    "package_goreleaser",
				Variant: "goreleaser_atlascli_snapshot",
			}).Function("clone").
				FunctionWithVars("docker build meta", map[string]string{
					"image":          postPkgImg[os],
					"server_version": sv,
				})

			// TODO: Re-enable meta package tests in 8.0 until mongosh is added.
			if sv == "8.0" {
				disable := true
				t.Disable = &disable
			}
			c.Tasks = append(c.Tasks, t)
			v.AddTasks(t.Name)
		}
	}

	c.Variants = append(c.Variants, v)
}

const buildNamePrefix = "generated_release_atlascli_publish_"

func PublishStableTasks(c *shrub.Configuration) {
	dependency := []shrub.TaskDependency{
		{
			Name:    "compile",
			Variant: "code_health",
		},
		{
			Name:    "package_goreleaser",
			Variant: "release_atlascli_github",
		},
	}

	for _, sv := range serverVersions {
		v := &shrub.Variant{
			BuildName:        buildNamePrefix + strings.ReplaceAll(sv, ".", ""),
			BuildDisplayName: "Publish atlascli yum/apt " + sv,
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
			Variant: "goreleaser_atlascli_snapshot",
		},
	}
	v := c.Variant("publish_atlascli_snapshot")
	publishVariant(
		c,
		v,
		"8.0",
		"",
		dependency,
		false,
	)
}
