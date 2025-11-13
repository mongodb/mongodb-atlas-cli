// Copyright 2021 MongoDB Inc
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

package availableregions

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312009/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=list_autocomplete_mock_test.go -package=availableregions . CloudProviderRegionsLister

type CloudProviderRegionsLister interface {
	CloudProviderRegions(string, string, []string) (*atlasv2.PaginatedApiAtlasProviderRegions, error)
}

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store    CloudProviderRegionsLister
	provider string
	tier     string
}

const (
	atlasM2                    = "M2"
	atlasM5                    = "M5"
	deprecateMessageSharedTier = "The '%s' tier is deprecated. Please use '--tier FLEX' instead. For the migration guide and timeline, visit: https://dochub.mongodb.org/core/flex-migration.\n"
)

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var listTemplate = `PROVIDER	INSTANCE SIZE	REGIONS{{range valueOrEmptySlice .Results}}{{ $providerName := .Provider }}{{range valueOrEmptySlice .InstanceSizes}}
{{$providerName}}	{{.Name}}	{{range valueOrEmptySlice .AvailableRegions}}{{.Name}} {{end}}{{end}}{{end}}
`

func (opts *ListOpts) Run() error {
	// Set provider if existent
	var provider []string
	if opts.provider != "" {
		provider = []string{opts.provider}
	}

	r, err := opts.store.CloudProviderRegions(opts.ConfigProjectID(), opts.tier, provider)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *ListOpts) validateTier() error {
	opts.tier = strings.ToUpper(opts.tier)
	if opts.tier == atlasM2 || opts.tier == atlasM5 {
		_, _ = fmt.Fprintf(os.Stderr, deprecateMessageSharedTier, opts.tier)
	}
	return nil
}

// ListBuilder atlas cluster(s) availableRegions list --provider provider --tier tier --projectId projectId.
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "List available regions that Atlas supports for new deployments.",
		Aliases: []string{"ls"},
		Args:    require.NoArgs,
		Example: `  # List available regions for a given cloud provider and tier:
  atlas cluster availableRegions list --provider AWS --tier M50

  # List available regions by tier for a given provider:
  atlas cluster availableRegions list --provider GCP`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			if opts.tier != "" && opts.provider == "" {
				return fmt.Errorf("tier search also requires a %s flag", flag.Provider)
			}

			return opts.PreRunE(
				opts.validateTier,
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
	cmd.Flags().StringVar(&opts.tier, flag.Tier, "", usage.Tier)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	return cmd
}
