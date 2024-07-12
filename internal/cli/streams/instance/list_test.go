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

//go:build unit

package instance

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockStreamsLister(ctrl)

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}
	listOpts.ProjectID = "list-project-id"

	listParams := new(atlasv2.ListStreamInstancesApiParams)
	listParams.ItemsPerPage = &listOpts.ItemsPerPage
	listParams.GroupId = listOpts.ProjectID
	listParams.PageNum = &listOpts.PageNum

	id := "1"
	name := "Test Tenant"

	tenant := atlasv2.NewStreamsTenant()
	tenant.Id = &id
	tenant.Name = &name
	tenant.DataProcessRegion = atlasv2.NewStreamsDataProcessRegion("AWS", "US-EAST-1")
	expected := atlasv2.NewPaginatedApiStreamsTenant()
	expected.Results = &[]atlasv2.StreamsTenant{*tenant}

	mockStore.
		EXPECT().
		ProjectStreams(listParams).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	t.Log(buf.String())
	test.VerifyOutputTemplate(t, listTemplate, expected)
	assert.Equal(t, `ID    NAME          CLOUD   REGION
1     Test Tenant   AWS     US-EAST-1
`, buf.String())
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.Page, flag.Limit},
	)
}
