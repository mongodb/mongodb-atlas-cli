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

package kubernetes

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type KubeCtl struct {
	config *api.Config
	client client.Client
}

func (ctl *KubeCtl) FindAtlasOperator(ctx context.Context) (*appsv1.Deployment, error) {
	namespaces := corev1.NamespaceList{}
	err := ctl.client.List(ctx, &namespaces, &client.ListOptions{})
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
		err := ctl.client.List(ctx, &deployments, &listOptions)
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

	return nil, errors.New("couldn't find an operator installed in any accessible namespace")
}

func (ctl *KubeCtl) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return ctl.client.Create(ctx, obj, opts...)
}

func (ctl *KubeCtl) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	return ctl.client.Get(ctx, key, obj, opts...)
}

func (ctl *KubeCtl) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return ctl.client.Update(ctx, obj, opts...)
}

func (ctl *KubeCtl) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return ctl.client.Delete(ctx, obj, opts...)
}

func (ctl *KubeCtl) List(ctx context.Context, obj client.ObjectList, opts ...client.ListOption) error {
	return ctl.client.List(ctx, obj, opts...)
}

func (ctl *KubeCtl) loadConfig(configFile string) error {
	pathOptions := clientcmd.NewDefaultPathOptions()

	if configFile != "" {
		pathOptions.LoadingRules.ExplicitPath = configFile
	}

	config, err := pathOptions.GetStartingConfig()
	if err != nil {
		return fmt.Errorf("unable to load kubernetes configuration: %w", err)
	}

	ctl.config = config

	return nil
}

func (ctl *KubeCtl) setupClient(witContext string) error {
	if ctl.config == nil {
		return errors.New("client configuration was not loaded")
	}

	configOverrides := &clientcmd.ConfigOverrides{}

	if witContext != "" {
		configOverrides.CurrentContext = witContext
	}

	clientConfig := clientcmd.NewDefaultClientConfig(*ctl.config, configOverrides)

	restConfig, err := clientConfig.ClientConfig()
	if err != nil {
		return fmt.Errorf("unable to prepare client configuration: %w", err)
	}

	err = akov2.AddToScheme(scheme.Scheme)
	if err != nil {
		return err
	}

	k8sClient, err := client.New(restConfig, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return fmt.Errorf("unable to setup kubernetes client: %w", err)
	}

	ctl.client = k8sClient

	return nil
}

func NewKubeCtl(fromKubeConfig string, withContext string) (*KubeCtl, error) {
	ctl := &KubeCtl{}

	if err := ctl.loadConfig(fromKubeConfig); err != nil {
		return nil, err
	}

	if err := ctl.setupClient(withContext); err != nil {
		return nil, err
	}

	return ctl, nil
}
