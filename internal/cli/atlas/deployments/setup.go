// Copyright 2023 MongoDB Inc
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

package deployments

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	startHostPort        = 37017
	internalMongodPort   = 27017
	internalMongotPort   = 27027
	localCluster         = "LOCAL"
	mdb6                 = "6.0"
	mongodHostnamePrefix = "mongod"
	mongotHostnamePrefix = "mongot"
	replicaSetName       = "rs-localdev"
	// based on https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster
	clusterNamePattern = "^[a-zA-Z0-9][a-zA-Z0-9-]*$"
)

var ErrPodmanNotReady = errors.New("podman not found in your system, check requirements at http://docpage")

type SetupOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	podmanClient   podman.Client
	deploymentName string
	deploymentType string
	deploymentID   string
	mdbVersion     string
	port           int
	debug          bool
}

const startTemplate = `local environment started at {{.ConnectionString}}
`

func (opts *SetupOpts) initPodmanClient() error {
	opts.podmanClient = podman.NewClient(opts.debug, opts.OutWriter)
	return nil
}

func (opts *SetupOpts) createLocalDeployment() error {
	if !opts.podmanClient.Ready() {
		return ErrPodmanNotReady
	}

	if err := opts.podmanClient.Setup(); err != nil {
		return err
	}

	containers, errList := opts.podmanClient.ListContainers(mongodHostnamePrefix)
	if errList != nil {
		return errList
	}

	if opts.port == 0 {
		opts.port = getPortForNewLocalCluster(containers)
	}

	if err := opts.validateLocalDeploymentsSettings(containers); err != nil {
		return err
	}

	if _, err := opts.podmanClient.CreateNetwork(opts.networkName()); err != nil {
		return err
	}

	if err := opts.configureMongod(); err != nil {
		return err
	}

	return opts.configureMongot()
}

func (opts *SetupOpts) configureMongod() error {
	mongodDataVolume := fmt.Sprintf("mongod-local-data-%s", opts.deploymentName)
	if _, err := opts.podmanClient.CreateVolume(mongodDataVolume); err != nil {
		return err
	}

	keyfile := "/data/configdb/keyfile"
	keyfilePerm := 400
	mongodArgs := []string{
		"--transitionToAuth",
		"--keyFile", keyfile,
		"--dbpath", "/data/db",
		"--replSet", replicaSetName,
		"--setParameter", fmt.Sprintf("mongotHost=%s:%d", opts.mongotHostname(), internalMongotPort),
		"--setParameter", fmt.Sprintf("searchIndexManagementHostAndPort=%s:%d", opts.mongotHostname(), internalMongotPort),
	}

	// wrap the entrypoint with a chain of commands that
	// creates the keyfile in the container and sets the 400 permissions for it,
	// then starts the entrypoint with the local dev config
	cmdTemplate := "echo '%[1]s' > %[2]s && chmod %[3]d %[2]s && python3 /usr/local/bin/docker-entrypoint.py %[4]s"
	cmd := fmt.Sprintf(cmdTemplate,
		base64.URLEncoding.EncodeToString([]byte(opts.deploymentID)),
		keyfile,
		keyfilePerm,
		strings.Join(mongodArgs, " "))

	if _, err := opts.podmanClient.RunContainer(
		podman.RunContainerOpts{
			Detach:   true,
			Image:    fmt.Sprintf("mongodb/mongodb-enterprise-server:%s-ubi8", opts.mdbVersion),
			Name:     opts.mongodHostname(),
			Hostname: opts.mongodHostname(),
			Volumes: map[string]string{
				mongodDataVolume: "/data/db",
			},
			Ports: map[int]int{
				opts.port: internalMongodPort,
			},
			Network: opts.networkName(),
			Args:    []string{"sh", "-c", cmd},
		}); err != nil {
		return err
	}

	// init ReplicaSet
	if err := opts.waitConnection(opts.port); err != nil {
		return err
	}

	seedRs := fmt.Sprintf(`try {
		rs.status();
	  } catch {
		rs.initiate({
		  _id: "%s",
		  version: 1,
		  configsvr: false,
		  members: [{ _id: 0, host: "%s:%d", horizons: { external: "localhost:%d" } }],
		});
	  }`,
		replicaSetName,
		opts.mongodHostname(),
		internalMongodPort,
		opts.port)
	if err := opts.seed(opts.port, seedRs); err != nil {
		return err
	}

	return opts.seed(opts.port, "db.getSiblingDB('admin').atlascli.insertOne({ managedClusterType: 'atlasCliLocalDevCluster' })")
}

func (opts *SetupOpts) configureMongot() error {
	mongotDataVolume := fmt.Sprintf("mongot-local-data-%s", opts.deploymentName)
	if _, err := opts.podmanClient.CreateVolume(mongotDataVolume); err != nil {
		return err
	}

	mongotMetricsVolume := fmt.Sprintf("mongot-local-metrics-%s", opts.deploymentName)
	if _, err := opts.podmanClient.CreateVolume(mongotMetricsVolume); err != nil {
		return err
	}

	_, err := opts.podmanClient.RunContainer(podman.RunContainerOpts{
		Detach:   true,
		Image:    "mongodb/apix_test:mongot",
		Name:     opts.mongotHostname(),
		Hostname: opts.mongotHostname(),
		Args: []string{
			"--mongodHostAndPort", fmt.Sprintf("%s:%d", opts.mongodHostname(), internalMongodPort),
			"--keyFileContent", base64.URLEncoding.EncodeToString([]byte(opts.deploymentID)),
		},
		Volumes: map[string]string{
			mongotDataVolume:    "/var/lib/mongot",
			mongotMetricsVolume: "/var/lib/mongot/metrics",
		},
		Network: opts.networkName(),
	})

	return err
}

func connString(port int) string {
	return fmt.Sprintf("mongodb://localhost:%d/admin", port)
}

func (opts *SetupOpts) seed(port int, script string) error {
	return mongosh.Exec(opts.debug, connString(port), "--eval", script)
}

func (opts *SetupOpts) validateLocalDeploymentsSettings(containers []podman.Container) error {
	if matched, _ := regexp.MatchString(clusterNamePattern, opts.deploymentName); !matched {
		return fmt.Errorf("%s is not a valid clusterName", opts.deploymentName)
	}

	mongodContainerName := opts.mongodHostname()
	for _, c := range containers {
		for _, n := range c.Names {
			if n == mongodContainerName {
				return fmt.Errorf("\"%s\" deployment was already created and is currently in \"%s\" state", opts.deploymentName, c.State)
			}
		}

		for _, p := range c.Ports {
			if p.HostPort == opts.port {
				return fmt.Errorf("port %d is already used by \"%s\" local deployment", opts.port, c.Names[0])
			}
		}
	}

	return nil
}

func (opts *SetupOpts) mongodHostname() string {
	return fmt.Sprintf("%s-%s", mongodHostnamePrefix, opts.deploymentName)
}

func (opts *SetupOpts) mongotHostname() string {
	return fmt.Sprintf("%s-%s", mongotHostnamePrefix, opts.deploymentName)
}

func (opts *SetupOpts) networkName() string {
	return fmt.Sprintf("mdb-local-%s", opts.deploymentName)
}

func (opts *SetupOpts) waitConnection(port int) error {
	const waitSeconds = 60
	for i := 0; i < waitSeconds; i++ {
		if err := mongosh.Exec(opts.debug, connString(port), "--eval", "db.runCommand('ping').ok"); err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return errors.New("waitConnection failed")
}

func getPortForNewLocalCluster(existingContainers []podman.Container) int {
	maxPort := startHostPort - 1
	for _, c := range existingContainers {
		for _, p := range c.Ports {
			if maxPort < p.HostPort {
				maxPort = p.HostPort
			}
		}
	}
	return maxPort + 1
}

func (opts *SetupOpts) Run(_ context.Context) error {
	if err := opts.createLocalDeployment(); err != nil {
		return err
	}

	return opts.Print(map[string]string{"ConnectionString": fmt.Sprintf("mongodb://localhost:%d", opts.port)})
}

// atlas deployments setup.
func SetupBuilder() *cobra.Command {
	opts := &SetupOpts{}
	cmd := &cobra.Command{
		Use:   "setup <clusterName>",
		Short: "Create a local deployment.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"clusterNameDesc": "Name of the cluster you want to setup.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.deploymentName = args[0]
			opts.deploymentID = uuid.NewString()

			return opts.PreRunE(opts.InitOutput(cmd.OutOrStdout(), startTemplate), opts.initPodmanClient)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.deploymentType, flag.TypeFlag, localCluster, usage.DeploymentType)
	cmd.Flags().IntVar(&opts.port, flag.Port, 0, usage.MongodPort)
	cmd.Flags().StringVar(&opts.mdbVersion, flag.MDBVersion, mdb6, usage.MDBVersion)

	cmd.Flags().BoolVarP(&opts.debug, flag.Debug, flag.DebugShort, false, usage.Debug)

	_ = cmd.RegisterFlagCompletionFunc(flag.MDBVersion, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{mdb6}, cobra.ShellCompDirectiveDefault
	})
	_ = cmd.RegisterFlagCompletionFunc(flag.TypeFlag, func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return []string{localCluster}, cobra.ShellCompDirectiveDefault
	})

	return cmd
}
