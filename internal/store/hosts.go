// Copyright 2020 MongoDB Inc
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

package store

import (
	"context"
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	"go.mongodb.org/ops-manager/opsmngr"
)

type HostLister interface {
	Hosts(string, *opsmngr.HostListOptions) (*opsmngr.Hosts, error)
}

type HostDescriber interface {
	Host(string, string) (*opsmngr.Host, error)
}

// Hosts encapsulate the logic to manage different cloud providers
func (s *Store) Hosts(groupID string, opts *opsmngr.HostListOptions) (*opsmngr.Hosts, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Hosts.List(context.Background(), groupID, opts)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Host encapsulate the logic to manage different cloud providers
func (s *Store) Host(groupID, hostID string) (*opsmngr.Host, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*opsmngr.Client).Hosts.Get(context.Background(), groupID, hostID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
