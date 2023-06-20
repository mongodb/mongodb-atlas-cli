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
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDatabaseRoleDescriber(ctrl)

	expected := atlasv2.CustomDBRole{
		Actions: []atlasv2.DBAction{
			{
				Action: "test",
				Resources: []atlasv2.DBResource{
					{
						Collection: "test",
						Db:         "test",
						Cluster:    true},
				},
			},
		},
		InheritedRoles: []atlasv2.InheritedRole{
			{
				Db:   "test",
				Role: "test",
			},
		},
		RoleName: "Test",
	}

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
		store:    mockStore,
		roleName: "",
	}

	mockStore.
		EXPECT().
		DatabaseRole(describeOpts.ConfigProjectID(), describeopts.roleName).
		Return(&expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `NAME   ACTION   DB     COLLECTION   CLUSTER 
Test   test     test   test         true
`, buf.String())
	t.Log(buf.String())
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
