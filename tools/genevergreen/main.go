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

	"github.com/andreangiolillo/mongocli-test/tools/genevergreen/generate"
	"github.com/evergreen-ci/shrub"
)

const (
	atlascli = "atlascli"
	mongocli = "mongocli"
)

var (
	ErrMissingOption = errors.New("missing option")
)

func run() error {
	var toolName, taskType string

	flag.StringVar(&taskType, "tasks", "", "type of task to be generated")
	flag.StringVar(&toolName, "tool_name", "", fmt.Sprintf("Tool to generate tasks for (%s or %s)", atlascli, mongocli))

	flag.Parse()

	if toolName == "" {
		return fmt.Errorf("%w: %s", ErrMissingOption, "tool_name")
	}

	if toolName != atlascli && toolName != mongocli {
		return fmt.Errorf("-tool_name must be either %q or %q", atlascli, mongocli)
	}

	if taskType == "" {
		return fmt.Errorf("%w: %s", ErrMissingOption, "tasks")
	}

	c := &shrub.Configuration{}

	switch taskType {
	case "repo":
		generate.RepoTasks(c, toolName)
	case "postpkg":
		generate.PostPkgTasks(c, toolName)
		generate.PostPkgMetaTasks(c, toolName)
	case "snapshot":
		generate.PublishSnapshotTasks(c, toolName)
	case "publish":
		generate.PublishStableTasks(c, toolName)
	case "local":
		generate.LocalDeploymentTasks(c, toolName)
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

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
