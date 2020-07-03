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

	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

func TestEvents(t *testing.T) {
	cliPath, err := cli()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	const eventsEntity = "events"

	t.Run("ListProjectEvent", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			eventsEntity,
			"list",
			"--projectId="+os.Getenv("MCLI_PROJECT_ID"),
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		events := mongodbatlas.EventResponse{}
		err = json.Unmarshal(resp, &events)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(events.Results) == 0 {
			t.Errorf("got=%#v\nwant>0\n", len(events.Results))
		}
	})

	t.Run("ListOrganizationEvent", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			entity,
			eventsEntity,
			"list",
			"--orgId="+os.Getenv("MCLI_ORG_ID"),
			"--minDate="+time.Now().Add(-time.Hour*time.Duration(24)).Format("2006-01-02"),
		)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		events := mongodbatlas.EventResponse{}
		err = json.Unmarshal(resp, &events)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(events.Results) == 0 {
			t.Errorf("got=%#v\nwant>0\n", len(events.Results))
		}
	})
}
