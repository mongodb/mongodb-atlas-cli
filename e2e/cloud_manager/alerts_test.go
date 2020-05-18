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

package cloud_manager_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

const (
	open                            = "OPEN"
	users_without_multi_factor_auth = "USERS_WITHOUT_MULTI_FACTOR_AUTH"
)

func TestCloudManagerAlerts(t *testing.T) {
	cliPath, err := filepath.Abs("../../bin/mongocli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	atlasEntity := "cloud-manager"
	alertsEntity := "alerts"

	t.Run("Describe", func(t *testing.T) {
		alertID := "5ec2ac941271767f21cbaefe"

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

		if alert.Status != open {
			t.Errorf("got=%#v\nwant=%#v\n", alert.Status, open)
		}

		if alert.EventTypeName != users_without_multi_factor_auth {
			t.Errorf("got=%#v\nwant=%#v\n", alert.EventTypeName, users_without_multi_factor_auth)
		}

	})

	t.Run("List with no status", func(t *testing.T) {

		cmd := exec.Command(cliPath,
			atlasEntity,
			alertsEntity,
			"list",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		alerts := mongodbatlas.AlertsResponse{}
		err = json.Unmarshal(resp, &alerts)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(alerts.Results) == 0 {
			t.Errorf("got=%#v\nwant>0\n", len(alerts.Results))
		}

	})
	t.Run("List with status OPEN", func(t *testing.T) {

		cmd := exec.Command(cliPath,
			atlasEntity,
			alertsEntity,
			"list",
			"--status",
			"OPEN",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		alerts := mongodbatlas.AlertsResponse{}
		err = json.Unmarshal(resp, &alerts)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(alerts.Results) == 0 {
			t.Errorf("got=%#v\nwant=0\n", len(alerts.Results))
		}

	})
	t.Run("List with status CLOSED", func(t *testing.T) {

		cmd := exec.Command(cliPath,
			atlasEntity,
			alertsEntity,
			"list",
			"--status",
			"CLOSED",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		alerts := mongodbatlas.AlertsResponse{}
		err = json.Unmarshal(resp, &alerts)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(alerts.Results) > 0 {
			t.Errorf("got=%#v\nwant>0\n", len(alerts.Results))
		}

	})
}
