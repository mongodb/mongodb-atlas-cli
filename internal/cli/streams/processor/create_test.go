package processor

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

func TestCreateOpts_Run(t *testing.T) {
	t.Run("streams processor create should fail if no file is passed", func(t *testing.T) {
		createOpts := &CreateOpts{fs: afero.NewMemMapFs()}
		err := createOpts.Run()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "missing file")
	})

	t.Run("streams processor create should fail if no name is provided", func(t *testing.T) {
		fs := afero.NewMemMapFs()
		fileContents := `
{
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
  ]
}
`
		fileName := "no-name.json"
		require.NoError(t, afero.WriteFile(fs, fileName, []byte(fileContents), 0600))

		createOpts := &CreateOpts{
			fs:              fs,
			filename:        fileName,
			streamsInstance: "ExampleInstance",
		}

		err := createOpts.Run()
		require.Error(t, err)
		assert.Contains(t, err.Error(), "streams processor name missing")
	})

	t.Run("streams processor create", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockStore := mocks.NewMockProcessorCreator(ctrl)

		fs := afero.NewMemMapFs()

		fileContents := `
{
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
  ]
}
`

		fileName := "test-processor.json"
		require.NoError(t, afero.WriteFile(fs, fileName, []byte(fileContents), 0600))

		buf := new(bytes.Buffer)
		createOpts := &CreateOpts{
			store:           mockStore,
			fs:              fs,
			filename:        fileName,
			streamsInstance: "ExampleInstance",
			OutputOpts: cli.OutputOpts{
				Template:  createTemplate,
				OutWriter: buf,
			},
		}

		name := "ExampleSP"

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

		expected := atlasv2.NewStreamsProcessor()
		expected.Name = &name
		expected.Pipeline = &pipeline

		mockStore.
			EXPECT().
			CreateStreamProcessor(gomock.Any()).Return(expected, nil).
			Times(1)

		if err := createOpts.Run(); err != nil {
			t.Fatalf("Run() unexpected error: %v", err)
		}
		t.Log(buf.String())
		test.VerifyOutputTemplate(t, createTemplate, expected)
		assert.Equal(t, "Processor ExampleSP created.\n", buf.String())
	})
}
