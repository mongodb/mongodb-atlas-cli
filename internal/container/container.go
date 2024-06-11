package container

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
)

type Port struct {
	HostPort      int `json:"host_port"`
	ContainerPort int `json:"container_port"`
}

type ContainerRunFlags struct {
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

type Engine interface {
	ContainerLogs(context.Context, string) ([]string, error)
	ContainerRun(context.Context, string, *ContainerRunFlags) (string, error)
}

type Container struct {
	ID     string            `json:"ID"`
	Names  []string          `json:"Names"`
	State  string            `json:"State"`
	Image  string            `json:"Image"`
	Ports  []Port            `json:"Ports,omitempty"`
	Labels map[string]string `json:"Labels"`
}

func New() Engine {
	return newPodmanEngine()
}

type podmanImpl struct {
	client podman.Client
}

func newPodmanEngine() Engine {
	return &podmanImpl{
		client: podman.NewClient(),
	}
}

func (e *podmanImpl) ContainerLogs(ctx context.Context, name string) ([]string, error) {
	return e.client.ContainerLogs(ctx, name)
}

func (e *podmanImpl) ContainerRun(ctx context.Context, image string, flags *ContainerRunFlags) (string, error) {
	podmanOpts := podman.RunContainerOpts{
		Image: image,
	}
	if flags != nil {
		if flags.Detach != nil {
			podmanOpts.Detach = *flags.Detach
		}
		if flags.Remove != nil {
			podmanOpts.Remove = *flags.Remove
		}
		if flags.BindIPAll != nil {
			podmanOpts.BindIPAll = *flags.BindIPAll
		}
		if flags.Name != nil {
			podmanOpts.Name = *flags.Name
		}
		if flags.Hostname != nil {
			podmanOpts.Hostname = *flags.Hostname
		}
		if flags.Network != nil {
			podmanOpts.Network = *flags.Network
		}
		if flags.Entrypoint != nil {
			podmanOpts.Entrypoint = *flags.Entrypoint
		}
		if flags.Cmd != nil {
			podmanOpts.Cmd = *flags.Cmd
		}
		if flags.IP != nil {
			podmanOpts.IP = *flags.IP
		}
		for _, entry := range flags.Ports {
			podmanOpts.Ports[entry.HostPort] = entry.ContainerPort
		}
		podmanOpts.Volumes = flags.Volumes
		podmanOpts.Args = flags.Args
		podmanOpts.EnvVars = flags.Env
	}

	buf, err := e.client.RunContainer(ctx, podmanOpts)
	return string(buf), err
}
