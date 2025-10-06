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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
)

var ErrDockerNotFound = fmt.Errorf("%w: docker not found in your system, check requirements at https://dochub.mongodb.org/core/atlas-cli-deploy-local-reqs", ErrContainerEngineNotFound)
var ErrDeterminingDockerVersion = errors.New("could not determine docker version")
var errParseHealthCheck = errors.New("parsing image healthcheck failed")
var errListContainer = errors.New("container listing failed")
var errParsingContainer = errors.New("container parsing failed")
var errDecodingJSON = errors.New("container decoding failed")
var errParsingPorts = errors.New("parsing ports failed")
var errConvertHostPort = errors.New("converting host port failed")
var errConvertContainerPort = errors.New("converting container port failed")
var minDockerVersion = semver.New(27, 0, 0, "", "") //nolint:mnd

type dockerImpl struct {
}

func newDockerEngine() Engine {
	return &dockerImpl{}
}

func (*dockerImpl) Name() string {
	return "docker"
}

func (*dockerImpl) Ready() error {
	_, err := exec.LookPath("docker")
	if errors.Is(err, exec.ErrDot) {
		err = nil
	}
	if err != nil {
		return ErrDockerNotFound
	}
	return nil
}

func (e *dockerImpl) VerifyVersion(ctx context.Context) error {
	versionBytes, err := e.run(ctx, "version", "--format", "v{{.Client.Version}}")
	if err != nil {
		return errors.Join(ErrDeterminingDockerVersion, err)
	}

	version, err := semver.NewVersion(strings.TrimSpace(string(versionBytes)))
	if err != nil {
		return errors.Join(ErrDeterminingDockerVersion, err)
	}

	if version.Compare(minDockerVersion) == -1 {
		_, _ = log.Warningf("Detected docker version %s, the minimum supported docker version is %s.\n", version.String(), minDockerVersion.String())
	}

	return nil
}

func (*dockerImpl) run(ctx context.Context, args ...string) ([]byte, error) {
	cmd := exec.CommandContext(ctx, "docker", args...)
	buf, err := cmd.Output()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		err = fmt.Errorf("%w: %s", exitErr, exitErr.Stderr)
	}
	return buf, err
}

func (e *dockerImpl) ContainerLogs(ctx context.Context, name string) ([]string, error) {
	buf, err := e.run(ctx, "container", "logs", name)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(buf), "\n"), nil
}

func parsePortMapping(s string) ([]PortMapping, error) {
	if s == "" {
		return nil, nil
	}

	mappings := strings.Split(s, ",")
	result := []PortMapping{}
	for _, mapping := range mappings {
		hostMapping, containerMapping := "", mapping
		if strings.Contains(mapping, "->") {
			segments := strings.SplitN(mapping, "->", 2) //nolint //max 2 fields
			hostMapping = segments[0]
			containerMapping = segments[1]
		}

		hostStr, hostPort, err := splitHostPort(hostMapping)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", errConvertHostPort, err)
		}
		containerPort, containerProtocol, err := splitPortProtocol(containerMapping)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", errConvertContainerPort, err)
		}
		result = append(result, PortMapping{
			HostAddress:       hostStr,
			HostPort:          hostPort,
			ContainerPort:     containerPort,
			ContainerProtocol: containerProtocol,
		})
	}
	return result, nil
}

func splitPortProtocol(s string) (int, string, error) {
	protocol := ""
	if strings.ContainsRune(s, '/') {
		index := strings.LastIndex(s, "/")
		protocol = s[index+1:]
		s = s[:index]
	}

	port, err := strconv.Atoi(s)

	return port, protocol, err
}

func splitHostPort(s string) (string, int, error) {
	host, port := "", s

	if strings.ContainsRune(s, ':') {
		index := strings.LastIndex(s, ":")
		host = s[:index]
		port = s[index+1:]
	}

	iport, err := strconv.Atoi(port)
	if err != nil {
		host = port
		if !strings.ContainsRune(s, ':') { // single value can be either host or port
			err = nil
		}
	}
	return host, iport, err
}

func portMappingFlag(pm PortMapping) string {
	result := ""

	if pm.HostAddress != "" {
		result += pm.HostAddress + ":"
	}

	if pm.HostPort != 0 {
		result += strconv.Itoa(pm.HostPort)
	}

	result += ":" + strconv.Itoa(pm.ContainerPort)

	if pm.ContainerProtocol != "" {
		result += "/" + pm.ContainerProtocol
	}

	return result
}

func portsFlags(flags *RunFlags) []string {
	if flags == nil {
		return nil
	}
	args := []string{}
	if flags.Ports != nil {
		for _, mapping := range flags.Ports {
			mapping.HostAddress = "127.0.0.1"
			if flags.BindIPAll != nil && *flags.BindIPAll {
				mapping.HostAddress = ""
			}
			args = append(args, "-p", portMappingFlag(mapping))
		}
	}
	return args
}

func runFlags(flags *RunFlags) []string {
	if flags == nil {
		return nil
	}
	args := portsFlags(flags)

	if flags.Detach != nil && *flags.Detach {
		args = append(args, "--detach")
	}

	if flags.Remove != nil && *flags.Remove {
		args = append(args, "--rm")
	}

	if flags.Name != nil {
		args = append(args, "--name", *flags.Name)
	}

	if flags.Hostname != nil {
		args = append(args, "--hostname", *flags.Hostname)
	}

	if flags.Env != nil {
		for key, value := range flags.Env {
			args = append(args, "-e", fmt.Sprintf("%s=%s", key, value))
		}
	}

	if flags.Network != nil {
		args = append(args, "--network", *flags.Network)
	}

	if flags.IP != nil {
		args = append(args, "--ip", *flags.IP)
	}

	if flags.Entrypoint != nil {
		args = append(args, "--entrypoint", *flags.Entrypoint)
	}

	if flags.Volumes != nil {
		for _, value := range flags.Volumes {
			args = append(args, "-v", fmt.Sprintf("%s:%s", value.HostPath, value.ContainerPath))
		}
	}

	return args
}

func (e *dockerImpl) ContainerRun(ctx context.Context, image string, flags *RunFlags) (string, error) {
	args := []string{"run"}
	args = append(args, runFlags(flags)...)
	args = append(args, image)
	if flags != nil && flags.Cmd != nil {
		args = append(args, *flags.Cmd)
	}

	if flags != nil && flags.Args != nil {
		args = append(args, flags.Args...)
	}

	buf, err := e.run(ctx, args...)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func parseContainers(buf []byte) ([]Container, error) {
	_, _ = log.Debugf("parsing containers: %s", string(buf))
	result := []Container{}
	decoder := json.NewDecoder(bytes.NewBuffer(buf))
	for decoder.More() {
		c := map[string]any{}

		if err := decoder.Decode(&c); err != nil {
			return nil, fmt.Errorf("%w: %w", errDecodingJSON, err)
		}

		cont := Container{
			ID:    c["ID"].(string),
			Names: []string{c["Names"].(string)},
			State: c["State"].(string),
			Image: c["Image"].(string),
		}

		pm, err := parsePortMapping(c["Ports"].(string))
		if err != nil {
			return nil, fmt.Errorf("%w: %w", errParsingPorts, err)
		}
		cont.Ports = pm

		labels := c["Labels"].(string)
		cont.Labels = map[string]string{}
		for label := range strings.SplitSeq(labels, ",") {
			segments := strings.SplitN(label, "=", 2) //nolint //max 2 fields
			if len(segments) == 2 {                   //nolint //max 2 fields
				cont.Labels[segments[0]] = segments[1]
			} else {
				cont.Labels[segments[0]] = ""
			}
		}

		result = append(result, cont)
	}

	return result, nil
}

func (e *dockerImpl) ContainerList(ctx context.Context, labels ...string) ([]Container, error) {
	args := []string{"container", "ls", "--all", "--format", "json"}

	if len(labels) > 0 {
		for _, label := range labels {
			args = append(args, "-f", "label="+label)
		}
	}
	buf, err := e.run(ctx, args...)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errListContainer, err)
	}

	if len(buf) == 0 {
		return nil, nil
	}

	list, err := parseContainers(buf)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errParsingContainer, err)
	}
	return list, nil
}

func (e *dockerImpl) ContainerRm(ctx context.Context, names ...string) error {
	args := []string{"container", "rm", "-v", "-f"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *dockerImpl) ContainerStart(ctx context.Context, names ...string) error {
	args := []string{"container", "start"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *dockerImpl) ContainerStop(ctx context.Context, names ...string) error {
	args := []string{"container", "stop"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *dockerImpl) ContainerUnpause(ctx context.Context, names ...string) error {
	args := []string{"container", "unpause"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *dockerImpl) ContainerInspect(ctx context.Context, names ...string) ([]*InspectData, error) {
	args := []string{"container", "inspect", "--format", "json"}
	args = append(args, names...)

	buf, err := e.run(ctx, args...)
	if err != nil {
		return nil, err
	}

	result := []*InspectData{}
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (e *dockerImpl) ImageList(ctx context.Context, references ...string) ([]Image, error) {
	args := []string{"image", "ls", "--format", "{{. | json}}"}

	if len(references) > 0 {
		for _, name := range references {
			args = append(args, "-f", "reference="+name)
		}
	}
	buf, err := e.run(ctx, args...)
	if err != nil {
		return nil, err
	}

	if len(buf) == 0 {
		return nil, nil
	}

	result, err := readJsonl[Image](bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func readJsonl[T any](r io.Reader) ([]T, error) {
	data := []T{}
	decoder := json.NewDecoder(r)
	for decoder.More() {
		var entry T
		if err := decoder.Decode(&entry); err != nil {
			return data, err
		}
		data = append(data, entry)
	}

	if len(data) == 0 {
		return nil, nil
	}

	return data, nil
}

func (e *dockerImpl) ImagePull(ctx context.Context, name string) error {
	_, err := e.run(ctx, "image", "pull", name)
	return err
}

func (e *dockerImpl) ImageHealthCheck(ctx context.Context, name string) (*ImageHealthCheck, error) {
	b, err := e.run(ctx, "image", "inspect", "--format", "json", name)
	if err != nil {
		return nil, err
	}

	type PartialImageInspectConfigHealth struct {
		Test        []string      `json:",omitempty"`
		StartPeriod time.Duration `json:",omitempty"`
		Interval    time.Duration `json:",omitempty"`
		Timeout     time.Duration `json:",omitempty"`
		Retries     int           `json:",omitempty"`
	}

	type PartialImageInspectConfig struct {
		Healthcheck *PartialImageInspectConfigHealth
	}

	type PartialImageInspect struct {
		Config PartialImageInspectConfig
	}

	var inspectOutput []PartialImageInspect
	if err := json.Unmarshal(b, &inspectOutput); err != nil {
		_, _ = log.Debug("failed json parsing: " + string(b))
		return nil, fmt.Errorf("%w: %w", errParseHealthCheck, err)
	}

	if len(inspectOutput) != 1 {
		return nil, fmt.Errorf("expected 1 output, got %v", len(inspectOutput))
	}

	healthCheck := inspectOutput[0].Config.Healthcheck

	if healthCheck == nil {
		return nil, nil
	}

	return &ImageHealthCheck{
		Test:        healthCheck.Test,
		Interval:    &healthCheck.Interval,
		Timeout:     &healthCheck.Timeout,
		StartPeriod: &healthCheck.StartPeriod,
		Retries:     &healthCheck.Retries,
	}, nil
}

func (e *dockerImpl) ContainerHealthStatus(ctx context.Context, name string) (DockerHealthcheckStatus, error) {
	buf, err := e.run(ctx, "inspect", "--format", "{{.State.Health.Status}}", name)
	if err != nil {
		return "", err
	}

	statusString := strings.TrimSpace(string(buf))
	switch statusString {
	case string(DockerHealthcheckStatusNone):
		return DockerHealthcheckStatusNone, nil
	case string(DockerHealthcheckStatusStarting):
		return DockerHealthcheckStatusStarting, nil
	case string(DockerHealthcheckStatusHealthy):
		return DockerHealthcheckStatusHealthy, nil
	case string(DockerHealthcheckStatusUnhealthy):
		return DockerHealthcheckStatusUnhealthy, nil
	case "": // Empty string means no health check configured or status not available
		return DockerHealthcheckStatusNone, nil
	default:
		return "", fmt.Errorf("unknown health status: %s", statusString)
	}
}

func (e *dockerImpl) Version(ctx context.Context) (map[string]any, error) {
	buf, err := e.run(ctx, "version", "--format", "json")
	if err != nil {
		return nil, err
	}

	result := map[string]any{}
	if err := json.Unmarshal(buf, &result); err != nil {
		return nil, err
	}
	return result, nil
}
