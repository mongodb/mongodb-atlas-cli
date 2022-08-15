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
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/evergreen-ci/shrub"
)

const (
	atlascli = "atlascli"
	mongocli = "mongocli"
)

var (
	serverVersions = []string{"4.2", "4.4", "5.0"}
	oses           = []string{"amazonlinux2", "centos7", "centos8", "rhel9", "debian9", "debian10", "ubuntu18.04", "ubuntu20.04", "ubuntu22.04"}
	repos          = []string{"org", "enterprise"}
)

func buildDependency(toolName, os, serverVersion, repo string) shrub.TaskDependency {
	newOs := map[string]string{
		"centos7":      "rhel70",
		"centos8":      "rhel80",
		"rhel90":       "rhel90",
		"amazonlinux2": "amazon2",
		"ubuntu18.04":  "ubuntu1804",
		"ubuntu20.04":  "ubuntu2004",
		"ubuntu22.04":  "ubuntu2204",
		"debian9":      "debian92",
		"debian10":     "debian10",
	}

	return shrub.TaskDependency{
		Name:    fmt.Sprintf("push_%v_%v_%v_stable", toolName, newOs[os], repo),
		Variant: fmt.Sprintf("release_%v_publish_%v", toolName, strings.ReplaceAll(serverVersion, ".", "")),
	}
}

func generateRepoTasks(toolName string) *shrub.Configuration {
	c := &shrub.Configuration{}

	for _, serverVersion := range serverVersions {
		v := &shrub.Variant{
			BuildName:        fmt.Sprintf("test_repo_%v_%v", toolName, serverVersion),
			BuildDisplayName: fmt.Sprintf("Test %v on repo %v", toolName, serverVersion),
			DistroRunOn:      []string{"ubuntu1804-small"},
		}

		pkg := "mongodb-atlas-cli"
		entrypoint := "atlas"
		if toolName == mongocli {
			pkg = mongocli
			entrypoint = mongocli
		}

		for _, os := range oses {
			for _, repo := range repos {
				mongoRepo := "https://repo.mongodb.com"
				if repo == "org" {
					mongoRepo = "https://repo.mongodb.org"
				}

				t := &shrub.Task{
					Name: fmt.Sprintf("test_repo_%v_%v_%v_%v", toolName, os, repo, serverVersion),
				}
				t = t.Stepback(false).GitTagOnly(true).Dependency(buildDependency(toolName, os, serverVersion, repo)).Function("clone").FunctionWithVars("docker build repo", map[string]string{
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

	return c
}

func run() error {
	var toolName string

	flag.StringVar(&toolName, "tool_name", "", fmt.Sprintf("Tool to generate tasks for (%v or %v)", atlascli, mongocli))

	flag.Parse()

	if toolName == "" {
		return errors.New("-tool_name missing")
	}

	if toolName != atlascli && toolName != mongocli {
		return fmt.Errorf("-tool_name must be either '%v' or '%v'", atlascli, mongocli)
	}

	c := generateRepoTasks(toolName)
	var b []byte
	b, err := json.MarshalIndent(c, "", "\t")

	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
