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

package streamsprocessing

import (
	"fmt"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/secrets"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	testNamespace       = "test"
	testProjectName     = "my-project"
	testOperatorVersion = "2.4.0"
	testInstanceName    = "instance-0"
	testCertificate     = "-----BEGIN CERTIFICATE-----\nMIIEITCCAwmgAwIBAgIUTLX+HHPxjMxw1pOXEu/+m+aXrgIwDQYJKoZIhvcNAQEL\nBQAwgZ8xCzAJBgNVBAYTAkRFMQ8wDQYDVQQIDAZCZXJsaW4xDzANBgNVBAcMBkJl\ncmxpbjEVMBMGA1UECgwMTW9uZ29EQiBHbWJoMRMwEQYDVQQLDApLdWJlcm5ldGVz\nMRcwFQYDVQQDDA5BdGxhcyBPcGVyYXRvcjEpMCcGCSqGSIb3DQEJARYaaGVsZGVy\nLnNhbnRhbmFAbW9uZ29kYi5jb20wHhcNMjQwNDIzMTE0NzI2WhcNMjcwMTE4MTE0\nNzI2WjCBnzELMAkGA1UEBhMCREUxDzANBgNVBAgMBkJlcmxpbjEPMA0GA1UEBwwG\nQmVybGluMRUwEwYDVQQKDAxNb25nb0RCIEdtYmgxEzARBgNVBAsMCkt1YmVybmV0\nZXMxFzAVBgNVBAMMDkF0bGFzIE9wZXJhdG9yMSkwJwYJKoZIhvcNAQkBFhpoZWxk\nZXIuc2FudGFuYUBtb25nb2RiLmNvbTCCASIwDQYJKoZIhvcNAQEBBQADggEPADCC\nAQoCggEBAKoBtN0V9F8ZnbPJMKDZ0jHRw35Y/jtZpdN6z824nyRh4U4FeLaAOzex\nEiHrxDt9IccxKcVc/9WAq7Pn1C42YJFy9dgLSD94TW4lJwLhAsGxI5bVy+ls6c3u\ncpiPzaoUU1vx+Gg5ob+UefjAf7WxaRnuSiUpYPVVueZ218Hhc1W8yajfwLdshXiN\nNaBox2Pu+ofsq5aM1T4MARsLODUJqzoQHR2275oFPNaz2BgBgRUDkICw+RPfjQ0X\nlCkCtHy2QeBb5hGOi0lG89C9lbuEXb5YOzGG4Cc6snZGf21MGxXAXiL/KsBZrP5i\nedABbwkXEgLk41OcwNgshuADM7iOd9sCAwEAAaNTMFEwHQYDVR0OBBYEFBiwIuyh\n3sqgzfcgKb80FF1WByAIMB8GA1UdIwQYMBaAFBiwIuyh3sqgzfcgKb80FF1WByAI\nMA8GA1UdEwEB/wQFMAMBAf8wDQYJKoZIhvcNAQELBQADggEBAB0iWV/hpK1WuxjS\nh5HAfRxBCyWFIU14S7tQHTPuQANQAh3Zktkghpmc6hdNb3VjKzVUSTv9Ye6V22mh\nResf7PVWFvOdPoiJnmJjUQ5W3FUVZWOgx3rFlKO/5HOi5wRvBDyuZsTjIEJP5MOl\n3lBs17FOVqM3iT785oabOEj/8LhkvdG9brobG8oAttUSPChiYbEtH83WqgeHnCWI\nreLAKIvG8bFVaokdInEgoRt5uque70g0tqAje9MXqCodB96Lo1tk8yyvX4jWI2Pb\npe7aAzw79hIH3tyw+FHjZLgHAq77E14xBxMxvamSnsqGhvCkb7pRHD5+l4tg2k/N\nYJZC5C0=\n-----END CERTIFICATE-----\n"
)

func TestBuildAtlasStreamKafkaConnection(t *testing.T) {
	t.Run("should build Kafka stream connection without SSL config", func(t *testing.T) {
		connection := admin.StreamsConnection{
			Name: pointer.Get("kafka-config"),
			Type: pointer.Get("Kafka"),
			Authentication: &admin.StreamsKafkaAuthentication{
				Mechanism: pointer.Get("SCRAM-SHA512"),
				Username:  pointer.Get("kafka-user"),
			},
			BootstrapServers: pointer.Get("kafka://server1:9001,kafka://server:9002"),
			Config:           pointer.Get(map[string]string{"config": "value"}),
			Security: &admin.StreamsKafkaSecurity{
				Protocol: pointer.Get("PLAINTEXT"),
			},
		}
		resource, resourceSecrets := buildAtlasStreamKafkaConnection(
			testNamespace,
			testOperatorVersion,
			testProjectName,
			testInstanceName,
			&connection,
			map[string]string{},
		)

		expectedResource := akov2.AtlasStreamConnection{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamConnection",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connection.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamConnectionSpec{
				Name:           "kafka-config",
				ConnectionType: "Kafka",
				KafkaConfig: &akov2.StreamsKafkaConnection{
					Authentication: akov2.StreamsKafkaAuthentication{
						Mechanism: "SCRAM-SHA512",
						Credentials: akov2common.ResourceRefNamespaced{
							Name:      fmt.Sprintf("%s-%s-%s-userpass", testProjectName, testInstanceName, connection.GetName()),
							Namespace: testNamespace,
						},
					},
					BootstrapServers: "kafka://server1:9001,kafka://server:9002",
					Security: akov2.StreamsKafkaSecurity{
						Protocol: "PLAINTEXT",
					},
					Config: map[string]string{"config": "value"},
				},
			},
			Status: akov2status.AtlasStreamConnectionStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}
		expectedSecrets := []*corev1.Secret{
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s-userpass", testProjectName, testInstanceName, connection.GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{secrets.UsernameField: []byte("kafka-user"), secrets.PasswordField: []byte("")},
			},
		}

		assert.Equal(t, expectedResource, *resource)
		assert.Equal(t, expectedSecrets, resourceSecrets)
	})

	t.Run("should build Kafka stream connection with SSL config", func(t *testing.T) {
		connection := admin.StreamsConnection{
			Name: pointer.Get("kafka-config"),
			Type: pointer.Get("Kafka"),
			Authentication: &admin.StreamsKafkaAuthentication{
				Mechanism: pointer.Get("SCRAM-SHA512"),
				Username:  pointer.Get("kafka-user"),
			},
			BootstrapServers: pointer.Get("kafka://server1:9001,kafka://server:9002"),
			Config:           pointer.Get(map[string]string{"config": "value"}),
			Security: &admin.StreamsKafkaSecurity{
				Protocol:                pointer.Get("SSL"),
				BrokerPublicCertificate: pointer.Get(testCertificate),
			},
		}
		resource, resourceSecrets := buildAtlasStreamKafkaConnection(
			testNamespace,
			testOperatorVersion,
			testProjectName,
			testInstanceName,
			&connection,
			map[string]string{},
		)

		expectedResource := akov2.AtlasStreamConnection{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamConnection",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connection.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamConnectionSpec{
				Name:           "kafka-config",
				ConnectionType: "Kafka",
				KafkaConfig: &akov2.StreamsKafkaConnection{
					Authentication: akov2.StreamsKafkaAuthentication{
						Mechanism: "SCRAM-SHA512",
						Credentials: akov2common.ResourceRefNamespaced{
							Name:      fmt.Sprintf("%s-%s-%s-userpass", testProjectName, testInstanceName, connection.GetName()),
							Namespace: testNamespace,
						},
					},
					BootstrapServers: "kafka://server1:9001,kafka://server:9002",
					Security: akov2.StreamsKafkaSecurity{
						Protocol: "SSL",
						Certificate: akov2common.ResourceRefNamespaced{
							Name:      fmt.Sprintf("%s-%s-%s-certificate", testProjectName, testInstanceName, connection.GetName()),
							Namespace: testNamespace,
						},
					},
					Config: map[string]string{"config": "value"},
				},
			},
			Status: akov2status.AtlasStreamConnectionStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}
		expectedSecrets := []*corev1.Secret{
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s-userpass", testProjectName, testInstanceName, connection.GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{secrets.UsernameField: []byte("kafka-user"), secrets.PasswordField: []byte("")},
			},
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s-certificate", testProjectName, testInstanceName, connection.GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{secrets.CertificateField: []byte(testCertificate)},
			},
		}

		assert.Equal(t, expectedResource, *resource)
		assert.Equal(t, expectedSecrets, resourceSecrets)
	})
}

func TestBuildAtlasStreamClusterConnection(t *testing.T) {
	t.Run("should build cluster stream connection", func(t *testing.T) {
		connection := &admin.StreamsConnection{
			Name:        pointer.Get("cluster-config"),
			Type:        pointer.Get("Cluster"),
			ClusterName: pointer.Get("my-cluster"),
			DbRoleToExecute: &admin.DBRoleToExecute{
				Role: pointer.Get("readWrite"),
				Type: pointer.Get("BUILT_IN"),
			},
		}
		expectedResource := akov2.AtlasStreamConnection{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamConnection",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connection.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamConnectionSpec{
				Name:           "cluster-config",
				ConnectionType: "Cluster",
				ClusterConfig: &akov2.ClusterConnectionConfig{
					Name: "my-cluster",
					Role: akov2.StreamsClusterDBRole{
						Name:     "readWrite",
						RoleType: "BUILT_IN",
					},
				},
			},
			Status: akov2status.AtlasStreamConnectionStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		assert.Equal(
			t,
			expectedResource,
			*buildAtlasStreamClusterConnection(testNamespace, testOperatorVersion, testProjectName, testInstanceName, connection, map[string]string{}),
		)
	})
}

func TestBuildAtlasStreamSampleConnection(t *testing.T) {
	t.Run("should build sample stream connection", func(t *testing.T) {
		connection := &admin.StreamsConnection{
			Name: pointer.Get("sample-config"),
			Type: pointer.Get("Sample"),
		}
		expectedResource := akov2.AtlasStreamConnection{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamConnection",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connection.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamConnectionSpec{
				Name:           "sample-config",
				ConnectionType: "Sample",
			},
			Status: akov2status.AtlasStreamConnectionStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		assert.Equal(
			t,
			expectedResource,
			*buildAtlasStreamSampleConnection(testNamespace, testOperatorVersion, testProjectName, testInstanceName, connection, map[string]string{}),
		)
	})
}

func TestBuildAtlasStreamConnection(t *testing.T) {
	t.Run("should return error when building an unsupported type", func(t *testing.T) {
		connection := &admin.StreamsConnection{
			Name: pointer.Get("config"),
			Type: pointer.Get("RabbitMQ"),
		}

		conn, sec, err := buildAtlasStreamConnection(testNamespace, testOperatorVersion, testProjectName, testInstanceName, connection, map[string]string{})
		require.ErrorContains(t, err, "trying to generate an unsupported connection type")
		assert.Nil(t, conn)
		assert.Nil(t, sec)
	})

	t.Run("should build sample stream connection", func(t *testing.T) {
		connection := &admin.StreamsConnection{
			Name: pointer.Get("sample-config"),
			Type: pointer.Get("Sample"),
		}
		expectedResource := akov2.AtlasStreamConnection{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamConnection",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connection.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamConnectionSpec{
				Name:           "sample-config",
				ConnectionType: "Sample",
			},
			Status: akov2status.AtlasStreamConnectionStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		conn, sec, err := buildAtlasStreamConnection(testNamespace, testOperatorVersion, testProjectName, testInstanceName, connection, map[string]string{})
		require.NoError(t, err)
		assert.Nil(t, sec)
		assert.Equal(t, expectedResource, *conn)
	})

	t.Run("should build cluster stream connection", func(t *testing.T) {
		connection := &admin.StreamsConnection{
			Name:        pointer.Get("cluster-config"),
			Type:        pointer.Get("Cluster"),
			ClusterName: pointer.Get("my-cluster"),
			DbRoleToExecute: &admin.DBRoleToExecute{
				Role: pointer.Get("readWrite"),
				Type: pointer.Get("BUILT_IN"),
			},
		}
		expectedResource := akov2.AtlasStreamConnection{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamConnection",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connection.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamConnectionSpec{
				Name:           "cluster-config",
				ConnectionType: "Cluster",
				ClusterConfig: &akov2.ClusterConnectionConfig{
					Name: "my-cluster",
					Role: akov2.StreamsClusterDBRole{
						Name:     "readWrite",
						RoleType: "BUILT_IN",
					},
				},
			},
			Status: akov2status.AtlasStreamConnectionStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		conn, sec, err := buildAtlasStreamConnection(testNamespace, testOperatorVersion, testProjectName, testInstanceName, connection, map[string]string{})
		require.NoError(t, err)
		assert.Nil(t, sec)
		assert.Equal(t, expectedResource, *conn)
	})

	t.Run("should build Kafka stream connection", func(t *testing.T) {
		connection := &admin.StreamsConnection{
			Name: pointer.Get("kafka-config"),
			Type: pointer.Get("Kafka"),
			Authentication: &admin.StreamsKafkaAuthentication{
				Mechanism: pointer.Get("SCRAM-SHA512"),
				Username:  pointer.Get("kafka-user"),
			},
			BootstrapServers: pointer.Get("kafka://server1:9001,kafka://server:9002"),
			Config:           pointer.Get(map[string]string{"config": "value"}),
			Security: &admin.StreamsKafkaSecurity{
				Protocol:                pointer.Get("SSL"),
				BrokerPublicCertificate: pointer.Get(testCertificate),
			},
		}

		expectedResource := akov2.AtlasStreamConnection{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamConnection",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connection.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamConnectionSpec{
				Name:           "kafka-config",
				ConnectionType: "Kafka",
				KafkaConfig: &akov2.StreamsKafkaConnection{
					Authentication: akov2.StreamsKafkaAuthentication{
						Mechanism: "SCRAM-SHA512",
						Credentials: akov2common.ResourceRefNamespaced{
							Name:      fmt.Sprintf("%s-%s-%s-userpass", testProjectName, testInstanceName, connection.GetName()),
							Namespace: testNamespace,
						},
					},
					BootstrapServers: "kafka://server1:9001,kafka://server:9002",
					Security: akov2.StreamsKafkaSecurity{
						Protocol: "SSL",
						Certificate: akov2common.ResourceRefNamespaced{
							Name:      fmt.Sprintf("%s-%s-%s-certificate", testProjectName, testInstanceName, connection.GetName()),
							Namespace: testNamespace,
						},
					},
					Config: map[string]string{"config": "value"},
				},
			},
			Status: akov2status.AtlasStreamConnectionStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}
		expectedSecrets := []*corev1.Secret{
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s-userpass", testProjectName, testInstanceName, connection.GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{secrets.UsernameField: []byte("kafka-user"), secrets.PasswordField: []byte("")},
			},
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s-certificate", testProjectName, testInstanceName, connection.GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{secrets.CertificateField: []byte(testCertificate)},
			},
		}

		conn, sec, err := buildAtlasStreamConnection(testNamespace, testOperatorVersion, testProjectName, testInstanceName, connection, map[string]string{})
		require.NoError(t, err)
		assert.Equal(t, expectedSecrets, sec)
		assert.Equal(t, expectedResource, *conn)
	})
}

func TestBuildAtlasStreamInstance(t *testing.T) {
	t.Run("should build stream instance", func(t *testing.T) {
		instance := &admin.StreamsTenant{
			Id:   pointer.Get("instance-0-id"),
			Name: pointer.Get(testInstanceName),
			DataProcessRegion: &admin.StreamsDataProcessRegion{
				CloudProvider: "AWS",
				Region:        "IRL_DUBLIN",
			},
			StreamConfig: &admin.StreamConfig{
				Tier: pointer.Get("SP30"),
			},
			Hostnames: &[]string{"server1", "server2"},
			GroupId:   pointer.Get("my-project-id"),
		}
		connections := []*akov2.AtlasStreamConnection{
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "connection1",
					Namespace: testNamespace,
				},
			},
			{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "connection2",
					Namespace: testNamespace,
				},
			},
		}
		expectedResource := akov2.AtlasStreamInstance{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamInstance",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s", testProjectName, instance.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamInstanceSpec{
				Name: testInstanceName,
				Config: akov2.Config{
					Provider: "AWS",
					Region:   "IRL_DUBLIN",
					Tier:     "SP30",
				},
				Project: akov2common.ResourceRefNamespaced{
					Name:      testProjectName,
					Namespace: testNamespace,
				},
				ConnectionRegistry: []akov2common.ResourceRefNamespaced{
					{
						Name:      "connection1",
						Namespace: testNamespace,
					},
					{
						Name:      "connection2",
						Namespace: testNamespace,
					},
				},
			},
			Status: akov2status.AtlasStreamInstanceStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		assert.Equal(
			t,
			expectedResource,
			*buildAtlasStreamInstance(testNamespace, testOperatorVersion, testProjectName, instance, connections, map[string]string{}),
		)
	})
}

func TestBuildAtlasStreamsProcessing(t *testing.T) {
	t.Run("should build stream processing resources", func(t *testing.T) {
		connections := []admin.StreamsConnection{
			{
				Name: pointer.Get("kafka-config"),
				Type: pointer.Get("Kafka"),
				Authentication: &admin.StreamsKafkaAuthentication{
					Mechanism: pointer.Get("SCRAM-SHA512"),
					Username:  pointer.Get("kafka-user"),
				},
				BootstrapServers: pointer.Get("kafka://server1:9001,kafka://server:9002"),
				Config:           pointer.Get(map[string]string{"config": "value"}),
				Security: &admin.StreamsKafkaSecurity{
					Protocol:                pointer.Get("SSL"),
					BrokerPublicCertificate: pointer.Get(testCertificate),
				},
			},
			{
				Name:        pointer.Get("cluster-config"),
				Type:        pointer.Get("Cluster"),
				ClusterName: pointer.Get("my-cluster"),
				DbRoleToExecute: &admin.DBRoleToExecute{
					Role: pointer.Get("readWrite"),
					Type: pointer.Get("BUILT_IN"),
				},
			},
		}
		expectedConnectionsResources := []*akov2.AtlasStreamConnection{
			{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "atlas.mongodb.com/v1",
					Kind:       "AtlasStreamConnection",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connections[0].GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						features.ResourceVersion: testOperatorVersion,
					},
				},
				Spec: akov2.AtlasStreamConnectionSpec{
					Name:           "kafka-config",
					ConnectionType: "Kafka",
					KafkaConfig: &akov2.StreamsKafkaConnection{
						Authentication: akov2.StreamsKafkaAuthentication{
							Mechanism: "SCRAM-SHA512",
							Credentials: akov2common.ResourceRefNamespaced{
								Name:      fmt.Sprintf("%s-%s-%s-userpass", testProjectName, testInstanceName, connections[0].GetName()),
								Namespace: testNamespace,
							},
						},
						BootstrapServers: "kafka://server1:9001,kafka://server:9002",
						Security: akov2.StreamsKafkaSecurity{
							Protocol: "SSL",
							Certificate: akov2common.ResourceRefNamespaced{
								Name:      fmt.Sprintf("%s-%s-%s-certificate", testProjectName, testInstanceName, connections[0].GetName()),
								Namespace: testNamespace,
							},
						},
						Config: map[string]string{"config": "value"},
					},
				},
				Status: akov2status.AtlasStreamConnectionStatus{
					Common: akoapi.Common{
						Conditions: []akoapi.Condition{},
					},
				},
			},
			{
				TypeMeta: metav1.TypeMeta{
					APIVersion: "atlas.mongodb.com/v1",
					Kind:       "AtlasStreamConnection",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connections[1].GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						features.ResourceVersion: testOperatorVersion,
					},
				},
				Spec: akov2.AtlasStreamConnectionSpec{
					Name:           "cluster-config",
					ConnectionType: "Cluster",
					ClusterConfig: &akov2.ClusterConnectionConfig{
						Name: "my-cluster",
						Role: akov2.StreamsClusterDBRole{
							Name:     "readWrite",
							RoleType: "BUILT_IN",
						},
					},
				},
				Status: akov2status.AtlasStreamConnectionStatus{
					Common: akoapi.Common{
						Conditions: []akoapi.Condition{},
					},
				},
			},
		}
		expectedSecrets := []*corev1.Secret{
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s-userpass", testProjectName, testInstanceName, connections[0].GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{secrets.UsernameField: []byte("kafka-user"), secrets.PasswordField: []byte("")},
			},
			{
				TypeMeta: metav1.TypeMeta{
					Kind:       "Secret",
					APIVersion: "v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      fmt.Sprintf("%s-%s-%s-certificate", testProjectName, testInstanceName, connections[0].GetName()),
					Namespace: testNamespace,
					Labels: map[string]string{
						secrets.TypeLabelKey: secrets.CredLabelVal,
					},
				},
				Data: map[string][]byte{secrets.CertificateField: []byte(testCertificate)},
			},
		}

		instance := &admin.StreamsTenant{
			Id:   pointer.Get("instance-0-id"),
			Name: pointer.Get(testInstanceName),
			DataProcessRegion: &admin.StreamsDataProcessRegion{
				CloudProvider: "AWS",
				Region:        "IRL_DUBLIN",
			},
			StreamConfig: &admin.StreamConfig{
				Tier: pointer.Get("SP30"),
			},
			Hostnames: &[]string{"server1", "server2"},
			GroupId:   pointer.Get("my-project-id"),
		}
		expectedInstanceResource := akov2.AtlasStreamInstance{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "atlas.mongodb.com/v1",
				Kind:       "AtlasStreamInstance",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-%s", testProjectName, instance.GetName()),
				Namespace: testNamespace,
				Labels: map[string]string{
					features.ResourceVersion: testOperatorVersion,
				},
			},
			Spec: akov2.AtlasStreamInstanceSpec{
				Name: testInstanceName,
				Config: akov2.Config{
					Provider: "AWS",
					Region:   "IRL_DUBLIN",
					Tier:     "SP30",
				},
				Project: akov2common.ResourceRefNamespaced{
					Name:      testProjectName,
					Namespace: testNamespace,
				},
				ConnectionRegistry: []akov2common.ResourceRefNamespaced{
					{
						Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connections[0].GetName()),
						Namespace: testNamespace,
					},
					{
						Name:      fmt.Sprintf("%s-%s-%s", testProjectName, testInstanceName, connections[1].GetName()),
						Namespace: testNamespace,
					},
				},
			},
			Status: akov2status.AtlasStreamInstanceStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		instanceResource, connectionsResources, secretsResource, err := BuildAtlasStreamsProcessing(
			testNamespace,
			testOperatorVersion,
			testProjectName,
			instance,
			connections,
			map[string]string{},
		)
		require.NoError(t, err)
		assert.Equal(t, expectedInstanceResource, *instanceResource)
		assert.Equal(t, expectedConnectionsResources, connectionsResources)
		assert.Equal(t, expectedSecrets, secretsResource)
	})

	t.Run("should return error when unable to build stream processing resources", func(t *testing.T) {
		connections := []admin.StreamsConnection{
			{
				Name: pointer.Get("rabbit-mq-config"),
				Type: pointer.Get("RabbitMQ"),
			},
		}
		instance := &admin.StreamsTenant{
			Id:   pointer.Get("instance-0-id"),
			Name: pointer.Get(testInstanceName),
			DataProcessRegion: &admin.StreamsDataProcessRegion{
				CloudProvider: "AWS",
				Region:        "IRL_DUBLIN",
			},
			StreamConfig: &admin.StreamConfig{
				Tier: pointer.Get("SP30"),
			},
			Hostnames: &[]string{"server1", "server2"},
			GroupId:   pointer.Get("my-project-id"),
		}

		instanceResource, connectionsResources, secretsResource, err := BuildAtlasStreamsProcessing(
			testNamespace,
			testOperatorVersion,
			testProjectName,
			instance,
			connections,
			map[string]string{},
		)
		require.ErrorContains(t, err, "trying to generate an unsupported connection type")
		assert.Nil(t, instanceResource)
		assert.Nil(t, connectionsResources)
		assert.Nil(t, secretsResource)
	})
}
