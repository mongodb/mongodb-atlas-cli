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
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/secrets"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	connectionTypeSample  = "Sample"
	connectionTypeCluster = "Cluster"
	connectionTypeKafka   = "Kafka"
)

func BuildAtlasStreamsProcessing(
	targetNamespace,
	operatorVersion,
	projectName string,
	instance *admin.StreamsTenant,
	connections []admin.StreamsConnection,
	dictionary map[string]string,
) (*akov2.AtlasStreamInstance, []*akov2.AtlasStreamConnection, []*corev1.Secret, error) {
	akoConnections := make([]*akov2.AtlasStreamConnection, 0, len(connections))
	akoSecrets := make([]*corev1.Secret, 0, len(connections))
	for i := range connections {
		genConnection, genSecrets, err := buildAtlasStreamConnection(
			targetNamespace,
			operatorVersion,
			projectName,
			instance.GetName(),
			&connections[i],
			dictionary,
		)
		if err != nil {
			return nil, nil, nil, err
		}

		akoConnections = append(akoConnections, genConnection)
		akoSecrets = append(akoSecrets, genSecrets...)
	}

	akoInstance := buildAtlasStreamInstance(
		targetNamespace,
		operatorVersion,
		projectName,
		instance,
		akoConnections,
		dictionary,
	)

	return akoInstance, akoConnections, akoSecrets, nil
}

func buildAtlasStreamInstance(
	targetNamespace,
	operatorVersion,
	projectName string,
	instance *admin.StreamsTenant,
	connections []*akov2.AtlasStreamConnection,
	dictionary map[string]string,
) *akov2.AtlasStreamInstance {
	processRegion := instance.GetDataProcessRegion()
	config := instance.GetStreamConfig()
	resource := &akov2.AtlasStreamInstance{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasStreamInstance",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s", projectName, instance.GetName()), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: akov2.AtlasStreamInstanceSpec{
			Name: instance.GetName(),
			Config: akov2.Config{
				Provider: processRegion.GetCloudProvider(),
				Region:   processRegion.GetRegion(),
				Tier:     config.GetTier(),
			},
			Project: akov2common.ResourceRefNamespaced{
				Name:      projectName,
				Namespace: targetNamespace,
			},
			ConnectionRegistry: make([]akov2common.ResourceRefNamespaced, 0, len(connections)),
		},
		Status: akov2status.AtlasStreamInstanceStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}

	for i := range connections {
		resource.Spec.ConnectionRegistry = append(
			resource.Spec.ConnectionRegistry,
			akov2common.ResourceRefNamespaced{Name: connections[i].Name, Namespace: connections[i].Namespace},
		)
	}

	return resource
}

func buildAtlasStreamConnection(
	targetNamespace,
	operatorVersion,
	projectName,
	instanceName string,
	connection *admin.StreamsConnection,
	dictionary map[string]string,
) (*akov2.AtlasStreamConnection, []*corev1.Secret, error) {
	switch connection.GetType() {
	case connectionTypeSample:
		return buildAtlasStreamSampleConnection(
			targetNamespace,
			operatorVersion,
			projectName,
			instanceName,
			connection,
			dictionary,
		), nil, nil
	case connectionTypeCluster:
		return buildAtlasStreamClusterConnection(
			targetNamespace,
			operatorVersion,
			projectName,
			instanceName,
			connection,
			dictionary,
		), nil, nil
	case connectionTypeKafka:
		resource, resourceSecrets := buildAtlasStreamKafkaConnection(
			targetNamespace,
			operatorVersion,
			projectName,
			instanceName,
			connection,
			dictionary,
		)

		return resource, resourceSecrets, nil
	}

	return nil, nil, errors.New("trying to generate an unsupported connection type")
}

func buildAtlasStreamSampleConnection(
	targetNamespace,
	operatorVersion,
	projectName,
	instanceName string,
	connection *admin.StreamsConnection,
	dictionary map[string]string,
) *akov2.AtlasStreamConnection {
	return &akov2.AtlasStreamConnection{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasStreamConnection",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-%s", projectName, instanceName, connection.GetName()), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: akov2.AtlasStreamConnectionSpec{
			Name:           connection.GetName(),
			ConnectionType: connection.GetType(),
		},
		Status: akov2status.AtlasStreamConnectionStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}
}

func buildAtlasStreamClusterConnection(
	targetNamespace,
	operatorVersion,
	projectName,
	instanceName string,
	connection *admin.StreamsConnection,
	dictionary map[string]string,
) *akov2.AtlasStreamConnection {
	dbRole := connection.GetDbRoleToExecute()
	return &akov2.AtlasStreamConnection{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasStreamConnection",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-%s", projectName, instanceName, connection.GetName()), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: akov2.AtlasStreamConnectionSpec{
			Name:           connection.GetName(),
			ConnectionType: connection.GetType(),
			ClusterConfig: &akov2.ClusterConnectionConfig{
				Name: connection.GetClusterName(),
				Role: akov2.StreamsClusterDBRole{
					Name:     dbRole.GetRole(),
					RoleType: dbRole.GetType(),
				},
			},
		},
		Status: akov2status.AtlasStreamConnectionStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}
}

func buildAtlasStreamKafkaConnection(
	targetNamespace,
	operatorVersion,
	projectName,
	instanceName string,
	connection *admin.StreamsConnection,
	dictionary map[string]string,
) (*akov2.AtlasStreamConnection, []*corev1.Secret) {
	authentication := connection.GetAuthentication()
	security := connection.GetSecurity()

	connSecrets := []*corev1.Secret{
		secrets.NewAtlasSecretBuilder(fmt.Sprintf("%s-%s-%s-userpass", projectName, instanceName, connection.GetName()), targetNamespace, dictionary).
			WithData(map[string][]byte{secrets.UsernameField: []byte(authentication.GetUsername()), secrets.PasswordField: []byte("")}).
			Build(),
	}

	resource := &akov2.AtlasStreamConnection{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "atlas.mongodb.com/v1",
			Kind:       "AtlasStreamConnection",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(fmt.Sprintf("%s-%s-%s", projectName, instanceName, connection.GetName()), dictionary),
			Namespace: targetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: operatorVersion,
			},
		},
		Spec: akov2.AtlasStreamConnectionSpec{
			Name:           connection.GetName(),
			ConnectionType: connection.GetType(),
			KafkaConfig: &akov2.StreamsKafkaConnection{
				Authentication: akov2.StreamsKafkaAuthentication{
					Mechanism: authentication.GetMechanism(),
					Credentials: akov2common.ResourceRefNamespaced{
						Name:      connSecrets[0].Name,
						Namespace: connSecrets[0].Namespace,
					},
				},
				BootstrapServers: connection.GetBootstrapServers(),
				Security: akov2.StreamsKafkaSecurity{
					Protocol: security.GetProtocol(),
				},
				Config: connection.GetConfig(),
			},
		},
		Status: akov2status.AtlasStreamConnectionStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}

	if security.GetProtocol() == "SSL" {
		connSecrets = append(
			connSecrets,
			secrets.NewAtlasSecretBuilder(fmt.Sprintf("%s-%s-%s-certificate", projectName, instanceName, connection.GetName()), targetNamespace, dictionary).
				WithData(map[string][]byte{secrets.CertificateField: []byte(security.GetBrokerPublicCertificate())}).
				Build(),
		)

		resource.Spec.KafkaConfig.Security.Certificate = akov2common.ResourceRefNamespaced{
			Name:      connSecrets[1].Name,
			Namespace: connSecrets[1].Namespace,
		}
	}

	return resource, connSecrets
}
