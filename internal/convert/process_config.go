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
	mongod                  = "mongod"
	authSchemaVersion       = 5
	defaultSizeThresholdMB  = 1000
	defaultTimeThresholdHrs = 24
)

// ProcessConfig that belongs to a cluster
type ProcessConfig struct {
	AuditLogPath           string                  `yaml:"auditLogPath,omitempty" json:"auditLogPath,omitempty"`
	AuditLogDestination    string                  `yaml:"auditLogDestination,omitempty" json:"auditLogDestination,omitempty"`
	AuditLogFormat         string                  `yaml:"auditLogFormat,omitempty" json:"auditLogFormat,omitempty"`
	AuditLogFilter         string                  `yaml:"auditLogFilter,omitempty" json:"auditLogFilter,omitempty"`
	BuildIndexes           *bool                   `yaml:"buildIndexes,omitempty" json:"buildIndexes,omitempty"`
	DBPath                 string                  `yaml:"dbPath,omitempty" json:"dbPath,omitempty"`
	BindIP                 *string                 `yaml:"bindIp,omitempty" json:"bindIp,omitempty"`
	BindIPAll              *bool                   `yaml:"bindIpAll,omitempty" json:"bindIpAll,omitempty"`
	DirectoryPerDB         *bool                   `yaml:"directoryPerDB,omitempty" json:"directoryPerDB,omitempty"`
	Engine                 string                  `yaml:"engine,omitempty" json:"engine,omitempty"`
	FCVersion              string                  `yaml:"featureCompatibilityVersion,omitempty" json:"featureCompatibilityVersion,omitempty"`
	Hostname               string                  `yaml:"hostname" json:"hostname"`
	InMemory               *map[string]interface{} `yaml:"inMemory,omitempty" json:"inMemory,omitempty"`
	IndexBuildRetry        *bool                   `yaml:"indexBuildRetry,omitempty" json:"indexBuildRetry,omitempty"`
	IPV6                   *bool                   `yaml:"ipv6,omitempty" json:"ipv6,omitempty"`
	Journal                *map[string]interface{} `yaml:"journal,omitempty" json:"journal,omitempty"`
	LogAppend              bool                    `yaml:"logAppend,omitempty" json:"logAppend,omitempty"`
	LogDestination         string                  `yaml:"logDestination,omitempty" json:"logDestination,omitempty"`
	LogPath                string                  `yaml:"logPath" json:"logPath"`
	LogRotate              string                  `yaml:"logRotate,omitempty" json:"logRotate,omitempty"`
	LogVerbosity           int                     `yaml:"logVerbosity,omitempty" json:"logVerbosity,omitempty"`
	Name                   string                  `yaml:"name,omitempty" json:"name,omitempty"`
	OplogMinRetentionHours *float64                `yaml:"oplogMinRetentionHours,omitempty" json:"oplogMinRetentionHours,omitempty"`
	Port                   int                     `yaml:"port" json:"port"`
	Priority               *float64                `yaml:"priority,omitempty" json:"priority,omitempty"`
	ProcessType            string                  `yaml:"processType" json:"processType"`
	SlaveDelay             *float64                `yaml:"slaveDelay,omitempty" json:"slaveDelay,omitempty"`
	SyncPeriodSecs         *float64                `yaml:"syncPeriodSecs,omitempty" json:"syncPeriodSecs,omitempty"`
	Votes                  *float64                `yaml:"votes,omitempty" json:"votes,omitempty"`
	ArbiterOnly            *bool                   `yaml:"arbiterOnly,omitempty" json:"arbiterOnly,omitempty"`
	Disabled               bool                    `yaml:"disabled" json:"disabled"`
	Hidden                 *bool                   `yaml:"hidden,omitempty" json:"hidden,omitempty"`
	Security               *map[string]interface{} `yaml:"security,omitempty" json:"security,omitempty"`
	TLS                    *TLS                    `yaml:"tls,omitempty" json:"tls,omitempty"`
	Version                string                  `yaml:"version,omitempty" json:"version,omitempty"`
	WiredTiger             *map[string]interface{} `yaml:"wiredTiger,omitempty" json:"wiredTiger,omitempty"`
}

// TLS defines TLS parameters for Net
type TLS struct {
	CAFile                     string `yaml:"CAFile,omitempty" json:"CAFile,omitempty"`
	CertificateKeyFile         string `yaml:"certificateKeyFile,omitempty" json:"certificateKeyFile,omitempty"`
	CertificateKeyFilePassword string `yaml:"certificateKeyFilePassword,omitempty" json:"certificateKeyFilePassword,omitempty"`
	CertificateSelector        string `yaml:"certificateSelector,omitempty" json:"certificateSelector,omitempty"`
	ClusterCertificateSelector string `yaml:"clusterCertificateSelector,omitempty" json:"clusterCertificateSelector,omitempty"`
	ClusterFile                string `yaml:"clusterFile,omitempty" json:"clusterFile,omitempty"`
	ClusterPassword            string `yaml:"clusterPassword,omitempty" json:"clusterPassword,omitempty"`
	CRLFile                    string `yaml:"CRLFile,omitempty" json:"CRLFile,omitempty"`
	DisabledProtocols          string `yaml:"disabledProtocols,omitempty" json:"disabledProtocols,omitempty"`
	FIPSMode                   string `yaml:"FIPSMode,omitempty" json:"FIPSMode,omitempty"`
	Mode                       string `yaml:"mode,omitempty" json:"mode,omitempty"`
	PEMKeyFile                 string `yaml:"PEMKeyFile,omitempty" json:"PEMKeyFile,omitempty"`
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
		SizeThresholdMB:  defaultSizeThresholdMB,
		TimeThresholdHrs: defaultTimeThresholdHrs,
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
		LogPath:        p.Args26.SystemLog.Path,
		LogDestination: p.Args26.SystemLog.Destination,
		LogAppend:      p.Args26.SystemLog.LogAppend,
		LogVerbosity:   p.Args26.SystemLog.Verbosity,
		LogRotate:      p.Args26.SystemLog.LogRotate,
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

	if p.Args26.Storage != nil {
		pc.DBPath = p.Args26.Storage.DBPath
		pc.DirectoryPerDB = p.Args26.Storage.DirectoryPerDB
		pc.SyncPeriodSecs = p.Args26.Storage.SyncPeriodSecs
		pc.Engine = p.Args26.Storage.Engine
		pc.WiredTiger = p.Args26.Storage.WiredTiger
		pc.OplogMinRetentionHours = p.Args26.Storage.OplogMinRetentionHours
		pc.Journal = p.Args26.Storage.Journal
	}

	if p.Args26.AuditLog != nil {
		pc.AuditLogDestination = p.Args26.AuditLog.Destination
		pc.AuditLogFormat = p.Args26.AuditLog.Format
		pc.AuditLogPath = p.Args26.AuditLog.Path
		pc.AuditLogFilter = p.Args26.AuditLog.Filter
	}

	if p.Args26.NET.TLS != nil {
		pc.TLS = &TLS{
			CAFile:                     p.Args26.NET.TLS.CAFile,
			CertificateKeyFile:         p.Args26.NET.TLS.CertificateKeyFile,
			CertificateKeyFilePassword: p.Args26.NET.TLS.CertificateKeyFilePassword,
			CertificateSelector:        p.Args26.NET.TLS.CertificateSelector,
			ClusterCertificateSelector: p.Args26.NET.TLS.ClusterCertificateSelector,
			ClusterFile:                p.Args26.NET.TLS.ClusterFile,
			ClusterPassword:            p.Args26.NET.TLS.ClusterPassword,
			CRLFile:                    p.Args26.NET.TLS.CRLFile,
			DisabledProtocols:          p.Args26.NET.TLS.DisabledProtocols,
			FIPSMode:                   p.Args26.NET.TLS.FIPSMode,
			Mode:                       p.Args26.NET.TLS.Mode,
			PEMKeyFile:                 p.Args26.NET.TLS.PEMKeyFile,
		}
	}
	if p.Args26.Security != nil {
		pc.Security = p.Args26.Security
	}
	return pc
}

// newMongosProcessConfig maps opsmngr.Process -> convert.ProcessConfig
func newMongosProcessConfig(p *opsmngr.Process) *ProcessConfig {
	pc := &ProcessConfig{
		LogPath:        p.Args26.SystemLog.Path,
		LogDestination: p.Args26.SystemLog.Destination,
		LogAppend:      p.Args26.SystemLog.LogAppend,
		LogVerbosity:   p.Args26.SystemLog.Verbosity,
		LogRotate:      p.Args26.SystemLog.LogRotate,
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
		pc.AuditLogFilter = p.Args26.AuditLog.Filter
	}
	if p.Args26.NET.TLS != nil {
		pc.TLS = &TLS{
			CAFile:                     p.Args26.NET.TLS.CAFile,
			CertificateKeyFile:         p.Args26.NET.TLS.CertificateKeyFile,
			CertificateKeyFilePassword: p.Args26.NET.TLS.CertificateKeyFilePassword,
			CertificateSelector:        p.Args26.NET.TLS.CertificateSelector,
			ClusterCertificateSelector: p.Args26.NET.TLS.ClusterCertificateSelector,
			ClusterFile:                p.Args26.NET.TLS.ClusterFile,
			ClusterPassword:            p.Args26.NET.TLS.ClusterPassword,
			CRLFile:                    p.Args26.NET.TLS.CRLFile,
			DisabledProtocols:          p.Args26.NET.TLS.DisabledProtocols,
			FIPSMode:                   p.Args26.NET.TLS.FIPSMode,
			Mode:                       p.Args26.NET.TLS.Mode,
			PEMKeyFile:                 p.Args26.NET.TLS.PEMKeyFile,
		}
	}
	if p.Args26.Security != nil {
		pc.Security = p.Args26.Security
	}
	return pc
}

// newMongosProcess generates a mongo process for a mongos
func newMongosProcess(p *ProcessConfig, cluster string) *opsmngr.Process {
	process := p.process()
	process.Cluster = cluster
	process.Args26 = opsmngr.Args26{
		NET:       p.net(),
		SystemLog: p.systemLog(),
	}
	if p.Security != nil {
		process.Args26.Security = p.Security
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
		Storage:   p.storage(),
		SystemLog: p.systemLog(),
	}
	if p.Security != nil {
		process.Args26.Security = p.Security
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
		Storage:   p.storage(),
		Sharding:  &opsmngr.Sharding{ClusterRole: "configsvr"},
		SystemLog: p.systemLog(),
	}
	if p.Security != nil {
		process.Args26.Security = p.Security
	}
	process.LogRotate = newLogRotate()
	if p.AuditLogPath != "" {
		process.Args26.AuditLog = p.auditLog()
	}

	return process
}

// net maps convert.ProcessConfig -> opsmngr.Net
func (p *ProcessConfig) net() opsmngr.Net {
	net := opsmngr.Net{
		Port:      p.Port,
		BindIP:    p.BindIP,
		BindIPAll: p.BindIPAll,
	}

	if p.TLS != nil {
		net.TLS = &opsmngr.TLS{
			CAFile:                     p.TLS.CAFile,
			CertificateKeyFile:         p.TLS.CertificateKeyFile,
			CertificateKeyFilePassword: p.TLS.CertificateKeyFilePassword,
			CertificateSelector:        p.TLS.CertificateSelector,
			ClusterCertificateSelector: p.TLS.ClusterCertificateSelector,
			ClusterFile:                p.TLS.ClusterFile,
			ClusterPassword:            p.TLS.ClusterPassword,
			CRLFile:                    p.TLS.CRLFile,
			DisabledProtocols:          p.TLS.DisabledProtocols,
			FIPSMode:                   p.TLS.FIPSMode,
			Mode:                       p.TLS.Mode,
			PEMKeyFile:                 p.TLS.PEMKeyFile,
		}
	}
	return net
}

// storage maps convert.ProcessConfig -> opsmngr.Storage
func (p *ProcessConfig) storage() *opsmngr.Storage {
	return &opsmngr.Storage{
		DBPath:                 p.DBPath,
		DirectoryPerDB:         p.DirectoryPerDB,
		SyncPeriodSecs:         p.SyncPeriodSecs,
		Engine:                 p.Engine,
		WiredTiger:             p.WiredTiger,
		InMemory:               p.InMemory,
		OplogMinRetentionHours: p.OplogMinRetentionHours,
		Journal:                p.Journal,
	}
}

// systemLog maps convert.ProcessConfig -> opsmngr.SystemLog
func (p *ProcessConfig) systemLog() opsmngr.SystemLog {
	return opsmngr.SystemLog{
		Destination: p.systemLogDestination(),
		Path:        p.LogPath,
		LogAppend:   p.LogAppend,
		Verbosity:   p.LogVerbosity,
		LogRotate:   p.LogRotate,
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
		Destination: p.auditLogDestination(),
		Path:        p.AuditLogPath,
		Format:      p.AuditLogFormat,
	}
}

func (p *ProcessConfig) auditLogDestination() string {
	if p.AuditLogDestination != "" {
		return p.AuditLogDestination
	}
	return file
}

// process maps convert.ProcessConfig -> opsmngr.Process
func (p *ProcessConfig) process() *opsmngr.Process {
	process := &opsmngr.Process{
		AuthSchemaVersion:           authSchemaVersion,
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
