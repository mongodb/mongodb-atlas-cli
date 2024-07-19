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
	"errors"
	"time"
)

type PortMapping struct {
	HostAddress       string
	HostPort          int
	ContainerPort     int
	ContainerProtocol string
}

type RunFlags struct {
	Name              *string
	Detach            *bool
	Remove            *bool
	Hostname          *string
	Ports             []PortMapping
	Env               map[string]string
	Cmd               *string
	Args              []string
	Network           *string
	IP                *string
	Entrypoint        *string
	BindIPAll         *bool
	HealthCmd         *[]string
	HealthInterval    *time.Duration
	HealthTimeout     *time.Duration
	HealthStartPeriod *time.Duration
	HealthRetries     *int
}

//go:generate mockgen -destination=../mocks/mock_container.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/container Engine

type Engine interface {
	Name() string
	Ready() error
	ContainerLogs(context.Context, string) ([]string, error)
	ContainerRun(context.Context, string, *RunFlags) (string, error)
	ContainerList(context.Context, ...string) ([]Container, error)
	ContainerRm(context.Context, ...string) error
	ContainerStart(context.Context, ...string) error
	ContainerStop(context.Context, ...string) error
	ContainerUnpause(context.Context, ...string) error
	ContainerInspect(context.Context, ...string) ([]*InspectData, error)
	ContainerHealthStatus(context.Context, string) (DockerHealthcheckStatus, error)
	ImageList(context.Context, ...string) ([]Image, error)
	ImagePull(context.Context, string) error
	ImageHealthCheck(context.Context, string) (*ImageHealthCheck, error)
	Version(context.Context) (map[string]any, error)
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

type ImageHealthCheck struct {
	Test        []string
	Interval    *time.Duration
	Timeout     *time.Duration
	StartPeriod *time.Duration
	Retries     *int
}

type Container struct {
	ID     string
	Names  []string
	State  string
	Image  string
	Ports  []PortMapping
	Labels map[string]string
}

type InspectDataConfig struct {
	Labels map[string]string `json:"Labels"`
}

type InspectData struct {
	ID         string                 `json:"Id"`
	Name       string                 `json:"Name"`
	Config     *InspectDataConfig     `json:"Config"`
	HostConfig *InspectDataHostConfig `json:"HostConfig"`
}

type InspectDataHostConfig struct {
	PortBindings map[string][]InspectDataHostPort `json:"PortBindings"`
}

type InspectDataHostPort struct {
	HostIP   string `json:"HostIp"`
	HostPort string `json:"HostPort"`
}

type DockerHealthcheckStatus string

const (
	DockerHealthcheckStatusStarting  DockerHealthcheckStatus = "starting"
	DockerHealthcheckStatusHealthy   DockerHealthcheckStatus = "healthy"
	DockerHealthcheckStatusUnhealthy DockerHealthcheckStatus = "unhealthy"
	DockerHealthcheckStatusNone      DockerHealthcheckStatus = "none"
)

var (
	ErrContainerEngineNotFound = errors.New("container engine not found")
)
