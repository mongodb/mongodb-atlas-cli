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
//go:build e2e || (atlas && cluster && kubernetes)

package atlas_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/internal/kubernetes/operator/pointers"
	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	atlasV1 "github.com/mongodb/mongodb-atlas-kubernetes/pkg/api/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlas "go.mongodb.org/atlas/mongodbatlas"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	k8syaml "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

func getK8SEntities(data []byte) ([]runtime.Object, error) {
	b := bufio.NewReader(bytes.NewReader(data))
	r := k8syaml.NewYAMLReader(b)

	var result []runtime.Object

	for {
		doc, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		d := scheme.Codecs.UniversalDeserializer()
		obj, _, err := d.Decode(doc, nil, nil)
		if err != nil {
			// if document is not a K8S object, skip it
			continue
		}
		if obj != nil {
			result = append(result, obj)
		}
	}
	return result, nil
}

func TestKubernetesConfigGenerate(t *testing.T) {
	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)

	n, err := e2e.RandInt(255)
	require.NoError(t, err)

	targetNamespace := "importer-namespace"

	projectName := fmt.Sprintf("test-project-%s", n)
	clusterReplicaSetName := fmt.Sprintf("test-cluster-replicaset-%s", n)
	clusterServerlessName := fmt.Sprintf("test-cluster-serverless-%s", n)

	helper := NewK8sHelper(t)
	helper.NewProject(&atlas.Project{Name: projectName, OrgID: helper.orgID})
	helper.NewCluster(&atlas.Cluster{
		BiConnector: &atlas.BiConnector{
			Enabled: pointers.MakePtr(true),
		},
		ClusterType:           "REPLICASET",
		DiskSizeGB:            pointers.MakePtr[float64](10),
		GroupID:               helper.project.ID,
		Name:                  clusterReplicaSetName,
		ProviderBackupEnabled: pointers.MakePtr(false),
		ProviderSettings: &atlas.ProviderSettings{
			InstanceSizeName: "M10",
			ProviderName:     "AWS",
			RegionName:       "US_EAST_1",
		},
		ReplicationFactor: pointers.MakePtr[int64](3),
		ReplicationSpecs: []atlas.ReplicationSpec{
			{
				NumShards: pointers.MakePtr[int64](1),
				ZoneName:  "Zone 1",
				RegionsConfig: map[string]atlas.RegionsConfig{
					"US_EAST_1": {
						Priority:       pointers.MakePtr[int64](7),
						AnalyticsNodes: pointers.MakePtr[int64](0),
						ElectableNodes: pointers.MakePtr[int64](3),
						ReadOnlyNodes:  pointers.MakePtr[int64](0),
					},
				},
			},
		},
	})
	helper.NewServerlessInstance(&atlas.ServerlessCreateRequestParams{
		Name: clusterServerlessName,
		ProviderSettings: &atlas.ServerlessProviderSettings{
			BackingProviderName: "AWS",
			ProviderName:        "SERVERLESS",
			RegionName:          "US_EAST_1",
		},
	})

	// always register atlas entities
	require.NoError(t, atlasV1.AddToScheme(scheme.Scheme))

	t.Run("Generate valid resources of ONE project and ONE cluster", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			helper.project.ID,
			"--clusterName",
			clusterReplicaSetName,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))

		a := assert.New(t)
		a.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
		})

		t.Run("Deployment present with valid name", func(t *testing.T) {
			found := false
			var deployment *atlasV1.AtlasDeployment
			var ok bool
			for i := range objects {
				deployment, ok = objects[i].(*atlasV1.AtlasDeployment)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasDeployment is not found in results")
			}
			assert.Equal(t, deployment.Namespace, targetNamespace)
			assert.Equal(t, deployment.Name, clusterReplicaSetName)
		})

		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			found := false
			var secret *corev1.Secret
			var ok bool
			for i := range objects {
				secret, ok = objects[i].(*corev1.Secret)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("Secret is not found in results")
			}
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})

	t.Run("Generate valid resources of ONE project and TWO clusters", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			helper.project.ID,
			"--clusterName",
			fmt.Sprintf("%s,%s", clusterServerlessName, clusterReplicaSetName),
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))

		a := assert.New(t)
		a.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
		})

		t.Run("Deployments present with valid names", func(t *testing.T) {
			var deployments []*atlasV1.AtlasDeployment
			for i := range objects {
				deployment, ok := objects[i].(*atlasV1.AtlasDeployment)
				if ok {
					deployments = append(deployments, deployment)
				}
			}
			clustersCount := len(deployments)
			require.True(t, clustersCount == 2, "result should contain two clusters. actual: ", clustersCount)
			for i := range deployments {
				assert.Equal(t, deployments[i].Namespace, targetNamespace)
			}
			clusterNames := []string{deployments[0].Name, deployments[1].Name}
			assert.Contains(t, clusterNames, clusterReplicaSetName, "result doesn't contain replicaset cluster")
			assert.Contains(t, clusterNames, clusterServerlessName, "result doesn't contain serverless instance")
		})

		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			found := false
			var secret *corev1.Secret
			var ok bool
			for i := range objects {
				secret, ok = objects[i].(*corev1.Secret)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("Secret is not found in results")
			}
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})

	t.Run("Generate valid resources of ONE project and TWO clusters without listing clusters", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			helper.project.ID,
			"--targetNamespace",
			targetNamespace,
			"--includeSecrets")
		cmd.Env = os.Environ()

		resp, err := cmd.CombinedOutput()
		t.Log(string(resp))

		a := assert.New(t)
		a.NoError(err, string(resp))

		var objects []runtime.Object
		t.Run("Output can be decoded", func(t *testing.T) {
			objects, err = getK8SEntities(resp)
			require.NoError(t, err, "should not fail on decode")
			require.True(t, len(objects) > 0, "result should not be empty. got", len(objects))
		})

		t.Run("Project present with valid name", func(t *testing.T) {
			found := false
			var project *atlasV1.AtlasProject
			var ok bool
			for i := range objects {
				project, ok = objects[i].(*atlasV1.AtlasProject)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("AtlasProject is not found in results")
			}
			assert.Equal(t, project.Namespace, targetNamespace)
		})

		t.Run("Deployments present with valid names", func(t *testing.T) {
			var deployments []*atlasV1.AtlasDeployment
			for i := range objects {
				deployment, ok := objects[i].(*atlasV1.AtlasDeployment)
				if ok {
					deployments = append(deployments, deployment)
				}
			}
			clustersCount := len(deployments)
			require.True(t, clustersCount == 2, "result should contain two clusters. actual: ", clustersCount)
			for i := range deployments {
				assert.Equal(t, deployments[i].Namespace, targetNamespace)
			}
			clusterNames := []string{deployments[0].Name, deployments[1].Name}
			assert.Contains(t, clusterNames, clusterReplicaSetName, "result doesn't contain replicaset cluster")
			assert.Contains(t, clusterNames, clusterServerlessName, "result doesn't contain serverless instance")
		})

		t.Run("Connection Secret present with non-empty credentials", func(t *testing.T) {
			found := false
			var secret *corev1.Secret
			var ok bool
			for i := range objects {
				secret, ok = objects[i].(*corev1.Secret)
				if ok {
					found = true
					break
				}
			}
			if !found {
				t.Fatal("Secret is not found in results")
			}
			assert.Equal(t, secret.Namespace, targetNamespace)
		})
	})
}
