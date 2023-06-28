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
//go:build e2e || (atlas && search)

package atlas_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"text/template"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/admin"
)

func TestSearch(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProjectAndCluster("search")
	r := require.New(t)

	cliPath, err := e2e.AtlasCLIBin()
	r.NoError(err)

	n, err := e2e.RandInt(1000)
	r.NoError(err)
	indexName := fmt.Sprintf("index-%v", n)
	collectionName := fmt.Sprintf("collection-%v", n)
	var indexID string

	t.Run("Load Sample data", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"load",
			g.clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		err := cmd.Run()
		r.NoError(err)
	})

	t.Run("Create via file", func(t *testing.T) {
		fileName := fmt.Sprintf("create_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer func() {
			if e := os.Remove(fileName); e != nil {
				t.Errorf("error deleting file '%v': %v", fileName, e)
			}
		}()

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "{{ .collectionName }}",
	"database": "test",
	"name": "{{ .indexName }}",
	"mappings": {
		"dynamic": true
	}
}`))
		err = tpl.Execute(file, map[string]string{
			"collectionName": collectionName,
			"indexName":      indexName,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"create",
			"--clusterName", g.clusterName,
			"--file",
			fileName,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var index atlasv2.ClusterSearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			assert.Equal(t, index.GetName(), indexName)
			indexID = index.GetIndexID()
		}
	})

	t.Run("Describe", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"describe",
			indexID,
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var index atlasv2.ClusterSearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			assert.Equal(t, indexID, index.GetIndexID())
		}
	})

	t.Run("Update via file", func(t *testing.T) {
		analyzer := "lucene.simple"
		fileName := fmt.Sprintf("update_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer func() {
			if e := os.Remove(fileName); e != nil {
				t.Errorf("error deleting file '%v': %v", fileName, e)
			}
		}()

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "{{ .collectionName }}",
	"database": "test",
	"name": "{{ .indexName }}",
	"analyzer": "{{ .analyzer }}",
	"mappings": {
		"dynamic": true
	}
}`))
		err = tpl.Execute(file, map[string]string{
			"collectionName": collectionName,
			"indexName":      indexName,
			"analyzer":       analyzer,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"update",
			indexID,
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"--file",
			fileName,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var index atlasv2.ClusterSearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			a := assert.New(t)
			a.Equal(indexID, index.GetIndexID())
			a.Equal(analyzer, index.GetAnalyzer())
		}
	})

	t.Run("Delete", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"delete",
			indexID,
			"--clusterName", g.clusterName,
			"--projectId", g.projectID,
			"--force")
		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		expected := fmt.Sprintf("Index '%s' deleted\n", indexID)
		assert.Equal(t, string(resp), expected)
	})

	t.Run("Create combinedMapping", func(t *testing.T) {
		fileName := fmt.Sprintf("create_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			if e := os.Remove(fileName); e != nil {
				t.Errorf("error deleting file '%v': %v", fileName, e)
			}
		})

		tpl := template.Must(template.New("").Parse(`
{
  "collectionName": "planets",
  "database": "sample_guides",
  "name": "{{ .indexName }}",
  "analyzer": "lucene.standard",
  "searchAnalyzer": "lucene.standard",
  "mappings": {
    "dynamic": false,
    "fields": {
      "name": {
        "type": "string",
        "analyzer": "lucene.whitespace",
        "multi": {
          "mySecondaryAnalyzer": {
            "type": "string",
            "analyzer": "lucene.french"
          }
        }
      },
      "mainAtmosphere": {
        "type": "string",
        "analyzer": "lucene.standard"
      },
      "surfaceTemperatureC": {
        "type": "document",
        "dynamic": true,
        "analyzer": "lucene.standard"
      }
    }
  }
}`))
		err = tpl.Execute(file, map[string]string{
			"indexName": indexName,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"create",
			"--clusterName", g.clusterName,
			"--file",
			fileName,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var index atlasv2.ClusterSearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			assert.Equal(t, index.Name, indexName)
		}
	})

	t.Run("Create staticMapping", func(t *testing.T) {
		fileName := fmt.Sprintf("create_index_search_test-array-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			if e := os.Remove(fileName); e != nil {
				t.Errorf("error deleting file '%v': %v", fileName, e)
			}
		})

		tpl := template.Must(template.New("").Parse(`
{
  "collectionName": "posts",
  "database": "sample_training",
  "name": "{{ .indexName }}",
  "analyzer": "lucene.standard",
  "searchAnalyzer": "lucene.standard",
  "mappings": {
    "dynamic": false,
    "fields": {
      "comments": {
        "type": "document",
        "fields": {
          "body": {
            "type": "string",
            "analyzer": "lucene.simple",
            "ignoreAbove": 255
          },
          "author": {
            "type": "string",
            "analyzer": "keywordLowerCase"
          }
        }
      },
      "body": {
        "type": "string",
        "analyzer": "lucene.whitespace",
        "multi": {
          "mySecondaryAnalyzer": {
            "type": "string",
            "analyzer": "keywordLowerCase"
          }
        }
      },
      "tags": {
        "type": "string",
        "analyzer": "standardLowerCase"
      }
    }
  },
"analyzers":[
      {
         "charFilters":[
            
         ],
         "name":"keywordLowerCase",
         "tokenFilters":[
            {
               "type":"lowercase"
            }
         ],
         "tokenizer":{
            "type":"keyword"
         }
      },
      {
         "charFilters":[
            
         ],
         "name":"standardLowerCase",
         "tokenFilters":[
            {
               "type":"lowercase"
            }
         ],
         "tokenizer":{
            "type":"standard"
         }
      }
   ]
}`))
		err = tpl.Execute(file, map[string]string{
			"indexName": indexName,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"create",
			"--clusterName", g.clusterName,
			"--file",
			fileName,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var index atlasv2.ClusterSearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			assert.Equal(t, index.Name, indexName)
		}
	})

	t.Run("Create array mapping", func(t *testing.T) {
		n, err := e2e.RandInt(1000)
		r.NoError(err)
		indexName := fmt.Sprintf("index-array-%v", n)
		fileName := fmt.Sprintf("create_index_search_test-array-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			if e := os.Remove(fileName); e != nil {
				t.Errorf("error deleting file '%v': %v", fileName, e)
			}
		})

		tpl := template.Must(template.New("").Parse(`
{
  "collectionName": "posts",
  "database": "sample_training",
  "name": "{{ .indexName }}",
  "analyzer": "lucene.standard",
  "searchAnalyzer": "lucene.standard",
  "mappings": {
    "dynamic": false,
    "fields": {
      "comments": [
		{
			"dynamic": true,
			"type": "document"
		},
		{
			"type": "string"
		}]
    }
  }
}`))
		err = tpl.Execute(file, map[string]string{
			"indexName": indexName,
		})
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"create",
			"--clusterName", g.clusterName,
			"--file",
			fileName,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}
		var index atlasv2.ClusterSearchIndex
		if err := json.Unmarshal(resp, &index); assert.NoError(t, err) {
			assert.Equal(t, index.Name, indexName)
		}
	})

	t.Run("list", func(t *testing.T) {
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"list",
			"--clusterName", g.clusterName,
			"--db=test",
			"--collection", collectionName,
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := cmd.CombinedOutput()

		if err != nil {
			t.Fatalf("unexpected error: %v, resp: %v", err, string(resp))
		}

		var indexes []atlasv2.ClusterSearchIndex
		if err := json.Unmarshal(resp, &indexes); assert.NoError(t, err) {
			assert.NotEmpty(t, indexes)
		}
	})
}
