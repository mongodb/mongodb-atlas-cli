package processor

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestStartOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessorStarter(ctrl)

	buf := new(bytes.Buffer)
	startOpts := &StartOpts{
		store:           mockStore,
		streamsInstance: "ExampleInstance",
		processorName:   "ExampleProcessor",
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		StartStreamProcessor(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil).
		Times(1)

	if err := startOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	assert.Equal(t, "Successfully started Stream Processor\n", buf.String())
}
