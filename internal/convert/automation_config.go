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

	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/mongocli/internal/search"
)

const (
	mongod                         = "mongod"
	atmAgentWindowsKeyFilePath     = "%SystemDrive%\\MMSAutomation\\versions\\keyfile"
	atmAgentKeyFilePathInContainer = "/var/lib/mongodb-mms-automation/keyfile"
)

// FromAutomationConfig convert from cloud format to mCLI format
func FromAutomationConfig(in *om.AutomationConfig) (out []ClusterConfig) {
	out = make([]ClusterConfig, len(in.ReplicaSets))

	for i, rs := range in.ReplicaSets {
		out[i].Name = rs.ID
		out[i].ProcessConfigs = make([]ProcessConfig, len(rs.Members))

		for j, m := range rs.Members {
			convertCloudMember(&out[i].ProcessConfigs[j], m)
			for k, p := range in.Processes {
				if p.Name == m.Host {
					convertCloudProcess(&out[i].ProcessConfigs[j], p)
					if out[i].MongoURI == "" {
						out[i].MongoURI = fmt.Sprintf("mongodb://%s:%d", p.Hostname, p.Args26.NET.Port)
					} else {
						out[i].MongoURI = fmt.Sprintf("%s,%s:%d", out[i].MongoURI, p.Hostname, p.Args26.NET.Port)
					}
					in.Processes = append(in.Processes[:k], in.Processes[k+1:]...)
					break
				}
			}
		}
	}

	return
}

func setDisabledByClusterName(out *om.AutomationConfig, name string, disabled bool) {
	// This value may not be present and is mandatory
	if out.Auth.DeploymentAuthMechanisms == nil {
		out.Auth.DeploymentAuthMechanisms = make([]string, 0)
	}
	for _, rs := range out.ReplicaSets {
		if rs.ID == name {
			for _, m := range rs.Members {
				for k, p := range out.Processes {
					if p.Name == m.Host {
						out.Processes[k].Disabled = disabled
					}
				}
			}
			break
		}
	}
}

// Shutdown a cluster processes
func Shutdown(out *om.AutomationConfig, name string) {
	setDisabledByClusterName(out, name, true)
}

// Startup a cluster processes
func Startup(out *om.AutomationConfig, name string) {
	setDisabledByClusterName(out, name, false)
}

func EnableMechanism(out *om.AutomationConfig, m []string) error {
	out.Auth.DeploymentAuthMechanisms = append(out.Auth.DeploymentAuthMechanisms, m...)
	out.Auth.AutoAuthMechanisms = append(out.Auth.AutoAuthMechanisms, m...)
	out.Auth.Disabled = false

	var err error
	if out.Auth.AutoUser == "" {
		if err := setAutoUser(out); err != nil {
			return err
		}

	}
	addMonitoringUser(out)
	addBackupUser(out)

	if out.Auth.Key == "" {
		if out.Auth.Key, err = generateRandomBase64String(500); err != nil {
			return err
		}
	}
	if out.Auth.KeyFile == "" {
		out.Auth.KeyFile = atmAgentKeyFilePathInContainer
	}
	if out.Auth.KeyFileWindows == "" {
		out.Auth.KeyFileWindows = atmAgentWindowsKeyFilePath
	}

	return nil
}

func addMonitoringUser(out *om.AutomationConfig) {
	_, exists := search.MongoDBUsers(out.Auth.Users, func(user *om.MongoDBUser) bool {
		return user.Username == monitoringAgentName
	})
	if !exists {
		AddUser(out, newMonitoringUser(out.Auth.AutoPwd))
	}
}

func addBackupUser(out *om.AutomationConfig) {
	_, exists := search.MongoDBUsers(out.Auth.Users, func(user *om.MongoDBUser) bool {
		return user.Username == backupAgentName
	})
	if !exists {
		AddUser(out, newBackupUser(out.Auth.AutoPwd))
	}
}

func setAutoUser(out *om.AutomationConfig) error {
	var err error
	out.Auth.AutoUser = automationAgentName
	if out.Auth.AutoPwd, err = generateRandomASCIIString(500); err != nil {
		return err
	}

	return nil
}

// AddUser adds a MongoDBUser to the config
func AddUser(out *om.AutomationConfig, u *om.MongoDBUser) {
	out.Auth.Users = append(out.Auth.Users, u)
}

// RemoveUser removes a MongoDBUser from the config
func RemoveUser(out *om.AutomationConfig, username string, database string) {
	pos, found := search.MongoDBUsers(out.Auth.Users, func(p *om.MongoDBUser) bool {
		return p.Username == username && p.Database == database
	})
	if found {
		out.Auth.Users = append(out.Auth.Users[:pos], out.Auth.Users[pos+1:]...)
	}
}

// convertCloudMember map cloudmanager.Member -> convert.ProcessConfig
func convertCloudMember(out *ProcessConfig, in om.Member) {
	out.Votes = in.Votes
	out.Priority = in.Priority
	out.SlaveDelay = in.SlaveDelay
	out.BuildIndexes = &in.BuildIndexes
}

// convertCloudProcess map cloudmanager.Process -> convert.ProcessConfig
func convertCloudProcess(out *ProcessConfig, in *om.Process) {
	out.DBPath = in.Args26.Storage.DBPath
	out.LogPath = in.Args26.SystemLog.Path
	out.Port = in.Args26.NET.Port
	out.ProcessType = in.ProcessType
	out.Version = in.Version
	out.FCVersion = in.FeatureCompatibilityVersion
	out.Hostname = in.Hostname
	out.Name = in.Name
}
