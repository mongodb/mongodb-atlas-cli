// Copyright 2024 MongoDB Inc
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

// file inspired by https://github.com/containers/podman/blob/v4.7.2/libpod/define/container_inspect.go

// InspectContainerConfig holds further data about how a container was initially
// configured.
type InspectContainerConfig struct {
	// Container labels
	Labels map[string]string `json:"Labels"`
}

// InspectHostPort provides information on a port on the host that a container's
// port is bound to.
type InspectHostPort struct {
	// IP on the host we are bound to. "" if not specified (binding to all
	// IPs).
	HostIP string `json:"HostIp"`
	// Port on the host we are bound to. No special formatting - just an
	// integer stuffed into a string.
	HostPort string `json:"HostPort"`
}

// InspectMount provides a record of a single mount in a container. It contains
// fields for both named and normal volumes. Only user-specified volumes will be
// included, and tmpfs volumes are not included even if the user specified them.
type InspectMount struct {
	// The name of the volume. Empty for bind mounts.
	Name string `json:"Name,omitempty"`
}

// InspectContainerHostConfig holds information used when the container was
// created.
// It's very much a Docker-specific struct, retained (mostly) as-is for
// compatibility. We fill individual fields as best as we can, inferring as much
// as possible from the spec and container config.
// Some things cannot be inferred. These will be populated by spec annotations
// (if available).
type InspectContainerHostConfig struct {
	// PortBindings contains the container's port bindings.
	// It is formatted as map[string][]InspectHostPort.
	// The string key here is formatted as <integer port number>/<protocol>
	// and represents the container port. A single container port may be
	// bound to multiple host ports (on different IPs).
	PortBindings map[string][]InspectHostPort `json:"PortBindings"`
}

// InspectBasicNetworkConfig holds basic configuration information (e.g. IP
// addresses, MAC address, subnet masks, etc) that are common for all networks
// (both additional and main).
type InspectBasicNetworkConfig struct {
	// IPAddress is the IP address for this network.
	IPAddress string `json:"IPAddress"`
}

// InspectAdditionalNetwork holds information about non-default networks the
// container has been connected to.
// As with InspectNetworkSettings, many fields are unused and maintained only
// for compatibility with Docker.
type InspectAdditionalNetwork struct {
	InspectBasicNetworkConfig
}

// InspectNetworkSettings holds information about the network settings of the
// container.
// Many fields are maintained only for compatibility with `docker inspect` and
// are unused within Libpod.
type InspectNetworkSettings struct {
	InspectBasicNetworkConfig

	// Networks contains information on non-default networks this
	// container has joined.
	// It is a map of network name to network information.
	Networks map[string]*InspectAdditionalNetwork `json:"Networks,omitempty"`
}

// InspectContainerData provides a detailed record of a container's configuration
// and state as viewed by Libpod.
// Large portions of this structure are defined such that the output is
// compatible with `docker inspect` JSON, but additional fields have been added
// as required to share information not in the original output.
type InspectContainerData struct {
	// Name is the name of the secret
	ID              string                      `json:"Id"`
	Name            string                      `json:"Name"`
	Mounts          []InspectMount              `json:"Mounts"`
	NetworkSettings *InspectNetworkSettings     `json:"NetworkSettings"`
	Config          *InspectContainerConfig     `json:"Config"`
	HostConfig      *InspectContainerHostConfig `json:"HostConfig"`
}
