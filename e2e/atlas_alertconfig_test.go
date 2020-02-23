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

package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const (
	group         = "GROUP"
	eventTypeName = "NO_PRIMARY"
	interval_min  = 5
	delay_min     = 0
)

func TestAtlasAlertConfig(t *testing.T) {
	cliPath, err := filepath.Abs("../bin/mcli")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	_, err = os.Stat(cliPath)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	atlasEntity := "atlas"
	alertConfigEntity := "alert-config"

	//t.Run("Create", func(t *testing.T) {
	//	cmd := exec.Command(cliPath,
	//		atlasEntity,
	//		alertConfigEntity,
	//		"create",
	//		"--event",
	//		eventTypeName,
	//		"--enabled=true",
	//		"--notificationTypeName",
	//		group,
	//		"--notificationIntervalMin",
	//		strconv.Itoa(interval_min),
	//		"--notificationDelayMin",
	//		strconv.Itoa(delay_min),
	//		"--notificationSmsEnabled=false",
	//		"--notificationEmailEnabled=true")
	//	cmd.Env = os.Environ()
	//	resp, err := cmd.CombinedOutput()
	//
	//	if err != nil {
	//		t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
	//	}
	//
	//	alert := mongodbatlas.AlertConfiguration{}
	//	err = json.Unmarshal(resp, &alert)
	//	if err != nil {
	//		t.Fatalf("unexpected error: %v", err)
	//	}
	//
	//	if alert.EventTypeName != eventTypeName {
	//		t.Errorf("got=%#v\nwant=%#v\n", alert.EventTypeName, eventTypeName)
	//	}
	//
	//	if len(alert.Notifications) != 1 {
	//		t.Errorf("len(alert.Notifications) got=%#v\nwant=%#v\n", len(alert.Notifications), 1)
	//	}
	//
	//	if *alert.Notifications[0].DelayMin != delay_min {
	//		t.Errorf("got=%#v\nwant=%#v\n", alert.Notifications[0].DelayMin, delay_min)
	//	}
	//
	//	if alert.Notifications[0].TypeName != group {
	//		t.Errorf("got=%#v\nwant=%#v\n", alert.Notifications[0].TypeName, group)
	//	}
	//
	//	if alert.Notifications[0].IntervalMin != interval_min {
	//		t.Errorf("got=%#v\nwant=%#v\n", alert.Notifications[0].IntervalMin, interval_min)
	//	}
	//
	//	if *alert.Notifications[0].SMSEnabled != false {
	//		t.Errorf("got=%#v\nwant=%#v\n", alert.Notifications[0].SMSEnabled, false)
	//	}
	//
	//})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath, atlasEntity, alertConfigEntity, "ls")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
	})

}
