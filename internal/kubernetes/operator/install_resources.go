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
	"io"
	"os"
	"strings"

	"github.com/andreaangiolillo/mongocli-test/internal/kubernetes"
	"github.com/andreaangiolillo/mongocli-test/internal/kubernetes/operator/resources"
	"github.com/andreaangiolillo/mongocli-test/internal/kubernetes/operator/version"
	"gopkg.in/yaml.v3"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apisv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

const (
	leaderElectionRoleName        = "mongodb-atlas-leader-election-role"
	leaderElectionRoleBindingName = "mongodb-atlas-leader-election-rolebinding"
	installationTargetClusterWide = "clusterwide"
	installationTargetNamespaced  = "namespaced"
	credentialsGlobalName         = "mongodb-atlas-operator-api-key" //nolint:gosec
	credentialsProjectScopedName  = "mongodb-atlas-%s-api-key"       //nolint:gosec
)

type Installer interface {
	InstallCRDs(ctx context.Context, version string, namespaced bool) error
	InstallConfiguration(ctx context.Context, version, namespace string, watch []string, atlasGov bool) error
	InstallCredentials(ctx context.Context, namespace, orgID, publicKey, privateKey string, projectName string) error
}

type InstallResources struct {
	versionProvider version.AtlasOperatorVersionProvider
	kubeCtl         *kubernetes.KubeCtl
	objConverter    runtime.UnstructuredConverter
}

func (ir *InstallResources) InstallCRDs(ctx context.Context, version string, namespaced bool) error {
	target := installationTargetClusterWide

	if namespaced {
		target = installationTargetNamespaced
	}

	data, err := ir.versionProvider.DownloadResource(ctx, version, fmt.Sprintf("deploy/%s/crds.yaml", target))
	if err != nil {
		return fmt.Errorf("unable to retrieve CRDs from repository: %w", err)
	}

	crdsData, err := parseYaml(data)
	if err != nil {
		return err
	}

	err = apiextensionsv1.AddToScheme(scheme.Scheme)
	if err != nil {
		return fmt.Errorf("unable to handle CRDs: %w", err)
	}

	for _, crdData := range crdsData {
		crd := &apiextensionsv1.CustomResourceDefinition{}
		err = ir.objConverter.FromUnstructured(crdData, crd)
		if err != nil {
			return fmt.Errorf("failed to convert CRD object: %w", err)
		}

		err = ir.kubeCtl.Create(ctx, crd)
		if err != nil {
			return fmt.Errorf("failed to add CRD into cluster: %w", err)
		}
	}

	return nil
}

func (ir *InstallResources) InstallConfiguration(ctx context.Context, version, namespace string, watch []string, atlasGov bool) error {
	target := installationTargetClusterWide

	if len(watch) > 0 {
		target = installationTargetNamespaced
	}

	data, err := ir.versionProvider.DownloadResource(ctx, version, fmt.Sprintf("deploy/%s/%s-config.yaml", target, target))
	if err != nil {
		return fmt.Errorf("unable to retrieve configuration from repository: %w", err)
	}

	configData, err := parseYaml(data)
	if err != nil {
		return err
	}

	for _, config := range configData {
		switch config["kind"] {
		case "ServiceAccount":
			err = ir.addServiceAccount(ctx, config, namespace)
			if err != nil {
				return err
			}
		case "Role":
			err = ir.addRoles(ctx, config, namespace, watch)
			if err != nil {
				return err
			}
		case "ClusterRole":
			err = ir.addClusterRole(ctx, config, namespace)
			if err != nil {
				return err
			}
		case "RoleBinding":
			err = ir.addRoleBindings(ctx, config, namespace, watch)
			if err != nil {
				return err
			}
		case "ClusterRoleBinding":
			err = ir.addClusterRoleBinding(ctx, config, namespace)
			if err != nil {
				return err
			}
		case "Deployment":
			err = ir.addDeployment(ctx, config, namespace, watch, atlasGov)
			if err != nil {
				return err
			}
		default:
			continue
		}
	}

	return nil
}

func (ir *InstallResources) InstallCredentials(ctx context.Context, namespace, orgID, publicKey, privateKey string, projectName string) error {
	name := credentialsGlobalName

	if projectName != "" {
		name = fmt.Sprintf(credentialsProjectScopedName, resources.NormalizeAtlasName(projectName, resources.AtlasNameToKubernetesName()))
	}

	obj := &corev1.Secret{
		ObjectMeta: apisv1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels: map[string]string{
				"atlas.mongodb.com/type": "credentials",
			},
		},
		StringData: map[string]string{
			"orgId":         orgID,
			"publicApiKey":  publicKey,
			"privateApiKey": privateKey,
		},
	}

	err := ir.kubeCtl.Create(ctx, obj)
	if err != nil {
		return fmt.Errorf("failed to add Secret into cluster: %w", err)
	}

	return nil
}

func (ir *InstallResources) addServiceAccount(ctx context.Context, config map[string]interface{}, namespace string) error {
	obj := &corev1.ServiceAccount{}
	err := ir.objConverter.FromUnstructured(config, obj)
	if err != nil {
		return fmt.Errorf("failed to convert ServiceAccount object: %w", err)
	}

	obj.SetNamespace(namespace)

	err = ir.kubeCtl.Create(ctx, obj)
	if err != nil {
		return fmt.Errorf("failed to add ServiceAccount into cluster: %w", err)
	}

	return nil
}

func (ir *InstallResources) addRoles(ctx context.Context, config map[string]interface{}, namespace string, watch []string) error {
	namespaces := map[string]struct{}{namespace: {}}

	if !isLeaderElectionResource(config, leaderElectionRoleName) {
		for _, watchedNamespace := range watch {
			namespaces[watchedNamespace] = struct{}{}
		}
	}

	for watchNamespace := range namespaces {
		obj := &rbacv1.Role{}
		err := ir.objConverter.FromUnstructured(config, obj)
		if err != nil {
			return fmt.Errorf("failed to convert Role object: %w", err)
		}

		obj.SetNamespace(watchNamespace)

		err = ir.kubeCtl.Create(ctx, obj)
		if err != nil {
			return fmt.Errorf("failed to add Role into cluster: %w", err)
		}
	}

	return nil
}

func (ir *InstallResources) addClusterRole(ctx context.Context, config map[string]interface{}, namespace string) error {
	obj := &rbacv1.ClusterRole{}
	err := ir.objConverter.FromUnstructured(config, obj)
	if err != nil {
		return fmt.Errorf("failed to convert ClusterRole object: %w", err)
	}

	obj.SetNamespace(namespace)

	err = ir.kubeCtl.Create(ctx, obj)
	if err != nil {
		return fmt.Errorf("failed to add ClusterRole into cluster: %w", err)
	}

	return nil
}

func (ir *InstallResources) addRoleBindings(ctx context.Context, config map[string]interface{}, namespace string, watch []string) error {
	namespaces := map[string]struct{}{namespace: {}}

	if !isLeaderElectionResource(config, leaderElectionRoleBindingName) {
		for _, watchedNamespace := range watch {
			namespaces[watchedNamespace] = struct{}{}
		}
	}

	for watchNamespace := range namespaces {
		obj := &rbacv1.RoleBinding{}
		err := ir.objConverter.FromUnstructured(config, obj)
		if err != nil {
			return fmt.Errorf("failed to convert RoleBinding object: %w", err)
		}

		obj.SetNamespace(watchNamespace)
		obj.Subjects[0].Namespace = namespace

		err = ir.kubeCtl.Create(ctx, obj)
		if err != nil {
			return fmt.Errorf("failed to add RoleBinding into cluster: %w", err)
		}
	}

	return nil
}

func (ir *InstallResources) addClusterRoleBinding(ctx context.Context, config map[string]interface{}, namespace string) error {
	obj := &rbacv1.ClusterRoleBinding{}
	err := ir.objConverter.FromUnstructured(config, obj)
	if err != nil {
		return fmt.Errorf("failed to convert ClusterRoleBinding object: %w", err)
	}

	obj.SetNamespace(namespace)
	obj.Subjects[0].Namespace = namespace

	err = ir.kubeCtl.Create(ctx, obj)
	if err != nil {
		return fmt.Errorf("failed to add ClusterRoleBinding into cluster: %w", err)
	}

	return nil
}

func (ir *InstallResources) addDeployment(ctx context.Context, config map[string]interface{}, namespace string, watch []string, atlasGov bool) error {
	obj := &appsv1.Deployment{}
	err := ir.objConverter.FromUnstructured(config, obj)
	if err != nil {
		return fmt.Errorf("failed to convert Deployment object: %w", err)
	}

	obj.SetNamespace(namespace)

	if len(obj.Spec.Template.Spec.Containers) > 0 {
		atlasDomain := os.Getenv("MCLI_OPS_MANAGER_URL")
		if atlasDomain == "" {
			atlasDomain = "https://cloud.mongodb.com/"
			if atlasGov {
				atlasDomain = "https://cloud.mongodbgov.com/"
			}
		}
		for i := range obj.Spec.Template.Spec.Containers[0].Args {
			arg := obj.Spec.Template.Spec.Containers[0].Args[i]
			if arg == "--atlas-domain=https://cloud.mongodb.com/" {
				obj.Spec.Template.Spec.Containers[0].Args[i] = fmt.Sprintf("--atlas-domain=%s", atlasDomain)
			}
		}

		if len(watch) > 0 {
			for i := range obj.Spec.Template.Spec.Containers[0].Env {
				env := obj.Spec.Template.Spec.Containers[0].Env[i]

				if env.Name == "WATCH_NAMESPACE" {
					env.ValueFrom = nil
					env.Value = joinNamespaces(namespace, watch)
				}

				obj.Spec.Template.Spec.Containers[0].Env[i] = env
			}
		}
	}

	err = ir.kubeCtl.Create(ctx, obj)
	if err != nil {
		return fmt.Errorf("failed to add Deployment into cluster: %w", err)
	}

	return nil
}

func parseYaml(data io.ReadCloser) ([]map[string]interface{}, error) {
	var k8sResources []map[string]interface{}

	decoder := yaml.NewDecoder(data)

	for {
		obj := map[string]interface{}{}
		err := decoder.Decode(obj)
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, fmt.Errorf("object decode failed: %w", err)
		}

		k8sResources = append(k8sResources, obj)
	}

	return k8sResources, nil
}

func NewInstaller(versionProvider version.AtlasOperatorVersionProvider, kubectl *kubernetes.KubeCtl) *InstallResources {
	return &InstallResources{
		versionProvider: versionProvider,
		kubeCtl:         kubectl,
		objConverter:    runtime.DefaultUnstructuredConverter,
	}
}

func joinNamespaces(namespace string, watched []string) string {
	unique := map[string]struct{}{namespace: {}}
	for _, watch := range watched {
		unique[watch] = struct{}{}
	}

	list := make([]string, 0, len(unique))
	for item := range unique {
		list = append(list, item)
	}

	return strings.Join(list, ",")
}

func isLeaderElectionResource(config map[string]interface{}, leaderElectionID string) bool {
	value, ok := config["metadata"]
	if !ok {
		return false
	}

	metadata, ok := value.(map[string]interface{})
	if !ok {
		return false
	}

	name, ok := metadata["name"]
	if !ok {
		return false
	}

	return name == leaderElectionID
}
