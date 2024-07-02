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

package restores

import (
	"bytes"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockServerlessRestoreJobsDescriber(ctrl)

	expiresAt, _ := time.Parse("01-02-2006", "01-01-2023")
	expected := &atlasv2.ServerlessBackupRestoreJob{
		Id:                pointer.Get("test"),
		SnapshotId:        pointer.Get("test2"),
		TargetClusterName: "ClusterTest",
		DeliveryType:      "test type",
		ExpiresAt:         pointer.Get(expiresAt),
		DeliveryUrl:       &[]string{"test url"},
	}

	buf := new(bytes.Buffer)

	describeOpts := &DescribeOpts{
		store:       mockStore,
		clusterName: "Cluster0",
		id:          "1",
		OutputOpts: cli.OutputOpts{
			Template:  restoreDescribeTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		ServerlessRestoreJob(describeOpts.ProjectID, describeOpts.clusterName, describeOpts.id).
		Return(expected, nil).
		Times(1)

	require.NoError(t, describeOpts.Run())

	assert.Equal(t, `ID     SNAPSHOT   CLUSTER       TYPE        EXPIRES AT                      URLs
test   test2      ClusterTest   test type   2023-01-01 00:00:00 +0000 UTC   test url
`, buf.String())
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{
			flag.ClusterName,
			flag.ProjectID,
			flag.Output,
		},
	)
}
