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
	"reflect"
	"testing"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

func Test_getCRDRoot(t *testing.T) {
	properties := map[string]apiextensions.JSONSchemaProps{
		"spec": {
			Items: &apiextensions.JSONSchemaPropsOrArray{},
		},
	}
	specPtr := properties["spec"]

	type args struct {
		document *apiextensions.CustomResourceDefinition
	}

	tests := []struct {
		name    string
		args    args
		want    *apiextensions.JSONSchemaProps
		wantErr bool
	}{
		{
			name: "Can get document Root for a valid CRD",
			args: args{
				document: &apiextensions.CustomResourceDefinition{
					Spec: apiextensions.CustomResourceDefinitionSpec{
						Versions: []apiextensions.CustomResourceDefinitionVersion{
							{
								Schema: &apiextensions.CustomResourceValidation{
									OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
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
				document: &apiextensions.CustomResourceDefinition{
					Spec: apiextensions.CustomResourceDefinitionSpec{
						Versions: []apiextensions.CustomResourceDefinitionVersion{},
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Return an error if there is no Schema",
			args: args{
				document: &apiextensions.CustomResourceDefinition{
					Spec: apiextensions.CustomResourceDefinitionSpec{
						Versions: []apiextensions.CustomResourceDefinitionVersion{
							{
								Schema: &apiextensions.CustomResourceValidation{
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
				document: &apiextensions.CustomResourceDefinition{
					Spec: apiextensions.CustomResourceDefinitionSpec{
						Versions: []apiextensions.CustomResourceDefinitionVersion{
							{
								Schema: &apiextensions.CustomResourceValidation{
									OpenAPIV3Schema: &apiextensions.JSONSchemaProps{
										Properties: map[string]apiextensions.JSONSchemaProps{},
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
		data *apiextensions.JSONSchemaProps
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
				data: &apiextensions.JSONSchemaProps{
					Properties: map[string]apiextensions.JSONSchemaProps{
						"level1": {
							Properties: map[string]apiextensions.JSONSchemaProps{
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
				data: &apiextensions.JSONSchemaProps{
					Properties: map[string]apiextensions.JSONSchemaProps{
						"level1": {
							Properties: map[string]apiextensions.JSONSchemaProps{
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
				data: &apiextensions.JSONSchemaProps{
					Properties: map[string]apiextensions.JSONSchemaProps{
						"level1": {
							Items: &apiextensions.JSONSchemaPropsOrArray{
								Schema: nil,
								JSONSchemas: []apiextensions.JSONSchemaProps{
									{
										Properties: map[string]apiextensions.JSONSchemaProps{
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
				data: &apiextensions.JSONSchemaProps{
					Properties: map[string]apiextensions.JSONSchemaProps{
						"level1": {
							Items: &apiextensions.JSONSchemaPropsOrArray{
								Schema: nil,
								JSONSchemas: []apiextensions.JSONSchemaProps{
									{
										Properties: map[string]apiextensions.JSONSchemaProps{
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
				data: &apiextensions.JSONSchemaProps{
					Properties: map[string]apiextensions.JSONSchemaProps{
						"level1": {
							Items: &apiextensions.JSONSchemaPropsOrArray{
								Schema: nil,
								JSONSchemas: []apiextensions.JSONSchemaProps{
									{
										Properties: map[string]apiextensions.JSONSchemaProps{
											"level2": {
												Items: &apiextensions.JSONSchemaPropsOrArray{
													JSONSchemas: nil,
													Schema: &apiextensions.JSONSchemaProps{
														Properties: map[string]apiextensions.JSONSchemaProps{
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
