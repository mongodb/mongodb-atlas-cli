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
	"fmt"

	"go.mongodb.org/ops-manager/opsmngr"
)

// ProcessConfig that belongs to a cluster
type ProcessConfig struct {
	BuildIndexes *bool   `yaml:"buildIndexes,omitempty" json:"buildIndexes,omitempty"`
	DBPath       string  `yaml:"dbPath" json:"dbPath"`
	FCVersion    string  `yaml:"featureCompatibilityVersion,omitempty" json:"featureCompatibilityVersion,omitempty"`
	Hostname     string  `yaml:"hostname" json:"hostname"`
	LogPath      string  `yaml:"logPath" json:"logPath"`
	Name         string  `yaml:"name,omitempty" json:"name,omitempty"`
	Port         int     `yaml:"port" json:"port"`
	Priority     float64 `yaml:"priority" json:"priority"`
	ProcessType  string  `yaml:"processType" json:"processType"`
	SlaveDelay   float64 `yaml:"slaveDelay" json:"slaveDelay"`
	Version      string  `yaml:"version,omitempty" json:"version,omitempty"`
	Votes        float64 `yaml:"votes" json:"votes"`
	ArbiterOnly  bool    `yaml:"arbiterOnly" json:"arbiterOnly"`
	Disabled     bool    `yaml:"disabled" json:"disabled"`
	Hidden       bool    `yaml:"hidden" json:"hidden"`
}

// setDefaults set default values based on the parent config
func (p *ProcessConfig) setDefaults(c *ClusterConfig) {
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
func (p *ProcessConfig) setProcessName(clusterName string, processes []*opsmngr.Process, i int) {
	if p.Name != "" {
		return
	}

	p.Name = fmt.Sprintf("%s_%d", clusterName, len(processes)+i)
	for _, pp := range processes {
		if pp.Args26.NET.Port == p.Port && pp.Hostname == p.Hostname {
			p.Name = pp.Name
			return
		}
	}
}

func (p *ProcessConfig) toCMProcess(replSetName string) *opsmngr.Process {
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

func (p *ProcessConfig) toCMMember(i int) opsmngr.Member {
	return opsmngr.Member{
		ID:           i,
		ArbiterOnly:  p.ArbiterOnly,
		BuildIndexes: *p.BuildIndexes,
		Hidden:       p.Hidden,
		Host:         p.Name,
		Priority:     p.Priority,
		SlaveDelay:   p.SlaveDelay,
		Votes:        p.Votes,
	}
}
