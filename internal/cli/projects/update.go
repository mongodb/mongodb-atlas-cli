// Copyright 2024 MongoDB Inc
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

package projects

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const updateTemplate = "Project '{{.Id}}' updated.\n"

type UpdateOpts struct {
	cli.OutputOpts
	projectID string
	filename  string
	fs        afero.Fs
	store     store.ProjectUpdater
}

func (opts *UpdateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *UpdateOpts) Run() error {
	params, err := opts.newUpdateProjectParams()
	if err != nil {
		return err
	}

	r, err := opts.store.UpdateProject(params)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *UpdateOpts) newUpdateProjectParams() (*atlasv2.UpdateProjectApiParams, error) {
	groupUpdate := new(atlasv2.GroupUpdate)
	if err := file.Load(opts.fs, opts.filename, groupUpdate); err != nil {
		return nil, err
	}

	return &atlasv2.UpdateProjectApiParams{
		GroupId: opts.projectID, GroupUpdate: groupUpdate}, nil
}

func (opts *UpdateOpts) validateProjectID() error {
	return validate.ObjectID(opts.projectID)
}

// atlas project(s) update <projectId> [--file filePath].
func UpdateBuilder() *cobra.Command {
	opts := &UpdateOpts{fs: afero.NewOsFs()}
	opts.Template = updateTemplate
	cmd := &cobra.Command{
		Use:   "update <ID>",
		Short: "Update a project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.ExactArgs(1),
		Annotations: map[string]string{
			"IDDesc": "ID of the project you want to update.",
			"output": updateTemplate,
		},
		Example: `  # Update a project with the ID 5e2211c17a3e5a48f5497de3 using the JSON file named myProject.json:
  atlas projects update 5f4007f327a3bd7b6f4103c5 --file myProject.json --output json`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			opts.projectID = args[0]
			return prerun.ExecuteE(
				opts.validateProjectID,
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.ProjectConfigFilename)
	_ = cmd.MarkFlagRequired(flag.File)
	_ = cmd.MarkFlagFilename(flag.File)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)
	_ = cmd.RegisterFlagCompletionFunc(flag.Output, opts.AutoCompleteOutputFlag())

	return cmd
}
