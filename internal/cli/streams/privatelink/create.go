// Copyright 2025 MongoDB Inc
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

package privatelink

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113005/admin"
)

var createTemplate = "Atlas Stream Processing PrivateLink endpoint {{.InterfaceEndpointId}} created.\n"

type CreateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	store    store.PrivateLinkCreator
	filename string
	fs       afero.Fs
}

func (opts *CreateOpts) Run() error {
	privateLinkEndpoint := atlasv2.NewStreamsPrivateLinkConnection()
	if err := file.Load(opts.fs, opts.filename, privateLinkEndpoint); err != nil {
		return err
	}

	// Remaining validation will be done by the API
	if privateLinkEndpoint.GetProvider() == "" {
		return errors.New("provider missing")
	}

	result, err := opts.store.CreatePrivateLinkEndpoint(opts.ConfigProjectID(), privateLinkEndpoint)
	if err != nil {
		return err
	}

	return opts.Print(result)
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams privateLink create
// -f filename: file containing the private link endpoint configuration.
// Create a PrivateLink endpoint that can be used as an Atlas Stream Processor connection.
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a PrivateLink endpoint that can be used as an Atlas Stream Processor connection.",
		Long:  fmt.Sprintf(usage.RequiredOneOfRoles, commandRoles),
		Args:  require.NoArgs,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		Example: `# create a new PrivateLink endpoint for Atlas Stream Processing:
  atlas streams privateLink create -f endpointConfig.json
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
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

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.StreamsPrivateLinkFilename)

	opts.AddProjectOptsFlags(cmd)
	opts.AddOutputOptFlags(cmd)

	_ = cmd.MarkFlagFilename(flag.File)
	_ = cmd.MarkFlagRequired(flag.File)

	return cmd
}
