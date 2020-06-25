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
	"strings"

	"go.mongodb.org/ops-manager/opsmngr"
)

const (
	mongod = "mongod"
)

// ProcessConfig that belongs to a cluster
type ProcessConfig struct {
	BuildIndexes *bool    `yaml:"buildIndexes,omitempty" json:"buildIndexes,omitempty"`
	DBPath       string   `yaml:"dbPath,omitempty" json:"dbPath,omitempty"`
	FCVersion    string   `yaml:"featureCompatibilityVersion,omitempty" json:"featureCompatibilityVersion,omitempty"`
	Hostname     string   `yaml:"hostname" json:"hostname"`
	LogPath      string   `yaml:"logPath" json:"logPath"`
	Name         string   `yaml:"name,omitempty" json:"name,omitempty"`
	Port         int      `yaml:"port" json:"port"`
	Priority     *float64 `yaml:"priority,omitempty" json:"priority,omitempty"`
	ProcessType  string   `yaml:"processType" json:"processType"`
	SlaveDelay   *float64 `yaml:"slaveDelay,omitempty" json:"slaveDelay,omitempty"`
	Version      string   `yaml:"version,omitempty" json:"version,omitempty"`
	Votes        *float64 `yaml:"votes,omitempty" json:"votes,omitempty"`
	ArbiterOnly  *bool    `yaml:"arbiterOnly,omitempty" json:"arbiterOnly,omitempty"`
	Disabled     bool     `yaml:"disabled" json:"disabled"`
	Hidden       *bool    `yaml:"hidden,omitempty" json:"hidden,omitempty"`
}

// setDefaults set default values based on the parent config
func (p *ProcessConfig) setDefaults(c *RSConfig) {
	if p.ProcessType == "" {
		p.ProcessType = mongod
	}
	if p.Version == "" {
		p.Version = c.Version
	}
	if p.FCVersion == "" {
		p.FCVersion = c.FCVersion
	}
	if p.BuildIndexes == nil {
		defaultValue := true
		p.BuildIndexes = &defaultValue
	}
}

// setProcessName reuse Name from an existing process
// this is based on hostname:port matching
func (p *ProcessConfig) setProcessName(processes []*opsmngr.Process, nameOpts ...string) {
	if p.Name != "" {
		return
	}

	for _, pp := range processes {
		if pp.Args26.NET.Port == p.Port && pp.Hostname == p.Hostname {
			p.Name = pp.Name
			return
		}
	}
	p.Name = strings.Join(nameOpts, "_")
}

// newReplicaSetProcessConfig maps opsmngr.member -> convert.ProcessConfig
func newReplicaSetProcessConfig(rs opsmngr.Member, p *opsmngr.Process) *ProcessConfig {
	return &ProcessConfig{
		BuildIndexes: &rs.BuildIndexes,
		Priority:     &rs.Priority,
		SlaveDelay:   &rs.SlaveDelay,
		Votes:        &rs.Votes,
		ArbiterOnly:  &rs.ArbiterOnly,
		Hidden:       &rs.Hidden,
		DBPath:       p.Args26.Storage.DBPath,
		LogPath:      p.Args26.SystemLog.Path,
		Port:         p.Args26.NET.Port,
		ProcessType:  p.ProcessType,
		Version:      p.Version,
		FCVersion:    p.FeatureCompatibilityVersion,
		Hostname:     p.Hostname,
		Name:         p.Name,
	}
}

// newReplicaSetProcessConfig maps opsmngr.Process -> convert.ProcessConfig
func newMongosProcessConfig(p *opsmngr.Process) *ProcessConfig {
	return &ProcessConfig{
		LogPath:     p.Args26.SystemLog.Path,
		Port:        p.Args26.NET.Port,
		ProcessType: p.ProcessType,
		Version:     p.Version,
		FCVersion:   p.FeatureCompatibilityVersion,
		Hostname:    p.Hostname,
		Name:        p.Name,
	}
}

// process maps convert.ProcessConfig -> opsmngr.Process
func (p *ProcessConfig) process() *opsmngr.Process {
	process := &opsmngr.Process{
		AuthSchemaVersion:           5,
		Disabled:                    p.Disabled,
		ManualMode:                  false,
		ProcessType:                 p.ProcessType,
		Version:                     p.Version,
		FeatureCompatibilityVersion: p.FCVersion,
		Hostname:                    p.Hostname,
		Name:                        p.Name,
	}
	return process
}

// newMongosProcess
func newMongosProcess(p *ProcessConfig, cluster string) *opsmngr.Process {
	process := p.process()
	process.Cluster = cluster
	process.Args26 = opsmngr.Args26{
		NET: opsmngr.Net{
			Port: p.Port,
		},
		SystemLog: opsmngr.SystemLog{
			Destination: file,
			Path:        p.LogPath,
		},
	}
	process.LogRotate = &opsmngr.LogRotate{
		SizeThresholdMB:  1000,
		TimeThresholdHrs: 24,
	}

	return process
}

func newReplicaSetProcess(p *ProcessConfig, replSetName string) *opsmngr.Process {
	process := p.process()

	process.Args26 = opsmngr.Args26{
		NET: opsmngr.Net{
			Port: p.Port,
		},
		Replication: &opsmngr.Replication{
			ReplSetName: replSetName,
		},
		Storage: &opsmngr.Storage{
			DBPath: p.DBPath,
		},
		SystemLog: opsmngr.SystemLog{
			Destination: file,
			Path:        p.LogPath,
		},
	}
	process.LogRotate = &opsmngr.LogRotate{
		SizeThresholdMB:  1000,
		TimeThresholdHrs: 24,
	}

	return process
}

func newConfigRSProcess(p *ProcessConfig, rsSetName string) *opsmngr.Process {
	process := p.process()

	process.Args26 = opsmngr.Args26{
		NET: opsmngr.Net{
			Port: p.Port,
		},
		Replication: &opsmngr.Replication{
			ReplSetName: rsSetName,
		},
		Storage: &opsmngr.Storage{
			DBPath: p.DBPath,
		},
		Sharding: &opsmngr.Sharding{ClusterRole: "configsvr"},
		SystemLog: opsmngr.SystemLog{
			Destination: file,
			Path:        p.LogPath,
		},
	}
	process.LogRotate = &opsmngr.LogRotate{
		SizeThresholdMB:  1000,
		TimeThresholdHrs: 24,
	}

	return process
}

func (p *ProcessConfig) member(i int) opsmngr.Member {
	m := opsmngr.Member{
		ID:           i,
		ArbiterOnly:  false,
		BuildIndexes: true,
		Hidden:       false,
		Host:         p.Name,
		Priority:     1,
		SlaveDelay:   0,
		Votes:        1,
	}
	if p.ArbiterOnly != nil {
		m.ArbiterOnly = *p.ArbiterOnly
	}
	if p.BuildIndexes != nil {
		m.BuildIndexes = *p.BuildIndexes
	}
	if p.Hidden != nil {
		m.Hidden = *p.Hidden
	}
	if p.Priority != nil {
		m.Priority = *p.Priority
	}
	if p.SlaveDelay != nil {
		m.SlaveDelay = *p.SlaveDelay
	}
	if p.Votes != nil {
		m.Votes = *p.Votes
	}
	return m
}
