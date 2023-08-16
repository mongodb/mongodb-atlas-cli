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
	_ "embed"
	"errors"
	"fmt"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/atlas/deployments/podman"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mongosh"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

var (
	//go:embed files/ca.pem
	CaContents []byte

	//go:embed files/seedRs.js
	SeedRsContents []byte

	//go:embed files/seedUser.js
	SeedUserContents []byte
)

type StartOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	debug bool
}

var startTemplate = `local environment started at {{.ConnectionString}}
`

func (opts *StartOpts) startWithPodman() error {
	if err := podman.CreateNetwork(opts.debug, "mdb-local-1"); err != nil {
		return err
	}
	fmt.Println("network created")

	if err := podman.CreateVolume(opts.debug, "mms-data-1"); err != nil {
		return err
	}

	if err := podman.CreateVolume(opts.debug, "mongo-data-1"); err != nil {
		return err
	}

	if err := podman.CreateVolume(opts.debug, "mongot-data-1"); err != nil {
		return err
	}

	if err := podman.CreateVolume(opts.debug, "mongot-metrics-1"); err != nil {
		return err
	}

	fmt.Println("volumes created")

	hostPort := 37017
	if err := podman.RunContainer(opts.debug,
		podman.RunContainerOpts{
			Detach:   true,
			Image:    "mongodb/mongodb-enterprise-server:6.0.9-ubi8",
			Name:     "mongod1",
			Hostname: "mongod1",
			Volumes: map[string]string{
				"mongo-data-1": "/data/db",
			},
			Ports: map[int]int{
				hostPort: 27017,
			},
			Network: "mdb-local-1",
			Args: []string{
				"--noauth",
				"--dbpath", "/data/db",
				"--replSet", "rs0",
				"--setParameter", "mongotHost=mongot1:27027",
				"--setParameter", "searchIndexManagementHostAndPort=mongot1:27027",
				"--setParameter", "skipAuthenticationToMongot=true",
			},
		}); err != nil {
		return err
	}
	fmt.Println("mongod container started")
	opts.waitConnection(hostPort)

	for _, seedScriptContents := range []string{string(SeedRsContents), string(SeedUserContents)} {
		if err := opts.seed(hostPort, seedScriptContents); err != nil {
			fmt.Println(err)
			return err
		}
	}
	fmt.Println("seed RS completed")

	if err := podman.RunContainer(opts.debug, podman.RunContainerOpts{
		Detach:   true,
		Image:    "mongodb/apix_test:mongot-noauth",
		Name:     "mongot1",
		Hostname: "mongot1",
		Volumes: map[string]string{
			"mongot-data-1":    "/var/lib/mongot",
			"mongot-metrics-1": "/var/lib/mongot/metrics",
		},
		Network: "mdb-local-1",
	}); err != nil {
		return err
	}
	fmt.Println("mongot container started")

	return nil
}

func connString(port int) string {
	return fmt.Sprintf("mongodb://localhost:%d/admin", port)
}

func (opts *StartOpts) seed(port int, script string) error {
	return mongosh.Exec(opts.debug, connString(port), "--eval", script)
}

func (opts *StartOpts) waitConnection(port int) error {
	for i := 0; i < 60; i++ { // 60 seconds
		if err := mongosh.Exec(opts.debug, connString(port), "--eval", "db.runCommand('ping').ok"); err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
	return errors.New("waitConnection failed")
}

func (opts *StartOpts) Run(_ context.Context) error {
	if err := opts.startWithPodman(); err != nil {
		return err
	}

	return opts.Print(localData)
}

// atlas local start.
func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	cmd := &cobra.Command{
		Use:   "start",
		Short: "Starts a local instance.",
		Args:  require.NoArgs,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), startTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	cmd.Flags().BoolVarP(&opts.debug, flag.Debug, flag.DebugShort, false, usage.Debug)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
