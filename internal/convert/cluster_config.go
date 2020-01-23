// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package convert

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/10gen/mcli/internal/utils"
	"github.com/Masterminds/semver"
	"github.com/mongodb-labs/pcgc/cloudmanager"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
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

var supportedExts = []string{"json", "yaml", "yml"}

// NewClusterConfigFromFile load a ClusterConfig from a YAML or JSON file
func NewClusterConfigFromFile(fs afero.Fs, filename string) (*ClusterConfig, error) {
	if exists, err := afero.Exists(fs, filename); !exists || err != nil {
		return nil, fmt.Errorf("file not found: %s", filename)
	}

	ext := filepath.Ext(filename)
	if len(ext) <= 1 {
		return nil, fmt.Errorf("filename: %s requires valid extension", filename)
	}
	configType := ext[1:]
	if !utils.StringInSlice(supportedExts, configType) {
		return nil, fmt.Errorf("unsupported file type: %s", configType)
	}

	file, err := afero.ReadFile(fs, filename)
	if err != nil {
		return nil, err
	}

	config := new(ClusterConfig)
	switch configType {
	case "yaml", "yml":
		if err := yaml.Unmarshal(file, config); err != nil {
			return nil, err
		}
	case "json":
		if err := json.Unmarshal(file, config); err != nil {
			return nil, err
		}
	}

	return config, nil
}

// PatchAutomationConfig add the ClusterConfig to a cloudmanager.AutomationConfig
// this method will modify the given AutomationConfig to add the new replica set information
func (c *ClusterConfig) PatchAutomationConfig(out *cloudmanager.AutomationConfig) error {
	newProcesses := make([]*cloudmanager.Process, len(c.ProcessConfigs))

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
func (c *ClusterConfig) toReplicaSet() (*cloudmanager.ReplicaSet, error) {
	protocolVer, err := protocolVer(c.Version)
	if err != nil {
		return nil, err
	}

	rs := &cloudmanager.ReplicaSet{
		ID:              c.Name,
		Members:         make([]cloudmanager.Member, len(c.ProcessConfigs)),
		ProtocolVersion: protocolVer,
	}

	return rs, nil
}

// patchProcesses replace replica set processes with new configuration
// this will disable all existing processes for the given replica set and remove the association
// Then try to patch then with the new config if one config exists for the same host:port
func patchProcesses(out *cloudmanager.AutomationConfig, newReplicaSetID string, newProcesses []*cloudmanager.Process) {
	for i, oldProcess := range out.Processes {
		if oldProcess.Args26.Replication != nil && oldProcess.Args26.Replication.ReplSetName == newReplicaSetID {
			oldProcess.Disabled = true
			oldProcess.Args26.Replication = new(cloudmanager.Replication)
		}
		pos := SearchProcesses(newProcesses, func(p *cloudmanager.Process) bool {
			return p.Name == oldProcess.Name
		})
		if pos != -1 {
			out.Processes[i] = newProcesses[pos]
			newProcesses = append(newProcesses[:pos], newProcesses[pos+1:]...)
		}
	}
	if len(newProcesses) > 0 {
		out.Processes = append(out.Processes, newProcesses...)
	}
}

// patchReplicaSet if the replica set exists try to patch it if not add it
func patchReplicaSet(out *cloudmanager.AutomationConfig, newReplicaSet *cloudmanager.ReplicaSet) {
	pos := SearchReplicaSets(out.ReplicaSets, func(r *cloudmanager.ReplicaSet) bool {
		return r.ID == newReplicaSet.ID
	})

	if pos == -1 {
		out.ReplicaSets = append(out.ReplicaSets, newReplicaSet)
		return
	}

	oldReplicaSet := out.ReplicaSets[pos]
	lastID := oldReplicaSet.Members[len(oldReplicaSet.Members)-1].ID
	for j, newMember := range newReplicaSet.Members {
		k := SearchMembers(oldReplicaSet.Members, func(m cloudmanager.Member) bool {
			return m.Host == newMember.Host
		})
		if k != -1 {
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
