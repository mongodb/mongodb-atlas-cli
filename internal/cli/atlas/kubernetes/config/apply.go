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

	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator"
	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/internal/usage"
	akov1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/spf13/cobra"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func (opts *ApplyOpts) autoDetectParams(k8sClient client.Client) error {
	if opts.targetNamespace != "" && opts.operatorVersion != "" {
		return nil
	}

	operatorDeployment, err := findOperatorInstallation(k8sClient)
	if err != nil {
		return fmt.Errorf("unable to auto detect params: %w", err)
	}

	if opts.targetNamespace == "" {
		opts.targetNamespace = operatorDeployment.Namespace
	}

	if opts.operatorVersion == "" {
		if !strings.HasPrefix(operatorDeployment.Spec.Template.Spec.Containers[0].Image, fmt.Sprintf("%s:", containerImage)) {
			return errors.New("unable to auto detect operator version. you should explicitly set operator version if you are running an openshift certified installation")
		}

		opts.operatorVersion = getOperatorMajorVersion(operatorDeployment.Spec.Template.Spec.Containers[0].Image)
	}

	return nil
}

func (opts *ApplyOpts) Run() error {
	kubeConfig, err := loadKubeConfig(opts.KubeConfig)
	if err != nil {
		return err
	}

	kubeClient, err := newK8sClient(kubeConfig, opts.KubeConfig)
	if err != nil {
		return err
	}

	err = opts.autoDetectParams(kubeClient)
	if err != nil {
		return err
	}

	featureValidator, err := features.NewAtlasCRDs(opts.crdsProvider, opts.operatorVersion)
	if err != nil {
		return err
	}

	exporter := operator.NewConfigExporter(opts.store, opts.credsStore, opts.ProjectID, opts.OrgID)
	err = operator.NewConfigApply(
		operator.NewConfigApplyParams{
			OrgID:     opts.OrgID,
			ProjectID: opts.ProjectID,
			K8sClient: kubeClient,
			Exporter:  exporter,
			Validator: featureValidator,
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
		Short:   "Apply Kubernetes configuration resources to a cluster.",
		Long:    `This command provides your Kubernetes configuration access to Atlas. You can export and apply Atlas Operator resources for Atlas objects, including Projects, Deployments, and Users into your Kubernetes cluster.`,
		Example: `# Export and apply all supported resources of a specific project:
  atlas kubernetes config apply --projectId=<projectId>

  # Export and apply all supported resources of a specific project and to a specific namespace:
  atlas kubernetes config apply --projectId=<projectId> --includeSecrets --targetNamespace=<namespace>
  
  # Export and apply all supported Project resource, and only the described Deployment resources of a specific project to a specific namespace:
  atlas kubernetes config apply --projectId=<projectId> --clusterName=<cluster-name-1, cluster-name-2> --includeSecrets --targetNamespace=<namespace>

  #Export and apply all supported resources of a specific project to a specific namespace restricting the version of the Atlas Kubernetes Operator:
  atlas kubernetes config apply --projectId=<projectId> --targetNamespace=<namespace> --operatorVersion=1.5.1`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidateOrgID,
				opts.ValidateTargetNamespace,
				opts.ValidateOperatorVersion,
				opts.initStores(cmd.Context()),
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
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

func loadKubeConfig(kubeConfig string) (*api.Config, error) {
	pathOptions := clientcmd.NewDefaultPathOptions()

	if kubeConfig != "" {
		pathOptions.LoadingRules.ExplicitPath = kubeConfig
	}

	config, err := pathOptions.GetStartingConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to load kubernetes configuration: %w", err)
	}

	return config, nil
}

func newK8sClient(config *api.Config, kubeContext string) (client.Client, error) {
	configOverrides := &clientcmd.ConfigOverrides{}

	if kubeContext != "" {
		configOverrides.CurrentContext = kubeContext
	}

	clientConfig := clientcmd.NewDefaultClientConfig(*config, configOverrides)

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to prepare client configuration: %w", err)
	}

	err = akov1.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	k8sClient, err := client.New(restConfig, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return nil, fmt.Errorf("unable to setup kubernetes client: %w", err)
	}

	return k8sClient, nil
}

func findOperatorInstallation(k8sClient client.Client) (*appsv1.Deployment, error) {
	namespaces := corev1.NamespaceList{}
	err := k8sClient.List(context.Background(), &namespaces, &client.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("unable to find operator installation. failed to list cluster namespaces: %w", err)
	}

	listOptions := client.ListOptions{
		LabelSelector: labels.SelectorFromSet(
			map[string]string{
				"app.kubernetes.io/component": "controller",
				"app.kubernetes.io/instance":  "mongodb-atlas-kubernetes-operator",
				"app.kubernetes.io/name":      "mongodb-atlas-kubernetes-operator",
			},
		),
	}

	for _, namespace := range namespaces.Items {
		deployments := appsv1.DeploymentList{}
		err := k8sClient.List(context.Background(), &deployments, &listOptions)
		if err != nil {
			_, _ = log.Warningf("failed to look into namespace %s: %v", namespace.GetName(), err)
		}

		if len(deployments.Items) > 0 {
			if err != nil {
				return nil, err
			}

			return &deployments.Items[0], nil
		}
	}

	return nil, errors.New("couldn't to find operator installed in any accessible namespace")
}

func getOperatorMajorVersion(image string) string {
	version := strings.Trim(image, fmt.Sprintf("%s:", containerImage))

	semVer := strings.Split(version, ".")
	semVer[2] = "0"

	return strings.Join(semVer, ".")
}
