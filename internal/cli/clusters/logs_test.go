// Copyright 2026 MongoDB Inc
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

package clusters

import (
	"io"
	"strings"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestLogsOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockLogsDownloader(ctrl)

	// Sample gzip content for testing
	sampleContent := "sample log content"

	validLogsToDownload := []string{
		"mongodb.gz",
		"mongos.gz",
		"mongosqld.gz",
		"mongodb-audit-log.gz",
		"mongos-audit-log.gz",
	}

	for _, validLogToDownload := range validLogsToDownload {
		t.Run(validLogToDownload, func(t *testing.T) {
			t.Parallel()
			opts := &LogsOpts{
				name:  validLogToDownload,
				host:  "atlas-123abc-shard-00-00.111xx.mongodb.net",
				store: mockStore,
			}
			opts.Out = opts.name
			opts.Fs = afero.NewMemMapFs()

			mockStore.
				EXPECT().
				DownloadLog(opts.newHostLogsParams()).
				Return(io.NopCloser(strings.NewReader(sampleContent)), nil).
				Times(1)
			require.NoError(t, opts.Run())
		})
	}
}

func TestLogsOpts_Run_EmptyLog(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockLogsDownloader(ctrl)

	opts := &LogsOpts{
		name:  "mongodb.gz",
		host:  "atlas-123abc-shard-00-00.111xx.mongodb.net",
		store: mockStore,
	}
	opts.Out = opts.name
	opts.Fs = afero.NewMemMapFs()

	mockStore.
		EXPECT().
		DownloadLog(opts.newHostLogsParams()).
		Return(io.NopCloser(strings.NewReader("")), nil).
		Times(1)
	err := opts.Run()
	require.Error(t, err)
	assert.Equal(t, errEmptyLog, err)
}

func TestLogsOpts_initDefaultOut(t *testing.T) {
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
			t.Parallel()
			opts := &LogsOpts{
				name: logName,
			}
			opts.Out = out
			require.NoError(t, opts.initDefaultOut())
			assert.Equal(t, want, opts.Out)
		})
	}
}

func TestLogsBuilder(t *testing.T) {
	cmd := LogsBuilder()
	require.NotNil(t, cmd)
	assert.Equal(t, "logs <hostname> <mongodb.gz|mongos.gz|mongosqld.gz|mongodb-audit-log.gz|mongos-audit-log.gz>", cmd.Use)
	assert.Contains(t, cmd.Aliases, "log")
	assert.NotEmpty(t, cmd.Short)
	assert.NotEmpty(t, cmd.Long)
	assert.NotEmpty(t, cmd.Example)
}
