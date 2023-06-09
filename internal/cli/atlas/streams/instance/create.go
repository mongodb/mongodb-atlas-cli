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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	name     string
	provider string
	region   string
	store    store.StreamsCreator
}

const (
	createTemplate = "Atlas Streams Processor Instance '{{.Name}}' successfully created.\n"
)

func (opts *CreateOpts) Run() error {

	streamProcessor := new(store.StreamProcessorInstance)
	streamProcessor.Name = opts.name
	streamProcessor.GroupID = opts.ProjectID
	streamProcessor.DataProcessRegion.CloudProvider = opts.provider
	streamProcessor.DataProcessRegion.Region = opts.region

	r, err := opts.store.CreateStream(streamProcessor)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) validate() error {
	return nil
}

// CreateBuilder
// atlas streams instance create [name]
// --provider AWS
// --region US-EAST-1
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create [name]...",
		Short: "Create an Atlas Stream Processor Instance for your project",
		Long:  `To get started quickly, specify a name, a cloud provider, and a region to configure an Atlas Streams processor instance.` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: fmt.Sprintf(`  # Deploy an Atlas Streams provider instance called myProcessor for the project with the ID 5e2211c17a3e5a48f5497de3:
  %[1]s streams create myProcessor --projectId 5e2211c17a3e5a48f5497de3 --provider AWS --region US_EAST_1`, cli.ExampleAtlasEntryPoint()),
		Args: require.MaximumNArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the Atlas Streams processor instance. The instance name cannot be changed after the processor is created. The name can contain ASCII letters, numbers, and hyphens.",
			"output":   createTemplate,
		},
		ValidArgs: []string{"name"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("Atlas Streams Processor instance name missing")
			}

			if len(args) != 0 {
				opts.name = args[0]
			}

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
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
	cmd.MarkFlagRequired(flag.Provider)
	cmd.MarkFlagRequired(flag.Region)

	return cmd
}
