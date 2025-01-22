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
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessorDescriber(ctrl)

	pipeline := []any{
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

	expected := atlasv2.NewStreamsProcessorWithStats("1", "ExampleSP", pipeline, "STOPPED")
	expected.Stats = map[string]any{
		"dlqMessageCount":   0,
		"inputMessageCount": 150,
		"inputMessageSize":  500,
	}

	mockStore.
		EXPECT().
		StreamProcessor(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(expected, nil).
		Times(2)

	t.Run("streams processors describe without stats", func(t *testing.T) {
		buf := new(bytes.Buffer)
		describeOpts := &DescribeOpts{
			store:         mockStore,
			processorName: "ExampleSP",
			OutputOpts: cli.OutputOpts{
				OutWriter: buf,
			},
			StreamsOpts: cli.StreamsOpts{Instance: "ExampleInstance"},
		}

		if err := describeOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
		t.Log(buf.String())

		expectedOutput := `{
  "_id": "1",
  "name": "ExampleSP",
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
}
`

		assert.Equal(t, expectedOutput, buf.String())
	})

	t.Run("streams processors describe with stats", func(t *testing.T) {
		buf := new(bytes.Buffer)
		describeOpts := &DescribeOpts{
			store:         mockStore,
			processorName: "ExampleSP",
			OutputOpts: cli.OutputOpts{
				OutWriter: buf,
			},
			StreamsOpts:  cli.StreamsOpts{Instance: "ExampleInstance"},
			includeStats: true,
		}

		if err := describeOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}

		expectedOutput := `{
  "_id": "1",
  "name": "ExampleSP",
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
}
`
		assert.Equal(t, expectedOutput, buf.String())
	})
}
