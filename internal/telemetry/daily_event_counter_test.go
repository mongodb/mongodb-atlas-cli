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

package telemetry

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testCacheDir = "/tmp/test-telemetry"

func newTestCounter(t *testing.T, dailyCap int) (*dailyEventCounter, afero.Fs) {
	t.Helper()
	fs := afero.NewMemMapFs()
	c := newDailyEventCounter(fs, testCacheDir, dailyCap)
	return c, fs
}

func TestDailyEventCounter_AllowsEventsUnderCap(t *testing.T) {
	c, _ := newTestCounter(t, 5)

	for i := range 5 {
		assert.True(t, c.CheckAndIncrement(), "event %d should be allowed", i+1)
	}
	assert.Equal(t, 5, c.Count())
}

func TestDailyEventCounter_BlocksEventsAtCap(t *testing.T) {
	c, _ := newTestCounter(t, 3)

	for range 3 {
		require.True(t, c.CheckAndIncrement())
	}

	assert.False(t, c.CheckAndIncrement(), "event beyond cap should be blocked")
	assert.False(t, c.CheckAndIncrement(), "subsequent events should also be blocked")
	assert.Equal(t, 3, c.Count(), "count should not exceed cap")
}

func TestDailyEventCounter_ResetsOnNewDay(t *testing.T) {
	c, _ := newTestCounter(t, 3)
	now := time.Date(2025, 3, 6, 10, 0, 0, 0, time.UTC)
	c.nowFn = func() time.Time { return now }

	for range 3 {
		require.True(t, c.CheckAndIncrement())
	}
	assert.False(t, c.CheckAndIncrement(), "should be capped on day 1")

	// Advance to next day
	now = now.Add(24 * time.Hour)
	assert.True(t, c.CheckAndIncrement(), "should be allowed on new day")
	assert.Equal(t, 1, c.Count())
}

func TestDailyEventCounter_IsCapped(t *testing.T) {
	c, _ := newTestCounter(t, 2)

	assert.False(t, c.IsCapped())
	require.True(t, c.CheckAndIncrement())
	assert.False(t, c.IsCapped())
	require.True(t, c.CheckAndIncrement())
	assert.True(t, c.IsCapped())
}

func TestDailyEventCounter_Reset(t *testing.T) {
	c, _ := newTestCounter(t, 5)

	for range 3 {
		require.True(t, c.CheckAndIncrement())
	}
	assert.Equal(t, 3, c.Count())

	require.NoError(t, c.Reset())
	assert.Equal(t, 0, c.Count())
	assert.True(t, c.CheckAndIncrement())
}

func TestDailyEventCounter_PersistsAcrossInstances(t *testing.T) {
	fs := afero.NewMemMapFs()

	c1 := newDailyEventCounter(fs, testCacheDir, 10)
	for range 5 {
		require.True(t, c1.CheckAndIncrement())
	}

	// New instance reading from the same filesystem
	c2 := newDailyEventCounter(fs, testCacheDir, 10)
	assert.Equal(t, 5, c2.Count())
	assert.True(t, c2.CheckAndIncrement())
	assert.Equal(t, 6, c2.Count())
}

func TestDailyEventCounter_CorruptFileResetsCounter(t *testing.T) {
	c, fs := newTestCounter(t, 5)

	require.NoError(t, fs.MkdirAll(testCacheDir, dirPermissions))
	require.NoError(t, afero.WriteFile(fs, c.counterPath(), []byte("not json"), filePermissions))

	// Should treat corrupt file as fresh start
	assert.Equal(t, 0, c.Count())
	assert.True(t, c.CheckAndIncrement())
}

func TestDailyEventCounter_StaleFileResetsCounter(t *testing.T) {
	c, fs := newTestCounter(t, 5)

	// Write a counter from yesterday
	yesterday := dailyCounter{Date: "2020-01-01", Count: 100}
	data, err := json.Marshal(yesterday)
	require.NoError(t, err)
	require.NoError(t, fs.MkdirAll(testCacheDir, dirPermissions))
	require.NoError(t, afero.WriteFile(fs, c.counterPath(), data, filePermissions))

	// Should ignore the stale counter
	assert.Equal(t, 0, c.Count())
	assert.True(t, c.CheckAndIncrement())
}

func TestDailyEventCounter_DefaultCap(t *testing.T) {
	assert.Equal(t, 10_000, DefaultDailyEventCap, "default cap should match agreed threshold")
}
