// Copyright 2021 MongoDB Inc
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

package customdbroles

import (
	"bytes"
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/golang/mock/gomock"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseRoleLister(ctrl)

	expected := []atlasv2.UserCustomDBRole{
		{
			Actions: &[]atlasv2.DatabasePrivilegeAction{
				{
					Action: "test",
					Resources: &[]atlasv2.DatabasePermittedNamespaceResource{
						{
							Collection: "test",
							Db:         "test",
							Cluster:    true,
						},
					},
				},
			},
			InheritedRoles: &[]atlasv2.DatabaseInheritedRole{
				{
					Db:   "pandas",
					Role: "dbAdmin",
				},
			},
			RoleName: "Test",
		},
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		DatabaseRoles(listOpts.ProjectID).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, listTemplate, expected)
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
