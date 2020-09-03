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

func TestTeamUsers(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	teamName := fmt.Sprintf("teams%v", n)
	teamID, err := createTeam(teamName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer func() {
		if e := deleteTeam(teamID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	username, userID, err := OrgNUser(1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("Add", func(t *testing.T) {
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			iamEntity,
			teamsEntity,
			usersEntity,
			"add",
			userID,
			"--teamId",
			teamID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var users []mongodbatlas.AtlasUser
		if err := json.Unmarshal(resp, &users); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		found := false
		for _, user := range users {
			if user.Username == username {
				found = true
				break
			}
		}

		assert.True(t, found)
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			teamsEntity,
			usersEntity,
			"ls",
			"--teamId",
			teamID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		var teams []mongodbatlas.AtlasUser
		if err := json.Unmarshal(resp, &teams); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assert.NotEmpty(t, teams)
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			teamsEntity,
			usersEntity,
			"delete",
			userID,
			"--teamId",
			teamID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))
	})
}
