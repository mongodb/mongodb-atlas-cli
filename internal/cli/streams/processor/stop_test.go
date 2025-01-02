package processor

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestStopOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockProcessorStopper(ctrl)

	buf := new(bytes.Buffer)
	startOpts := &StopOpts{
		store:           mockStore,
		streamsInstance: "ExampleInstance",
		processorName:   "ExampleProcessor",
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
		},
	}

	mockStore.
		EXPECT().
		StopStreamProcessor(gomock.Any()).
		Return(nil).
		Times(1)

	if err := startOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}
	t.Log(buf.String())
	assert.Equal(t, "Successfully stopped Stream Processor\n", buf.String())
}
