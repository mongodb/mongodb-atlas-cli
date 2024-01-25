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

import "time"

// file partially copied from https://github.com/containers/podman/blob/v4.7.2/pkg/machine/config.go

type InspectInfo struct {
	ConfigPath         VMFile
	ConnectionInfo     ConnectionConfig
	Created            time.Time
	Image              ImageConfig
	LastUp             time.Time
	Name               string
	Resources          ResourceConfig
	SSHConfig          SSHConfig
	State              Status
	UserModeNetworking bool
	Rootful            bool
}

type VMFile struct {
	// Path is the fully qualified path to a file
	Path string
	// Symlink is a shortened version of Path by using
	// a symlink
	Symlink *string `json:"symlink,omitempty"`
}

// ConnectionConfig contains connections like sockets, etc.
type ConnectionConfig struct {
	// PodmanSocket is the exported podman service socket
	PodmanSocket *VMFile `json:"PodmanSocket"`
	// PodmanPipe is the exported podman service named pipe (Windows hosts only)
	PodmanPipe *VMFile `json:"PodmanPipe"`
}

// ImageConfig describes the bootable image for the VM.
type ImageConfig struct {
	// IgnitionFile is the path to the filesystem where the
	// ignition file was written (if needs one)
	IgnitionFile VMFile `json:"IgnitionFilePath"`
	// ImageStream is the update stream for the image
	ImageStream string
	// ImageFile is the fq path to
	ImagePath VMFile `json:"ImagePath"`
}

// ResourceConfig describes physical attributes of the machine.
type ResourceConfig struct {
	// CPUs to be assigned to the VM
	CPUs uint64
	// Disk size in gigabytes assigned to the vm
	DiskSize uint64
	// Memory in megabytes assigned to the vm
	Memory uint64
}

// SSHConfig contains remote access information for SSH.
type SSHConfig struct {
	// IdentityPath is the fq path to the ssh priv key
	IdentityPath string
	// SSH port for user networking
	Port int
	// RemoteUsername of the vm user
	RemoteUsername string
}

type Status = string

const (
	Running            Status = "running"
	DefaultMachineName string = "podman-machine-default"
)
