package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/root"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	errFlagDeleted        = errors.New("flag was deleted")
	errFlagTypeChanged    = errors.New("flag type changed")
	errFlagDefaultChanged = errors.New("flag default value changed")
	errCmdDeleted         = errors.New("command deleted")
	errCmdRemovedAlias    = errors.New("command alias removed")
	errFlagShortChanged   = errors.New("flag shorthand changed")
)

type flagData struct {
	Type    string `json:"type"`
	Default string `json:"default"`
	Short   string `json:"short"`
}

type cmdData struct {
	Aliases []string            `json:"aliases"`
	Flags   map[string]flagData `json:"flags"`
}

func generateCmd(cmd *cobra.Command) cmdData {
	data := cmdData{}
	if len(cmd.Aliases) > 0 {
		data.Aliases = cmd.Aliases
	}
	flags := false
	data.Flags = map[string]flagData{}
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		data.Flags[f.Name] = flagData{
			Type:    f.Value.Type(),
			Default: f.DefValue,
			Short:   f.Shorthand,
		}
		flags = true
	})
	if !flags {
		data.Flags = nil
	}
	return data
}

func generateCmds(cmd *cobra.Command) map[string]cmdData {
	data := map[string]cmdData{}
	data[cmd.CommandPath()] = generateCmd(cmd)
	for _, c := range cmd.Commands() {
		for k, v := range generateCmds(c) {
			data[k] = v
		}
	}
	return data
}

func generateCmdRun(output io.Writer) error {
	cliCmd := root.Builder()
	data := generateCmds(cliCmd)
	return json.NewEncoder(output).Encode(data)
}

func compareFlags(cmdPath string, mainFlags, changedFlags map[string]flagData) []error {
	if mainFlags == nil {
		return nil
	}

	changes := []error{}

	for flagName, flagValue := range mainFlags {
		changedFlagValue, ok := changedFlags[flagName]
		if !ok {
			changes = append(changes, fmt.Errorf("%w: %s --%s", errFlagDeleted, cmdPath, flagName))
			continue
		}

		if flagValue.Type != changedFlagValue.Type {
			changes = append(changes, fmt.Errorf("%w: %s --%s", errFlagTypeChanged, cmdPath, flagName))
		}

		if flagValue.Default != changedFlagValue.Default {
			changes = append(changes, fmt.Errorf("%w: %s --%s", errFlagDefaultChanged, cmdPath, flagName))
		}

		if flagValue.Short != changedFlagValue.Short {
			changes = append(changes, fmt.Errorf("%w: %s --%s", errFlagShortChanged, cmdPath, flagName))
		}
	}

	return changes
}

func compareCmds(changedData, mainData map[string]cmdData) error {
	changes := []error{}
	for cmdPath, mv := range mainData {
		cv, ok := changedData[cmdPath]
		if !ok {
			changes = append(changes, fmt.Errorf("%w: %s", errCmdDeleted, cmdPath))
			continue
		}

		if mv.Aliases != nil {
			mainAliases := mv.Aliases
			changedAliases := cv.Aliases

			for _, alias := range mainAliases {
				if !slices.Contains(changedAliases, alias) {
					changes = append(changes, fmt.Errorf("%w: %s", errCmdRemovedAlias, cmdPath))
				}
			}
		}

		changes = append(changes, compareFlags(cmdPath, mv.Flags, cv.Flags)...)
	}

	if len(changes) > 0 {
		return errors.Join(changes...)
	}

	return nil
}

func validateCmdRun(output io.Writer, input io.Reader) error {
	var inputData map[string]cmdData
	if err := json.NewDecoder(input).Decode(&inputData); err != nil {
		return err
	}

	cliCmd := root.Builder()
	generatedData := generateCmds(cliCmd)

	err := compareCmds(generatedData, inputData)
	if err != nil {
		return err
	}

	fmt.Fprintln(output, "no breaking changes detected")
	return nil
}

func buildCmd() *cobra.Command {
	generateCmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate the CLI command structure.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return generateCmdRun(cmd.OutOrStdout())
		},
	}

	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate the CLI command structure.",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return validateCmdRun(cmd.OutOrStdout(), cmd.InOrStdin())
		},
	}

	rootCmd := &cobra.Command{
		Use:   "breakvalidator",
		Short: "CLI tool to validate breaking changes in the CLI.",
	}
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(validateCmd)

	return rootCmd
}

func main() {
	rootCmd := buildCmd()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
