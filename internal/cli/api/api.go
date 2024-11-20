// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/spf13/cobra"
)

func Builder() *cobra.Command {
	apiCmd := createRootAPICommand()

	for _, tag := range api.Commands {
		tagCmd := createAPICommandGroupToCobraCommand(tag)

		for _, command := range tag.Commands {
			cobraCommand, err := convertAPIToCobraCommand(command)
			// err should always be nil, errors happening here should be covered by the generate api commands tool
			// if err != nil there is a bug in the converter tool
			if err != nil {
				_, _ = log.Warningf("failed to add command for operationId: %s, err: %s", command.OperationID, err)
				continue
			}

			tagCmd.AddCommand(cobraCommand)
		}

		apiCmd.AddCommand(tagCmd)
	}

	return apiCmd
}

func createRootAPICommand() *cobra.Command {
	return &cobra.Command{
		Use: "api",
	}
}

func createAPICommandGroupToCobraCommand(group api.Group) *cobra.Command {
	groupName := strcase.ToLowerCamel(group.Name)
	shortDescription, longDescription := splitShortAndLongDescription(group.Description)

	return &cobra.Command{
		Use:    groupName,
		Short:  shortDescription,
		Long:   longDescription,
		Hidden: true, // TODO: part of CLOUDP-280743 Polish autogenerated docs
	}
}

func convertAPIToCobraCommand(command api.Command) (*cobra.Command, error) {
	commandName := strcase.ToLowerCamel(command.OperationID)
	shortDescription, longDescription := splitShortAndLongDescription(command.Description)

	cmd := &cobra.Command{
		Use:   commandName,
		Short: shortDescription,
		Long:  longDescription,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			flag := cmd.Flag("groupId")
			if flag != nil {
				println("flag: %#v\n", flag)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			// TODO: wire up CLI and execution framework
			// These lines are just here for testing purposes (only works for commands which don't need flags like `atlas api cluster listClusters`)
			// Call "glue" layer which will glue commands and executor together: ticket CLOUDP-280745
			// This layer will call something among the lines of:
			executor, err := api.NewDefaultExecutor()
			if err != nil {
				return err
			}

			result, err := executor.ExecuteCommand(cmd.Context(), api.CommandRequest{
				Command:    command,
				Content:    cmd.InOrStdin(),
				Format:     "json",                // get from flags
				Parameters: map[string][]string{}, // get from flags + profile
				Version:    "2023-01-01",          // get from flags
			})

			if err != nil {
				return err
			}

			defer result.Output.Close()
			_, err = io.Copy(os.Stdout, result.Output)
			if err != nil {
				return err
			}

			// All this code is todo, will find a better way to get a non-zero error code as part of CLOUDP-280745
			if !result.IsSuccess {
				// Return an empty error so cobra doesn't print anything, temporary, will be part of CLOUDP-280745
				return errors.New("")
			}

			return nil
		},
	}

	if err := addParameters(cmd, command.RequestParameters.URLParameters); err != nil {
		return nil, err
	}

	if err := addParameters(cmd, command.RequestParameters.QueryParameters); err != nil {
		return nil, err
	}

	return cmd, nil
}

func addParameters(cmd *cobra.Command, parameters []api.Parameter) error {
	for _, parameter := range parameters {
		if cmd.Flag(parameter.Name) != nil {
			// this should never happen, the api command generation tool should cover this
			return fmt.Errorf("there is already a parameter with that name, name='%s'", parameter.Name)
		}

		flag := parameterToPFlag(parameter)
		cmd.Flags().AddFlag(flag)

		// TODO: part of CLOUDP-280745 handle values coming from profile
		if parameter.Required {
			_ = cmd.MarkFlagRequired(flag.Name)
		}
	}

	return nil
}

func splitShortAndLongDescription(description string) (string, string) {
	// Split on periods that are followed by a space or end of string
	// This approach allows us to not accidentally split verion numbers like 8.0
	split := regexp.MustCompile(`\.(?:\s+|$)`).Split(description, -1)

	// Short description is the first sentence
	shortDescription := split[0]

	// Add the dot back, if needed
	if shortDescription != "" && !strings.HasSuffix(shortDescription, ".") && !strings.HasSuffix(shortDescription, ". ") {
		shortDescription += "."
	}

	// Long descriptions is everything after the first sentence
	longDescription := ""

	if len(split) > 1 {
		// Remove all empty whitespace around sentences
		// This turns multi-line descriptions into single line
		for i, s := range split[1:] {
			split[i+1] = strings.TrimSpace(s)
		}

		// Add the ". " back
		longDescription = strings.Join(split[1:], ". ")
	}

	// Get rid of the last space after ". "
	longDescription = strings.TrimSpace(longDescription)

	return shortDescription, longDescription
}
