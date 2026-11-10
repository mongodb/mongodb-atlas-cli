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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/cmd/genevergreen/generate"
)

var (
	ErrMissingOption = errors.New("missing option")
)

func run() error {
	var taskType string
	var variant string
	var taskPrefix string
	var includeGroups string
	var excludeGroups string
	var requireFixtures string
	var exclude string
	var profile string

	flag.StringVar(&taskType, "tasks", "", "type of task to be generated")
	flag.StringVar(&variant, "variant", "", "e2e: variant to append generated tasks to")
	flag.StringVar(&taskPrefix, "task-prefix", "", "e2e: prefix for generated task names (defaults to variant + \"_\")")
	flag.StringVar(&includeGroups, "include-groups", "", "e2e: comma-separated groups to include (packages with any of these groups)")
	flag.StringVar(&excludeGroups, "exclude-groups", "", "e2e: comma-separated groups to exclude (packages with any of these groups)")
	flag.StringVar(&requireFixtures, "require-fixtures", "", "e2e: comma-separated fixtures a package must declare to be selected")
	flag.StringVar(&exclude, "exclude", "", "e2e: comma-separated fixtures to exclude")
	flag.StringVar(&profile, "profile", "", "e2e: E2E_PROFILE_NAME to set on generated tasks")

	flag.Parse()

	if taskType == "" {
		return fmt.Errorf("%w: %s", ErrMissingOption, "tasks")
	}

	c := &shrub.Configuration{}

	switch taskType {
	case "repo":
		generate.RepoTasks(c)
	case "postpkg":
		generate.PostPkgTasks(c)
		generate.PostPkgMetaTasks(c)
	case "snapshot":
		generate.PublishSnapshotTasks(c)
	case "publish":
		generate.PublishStableTasks(c)
	case "e2e":
		if variant == "" {
			return fmt.Errorf("%w: %s", ErrMissingOption, "variant")
		}
		if err := generate.E2ETasks(c, generate.E2EOptions{
			Variant:          variant,
			TaskPrefix:       taskPrefix,
			Profile:          profile,
			IncludeAnyGroups: splitFlag(includeGroups),
			ExcludeGroups:    splitFlag(excludeGroups),
			RequireFixtures:  splitFlag(requireFixtures),
			ExcludeFixtures:  splitFlag(exclude),
		}); err != nil {
			return err
		}
	case "e2e-manifests":
		manifests, err := generate.E2EManifests(generate.E2EOptions{
			IncludeAnyGroups: splitFlag(includeGroups),
			ExcludeGroups:    splitFlag(excludeGroups),
			RequireFixtures:  splitFlag(requireFixtures),
			ExcludeFixtures:  splitFlag(exclude),
		})
		if err != nil {
			return err
		}
		fmt.Println(string(manifests))
		return nil
	default:
		return errors.New("-tasks is invalid")
	}

	var b []byte
	b, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

func splitFlag(value string) []string {
	values := make([]string, 0)
	for _, entry := range strings.Split(value, ",") {
		entry = strings.TrimSpace(entry)
		if entry != "" {
			values = append(values, entry)
		}
	}
	return values
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
