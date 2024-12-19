package processor

import (
	"context"
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/file"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20241113004/admin"
)

var createTemplate = "Processor {{.Name}} created.\n"

type CreateOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	streamsInstance string
	processorName   string
	store           store.ProcessorCreator
	filename        string
	fs              afero.Fs
}

func (opts *CreateOpts) Run() error {
	createParams, err := opts.newCreateRequest()
	if err != nil {
		return err
	}

	result, err := opts.store.CreateStreamProcessor(createParams)
	if err != nil {
		return err
	}

	return opts.Print(result)
}

func (opts *CreateOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

func (opts *CreateOpts) newCreateRequest() (*atlasv2.CreateStreamProcessorApiParams, error) {
	processor := atlasv2.NewStreamsProcessorWithDefaults()
	if err := file.Load(opts.fs, opts.filename, processor); err != nil {
		return nil, err
	}

	if opts.processorName != "" {
		processor.Name = &opts.processorName
	}

	if opts.processorName == "" && processor.Name == nil {
		return nil, errors.New("streams processor name missing")
	}

	createParams := new(atlasv2.CreateStreamProcessorApiParams)
	createParams.GroupId = opts.ProjectID
	createParams.TenantName = opts.streamsInstance
	createParams.StreamsProcessor = processor

	return createParams, nil
}

// atlas streams processor create <processorName> [--projectId projectId].
func CreateBuilder() *cobra.Command {
	opts := &CreateOpts{
		fs: afero.NewOsFs(),
	}
	cmd := &cobra.Command{
		Use:   "create <processorName>",
		Short: "Creates a stream processor for an Atlas Stream Processing instance.",
		Long:  fmt.Sprintf(usage.RequiredRole, "Project Owner"),
		Args:  require.MaximumNArgs(1),
		Annotations: map[string]string{
			"processorNameDesc": "Name of the processor",
			"output":            createTemplate,
		},
		Example: `# create a new stream processor for Atlas Stream Processing Instance:
  atlas streams processor create kafkaprod -i test01 -f processorConfig.json

# create a new stream processor using the name from a cluster configuration file
  atlas streams processor create -i test01 -f clusterConfig.json
`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 0 {
				opts.processorName = args[0]
			}
			return opts.PreRunE(
				opts.ValidateProjectID,
				opts.initStore(cmd.Context()),
				opts.InitOutput(cmd.OutOrStdout(), createTemplate),
			)
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)
	opts.AddOutputOptFlags(cmd)
	cmd.Flags().StringVarP(&opts.streamsInstance, flag.Instance, flag.InstanceShort, "", usage.StreamsInstance)

	cmd.Flags().StringVarP(&opts.filename, flag.File, flag.FileShort, "", usage.StreamsConnectionFilename)
	_ = cmd.MarkFlagFilename(flag.File)

	_ = cmd.MarkFlagRequired(flag.Instance)
	_ = cmd.MarkFlagRequired(flag.File)

	return cmd
}
