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

package operator

import (
	"context"
	"fmt"

	"github.com/google/go-github/v50/github"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/crds"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/version"
	"github.com/mongodb/mongodb-atlas-cli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/validation"
)

const defaultInstallNamespace = "default"

type InstallOpts struct {
	cli.GlobalOpts
	cli.OutputOpts

	versionProvider version.AtlasOperatorVersionProvider

	operatorVersion string
	targetNamespace string
	watchNamespace  []string
	projectName     string
	importResources bool
	KubeConfig      string
	KubeContext     string
}

func (opts *InstallOpts) defaults() error {
	if opts.operatorVersion == "" {
		latest, err := opts.versionProvider.GetLatest()
		if err != nil {
			return err
		}

		opts.operatorVersion = latest
	}

	if opts.targetNamespace == "" {
		opts.targetNamespace = defaultInstallNamespace
	}

	return nil
}

func (opts *InstallOpts) ValidateTargetNamespace() error {
	if errs := validation.IsDNS1123Label(opts.targetNamespace); len(errs) != 0 {
		return fmt.Errorf("%s parameter is invalid: %v", flag.OperatorTargetNamespace, errs)
	}

	return nil
}

func (opts *InstallOpts) ValidateOperatorVersion() error {
	isSupported, err := opts.versionProvider.IsSupported(opts.operatorVersion)
	if err != nil {
		return err
	}

	if !isSupported {
		return fmt.Errorf("version %s is not supported", opts.operatorVersion)
	}

	return nil
}

func (opts *InstallOpts) ValidateWatchNamespace() error {
	for _, ns := range opts.watchNamespace {
		if errs := validation.IsDNS1123Label(ns); len(errs) != 0 {
			return fmt.Errorf("item %s of %s parameter is invalid: %v", ns, flag.OperatorWatchNamespace, errs)
		}
	}

	return nil
}

func (opts *InstallOpts) Run(ctx context.Context) error {
	kubeCtl, err := kubernetes.NewKubeCtl(opts.KubeConfig, opts.KubeContext)
	if err != nil {
		return err
	}

	installer := operator.NewInstaller(opts.versionProvider, kubeCtl)

	profile := config.Default()
	atlasStore, err := store.New(store.AuthenticatedPreset(profile), store.WithContext(ctx))
	if err != nil {
		return err
	}

	credStore := profile

	crdVersion, err := features.CRDCompatibleVersion(opts.operatorVersion)
	if err != nil {
		return err
	}

	featureValidator, err := features.NewAtlasCRDs(crds.NewGithubAtlasCRDProvider(), crdVersion)
	if err != nil {
		return err
	}

	err = operator.NewInstall(installer, atlasStore, credStore, featureValidator, kubeCtl, opts.operatorVersion).
		WithNamespace(opts.targetNamespace).
		WithWatchNamespaces(opts.watchNamespace).
		WithWatchProjectName(opts.projectName).
		WithImportResources(opts.importResources).
		Run(ctx, opts.OrgID)

	if err != nil {
		return err
	}

	return opts.Print("Atlas Kubernetes Operator installed successfully")
}

func InstallBuilder() *cobra.Command {
	const use = "install"
	opts := &InstallOpts{}

	cmd := &cobra.Command{
		Use:     use,
		Args:    require.NoArgs,
		Aliases: cli.GenerateAliases(use),
		Short:   "Install Atlas Kubernetes Operator to a cluster.",
		Long:    `This command installs one of the supported versions of Atlas Kubernetes Operator to an existing cluster, as well as automatically import Atlas resources to be managed by the operator.`,
		Example: `# Install latest version of the operator into the default namespace:
  atlas kubernetes operator install

  # Install an specific version of the operator:
  atlas kubernetes operator install --operatorVersion=1.7.0

  # Install an specific version of the operator to a namespace and watch only this namespace and a second one
  atlas kubernetes operator install --operatorVersion=1.7.0 --targetNamespace=<namespace> --watchNamespace=<namespace>,<secondNamespace>`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			opts.versionProvider = version.NewOperatorVersion(github.NewClient(nil))

			return opts.PreRunE(
				opts.defaults,
				opts.ValidateOrgID,
				opts.ValidateOperatorVersion,
				opts.ValidateTargetNamespace,
				opts.ValidateWatchNamespace,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run(cmd.Context())
		},
	}

	flags := cmd.Flags()

	flags.StringVar(&opts.OrgID, flag.OrgID, "", usage.OrgID)
	flags.StringVar(&opts.operatorVersion, flag.OperatorVersion, "", usage.OperatorVersionInstall)
	flags.StringVar(&opts.targetNamespace, flag.OperatorTargetNamespace, "", usage.OperatorTargetNamespaceInstall)
	flags.StringSliceVar(&opts.watchNamespace, flag.OperatorWatchNamespace, []string{}, usage.OperatorWatchNamespace)
	flags.StringVar(&opts.projectName, flag.OperatorProjectName, "", usage.OperatorProjectName)
	flags.BoolVar(&opts.importResources, flag.OperatorImport, false, usage.OperatorImport)
	flags.StringVar(&opts.KubeConfig, flag.KubernetesClusterConfig, "", usage.KubernetesClusterConfig)
	flags.StringVar(&opts.KubeContext, flag.KubernetesClusterContext, "", usage.KubernetesClusterContext)

	return cmd
}
