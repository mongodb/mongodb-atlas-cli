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
	Network    *string
	IP         *string
	Entrypoint *string
	BindIPAll  *bool
}

//go:generate mockgen -destination=../mocks/mock_container.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container Engine

type Engine interface {
	Ready(context.Context) error
	ContainerLogs(context.Context, string) ([]string, error)
	ContainerRun(context.Context, string, *RunFlags) (string, error)
	ContainerList(context.Context, ...string) ([]Container, error)
	ContainerRm(context.Context, ...string) error
	ImageList(context.Context, ...string) ([]Image, error)
	ImagePull(context.Context, string) error
}

type Image struct {
	ID          string
	RepoTags    string
	RepoDigests []string
	Created     int
	CreatedAt   string
	Size        int
	SharedSize  int
	VirtualSize int
	Labels      struct {
		Architecture string `json:"architecture"`
		BuildDate    string `json:"build-date"`
		Description  string `json:"description"`
		Name         string `json:"name"`
		Version      string `json:"version"`
	}
	Containers int
	Names      []string
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
