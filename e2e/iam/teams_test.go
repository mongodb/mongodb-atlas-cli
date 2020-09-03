// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// +build e2e iam

package iam_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestTeams(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	teamName := fmt.Sprintf("teams%v", n)
	var teamID string

	t.Run("Create", func(t *testing.T) {
		username, _, err := getUserFromOrg(0)

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			iamEntity,
			teamsEntity,
			"create",
			teamName,
			"--username",
			username,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var team mongodbatlas.Team
		if err := json.Unmarshal(resp, &team); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if team.Name != teamName {
			t.Errorf("got=%#v\nwant=%#v\n", team.Name, teamName)
		}

		teamID = team.ID
	})

	t.Run("Describe By ID", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			teamsEntity,
			"describe",
			"--id",
			teamID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var team mongodbatlas.Team
		if err := json.Unmarshal(resp, &team); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if team.ID != teamID {
			t.Errorf("got=%#v\nwant=%#v\n", team.ID, teamID)
		}
	})

	t.Run("Describe By Name", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			teamsEntity,
			"describe",
			"--name",
			teamName,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var team mongodbatlas.Team
		if err := json.Unmarshal(resp, &team); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if team.Name != teamName {
			t.Errorf("got=%#v\nwant=%#v\n", team.Name, teamName)
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			teamsEntity,
			"ls",
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var teams []mongodbatlas.Team
		if err := json.Unmarshal(resp, &teams); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(teams) == 0 {
			t.Errorf("got=%#v\nwant=%#v\n", len(teams), ">0")
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			teamsEntity,
			"delete",
			teamID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))
	})
}
