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

import (
	"net"
	"time"
)

// file partially inspired by from https://github.com/containers/common/blob/v0.56.0/libnetwork/types/network.go

// Network describes the Network attributes.
type Network struct {
	// Name of the Network.
	Name string `json:"name"`
	// ID of the Network.
	ID string `json:"id"`
	// Driver for this Network, e.g. bridge, macvlan...
	Driver string `json:"driver"`
	// NetworkInterface is the network interface name on the host.
	NetworkInterface string `json:"network_interface,omitempty"`
	// Created contains the timestamp when this network was created.
	Created time.Time `json:"created,omitempty"`
	// Subnets to use for this network.
	Subnets []Subnet `json:"subnets,omitempty"`
	// Routes to use for this network.
	Routes []Route `json:"routes,omitempty"`
	// IPv6Enabled if set to true an ipv6 subnet should be created for this net.
	IPv6Enabled bool `json:"ipv6_enabled"`
	// Internal is whether the Network should not have external routes
	// to public or other Networks.
	Internal bool `json:"internal"`
	// DNSEnabled is whether name resolution is active for container on
	// this Network. Only supported with the bridge driver.
	DNSEnabled bool `json:"dns_enabled"`
	// List of custom DNS server for podman's DNS resolver at network level,
	// all the containers attached to this network will consider resolvers
	// configured at network level.
	NetworkDNSServers []string `json:"network_dns_servers,omitempty"`
	// Labels is a set of key-value labels that have been applied to the
	// Network.
	Labels map[string]string `json:"labels,omitempty"`
	// Options is a set of key-value options that have been applied to
	// the Network.
	Options map[string]string `json:"options,omitempty"`
	// IPAMOptions contains options used for the ip assignment.
	IPAMOptions map[string]string `json:"ipam_options,omitempty"`
}

type Subnet struct {
	// Subnet for this Network in CIDR form.
	// swagger:strfmt string
	Subnet IPNet `json:"subnet"`
	// Gateway IP for this Network.
	// swagger:strfmt string
	Gateway net.IP `json:"gateway,omitempty"`
	// LeaseRange contains the range where IP are leased. Optional.
	LeaseRange *LeaseRange `json:"lease_range,omitempty"`
}

// IPNet is used as custom net.IPNet type to add Marshal/Unmarshal methods.
type IPNet struct {
	net.IPNet
}

func (n *IPNet) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

func (n *IPNet) UnmarshalText(text []byte) error {
	subnet, err := ParseCIDR(string(text))
	if err != nil {
		return err
	}
	*n = subnet
	return nil
}

// ParseCIDR parse a string to IPNet.
func ParseCIDR(cidr string) (IPNet, error) {
	ip, subnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return IPNet{}, err
	}
	// convert to 4 bytes if ipv4
	ipv4 := ip.To4()
	if ipv4 != nil {
		ip = ipv4
	}
	subnet.IP = ip
	return IPNet{*subnet}, err
}

type Route struct {
	// Destination for this route in CIDR form.
	// swagger:strfmt string
	Destination IPNet `json:"destination"`
	// Gateway IP for this route.
	// swagger:strfmt string
	Gateway net.IP `json:"gateway"`
	// Metric for this route. Optional.
	Metric *uint32 `json:"metric,omitempty"`
}

// LeaseRange contains the range where IP are leased.
type LeaseRange struct {
	// StartIP first IP in the subnet which should be used to assign ips.
	// swagger:strfmt string
	StartIP net.IP `json:"start_ip,omitempty"`
	// EndIP last IP in the subnet which should be used to assign ips.
	// swagger:strfmt string
	EndIP net.IP `json:"end_ip,omitempty"`
}
