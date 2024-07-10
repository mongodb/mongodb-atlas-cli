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

package alerts

import (
	"bytes"
	"errors"
	"testing"

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

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{
			flag.ProjectID,
			flag.Output,
		},
	)
}

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAlertDescriber(ctrl)

	expected := &atlasv2.AlertViewForNdsGroup{
		Id:            pointer.Get("test"),
		EventTypeName: pointer.Get("NO_PRIMARY"),
		Status:        pointer.Get("test"),
		MetricName:    pointer.Get("test"),
	}
	buf := new(bytes.Buffer)

	tests := []struct {
		name    string
		cmd     *DescribeOpts
		wantErr bool
	}{
		{
			name: "default",
			cmd: &DescribeOpts{
				OutputOpts: cli.OutputOpts{
					Template:  describeTemplate,
					OutWriter: buf,
				},
				alertID: "533dc40ae4b00835ff81eaee",
				store:   mockStore,
			},

			wantErr: false,
		},
		{
			name: "default",
			cmd: &DescribeOpts{
				alertID: "533dc40ae4b00835ff81eaee",
				store:   mockStore,
				OutputOpts: cli.OutputOpts{
					Template:  describeTemplate,
					OutWriter: buf,
				},
			},
			wantErr: true,
		},
	}
	test.VerifyOutputTemplate(t, describeTemplate, expected)
	for _, tt := range tests {
		cmd := tt.cmd
		params := &atlasv2.GetAlertApiParams{
			GroupId: cmd.ProjectID,
			AlertId: cmd.alertID,
		}
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			if wantErr {
				mockStore.
					EXPECT().
					Alert(params).
					Return(nil, errors.New("fake")).
					Times(1)
				require.Error(t, cmd.Run())
			} else {
				mockStore.
					EXPECT().
					Alert(params).
					Return(expected, nil).
					Times(1)
				require.NoError(t, cmd.Run())
				assert.Equal(t, `ID     TYPE         METRIC   STATUS
test   NO_PRIMARY   test     test
`, buf.String())
				t.Log(buf.String())
			}
		})
	}
}
