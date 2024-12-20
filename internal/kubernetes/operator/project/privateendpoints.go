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

package project

import (
	"fmt"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1/status"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PrivateEndpointRequest struct {
	ProjectName         string
	ProjectID           string
	TargetNamespace     string
	Version             string
	Credentials         string
	IndependentResource bool
	Dictionary          map[string]string
}

func BuildPrivateEndpointCustomResources(
	provider store.OperatorPrivateEndpointStore,
	request PrivateEndpointRequest,
) ([]akov2.AtlasPrivateEndpoint, error) {
	services := make([]atlasv2.EndpointService, 0)

	for _, cloud := range []string{"AWS", "AZURE", "GCP"} {
		cloudServices, err := provider.PrivateEndpoints(request.ProjectID, cloud)
		if err != nil {
			return nil, err
		}

		services = append(services, cloudServices...)
	}

	privateEndpoints := make([]akov2.AtlasPrivateEndpoint, 0, len(services))
	for _, service := range services {
		resourceName := resources.NormalizeAtlasName(
			fmt.Sprintf("%s-pe-%s-%s", request.ProjectName, service.GetCloudProvider(), strings.ToLower(strings.ReplaceAll(service.GetRegionName(), "_", ""))),
			request.Dictionary,
		)
		resource := akov2.AtlasPrivateEndpoint{
			TypeMeta: metav1.TypeMeta{
				Kind:       "AtlasPrivateEndpoint",
				APIVersion: "atlas.mongodb.com/v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      resourceName,
				Namespace: request.TargetNamespace,
				Labels: map[string]string{
					features.ResourceVersion: request.Version,
				},
			},
			Spec: akov2.AtlasPrivateEndpointSpec{
				Provider: service.GetCloudProvider(),
				Region:   service.GetRegionName(),
			},
			Status: akov2status.AtlasPrivateEndpointStatus{
				Common: akoapi.Common{
					Conditions: []akoapi.Condition{},
				},
			},
		}

		if request.IndependentResource {
			resource.Spec.ExternalProject = &akov2.ExternalProjectReference{
				ID: request.ProjectID,
			}
			resource.Spec.LocalCredentialHolder = akoapi.LocalCredentialHolder{
				ConnectionSecret: &akoapi.LocalObjectReference{
					Name: resources.NormalizeAtlasName(request.Credentials, request.Dictionary),
				},
			}
		} else {
			resource.Spec.Project = &akov2common.ResourceRefNamespaced{
				Name:      request.ProjectName,
				Namespace: request.TargetNamespace,
			}
		}

		awsConfigs, err := buildAWSInterfaces(provider, request.ProjectID, service.GetId(), service.GetInterfaceEndpoints())
		if err != nil {
			return nil, err
		}
		resource.Spec.AWSConfiguration = awsConfigs

		azureConfigs, err := buildAzureInterfaces(provider, request.ProjectID, service.GetId(), service.GetPrivateEndpoints())
		if err != nil {
			return nil, err
		}
		resource.Spec.AzureConfiguration = azureConfigs

		gcpConfigs, err := buildGCPInterfaces(provider, request.ProjectID, service.GetId(), service.GetEndpointGroupNames())
		if err != nil {
			return nil, err
		}
		resource.Spec.GCPConfiguration = gcpConfigs

		privateEndpoints = append(privateEndpoints, resource)
	}

	return privateEndpoints, nil
}

func buildAWSInterfaces(
	provider store.OperatorPrivateEndpointStore,
	projectID string,
	serviceID string,
	interfaceIDs []string,
) ([]akov2.AWSPrivateEndpointConfiguration, error) {
	if len(interfaceIDs) == 0 {
		return nil, nil
	}

	configs := make([]akov2.AWSPrivateEndpointConfiguration, 0, len(interfaceIDs))

	for _, interfaceID := range interfaceIDs {
		pe, err := provider.InterfaceEndpoint(projectID, "AWS", serviceID, interfaceID)
		if err != nil {
			return nil, err
		}

		configs = append(configs, akov2.AWSPrivateEndpointConfiguration{ID: pe.GetInterfaceEndpointId()})
	}

	return configs, nil
}

func buildAzureInterfaces(
	provider store.OperatorPrivateEndpointStore,
	projectID string,
	serviceID string,
	interfaceIDs []string,
) ([]akov2.AzurePrivateEndpointConfiguration, error) {
	if len(interfaceIDs) == 0 {
		return nil, nil
	}

	configs := make([]akov2.AzurePrivateEndpointConfiguration, 0, len(interfaceIDs))

	for _, interfaceID := range interfaceIDs {
		pe, err := provider.InterfaceEndpoint(projectID, "AZURE", serviceID, interfaceID)
		if err != nil {
			return nil, err
		}

		configs = append(
			configs,
			akov2.AzurePrivateEndpointConfiguration{
				ID: pe.GetPrivateEndpointResourceId(),
				IP: pe.GetPrivateEndpointIPAddress(),
			},
		)
	}

	return configs, nil
}

func buildGCPInterfaces(
	provider store.OperatorPrivateEndpointStore,
	projectID string,
	serviceID string,
	interfaceIDs []string,
) ([]akov2.GCPPrivateEndpointConfiguration, error) {
	if len(interfaceIDs) == 0 {
		return nil, nil
	}

	configs := make([]akov2.GCPPrivateEndpointConfiguration, 0, len(interfaceIDs))

	for _, interfaceID := range interfaceIDs {
		pe, err := provider.InterfaceEndpoint(projectID, "GCP", serviceID, interfaceID)
		if err != nil {
			return nil, err
		}

		gcpEPs := make([]akov2.GCPPrivateEndpoint, 0, len(pe.GetEndpoints()))
		for _, gcpEP := range pe.GetEndpoints() {
			gcpEPs = append(
				gcpEPs,
				akov2.GCPPrivateEndpoint{
					Name: gcpEP.GetEndpointName(),
					IP:   gcpEP.GetIpAddress(),
				},
			)
		}

		configs = append(
			configs,
			akov2.GCPPrivateEndpointConfiguration{
				// GCP ProjectID is not returned by Atlas and therefore, need to be set by the user after resource is exported
				ProjectID: "",
				GroupName: pe.GetEndpointGroupName(),
				Endpoints: gcpEPs,
			},
		)
	}

	return configs, nil
}
