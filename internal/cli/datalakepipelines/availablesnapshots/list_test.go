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

// This code was autogenerated at 2023-04-27T17:56:13+01:00. Note: Manual updates are allowed, but may be overwritten.

package availablesnapshots

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20231115013/admin"
)

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPipelineAvailableSnapshotsLister(ctrl)

	expected := &atlasv2.PaginatedBackupSnapshot{
		Results: &[]atlasv2.DiskBackupSnapshot{
			{
				Id:          pointer.Get("5e4e593f70dfbf1010295836"),
				Description: pointer.Get("test rs"),
				Status:      pointer.Get("IDLE"),
			},
			{
				Id:          pointer.Get("5e4e593f70dfbf1010295638"),
				Description: pointer.Get("test cluster"),
				Status:      pointer.Get("IDLE"),
			},
		},
		TotalCount: pointer.Get(2),
	}

	buf := new(bytes.Buffer)
	listOpts := &ListOpts{
		store:          mockStore,
		pipelineName:   "Pipeline1",
		completedAfter: "2022-05-01",
		ListOpts: cli.ListOpts{
			PageNum:      1,
			ItemsPerPage: 20,
		},
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		PipelineAvailableSnapshots(listOpts.ProjectID, listOpts.pipelineName, convertTime(listOpts.completedAfter), listOpts.NewListOptions()).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID                         DESCRIPTION    STATUS
5e4e593f70dfbf1010295836   test rs        IDLE
5e4e593f70dfbf1010295638   test cluster   IDLE`, buf.String())
	t.Log(buf.String())
}

func TestListBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.OmitCount, flag.Page, flag.Limit, flag.CompletedAfter, flag.Pipeline},
	)
}
