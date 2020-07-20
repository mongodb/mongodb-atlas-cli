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
// +build e2e atlas,generic

package atlas_test

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
	var alertID string

	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// This test should be run before all other tests to grab an alert ID for all other tests
	t.Run("List", func(t *testing.T) {
		t.Run("with no status", func(t *testing.T) {
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

			var alerts mongodbatlas.AlertsResponse
			if err := json.Unmarshal(resp, &alerts); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(alerts.Results) == 0 {
				t.Errorf("got=%#v\nwant>0\n", len(alerts.Results))
			}
			alertID = alerts.Results[0].ID
		})

		t.Run("with status OPEN", func(t *testing.T) {
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

			var alerts mongodbatlas.AlertsResponse
			if err := json.Unmarshal(resp, &alerts); err != nil {
				t.Fatalf("unexpected error: %v", err)
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

			var alerts mongodbatlas.AlertsResponse
			if err := json.Unmarshal(resp, &alerts); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	})

	t.Run("Describe", func(t *testing.T) {
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

	t.Run("Acknowledge", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
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
			atlasEntity,
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
		if err := json.Unmarshal(resp, &alert); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if alert.ID != alertID {
			t.Errorf("got=%#v\nwant%v\n", alert.ID, alertID)
		}
	})

	t.Run("UnAcknowledge", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
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
