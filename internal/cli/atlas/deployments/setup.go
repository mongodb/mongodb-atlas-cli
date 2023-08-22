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
	"errors"
	"fmt"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

const (
	startHostPort          = 37017
	internalMongodPort     = 27017
	internalMongotPort     = 27027
	localCluster           = "LOCAL"
	mdb6                   = "6.0"
	InternalMongodHostname = "mongod"
	InternalMongotHostname = "mongot"
	ReplicaSetName         = "rs-localdev"
)

type SetupOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	deploymentName string
	deploymentType string
	mdbVersion     string
	port           int
	debug          bool
}

const startTemplate = `local environment started at {{.ConnectionString}}
`

func (opts *SetupOpts) createLocalDeployment() error {
	podmanOpts := &podman.Opts{
		Debug:     opts.debug,
		OutWriter: opts.OutWriter,
	}

	containers, errList := podmanOpts.ListContainers(InternalMongodHostname)
	if errList != nil {
		return errList
	}

	mongodContainerName := fmt.Sprintf("%s-%s", InternalMongodHostname, opts.deploymentName)
	mongotContainerName := fmt.Sprintf("%s-%s", InternalMongotHostname, opts.deploymentName)

	for _, c := range containers {
		for _, n := range c.Names {
			if n == mongodContainerName {
				return fmt.Errorf("\"%s\" deployment was already created and is currently in \"%s\" state", opts.deploymentName, c.State)
			}
		}
	}

	if opts.port == 0 {
		opts.port = getPortForNewLocalCluster(containers)
	}

	networkName := fmt.Sprintf("mdb-local-%s", opts.deploymentName)
	if _, err := podmanOpts.CreateNetwork(networkName); err != nil {
		return err
	}

	// setup mongod
	mongodDataVolume := fmt.Sprintf("mongod-local-data-%s", opts.deploymentName)
	if _, err := podmanOpts.CreateVolume(mongodDataVolume); err != nil {
		return err
	}

	if _, err := podmanOpts.RunContainer(
		podman.RunContainerOpts{
			Detach:   true,
			Image:    fmt.Sprintf("mongodb/mongodb-enterprise-server:%s-ubi8", opts.mdbVersion),
			Name:     mongodContainerName,
			Hostname: mongodContainerName,
			Volumes: map[string]string{
				mongodDataVolume: "/data/db",
			},
			Ports: map[int]int{
				opts.port: internalMongodPort,
			},
			Network: networkName,
			Args: []string{
				"--noauth",
				"--dbpath", "/data/db",
				"--replSet", ReplicaSetName,
				"--setParameter", fmt.Sprintf("mongotHost=%s:%d", mongotContainerName, internalMongotPort),
				"--setParameter", fmt.Sprintf("searchIndexManagementHostAndPort=%s:%d", mongotContainerName, internalMongotPort),
				"--setParameter", "skipAuthenticationToMongot=true",
			},
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
		ReplicaSetName,
		mongodContainerName,
		internalMongodPort,
		opts.port)
	if err := opts.seed(opts.port, seedRs); err != nil {
		return err
	}

	// setup mongot
	mongotDataVolume := fmt.Sprintf("mongot-local-data-%s", opts.deploymentName)
	if _, err := podmanOpts.CreateVolume(mongotDataVolume); err != nil {
		return err
	}

	mongotMetricsVolume := fmt.Sprintf("mongot-local-metrics-%s", opts.deploymentName)
	if _, err := podmanOpts.CreateVolume(mongotMetricsVolume); err != nil {
		return err
	}

	_, err := podmanOpts.RunContainer(podman.RunContainerOpts{
		Detach:   true,
		Image:    "mongodb/apix_test:mongot-noauth",
		Name:     mongotContainerName,
		Hostname: mongotContainerName,
		Args:     []string{"--mongodHostAndPort", fmt.Sprintf("%s:%d", mongodContainerName, internalMongodPort)},
		Volumes: map[string]string{
			mongotDataVolume:    "/var/lib/mongot",
			mongotMetricsVolume: "/var/lib/mongot/metrics",
		},
		Network: networkName,
	})

	return err
}

func connString(port int) string {
	return fmt.Sprintf("mongodb://localhost:%d/admin", port)
}

func (opts *SetupOpts) seed(port int, script string) error {
	return mongosh.Exec(opts.debug, connString(port), "--eval", script)
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

			return opts.PreRunE(opts.InitOutput(cmd.OutOrStdout(), startTemplate))
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().StringVar(&opts.deploymentType, flag.TypeFlag, "LOCAL", usage.DeploymentType)
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
