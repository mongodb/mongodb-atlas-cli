// Copyright 2023 MongoDB Inc
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

package podman

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
)

var (
	ErrPodmanNotFound = errors.New("podman not found in your system, check requirements at https://dochub.mongodb.org/core/atlas-cli-deploy-local-reqs")
)

type RunContainerOpts struct {
	Detach            bool
	Remove            bool
	Image             string
	Name              string
	Hostname          string
	Volumes           map[string]string
	Ports             map[int]int
	BindIPAll         bool
	Network           string
	EnvVars           map[string]string
	Args              []string
	Entrypoint        string
	Cmd               string
	IP                string
	HealthCmd         *[]string
	HealthInterval    *time.Duration
	HealthTimeout     *time.Duration
	HealthStartPeriod *time.Duration
	HealthRetries     *int
}

type Container struct {
	ID    string   `json:"ID"`
	Names []string `json:"Names"`
	State string   `json:"State"`
	Image string   `json:"Image"`
	Ports []struct {
		HostPort      int `json:"host_port"`
		ContainerPort int `json:"container_port"`
	} `json:"Ports,omitempty"`
	Labels map[string]string `json:"Labels"`
}

type Image struct {
	ID          string   `json:"ID"`
	RepoTags    string   `json:"RepoTags"`
	RepoDigests []string `json:"RepoDigests"`
	Created     int      `json:"Created"`
	CreatedAt   string   `json:"CreatedAt"`
	Size        int      `json:"Size"`
	SharedSize  int      `json:"SharedSize"`
	VirtualSize int      `json:"VirtualSize"`
	Labels      struct {
		Architecture string `json:"architecture"`
		BuildDate    string `json:"build-date"`
		Description  string `json:"description"`
		Name         string `json:"name"`
		Version      string `json:"version"`
	} `json:"Labels"`
	Containers int `json:"Containers"`
	Names      []string
}

//go:generate mockgen -destination=../mocks/mock_podman.go -package=mocks github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman Client

type Client interface {
	Ready(ctx context.Context) error
	Version(ctx context.Context) (map[string]any, error)
	RunContainer(ctx context.Context, opts RunContainerOpts) ([]byte, error)
	RunHealthcheck(ctx context.Context, name string) error
	ContainerInspect(ctx context.Context, names ...string) ([]*InspectContainerData, error)
	StopContainers(ctx context.Context, names ...string) ([]byte, error)
	StartContainers(ctx context.Context, names ...string) ([]byte, error)
	UnpauseContainers(ctx context.Context, names ...string) ([]byte, error)
	RemoveContainers(ctx context.Context, names ...string) ([]byte, error)
	ListContainers(ctx context.Context, nameFilter string) ([]*Container, error)
	ListImages(ctx context.Context, nameFilter string) ([]*Image, error)
	PullImage(ctx context.Context, name string) ([]byte, error)
	ImageHealthCheck(ctx context.Context, name string) (*Schema2HealthConfig, error)
	ContainerHealthStatus(ctx context.Context, name string) (string, error)
	Logs(ctx context.Context) (map[string]any, []error)
	ContainerLogs(ctx context.Context, name string) ([]string, error)
	ContainerStatusAndUptime(ctx context.Context, name string) (string, time.Duration, error)
}

type client struct{}

func Installed() error {
	if _, err := exec.LookPath("podman"); err != nil {
		return ErrPodmanNotFound
	}
	return nil
}

func (*client) Ready(_ context.Context) error {
	return Installed()
}

func extractErrorMessage(exitErr *exec.ExitError) error {
	stderrLines := strings.Split(string(exitErr.Stderr), "\n")
	if len(stderrLines) < 2 { //nolint // expected to have at least 2 lines
		return fmt.Errorf("%w: %s", exitErr, string(exitErr.Stderr))
	}
	stderrLastLine := stderrLines[len(stderrLines)-2] // 2nd last line because last line should be empty
	return fmt.Errorf("%w: %s", exitErr, stderrLastLine)
}

func (*client) runPodman(ctx context.Context, arg ...string) ([]byte, error) {
	_, _ = log.Debug(fmt.Sprintln(append([]string{"podman"}, arg...)))

	cmd := exec.CommandContext(ctx, "podman", arg...)

	output, err := cmd.Output() // ignore stderr

	_, _ = log.Debug(string(output))

	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		_, _ = log.Debug(string(exitErr.Stderr))
		err = extractErrorMessage(exitErr)
	}

	return output, err
}

func (o *client) ContainerInspect(ctx context.Context, names ...string) ([]*InspectContainerData, error) {
	args := append([]string{"container", "inspect"}, names...)
	buf, err := o.runPodman(ctx, args...)
	if err != nil {
		return nil, err
	}

	var containers []*InspectContainerData
	err = json.Unmarshal(buf, &containers)
	return containers, err
}

func (o *client) ContainerHealthStatus(ctx context.Context, name string) (string, error) {
	buf, err := o.runPodman(ctx, "inspect", "--format", "{{.State.Health.Status}}", name)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(buf)), nil
}

//nolint:gocyclo
func (o *client) RunContainer(ctx context.Context, opts RunContainerOpts) ([]byte, error) {
	arg := []string{"run",
		"--name", opts.Name,
		"--hostname", opts.Hostname,
	}

	if opts.Network != "" {
		arg = append(arg, "--network", opts.Network)
	}

	for hostVolume, pathInContainer := range opts.Volumes {
		arg = append(arg, "-v", hostVolume+":"+pathInContainer)
	}

	for hostPort, containerPort := range opts.Ports {
		portMapping := strconv.Itoa(hostPort) + ":" + strconv.Itoa(containerPort)
		if !opts.BindIPAll {
			portMapping = "127.0.0.1:" + portMapping
		}
		arg = append(arg, "-p", portMapping)
	}

	for envVar, value := range opts.EnvVars {
		arg = append(arg, "-e", envVar+"="+value)
	}

	if opts.IP != "" {
		arg = append(arg, "--ip", opts.IP)
	}

	if opts.Detach {
		arg = append(arg, "-d")
	}

	if opts.Remove {
		arg = append(arg, "--rm")
	}

	if opts.Entrypoint != "" {
		arg = append(arg, "--entrypoint", opts.Entrypoint)
	}

	if opts.HealthCmd != nil && len(*opts.HealthCmd) > 0 {
		cmd := ""
		for i, c := range *opts.HealthCmd {
			if i != 0 {
				cmd += ","
			}

			cmd = cmd + "\"" + c + "\""
		}

		arg = append(arg, fmt.Sprintf("--health-cmd=[%s]", cmd))
	}
	if opts.HealthInterval != nil {
		arg = append(arg, "--health-interval", toMsString(*opts.HealthInterval))
	}
	if opts.HealthTimeout != nil {
		arg = append(arg, "--health-timeout", toMsString(*opts.HealthTimeout))
	}
	if opts.HealthStartPeriod != nil {
		arg = append(arg, "--health-start-period", toMsString(*opts.HealthStartPeriod))
	}
	if opts.HealthRetries != nil {
		arg = append(arg, "--health-startup-retries", strconv.Itoa(*opts.HealthRetries))
	}

	arg = append(arg, opts.Image)

	if opts.Cmd != "" {
		arg = append(arg, opts.Cmd)
	}

	arg = append(arg, opts.Args...)

	return o.runPodman(ctx, arg...)
}

func (o *client) RunHealthcheck(ctx context.Context, name string) error {
	_, err := o.runPodman(ctx, "healthcheck", "run", name)
	return err
}

func toMsString(duration time.Duration) string {
	return strconv.FormatInt(duration.Milliseconds(), 10) + "ms"
}

func (o *client) StopContainers(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"stop"}, names...)...)
}

func (o *client) StartContainers(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"start"}, names...)...)
}

func (o *client) PauseContainers(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"pause"}, names...)...)
}

func (o *client) UnpauseContainers(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"unpause"}, names...)...)
}

func (o *client) RemoveContainers(ctx context.Context, names ...string) ([]byte, error) {
	return o.runPodman(ctx, append([]string{"rm", "-f", "-v"}, names...)...)
}

func (o *client) ListContainers(ctx context.Context, label string) ([]*Container, error) {
	args := []string{"ps", "--all", "--format", "json"}
	if label != "" {
		args = append(args, "--filter", "label="+label)
	}

	response, err := o.runPodman(ctx, args...)
	if err != nil {
		return nil, err
	}

	var containers []*Container
	err = json.Unmarshal(response, &containers)
	return containers, err
}

func (o *client) ListImages(ctx context.Context, reference string) ([]*Image, error) {
	args := []string{"image", "list", "--format", "json"}

	if reference != "" {
		args = append(args, "--filter", "reference="+reference)
	}

	response, err := o.runPodman(ctx, args...)
	if err != nil {
		return nil, err
	}

	var images []*Image
	if err = json.Unmarshal(response, &images); err != nil {
		return nil, err
	}
	return images, err
}

func (o *client) PullImage(ctx context.Context, name string) ([]byte, error) {
	return o.runPodman(ctx, "pull", name)
}

func (o *client) ImageHealthCheck(ctx context.Context, name string) (*Schema2HealthConfig, error) {
	bytes, err := o.runPodman(ctx, "image", "inspect", "--format", "json", name)
	if err != nil {
		return nil, err
	}

	type PartialImageInspectConfig struct {
		Healthcheck *Schema2HealthConfig
	}

	type PartialImageInspect struct {
		Config PartialImageInspectConfig
	}

	var inspectOutput []PartialImageInspect
	if err := json.Unmarshal(bytes, &inspectOutput); err != nil {
		return nil, err
	}

	if len(inspectOutput) != 1 {
		return nil, fmt.Errorf("expected 1 output, got %v", len(inspectOutput))
	}

	return inspectOutput[0].Config.Healthcheck, nil
}

func (o *client) Version(ctx context.Context) (map[string]any, error) {
	output, err := o.runPodman(ctx, "version", "--format", "json")
	if err != nil {
		return nil, err
	}

	var version map[string]any
	if err = json.Unmarshal(output, &version); err != nil {
		return nil, err
	}
	return version, err
}

func (o *client) Logs(ctx context.Context) (map[string]any, []error) {
	l := map[string]any{}
	var errs []error
	output, err := o.runPodman(ctx, "ps", "--all", "--format", "json", "--filter", "name=mongo")
	if err != nil {
		errs = append(errs, err)
	} else {
		var podmanLogs []any
		if err = json.Unmarshal(output, &podmanLogs); err != nil {
			errs = append(errs, err)
		}
		l["Containers"] = podmanLogs
	}

	output, err = o.runPodman(ctx, "network", "ls", "--format", "json", "--filter", "name=mdb")
	if err != nil {
		errs = append(errs, err)
	} else {
		var networks []any
		if err = json.Unmarshal(output, &networks); err != nil {
			errs = append(errs, err)
		}
		l["Networks"] = networks
	}

	return l, errs
}

func (o *client) ContainerLogs(ctx context.Context, name string) ([]string, error) {
	output, err := o.runPodman(ctx, "container", "logs", name)
	if err != nil {
		return []string{""}, err
	}

	logs := strings.Split(string(output), "\n")
	return logs, nil
}

func (o *client) ContainerStatusAndUptime(ctx context.Context, name string) (string, time.Duration, error) {
	output, err := o.runPodman(ctx, "inspect", "--format", "[\"{{.State.Status}}\",\"{{.State.StartedAt}}\"]", name)
	if err != nil {
		return "", 0, err
	}

	var statusAndStartedAt []string
	if err = json.Unmarshal(output, &statusAndStartedAt); err != nil {
		return "", 0, err
	}

	const expectedArrayLength = 2
	if len(statusAndStartedAt) != expectedArrayLength {
		return "", 0, fmt.Errorf("parsing status and uptime: expected 2 output, got %v", len(statusAndStartedAt))
	}

	status := statusAndStartedAt[0]
	startedAt, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", statusAndStartedAt[1])
	if err != nil {
		return "", 0, err
	}

	return status, time.Since(startedAt), nil
}

func NewClient() Client {
	return &client{}
}
