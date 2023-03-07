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
	"io/ioutil"
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

const createTemplate = `NAME	PORT	CONNECTION STRING
{{.Name}}	{{.Port}}	mongodb://localhost:{{.Port}}
`

func (opts *CreateOpts) Run(ctx context.Context) error {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	cli.NegotiateAPIVersion(ctx)

	port := strconv.Itoa(rand.IntnRange(10000, 65000))

	_, portBindings, err := nat.ParsePortSpecs([]string{fmt.Sprintf("%s:27017", port)})
	if err != nil {
		return err
	}

	img, err := cli.ImagePull(ctx, "mongodb/mongodb-community-server:6.0-ubuntu2204", types.ImagePullOptions{})
	if err != nil {
		return err
	}
	defer img.Close()
	io.Copy(ioutil.Discard, img) //send to debug logs

	resp, err := cli.ContainerCreate(ctx, &container.Config{Image: "mongodb/mongodb-community-server:6.0-ubuntu2204", Labels: map[string]string{"atlascli": "true"}}, &container.HostConfig{PortBindings: portBindings}, nil, nil, opts.name)
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
