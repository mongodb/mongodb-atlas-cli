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

package project

import (
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/features"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/kubernetes/operator/resources"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	akoapi "github.com/mongodb/mongodb-atlas-kubernetes/v2/api"
	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1"
	akov2common "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/common"
	akov2status "github.com/mongodb/mongodb-atlas-kubernetes/v2/api/v1/status"
	"go.mongodb.org/atlas-sdk/v20241113004/admin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IPAccessListRequest struct {
	ProjectName         string
	ProjectID           string
	TargetNamespace     string
	Version             string
	Credentials         string
	IndependentResource bool
	Dictionary          map[string]string
}

func BuildIPAccessList(
	provider store.ProjectIPAccessListLister,
	request IPAccessListRequest,
) (*akov2.AtlasIPAccessList, bool, error) {
	ipAccessLists, err := provider.ProjectIPAccessLists(request.ProjectID, &store.ListOptions{ItemsPerPage: MaxItems})
	if err != nil {
		return nil, false, err
	}

	if len(ipAccessLists.GetResults()) == 0 {
		return nil, true, nil
	}

	entries := make([]akov2.IPAccessEntry, 0, len(ipAccessLists.GetResults()))
	for _, ipAccessList := range ipAccessLists.GetResults() {
		entries = append(entries, fromAtlas(ipAccessList))
	}

	resource := akov2.AtlasIPAccessList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "AtlasIPAccessList",
			APIVersion: "atlas.mongodb.com/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      resources.NormalizeAtlasName(request.ProjectName+"-ip-access-list", request.Dictionary),
			Namespace: request.TargetNamespace,
			Labels: map[string]string{
				features.ResourceVersion: request.Version,
			},
		},
		Spec: akov2.AtlasIPAccessListSpec{
			Entries: entries,
		},
		Status: akov2status.AtlasIPAccessListStatus{
			Common: akoapi.Common{
				Conditions: []akoapi.Condition{},
			},
		},
	}

	if request.IndependentResource {
		resource.Spec.ProjectDualReference = akov2.ProjectDualReference{
			ExternalProjectRef: &akov2.ExternalProjectReference{
				ID: request.ProjectID,
			},
			ConnectionSecret: &akoapi.LocalObjectReference{
				Name: resources.NormalizeAtlasName(request.Credentials, request.Dictionary),
			},
		}
	} else {
		resource.Spec.ProjectDualReference = akov2.ProjectDualReference{
			ProjectRef: &akov2common.ResourceRefNamespaced{
				Name:      request.ProjectName,
				Namespace: request.TargetNamespace,
			},
		}
	}

	return &resource, false, nil
}

func fromAtlas(entry admin.NetworkPermissionEntry) akov2.IPAccessEntry {
	result := akov2.IPAccessEntry{
		AwsSecurityGroup: entry.GetAwsSecurityGroup(),
		Comment:          entry.GetComment(),
	}

	if _, ok := entry.GetDeleteAfterDateOk(); ok {
		deleteAfter := metav1.NewTime(entry.GetDeleteAfterDate())
		result.DeleteAfterDate = &deleteAfter
	}

	if _, ok := entry.GetIpAddressOk(); ok {
		result.IPAddress = entry.GetIpAddress()
	} else if _, ok := entry.GetCidrBlockOk(); ok {
		result.CIDRBlock = entry.GetCidrBlock()
	}

	return result
}
