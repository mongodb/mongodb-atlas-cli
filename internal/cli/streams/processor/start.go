package processor

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113001/admin"
)

type StartOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	streamsInstance string
	processorName   string
	store           store.ProcessorStarter
}

func (opts *StartOpts) Run() error {
	startParams := new(atlasv2.StartStreamProcessorApiParams)
	startParams.GroupId = opts.ProjectID
	startParams.TenantName = opts.streamsInstance
	startParams.ProcessorName = opts.processorName

	err := opts.store.StartStreamProcessor(startParams)
	if err != nil {
		return err
	}

	return opts.Print("Successfully started Stream Processor")
}

func (opts *StartOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams processor start <processorName>.
func StartBuilder() *cobra.Command {
	opts := &StartOpts{}
	cmd := &cobra.Command{
		Use:   "start <processorName>",
		Short: "Start a specific Atlas Stream Processor in a Stream Processing Instance.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  # start Stream Processor 'ExampleSP' for an instance 'ExampleInstance' for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas streams processors start ExampleSP --projectId 5e2211c17a3e5a48f5497de3 --instance ExampleInstance`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"processorNameDesc": "Name of the Stream Processor",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.processorName = args[0]
			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)

	_ = cmd.MarkFlagRequired(flag.Instance)

	return cmd
}
