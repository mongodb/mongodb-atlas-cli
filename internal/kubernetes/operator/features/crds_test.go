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

package features

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Test_getCRDRoot(t *testing.T) {
	properties := map[string]apiextensionsv1.JSONSchemaProps{
		"spec": {
			Items: &apiextensionsv1.JSONSchemaPropsOrArray{},
		},
	}
	specPtr := properties["spec"]

	type args struct {
		document *apiextensionsv1.CustomResourceDefinition
	}

	tests := []struct {
		name    string
		args    args
		want    *apiextensionsv1.JSONSchemaProps
		wantErr bool
	}{
		{
			name: "Can get document Root for a valid CRD",
			args: args{
				document: &apiextensionsv1.CustomResourceDefinition{
					Spec: apiextensionsv1.CustomResourceDefinitionSpec{
						Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
							{
								Schema: &apiextensionsv1.CustomResourceValidation{
									OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
										Properties: properties,
									},
								},
							},
						},
					},
				},
			},
			want:    &specPtr,
			wantErr: false,
		},
		{
			name: "Return a document is empty error",
			args: args{
				document: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Return an error if there are no versions",
			args: args{
				document: &apiextensionsv1.CustomResourceDefinition{
					Spec: apiextensionsv1.CustomResourceDefinitionSpec{
						Versions: []apiextensionsv1.CustomResourceDefinitionVersion{},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Return an error if there is no Schema",
			args: args{
				document: &apiextensionsv1.CustomResourceDefinition{
					Spec: apiextensionsv1.CustomResourceDefinitionSpec{
						Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
							{
								Schema: &apiextensionsv1.CustomResourceValidation{
									OpenAPIV3Schema: nil,
								},
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Return an error if there is no Spec",
			args: args{
				document: &apiextensionsv1.CustomResourceDefinition{
					Spec: apiextensionsv1.CustomResourceDefinitionSpec{
						Versions: []apiextensionsv1.CustomResourceDefinitionVersion{
							{
								Schema: &apiextensionsv1.CustomResourceValidation{
									OpenAPIV3Schema: &apiextensionsv1.JSONSchemaProps{
										Properties: map[string]apiextensionsv1.JSONSchemaProps{},
									},
								},
							},
						},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCRDRoot(tt.args.document)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCRDRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCRDRoot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pathExists(t *testing.T) {
	type args struct {
		path string
		data *apiextensionsv1.JSONSchemaProps
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Path is valid with Properties",
			args: args{
				path: "level1.level2",
				data: &apiextensionsv1.JSONSchemaProps{
					Properties: map[string]apiextensionsv1.JSONSchemaProps{
						"level1": {
							Properties: map[string]apiextensionsv1.JSONSchemaProps{
								"level2": {},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Path is NOT valid with Properties",
			args: args{
				path: "level1.level2",
				data: &apiextensionsv1.JSONSchemaProps{
					Properties: map[string]apiextensionsv1.JSONSchemaProps{
						"level1": {
							Properties: map[string]apiextensionsv1.JSONSchemaProps{
								"level3": {},
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "Path is valid with Items",
			args: args{
				path: "level1.level2",
				data: &apiextensionsv1.JSONSchemaProps{
					Properties: map[string]apiextensionsv1.JSONSchemaProps{
						"level1": {
							Items: &apiextensionsv1.JSONSchemaPropsOrArray{
								Schema: nil,
								JSONSchemas: []apiextensionsv1.JSONSchemaProps{
									{
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"level2": {},
										},
									},
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "Path is NOT valid with Items",
			args: args{
				path: "level1.level2",
				data: &apiextensionsv1.JSONSchemaProps{
					Properties: map[string]apiextensionsv1.JSONSchemaProps{
						"level1": {
							Items: &apiextensionsv1.JSONSchemaPropsOrArray{
								Schema: nil,
								JSONSchemas: []apiextensionsv1.JSONSchemaProps{
									{
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"level32": {},
										},
									},
								},
							},
						},
					},
				},
			},
			want: false,
		},
		{
			name: "Path is valid with Items and Props",
			args: args{
				path: "level1.level2.level3",
				data: &apiextensionsv1.JSONSchemaProps{
					Properties: map[string]apiextensionsv1.JSONSchemaProps{
						"level1": {
							Items: &apiextensionsv1.JSONSchemaPropsOrArray{
								Schema: nil,
								JSONSchemas: []apiextensionsv1.JSONSchemaProps{
									{
										Properties: map[string]apiextensionsv1.JSONSchemaProps{
											"level2": {
												Items: &apiextensionsv1.JSONSchemaPropsOrArray{
													JSONSchemas: nil,
													Schema: &apiextensionsv1.JSONSchemaProps{
														Properties: map[string]apiextensionsv1.JSONSchemaProps{
															"level3": {},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathExists(tt.args.path, tt.args.data); got != tt.want {
				t.Errorf("pathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CRDCompatibleVersion(t *testing.T) {
	t.Run("should return error when operator version is invalid", func(t *testing.T) {
		_, err := CRDCompatibleVersion("abc")
		require.Error(t, err)
	})

	t.Run("should return operator major version when it is less than supported CRD version", func(t *testing.T) {
		latestOperatorSemver, err := semver.NewVersion(LatestOperatorMajorVersion)
		require.NoError(t, err)
		operatorVersion := semver.New(latestOperatorSemver.Major()-1, latestOperatorSemver.Minor(), 2, "", "")

		expected := fmt.Sprintf("%d.%d.0", operatorVersion.Major(), operatorVersion.Minor())
		compatibleVersion, err := CRDCompatibleVersion(operatorVersion.String())
		require.NoError(t, err)
		assert.Equal(t, expected, compatibleVersion)
	})

	t.Run("should return operator major version when it is equal than supported CRD version", func(t *testing.T) {
		latestOperatorSemver, err := semver.NewVersion(LatestOperatorMajorVersion)
		require.NoError(t, err)
		operatorVersion := semver.New(latestOperatorSemver.Major(), latestOperatorSemver.Minor(), latestOperatorSemver.Patch(), "", "")

		expected := fmt.Sprintf("%d.%d.0", operatorVersion.Major(), operatorVersion.Minor())
		compatibleVersion, err := CRDCompatibleVersion(operatorVersion.String())
		require.NoError(t, err)
		assert.Equal(t, expected, compatibleVersion)
	})

	t.Run("should return CRD major version when it is less than operator version", func(t *testing.T) {
		latestOperatorSemver, err := semver.NewVersion(LatestOperatorMajorVersion)
		require.NoError(t, err)
		operatorVersion := semver.New(latestOperatorSemver.Major(), latestOperatorSemver.Minor()+1, 0, "", "")

		compatibleVersion, err := CRDCompatibleVersion(operatorVersion.String())
		require.NoError(t, err)
		assert.Equal(t, LatestOperatorMajorVersion, compatibleVersion)
	})
}
