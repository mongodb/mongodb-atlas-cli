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

	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/config"
)

type AgentLister interface {
	Agents(string, string) (*om.Agents, error)
}

type AgentUpgrader interface {
	UpgradeAgent(string) (*om.AutomationConfigAgent, error)
}

// Agents encapsulate the logic to manage different cloud providers
func (s *Store) Agents(projectID, agentType string) (*om.Agents, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*om.Client).Agents.List(context.Background(), projectID, agentType)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// Agents encapsulate the logic to manage different cloud providers
func (s *Store) UpgradeAgent(projectID string) (*om.AutomationConfigAgent, error) {
	switch s.service {
	case config.OpsManagerService, config.CloudManagerService:
		result, _, err := s.client.(*om.Client).AutomationConfig.UpdateAgent(context.Background(), projectID)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}
