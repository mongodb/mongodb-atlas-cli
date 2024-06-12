// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

import (
	"context"
)

type Port struct {
	HostPort      int
	ContainerPort int
}

type RunFlags struct {
	Name       *string
	Detach     *bool
	Remove     *bool
	Hostname   *string
	Ports      []Port
	Env        map[string]string
	Cmd        *string
	Args       []string
	Volumes    map[string]string
	Network    *string
	IP         *string
	Entrypoint *string
	BindIPAll  *bool
}

//go:generate mockgen -destination=../mocks/mock_container.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container Engine

type Engine interface {
	ContainerLogs(context.Context, string) ([]string, error)
	ContainerRun(context.Context, string, *RunFlags) (string, error)
	ContainerList(context.Context, ...string) ([]Container, error)
}

type Container struct {
	ID     string
	Names  []string
	State  string
	Image  string
	Ports  []Port
	Labels map[string]string
}

func New() Engine {
	return newPodmanEngine()
}
