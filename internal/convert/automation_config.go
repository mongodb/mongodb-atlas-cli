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

	"github.com/mongodb/go-client-mongodb-ops-manager/atmcfg"
	om "github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
	"github.com/mongodb/go-client-mongodb-ops-manager/search"
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

func AddIndexConfig(out *om.AutomationConfig, newIndex *om.IndexConfigs) {
	_, exists := search.MongoDBIndexes(out.IndexConfigs, func(index *om.IndexConfigs) bool {
		if index.RSName == newIndex.RSName && index.CollectionName == newIndex.CollectionName && index.DBName == newIndex.DBName && len(index.Key) == len(newIndex.Key) {
			// if keys are the equal the two indexes are considered to be the same
			for i := 0; i < len(index.Key); i++ {
				if index.Key[i][0] != newIndex.Key[i][0] || index.Key[i][1] != newIndex.Key[i][1] {
					return false
				}
			}

			return true
		}

		return false
	})
	if !exists {
		atmcfg.AddIndexConfig(out, newIndex)
	}
}

func addMonitoringUser(out *om.AutomationConfig) {
	_, exists := search.MongoDBUsers(out.Auth.Users, func(user *om.MongoDBUser) bool {
		return user.Username == monitoringAgentName
	})
	if !exists {
		atmcfg.AddUser(out, newMonitoringUser(out.Auth.AutoPwd))
	}
}

func addBackupUser(out *om.AutomationConfig) {
	_, exists := search.MongoDBUsers(out.Auth.Users, func(user *om.MongoDBUser) bool {
		return user.Username == backupAgentName
	})
	if !exists {
		atmcfg.AddUser(out, newBackupUser(out.Auth.AutoPwd))
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
