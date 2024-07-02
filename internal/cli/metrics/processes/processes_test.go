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

package processes

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

const oneMinute = "PT1M"

func TestProcess_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessMeasurementLister(ctrl)

	expected := &atlasv2.ApiMeasurementsGeneralViewAtlas{}

	listOpts := &Opts{
		host:  "hard-00-00.mongodb.net",
		port:  27017,
		store: mockStore,
	}
	listOpts.Granularity = oneMinute
	listOpts.Period = oneMinute

	params := listOpts.NewProcessMeasurementsAPIParams("", "hard-00-00.mongodb.net:27017")

	mockStore.
		EXPECT().ProcessMeasurements(params).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		0,
		[]string{
			flag.Page, flag.Limit, flag.Granularity, flag.Period, flag.Start,
			flag.End, flag.TypeFlag, flag.ProjectID, flag.Output},
	)
}
