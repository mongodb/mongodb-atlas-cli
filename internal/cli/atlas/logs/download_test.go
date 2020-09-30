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

// +build unit

package logs

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestLogsDownloadOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockLogsDownloader(ctrl)
	defer ctrl.Finish()

	opts := &DownloadOpts{
		name:  "mongo.gz",
		store: mockStore,
	}
	opts.Out = opts.name
	opts.Fs = afero.NewMemMapFs()

	mockStore.
		EXPECT().
		DownloadLog(opts.ProjectID, opts.host, opts.name, gomock.Any(), opts.newDateRangeOpts()).
		Return(nil).
		Times(1)

	if err := opts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
}

func TestDownloadBuilder(t *testing.T) {
	cli.CmdValidator(
		t,
		DownloadBuilder(),
		0,
		[]string{flag.Out, flag.ProjectID, flag.Force, flag.Start, flag.End},
	)
}

func TestDownloadOpts_initDefaultOut(t *testing.T) {
	type fields struct {
		logName string
		out     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "empty out and add log",
			fields: fields{
				logName: "mongo.gz",
				out:     "",
			},
			want: "mongo.log.gz",
		},
		{
			name: "with out",
			fields: fields{
				logName: "mongo.gz",
				out:     "myfile.gz",
			},
			want: "myfile.gz",
		},
	}
	for _, tt := range tests {
		logName := tt.fields.logName
		out := tt.fields.out
		want := tt.want
		t.Run(tt.name, func(t *testing.T) {
			opts := &DownloadOpts{
				name: logName,
			}
			opts.Out = out
			opts.initDefaultOut()
			assert.Equal(t, opts.Out, want)
		})
	}
}
