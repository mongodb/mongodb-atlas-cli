// Copyright 2025 MongoDB Inc
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

//go:build unit

package project

import (
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/status"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20241113004/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestBuildIPAccessList(t *testing.T) {
	projectID := "project-ial-id"
	projectName := "projectName-ial"
	targetNamespace := "ialNamespace"
	credentialName := "ial-creds"
	deleteAfter := metav1.NewTime(time.Now().Add(time.Hour))

	tests := map[string]struct {
		ipAccessList        []admin.NetworkPermissionEntry
		independentResource bool
		expectedEmpty       bool
		expectedResource    *akov2.AtlasIPAccessList
	}{
		"ip access list is empty": {
			ipAccessList:  []admin.NetworkPermissionEntry{},
			expectedEmpty: true,
		},
		"generate ip access list with kubernetes project reference": {
			ipAccessList: []admin.NetworkPermissionEntry{
				{
					IpAddress: pointer.Get("192.168.100.233"),
					Comment:   pointer.Get("My private access"),
				},
				{
					CidrBlock: pointer.Get("10.1.1.0/24"),
					Comment:   pointer.Get("Company network"),
				},
				{
					AwsSecurityGroup: pointer.Get("sg-123456"),
					Comment:          pointer.Get("Cloud network"),
				},
				{
					IpAddress:       pointer.Get("172.16.100.10"),
					DeleteAfterDate: pointer.Get(deleteAfter.Time),
					Comment:         pointer.Get("Third party temporary access"),
				},
			},
			expectedResource: &akov2.AtlasIPAccessList{
				TypeMeta: metav1.TypeMeta{
					Kind:       "AtlasIPAccessList",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "projectname-ial-ip-access-list",
					Namespace: targetNamespace,
					Labels: map[string]string{
						features.ResourceVersion: "2.7.0",
					},
				},
				Spec: akov2.AtlasIPAccessListSpec{
					ProjectDualReference: akov2.ProjectDualReference{
						ProjectRef: &akov2common.ResourceRefNamespaced{
							Name:      projectName,
							Namespace: targetNamespace,
						},
					},
					Entries: []akov2.IPAccessEntry{
						{
							IPAddress: "192.168.100.233",
							Comment:   "My private access",
						},
						{
							CIDRBlock: "10.1.1.0/24",
							Comment:   "Company network",
						},
						{
							AwsSecurityGroup: "sg-123456",
							Comment:          "Cloud network",
						},
						{
							IPAddress:       "172.16.100.10",
							DeleteAfterDate: pointer.Get(deleteAfter),
							Comment:         "Third party temporary access",
						},
					},
				},
				Status: akov2status.AtlasIPAccessListStatus{
					Common: akoapi.Common{
						Conditions: []akoapi.Condition{},
					},
				},
			},
		},
		"generate ip access list with external project reference": {
			ipAccessList: []admin.NetworkPermissionEntry{
				{
					IpAddress: pointer.Get("192.168.100.233"),
					Comment:   pointer.Get("My private access"),
				},
				{
					CidrBlock: pointer.Get("10.1.1.0/24"),
					Comment:   pointer.Get("Company network"),
				},
			},
			independentResource: true,
			expectedResource: &akov2.AtlasIPAccessList{
				TypeMeta: metav1.TypeMeta{
					Kind:       "AtlasIPAccessList",
					APIVersion: "atlas.mongodb.com/v1",
				},
				ObjectMeta: metav1.ObjectMeta{
					Name:      "projectname-ial-ip-access-list",
					Namespace: targetNamespace,
					Labels: map[string]string{
						features.ResourceVersion: "2.7.0",
					},
				},
				Spec: akov2.AtlasIPAccessListSpec{
					ProjectDualReference: akov2.ProjectDualReference{
						ExternalProjectRef: &akov2.ExternalProjectReference{
							ID: projectID,
						},
						ConnectionSecret: &akoapi.LocalObjectReference{
							Name: credentialName,
						},
					},
					Entries: []akov2.IPAccessEntry{
						{
							IPAddress: "192.168.100.233",
							Comment:   "My private access",
						},
						{
							CIDRBlock: "10.1.1.0/24",
							Comment:   "Company network",
						},
					},
				},
				Status: akov2status.AtlasIPAccessListStatus{
					Common: akoapi.Common{
						Conditions: []akoapi.Condition{},
					},
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			ctl := gomock.NewController(t)
			ialStore := mocks.NewMockProjectIPAccessListLister(ctl)
			dictionary := resources.AtlasNameToKubernetesName()

			ialStore.EXPECT().ProjectIPAccessLists(projectID, &store.ListOptions{ItemsPerPage: MaxItems}).
				Return(&admin.PaginatedNetworkAccess{Results: &tt.ipAccessList}, nil)

			atlasIPAccessList, isEmpty, err := BuildIPAccessList(
				ialStore,
				IPAccessListRequest{
					ProjectName:         projectName,
					ProjectID:           projectID,
					TargetNamespace:     targetNamespace,
					Version:             "2.7.0",
					Credentials:         credentialName,
					IndependentResource: tt.independentResource,
					Dictionary:          dictionary,
				},
			)
			require.NoError(t, err)
			assert.Equal(t, tt.expectedEmpty, isEmpty)
			assert.Equal(t, tt.expectedResource, atlasIPAccessList)
		})
	}
}
