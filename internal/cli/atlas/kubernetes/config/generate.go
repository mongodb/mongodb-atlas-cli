// Copyright 2022 MongoDB Inc
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

package config

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/validation"
)

type GenerateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	clusterName     []string
	includeSecrets  bool
	targetNamespace string
	store           store.AtlasOperatorGenericStore
	credsStore      store.CredentialsGetter
}

func (opts *GenerateOpts) ValidateTargetNamespace() error {
	if errs := validation.IsDNS1123Label(opts.targetNamespace); len(errs) != 0 {
		return fmt.Errorf("%s parameter is invalid: %v", flag.OperatorTargetNamespace, errs)
	}
	return nil
}

func (opts *GenerateOpts) initStores(ctx context.Context) func() error {
	return func() error {
		var err error

		profile := config.Default()
		opts.store, err = store.New(store.AuthenticatedPreset(profile), store.WithContext(ctx))
		if err != nil {
			return err
		}
		opts.credsStore = profile

		return nil
	}
}

func (opts *GenerateOpts) Run() error {
	result, err := operator.NewConfigExporter(opts.store, opts.credsStore, opts.ProjectID, opts.OrgID).
		WithClustersNames(opts.clusterName).
		WithTargetNamespace(opts.targetNamespace).
		WithSecretsData(opts.includeSecrets).
		Run()
	if err != nil {
		return err
	}
	return opts.Print(result)
}

// GenerateBuilder builds a cobra.Command that can run as:
// atlas kubernetes config generate --projectId=projectId --clusterName="cluster-1,cluster-2...cluster-N" --includeSecrets --targetNamespace=my-namespace.
func GenerateBuilder() *cobra.Command {
	const use = "generate"
	opts := &GenerateOpts{}

	cmd := &cobra.Command{
		Use:     use,
		Args:    require.NoArgs,
		Aliases: cli.GenerateAliases(use),
		Short:   "Generate Kubernetes configuration resources.",
		Long: "This command provides your Kubernetes configuration access to Atlas. " +
			"You can generate Atlas Operator resources for Atlas objects, including Projects, Deployments, and Users.",
		Example: `# Export Project, DatabaseUsers resources for a specific project without connection and integration secrets:
  atlas kubernetes config generate --projectId=<projectId>

  # Export Project, DatabaseUsers resources for a specific project including connection and integration secrets:
  atlas kubernetes config generate --projectId=<projectId> --includeSecrets

  # Export Project, DatabaseUsers resources for a specific project including connection and integrations secrets to a specific namespace
  atlas kubernetes config generate --projectId=<projectId> --includeSecrets --targetNamespace=<namespace>
  
  # Export Project, DatabaseUsers, Deployments resources for a specific project including connection and integrations secrets to a specific namespace
  atlas kubernetes config generate --projectId=<projectId> --clusterName=<cluster-name-1, cluster-name-2> --includeSecrets --targetNamespace=<namespace>`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidateOrgID,
				opts.ValidateTargetNamespace,
				opts.initStores(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringSliceVar(&opts.clusterName, flag.ClusterName, []string{}, usage.ExporterClusterName)
	cmd.Flags().BoolVar(&opts.includeSecrets, flag.OperatorIncludeSecrets, false, usage.OperatorIncludeSecrets)
	cmd.Flags().StringVar(&opts.targetNamespace, flag.OperatorTargetNamespace, "", usage.OperatorTargetNamespace)

	return cmd
}
