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
// +build e2e

package e2e_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	closed                               = "CLOSED"
	replication_oplog_window_running_out = "REPLICATION_OPLOG_WINDOW_RUNNING_OUT"
)

func TestAtlasAlerts(t *testing.T) {
	cliPath, err := filepath.Abs("../bin/mongocli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	atlasEntity := "atlas"
	alertsEntity := "alerts"

	t.Run("Describe", func(t *testing.T) {
		alertID := "5e4d20ff5cc174527c22c606"

		cmd := exec.Command(cliPath,
			atlasEntity,
			alertsEntity,
			"describe",
			alertID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		alert := mongodbatlas.Alert{}
		err = json.Unmarshal(resp, &alert)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if alert.ID != alertID {
			t.Errorf("got=%#v\nwant=%#v\n", alert.ID, alertID)
		}

		if alert.Status != closed {
			t.Errorf("got=%#v\nwant=%#v\n", alert.Status, closed)
		}

		if alert.EventTypeName != replication_oplog_window_running_out {
			t.Errorf("got=%#v\nwant=%#v\n", alert.EventTypeName, replication_oplog_window_running_out)
		}

	})
}
