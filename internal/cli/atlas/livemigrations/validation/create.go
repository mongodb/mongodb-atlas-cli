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
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
	"go.mongodb.org/atlas/mongodbatlas"
)

type CreateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	destinationClusterName      string
	destinationDropEnabled      bool
	migrationHosts              []string
	sourceCACertificatePath     string
	sourceClusterName           string
	sourceProjectID             string
	sourceSSL                   bool
	sourceManagedAuthentication bool
	sourceUsername              string
	sourcePassword              string
	store                       store.LiveMigrationValidationsCreator
}

func (opts *CreateOpts) initStore() error {
	var err error
	opts.store, err = store.New(store.AuthenticatedPreset(config.Default()))
	return err
}

var createTemplate = `ID	PROJECT ID	SOURCE PROJECT ID	STATUS
{{.ID}}	{{.GroupID}}	{{.SourceGroupID}}	{{.Status}}`

func (opts *CreateOpts) Run() error {
	createRequest := opts.newValidationCreateRequest()

	r, err := opts.store.CreateValidation(opts.ConfigProjectID(), createRequest)
	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *CreateOpts) newValidationCreateRequest() *mongodbatlas.LiveMigration {
	return &mongodbatlas.LiveMigration{
		Source: &mongodbatlas.Source{
			ClusterName:           opts.sourceClusterName,
			GroupID:               opts.sourceProjectID,
			Username:              opts.sourceUsername,
			Password:              opts.sourcePassword,
			SSL:                   &opts.sourceSSL,
			CACertificatePath:     opts.sourceCACertificatePath,
			ManagedAuthentication: &opts.sourceManagedAuthentication,
		},
		Destination: &mongodbatlas.Destination{
			ClusterName: opts.destinationClusterName,
			GroupID:     opts.ConfigProjectID(),
		},
		MigrationHosts: opts.migrationHosts,
		DropEnabled:    &opts.destinationDropEnabled,
	}
}

// mongocli atlas liveMigrations|lm validation create --clusterName clusterName --migrationHosts hosts --sourceClusterName clusterName --sourceProjectId projectId [--sourceSSL] [--sourceCACertificatePath path] [--sourceManagedAuthentication] [--sourceUsername userName] [--sourcePassword password] [--drop] [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{}
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create one new validation request.",
		Long:  "Your API Key must have the Organization Owner role to successfully run this command.",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateOrgID,
				opts.ValidateProjectID,
				opts.initStore,
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.OrgID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.sourceClusterName, flag.LiveMigrationSourceClusterName, "", usage.LiveMigrationSourceClusterName)
	cmd.Flags().StringVar(&opts.sourceProjectID, flag.LiveMigrationSourceProjectID, "", usage.LiveMigrationSourceProjectID)
	cmd.Flags().StringVar(&opts.sourceUsername, flag.LiveMigrationSourceUsername, "", usage.LiveMigrationSourceUsername)
	cmd.Flags().StringVar(&opts.sourcePassword, flag.LiveMigrationSourcePassword, "", usage.LiveMigrationSourcePassword)
	cmd.Flags().BoolVar(&opts.sourceSSL, flag.LiveMigrationSourceSSL, false, usage.LiveMigrationSourceSSL)
	cmd.Flags().StringVar(&opts.sourceCACertificatePath, flag.LiveMigrationSourceCACertificatePath, "", usage.LiveMigrationSourceCACertificatePath)
	cmd.Flags().BoolVar(&opts.sourceManagedAuthentication, flag.LiveMigrationSourceManagedAuthentication, false, usage.LiveMigrationSourceManagedAuthentication)
	cmd.Flags().StringVar(&opts.destinationClusterName, flag.LiveMigrationDestinationClusterName, "", usage.LiveMigrationDestinationClusterName)
	cmd.Flags().StringVar(&opts.ProjectID, flag.LiveMigrationDestinationProjectID, "", usage.LiveMigrationDestinationProjectID)
	cmd.Flags().StringSliceVar(&opts.migrationHosts, flag.LiveMigrationHost, []string{}, usage.LiveMigrationHostEntries)
	cmd.Flags().BoolVar(&opts.destinationDropEnabled, flag.LiveMigrationDropCollections, false, usage.LiveMigrationDropCollections)

	_ = cmd.MarkFlagRequired(flag.LiveMigrationSourceClusterName)
	_ = cmd.MarkFlagRequired(flag.LiveMigrationSourceProjectID)
	_ = cmd.MarkFlagRequired(flag.LiveMigrationDestinationClusterName)
	_ = cmd.MarkFlagRequired(flag.LiveMigrationHost)

	return cmd
}
