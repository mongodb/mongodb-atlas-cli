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
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312005/admin"
	"go.uber.org/mock/gomock"
)

func TestAcknowledgeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockAlertAcknowledger(ctrl)

	tests := []struct {
		name    string
		opts    AcknowledgeOpts
		wantErr bool
	}{
		{
			name: "default",
			opts: AcknowledgeOpts{
				alertID: "533dc40ae4b00835ff81eaee",
				comment: "Test",
				store:   mockStore,
				until:   "2123-06-30T13:13:09+01:00",
			},
			wantErr: false,
		},
		{
			name: "forever",
			opts: AcknowledgeOpts{
				alertID: "533dc40ae4b00835ff81eaee",
				comment: "Test",
				forever: true,
				store:   mockStore,
			},
			wantErr: false,
		},
		{
			name: "with error",
			opts: AcknowledgeOpts{
				alertID: "533dc40ae4b00835ff81eaee",
				comment: "Test",
				forever: true,
				store:   mockStore,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		opts := tt.opts
		wantErr := tt.wantErr
		t.Run(tt.name, func(t *testing.T) {
			ackReq, _ := opts.newAcknowledgeRequest()
			params := &atlasv2.AcknowledgeAlertApiParams{
				GroupId:          opts.ProjectID,
				AlertId:          opts.alertID,
				AcknowledgeAlert: ackReq,
			}
			if wantErr {
				mockStore.
					EXPECT().
					AcknowledgeAlert(params).
					Return(nil, errors.New("fake")).
					Times(1)
				require.Error(t, opts.Run())
			} else {
				expected := &atlasv2.AlertViewForNdsGroup{}
				mockStore.
					EXPECT().
					AcknowledgeAlert(params).
					Return(expected, nil).
					Times(1)
				require.NoError(t, opts.Run())
			}
		})
	}
}
