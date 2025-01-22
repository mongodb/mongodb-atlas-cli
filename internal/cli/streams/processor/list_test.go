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

func TestListOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessorLister(ctrl)

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

	mockStore.
		EXPECT().
		ListProcessors(gomock.Any(), gomock.Any()).
		Return(ret, nil).
		Times(2)

	t.Run("streams processors list without stats", func(t *testing.T) {
		buf := new(bytes.Buffer)
		listOpts := &ListOpts{
			store: mockStore,
			OutputOpts: cli.OutputOpts{
				OutWriter: buf,
			},
			streamsInstance: "ExampleInstance",
		}

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
		t.Log(buf.String())

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
		buf := new(bytes.Buffer)
		listOpts := &ListOpts{
			store: mockStore,
			OutputOpts: cli.OutputOpts{
				OutWriter: buf,
			},
			streamsInstance: "ExampleInstance",
			includeStats:    true,
		}

		if err := listOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
		t.Log(buf.String())

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
