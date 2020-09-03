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

func ProjectTeamsTest(t *testing.T) {
	cliPath, err := e2e.Bin()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	n, err := e2e.RandInt(1000)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projectName := fmt.Sprintf("e2e-proj-%v", n)
	projectID, err := createProject(projectName)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	teamName := fmt.Sprintf("teams%v", n)
	teamID, err := createTeam(teamName)

	defer func() {
		if e := deleteProject(projectID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
		if e := deleteTeam(teamID); e != nil {
			t.Errorf("error deleting project: %v", e)
		}
	}()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	t.Run("Add", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			projectEntity,
			teamsEntity,
			"add",
			teamID,
			"--role",
			"GROUP_READ_ONLY",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))

		var teams mongodbatlas.TeamsAssigned
		if err = json.Unmarshal(resp, &teams); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		found := false
		for _, team := range teams.Results {
			if team.TeamID == teamID {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("the team %v has not been added", teamID)
		}
	})

	t.Run("Update", func(t *testing.T) {
		roleName1 := "GROUP_READ_ONLY"
		roleName2 := "GROUP_DATA_ACCESS_READ_ONLY"
		cmd := exec.Command(cliPath,
			iamEntity,
			projectEntity,
			teamsEntity,
			"update",
			teamID,
			"--role",
			roleName1,
			"--role",
			roleName2,
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))

		var roles []mongodbatlas.TeamRoles
		if err = json.Unmarshal(resp, &roles); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(roles) != 1 {
			t.Errorf("got=%#v\nwant=%#v\n", len(roles), 1)
		}

		role := roles[0]

		if role.TeamID != teamID {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(role.RoleNames) != 2 {
			t.Errorf("got=%#v\nwant=%#v\n", len(role.RoleNames), 2)
		}

		for _, roleName := range role.RoleNames {
			if roleName != roleName1 && roleName != roleName2 {
				t.Errorf("got=%#v\nwant=%#v\n", roleName, roleName1+"|"+roleName2)
			}
		}
	})

	t.Run("List", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			projectEntity,
			teamsEntity,
			"ls",
			"--projectId",
			projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))

		var teams mongodbatlas.TeamsAssigned
		if err = json.Unmarshal(resp, &teams); err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if len(teams.Results) == 0 {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			iamEntity,
			projectEntity,
			teamsEntity,
			"delete",
			teamID,
			"--force",
			"--projectId",
			projectID)
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()
		assert.NoError(t, err, string(resp))
	})
}
