// Copyright 2024 MongoDB Inc
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
//go:build e2e || (iam && flp && atlas)

package atlas_test

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/test/e2e"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes/scheme"
)

const tnm = "importer-namespace"
const pjprefix = "Kubernetes-"

type KubPjSuite struct {
	generator       *atlasE2ETestGenerator
	expectedProject *akov2.AtlasProject
	cliPath         string
}

var expLabels = map[string]string{
	features.ResourceVersion: features.LatestOperatorMajorVersion,
}

func getK8SEntitie(data []byte) ([]runtime.Object, error) {
	b := bufio.NewReader(bytes.NewReader(data))
	r := yaml.NewYAMLReader(b)

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
func TIP(t *testing.T) {

	req := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	req.NoError(err)

	var federationSettingsID string
	var oidcIWorkforceIdpID string

	t.Run("Describe an org federation settings", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			"describe",
			"-o=json",
		)

		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))

		var settings atlasv2.OrgFederationSettings
		req.NoError(json.Unmarshal(resp, &settings))

		a := assert.New(t)
		a.NotEmpty(settings.GetId())
		a.NotEmpty(settings.GetIdentityProviderStatus())
		federationSettingsID = settings.GetId()
	})

	t.Run("Connect OIDC IdP WORKFORCE", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			federatedAuthenticationEntity,
			federationSettingsEntity,
			connectedOrgsConfigsEntity,
			"connect",
			"--identityProviderId",
			oidcIWorkforceIdpID,
			"--federationSettingsId",
			federationSettingsID,
			"-o=json",
		)
		s := STP(t)
		cliPath := s.cliPath
		generator := s.generator
		cmd.Env = os.Environ()
		resp, err := e2e.RunAndGetStdOut(cmd)
		req.NoError(err, string(resp))
		expectedProject := s.expectedProject
		var config atlasv2.ConnectedOrgConfig
		req.NoError(json.Unmarshal(resp, &config))

		assert.NotEmpty(t, config.DataAccessIdentityProviderIds)
		assert.Contains(t, config.GetDataAccessIdentityProviderIds(), oidcIWorkforceIdpID)

		secondCmd := exec.Command(cliPath,
			"kubernetes",
			"config",
			"generate",
			"--projectId",
			generator.projectID,
			"--orgId", "", // Empty org id does not make it fail
			"--targetNamespace",
			tnm,
			"--includeSecrets")
		cmd.Env = os.Environ()

		secondResp, err := e2e.RunAndGetStdOut(secondCmd)
		t.Log(string(secondResp))
		require.NoError(t, err, string(secondResp))

		var objects []runtime.Object
		objects, err = getK8SEntitie(secondResp)
		require.NoError(t, err, "should not fail on decode")
		require.NotEmpty(t, objects)

		checkPj(t, objects, expectedProject)
		secret, found := findSc(objects)
		require.True(t, found, "Secret is not found in results")
		assert.Equal(t, tnm, secret.Namespace)
	})
}

func checkPj(t *testing.T, output []runtime.Object, expected *akov2.AtlasProject) {
	t.Helper()
	found := false
	var p *akov2.AtlasProject
	var ok bool
	for i := range output {
		p, ok = output[i].(*akov2.AtlasProject)
		if ok {
			found = true
			break
		}
	}
	require.True(t, found, "AtlasProject is not found in results")

	// secretref names are randomly generated so we can't determine those in forehand
	expected.Spec.EncryptionAtRest.AwsKms = p.Spec.EncryptionAtRest.AwsKms
	expected.Spec.EncryptionAtRest.GoogleCloudKms = p.Spec.EncryptionAtRest.GoogleCloudKms
	expected.Spec.EncryptionAtRest.AzureKeyVault = p.Spec.EncryptionAtRest.AzureKeyVault

	for i := range p.Spec.AlertConfigurations {
		alertConfig := &p.Spec.AlertConfigurations[i]
		for j := range alertConfig.Notifications {
			expected.Spec.AlertConfigurations[i].Notifications[j].APITokenRef = p.Spec.AlertConfigurations[i].Notifications[j].APITokenRef
			expected.Spec.AlertConfigurations[i].Notifications[j].DatadogAPIKeyRef = p.Spec.AlertConfigurations[i].Notifications[j].DatadogAPIKeyRef
			expected.Spec.AlertConfigurations[i].Notifications[j].OpsGenieAPIKeyRef = p.Spec.AlertConfigurations[i].Notifications[j].OpsGenieAPIKeyRef
			expected.Spec.AlertConfigurations[i].Notifications[j].ServiceKeyRef = p.Spec.AlertConfigurations[i].Notifications[j].ServiceKeyRef
			expected.Spec.AlertConfigurations[i].Notifications[j].VictorOpsSecretRef = p.Spec.AlertConfigurations[i].Notifications[j].VictorOpsSecretRef
		}
		for k := range alertConfig.Matchers {
			expected.Spec.AlertConfigurations[i].Matchers[k].FieldName = p.Spec.AlertConfigurations[i].Matchers[k].FieldName
			expected.Spec.AlertConfigurations[i].Matchers[k].Operator = p.Spec.AlertConfigurations[i].Matchers[k].Operator
			expected.Spec.AlertConfigurations[i].Matchers[k].Value = p.Spec.AlertConfigurations[i].Matchers[k].Value
		}
	}

	assert.Equal(t, expected, p)
}
func findSc(objects []runtime.Object) (*corev1.Secret, bool) {
	for i := range objects {
		if secret, ok := objects[i].(*corev1.Secret); ok {
			return secret, ok
		}
	}
	return nil, false
}
func STP(t *testing.T) KubPjSuite {
	t.Helper()
	s := KubPjSuite{}
	s.generator = newAtlasE2ETestGenerator(t)
	s.generator.generateEmptyProject(pjprefix + s.generator.projectName)
	s.expectedProject = refPj(s.generator.projectName, tnm, expLabels)

	cliPath, err := e2e.AtlasCLIBin()
	require.NoError(t, err)
	s.cliPath = cliPath

	// always register atlas entities
	require.NoError(t, akov2.AddToScheme(scheme.Scheme))
	return s
}

func refPj(name, namespace string, labels map[string]string) *akov2.AtlasProject {
	dictionary := resources.AtlasNameToKubernetesName()
	return &akov2.AtlasProject{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasProject",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(name, dictionary),
			Namespace: namespace,
			Labels:    labels,
		},
		Status: akov2status.AtlasProjectStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
		Spec: akov2.AtlasProjectSpec{
			Name: name,
			ConnectionSecret: &akov2common.ResourceRefNamespaced{
				Name: resources.NormalizeAtlasName(name+"-credentials", dictionary),
			},
			Settings: &akov2.ProjectSettings{
				IsCollectDatabaseSpecificsStatisticsEnabled: pointer.Get(true),
				IsDataExplorerEnabled:                       pointer.Get(true),
				IsPerformanceAdvisorEnabled:                 pointer.Get(true),
				IsRealtimePerformancePanelEnabled:           pointer.Get(true),
				IsSchemaAdvisorEnabled:                      pointer.Get(true),
			},
			Auditing: &akov2.Auditing{
				AuditAuthorizationSuccess: false,
				Enabled:                   false,
			},
			EncryptionAtRest: &akov2.EncryptionAtRest{
				AwsKms: akov2.AwsKms{
					Enabled: pointer.Get(false),
					Valid:   pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(name+"-aws-credentials", dictionary),
						Namespace: namespace,
					},
				},
				AzureKeyVault: akov2.AzureKeyVault{
					Enabled: pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(name+"-azure-credentials", dictionary),
						Namespace: namespace,
					},
				},
				GoogleCloudKms: akov2.GoogleCloudKms{
					Enabled: pointer.Get(false),
					SecretRef: akov2common.ResourceRefNamespaced{
						Name:      resources.NormalizeAtlasName(name+"-gcp-credentials", dictionary),
						Namespace: namespace,
					},
				},
			},
		},
	}
}
