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

package dbusers

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/convert"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestDBUserDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDatabaseUserDescriber(ctrl)

	var expected atlasv2.CloudDatabaseUser

	username := "user"
	authDB := convert.AdminDB

	describeOpts := &DescribeOpts{
		store:    mockStore,
		username: username,
		authDB:   authDB,
	}

	mockStore.
		EXPECT().
		DatabaseUser(describeOpts.authDB, describeOpts.ProjectID, describeOpts.username).
		Return(&expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	test.VerifyOutputTemplate(t, describeTemplate, expected)
}
