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

package e2e_test

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312001/admin"
)

const analyzer = "lucene.simple"

func TestSearch(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProjectAndCluster("search")
	r := require.New(t)

	cliPath, err := AtlasCLIBin()
	r.NoError(err)

	n := g.memoryRand("rand", 1000)
	indexName := fmt.Sprintf("index-%v", n)
	var indexID string

	g.Run("Load Sample data", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"load",
			g.clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, resp)
		var r *atlasv2.SampleDatasetStatus
		require.NoError(t, json.Unmarshal(resp, &r))

		cmd = exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"watch",
			r.GetId(),
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		resp, err = RunAndGetStdOut(cmd)
		require.NoError(t, err, resp)
	})

	g.Run("Create via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		fileName := fmt.Sprintf("create_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
		})

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "movies",
	"database": "sample_mflix",
	"name": "{{ .indexName }}",
	"definition": {
		"mappings": {
			"dynamic": true
		}
	}
}`))
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, index.GetName(), indexName)
		indexID = index.GetIndexID()
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, indexID, index.GetIndexID())
	})

	g.Run("Update via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		fileName := fmt.Sprintf("update_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
		})

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "movies",
	"database": "sample_mflix",
	"name": "{{ .indexName }}",
	"definition": {
		"analyzer": "{{ .analyzer }}",
		"mappings": {
			"dynamic": true
		}
	}
}`))
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
			"analyzer":  analyzer,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index *atlasv2.SearchIndexResponse
		require.NoError(t, json.Unmarshal(resp, &index))
		a := assert.New(t)
		a.Equal(indexID, index.GetIndexID())
		a.NotNil(index.GetLatestDefinition().Analyzer)
		a.Equal(analyzer, *index.GetLatestDefinition().Analyzer)
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Index '%s' deleted\n", indexID)
		assert.Equal(t, expected, string(resp))
	})

	g.Run("Create combinedMapping", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		fileName := fmt.Sprintf("create_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
		})

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "planets",
	"database": "sample_guides",
	"name": "{{ .indexName }}",
	"definition": {
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
	}
}`))
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, indexName, index.Name)
	})

	g.Run("Create staticMapping", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		fileName := fmt.Sprintf("create_index_search_test-array-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
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
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, indexName, index.Name)
	})

	g.Run("Create array mapping", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		n := g.memoryRand("arrayRand", 1000)
		r.NoError(err)
		indexName := fmt.Sprintf("index-array-%v", n)
		fileName := fmt.Sprintf("create_index_search_test-array-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
		})

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "posts",
	"database": "sample_training",
	"name": "{{ .indexName }}",
	"name": "{{ .indexName }}",
	"definition": {
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
					}
				]
			}
		}
	}
}`))
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, indexName, index.Name)
	})

	g.Run("list", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"list",
			"--clusterName", g.clusterName,
			"--db=sample_mflix",
			"--collection=movies",
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var indexes []atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &indexes))
		assert.NotEmpty(t, indexes)
	})
}

func TestSearchDeprecated(t *testing.T) {
	g := newAtlasE2ETestGenerator(t, withSnapshot())
	g.generateProjectAndCluster("search")
	r := require.New(t)

	cliPath, err := AtlasCLIBin()
	r.NoError(err)

	n := g.memoryRand("rand", 1000)
	indexName := fmt.Sprintf("index-%v", n)
	var indexID string

	g.Run("Load Sample data", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"load",
			g.clusterName,
			"--projectId", g.projectID,
			"-o=json")
		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, resp)
		var r *atlasv2.SampleDatasetStatus
		require.NoError(t, json.Unmarshal(resp, &r))

		cmd = exec.Command(cliPath,
			clustersEntity,
			"sampleData",
			"watch",
			r.GetId(),
			"--projectId", g.projectID)
		cmd.Env = os.Environ()
		require.NoError(t, cmd.Run())
	})

	g.Run("Create via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		fileName := fmt.Sprintf("create_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
		})

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "movies",
	"database": "sample_mflix",
	"name": "{{ .indexName }}",
	"mappings": {
		"dynamic": true
	}
}`))
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, index.GetName(), indexName)
		indexID = index.GetIndexID()
	})

	g.Run("Describe", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, indexID, index.GetIndexID())
	})

	g.Run("Update via file", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		fileName := fmt.Sprintf("update_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		require.NoError(t, err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
		})

		tpl := template.Must(template.New("").Parse(`
{
	"collectionName": "movies",
	"database": "sample_mflix",
	"name": "{{ .indexName }}",
	"analyzer": "{{ .analyzer }}",
	"mappings": {
		"dynamic": true
	}
}`))
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
			"analyzer":  analyzer,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		a := assert.New(t)
		a.Equal(indexID, index.GetIndexID())
		a.Equal(analyzer, index.GetAnalyzer())
	})

	g.Run("Delete", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		expected := fmt.Sprintf("Index '%s' deleted\n", indexID)
		assert.Equal(t, expected, string(resp))
	})

	g.Run("Create combinedMapping", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		fileName := fmt.Sprintf("create_index_search_test-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
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
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, indexName, index.Name)
	})

	g.Run("Create staticMapping", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		fileName := fmt.Sprintf("create_index_search_test-array-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
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
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, indexName, index.Name)
	})

	g.Run("Create array mapping", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		n := g.memoryRand("arrayRand", 1000)
		r.NoError(err)
		indexName := fmt.Sprintf("index-array-%v", n)
		fileName := fmt.Sprintf("create_index_search_test-array-%v.json", n)

		file, err := os.Create(fileName)
		r.NoError(err)
		t.Cleanup(func() {
			require.NoError(t, os.Remove(fileName))
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
		require.NoError(t, tpl.Execute(file, map[string]string{
			"indexName": indexName,
		}))

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
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))
		var index atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &index))
		assert.Equal(t, indexName, index.Name)
	})

	g.Run("list", func(t *testing.T) { //nolint:thelper // g.Run replaces t.Run
		cmd := exec.Command(cliPath,
			clustersEntity,
			searchEntity,
			indexEntity,
			"list",
			"--clusterName", g.clusterName,
			"--db=sample_mflix",
			"--collection=movies",
			"--projectId", g.projectID,
			"-o=json")

		cmd.Env = os.Environ()
		resp, err := RunAndGetStdOut(cmd)
		require.NoError(t, err, string(resp))

		var indexes []atlasv2.ClusterSearchIndex
		require.NoError(t, json.Unmarshal(resp, &indexes))
		assert.NotEmpty(t, indexes)
	})
}
