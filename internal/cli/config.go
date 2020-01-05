package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type configOpts struct {
	config.Config
	Profile       string
	Service       string
	PublicAPIKey  string
	PrivateAPIKey string
	OpsManagerURL string
}

func (opts *configOpts) IsCloud() bool {
	return opts.Service == config.CloudService
}

func (opts *configOpts) IsOpsManager() bool {
	return opts.Service == config.OpsManagerService
}

func (opts *configOpts) IsCloudManager() bool {
	return opts.Service == config.CloudManagerService
}

func (opts *configOpts) Save() error {
	opts.SetService(opts.Service)
	if opts.PublicAPIKey != "" {
		opts.SetPublicAPIKey(opts.PublicAPIKey)
	}
	if opts.PrivateAPIKey != "" {
		opts.SetPrivateAPIKey(opts.PrivateAPIKey)
	}
	if opts.IsOpsManager() && opts.OpsManagerURL != "" {
		opts.SetOpsManagerURL(opts.OpsManagerURL)
	}

	return viper.WriteConfig()
}

func (opts *configOpts) Run() error {
	helpLink := "https://docs.atlas.mongodb.com/configure-api-access/"

	if opts.IsOpsManager() {
		helpLink = "https://docs.opsmanager.mongodb.com/current/tutorial/configure-public-api-access/"
	}

	var defaultQuestions = []*survey.Question{
		{
			Name: "publicAPIKey",
			Prompt: &survey.Input{
				Message: "Public API Key:",
				Help:    helpLink,
				Default: opts.Config.PublicAPIKey(),
			},
		},
		{
			Name: "privateAPIKey",
			Prompt: &survey.Password{
				Message: "Private API Key:",
				Help:    helpLink,
			},
		},
	}

	if opts.IsOpsManager() {
		var opsManagerQuestions = []*survey.Question{
			{
				Name: "opsManagerURL",
				Prompt: &survey.Input{
					Message: "Ops Manager Base URL:",
					Default: opts.Config.OpsManagerURL(),
					Help:    "Ops Manager host URL",
				},
				Validate: validURL,
			},
		}
		defaultQuestions = append(opsManagerQuestions, defaultQuestions...)
	}

	err := survey.Ask(defaultQuestions, opts)
	if err != nil {
		return err
	}

	return opts.Save()
}

func ConfigBuilder() *cobra.Command {
	opts := new(configOpts)
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure the tool",
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.Config = config.New(opts.Profile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.Service, flags.Service, config.CloudService, "service provider, Atlas, Cloud Manager or Ops Manager")
	cmd.Flags().StringVar(&opts.Profile, flags.Profile, config.DefaultProfile, "profile")

	return cmd
}
