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

type ListOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.ListOpts
	streamsInstance string
	includeStats    bool
	store           store.ProcessorLister
}

func (opts *ListOpts) Run() error {
	listParams := new(atlasv2.ListStreamProcessorsApiParams)
	listParams.ItemsPerPage = &opts.ItemsPerPage
	listParams.GroupId = opts.ProjectID
	listParams.PageNum = &opts.PageNum
	listParams.TenantName = opts.streamsInstance

	r, err := opts.store.ListProcessors(listParams)
	if err != nil {
		return err
	}

	sps := make([]atlasv2.StreamsProcessorWithStats, 0, len(*r.Results))
	opts.Output = jsonOutput
	if opts.includeStats {
		return opts.Print(*r.Results)
	}

	for _, res := range *r.Results {
		sps = append(sps, *atlasv2.NewStreamsProcessorWithStats(res.Id, res.Name, res.Pipeline, res.State))
	}
	return opts.Print(sps)
}

func (opts *ListOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams processor list.
func ListBuilder() *cobra.Command {
	opts := &ListOpts{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all the Atlas Stream Processing Processors for your project.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Read Only"),
		Example: `  # Return a JSON-formatted list of all Atlas Stream Processors for an instance 'ExampleInstance' for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas streams processors list --projectId 5e2211c17a3e5a48f5497de3 --instance ExampleInstance`,
		Args: require.NoArgs,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().IntVar(&opts.PageNum, flag.Page, cli.DefaultPage, usage.Page)
	cmd.Flags().IntVar(&opts.ItemsPerPage, flag.Limit, cli.DefaultPageLimit, usage.Limit)

	cmd.Flags().BoolVar(&opts.includeStats, flag.IncludeStats, false, usage.IncludeStreamProcessorStats)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)

	_ = cmd.MarkFlagRequired(flag.Instance)

	return cmd
}
