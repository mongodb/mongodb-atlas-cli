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
// +build e2e atlas,search

package atlas_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/mongodb/mongocli/e2e"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestSearch(t *testing.T) {
	clusterName, err := deployCluster()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	defer func() {
		if e := deleteCluster(clusterName); e != nil {
			t.Errorf("error deleting test cluster: %v", e)
		}
	}()

	cliPath, err := e2e.Bin()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	indexName := fmt.Sprintf("index-%v", r.Uint32())
	collectionName := fmt.Sprintf("collection-%v", r.Uint32())
	var indexID string

	t.Run("Create", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			searchEntity,
			indexEntity,
			"create",
			indexName,
			"--clusterName="+clusterName,
			"--db=test",
			"--collection="+collectionName,
			"--dynamic")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var index mongodbatlas.SearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			assert.Equal(t, index.Name, indexName)
			indexID = index.IndexID
		}
	})

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			searchEntity,
			indexEntity,
			"list",
			"--clusterName="+clusterName,
			"--db=test",
			"--collection="+collectionName)

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var indexes []mongodbatlas.SearchIndex
		if err := json.Unmarshal(resp, &indexes); assert.NoError(t, err) {
			assert.NotEmpty(t, indexes)
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			atlasEntity,
			clustersEntity,
			searchEntity,
			indexEntity,
			"delete",
			indexID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("Index '%s' deleted\n", indexID)
		assert.Equal(t, string(resp), expected)
	})
}
