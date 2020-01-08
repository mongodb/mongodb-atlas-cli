package convert

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/Masterminds/semver"
	"github.com/mongodb-labs/pcgc/cloudmanager"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

const (
	mongod = "mongod"
)

var supportedExts = []string{"json", "yaml", "yml"}

// ClusterConfig configuration for a cluster
// This cluster can be used to patch an automation config
type ClusterConfig struct {
	Name      string          `yaml:"name" json:"name"`
	Processes []ProcessConfig `yaml:"processes" json:"processes"`
	Version   string          `yaml:"version" json:"version"`
	FCVersion string          `yaml:"feature_compatibility_version" json:"feature_compatibility_version"`
}

// ProcessConfig that belongs to a cluster
type ProcessConfig struct {
	Hostname    string  `yaml:"hostname" json:"hostname"`
	DBPath      string  `yaml:"db_path" json:"db_path"`
	LogPath     string  `yaml:"log_path" json:"log_path"`
	Priority    float64 `yaml:"priority" json:"priority"`
	Votes       float64 `yaml:"votes" json:"votes"`
	Port        int     `yaml:"port" json:"port"`
	SlaveDelay  float64 `yaml:"slave_delay" json:"slave_delay"`
	ProcessType string  `yaml:"process_type" json:"process_type"`
}

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
	if !stringInSlice(configType, supportedExts) {
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

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// PatchReplicaSet add the ClusterConfig to a cloudmanager.AutomationConfig
// this method will modify the given AutomationConfig to add the new replica set information
func (c *ClusterConfig) PatchReplicaSet(out *cloudmanager.AutomationConfig) error {
	protocolVer, err := c.protocolVer()
	if err != nil {
		return err
	}
	newProcesses := make([]*cloudmanager.Process, len(c.Processes))

	newReplicaSet := &cloudmanager.ReplicaSet{
		ID:              c.Name,
		Members:         make([]cloudmanager.Member, len(c.Processes)),
		ProtocolVersion: protocolVer,
	}

	for i, process := range c.Processes {
		newProcesses[i] = process.toCMProcess(i, c.Name, c.Version, c.FCVersion)
		newReplicaSet.Members[i] = process.toCMMember(i, c.Name)
	}

	// TODO: remove when automation fixes this
	if out.Auth.DeploymentAuthMechanisms == nil {
		out.Auth.DeploymentAuthMechanisms = make([]string, 0)
	}

	out.Processes = append(out.Processes, newProcesses...)
	out.ReplicaSets = append(out.ReplicaSets, newReplicaSet)

	return nil
}

func (c *ClusterConfig) protocolVer() (string, error) {
	ver, err := semver.NewVersion(c.Version)
	if err != nil {
		return "", err
	}
	constrain, _ := semver.NewConstraint("< 4.0")

	if constrain.Check(ver) {
		return "0", nil
	}
	return "1", nil
}

func (p *ProcessConfig) toCMProcess(i int, name, version, fcVersion string) *cloudmanager.Process {
	processType := p.ProcessType
	if processType == "" {
		processType = mongod
	}
	process := &cloudmanager.Process{
		AuthSchemaVersion:           5,
		Disabled:                    false,
		ManualMode:                  false,
		ProcessType:                 processType,
		Version:                     version,
		FeatureCompatibilityVersion: fcVersion,
		Hostname:                    p.Hostname,
		Name:                        fmt.Sprintf("%s_%d", name, i),
	}

	process.Args26 = cloudmanager.Args26{
		NET: cloudmanager.Net{
			Port: p.Port,
		},
		Replication: &cloudmanager.Replication{
			ReplSetName: name,
		},
		Storage: cloudmanager.Storage{
			DBPath: p.DBPath,
		},
		SystemLog: cloudmanager.SystemLog{
			Destination: "file",
			Path:        p.LogPath,
		},
	}
	process.LogRotate = &cloudmanager.LogRotate{
		SizeThresholdMB:  1000,
		TimeThresholdHrs: 24,
	}

	return process
}

func (p *ProcessConfig) toCMMember(i int, name string) cloudmanager.Member {
	return cloudmanager.Member{
		ID:           i,
		ArbiterOnly:  false,
		BuildIndexes: true,
		Hidden:       false,
		Host:         fmt.Sprintf("%s_%d", name, i),
		Priority:     p.Priority,
		SlaveDelay:   p.SlaveDelay,
		Votes:        p.Votes,
	}
}
