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
// +build e2e cloudmanager,generic

package cloud_manager_test

import (
	"encoding/json"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongocli/e2e"
	"go.mongodb.org/atlas/mongodbatlas"
)

const (
	open                        = "OPEN"
	usersWithoutMultiFactorAuth = "USERS_WITHOUT_MULTI_FACTOR_AUTH"
)

func TestAlerts(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const alertsEntity = "alerts"
	alertID := "5ec2ac941271767f21cbaefe"

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"describe",
			alertID,
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var alert mongodbatlas.Alert
		if err := json.Unmarshal(resp, &alert); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if alert.ID != alertID {
			t.Errorf("got=%#v\nwant=%#v\n", alert.ID, alertID)
		}

		if alert.Status != open {
			t.Errorf("got=%#v\nwant=%#v\n", alert.Status, open)
		}

		if alert.EventTypeName != usersWithoutMultiFactorAuth {
			t.Errorf("got=%#v\nwant=%#v\n", alert.EventTypeName, usersWithoutMultiFactorAuth)
		}
	})

	t.Run("List with no status", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"list",
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var alerts mongodbatlas.AlertsResponse
		if err = json.Unmarshal(resp, &alerts); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(alerts.Results) == 0 {
			t.Errorf("got=%#v\nwant>0\n", len(alerts.Results))
		}
	})
	t.Run("List with status OPEN", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
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

		var alerts mongodbatlas.AlertsResponse
		if err := json.Unmarshal(resp, &alerts); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(alerts.Results) == 0 {
			t.Errorf("got=%#v\nwant>0\n", len(alerts.Results))
		}
	})
	t.Run("List with status CLOSED", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
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

		var alerts mongodbatlas.AlertsResponse
		if err := json.Unmarshal(resp, &alerts); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(alerts.Results) > 0 {
			t.Errorf("got=%#v\nwant=0\n", len(alerts.Results))
		}
	})

	t.Run("Acknowledge", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"ack",
			alertID,
			"--until",
			time.Now().Format(time.RFC3339))

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var alert mongodbatlas.Alert
		if err := json.Unmarshal(resp, &alert); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if alert.ID != alertID {
			t.Errorf("got=%#v\nwant%v\n", alert.ID, alertID)
		}
	})

	t.Run("Acknowledge Forever", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"ack",
			alertID,
			"--forever")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var alert mongodbatlas.Alert
		if err = json.Unmarshal(resp, &alert); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if alert.ID != alertID {
			t.Errorf("got=%#v\nwant%v\n", alert.ID, alertID)
		}
	})

	t.Run("UnaAcknowledge", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			alertsEntity,
			"unack",
			alertID)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var alert mongodbatlas.Alert
		if err := json.Unmarshal(resp, &alert); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if alert.ID != alertID {
			t.Errorf("got=%#v\nwant%v\n", alert.ID, alertID)
		}
	})
}
