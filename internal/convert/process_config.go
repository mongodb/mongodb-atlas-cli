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
	"fmt"

	"github.com/mongodb-labs/pcgc/cloudmanager"
)

// ProcessConfig that belongs to a cluster
type ProcessConfig struct {
	BuildIndexes *bool   `yaml:"buildIndexes,omitempty" json:"buildIndexes,omitempty"`
	DBPath       string  `yaml:"dbPath" json:"db_path"`
	FCVersion    string  `yaml:"featureCompatibilityVersion,omitempty" json:"featureCompatibilityVersion,omitempty"`
	Hostname     string  `yaml:"hostname" json:"hostname"`
	LogPath      string  `yaml:"logPath" json:"log_path"`
	Name         string  `yaml:"name,omitempty" json:"name,omitempty"`
	Port         int     `yaml:"port" json:"port"`
	Priority     float64 `yaml:"priority" json:"priority"`
	ProcessType  string  `yaml:"processType" json:"process_type"`
	SlaveDelay   float64 `yaml:"slaveDelay" json:"slave_delay"`
	Version      string  `yaml:"version,omitempty" json:"version,omitempty"`
	Votes        float64 `yaml:"votes" json:"votes"`
	ArbiterOnly  bool    `yaml:"arbiterOnly" json:"arbiter_only"`
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
func (p *ProcessConfig) setProcessName(clusterName string, processes []*cloudmanager.Process, i int) {
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

func (p *ProcessConfig) toCMProcess(replSetName string) *cloudmanager.Process {
	process := &cloudmanager.Process{
		AuthSchemaVersion:           5,
		Disabled:                    p.Disabled,
		ManualMode:                  false,
		ProcessType:                 p.ProcessType,
		Version:                     p.Version,
		FeatureCompatibilityVersion: p.FCVersion,
		Hostname:                    p.Hostname,
		Name:                        p.Name,
	}

	process.Args26 = cloudmanager.Args26{
		NET: cloudmanager.Net{
			Port: p.Port,
		},
		Replication: &cloudmanager.Replication{
			ReplSetName: replSetName,
		},
		Storage: &cloudmanager.Storage{
			DBPath: p.DBPath,
		},
		SystemLog: cloudmanager.SystemLog{
			Destination: file,
			Path:        p.LogPath,
		},
	}
	process.LogRotate = &cloudmanager.LogRotate{
		SizeThresholdMB:  1000,
		TimeThresholdHrs: 24,
	}

	return process
}

func (p *ProcessConfig) toCMMember(i int) cloudmanager.Member {
	return cloudmanager.Member{
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
