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

type DeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	streamsInstance string
	store           store.ProcessorDeleter
}

func (opts *DeleteOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *DeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteStreamProcessor, opts.ProjectID, opts.streamsInstance)
}

// atlas streams processor delete <processorName>.
func DeleteBuilder() *cobra.Command {
	opts := &DeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Atlas Stream Processor '%s' deleted\n", "Atlas Stream Processor not deleted"),
	}
	cmd := &cobra.Command{
		Use:   "delete <processorName>",
		Short: "Delete a specific Atlas Stream Processor in a Stream Processing Instance.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `# deletes stream processor 'ExampleSP' from instance 'ExampleInstance':
  atlas streams processors delete ExampleSP --instance ExampleInstance`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"processorNameDesc": "Name of the Stream Processor",
			"output":            opts.SuccessMessage(),
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)
	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)

	_ = cmd.MarkFlagRequired(flag.Instance)

	return cmd
}
