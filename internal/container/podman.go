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
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/podman"
)

type podmanImpl struct {
	client podman.Client
}

func newPodmanEngine() Engine {
	return &podmanImpl{
		client: podman.NewClient(),
	}
}

func (e *podmanImpl) Ready(ctx context.Context) error {
	return e.client.Ready(ctx)
}

func (e *podmanImpl) ContainerLogs(ctx context.Context, name string) ([]string, error) {
	return e.client.ContainerLogs(ctx, name)
}

func (e *podmanImpl) ContainerRun(ctx context.Context, image string, flags *RunFlags) (string, error) {
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
		if flags.Ports != nil {
			podmanOpts.Ports = map[int]int{}
			for _, entry := range flags.Ports {
				podmanOpts.Ports[entry.HostPort] = entry.ContainerPort
			}
		}
		podmanOpts.Args = flags.Args
		podmanOpts.EnvVars = flags.Env
	}

	buf, err := e.client.RunContainer(ctx, podmanOpts)
	return string(buf), err
}

func (e *podmanImpl) ContainerList(ctx context.Context, nameFilter ...string) ([]Container, error) {
	containers, err := e.client.ListContainers(ctx, strings.Join(nameFilter, " "))
	if err != nil {
		return nil, err
	}
	result := make([]Container, 0, len(containers))
	for _, c := range containers {
		ports := make([]Port, 0, len(c.Ports))
		for _, p := range c.Ports {
			ports = append(ports, Port{
				ContainerPort: p.ContainerPort,
				HostPort:      p.HostPort,
			})
		}
		result = append(result, Container{
			ID:     c.ID,
			Image:  c.Image,
			Names:  c.Names,
			State:  c.State,
			Ports:  ports,
			Labels: c.Labels,
		})
	}
	return result, nil
}

func (e *podmanImpl) ImageList(ctx context.Context, nameFilter ...string) ([]Image, error) {
	images, err := e.client.ListImages(ctx, strings.Join(nameFilter, " "))
	if err != nil {
		return nil, err
	}
	result := make([]Image, 0, len(images))
	for _, c := range images {
		result = append(result, Image{
			ID:          c.ID,
			RepoTags:    c.RepoTags,
			RepoDigests: c.RepoDigests,
			Created:     c.Created,
			CreatedAt:   c.CreatedAt,
			Size:        c.Size,
			SharedSize:  c.SharedSize,
			VirtualSize: c.VirtualSize,
			Labels:      c.Labels,
			Containers:  c.Containers,
			Names:       c.Names,
		})
	}
	return result, nil
}

func (e *podmanImpl) ImagePull(ctx context.Context, name string) error {
	_, err := e.client.PullImage(ctx, name)
	return err
}

func (e *podmanImpl) ContainerRm(ctx context.Context, names ...string) error {
	_, err := e.client.RemoveContainers(ctx, names...)
	return err
}

func (e *podmanImpl) ContainerStart(ctx context.Context, names ...string) error {
	_, err := e.client.StartContainers(ctx, names...)
	return err
}

func (e *podmanImpl) ContainerStop(ctx context.Context, names ...string) error {
	_, err := e.client.StopContainers(ctx, names...)
	return err
}

func (e *podmanImpl) ContainerUnpause(ctx context.Context, names ...string) error {
	_, err := e.client.UnpauseContainers(ctx, names...)
	return err
}

func (e *podmanImpl) ContainerInspect(ctx context.Context, names ...string) ([]*InspectData, error) {
	res, err := e.client.ContainerInspect(ctx, names...)
	if err != nil {
		return nil, err
	}

	results := []*InspectData{}
	for _, data := range res {
		portBidings := map[string][]InspectDataHostPort{}

		for key, values := range data.HostConfig.PortBindings {
			for _, value := range values {
				portBidings[key] = append(portBidings[key], InspectDataHostPort{
					HostIP:   value.HostIP,
					HostPort: value.HostPort,
				})
			}
		}

		results = append(results, &InspectData{
			ID:   data.ID,
			Name: data.Name,
			Config: &InspectDataConfig{
				Labels: data.Config.Labels,
			},
			HostConfig: &InspectDataHostConfig{
				PortBindings: portBidings,
			},
		})
	}

	return results, nil
}

func (e *podmanImpl) Version(ctx context.Context) (map[string]any, error) {
	return e.client.Version(ctx)
}

func (*podmanImpl) Name() string {
	return "podman"
}
