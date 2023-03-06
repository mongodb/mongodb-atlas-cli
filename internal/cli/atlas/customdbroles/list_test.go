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

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/openlyinc/pointy"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseRoleLister(ctrl)
	defer ctrl.Finish()

	expected := &[]mongodbatlas.CustomDBRole{
		{
			Actions: []mongodbatlas.Action{
				{
					Action: "test",
					Resources: []mongodbatlas.Resource{
						{
							Collection: pointy.String("test"),
							DB:         pointy.String("test"),
							Cluster:    pointy.Bool(true)},
					},
				},
			},
			InheritedRoles: []mongodbatlas.InheritedRole{
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
		DatabaseRoles(listOpts.ProjectID, listOpts.NewListOptions()).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `NAME   ACTION   INHERITED ROLES   DB       COLLECTION   CLUSTER 
Test   test     N/A               test     test         true
Test   N/A      dbAdmin           pandas   N/A          N/A
`, buf.String())
	t.Log(buf.String())
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.Page, flag.Limit},
	)
}
