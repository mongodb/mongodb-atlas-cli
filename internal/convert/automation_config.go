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

const (
	mongod = "mongod"
)

// FromAutomationConfig convert from cloud format to mCLI format
func FromAutomationConfig(in *cloudmanager.AutomationConfig) (out []ClusterConfig) {
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

// Shutdown a cluster processes
func Shutdown(out *cloudmanager.AutomationConfig, name string) {
	setDisabledByClusterName(out, name, true)
}

// Startup a cluster processes
func Startup(out *cloudmanager.AutomationConfig, name string) {
	setDisabledByClusterName(out, name, false)
}

func setDisabledByClusterName(out *cloudmanager.AutomationConfig, name string, disabled bool) {
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

// convertCloudMember map cloudmanager.Member -> convert.ProcessConfig
func convertCloudMember(out *ProcessConfig, in cloudmanager.Member) {
	out.Votes = in.Votes
	out.Priority = in.Priority
	out.SlaveDelay = in.SlaveDelay
	out.BuildIndexes = &in.BuildIndexes
}

// convertCloudProcess map cloudmanager.Process -> convert.ProcessConfig
func convertCloudProcess(out *ProcessConfig, in *cloudmanager.Process) {
	out.DBPath = in.Args26.Storage.DBPath
	out.LogPath = in.Args26.SystemLog.Path
	out.Port = in.Args26.NET.Port
	out.ProcessType = in.ProcessType
	out.Version = in.Version
	out.FCVersion = in.FeatureCompatibilityVersion
	out.Hostname = in.Hostname
	out.Name = in.Name
}
