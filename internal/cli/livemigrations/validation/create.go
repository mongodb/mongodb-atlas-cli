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

package validation

import (
	"context"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/livemigrations/options"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/spf13/cobra"
)

type CreateOpts struct {
	options.LiveMigrationsOpts
	store store.LiveMigrationValidationsCreator
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

var createTemplate = `ID	PROJECT ID	SOURCE PROJECT ID	STATUS
{{.Id}}	{{.GroupId}}	{{.SourceGroupId}}	{{.Status}}`

func (opts *CreateOpts) Run() error {
	if err := opts.Prompt(); err != nil {
		return err
	}

	createRequest := opts.LiveMigrationsOpts.NewCreateRequest()

	r, err := opts.store.CreateValidation(opts.LiveMigrationsOpts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

// atlas liveMigrations|lm validation create --clusterName clusterName --migrationHosts hosts --sourceClusterName clusterName --sourceProjectId projectId [--sourceSSL] [--sourceCACertificatePath path] [--sourceManagedAuthentication] [--sourceUsername userName] [--sourcePassword password] [--drop] [--force] [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new validation request for a push live migration.",
		Long:  `To migrate using scripts, use mongomirror instead of the Atlas CLI. To learn more about mongomirror, see https://www.mongodb.com/docs/atlas/reference/mongomirror/.`,
		Annotations: map[string]string{
			"output": createTemplate,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
				opts.InitInput(cmd.InOrStdin()),
				opts.Validate,
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.GenerateFlags(cmd)
	return cmd
}
