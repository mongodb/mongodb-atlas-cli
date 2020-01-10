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
	"github.com/mongodb-labs/pcgc/cloudmanager"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

const (
	mongod = "mongod"
)

var supportedExts = []string{"json", "yaml", "yml"}

// ReadInClusterConfig load a ClusterConfig from a YAML or JSON file
func ReadInClusterConfig(fs afero.Fs, filename string) (*ClusterConfig, error) {
	if exists, err := afero.Exists(fs, filename); !exists || err != nil {
		return nil, fmt.Errorf("file not found: %s", filename)
	}

	ext := filepath.Ext(filename)
	if len(ext) <= 1 {
		return nil, fmt.Errorf("filename: %s requires valid extension", filename)
	}
	configType := ext[1:]
	if !utils.StringInSlice(configType, supportedExts) {
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

// FromAutomationConfig convert from cloud format to mCLI format
func FromAutomationConfig(in *cloudmanager.AutomationConfig) (out []ClusterConfig) {
	out = make([]ClusterConfig, len(in.ReplicaSets))

	for i, rs := range in.ReplicaSets {
		out[i].Name = rs.ID
		out[i].Processes = make([]ProcessConfig, len(rs.Members))

		for j, m := range rs.Members {
			convertCloudMember(&out[i].Processes[j], m)
			for k, p := range in.Processes {
				if p.Name == m.Host {
					convertCloudProcess(&out[i].Processes[j], p)
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
}
