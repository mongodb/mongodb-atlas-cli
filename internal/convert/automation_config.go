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

const (
	mongod = "mongod"
)

// FromAutomationConfig convert from cloud format to mCLI format
func FromAutomationConfig(in *opsmngr.AutomationConfig) (out []ClusterConfig) {
	out = make([]ClusterConfig, len(in.ReplicaSets))

	for i, rs := range in.ReplicaSets {
		out[i].Name = rs.ID
		out[i].ProcessConfigs = make([]*ProcessConfig, len(rs.Members))

		for j, m := range rs.Members {
			out[i].ProcessConfigs[j] = convertCloudMember(m)
			for k, p := range in.Processes {
				if p.Name == m.Host {
					convertCloudProcess(out[i].ProcessConfigs[j], p)
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

// convertCloudMember map cloudmanager.Member -> convert.ProcessConfig
func convertCloudMember(in opsmngr.Member) *ProcessConfig {
	return &ProcessConfig{
		BuildIndexes: &in.BuildIndexes,
		Priority:     in.Priority,
		SlaveDelay:   in.SlaveDelay,
		Votes:        in.Votes,
	}
}

// convertCloudProcess map cloudmanager.Process -> convert.ProcessConfig
func convertCloudProcess(out *ProcessConfig, in *opsmngr.Process) {
	out.DBPath = in.Args26.Storage.DBPath
	out.LogPath = in.Args26.SystemLog.Path
	out.Port = in.Args26.NET.Port
	out.ProcessType = in.ProcessType
	out.Version = in.Version
	out.FCVersion = in.FeatureCompatibilityVersion
	out.Hostname = in.Hostname
	out.Name = in.Name
}
