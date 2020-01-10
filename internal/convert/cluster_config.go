package convert

import (
	"fmt"

	"github.com/Masterminds/semver"
	"github.com/mongodb-labs/pcgc/cloudmanager"
)

const (
	zero = "0"
	one  = "1"
)

// ClusterConfig configuration for a cluster
// This cluster can be used to patch an automation config
type ClusterConfig struct {
	FCVersion string          `yaml:"feature_compatibility_version,omitempty" json:"feature_compatibility_version,omitempty"`
	MongoURI  string          `yaml:"mongoURI,omitempty" json:"mongoURI,omitempty"`
	Name      string          `yaml:"name" json:"name"`
	Processes []ProcessConfig `yaml:"processes" json:"processes"`
	Version   string          `yaml:"version,omitempty" json:"version,omitempty"`
}

// ProcessConfig that belongs to a cluster
type ProcessConfig struct {
	BuildIndexes *bool   `yaml:"buildIndexes,omitempty" json:"buildIndexes,omitempty"`
	DBPath       string  `yaml:"db_path" json:"db_path"`
	FCVersion    string  `yaml:"feature_compatibility_version,omitempty" json:"feature_compatibility_version,omitempty"`
	Hostname     string  `yaml:"hostname" json:"hostname"`
	LogPath      string  `yaml:"log_path" json:"log_path"`
	Port         int     `yaml:"port" json:"port"`
	Priority     float64 `yaml:"priority" json:"priority"`
	ProcessType  string  `yaml:"process_type" json:"process_type"`
	SlaveDelay   float64 `yaml:"slave_delay" json:"slave_delay"`
	Version      string  `yaml:"version,omitempty" json:"version,omitempty"`
	Votes        float64 `yaml:"votes" json:"votes"`
	ArbiterOnly  bool    `yaml:"arbiter_only" json:"arbiter_only"`
	Disabled     bool    `yaml:"disabled" json:"disabled"`
	Hidden       bool    `yaml:"hidden" json:"hidden"`
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

	// TODO: remove when automation fixes this CLOUDP-55268
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
		return zero, nil
	}
	return one, nil
}

func (p *ProcessConfig) toCMProcess(i int, name, version, fcVersion string) *cloudmanager.Process {
	if p.ProcessType == "" {
		p.ProcessType = mongod
	}
	if p.Version == "" {
		p.Version = version
	}
	if p.FCVersion == "" {
		p.FCVersion = fcVersion
	}

	process := &cloudmanager.Process{
		AuthSchemaVersion:           5,
		Disabled:                    p.Disabled,
		ManualMode:                  false,
		ProcessType:                 p.ProcessType,
		Version:                     p.Version,
		FeatureCompatibilityVersion: p.FCVersion,
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
	if p.BuildIndexes == nil {
		defaultValue := true
		p.BuildIndexes = &defaultValue
	}

	return cloudmanager.Member{
		ID:           i,
		ArbiterOnly:  p.ArbiterOnly,
		BuildIndexes: *p.BuildIndexes,
		Hidden:       p.Hidden,
		Host:         fmt.Sprintf("%s_%d", name, i),
		Priority:     p.Priority,
		SlaveDelay:   p.SlaveDelay,
		Votes:        p.Votes,
	}
}
