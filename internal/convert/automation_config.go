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
