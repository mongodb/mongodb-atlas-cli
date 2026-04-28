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
	"path/filepath"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/spf13/afero"
)

const (
	dailyCounterFilename = "telemetry_daily_counter"

	// DefaultDailyEventCap is the maximum number of telemetry events a single
	// CLI installation can emit per calendar day (UTC). The value was chosen
	// to sit well above the p99 of legitimate heavy users (based on
	// Grafana/Mixpanel baselines) while blocking the distributed-proxy abuse
	// pattern observed on 2025-03-06 (CLOUDP-387377 postmortem).
	DefaultDailyEventCap = 10_000

	// DailyCapWarningPercent is the threshold (as a fraction of the cap) at
	// which a warning is logged.
	DailyCapWarningPercent    = 0.80
	dailyCapWarningPercentFmt = 80.0
)

// dailyCounter tracks how many telemetry events have been recorded today.
type dailyCounter struct {
	Date  string `json:"date"`  // UTC date in YYYY-MM-DD format
	Count int    `json:"count"` // events recorded so far today
}

// dailyEventCounter reads and writes a per-day event counter file to enforce
// the global daily event cap.
type dailyEventCounter struct {
	fs       afero.Fs
	cacheDir string
	cap      int
	nowFn    func() time.Time // injectable clock for tests
}

func newDailyEventCounter(fs afero.Fs, cacheDir string, dailyCap int) *dailyEventCounter {
	return &dailyEventCounter{
		fs:       fs,
		cacheDir: cacheDir,
		cap:      dailyCap,
		nowFn:    time.Now,
	}
}

func (d *dailyEventCounter) today() string {
	return d.nowFn().UTC().Format("2006-01-02")
}

func (d *dailyEventCounter) counterPath() string {
	return filepath.Join(d.cacheDir, dailyCounterFilename)
}

// load reads the counter file. If the file is missing, corrupt, or from a
// previous day, a fresh counter for today is returned.
func (d *dailyEventCounter) load() dailyCounter {
	today := d.today()

	data, err := afero.ReadFile(d.fs, d.counterPath())
	if err != nil {
		return dailyCounter{Date: today, Count: 0}
	}

	var c dailyCounter
	if err := json.Unmarshal(data, &c); err != nil || c.Date != today {
		return dailyCounter{Date: today, Count: 0}
	}
	return c
}

// persist writes the counter to disk atomically.
func (d *dailyEventCounter) persist(c dailyCounter) error {
	exists, err := afero.DirExists(d.fs, d.cacheDir)
	if err != nil {
		return err
	}
	if !exists {
		if err := d.fs.MkdirAll(d.cacheDir, dirPermissions); err != nil {
			return err
		}
	}

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return afero.WriteFile(d.fs, d.counterPath(), data, filePermissions)
}

// CheckAndIncrement loads today's counter and returns whether the event should
// be allowed. It increments the count and persists it when allowed. It logs
// warnings when the count passes the warning threshold and when the hard cap
// is hit.
func (d *dailyEventCounter) CheckAndIncrement() (allowed bool) {
	c := d.load()

	if c.Count >= d.cap {
		_, _ = log.Debugf("telemetry: daily event cap reached (%d/%d) — event dropped\n", c.Count, d.cap)
		return false
	}

	c.Count++

	warningThreshold := int(float64(d.cap) * DailyCapWarningPercent)
	if c.Count == warningThreshold {
		_, _ = log.Warningf("telemetry: approaching daily event cap (%d/%d, %.0f%%)\n", c.Count, d.cap, dailyCapWarningPercentFmt)
	}
	if c.Count == d.cap {
		_, _ = log.Warningf("telemetry: daily event cap reached (%d/%d) — subsequent events today will be dropped\n", c.Count, d.cap)
	}

	if err := d.persist(c); err != nil {
		_, _ = log.Debugf("telemetry: failed to persist daily counter: %v\n", err)
	}

	return true
}

// Count returns the current event count for today without modifying state.
func (d *dailyEventCounter) Count() int {
	return d.load().Count
}

// Reset writes a fresh zero counter for today. Useful in tests.
func (d *dailyEventCounter) Reset() error {
	return d.persist(dailyCounter{Date: d.today(), Count: 0})
}

// IsCapped returns true when the daily cap has already been reached.
func (d *dailyEventCounter) IsCapped() bool {
	return d.load().Count >= d.cap
}
