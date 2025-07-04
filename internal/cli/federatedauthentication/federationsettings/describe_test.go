// Copyright 2024 MongoDB Inc
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

package federationsettings

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDescriber(ctrl)

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	id := "id"
	status := "ACTIVE"
	expected := &atlasv2.OrgFederationSettings{
		Id:                     &id,
		IdentityProviderStatus: &status,
	}

	mockStore.EXPECT().FederationSetting(gomock.Any()).Return(expected, nil)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplate, *expected)
}
