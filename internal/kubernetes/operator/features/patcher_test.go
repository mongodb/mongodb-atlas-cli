// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package features

import (
	"testing"

	akov2 "github.com/mongodb/mongodb-atlas-kubernetes/v2/pkg/api/v1"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/equality"
)

func TestUnkownBackupPolicyFrequencyTypesPruner(t *testing.T) {
	crdSpec := &apiextensionsv1.JSONSchemaProps{
		Properties: map[string]apiextensionsv1.JSONSchemaProps{
			"items": {
				Items: &apiextensionsv1.JSONSchemaPropsOrArray{
					Schema: &apiextensionsv1.JSONSchemaProps{
						Properties: map[string]apiextensionsv1.JSONSchemaProps{
							"frequencyType": {
								Enum: []apiextensionsv1.JSON{
									{Raw: []byte(`"daily"`)},
									{Raw: []byte(`"monthly"`)},
								},
							},
						},
					},
				},
			},
		},
	}

	tests := []struct {
		name                  string
		crdSpec               *apiextensionsv1.JSONSchemaProps
		atlasBackupPolicy     *akov2.AtlasBackupPolicy
		wantAtlasBackupPolicy *akov2.AtlasBackupPolicy
		wantErr               string
	}{
		{
			name:    "nil object",
			crdSpec: crdSpec,
			wantErr: "invalid object: *v1.AtlasBackupPolicy: <nil>",
		},
		{
			name:                  "empty object",
			crdSpec:               crdSpec,
			atlasBackupPolicy:     &akov2.AtlasBackupPolicy{},
			wantAtlasBackupPolicy: &akov2.AtlasBackupPolicy{},
		},
		{
			name:    "known frequency types",
			crdSpec: crdSpec,
			atlasBackupPolicy: &akov2.AtlasBackupPolicy{
				Spec: akov2.AtlasBackupPolicySpec{
					Items: []akov2.AtlasBackupPolicyItem{
						{FrequencyType: "daily"},
					},
				},
			},
			wantAtlasBackupPolicy: &akov2.AtlasBackupPolicy{
				Spec: akov2.AtlasBackupPolicySpec{
					Items: []akov2.AtlasBackupPolicyItem{
						{FrequencyType: "daily"},
					},
				},
			},
		},
		{
			name:    "some unknown frequency types",
			crdSpec: crdSpec,
			atlasBackupPolicy: &akov2.AtlasBackupPolicy{
				Spec: akov2.AtlasBackupPolicySpec{
					Items: []akov2.AtlasBackupPolicyItem{
						{FrequencyType: "daily"},
						{FrequencyType: "unknown"},
						{FrequencyType: "monthly"},
					},
				},
			},
			wantAtlasBackupPolicy: &akov2.AtlasBackupPolicy{
				Spec: akov2.AtlasBackupPolicySpec{
					Items: []akov2.AtlasBackupPolicyItem{
						{FrequencyType: "daily"},
						{FrequencyType: "monthly"},
					},
				},
			},
		},
		{
			name:    "all unknown frequency types",
			crdSpec: crdSpec,
			atlasBackupPolicy: &akov2.AtlasBackupPolicy{
				Spec: akov2.AtlasBackupPolicySpec{
					Items: []akov2.AtlasBackupPolicyItem{
						{FrequencyType: "unknown1"},
						{FrequencyType: "unknown2"},
						{FrequencyType: "unknown3"},
					},
				},
			},
			wantAtlasBackupPolicy: &akov2.AtlasBackupPolicy{
				Spec: akov2.AtlasBackupPolicySpec{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := ""
			if err := UnknownBackupPolicyFrequencyTypesPruner(tt.crdSpec, tt.atlasBackupPolicy); err != nil {
				gotErr = err.Error()
			}

			if gotErr != tt.wantErr {
				t.Errorf("want error %q, got %q", tt.wantErr, gotErr)
			}

			if !equality.Semantic.DeepEqual(tt.atlasBackupPolicy, tt.wantAtlasBackupPolicy) {
				t.Errorf("want %+v, got %+v", tt.wantAtlasBackupPolicy, tt.atlasBackupPolicy)
			}
		})
	}
}
