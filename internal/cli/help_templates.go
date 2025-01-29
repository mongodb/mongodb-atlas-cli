package cli

const (
	// Help template
	// Modified version of https://github.com/spf13/cobra/blob/01ffff4eca5a08384ef2b85f39ec0dac192a5f7b/command.go#L595 which shows both .Short and .Long help descriptions.
	HelpTemplate = `{{.Short | trimTrailingWhitespaces}} {{.Long | trimTrailingWhitespaces}}

{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

	ExperimentalHelpTemplate = "experimental: " + HelpTemplate
)
