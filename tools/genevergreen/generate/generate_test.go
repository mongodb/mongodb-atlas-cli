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

//go:build unit

package generate

import (
	"strings"
	"testing"

	"github.com/evergreen-ci/shrub"
	"github.com/stretchr/testify/assert"
)

const (
	push = "push"
)

func TestPublishSnapshotTasks(t *testing.T) {
	c := &shrub.Configuration{}
	PublishSnapshotTasks(c)
	assert.Len(t, c.Tasks, 28)
	commandFound := false
	for _, task := range c.Tasks {
		for _, c := range task.Commands {
			if c.FunctionName != push {
				continue
			}
			commandFound = true
			distro := c.Vars["distro"]
			serverVersion := c.Vars["server_version"]
			assert.NotContains(t, unsupportedNewOsByVersion[serverVersion], distro)
			assert.NotEmpty(t, distro)
		}
	}

	assert.True(t, commandFound, "expected to find a push command")
	assert.Len(t, c.Variants, 2)
}

func TestPublishStableTasks(t *testing.T) {
	c := &shrub.Configuration{}
	PublishStableTasks(c)

	commandFound := false
	for _, task := range c.Tasks {
		for _, c := range task.Commands {
			if c.FunctionName != push {
				continue
			}
			commandFound = true
			distro := c.Vars["distro"]
			serverVersion := c.Vars["server_version"]
			// ensure unsupportedNewOs is not used
			assert.NotContains(t, unsupportedNewOsByVersion[serverVersion], distro)
			assert.NotEmpty(t, distro)
		}
	}

	assert.True(t, commandFound, "expected to find a push command")
	assert.Len(t, c.Variants, 4)
	assert.Len(t, c.Tasks, 112)
}

func TestPostPkgMetaTasks(t *testing.T) {
	c := &shrub.Configuration{}
	PostPkgMetaTasks(c)
	// validate server / distro
	for _, task := range c.Tasks {
		for _, c := range task.Commands {
			if !strings.Contains(c.FunctionName, "docker build meta") {
				continue
			}
			image := c.Vars["image"]
			serverVersion := c.Vars["server_version"]
			// find the key from the image
			for key, value := range postPkgImg {
				if value == image {
					assert.NotContains(t, unsupportedNewOsByVersion[serverVersion], newOs[key])
				}
			}
		}
	}
	assert.Len(t, c.Variants, 1)
	assert.Len(t, c.Tasks, 24)
}

func TestRepoTasks(t *testing.T) {
	c := &shrub.Configuration{}
	RepoTasks(c)
	// validate server / distro
	for _, task := range c.Tasks {
		for _, c := range task.Commands {
			if c.FunctionName != "docker build repo" {
				continue
			}
			image := c.Vars["image"]
			serverVersion := c.Vars["server_version"]
			// ensure unsupportedNewOs is not used
			assert.NotContains(t, unsupportedNewOsByVersion[serverVersion], image)
			assert.NotEmpty(t, image)
		}
	}

	assert.Len(t, c.Variants, 4)
	assert.Len(t, c.Tasks, 48)
}
