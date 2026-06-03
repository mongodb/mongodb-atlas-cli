// Copyright 2022 MongoDB Inc
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

package telemetry

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestTrackCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventsSender(ctrl)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {
		},
	}
	_ = cmd.ExecuteContext(NewContext())

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	mockStore.
		EXPECT().
		SendEvents(gomock.Any()).
		Return(nil).
		Times(1)
	require.NoError(t, tr.trackCommand(TrackOptions{}))
}

func TestTrackCommandWithError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventsSender(ctrl)

	cacheDir := t.TempDir()
	cmd := &cobra.Command{
		Use: "test-command",
		RunE: func(_ *cobra.Command, _ []string) error {
			return errors.New("test command error")
		},
	}
	errCmd := cmd.ExecuteContext(NewContext())
	require.Error(t, errCmd)

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	mockStore.
		EXPECT().
		SendEvents(gomock.Any()).
		Return(nil).
		Times(1)

	require.NoError(t, tr.trackCommand(TrackOptions{
		Err: errCmd,
	}))
}

func TestTrackCommandWithSendError(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventsSender(ctrl)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {
		},
	}
	errCmd := cmd.ExecuteContext(NewContext())
	require.NoError(t, errCmd)

	cacheDir := t.TempDir()

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	mockStore.
		EXPECT().
		SendEvents(gomock.Any()).
		Return(errors.New("test send error")).
		Times(1)

	require.NoError(t, tr.trackCommand(TrackOptions{
		Err: errCmd,
	}))

	// Verify that the file exists
	filename := filepath.Join(cacheDir, cacheFilename)
	info, statError := tr.fs.Stat(filename)
	require.NoError(t, statError)
	// Verify the file name
	assert.Equal(t, cacheFilename, info.Name())
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	assert.Greater(t, info.Size(), minExpectedSize)
}

func TestSave(t *testing.T) {
	cacheDir := os.TempDir()
	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
	}

	var properties = map[string]any{
		"command": "mock-command",
	}
	var event = Event{
		Timestamp:  time.Now(),
		Source:     config.AtlasCLI,
		Properties: properties,
	}
	require.NoError(t, tr.save(event))
	// Verify that the file exists
	filename := path.Join(cacheDir, cacheFilename)
	info, statError := tr.fs.Stat(filename)
	require.NoError(t, statError)
	// Verify the file name
	a := assert.New(t)
	a.Equal(cacheFilename, info.Name())
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.Greater(info.Size(), minExpectedSize)
}

func TestSaveOverMaxCacheFileSize(t *testing.T) {
	cacheDir := os.TempDir()
	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: 10, // 10 bytes
		cacheDir:         cacheDir,
	}

	var properties = map[string]any{
		"command": "mock-command",
	}
	var event = Event{
		Timestamp:  time.Now(),
		Source:     config.AtlasCLI,
		Properties: properties,
	}

	// First save will work as the cache file will be new
	require.NoError(t, tr.save(event))
	// Second save should fail as the file will be larger than 10 bytes
	require.Error(t, tr.save(event))
}

func TestOpenCacheFile(t *testing.T) {
	cacheDir := t.TempDir()
	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: 10, // 10 bytes
		cacheDir:         cacheDir,
	}

	_, err := tr.openCacheFile()
	require.NoError(t, err)
	// Verify that the file exists
	filename := path.Join(cacheDir, cacheFilename)
	info, statError := tr.fs.Stat(filename)
	require.NoError(t, statError)
	// Verify the file name
	a := assert.New(t)
	a.Equal(cacheFilename, info.Name())
	// Verify that the file is empty
	var expectedSize int64 // The nil value is zero
	a.Equal(expectedSize, info.Size())
}

func TestTrackCommandBatchLimit(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventsSender(ctrl)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {},
	}
	_ = cmd.ExecuteContext(NewContext())

	cacheDir := t.TempDir()
	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	// Pre-fill cache with maxBatchSize events so total = maxBatchSize+1 after new event.
	for i := range maxBatchSize {
		require.NoError(t, tr.save(Event{Properties: map[string]any{"i": i}}))
	}

	// Expect exactly maxBatchSize events in the batch (not maxBatchSize+1).
	mockStore.EXPECT().
		SendEvents(gomock.Len(maxBatchSize)).
		Return(nil).
		Times(1)

	require.NoError(t, tr.trackCommand(TrackOptions{}))

	// The new event (allEvents[maxBatchSize]) was not in the batch and should be written back to cache.
	remaining, err := tr.read()
	require.NoError(t, err)
	assert.Len(t, remaining, 1)
}

func TestTrackCommandRemainingEventsPersistedAfterSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventsSender(ctrl)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {},
	}
	_ = cmd.ExecuteContext(NewContext())

	cacheDir := t.TempDir()
	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	totalCached := maxBatchSize + 10
	for i := range totalCached {
		require.NoError(t, tr.save(Event{Properties: map[string]any{"i": i}}))
	}

	mockStore.EXPECT().
		SendEvents(gomock.Len(maxBatchSize)).
		Return(nil).
		Times(1)

	require.NoError(t, tr.trackCommand(TrackOptions{}))

	// totalCached+1 (new event) - maxBatchSize should remain.
	remaining, err := tr.read()
	require.NoError(t, err)
	assert.Len(t, remaining, totalCached+1-maxBatchSize)
}

func TestTrackCommandRateLimitedPreservesEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventsSender(ctrl)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {},
	}
	_ = cmd.ExecuteContext(NewContext())

	cacheDir := t.TempDir()
	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	// Pre-fill cache with some events.
	preCached := 5
	for i := range preCached {
		require.NoError(t, tr.save(Event{Properties: map[string]any{"i": i}}))
	}

	mockStore.EXPECT().
		SendEvents(gomock.Any()).
		Return(errors.New("http 429: too many requests")).
		Times(1)

	require.NoError(t, tr.trackCommand(TrackOptions{}))

	// Old events still in cache + new event appended = preCached+1.
	remaining, err := tr.read()
	require.NoError(t, err)
	assert.Len(t, remaining, preCached+1)
}

func TestSaveAndIsBackedOff(t *testing.T) {
	cacheDir := t.TempDir()
	tr := &tracker{fs: afero.NewMemMapFs(), cacheDir: cacheDir}

	// No backoff file yet.
	assert.False(t, tr.isBackedOff())

	// Create backoff sentinel file — mtime = now, within backoffDuration.
	require.NoError(t, tr.saveBackoff())
	assert.True(t, tr.isBackedOff())
}

func TestRemoveBackoff(t *testing.T) {
	cacheDir := t.TempDir()
	tr := &tracker{fs: afero.NewMemMapFs(), cacheDir: cacheDir}

	require.NoError(t, tr.saveBackoff())
	assert.True(t, tr.isBackedOff())

	require.NoError(t, tr.removeBackoff())
	assert.False(t, tr.isBackedOff())
}

func TestBackoffExpired(t *testing.T) {
	cacheDir := t.TempDir()
	fs := afero.NewMemMapFs()
	tr := &tracker{fs: fs, cacheDir: cacheDir}

	require.NoError(t, tr.saveBackoff())

	// Wind the file's mtime back past backoffDuration.
	filename := filepath.Join(cacheDir, backoffFilename)
	past := time.Now().Add(-(backoffDuration + time.Second))
	require.NoError(t, fs.Chtimes(filename, past, past))

	assert.False(t, tr.isBackedOff())
}

func TestTrackCommandSkipsSendWhenBackedOff(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventsSender(ctrl)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {},
	}
	_ = cmd.ExecuteContext(NewContext())

	cacheDir := t.TempDir()
	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	require.NoError(t, tr.saveBackoff())

	// SendEvents must NOT be called.
	mockStore.EXPECT().SendEvents(gomock.Any()).Times(0)

	require.NoError(t, tr.trackCommand(TrackOptions{}))

	// Event should be cached for later.
	events, err := tr.read()
	require.NoError(t, err)
	assert.Len(t, events, 1)
}

func TestTrackCommandSetsBackoffOnSendError(t *testing.T) {
	for _, tc := range []struct {
		name string
		err  error
	}{
		{"rate limit 429", errors.New("http 429")},
		{"connection refused", errors.New("connection refused")},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockStore := NewMockEventsSender(ctrl)

			cmd := &cobra.Command{
				Use: "test-command",
				Run: func(_ *cobra.Command, _ []string) {},
			}
			_ = cmd.ExecuteContext(NewContext())

			cacheDir := t.TempDir()
			tr := &tracker{
				fs:               afero.NewMemMapFs(),
				maxCacheFileSize: defaultMaxCacheFileSize,
				cacheDir:         cacheDir,
				store:            mockStore,
				storeSet:         true,
				cmd:              cmd,
			}

			mockStore.EXPECT().SendEvents(gomock.Any()).Return(tc.err).Times(1)

			require.NoError(t, tr.trackCommand(TrackOptions{}))

			assert.True(t, tr.isBackedOff())
		})
	}
}

func TestTrackCommandClearsBackoffOnSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockEventsSender(ctrl)

	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {},
	}
	_ = cmd.ExecuteContext(NewContext())

	cacheDir := t.TempDir()
	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		store:            mockStore,
		storeSet:         true,
		cmd:              cmd,
	}

	mockStore.EXPECT().SendEvents(gomock.Any()).Return(nil).Times(1)

	require.NoError(t, tr.trackCommand(TrackOptions{}))

	assert.False(t, tr.isBackedOff())
}

func TestTrackSurvey(t *testing.T) {
	cacheDir := t.TempDir()
	cmd := &cobra.Command{
		Use: "test-command",
		Run: func(_ *cobra.Command, _ []string) {
		},
	}
	errCmd := cmd.ExecuteContext(NewContext())
	require.NoError(t, errCmd)

	tr := &tracker{
		fs:               afero.NewMemMapFs(),
		maxCacheFileSize: defaultMaxCacheFileSize,
		cacheDir:         cacheDir,
		cmd:              cmd,
	}

	response := true
	err := tr.trackSurvey(
		&survey.Confirm{Message: "test"},
		&response,
		nil,
	)
	require.NoError(t, err)
	// Verify that the file exists
	filename := filepath.Join(cacheDir, cacheFilename)
	info, statError := tr.fs.Stat(filename)
	require.NoError(t, statError)
	// Verify the file name
	a := assert.New(t)
	a.Equal(cacheFilename, info.Name())
	// Verify that the file contains some data
	var minExpectedSize int64 = 10
	a.Greater(info.Size(), minExpectedSize)
}
