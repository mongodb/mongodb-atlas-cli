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
	AuditLogPath        string   `yaml:"auditLogPath,omitempty" json:"auditLogPath,omitempty"`
	AuditLogDestination string   `yaml:"auditLogDestination,omitempty" json:"auditLogDestination,omitempty"`
	AuditLogFormat      string   `yaml:"auditLogFormat,omitempty" json:"auditLogFormat,omitempty"`
	BuildIndexes        *bool    `yaml:"buildIndexes,omitempty" json:"buildIndexes,omitempty"`
	DBPath              string   `yaml:"dbPath,omitempty" json:"dbPath,omitempty"`
	BindIP              *string  `yaml:"bindIp,omitempty" json:"bindIp,omitempty"`
	BindIPAll           *bool    `yaml:"bindIpAll,omitempty" json:"bindIpAll,omitempty"`
	FCVersion           string   `yaml:"featureCompatibilityVersion,omitempty" json:"featureCompatibilityVersion,omitempty"`
	Hostname            string   `yaml:"hostname" json:"hostname"`
	IPV6                *bool    `yaml:"ipv6,omitempty" json:"ipv6,omitempty"`
	LogPath             string   `yaml:"logPath" json:"logPath"`
	LogDestination      string   `yaml:"logDestination,omitempty" json:"logDestination,omitempty"`
	Name                string   `yaml:"name,omitempty" json:"name,omitempty"`
	Port                int      `yaml:"port" json:"port"`
	Priority            *float64 `yaml:"priority,omitempty" json:"priority,omitempty"`
	ProcessType         string   `yaml:"processType" json:"processType"`
	SlaveDelay          *float64 `yaml:"slaveDelay,omitempty" json:"slaveDelay,omitempty"`
	Version             string   `yaml:"version,omitempty" json:"version,omitempty"`
	Votes               *float64 `yaml:"votes,omitempty" json:"votes,omitempty"`
	ArbiterOnly         *bool    `yaml:"arbiterOnly,omitempty" json:"arbiterOnly,omitempty"`
	Disabled            bool     `yaml:"disabled" json:"disabled"`
	Hidden              *bool    `yaml:"hidden,omitempty" json:"hidden,omitempty"`
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

// newLogRotate default log rotation in LogRotate
func newLogRotate() *opsmngr.LogRotate {
	return &opsmngr.LogRotate{
		SizeThresholdMB:  1000,
		TimeThresholdHrs: 24,
	}
}

// newReplicaSetProcessConfig maps opsmngr.member -> convert.ProcessConfig
func newReplicaSetProcessConfig(rs opsmngr.Member, p *opsmngr.Process) *ProcessConfig {
	pc := &ProcessConfig{
		BuildIndexes:   &rs.BuildIndexes,
		Priority:       &rs.Priority,
		SlaveDelay:     &rs.SlaveDelay,
		Votes:          &rs.Votes,
		ArbiterOnly:    &rs.ArbiterOnly,
		Hidden:         &rs.Hidden,
		DBPath:         p.Args26.Storage.DBPath,
		LogPath:        p.Args26.SystemLog.Path,
		LogDestination: p.Args26.SystemLog.Destination,
		Port:           p.Args26.NET.Port,
		BindIP:         p.Args26.NET.BindIP,
		BindIPAll:      p.Args26.NET.BindIPAll,
		IPV6:           p.Args26.NET.IPV6,
		ProcessType:    p.ProcessType,
		Version:        p.Version,
		FCVersion:      p.FeatureCompatibilityVersion,
		Hostname:       p.Hostname,
		Name:           p.Name,
	}
	if p.Args26.AuditLog != nil {
		pc.AuditLogDestination = p.Args26.AuditLog.Destination
		pc.AuditLogFormat = p.Args26.AuditLog.Format
		pc.AuditLogPath = p.Args26.AuditLog.Path
	}
	return pc
}

// newReplicaSetProcessConfig maps opsmngr.Process -> convert.ProcessConfig
func newMongosProcessConfig(p *opsmngr.Process) *ProcessConfig {
	return &ProcessConfig{
		LogPath:        p.Args26.SystemLog.Path,
		LogDestination: p.Args26.SystemLog.Destination,
		Port:           p.Args26.NET.Port,
		ProcessType:    p.ProcessType,
		Version:        p.Version,
		FCVersion:      p.FeatureCompatibilityVersion,
		Hostname:       p.Hostname,
		Name:           p.Name,
	}
}

// newMongosProcess generates a mongo process for a mongos
func newMongosProcess(p *ProcessConfig, cluster string) *opsmngr.Process {
	process := p.process()
	process.Cluster = cluster
	process.Args26 = opsmngr.Args26{
		NET:       p.net(),
		SystemLog: p.systemLog(),
	}
	process.LogRotate = newLogRotate()
	if p.AuditLogPath != "" {
		process.Args26.AuditLog = p.auditLog()
	}
	return process
}

// newReplicaSetProcess generates a mongo process for a replica set mongod
func newReplicaSetProcess(p *ProcessConfig, replSetName string) *opsmngr.Process {
	process := p.process()

	process.Args26 = opsmngr.Args26{
		NET: p.net(),
		Replication: &opsmngr.Replication{
			ReplSetName: replSetName,
		},
		Storage: &opsmngr.Storage{
			DBPath: p.DBPath,
		},
		SystemLog: p.systemLog(),
	}
	process.LogRotate = newLogRotate()
	if p.AuditLogPath != "" {
		process.Args26.AuditLog = p.auditLog()
	}
	return process
}

// newConfigRSProcess generates a mongo process for a replica set config server
func newConfigRSProcess(p *ProcessConfig, rsSetName string) *opsmngr.Process {
	process := p.process()

	process.Args26 = opsmngr.Args26{
		NET: p.net(),
		Replication: &opsmngr.Replication{
			ReplSetName: rsSetName,
		},
		Storage: &opsmngr.Storage{
			DBPath: p.DBPath,
		},
		Sharding:  &opsmngr.Sharding{ClusterRole: "configsvr"},
		SystemLog: p.systemLog(),
	}
	process.LogRotate = newLogRotate()
	if p.AuditLogPath != "" {
		process.Args26.AuditLog = p.auditLog()
	}

	return process
}

// systemLog maps convert.ProcessConfig -> opsmngr.Net
func (p *ProcessConfig) net() opsmngr.Net {
	return opsmngr.Net{
		Port:      p.Port,
		BindIP:    p.BindIP,
		BindIPAll: p.BindIPAll,
	}
}

// systemLog maps convert.ProcessConfig -> opsmngr.SystemLog
func (p *ProcessConfig) systemLog() opsmngr.SystemLog {
	return opsmngr.SystemLog{
		Destination: p.systemLogDestination(),
		Path:        p.LogPath,
	}
}

func (p *ProcessConfig) systemLogDestination() string {
	if p.LogDestination != "" {
		return p.LogDestination
	}
	return file
}

// auditLog maps convert.ProcessConfig -> opsmngr.AuditLog
func (p *ProcessConfig) auditLog() *opsmngr.AuditLog {
	return &opsmngr.AuditLog{
		Destination: p.AuditLogDestination,
		Path:        p.AuditLogPath,
		Format:      p.AuditLogFormat,
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

// member maps convert.ProcessConfig -> opsmngr.Member
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
