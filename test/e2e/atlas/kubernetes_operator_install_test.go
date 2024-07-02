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

//go:build e2e || (atlas && cluster && kubernetes && local)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	operatorNamespace     = "atlas-operator"
	maxAttempts           = 12
	deploymentMaxAttempts = 36
	poolInterval          = 10 * time.Second
)

func TestKubernetesOperatorInstall(t *testing.T) {
	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)
	const contextPrefix = "kind-"
	t.Run("should failed to install old and not supported version of the operator", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--operatorVersion", "1.1.0")
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		require.Error(t, inErr, string(resp))
		assert.Equal(t, "Error: version 1.1.0 is not supported\n", string(resp))
	})

	t.Run("should failed to install a non-existing version of the operator", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--operatorVersion", "100.0.0")
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		require.Error(t, inErr, string(resp))
		assert.Equal(t, "Error: version 100.0.0 is not supported\n", string(resp))
	})

	t.Run("should failed when unable to setup connection to the cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--kubeconfig", "/path/to/non/existing/config")
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		require.Error(t, inErr, string(resp))
		assert.Equal(t, "Error: unable to prepare client configuration: invalid configuration: no configuration has been provided, try setting KUBERNETES_MASTER environment variable\n", string(resp))
	})

	t.Run("should install operator with default options", func(t *testing.T) {
		clusterName := "install-default"
		operator := setupCluster(t, clusterName)
		context := contextPrefix + clusterName

		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--kubeContext", context)
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, inErr, string(resp))
		assert.Equal(t, "Atlas Kubernetes Operator installed successfully\n", string(resp))

		checkDeployment(t, operator, "default")
	})

	t.Run("should install latest major version of operator in its own namespace with cluster-wide config", func(t *testing.T) {
		clusterName := "install-clusterwide"
		operator := setupCluster(t, clusterName, operatorNamespace)
		context := contextPrefix + clusterName

		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--operatorVersion", features.LatestOperatorMajorVersion,
			"--targetNamespace", operatorNamespace,
			"--kubeContext", context)
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, inErr, string(resp))
		assert.Equal(t, "Atlas Kubernetes Operator installed successfully\n", string(resp))

		checkDeployment(t, operator, operatorNamespace)
	})

	t.Run("should install latest major version of operator in its own namespace with namespaced config", func(t *testing.T) {
		clusterName := "single-namespace"
		operatorWatch1 := "atlas-watch1"
		operatorWatch2 := "atlas-watch2"
		operator := setupCluster(t, clusterName, operatorNamespace, operatorWatch1, operatorWatch2)
		context := contextPrefix + clusterName

		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--operatorVersion", features.LatestOperatorMajorVersion,
			"--targetNamespace", operatorNamespace,
			"--watchNamespace", operatorWatch1,
			"--watchNamespace", operatorWatch2,
			"--kubeContext", context)
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		req.NoError(inErr, string(resp))
		assert.Equal(t, "Atlas Kubernetes Operator installed successfully\n", string(resp))
		checkDeployment(t, operator, operatorNamespace)
	})

	t.Run("should install latest major version of operator in a single namespaced config", func(t *testing.T) {
		clusterName := "install-namespaced"
		operator := setupCluster(t, clusterName, operatorNamespace)
		context := contextPrefix + clusterName

		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--operatorVersion", features.LatestOperatorMajorVersion,
			"--targetNamespace", operatorNamespace,
			"--watchNamespace", operatorNamespace,
			"--kubeContext", context)
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		req.NoError(inErr, string(resp))
		assert.Equal(t, "Atlas Kubernetes Operator installed successfully\n", string(resp))

		checkDeployment(t, operator, operatorNamespace)
	})

	t.Run("should install operator starting a new project", func(t *testing.T) {
		clusterName := "install-new-project"
		operator := setupCluster(t, clusterName, operatorNamespace)
		context := contextPrefix + clusterName
		projectName := "MyK8sProject"

		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--targetNamespace", operatorNamespace,
			"--projectName", projectName,
			"--import",
			"--kubeContext", context)
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, inErr, string(resp))
		assert.Equal(t, "Atlas Kubernetes Operator installed successfully\n", string(resp))

		checkDeployment(t, operator, operatorNamespace)

		projectSecret := &corev1.Secret{}
		inErr = operator.getK8sObject(
			client.ObjectKey{Name: fmt.Sprintf("mongodb-atlas-%s-api-key", prepareK8sName(projectName)), Namespace: operatorNamespace},
			projectSecret,
			false,
		)
		require.NoError(t, inErr)

		orgSecret := &corev1.Secret{}
		inErr = operator.getK8sObject(
			client.ObjectKey{Name: "mongodb-atlas-operator-api-key", Namespace: operatorNamespace},
			orgSecret,
			false,
		)
		require.Error(t, inErr)

		checkK8sAtlasProject(t, operator, client.ObjectKey{Name: prepareK8sName(projectName), Namespace: operatorNamespace})

		akoProject := &akov2.AtlasProject{}
		err = operator.getK8sObject(
			client.ObjectKey{Name: prepareK8sName(projectName), Namespace: operatorNamespace},
			akoProject,
			true,
		)
		req.NoError(err)
		req.NoError(operator.deleteK8sObject(akoProject))

		projectDeleted := false
		for i := 0; i < maxAttempts; i++ {
			err = operator.getK8sObject(
				client.ObjectKey{Name: prepareK8sName(projectName), Namespace: operatorNamespace},
				akoProject,
				true,
			)

			if err != nil {
				projectDeleted = true
				break
			}

			time.Sleep(poolInterval)
		}

		if !projectDeleted {
			t.Errorf("project %s was not cleaned up", projectName)
		}

		cleanUpKeys(t, operator, operatorNamespace, cliPath)
	})

	t.Run("should install operator importing atlas existing resources", func(t *testing.T) {
		g := newAtlasE2ETestGenerator(t)
		g.enableBackup = true
		g.generateProject("k8sOperatorInstall")
		g.generateCluster()

		clusterName := "install-import"
		operator := setupCluster(t, clusterName, operatorNamespace)
		context := contextPrefix + clusterName

		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--targetNamespace", operatorNamespace,
			"--projectName", g.projectName,
			"--import",
			"--kubeContext", context)
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		req.NoError(inErr, string(resp))
		assert.Equal(t, "Atlas Kubernetes Operator installed successfully\n", string(resp))

		checkDeployment(t, operator, operatorNamespace)
		checkK8sAtlasProject(t, operator, client.ObjectKey{Name: prepareK8sName(g.projectName), Namespace: operatorNamespace})
		checkK8sAtlasDeployment(t, operator, client.ObjectKey{Name: prepareK8sName(fmt.Sprintf("%s-%s", g.projectName, g.clusterName)), Namespace: operatorNamespace})

		cleanUpKeys(t, operator, operatorNamespace, cliPath)
	})

	t.Run("should install operator with deletion protection and sub resource protection disabled", func(t *testing.T) {
		g := newAtlasE2ETestGenerator(t)
		g.enableBackup = true
		g.generateProject("k8sOperatorInstall")
		g.generateCluster()

		clusterName := "install-import-without-dp"
		operator := setupCluster(t, clusterName, operatorNamespace)
		context := contextPrefix + clusterName

		cmd := exec.Command(cliPath,
			"kubernetes",
			"operator",
			"install",
			"--resourceDeletionProtection=false",
			"--subresourceDeletionProtection=false",
			"--targetNamespace", operatorNamespace,
			"--kubeContext", context)
		cmd.Env = os.Environ()
		resp, inErr := e2e.RunAndGetStdOut(cmd)
		require.NoError(t, inErr, string(resp))
		assert.Equal(t, "Atlas Kubernetes Operator installed successfully\n", string(resp))

		checkDeployment(t, operator, operatorNamespace)

		deployment := &appsv1.Deployment{}
		require.NoError(t, operator.getK8sObject(
			client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: operatorNamespace},
			deployment,
			false,
		))
		assert.Contains(t, deployment.Spec.Template.Spec.Containers[0].Args, "--resourceDeletionProtection=false")
		assert.Contains(t, deployment.Spec.Template.Spec.Containers[0].Args, "--subresourceDeletionProtection=false")

		cleanUpKeys(t, operator, operatorNamespace, cliPath)
	})
}

func checkDeployment(t *testing.T, operator *operatorHelper, namespace string) {
	t.Helper()

	deployment := &appsv1.Deployment{}

	var deploymentReady bool

	for i := 0; i < maxAttempts; i++ {
		deploymentReady = true

		err := operator.getK8sObject(
			client.ObjectKey{Name: "mongodb-atlas-operator", Namespace: namespace},
			deployment,
			false,
		)
		require.NoError(t, err)

		for _, condition := range deployment.Status.Conditions {
			if condition.Status != corev1.ConditionTrue {
				deploymentReady = false
			}
		}

		if deployment.Status.ReadyReplicas != 1 {
			deploymentReady = false
		}

		if deploymentReady {
			break
		}

		time.Sleep(poolInterval)
	}

	if !deploymentReady {
		t.Error("operator install failed: deployment is not ready")
	}

	var podReady bool

	for i := 0; i < maxAttempts; i++ {
		podReady = true

		pods, err := operator.getPodFromDeployment(deployment)
		require.NoError(t, err)

		if len(pods) == 1 {
			pod := pods[0]

			for _, condition := range pod.Status.Conditions {
				if condition.Status != corev1.ConditionTrue {
					podReady = false
				}
			}

			if podReady {
				break
			}
		}

		time.Sleep(poolInterval)
	}

	if !podReady {
		t.Error("operator install failed: pod is not ready")
	}
}

func checkK8sAtlasProject(t *testing.T, operator *operatorHelper, key client.ObjectKey) {
	t.Helper()

	var ready bool
	project := &akov2.AtlasProject{}

	for i := 0; i < maxAttempts; i++ {
		ready = true

		err := operator.getK8sObject(key, project, false)
		require.NoError(t, err)

		for _, condition := range project.Status.Conditions {
			if condition.Status != corev1.ConditionTrue {
				ready = false
			}
		}

		if ready {
			break
		}

		time.Sleep(poolInterval)
	}

	if !ready {
		t.Error("import resources failed: project is not ready")
	}
}

func checkK8sAtlasDeployment(t *testing.T, operator *operatorHelper, key client.ObjectKey) {
	t.Helper()

	var ready bool
	deployment := &akov2.AtlasDeployment{}

	for i := 0; i < deploymentMaxAttempts; i++ {
		ready = true

		err := operator.getK8sObject(key, deployment, false)
		require.NoError(t, err)

		for _, condition := range deployment.Status.Conditions {
			if condition.Status != corev1.ConditionTrue {
				ready = false
			}
		}

		if ready {
			break
		}

		time.Sleep(poolInterval)
	}

	if !ready {
		t.Error("import resources failed: deployment is not ready")
	}
}

func cleanUpKeys(t *testing.T, operator *operatorHelper, namespace string, cliPath string) {
	t.Helper()

	secrets, err := operator.getOperatorSecretes(namespace)
	require.NoError(t, err)

	toDelete := map[string]struct{}{}
	for _, secret := range secrets {
		toDelete[secret.Name] = struct{}{}
	}

	cmd := exec.Command(cliPath,
		orgEntity,
		"apiKeys",
		"ls",
		"-o=json")
	cmd.Env = os.Environ()
	resp, err := e2e.RunAndGetStdOut(cmd)
	require.NoError(t, err, string(resp))

	var keys atlasv2.PaginatedApiApiUser
	err = json.Unmarshal(resp, &keys)
	require.NoError(t, err)

	for _, key := range keys.GetResults() {
		keyID := *key.Id
		desc := *key.Desc

		if _, ok := toDelete[desc]; ok {
			cmd = exec.Command(cliPath,
				orgEntity,
				"apiKeys",
				"rm",
				keyID,
				"--force")
			cmd.Env = os.Environ()
			_, err = e2e.RunAndGetStdOut(cmd)
			require.NoError(t, err)
		}
	}
}
