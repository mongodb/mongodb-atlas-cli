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
			streamsInstance: "ExampleInstance",
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
			streamsInstance: "ExampleInstance",
			includeStats:    true,
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
