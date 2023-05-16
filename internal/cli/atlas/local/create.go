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

package local

import (
	"context"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/rand"
)

type CreateOpts struct {
	cli.OutputOpts
	cli.GlobalOpts
	name string
}

const minPort = 10000
const maxPort = 65000

func randomPort() int {
	return rand.IntnRange(minPort, maxPort)
}

const createTemplate = `NAME	PORT	CONNECTION STRING
{{.Name}}	{{.Port}}	mongodb://localhost:{{.Port}}
`

func dockerImage() string {
	env := os.Getenv("ATLASCLI_LOCAL_DOCKER_IMAGE_MONGOD")
	if env != "" {
		return env
	}
	// return "mongodb/mongodb-community-server:6.0-ubuntu2204"
	return "mongo:6.0"
}

func (opts *CreateOpts) Run(ctx context.Context) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)

	port := strconv.Itoa(randomPort())

	_, portBindings, err := nat.ParsePortSpecs([]string{fmt.Sprintf("%s:27017", port)})
	if err != nil {
		return err
	}

	img, err := cli.ImagePull(ctx, dockerImage(), types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer img.Close()
	_, _ = io.Copy(io.Discard, img) // send to debug logs

	resp, err := cli.ContainerCreate(ctx, &container.Config{Image: dockerImage(), Labels: map[string]string{"atlascli": "true"}}, &container.HostConfig{PortBindings: portBindings}, nil, nil, opts.name)
	if err != nil {
		return err
	}

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}

	return opts.Print(map[string]string{"Name": opts.name, "Port": port})
}

// atlas local create <instanceName>.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <instanceName>",
		Short: "Creates a new local instance.",
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"instanceNameDesc": "Name of the local instance.",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.Run(cmd.Context())
		},
	}

	return cmd
}
