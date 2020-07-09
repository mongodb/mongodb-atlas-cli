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
// +build e2e atlas,onlinearchive

package atlas_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestOnlineArchives(t *testing.T) {
	const entity = "clusters onlineArchives"

	clusterName, err := deployCluster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	cliPath, err := cli()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			entity,
			"list",
			"--clusterName="+clusterName)

		cmd.Env = os.Environ()
		_, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	if err := deleteCluster(clusterName); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}
