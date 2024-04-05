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

// This code was autogenerated at 2023-06-23T15:50:56+01:00. Note: Manual updates are allowed, but may be overwritten.

package querylimits

import (
	"bytes"
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/golang/mock/gomock"
	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockDataFederationQueryLimitDescriber(ctrl)

	expected := &admin.DataFederationTenantQueryLimit{
		Name:          "bytesProcessed.query",
		TenantName:    pointer.Get("DataFederation1"),
		Value:         1000,
		OverrunPolicy: pointer.Get("BLOCK"),
	}

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		limitName: "id",
		store:     mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		DataFederationQueryLimit(describeOpts.ProjectID, describeOpts.tenantName, describeOpts.limitName).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
