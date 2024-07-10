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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	cli.InputOpts
	name     string
	provider string
	region   string
	tier     string
	store    store.StreamsCreator
}

const (
	createTemplate = "Atlas Streams Processor Instance '{{.Name}}' successfully created.\n"
	defaultTier    = "SP30"
)

func (opts *CreateOpts) Run() error {
	streamProcessor := atlasv2.NewStreamsTenant()
	streamProcessor.Name = &opts.name
	streamProcessor.GroupId = &opts.ProjectID
	streamProcessor.DataProcessRegion = atlasv2.NewStreamsDataProcessRegion(opts.provider, opts.region)

	tierOrDefault := defaultTier
	if opts.tier != "" {
		tierOrDefault = opts.tier
	}
	streamConfig := streamProcessor.GetStreamConfig()
	streamConfig.Tier = &tierOrDefault
	streamProcessor.StreamConfig = &streamConfig

	r, err := opts.store.CreateStream(opts.ProjectID, streamProcessor)

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

// CreateBuilder
// atlas streams instance create [name]
// --provider AWS
// --region VIRGINIA_USA.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <name>",
		Short: "Create an Atlas Stream Processing instance for your project",
		Long:  `To get started quickly, specify a name, a cloud provider, and a region to configure an Atlas Stream Processing instance.` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Example: `  # Deploy an Atlas Stream Processing instance called myProcessor for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas streams instance create myProcessor --projectId 5e2211c17a3e5a48f5497de3 --provider AWS --region VIRGINIA_USA --tier SP30`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the Atlas Stream Processing instance. After creation, you can't change the name of the instance. The name can contain ASCII letters, numbers, and hyphens.",
			"output":   createTemplate,
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.provider, flag.Provider, "AWS", usage.StreamsProvider)
	cmd.Flags().StringVarP(&opts.region, flag.Region, flag.RegionShort, "", usage.StreamsRegion)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	cmd.Flags().StringVar(&opts.tier, flag.Tier, "SP30", usage.StreamsInstanceTier)

	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())
	_ = cmd.MarkFlagRequired(flag.Provider)
	_ = cmd.MarkFlagRequired(flag.Region)

	return cmd
}
