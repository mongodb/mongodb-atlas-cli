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

package datalakepipelines

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/internal/flag"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/internal/test"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestTrigger_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockPipelinesTriggerer(ctrl)

	expected := &atlasv2.IngestionPipelineRun{
		Id:          pointer.Get("1a5cbd92c036a0eb288"),
		DatasetName: pointer.Get("pipeline 1"),
		State:       pointer.Get("PENDING"),
	}

	buf := new(bytes.Buffer)
	triggerOpts := &TriggerOpts{
		id:         "id",
		snapshotID: "snapshotID",
		store:      mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  triggerTemplate,
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		PipelineTrigger(triggerOpts.ProjectID, triggerOpts.id, triggerOpts.snapshotID).
		Return(expected, nil).
		Times(1)

	if err := triggerOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	assert.Equal(t, `ID                    DATASET NAME   STATE
1a5cbd92c036a0eb288   pipeline 1     PENDING
`, buf.String())
	t.Log(buf.String())
}

func TestTriggerBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		TriggerBuilder(),
		0,
		[]string{flag.ProjectID, flag.Output},
	)
}
