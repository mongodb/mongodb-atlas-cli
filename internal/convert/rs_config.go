package convert

import (
	"strconv"

	"github.com/Masterminds/semver/v3"

	"go.mongodb.org/ops-manager/opsmngr"
	"go.mongodb.org/ops-manager/search"
)

// RSConfig shared properties of replica sets, config servers, and sharded clusters
type RSConfig struct {
	Name           string           `yaml:"name,omitempty" json:"name,omitempty"`
	FCVersion      string           `yaml:"featureCompatibilityVersion,omitempty" json:"featureCompatibilityVersion,omitempty"`
	ProcessConfigs []*ProcessConfig `yaml:"processes,omitempty" json:"processes,omitempty"`
	Tags           []string         `yaml:"tags,omitempty" json:"tags,omitempty"`
	Version        string           `yaml:"version,omitempty" json:"version,omitempty"`
}

type patcher func(*ProcessConfig, string) *opsmngr.Process

// path is a generic replica set patcher, you'll need to provide a function that
func (c *RSConfig) path(out *opsmngr.AutomationConfig, f patcher, names ...string) error {
	newProcesses := make([]*opsmngr.Process, len(c.ProcessConfigs))
	rs, err := newReplicaSet(c)
	if err != nil {
		return err
	}
	// transform cli config to automation config
	for i, pc := range c.ProcessConfigs {
		id := strconv.Itoa(len(out.Processes) + i)
		pc.setDefaults(c)
		pn := append(names, c.Name, id)
		pc.setProcessName(out.Processes, pn...)
		newProcesses[i] = f(pc, c.Name)
		rs.Members[i] = pc.member(i)
	}
	// This value may not be present and is mandatory
	if out.Auth.DeploymentAuthMechanisms == nil {
		out.Auth.DeploymentAuthMechanisms = make([]string, 0)
	}
	patchProcesses(out, rs.ID, newProcesses)
	patchReplicaSet(out, rs)
	return nil
}

func (c *RSConfig) pathReplicaSet(out *opsmngr.AutomationConfig) error {
	return c.path(out, newReplicaSetProcess)
}

func (c *RSConfig) pathShard(out *opsmngr.AutomationConfig, name string) error {
	return c.path(out, newReplicaSetProcess, name)
}

func (c *RSConfig) pathConfigServer(out *opsmngr.AutomationConfig, name string) error {
	return c.path(out, newConfigRSProcess, name)
}

// protocolVer determines the appropriate protocol based on FCV
// returns "0" for versions <4.0 or "1" otherwise
func (c *RSConfig) protocolVer() (string, error) {
	if c.FCVersion == "" {
		return "", nil
	}
	ver, err := semver.NewVersion(c.FCVersion)
	if err != nil {
		return "", err
	}
	constrain, _ := semver.NewConstraint(fcvLessThanFour)

	if constrain.Check(ver) {
		return zero, nil
	}
	return one, nil
}

// newReplicaSet convert from cli config to opsmngr.ReplicaSet
func newReplicaSet(c *RSConfig) (*opsmngr.ReplicaSet, error) {
	pv, err := c.protocolVer()
	if err != nil {
		return nil, err
	}

	rs := &opsmngr.ReplicaSet{
		ID:              c.Name,
		Members:         make([]opsmngr.Member, len(c.ProcessConfigs)),
		ProtocolVersion: pv,
	}

	return rs, nil
}

// newRSConfig
func newRSConfig(in *opsmngr.AutomationConfig, id string) *RSConfig {
	rsi, found := search.ReplicaSets(in.ReplicaSets, func(rs *opsmngr.ReplicaSet) bool {
		return rs.ID == id
	})
	if !found {
		return nil
	}
	rs := in.ReplicaSets[rsi]
	out := &RSConfig{
		Name:           rs.ID,
		ProcessConfigs: make([]*ProcessConfig, len(rs.Members)),
	}

	for i, m := range rs.Members {
		for l, p := range in.Processes {
			if p.Name == m.Host {
				out.ProcessConfigs[i] = newReplicaSetProcessConfig(m, p)
				in.Processes = append(in.Processes[:l], in.Processes[l+1:]...)
				break
			}
		}
	}
	in.ReplicaSets = append(in.ReplicaSets[:rsi], in.ReplicaSets[rsi+1:]...)
	return out
}
