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

package serverless

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const providerName = "SERVERLESS"

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	instanceName string
	provider     string
	region       string
	tag          map[string]string
	store        store.ServerlessInstanceCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = "Serverless instance {{.Name}} created.\n"

func (opts *CreateOpts) Run() error {
	r, err := opts.store.CreateServerlessInstance(opts.ConfigProjectID(), opts.newServerlessCreateRequestParams())
	target, ok := atlasv2.AsError(err)
	if ok && target.GetErrorCode() == "INVALID_REGION" {
		return cli.ErrNoRegionExistsTryCommand
	} else if err != nil {
		return err
	}
	return opts.Print(r)
}

func (opts *CreateOpts) newServerlessCreateRequestParams() *atlasv2.ServerlessInstanceDescriptionCreate {
	req := &atlasv2.ServerlessInstanceDescriptionCreate{
		Name: opts.instanceName,
		ProviderSettings: atlasv2.ServerlessProviderSettings{
			BackingProviderName: opts.provider,
			ProviderName:        pointer.Get(providerName),
			RegionName:          opts.region,
		},
	}

	if len(opts.tag) > 0 {
		var tags []atlasv2.ResourceTag
		for k, v := range opts.tag {
			if k != "" && v != "" {
				tags = append(tags, atlasv2.ResourceTag{Key: k, Value: v})
			}
		}
		req.Tags = &tags
	}

	return req
}

// atlas serverless|sl create <instanceName> --backingProviderName backingProviderName --providerName providerName --regionName regionName [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create <instanceName>",
		Short: "Creates one serverless instance in the specified project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Example: `  # Deploy a serverlessInstance named myInstance for the project with the ID 5e2211c17a3e5a48f5497de3:
  atlas serverless create myInstance --provider AWS --region US_EAST_1 --projectId 5e2211c17a3e5a48f5497de3`,
		Annotations: map[string]string{
			"instanceNameDesc": "Human-readable label that identifies your serverless instance.",
			"output":           createTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, args []string) error {
			opts.instanceName = args[0]
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.provider, flag.Provider, "", usage.ServerlessProvider)
	cmd.Flags().StringVar(&opts.region, flag.Region, "", usage.ServerlessRegion)
	cmd.Flags().StringToStringVar(&opts.tag, flag.Tag, nil, usage.ServerlessTag)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	_ = cmd.MarkFlagRequired(flag.Provider)
	_ = cmd.MarkFlagRequired(flag.Region)

	return cmd
}
