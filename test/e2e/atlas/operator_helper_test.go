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
//go:build e2e || atlas

package atlas_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2toptr "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/util/toptr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
	"sigs.k8s.io/kind/pkg/cluster"
)

type operatorHelper struct {
	t         *testing.T
	k8sClient client.Client

	resourcesTracked []client.Object
	deployment       *appsv1.Deployment
}

func newOperatorHelper(t *testing.T) (*operatorHelper, error) {
	t.Helper()

	cfg, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	err = akov2.AddToScheme(scheme.Scheme)
	if err != nil {
		return nil, err
	}

	k8sClient, err := client.New(cfg, client.Options{Scheme: scheme.Scheme})
	if err != nil {
		return nil, err
	}

	return &operatorHelper{
		t:         t,
		k8sClient: k8sClient,
	}, nil
}

func createK8SCluster(name string) error {
	clusterConfig := &v1alpha4.Cluster{
		Networking: v1alpha4.Networking{
			IPFamily: v1alpha4.IPv4Family,
		},
	}

	provider := cluster.NewProvider(cluster.ProviderWithDocker())
	return provider.Create(
		name,
		cluster.CreateWithV1Alpha4Config(clusterConfig),
		cluster.CreateWithWaitForReady(1*time.Minute),
		cluster.CreateWithDisplayUsage(false),
		cluster.CreateWithDisplaySalutation(false),
	)
}

func deleteK8SCluster(name string) error {
	provider := cluster.NewProvider(cluster.ProviderWithDocker())
	return provider.Delete(name, "")
}

func (oh *operatorHelper) getK8sObject(key client.ObjectKey, object client.Object, track bool) error {
	err := oh.k8sClient.Get(context.Background(), key, object, &client.GetOptions{})
	if err != nil {
		return err
	}

	if track {
		oh.resourcesTracked = append(oh.resourcesTracked, object)
	}

	return nil
}

func (oh *operatorHelper) createK8sObject(object client.Object) error {
	return oh.k8sClient.Create(context.Background(), object, &client.CreateOptions{})
}

func (oh *operatorHelper) deleteK8sObject(object client.Object) error {
	return oh.k8sClient.Delete(context.Background(), object, &client.DeleteOptions{})
}

func (oh *operatorHelper) getPodFromDeployment(deployment *appsv1.Deployment) ([]corev1.Pod, error) {
	podList := &corev1.PodList{}
	err := oh.k8sClient.List(
		context.Background(),
		podList,
		&client.ListOptions{
			Namespace:     deployment.Namespace,
			LabelSelector: labels.SelectorFromSet(deployment.Labels),
		},
	)
	if err != nil {
		return nil, err
	}

	if len(podList.Items) == 0 {
		return nil, errors.New("pod not found")
	}

	return podList.Items, nil
}

func (oh *operatorHelper) getOperatorSecretes(namespace string) ([]corev1.Secret, error) {
	secretList := &corev1.SecretList{}
	err := oh.k8sClient.List(
		context.Background(),
		secretList,
		&client.ListOptions{
			Namespace:     namespace,
			LabelSelector: labels.SelectorFromSet(map[string]string{"atlas.mongodb.com/type": "credentials"}),
		},
	)
	if err != nil {
		return nil, err
	}

	return secretList.Items, nil
}

func (oh *operatorHelper) stopOperator() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	deployment.Spec.Replicas = akov2toptr.MakePtr(int32(0))

	err = oh.k8sClient.Update(context.Background(), &deployment, &client.UpdateOptions{})
	if err != nil {
		oh.t.Errorf("unable to stop operator: %v", err)
	}
}

func (oh *operatorHelper) startOperator() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	deployment.Spec.Replicas = akov2toptr.MakePtr(int32(1))

	err = oh.k8sClient.Update(context.Background(), &deployment, &client.UpdateOptions{})
	if err != nil {
		oh.t.Errorf("unable to start operator: %v", err)
	}
}

func (oh *operatorHelper) deleteOperator() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	oh.deployment = &deployment

	err = oh.k8sClient.Delete(context.Background(), &deployment, &client.DeleteOptions{})
	if err != nil {
		oh.t.Errorf("unable to delete operator: %v", err)
	}
}

func (oh *operatorHelper) restoreOperator() {
	if oh.deployment == nil {
		oh.t.Errorf("unable to restore operator. unknown previous state")
	}

	oh.deployment.ResourceVersion = ""
	err := oh.k8sClient.Create(context.Background(), oh.deployment, &client.CreateOptions{})
	if err != nil {
		oh.t.Errorf("unable to restore operator: %v", err)
	}
}

func (oh *operatorHelper) emulateCertifiedOperator() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	container := deployment.Spec.Template.Spec.Containers[0]
	container.Image = "quay.io/" + container.Image
	deployment.Spec.Template.Spec.Containers[0] = container

	err = oh.k8sClient.Update(context.Background(), &deployment, &client.UpdateOptions{})
	if err != nil {
		oh.t.Errorf("unable to emulate certified operator: %v", err)
	}
}

func (oh *operatorHelper) restoreOperatorImage() {
	deployment := appsv1.Deployment{}
	err := oh.getK8sObject(
		client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: "mongodb-atlas-system"},
		&deployment,
		false,
	)
	if err != nil {
		oh.t.Errorf("unable to retrieve operator deployment: %v", err)
	}

	container := deployment.Spec.Template.Spec.Containers[0]
	container.Image = strings.Trim(container.Image, "quay.io/")
	deployment.Spec.Template.Spec.Containers[0] = container

	err = oh.k8sClient.Update(context.Background(), &deployment, &client.UpdateOptions{})
	if err != nil {
		oh.t.Errorf("unable to restore operator image: %v", err)
	}
}

func (oh *operatorHelper) cleanUpResources() {
	for _, object := range oh.resourcesTracked {
		if len(object.GetFinalizers()) > 0 {
			object.SetFinalizers([]string{})

			err := oh.k8sClient.Update(context.Background(), object, &client.UpdateOptions{})
			if err != nil {
				oh.t.Errorf("unable to update k8s resource: %v", err)
			}
		}

		err := oh.k8sClient.Delete(context.Background(), object)
		if err != nil {
			oh.t.Errorf("unable to delete k8s resource: %v", err)
		}
	}
}
