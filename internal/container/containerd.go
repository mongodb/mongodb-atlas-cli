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
	"net"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
)

const (
	darwinOS    = "darwin"
	localhostIP = "127.0.0.1"
)

var ErrNerdctlNotFound = fmt.Errorf("%w: container engine not found in your system, check requirements at https://dochub.mongodb.org/core/atlas-cli-deploy-local-reqs", ErrContainerEngineNotFound)
var ErrDeterminingContainerdVersion = errors.New("could not determine containerd version")
var minContainerdVersion = semver.New(0, 0, 0, "", "")

type containerdImpl struct {
	useDirectNerdctl bool
}

func newContainerdEngine() Engine {
	impl := &containerdImpl{}
	impl.detectNerdctlMethod()
	return impl
}

func (*containerdImpl) Name() string {
	return "containerd"
}

func (e *containerdImpl) detectNerdctlMethod() {
	if runtime.GOOS == darwinOS {
		// On macOS, use lima nerdctl (direct nerdctl won't work)
		e.useDirectNerdctl = false
	} else {
		// On non-macOS, use direct nerdctl
		e.useDirectNerdctl = true
	}
}

func (e *containerdImpl) Ready() error {
	if e.useDirectNerdctl {
		_, err := exec.LookPath("nerdctl")
		if errors.Is(err, exec.ErrDot) {
			err = nil
		}
		if err != nil {
			return ErrNerdctlNotFound
		}
	} else {
		_, err := exec.LookPath("lima")
		if errors.Is(err, exec.ErrDot) {
			err = nil
		}
		if err != nil {
			return ErrNerdctlNotFound
		}
	}
	return nil
}

func (e *containerdImpl) VerifyVersion(ctx context.Context) error {
	versionBytes, err := e.run(ctx, "version", "--format", "{{.Client.Version}}")
	if err != nil {
		if !e.useDirectNerdctl && runtime.GOOS == darwinOS {
			// On macOS using Lima
			return fmt.Errorf("failed to connect to Lima nerdctl: %w.\n"+
				"Try: limactl start", err)
		}

		// On Linux/Windows using direct nerdctl
		if strings.Contains(err.Error(), "executable file not found") {
			return fmt.Errorf("nerdctl not found. Please install nerdctl: %w", err)
		}
		return fmt.Errorf("failed to connect to nerdctl daemon: %w. Make sure containerd is running", err)
	}

	version, err := semver.NewVersion(strings.TrimSpace(string(versionBytes)))
	if err != nil {
		return errors.Join(ErrDeterminingContainerdVersion, err)
	}

	if version.Compare(minContainerdVersion) == -1 {
		_, _ = log.Warningf("Detected containerd version %s, the minimum supported containerd version is %s.\n", version.String(), minContainerdVersion.String())
	}

	return nil
}

func (e *containerdImpl) run(ctx context.Context, args ...string) ([]byte, error) {
	var cmd *exec.Cmd
	if e.useDirectNerdctl {
		cmd = exec.CommandContext(ctx, "nerdctl", args...)
	} else {
		cmd = exec.CommandContext(ctx, "lima", append([]string{"nerdctl"}, args...)...) //nolint:gosec // lima command with nerdctl args is expected
	}

	buf, err := cmd.Output()
	var exitErr *exec.ExitError
	if errors.As(err, &exitErr) {
		err = fmt.Errorf("%w: %s", exitErr, exitErr.Stderr)
	}
	return buf, err
}

func (e *containerdImpl) ContainerLogs(ctx context.Context, name string) ([]string, error) {
	buf, err := e.run(ctx, "container", "logs", name)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(buf), "\n"), nil
}

// findRandomUnusedPort finds a random unused port on the system.
// This mimics Docker's behavior of automatically allocating random ports
// when no host port is specified, which containerd/nerdctl doesn't do in rootless mode.
func findRandomUnusedPort() (int, error) {
	listener, err := net.Listen("tcp", ":0") //nolint:gosec // binding to all interfaces is intentional for random port allocation
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	addr := listener.Addr().(*net.TCPAddr)
	return addr.Port, nil
}

func containerdPortsFlags(flags *RunFlags) []string {
	if flags == nil {
		return nil
	}
	args := []string{}
	if flags.Ports != nil {
		for _, mapping := range flags.Ports {
			mapping.HostAddress = localhostIP
			if flags.BindIPAll != nil && *flags.BindIPAll {
				mapping.HostAddress = ""
			}

			// For rootless mode, always specify a host port
			// If no host port is specified, find a random unused port
			if mapping.HostPort == 0 {
				randomPort, err := findRandomUnusedPort()
				if err != nil {
					// Fallback to using the container port if we can't find a random port
					_, _ = log.Warningf("Failed to find random unused port, using container port %d: %v", mapping.ContainerPort, err)
					mapping.HostPort = mapping.ContainerPort
				} else {
					mapping.HostPort = randomPort
					_, _ = log.Debugf("Assigned random port %d for container port %d", randomPort, mapping.ContainerPort)
				}
			}

			args = append(args, "-p", portMappingFlag(mapping))
		}
	}
	return args
}

func containerdRunFlags(flags *RunFlags) []string {
	if flags == nil {
		return nil
	}
	args := containerdPortsFlags(flags)

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

func (e *containerdImpl) ContainerRun(ctx context.Context, image string, flags *RunFlags) (string, error) {
	args := []string{"run"}
	args = append(args, containerdRunFlags(flags)...)
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

func (e *containerdImpl) ContainerList(ctx context.Context, labels ...string) ([]Container, error) {
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

	list, err := parseNerdctlContainers(buf)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errParsingContainer, err)
	}
	return list, nil
}

// parseNerdctlContainers parses nerdctl's JSON output format which differs from Docker's format
func parseNerdctlContainers(data []byte) ([]Container, error) {
	// Pre-allocate slice based on estimated number of lines
	lines := strings.Count(strings.TrimSpace(string(data)), "\n") + 1
	containers := make([]Container, 0, lines)

	// Split by lines since nerdctl outputs one JSON object per line
	for line := range strings.SplitSeq(strings.TrimSpace(string(data)), "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}

		var c map[string]any
		if err := json.Unmarshal([]byte(line), &c); err != nil {
			return nil, fmt.Errorf("%w: %w", errDecodingJSON, err)
		}

		// Safely extract string fields with nil checks
		cont := Container{
			ID:    safeStringFromMap(c, "ID"),
			Names: []string{safeStringFromMap(c, "Names")},
			State: safeStringFromMap(c, "Status"), // nerdctl uses "Status" instead of "State"
			Image: safeStringFromMap(c, "Image"),
		}

		// Parse port mappings
		portsStr := safeStringFromMap(c, "Ports")
		if portsStr != "" {
			pm, err := parsePortMapping(portsStr)
			if err != nil {
				return nil, fmt.Errorf("%w: %w", errParsingPorts, err)
			}
			cont.Ports = pm
		}

		// Parse labels
		labelsStr := safeStringFromMap(c, "Labels")
		cont.Labels = map[string]string{}
		if labelsStr != "" {
			for label := range strings.SplitSeq(labelsStr, ",") {
				segments := strings.SplitN(label, "=", 2) //nolint:mnd // split label into key=value (2 parts)
				if len(segments) == 2 {                   //nolint:mnd // check for key=value pair (2 parts)
					cont.Labels[segments[0]] = segments[1]
				}
			}
		}

		containers = append(containers, cont)
	}

	return containers, nil
}

// safeStringFromMap safely extracts a string value from a map, returning empty string if nil or not a string
func safeStringFromMap(m map[string]any, key string) string {
	if val, exists := m[key]; exists && val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func (e *containerdImpl) ContainerRm(ctx context.Context, names ...string) error {
	args := []string{"container", "rm", "-v", "-f"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *containerdImpl) ContainerStart(ctx context.Context, names ...string) error {
	args := []string{"container", "start"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *containerdImpl) ContainerStop(ctx context.Context, names ...string) error {
	args := []string{"container", "stop"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *containerdImpl) ContainerUnpause(ctx context.Context, names ...string) error {
	args := []string{"container", "unpause"}
	args = append(args, names...)

	_, err := e.run(ctx, args...)
	return err
}

func (e *containerdImpl) ContainerInspect(ctx context.Context, names ...string) ([]*InspectData, error) {
	args := []string{"container", "inspect", "--format", "json"}
	args = append(args, names...)

	buf, err := e.run(ctx, args...)
	if err != nil {
		return nil, err
	}

	result := []*InspectData{}
	if err := json.Unmarshal(buf, &result); err != nil {
		// nerdctl returns single object instead of array, try unmarshaling as single object
		var singleResult *InspectData
		if singleErr := json.Unmarshal(buf, &singleResult); singleErr != nil {
			return nil, err // return original array error
		}
		result = []*InspectData{singleResult}
	}
	return result, nil
}

func (e *containerdImpl) ImageList(ctx context.Context, references ...string) ([]Image, error) {
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

func (e *containerdImpl) ImagePull(ctx context.Context, name string) error {
	_, err := e.run(ctx, "image", "pull", name)
	return err
}

func (e *containerdImpl) ImageHealthCheck(ctx context.Context, name string) (*ImageHealthCheck, error) {
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
		var singleOutput PartialImageInspect
		if err2 := json.Unmarshal(b, &singleOutput); err2 != nil {
			return nil, fmt.Errorf("%w: %w", errParseHealthCheck, err)
		}
		inspectOutput = []PartialImageInspect{singleOutput}
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

func (e *containerdImpl) ContainerHealthStatus(ctx context.Context, name string) (DockerHealthcheckStatus, error) { //nolint:gocyclo // complex health check logic is necessary
	buf, err := e.run(ctx, "inspect", "--format", "{{.State.Health.Status}}", name)
	if err != nil {
		if strings.Contains(err.Error(), "Health") || strings.Contains(err.Error(), "healthcheck") {
			return DockerHealthcheckStatusNone, nil
		}

		healthBuf, healthErr := e.run(ctx, "inspect", "--format", "{{if .State.Health}}{{.State.Health.Status}}{{else}}none{{end}}", name)
		if healthErr != nil {
			return "", err
		}

		statusString := strings.TrimSpace(string(healthBuf))
		if statusString == "none" || statusString == "" {
			return DockerHealthcheckStatusNone, nil
		}
		buf = healthBuf
	}

	statusString := strings.TrimSpace(string(buf))

	if statusString == "" {
		imageBuf, err := e.run(ctx, "inspect", "--format", "{{.Image}}", name)
		if err != nil {
			return DockerHealthcheckStatusNone, nil
		}
		imageName := strings.TrimSpace(string(imageBuf))

		imageHealthCheck, err := e.ImageHealthCheck(ctx, imageName)
		if err != nil {
			return DockerHealthcheckStatusNone, nil
		}

		if imageHealthCheck == nil {
			return DockerHealthcheckStatusNone, nil
		}

		startedAtBuf, err := e.run(ctx, "inspect", "--format", "{{.State.StartedAt}}", name)
		if err == nil {
			startedAtStr := strings.TrimSpace(string(startedAtBuf))
			if startedAt, err := time.Parse(time.RFC3339Nano, startedAtStr); err == nil {
				uptime := time.Since(startedAt)

				// Trigger healthcheck after 5 seconds to initialize it (same workaround as podman)
				if uptime > 5*time.Second {
					_, _ = e.run(ctx, "healthcheck", name)

					buf, err = e.run(ctx, "inspect", "--format", "{{.State.Health.Status}}", name)
					if err == nil {
						statusString = strings.TrimSpace(string(buf))
					}
				}
			}
		}

		if statusString == "" {
			return DockerHealthcheckStatusStarting, nil
		}
	}

	switch statusString {
	case string(DockerHealthcheckStatusNone):
		return DockerHealthcheckStatusNone, nil
	case string(DockerHealthcheckStatusStarting):
		return DockerHealthcheckStatusStarting, nil
	case string(DockerHealthcheckStatusHealthy):
		return DockerHealthcheckStatusHealthy, nil
	case string(DockerHealthcheckStatusUnhealthy):
		return DockerHealthcheckStatusUnhealthy, nil
	case "":
		return DockerHealthcheckStatusNone, nil
	default:
		return "", fmt.Errorf("unknown health status: %s", statusString)
	}
}

func (e *containerdImpl) Version(ctx context.Context) (map[string]any, error) {
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
