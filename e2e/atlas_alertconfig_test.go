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
	"strconv"
	"testing"

	"github.com/go-test/deep"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongocli/internal/fixtures"
)

const (
	group         = "GROUP"
	eventTypeName = "NO_PRIMARY"
	intervalMin   = 5
	delayMin      = 0
)

func TestAtlasAlertConfig(t *testing.T) {
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
	configEntity := "configs"

	var alertID string

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			alertsEntity,
			configEntity,
			"create",
			"--event",
			eventTypeName,
			"--enabled=true",
			"--notificationType",
			group,
			"--notificationIntervalMin",
			strconv.Itoa(intervalMin),
			"--notificationDelayMin",
			strconv.Itoa(delayMin),
			"--notificationSmsEnabled=false",
			"--notificationEmailEnabled=true")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		alert := mongodbatlas.AlertConfiguration{}
		err = json.Unmarshal(resp, &alert)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if alert.EventTypeName != eventTypeName {
			t.Errorf("got=%#v\nwant=%#v\n", alert.EventTypeName, eventTypeName)
		}

		if len(alert.Notifications) != 1 {
			t.Errorf("len(alert.Notifications) got=%#v\nwant=%#v\n", len(alert.Notifications), 1)
		}

		if *alert.Notifications[0].DelayMin != delayMin {
			t.Errorf("got=%#v\nwant=%#v\n", alert.Notifications[0].DelayMin, delayMin)
		}

		if alert.Notifications[0].TypeName != group {
			t.Errorf("got=%#v\nwant=%#v\n", alert.Notifications[0].TypeName, group)
		}

		if alert.Notifications[0].IntervalMin != intervalMin {
			t.Errorf("got=%#v\nwant=%#v\n", alert.Notifications[0].IntervalMin, intervalMin)
		}

		if *alert.Notifications[0].SMSEnabled != false {
			t.Errorf("got=%#v\nwant=%#v\n", alert.Notifications[0].SMSEnabled, false)
		}

		alertID = alert.ID

	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, alertsEntity, configEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("Update", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			alertsEntity,
			configEntity,
			"update",
			alertID,
			"--event",
			eventTypeName,
			"--notificationType",
			group,
			"--notificationIntervalMin",
			strconv.Itoa(intervalMin),
			"--notificationDelayMin",
			strconv.Itoa(delayMin),
			"--notificationSmsEnabled=true",
			"--notificationEmailEnabled=true")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		alert := mongodbatlas.AlertConfiguration{}
		err = json.Unmarshal(resp, &alert)

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if *alert.Enabled {
			t.Errorf("got=%#v\nwant=%#v\n", true, false)
		}

		if len(alert.Notifications) != 1 {
			t.Errorf("got=%#v\nwant=%#v\n", len(alert.Notifications), 1)
		}

		if !*alert.Notifications[0].SMSEnabled {
			t.Errorf("got=%#v\nwant=%#v\n", false, true)
		}

		if !*alert.Notifications[0].EmailEnabled {
			t.Errorf("got=%#v\nwant=%#v\n", false, true)
		}

	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, alertsEntity, configEntity, "delete", alertID, "--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

	t.Run("List Matcher Fields", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, alertsEntity, configEntity, "fields", "type")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		fields := []string{}
		err = json.Unmarshal(resp, &fields)

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		if diff := deep.Equal(fields, fixtures.MatcherFieldsType()); diff != nil {
			t.Error(diff)
		}

	})
}
