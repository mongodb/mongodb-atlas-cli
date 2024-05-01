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

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/crds"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/validation"
)

var ErrUnsupportedOperatorVersionFmt = "version %q is not supported. Supported versions: %v"

type GenerateOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	clusterName        []string
	dataFederationName []string
	includeSecrets     bool
	targetNamespace    string
	operatorVersion    string
	store              store.OperatorGenericStore
	credsStore         store.CredentialsGetter
	crdsProvider       crds.AtlasOperatorCRDProvider
}

func (opts *GenerateOpts) ValidateTargetNamespace() error {
	if errs := validation.IsDNS1123Label(opts.targetNamespace); len(errs) != 0 {
		return fmt.Errorf("%s parameter is invalid: %v", flag.OperatorTargetNamespace, errs)
	}
	return nil
}

func (opts *GenerateOpts) ValidateOperatorVersion() error {
	if _, versionFound := features.GetResourcesForVersion(opts.operatorVersion); versionFound {
		return nil
	}
	return fmt.Errorf(ErrUnsupportedOperatorVersionFmt, opts.operatorVersion, features.SupportedVersions())
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

		opts.crdsProvider = crds.NewGithubAtlasCRDProvider()

		return nil
	}
}

func (opts *GenerateOpts) Run() error {
	atlasCRDs, err := features.NewAtlasCRDs(opts.crdsProvider, opts.operatorVersion)
	if err != nil {
		return err
	}
	result, err := operator.NewConfigExporter(opts.store, opts.credsStore, opts.ProjectID, opts.OrgID).
		WithClustersNames(opts.clusterName).
		WithTargetNamespace(opts.targetNamespace).
		WithSecretsData(opts.includeSecrets).
		WithTargetOperatorVersion(opts.operatorVersion).
		WithFeatureValidator(atlasCRDs).
		WithPatcher(atlasCRDs).
		WithDataFederationNames(opts.dataFederationName).
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
		Short:   "Generate Kubernetes configuration resources for use with Atlas Kubernetes Operator.",
		Long:    `This command exports configurations for Atlas objects including projects, deployments, and users in a Kubernetes-compatible format, allowing you to manage these resources using the Atlas Kubernetes Operator. For more information, see https://www.mongodb.com/docs/atlas/atlas-operator/`,
		Example: `# Export Project, DatabaseUsers, Deployments resources for a specific project without connection and integration secrets:
  atlas kubernetes config generate --projectId=<projectId>

  # Export Project, DatabaseUsers, Deployments resources for a specific project including connection and integration secrets:
  atlas kubernetes config generate --projectId=<projectId> --includeSecrets

  # Export Project, DatabaseUsers, Deployments resources for a specific project including connection and integration secrets to a specific namespace:
  atlas kubernetes config generate --projectId=<projectId> --includeSecrets --targetNamespace=<namespace>

  # Export Project, DatabaseUsers, DataFederations and specific Deployment resources for a specific project including connection and integration secrets to a specific namespace:
  atlas kubernetes config generate --projectId=<projectId> --clusterName=<cluster-name-1, cluster-name-2> --includeSecrets --targetNamespace=<namespace>

  # Export resources for a specific version of the Atlas Kubernetes Operator:
  atlas kubernetes config generate --projectId=<projectId> --targetNamespace=<namespace> --operatorVersion=1.5.1

  # Export Project, DatabaseUsers, Clusters and specific DataFederation resources for a specific project to a specific namespace:
  atlas kubernetes config generate --projectId=<projectId> --dataFederationName=<data-federation-name-1, data-federation-name-2> --targetNamespace=<namespace>`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidateTargetNamespace,
				opts.ValidateOperatorVersion,
				opts.initStores(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	cmd.Flags().StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	cmd.Flags().StringSliceVar(&opts.clusterName, flag.ClusterName, []string{}, usage.ExporterClusterName)
	cmd.Flags().BoolVar(&opts.includeSecrets, flag.OperatorIncludeSecrets, false, usage.OperatorIncludeSecrets)
	cmd.Flags().StringVar(&opts.targetNamespace, flag.OperatorTargetNamespace, "", usage.OperatorTargetNamespace)
	cmd.Flags().StringVar(&opts.operatorVersion, flag.OperatorVersion, features.LatestOperatorMajorVersion, usage.OperatorVersion)
	cmd.Flags().StringSliceVar(&opts.dataFederationName, flag.DataFederationName, []string{}, usage.ExporterDataFederationName)

	return cmd
}
