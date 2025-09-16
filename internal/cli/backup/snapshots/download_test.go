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

package snapshots

import (
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312007/admin"
	"go.uber.org/mock/gomock"
)

func TestSnapshotDownloadOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDownloader(ctrl)

	opts := &DownloadOpts{
		id:          "test.tgz",
		store:       mockStore,
		clusterName: "test",
	}
	opts.Out = opts.id
	opts.Fs = afero.NewMemMapFs()

	expected := &atlasv2.FlexBackupRestoreJob20241113{
		SnapshotUrl: pointer.Get("test.tgz"),
	}

	mockStore.
		EXPECT().
		DownloadFlexClusterSnapshot(opts.ConfigProjectID(), opts.clusterName, opts.newFlexBackupSnapshotDownloadCreate()).
		Return(expected, nil).
		Times(1)

	require.Error(t, opts.Run(), errEmptyURL.Error())
}
