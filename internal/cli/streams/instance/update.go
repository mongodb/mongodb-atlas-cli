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

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312012/admin"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=update_mock_test.go -package=instance . StreamsUpdater

type StreamsUpdater interface {
	UpdateStream(string, string, *atlasv2.StreamsTenantUpdateRequest) (*atlasv2.StreamsTenant, error)
}

type UpdateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.InputOpts
	name     string
	provider string
	region   string
	store    StreamsUpdater
}

const (
	updateTemplate = "Atlas Streams Processor Instance '{{.Name}}' successfully updated.\n"
)

func (opts *UpdateOpts) Run() error {
	updateRequest := opts.streamsUpdateRequest()
	r, err := opts.store.UpdateStream(opts.ProjectID, opts.name, updateRequest)

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) streamsUpdateRequest() *atlasv2.StreamsTenantUpdateRequest {
	request := atlasv2.NewStreamsTenantUpdateRequestWithDefaults()

	if opts.provider != "" {
		request.CloudProvider = &opts.provider
	}

	if opts.region != "" {
		request.Region = &opts.region
	}

	return request
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// CreateBuilder
// atlas streams instance update [name]
// --provider AWS
// --region VIRGINIA_USA.
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{}
	cmd := &cobra.Command{
		Use:   "update <name>",
		Short: "Updates an Atlas Stream Processing instance for your project.",
		Long: `Before updating an Atlas Streams Processing instance, you must first stop all processes associated with it.
` + fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"nameDesc": "Name of the Atlas Stream Processing instance. After creation, you can't change the name of the instance. The name can contain ASCII letters, numbers, and hyphens.",
			"output":   updateTemplate,
		},
		Example: `  # Modify the Atlas Stream Processing instance configuration with the name MyInstance:
  atlas streams instance update MyInstance --provider AWS --region VIRGINIA_USA`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.name = args[0]

			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), updateTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.provider, flag.Provider, "AWS", usage.StreamsProvider)
	cmd.Flags().StringVarP(&opts.region, flag.Region, flag.RegionShort, "", usage.StreamsRegion)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagRequired(flag.Provider)
	_ = cmd.MarkFlagRequired(flag.Region)

	return cmd
}
