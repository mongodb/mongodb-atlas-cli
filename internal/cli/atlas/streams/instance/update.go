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

package instance

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201008/admin"
)

type UpdateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	name     string
	provider string
	region   string
	store    store.StreamsUpdater
}

const (
	updateTemplate = "Atlas Streams Processor Instance '{{.Name}}' successfully updated.\n"
)

func (opts *UpdateOpts) Run() error {
	stream, err := opts.streams()
	if err != nil {
		return err
	}
	r, err := opts.store.UpdateStream(opts.ProjectID, opts.name, stream.DataProcessRegion)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) streams() (*atlasv2.StreamsTenant, error) {
	processor := atlasv2.NewStreamsTenant()
	processor.Name = &opts.name
	processor.GroupId = &opts.ProjectID

	processor.DataProcessRegion = atlasv2.NewStreamsDataProcessRegionWithDefaults()
	if opts.provider != "" {
		processor.DataProcessRegion.CloudProvider = opts.provider
	}

	if opts.region != "" {
		processor.DataProcessRegion.Region = opts.region
	}

	return processor, nil
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) validate() error {
	if opts.provider == "" && opts.region == "" {
		return fmt.Errorf("Either a provider or region must be provided")
	}
	//TODO: Check if the providers are correct
	//TODO: Check if I can restrict the region as well
	return nil
}

// CreateBuilder
// atlas streams instance update [name]
// --provider AWS
// --region VIRGINIA_USA
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <name>",
		Short: "Updates an Atlas Stream Processor Instance for your project",
		Long: `An Atlas Streams processor instance with running processors cannot be updated without stopping the processes first.
` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the Atlas Streams processor instance. The instance name cannot be changed after the processor is created. The name can contain ASCII letters, numbers, and hyphens.",
			"output":   updateTemplate,
		},
		Example: fmt.Sprintf(`  # Modify the Atlas Streams processor instance configuration with the ID 5d1113b25a115342cca2d1aa:
  %s streams instances update 5d1113b25a115342cca2d1aa --provider AWS --provider VIRGINIA_USA`, cli.ExampleAtlasEntryPoint()),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
				opts.validate,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVarP(&opts.provider, flag.Provider, flag.ProviderShort, "", usage.StreamsProvider)
	cmd.Flags().StringVarP(&opts.region, flag.Region, flag.RegionShort, "", usage.StreamsRegion)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
