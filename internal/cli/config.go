package cli

import (
	"github.com/10gen/mcli/internal/config"
	"github.com/10gen/mcli/internal/flags"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigOpts struct {
	Profile       string
	Service       string
	PublicAPIKey  string
	PrivateAPIKey string
	OpsManagerURL string
	config        config.Config
}

func (opts *ConfigOpts) IsCloud() bool {
	return opts.Service == config.CloudService
}

func (opts *ConfigOpts) IsOpsManager() bool {
	return opts.Service == config.OpsManagerService
}

func (opts *ConfigOpts) IsCloudManager() bool {
	return opts.Service == config.CloudManagerService
}

func (opts *ConfigOpts) Save() error {
	opts.config.SetService(opts.Service)
	if opts.PublicAPIKey != "" {
		opts.config.SetPublicAPIKey(opts.PublicAPIKey)
	}
	if opts.PrivateAPIKey != "" {
		opts.config.SetPrivateAPIKey(opts.PrivateAPIKey)
	}
	if opts.IsOpsManager() && opts.OpsManagerURL != "" {
		opts.config.SetOpsManagerURL(opts.OpsManagerURL)
	}

	return viper.WriteConfig()
}

func (opts *ConfigOpts) Run() error {
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
				Default: opts.config.GetPublicAPIKey(),
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
					Default: opts.config.GetOpsManagerURL(),
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
	opts := new(ConfigOpts)
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure the tool",
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.config = config.New(opts.Profile)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}
	cmd.Flags().StringVar(&opts.Service, flags.Service, config.CloudService, "Service provider, Atlas, Cloud Manager or Ops Manager")
	cmd.Flags().StringVar(&opts.Profile, flags.Profile, config.DefaultProfile, "Profile")

	return cmd
}
