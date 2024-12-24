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
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

const jsonOutput = "json"

type DescribeOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	streamsInstance string
	processorName   string
	includeStats    bool
	store           store.ProcessorDescriber
}

func (opts *DescribeOpts) Run() error {
	describeParams := new(atlasv2.GetStreamProcessorApiParams)
	describeParams.GroupId = opts.ProjectID
	describeParams.TenantName = opts.streamsInstance
	describeParams.ProcessorName = opts.processorName

	r, err := opts.store.StreamProcessor(describeParams)
	if err != nil {
		return err
	}

	opts.Output = jsonOutput
	if opts.includeStats {
		return opts.Print(r)
	}

	sp := atlasv2.NewStreamsProcessorWithStats(r.Id, r.Name, r.Pipeline, r.State)
	return opts.Print(sp)
}

func (opts *DescribeOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams processor describe <processorName>.
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:   "describe <processorName>",
		Short: "Get details about a specific Atlas Stream Processor in a Stream Processing Instance.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `# Return a JSON-formatted view of stream processor 'ExampleProcessor' for an instance 'ExampleInstance':
  atlas streams processors describe ExampleProcessor --instance ExampleInstance`,
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

	cmd.Flags().BoolVar(&opts.includeStats, flag.IncludeStats, false, usage.IncludeStreamProcessorStats)

	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)

	_ = cmd.MarkFlagRequired(flag.Instance)

	return cmd
}
