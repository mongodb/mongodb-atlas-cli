// Copyright 2020 MongoDB Inc
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

package containers

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_mock_test.go -package=containers . Lister

type Lister interface {
	ContainersByProvider(string, *store.ContainersListOptions) ([]atlasv2.CloudProviderContainer, error)
	AllContainers(string, *store.ListOptions) ([]atlasv2.CloudProviderContainer, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	provider string
	store    Lister
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `ID	PROVIDER	REGION	ATLAS CIDR	PROVISIONED{{range valueOrEmptySlice .}}
{{.Id}}	{{.ProviderName}}	{{if .RegionName}}{{.RegionName}}{{else}}{{.Region}}{{end}}	{{.AtlasCidrBlock}}	{{.Provisioned}}{{end}}
`

func (opts *ListOpts) Run() error {
	var r []atlasv2.CloudProviderContainer
	var err error
	if opts.provider == "" {
		r, err = opts.store.AllContainers(opts.ConfigProjectID(), opts.NewAtlasListOptions())
	} else {
		listOpts := opts.newContainerListOptions()
		r, err = opts.store.ContainersByProvider(opts.ConfigProjectID(), listOpts)
	}
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) newContainerListOptions() *store.ContainersListOptions {
	lst := opts.NewAtlasListOptions()
	return &store.ContainersListOptions{
		ProviderName: opts.provider,
		ListOptions: store.ListOptions{
			PageNum:      lst.PageNum,
			ItemsPerPage: lst.ItemsPerPage,
			IncludeCount: lst.IncludeCount,
		},
	}
}

// atlas networking container(s) list [--projectId projectId] [--page N] [--limit N] [--minDate minDate] [--maxDate maxDate].
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Return all network peering containers for your project.",
		Long:    fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # Return a JSON-formatted list of network peering containers in the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas networking containers list --projectId 5e2211c17a3e5a48f5497de3 --output json`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.provider, flag.Provider, "", usage.Provider)
	opts.AddListOptsFlags(cmd)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
