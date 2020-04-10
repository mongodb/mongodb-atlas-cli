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

package convert

import (
	"github.com/Masterminds/semver"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/go-client-mongodb-ops-manager/search"
)

const (
	zero            = "0"
	one             = "1"
	file            = "file"
	fcvLessThanFour = "< 4.0"
)

// ClusterConfig configuration for a cluster
// This cluster can be used to patch an automation config
type ClusterConfig struct {
	FCVersion      string          `yaml:"featureCompatibilityVersion,omitempty" json:"featureCompatibilityVersion,omitempty"`
	MongoURI       string          `yaml:"mongoURI,omitempty" json:"mongoURI,omitempty"`
	Name           string          `yaml:"name" json:"name"`
	ProcessConfigs []ProcessConfig `yaml:"processes" json:"processes"`
	Version        string          `yaml:"version,omitempty" json:"version,omitempty"`
}

// PatchAutomationConfig add the ClusterConfig to a cloudmanager.AutomationConfig
// this method will modify the given AutomationConfig to add the new replica set information
func (c *ClusterConfig) PatchAutomationConfig(out *om.AutomationConfig) error {
	newProcesses := make([]*om.Process, len(c.ProcessConfigs))

	newReplicaSet, err := c.toReplicaSet()
	if err != nil {
		return err
	}

	// transform cli config to automation config
	for i, pc := range c.ProcessConfigs {
		pc.setDefaults(c)
		pc.setProcessName(c.Name, out.Processes, i)
		newProcesses[i] = pc.toCMProcess(c.Name)
		newReplicaSet.Members[i] = pc.toCMMember(i)
	}

	// This value may not be present and is mandatory
	if out.Auth.DeploymentAuthMechanisms == nil {
		out.Auth.DeploymentAuthMechanisms = make([]string, 0)
	}

	patchProcesses(out, newReplicaSet.ID, newProcesses)
	patchReplicaSet(out, newReplicaSet)

	return nil
}

// toReplicaSet convert from cli config to cloudmanager.ReplicaSet
func (c *ClusterConfig) toReplicaSet() (*om.ReplicaSet, error) {
	protocolVer, err := protocolVer(c.FCVersion)
	if err != nil {
		return nil, err
	}

	rs := &om.ReplicaSet{
		ID:              c.Name,
		Members:         make([]om.Member, len(c.ProcessConfigs)),
		ProtocolVersion: protocolVer,
	}

	return rs, nil
}

// patchProcesses replace replica set processes with new configuration
// this will disable all existing processes for the given replica set and remove the association
// Then try to patch then with the new config if one config exists for the same host:port
func patchProcesses(out *om.AutomationConfig, newReplicaSetID string, newProcesses []*om.Process) {
	for i, oldProcess := range out.Processes {
		if oldProcess.Args26.Replication != nil && oldProcess.Args26.Replication.ReplSetName == newReplicaSetID {
			oldProcess.Disabled = true
			oldProcess.Args26.Replication = new(om.Replication)
		}
		pos, found := search.Processes(newProcesses, func(p *om.Process) bool {
			return p.Name == oldProcess.Name
		})
		if found {
			out.Processes[i] = newProcesses[pos]
			newProcesses = append(newProcesses[:pos], newProcesses[pos+1:]...)
		}
	}
	if len(newProcesses) > 0 {
		out.Processes = append(out.Processes, newProcesses...)
	}
}

// patchReplicaSet if the replica set exists try to patch it if not add it
func patchReplicaSet(out *om.AutomationConfig, newReplicaSet *om.ReplicaSet) {
	pos, found := search.ReplicaSets(out.ReplicaSets, func(r *om.ReplicaSet) bool {
		return r.ID == newReplicaSet.ID
	})

	if !found {
		out.ReplicaSets = append(out.ReplicaSets, newReplicaSet)
		return
	}

	oldReplicaSet := out.ReplicaSets[pos]
	lastID := oldReplicaSet.Members[len(oldReplicaSet.Members)-1].ID
	for j, newMember := range newReplicaSet.Members {
		k, found := search.Members(oldReplicaSet.Members, func(m om.Member) bool {
			return m.Host == newMember.Host
		})
		if found {
			newMember.ID = oldReplicaSet.Members[k].ID
		} else {
			lastID++
			newMember.ID = lastID
		}
		newReplicaSet.Members[j] = newMember
	}
	out.ReplicaSets[pos] = newReplicaSet
}

// protocolVer determines the appropriate protocol based on FCV
// return "0" for versions <4.0 or "1" otherwise
func protocolVer(version string) (string, error) {
	ver, err := semver.NewVersion(version)
	if err != nil {
		return "", err
	}
	constrain, _ := semver.NewConstraint(fcvLessThanFour)

	if constrain.Check(ver) {
		return zero, nil
	}
	return one, nil
}
