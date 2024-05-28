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

package config

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

const containerImage = "mongodb/mongodb-atlas-kubernetes-operator"

type ApplyOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	GenerateOpts

	KubeConfig  string
	KubeContext string
}

func (opts *ApplyOpts) ValidateTargetNamespace() error {
	if opts.targetNamespace != "" {
		return opts.GenerateOpts.ValidateTargetNamespace()
	}

	return nil
}

func (opts *ApplyOpts) ValidateOperatorVersion() error {
	if opts.operatorVersion != "" {
		return opts.GenerateOpts.ValidateOperatorVersion()
	}

	return nil
}

func (opts *ApplyOpts) autoDetectParams(kubeCtl *kubernetes.KubeCtl) error {
	if opts.targetNamespace != "" && opts.operatorVersion != "" {
		return nil
	}

	operatorDeployment, err := kubeCtl.FindAtlasOperator(context.Background())
	if err != nil {
		return fmt.Errorf("unable to auto detect params: %w", err)
	}

	if opts.targetNamespace == "" {
		opts.targetNamespace = operatorDeployment.Namespace
	}

	if opts.operatorVersion == "" {
		if !strings.HasPrefix(operatorDeployment.Spec.Template.Spec.Containers[0].Image, containerImage+":") {
			return errors.New("unable to auto detect operator version. you should explicitly set operator version if you are running an openshift certified installation")
		}

		opts.operatorVersion = getOperatorMajorVersion(operatorDeployment.Spec.Template.Spec.Containers[0].Image)
	}

	return nil
}

func (opts *ApplyOpts) Run() error {
	kubeCtl, err := kubernetes.NewKubeCtl(opts.KubeConfig, opts.KubeContext)
	if err != nil {
		return err
	}

	err = opts.autoDetectParams(kubeCtl)
	if err != nil {
		return err
	}

	atlasCRDs, err := features.NewAtlasCRDs(opts.crdsProvider, opts.operatorVersion)
	if err != nil {
		return err
	}

	exporter := operator.NewConfigExporter(opts.store, opts.credsStore, opts.ProjectID, opts.OrgID).
		WithClustersNames(opts.clusterName).
		WithTargetNamespace(opts.targetNamespace).
		WithTargetOperatorVersion(opts.operatorVersion).
		WithSecretsData(true).
		WithFeatureValidator(atlasCRDs).
		WithPatcher(atlasCRDs)
	err = operator.NewConfigApply(
		operator.NewConfigApplyParams{
			OrgID:     opts.OrgID,
			ProjectID: opts.ProjectID,
			KubeCtl:   kubeCtl,
			Exporter:  exporter,
		},
	).WithTargetOperatorVersion(opts.operatorVersion).
		WithNamespace(opts.targetNamespace).
		Run()

	if err != nil {
		return err
	}

	return opts.Print("Atlas Resources exported and applied to Kubernetes cluster successfully")
}

// ApplyBuilder builds a cobra.Command that can run as:
// atlas kubernetes config apply --orgId=orgId --projectId=projectId --clusterName="cluster-1,cluster-2...cluster-N" --targetNamespace=my-namespace.
func ApplyBuilder() *cobra.Command {
	const use = "apply"
	opts := &ApplyOpts{}

	cmd := &cobra.Command{
		Use:     use,
		Args:    require.NoArgs,
		Aliases: cli.GenerateAliases(use),
		Short:   "Generate and apply Kubernetes configuration resources for use with Atlas Kubernetes Operator.",
		Long:    `This command exports configurations for Atlas objects including projects, deployments, and users directly into Kubernetes, allowing you to manage these resources using the Atlas Kubernetes Operator. For more information, see https://www.mongodb.com/docs/atlas/atlas-operator/.`,
		Example: `# Export and apply all supported resources of a specific project:
  atlas kubernetes config apply --projectId=<projectId>

  # Export and apply all supported resources of a specific project and to a specific namespace:
  atlas kubernetes config apply --projectId=<projectId> --targetNamespace=<namespace>
  
  # Export and apply all supported Project resource, and only the described Deployment resources of a specific project to a specific namespace:
  atlas kubernetes config apply --projectId=<projectId> --clusterName=<cluster-name-1, cluster-name-2> --targetNamespace=<namespace>

  # Export and apply all supported resources of a specific project to a specific namespace restricting the version of the Atlas Kubernetes Operator:
  atlas kubernetes config apply --projectId=<projectId> --targetNamespace=<namespace> --operatorVersion=1.5.1`,
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

	flags := cmd.Flags()

	flags.StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	flags.StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	flags.StringSliceVar(&opts.clusterName, flag.ClusterName, []string{}, usage.ExporterClusterName)
	flags.StringVar(&opts.targetNamespace, flag.OperatorTargetNamespace, "", usage.OperatorTargetNamespace)
	flags.StringVar(&opts.operatorVersion, flag.OperatorVersion, "", usage.OperatorVersion)
	flags.StringVar(&opts.KubeConfig, flag.KubernetesClusterConfig, "", usage.KubernetesClusterConfig)
	flags.StringVar(&opts.KubeContext, flag.KubernetesClusterContext, "", usage.KubernetesClusterContext)

	return cmd
}

func getOperatorMajorVersion(image string) string {
	version := strings.TrimPrefix(image, containerImage+":")

	semVer := strings.Split(version, ".")
	semVer[2] = "0"

	return strings.Join(semVer, ".")
}
