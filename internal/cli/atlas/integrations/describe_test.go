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

package integrations

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201002/admin"
)

func TestDescribe_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockIntegrationDescriber(ctrl)

	buf := new(bytes.Buffer)
	describeOpts := &DescribeOpts{
		store:           mockStore,
		integrationType: "SLACK",
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplateSlack,
			OutWriter: buf,
		},
	}

	expected := &atlasv2.ThridPartyIntegration{
		ApiToken: pointer.Get("testToken"),
		TeamName: pointer.Get("testTeam"),
		Type:     pointer.Get("SLACK"),
	}
	expected.ChannelName = pointer.Get("testChannel")
	mockStore.
		EXPECT().
		Integration(describeOpts.ProjectID, describeOpts.integrationType).
		Return(expected, nil).
		Times(1)

	if err := describeOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	test.VerifyOutputTemplate(t, describeTemplateDatadogOpsGenie, expected)
	test.VerifyOutputTemplate(t, describeTemplateMicrosoftTeams, expected)
	test.VerifyOutputTemplate(t, describeTemplateNewRelic, expected)
	test.VerifyOutputTemplate(t, describeTemplatePagerDuty, expected)
	test.VerifyOutputTemplate(t, describeTemplateSlack, expected)
	test.VerifyOutputTemplate(t, describeTemplateVictorOps, expected)
	test.VerifyOutputTemplate(t, describeTemplateWebhook, expected)
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.ProjectID},
	)
}
