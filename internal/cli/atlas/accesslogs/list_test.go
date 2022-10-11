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
// +build unit

package accesslogs

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

func TestAccessLogListClusterName_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAccessLogsLister(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.AccessLogSettings{
		AccessLogs: []*mongodbatlas.AccessLogs{
			{
				GroupID:       "test",
				Hostname:      "test",
				ClusterName:   "test",
				IPAddress:     "test",
				AuthResult:    pointy.Bool(true),
				LogLine:       "test",
				Timestamp:     "test",
				Username:      "test",
				FailureReason: "test",
				AuthSource:    "test",
			},
		},
	}

	buf := new(bytes.Buffer)
	opts := &ListOpts{
		store:       mockStore,
		clusterName: "test",
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		AccessLogsByClusterName(opts.ConfigProjectID(), opts.clusterName, opts.newAccessLogOptions()).
		Return(expected, nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `HOSTNAME    CLUSTER NAME   AUTH RESULT   LOG LINE 
test test   test           true          test
`, buf.String())
	t.Log(buf.String())
}

func TestAccessLogListHostname_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAccessLogsLister(ctrl)
	defer ctrl.Finish()

	expected := &mongodbatlas.AccessLogSettings{}

	describeOpts := &ListOpts{
		store:    mockStore,
		hostname: "test",
	}

	mockStore.
		EXPECT().
		AccessLogsByHostname(describeOpts.ConfigProjectID(), describeOpts.hostname, describeOpts.newAccessLogOptions()).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		ListBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output, flag.Start, flag.End, flag.IP, flag.AuthResult, flag.NLog},
	)
}
