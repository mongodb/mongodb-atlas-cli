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
)

type StopOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	streamsInstance string
	processorName   string
	store           store.ProcessorStopper
}

func (opts *StopOpts) Run() error {
	err := opts.store.StopStreamProcessor(opts.ProjectID, opts.streamsInstance, opts.processorName)
	if err != nil {
		return err
	}

	return opts.Print("Successfully stopped Stream Processor")
}

func (opts *StopOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams processor stop <processorName>.
func StopBuilder() *cobra.Command {
	opts := &StopOpts{}
	cmd := &cobra.Command{
		Use:   "stop <processorName>",
		Short: "Stop a specific Atlas Stream Processor in a Stream Processing Instance.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  # stop Stream Processor 'ExampleProcessor' for an instance 'ExampleInstance' for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas streams processors stop ExampleProcessor --projectId 5e2211c17a3e5a48f5497de3 --instance ExampleInstance`,
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
