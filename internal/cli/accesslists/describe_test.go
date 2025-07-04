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

package accesslists

import (
	"bytes"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestWhitelistDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockProjectIPAccessListDescriber(ctrl)

	expected := &atlasv2.NetworkPermissionEntry{
		AwsSecurityGroup: pointer.Get("test"),
		CidrBlock:        pointer.Get("test"),
		Comment:          pointer.Get("test"),
		DeleteAfterDate:  &time.Time{},
		GroupId:          pointer.Get("test"),
		IpAddress:        pointer.Get("test"),
	}

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		name:  "test",
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		IPAccessList(describeOpts.ProjectID, describeOpts.name).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}
