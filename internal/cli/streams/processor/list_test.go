// Copyright 2025 MongoDB Inc
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

package processor

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/stretchr/testify/assert"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestListOpts_Run(t *testing.T) {
	pipeline1 := []any{
		map[string]any{
			"$source": map[string]any{
				"connectionName": "sample_stream_solar",
			},
		},
		map[string]any{
			"$match": map[string]any{
				"group_id": 10,
			},
		},
		map[string]any{
			"$merge": map[string]any{
				"into": map[string]any{
					"db":             "testdb",
					"coll":           "testcoll",
					"connectionName": "atlascluster",
				},
			},
		},
	}

	pipeline2 := []any{
		map[string]any{
			"$source": map[string]any{
				"connectionName": "second_connection",
			},
		},
		map[string]any{
			"$match": map[string]any{
				"group_id": 123,
			},
		},
		map[string]any{
			"$merge": map[string]any{
				"into": map[string]any{
					"db":             "seconddb",
					"coll":           "secondcoll",
					"connectionName": "secondcluster",
				},
			},
		},
	}

	expected1 := atlasv2.NewStreamsProcessorWithStats("1", "ExampleSP1", pipeline1, "STOPPED")
	expected1.Stats = map[string]any{
		"dlqMessageCount":   0,
		"inputMessageCount": 150,
		"inputMessageSize":  500,
	}

	expected2 := atlasv2.NewStreamsProcessorWithStats("2", "ExampleSP2", pipeline2, "STARTED")
	expected2.Stats = map[string]any{
		"dlqMessageCount":   0,
		"inputMessageCount": 30,
		"inputMessageSize":  6000,
	}

	ret := atlasv2.NewPaginatedApiStreamsStreamProcessorWithStats()
	ret.Results = &[]atlasv2.StreamsProcessorWithStats{*expected1, *expected2}

	t.Run("streams processors list without stats", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := mocks.NewMockProcessorLister(ctrl)

		buf := new(bytes.Buffer)
		listOpts := &ListOpts{
			store: mockStore,
			OutputOpts: cli.OutputOpts{
				OutWriter: buf,
			},
			StreamsOpts: cli.StreamsOpts{Instance: "ExampleInstance"},
			ProjectOpts: cli.ProjectOpts{ProjectID: primitive.NewObjectID().Hex()},
		}

		mockStore.
			EXPECT().
			ListProcessors(gomock.Eq(listOpts.ConfigProjectID()), gomock.Eq(listOpts.Instance)).
			Return(ret, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		expectedOutput := `[
  {
    "_id": "1",
    "name": "ExampleSP1",
    "pipeline": [
      {
        "$source": {
          "connectionName": "sample_stream_solar"
        }
      },
      {
        "$match": {
          "group_id": 10
        }
      },
      {
        "$merge": {
          "into": {
            "coll": "testcoll",
            "connectionName": "atlascluster",
            "db": "testdb"
          }
        }
      }
    ],
    "state": "STOPPED"
  },
  {
    "_id": "2",
    "name": "ExampleSP2",
    "pipeline": [
      {
        "$source": {
          "connectionName": "second_connection"
        }
      },
      {
        "$match": {
          "group_id": 123
        }
      },
      {
        "$merge": {
          "into": {
            "coll": "secondcoll",
            "connectionName": "secondcluster",
            "db": "seconddb"
          }
        }
      }
    ],
    "state": "STARTED"
  }
]
`
		assert.Equal(t, expectedOutput, buf.String())
	})

	t.Run("streams processors list with stats", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := mocks.NewMockProcessorLister(ctrl)

		buf := new(bytes.Buffer)
		listOpts := &ListOpts{
			store: mockStore,
			OutputOpts: cli.OutputOpts{
				OutWriter: buf,
			},
			StreamsOpts:  cli.StreamsOpts{Instance: "ExampleInstance"},
			ProjectOpts:  cli.ProjectOpts{ProjectID: primitive.NewObjectID().Hex()},
			includeStats: true,
		}

		mockStore.
			EXPECT().
			ListProcessors(gomock.Eq(listOpts.ConfigProjectID()), gomock.Eq(listOpts.Instance)).
			Return(ret, nil).
			Times(1)

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		expectedOutput := `[
  {
    "_id": "1",
    "name": "ExampleSP1",
    "pipeline": [
      {
        "$source": {
          "connectionName": "sample_stream_solar"
        }
      },
      {
        "$match": {
          "group_id": 10
        }
      },
      {
        "$merge": {
          "into": {
            "coll": "testcoll",
            "connectionName": "atlascluster",
            "db": "testdb"
          }
        }
      }
    ],
    "state": "STOPPED",
    "stats": {
      "dlqMessageCount": 0,
      "inputMessageCount": 150,
      "inputMessageSize": 500
    }
  },
  {
    "_id": "2",
    "name": "ExampleSP2",
    "pipeline": [
      {
        "$source": {
          "connectionName": "second_connection"
        }
      },
      {
        "$match": {
          "group_id": 123
        }
      },
      {
        "$merge": {
          "into": {
            "coll": "secondcoll",
            "connectionName": "secondcluster",
            "db": "seconddb"
          }
        }
      }
    ],
    "state": "STARTED",
    "stats": {
      "dlqMessageCount": 0,
      "inputMessageCount": 30,
      "inputMessageSize": 6000
    }
  }
]
`
		assert.Equal(t, expectedOutput, buf.String())
	})
}
