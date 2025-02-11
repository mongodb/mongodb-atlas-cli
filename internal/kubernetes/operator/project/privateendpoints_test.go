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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20241113004/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBuildPrivateEndpoints(t *testing.T) {
	projectID := "project-id-1"
	projectName := "projectName-1"
	targetNamespace := "testNamespace-1"
	credentialName := "pe-creds"

	tests := map[string]struct {
		awsServices            []admin.EndpointService
		azureServices          []admin.EndpointService
		gcpServices            []admin.EndpointService
		awsEndpointInterface   *admin.PrivateLinkEndpoint
		azureEndpointInterface *admin.PrivateLinkEndpoint
		gcpEndpointInterface   *admin.PrivateLinkEndpoint
		independentResource    bool
		expectedResources      []akov2.AtlasPrivateEndpoint
	}{
		"generate private endpoint without interfaces for supported providers": {
			awsServices: []admin.EndpointService{
				{
					CloudProvider: "AWS",
					Id:            pointer.Get("aws-pe-1"),
					RegionName:    pointer.Get("US_EAST_1"),
				},
			},
			azureServices: []admin.EndpointService{
				{
					CloudProvider: "AZURE",
					Id:            pointer.Get("azure-pe-1"),
					RegionName:    pointer.Get("EUROPE_NORTH"),
				},
			},
			gcpServices: []admin.EndpointService{
				{
					CloudProvider: "GCP",
					Id:            pointer.Get("gcp-pe-1"),
					RegionName:    pointer.Get("SOUTH_AMERICA_EAST_1"),
				},
			},
			expectedResources: []akov2.AtlasPrivateEndpoint{
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasPrivateEndpoint",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "projectname-1-pe-aws-useast1",
						Namespace: targetNamespace,
						Labels: map[string]string{
							features.ResourceVersion: "2.6.0",
						},
					},
					Spec: akov2.AtlasPrivateEndpointSpec{
						ProjectDualReference: akov2.ProjectDualReference{
							ProjectRef: &akov2common.ResourceRefNamespaced{
								Name:      projectName,
								Namespace: targetNamespace,
							},
						},
						Provider: "AWS",
						Region:   "US_EAST_1",
					},
					Status: akov2status.AtlasPrivateEndpointStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasPrivateEndpoint",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "projectname-1-pe-azure-europenorth",
						Namespace: targetNamespace,
						Labels: map[string]string{
							features.ResourceVersion: "2.6.0",
						},
					},
					Spec: akov2.AtlasPrivateEndpointSpec{
						ProjectDualReference: akov2.ProjectDualReference{
							ProjectRef: &akov2common.ResourceRefNamespaced{
								Name:      projectName,
								Namespace: targetNamespace,
							},
						},
						Provider: "AZURE",
						Region:   "EUROPE_NORTH",
					},
					Status: akov2status.AtlasPrivateEndpointStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasPrivateEndpoint",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "projectname-1-pe-gcp-southamericaeast1",
						Namespace: targetNamespace,
						Labels: map[string]string{
							features.ResourceVersion: "2.6.0",
						},
					},
					Spec: akov2.AtlasPrivateEndpointSpec{
						ProjectDualReference: akov2.ProjectDualReference{
							ProjectRef: &akov2common.ResourceRefNamespaced{
								Name:      projectName,
								Namespace: targetNamespace,
							},
						},
						Provider: "GCP",
						Region:   "SOUTH_AMERICA_EAST_1",
					},
					Status: akov2status.AtlasPrivateEndpointStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
			},
		},
		"generate private endpoint with interfaces for supported providers": {
			awsServices: []admin.EndpointService{
				{
					CloudProvider:      "AWS",
					Id:                 pointer.Get("aws-pe-1"),
					RegionName:         pointer.Get("US_EAST_1"),
					InterfaceEndpoints: &[]string{"vpcpe-123456"},
				},
			},
			awsEndpointInterface: &admin.PrivateLinkEndpoint{
				InterfaceEndpointId: pointer.Get("vpcpe-123456"),
			},
			azureServices: []admin.EndpointService{
				{
					CloudProvider:    "AZURE",
					Id:               pointer.Get("azure-pe-1"),
					RegionName:       pointer.Get("EUROPE_NORTH"),
					PrivateEndpoints: &[]string{"azure/resource/id"},
				},
			},
			azureEndpointInterface: &admin.PrivateLinkEndpoint{
				PrivateEndpointResourceId: pointer.Get("azure/resource/id"),
				PrivateEndpointIPAddress:  pointer.Get("10.0.0.10"),
			},
			gcpServices: []admin.EndpointService{
				{
					CloudProvider:      "GCP",
					Id:                 pointer.Get("gcp-pe-1"),
					RegionName:         pointer.Get("SOUTH_AMERICA_EAST_1"),
					EndpointGroupNames: &[]string{"groupName"},
				},
			},
			gcpEndpointInterface: &admin.PrivateLinkEndpoint{
				EndpointGroupName: pointer.Get("groupName"),
				Endpoints: &[]admin.GCPConsumerForwardingRule{
					{
						EndpointName: pointer.Get("groupName-1"),
						IpAddress:    pointer.Get("10.0.0.10"),
					},
					{
						EndpointName: pointer.Get("groupName-2"),
						IpAddress:    pointer.Get("10.0.0.20"),
					},
				},
			},
			independentResource: true,
			expectedResources: []akov2.AtlasPrivateEndpoint{
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasPrivateEndpoint",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "projectname-1-pe-aws-useast1",
						Namespace: targetNamespace,
						Labels: map[string]string{
							features.ResourceVersion: "2.6.0",
						},
					},
					Spec: akov2.AtlasPrivateEndpointSpec{
						ProjectDualReference: akov2.ProjectDualReference{
							ExternalProjectRef: &akov2.ExternalProjectReference{
								ID: projectID,
							},
							ConnectionSecret: &akoapi.LocalObjectReference{
								Name: credentialName,
							},
						},
						Provider: "AWS",
						Region:   "US_EAST_1",
						AWSConfiguration: []akov2.AWSPrivateEndpointConfiguration{
							{
								ID: "vpcpe-123456",
							},
						},
					},
					Status: akov2status.AtlasPrivateEndpointStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasPrivateEndpoint",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "projectname-1-pe-azure-europenorth",
						Namespace: targetNamespace,
						Labels: map[string]string{
							features.ResourceVersion: "2.6.0",
						},
					},
					Spec: akov2.AtlasPrivateEndpointSpec{
						ProjectDualReference: akov2.ProjectDualReference{
							ExternalProjectRef: &akov2.ExternalProjectReference{
								ID: projectID,
							},
							ConnectionSecret: &akoapi.LocalObjectReference{
								Name: credentialName,
							},
						},
						Provider: "AZURE",
						Region:   "EUROPE_NORTH",
						AzureConfiguration: []akov2.AzurePrivateEndpointConfiguration{
							{
								ID: "azure/resource/id",
								IP: "10.0.0.10",
							},
						},
					},
					Status: akov2status.AtlasPrivateEndpointStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
				{
					TypeMeta: metav1.TypeMeta{
						Kind:       "AtlasPrivateEndpoint",
						APIVersion: "atlas.mongodb.com/v1",
					},
					ObjectMeta: metav1.ObjectMeta{
						Name:      "projectname-1-pe-gcp-southamericaeast1",
						Namespace: targetNamespace,
						Labels: map[string]string{
							features.ResourceVersion: "2.6.0",
						},
					},
					Spec: akov2.AtlasPrivateEndpointSpec{
						ProjectDualReference: akov2.ProjectDualReference{
							ExternalProjectRef: &akov2.ExternalProjectReference{
								ID: projectID,
							},
							ConnectionSecret: &akoapi.LocalObjectReference{
								Name: credentialName,
							},
						},
						Provider: "GCP",
						Region:   "SOUTH_AMERICA_EAST_1",
						GCPConfiguration: []akov2.GCPPrivateEndpointConfiguration{
							{
								ProjectID: "",
								GroupName: "groupName",
								Endpoints: []akov2.GCPPrivateEndpoint{
									{
										Name: "groupName-1",
										IP:   "10.0.0.10",
									},
									{
										Name: "groupName-2",
										IP:   "10.0.0.20",
									},
								},
							},
						},
					},
					Status: akov2status.AtlasPrivateEndpointStatus{
						Common: akoapi.Common{
							Conditions: []akoapi.Condition{},
						},
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			peStore := mocks.NewMockOperatorProjectStore(ctl)
			dictionary := resources.AtlasNameToKubernetesName()

			peStore.EXPECT().PrivateEndpoints(projectID, "AWS").Return(tt.awsServices, nil)
			peStore.EXPECT().PrivateEndpoints(projectID, "AZURE").Return(tt.azureServices, nil)
			peStore.EXPECT().PrivateEndpoints(projectID, "GCP").Return(tt.gcpServices, nil)

			if tt.awsEndpointInterface != nil {
				peStore.EXPECT().InterfaceEndpoint(projectID, "AWS", "vpcpe-123456", "aws-pe-1").Return(tt.awsEndpointInterface, nil)
			}

			if tt.azureEndpointInterface != nil {
				peStore.EXPECT().InterfaceEndpoint(projectID, "AZURE", "azure/resource/id", "azure-pe-1").Return(tt.azureEndpointInterface, nil)
			}

			if tt.gcpEndpointInterface != nil {
				peStore.EXPECT().InterfaceEndpoint(projectID, "GCP", "groupName", "gcp-pe-1").Return(tt.gcpEndpointInterface, nil)
			}

			privateEndpoints, err := BuildPrivateEndpointCustomResources(
				peStore,
				PrivateEndpointRequest{
					ProjectName:         projectName,
					ProjectID:           projectID,
					TargetNamespace:     targetNamespace,
					Version:             "2.6.0",
					Credentials:         credentialName,
					IndependentResource: tt.independentResource,
					Dictionary:          dictionary,
				},
			)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedResources, privateEndpoints)
		})
	}
}
