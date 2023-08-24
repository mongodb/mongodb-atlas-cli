package options

import "fmt"

const (
	MongodHostnamePrefix = "mongod"
	MongotHostnamePrefix = "mongot"
)

type DeploymentOpts struct {
	DeploymentName string
	DeploymentType string
	DeploymentID   string
	MdbVersion     string
	Port           int
}

func (opts *DeploymentOpts) LocalMongodHostname() string {
	return fmt.Sprintf("%s-%s", MongodHostnamePrefix, opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongotHostname() string {
	return fmt.Sprintf("%s-%s", MongotHostnamePrefix, opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalNetworkName() string {
	return fmt.Sprintf("mdb-local-%s", opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongotDataVolume() string {
	return fmt.Sprintf("mongot-local-data-%s", opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongodDataVolume() string {
	return fmt.Sprintf("mongod-local-data-%s", opts.DeploymentName)
}

func (opts *DeploymentOpts) LocalMongoMetricsVolume() string {
	return fmt.Sprintf("mongot-local-metrics-%s", opts.DeploymentName)
}
