// Copyright 2020 MongoDB Inc
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

package dbusers

import (
	"testing"

	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestDBUserCreateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDatabaseUserCreator(ctrl)

	expected := &atlasv2.CloudDatabaseUser{}

	createOpts := &CreateOpts{
		username: "ProjectBar",
		password: "US",
		roles:    []string{"admin@admin"},
		store:    mockStore,
	}

	mockStore.
		EXPECT().
		CreateDatabaseUser(createOpts.newDatabaseUser()).Return(expected, nil).
		Times(1)

	if err := createOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestCreateOpts_validate(t *testing.T) {
	type fields struct {
		x509Type    string
		awsIamType  string
		oidcType    string
		ldapType    string
		password    string
		description string
		roles       []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "empty roles",
			fields:  fields{},
			wantErr: true,
		},
		{
			name: "invalid x509Type",
			fields: fields{
				roles:      []string{"test"},
				awsIamType: none,
				ldapType:   none,
				oidcType:   none,
				x509Type:   "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid ldapType",
			fields: fields{
				roles:      []string{"test"},
				awsIamType: none,
				oidcType:   none,
				ldapType:   "invalid",
				x509Type:   none,
			},
			wantErr: true,
		},
		{
			name: "invalid awsIamType",
			fields: fields{
				roles:      []string{"test"},
				awsIamType: "invalid",
				ldapType:   none,
				x509Type:   none,
				oidcType:   none,
			},
			wantErr: true,
		},
		{
			name: "invalid oidcType",
			fields: fields{
				roles:      []string{"test"},
				oidcType:   "invalid",
				awsIamType: none,
				ldapType:   none,
				x509Type:   none,
			},
			wantErr: true,
		},
		{
			name: "awsIamType and password",
			fields: fields{
				roles:       []string{"test"},
				awsIamType:  user,
				ldapType:    none,
				oidcType:    none,
				x509Type:    none,
				description: "test",
				password:    "password",
			},
			wantErr: true,
		},
		{
			name: "no external auth",
			fields: fields{
				roles:       []string{"test"},
				awsIamType:  none,
				ldapType:    none,
				x509Type:    none,
				description: "test",
				oidcType:    none,
			},
			wantErr: false,
		},
		{
			name: "no external auth with password",
			fields: fields{
				roles:      []string{"test"},
				awsIamType: none,
				ldapType:   none,
				oidcType:   none,
				x509Type:   none,
				password:   "password",
			},
			wantErr: false,
		},
		{
			name: "external auth no password",
			fields: fields{
				roles:      []string{"test"},
				awsIamType: user,
				ldapType:   none,
				oidcType:   none,
				x509Type:   none,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		fields := tt.fields
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			opts := &CreateOpts{
				x509Type:    fields.x509Type,
				awsIamType:  fields.awsIamType,
				ldapType:    fields.ldapType,
				oidcType:    fields.oidcType,
				roles:       fields.roles,
				password:    fields.password,
				description: fields.description,
			}
			if err := opts.validate(); (err != nil) != wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, wantErr)
			}
		})
	}
}
