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
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	ErrFailedToSetUntouchedFlags         = errors.New("failed to set untouched flags")
	ErrServerReturnedAnErrorResponseCode = errors.New("server returned an error response code")
	ErrAPICommandsHasNoVersions          = errors.New("api command has no versions")
	ErrFormattingOutput                  = errors.New("error formatting output")
	ErrRunningWatcher                    = errors.New("error while running watcher")
	BinaryOutputTypes                    = []string{"gzip"}
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
	rootCmd := &cobra.Command{
		Use:   "api",
		Short: `Access all features of the Atlas Administration API through the Atlas CLI by using the 'atlas api <tag> <operationId>' command.`,
		Long: `This feature streamlines script development by letting you interact directly with any Atlas Administration API endpoint through the Atlas CLI.

For more information on
- Atlas Administration API see: https://www.mongodb.com/docs/api/doc/atlas-admin-api-v2/
- Getting started with the Atlas Administration API: https://www.mongodb.com/docs/atlas/configure-api-access/#std-label-atlas-admin-api-access`,
	}

	rootCmd.SetHelpTemplate(cli.APICommandHelpTemplate)

	return rootCmd
}

func createAPICommandGroupToCobraCommand(group shared_api.Group) *cobra.Command {
	groupName := strcase.ToLowerCamel(group.Name)
	shortDescription, longDescription := splitShortAndLongDescription(group.Description)

	return &cobra.Command{
		Use:   groupName,
		Short: shortDescription,
		Long:  longDescription,
	}
}

//nolint:gocyclo
func convertAPIToCobraCommand(command shared_api.Command) (*cobra.Command, error) {
	// command properties
	commandName := strcase.ToLowerCamel(command.OperationID)
	commandOperationID := command.OperationID
	commandAliases := command.Aliases

	if command.ShortOperationID != "" {
		// Add original operation ID to aliases
		commandAliases = append(commandAliases, commandName)
		// Use shortOperationID as the command name
		commandName = strcase.ToLowerCamel(command.ShortOperationID)
	}

	shortDescription, longDescription := splitShortAndLongDescription(command.Description)

	// flag values
	file := ""
	format := ""
	outputFile := ""
	version, err := defaultAPIVersion(command)
	watch := false
	watchTimeout := int64(0)
	if err != nil {
		return nil, err
	}

	cmd := &cobra.Command{
		Use:     commandName,
		Aliases: commandAliases,
		Short:   shortDescription,
		Long:    longDescription,
		Annotations: map[string]string{
			"operationId": commandOperationID,
		},
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			// Go through all commands that have not been touched/modified by the user and try to populate them from the users profile
			// Common usecases:
			// - set orgId
			// - set projectId
			// - default api version
			if err := setUnTouchedFlags(NewProfileFlagValueProviderForDefaultProfile(), cmd); err != nil {
				return errors.Join(ErrFailedToSetUntouchedFlags, err)
			}

			// Remind the user to pin their api command to a specific version to avoid breaking changes
			remindUserToPinVersion(cmd)

			// Reset version to default if unsupported version was selected
			// This can happen when the profile contains a default version which is not supported for a specific endpoint
			ensureVersionIsSupported(command, &version)

			// Print a warning if the version is a preview version
			printPreviewWarning(command, &version)

			// Print a warning if the version has a sunset date
			printSunsetWarning(command, &version)

			// Print a warning if the version is deprecated
			printDeprecatedWarning(command, &version)

			// Detect if stdout is being piped (atlas api myTag myOperationId > output.json)
			isPiped, err := IsStdOutPiped()
			if err != nil {
				return err
			}

			// If the selected output format is binary and stdout is not being piped, mark output as required
			// This ensures that the console isn't flooded with binary contents (for example gzip contents)
			if slices.Contains(BinaryOutputTypes, format) && !isPiped {
				if err := cmd.MarkFlagRequired(flag.Output); err != nil {
					return err
				}
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Get the request input if needed
			// This is needed for most PATCH/POST/PUT requests
			var content io.ReadCloser
			if needsFileFlag(command) {
				content, err = handleInput(cmd)
				if err != nil {
					return err
				}
				defer content.Close()
			}

			// Create a new executor
			// This is the piece of code which knows how to execute api.Commands
			formatter := api.NewFormatter()
			executor, err := api.NewDefaultExecutor(formatter)
			if err != nil {
				return err
			}

			// Convert the version string into an api version
			apiVersion, err := shared_api.ParseVersion(version)
			if err != nil {
				// This should never happen, the version is already validated in the prerun function
				return err
			}

			// Convert the api command + cobra command into a api command request
			commandRequest, err := NewCommandRequestFromCobraCommand(cmd, command, content, format, apiVersion)
			if err != nil {
				return err
			}

			// Execute the api command request
			// This function will return an error if the http request failed
			// When the http request returns a non-success code error will still be nil
			result, err := executor.ExecuteCommand(cmd.Context(), *commandRequest)
			if err != nil {
				return err
			}

			// Properly free up result output
			defer result.Output.Close()

			// Output that will be send to stdout/file
			responseOutput := result.Output

			// Response body used for the watcher
			var watchResponseBody []byte

			// If the response was successful, handle --format
			if result.IsSuccess {
				// If we're watching, we need to cache the original output before formatting so we don't read twice from the same reader
				// In case we're not watching the http output will be piped straight into the formatter, which should be a little more memory efficient
				if watch {
					responseBytes, err := io.ReadAll(result.Output)
					if err != nil {
						return errors.Join(errors.New("failed to read output"), err)
					}
					watchResponseBody = responseBytes

					// Create a new reader for the formatter
					responseOutput = io.NopCloser(bytes.NewReader(responseBytes))
				}

				formattedOutput, err := formatter.Format(format, responseOutput)
				if err != nil {
					return errors.Join(ErrFormattingOutput, err)
				}

				responseOutput = formattedOutput
			}

			// Determine where to write the
			output, err := getOutputWriteCloser(outputFile)
			if err != nil {
				return err
			}
			// Properly free up output
			defer output.Close()

			// Write the output
			_, err = io.Copy(output, responseOutput)
			if err != nil {
				return err
			}

			// In case the http status code was non-success
			// Return an error, this causes the CLI to exit with a non-zero exit code while still running all telemetry code
			if !result.IsSuccess {
				return ErrServerReturnedAnErrorResponseCode
			}

			// In case watcher is set we wait for the watcher to succeed before we exit the program
			if watch {
				// Create a new watcher
				watcher, err := NewWatcher(executor, commandRequest.Parameters, watchResponseBody, *command.Watcher)
				if err != nil {
					return err
				}

				// Wait until we're in the desired state or until an error occures when watching
				if err := watcher.Wait(cmd.Context(), time.Duration(watchTimeout)); err != nil {
					return errors.Join(ErrRunningWatcher, err)
				}
			}

			return nil
		},
	}

	// Add deprecation message if all versions are sunset
	addDeprecationMessageIfNeeded(cmd, command)

	// Common flags
	addWatchFlagIfNeeded(cmd, command, &watch, &watchTimeout)
	addVersionFlag(cmd, command, &version)

	if needsFileFlag(command) {
		cmd.Flags().StringVar(&file, flag.File, "", "path to your API request file. Leave empty to use standard input instead - you must provide one or the other, but not both.")
	}

	// Add output flags:
	// - `--output`: desired output format, translates to ContentType. Can also be a go template
	// - `--output-file`: file where we want to write the output to
	if err := addOutputFlags(cmd, command, &format, &outputFile); err != nil {
		return nil, err
	}

	// Add URL parameters as flags
	if err := addParameters(cmd, command.RequestParameters.URLParameters); err != nil {
		return nil, err
	}

	// Add query parameters as flags
	if err := addParameters(cmd, command.RequestParameters.QueryParameters); err != nil {
		return nil, err
	}

	// Handle parameter aliases
	cmd.Flags().SetNormalizeFunc(normalizeFlagFunc(command))

	return cmd, nil
}

func normalizeFlagFunc(command shared_api.Command) func(f *pflag.FlagSet, name string) pflag.NormalizedName {
	return func(_ *pflag.FlagSet, name string) pflag.NormalizedName {
		name = normalizeFlagName(command.RequestParameters.QueryParameters, name)
		name = normalizeFlagName(command.RequestParameters.URLParameters, name)

		return pflag.NormalizedName(name)
	}
}

func normalizeFlagName(parameters []shared_api.Parameter, name string) string {
	for _, parameter := range parameters {
		if slices.Contains(parameter.Aliases, name) {
			return parameter.Name
		}
	}

	return name
}

func addParameters(cmd *cobra.Command, parameters []shared_api.Parameter) error {
	for _, parameter := range parameters {
		if err := addFlag(cmd, parameter); err != nil {
			return err
		}
	}

	return nil
}

func setUnTouchedFlags(flagValueProvider FlagValueProvider, cmd *cobra.Command) error {
	var visitErr error
	cmd.NonInheritedFlags().VisitAll(func(f *pflag.Flag) {
		// There is no VisitAll which accepts an error
		// because of this we set visitErr when an error occures
		// if visitErr != nil we stop processing other flags
		if visitErr != nil {
			return
		}

		// Only update flags thave have been un-touched by the user
		if !f.Changed {
			value, err := flagValueProvider.ValueForFlag(f.Name)
			if err != nil {
				visitErr = err
			}

			// If we get a value back from the FlagValueProvider:
			// - Set the value
			// - Mark the flag as changed, this is needed to mark required flags as set
			if value != nil {
				if err := f.Value.Set(*value); err != nil {
					visitErr = err
					return
				}

				f.Changed = true
			}
		}
	})

	return visitErr
}

func handleInput(cmd *cobra.Command) (io.ReadCloser, error) {
	isPiped, err := IsStdInPiped()
	if err != nil {
		return nil, err
	}

	// If not piped, get the file flag
	filePath, err := cmd.Flags().GetString(flag.File)
	if err != nil {
		return nil, fmt.Errorf("error getting file flag: %w", err)
	}

	if isPiped {
		if filePath != "" {
			return nil, errors.New("cannot use --file flag and also input from standard input")
		}
		// Use stdin as the input
		return os.Stdin, nil
	}

	// Require file flag if not piped
	if filePath == "" {
		return nil, errors.New("--file flag is required when not using piped input")
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", filePath, err)
	}

	return file, nil
}

func IsStdInPiped() (bool, error) {
	return isPiped(os.Stdin)
}

func IsStdOutPiped() (bool, error) {
	return isPiped(os.Stdout)
}

func isPiped(file *os.File) (bool, error) {
	// Check if data is being piped to stdin
	info, err := file.Stat()
	if err != nil {
		return false, fmt.Errorf("isPiped, error checking: %w", err)
	}

	// Check if there's data in stdin (piped input)
	isPiped := (info.Mode() & os.ModeCharDevice) == 0

	return isPiped, nil
}

func defaultAPIVersion(command shared_api.Command) (string, error) {
	// Command versions are sorted by the generation tool
	nVersions := len(command.Versions)
	if nVersions == 0 {
		return "", ErrAPICommandsHasNoVersions
	}

	lastVersion := command.Versions[nVersions-1]
	return lastVersion.Version.String(), nil
}

func remindUserToPinVersion(cmd *cobra.Command) {
	versionFlag := cmd.Flag(flag.Version)
	// if we fail to get the version flag (which should never happen), then quit
	if versionFlag == nil {
		return
	}

	// check if the version flag is still in it's default state:
	// - not set by the user
	// - not set using api_version on the users profile
	// in that case, print a warning
	if !versionFlag.Changed {
		fmt.Fprintf(os.Stderr, "warning: using default API version '%s'; consider pinning a version to ensure consisentcy when updating the CLI\n", versionFlag.Value.String())
	}
}

func ensureVersionIsSupported(apiCommand shared_api.Command, versionString *string) {
	version, err := shared_api.ParseVersion(*versionString)

	// If the version is valid, check if it's supported
	if err == nil {
		for _, commandVersion := range apiCommand.Versions {
			if commandVersion.Version.Equal(version) {
				return
			}
		}
	}

	// if we get here it means that the picked version is not supported
	defaultVersion, err := defaultAPIVersion(apiCommand)
	// if we fail to get a version (which should never happen), then quit
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in 'ensureVersionIsSupported': default version has an invalid format '%s'\n", *versionString)
		return
	}

	fmt.Fprintf(os.Stderr, "warning: version '%s' is not supported for this endpoint, using default API version '%s'; consider pinning a version to ensure consisentcy when updating the CLI\n", *versionString, defaultVersion)
	*versionString = defaultVersion
}

func printPreviewWarning(apiCommand shared_api.Command, versionString *string) {
	version, err := shared_api.ParseVersion(*versionString)

	// If the version is invalid return, this should never happen
	if err != nil {
		fmt.Fprintf(os.Stderr, "error in 'printPreviewWarning': received an invalid version '%s'\n", *versionString)
		return
	}

	// If the version is not a preview version, return
	if version.StabilityLevel() != shared_api.StabilityLevelPreview {
		return
	}

	// Find the version in the command versions
	var commandVersion *shared_api.CommandVersion
	for _, cv := range apiCommand.Versions {
		if cv.Version.Equal(version) {
			commandVersion = &cv
			break
		}
	}

	// If the version is not found, return (should also never happen)
	if commandVersion == nil {
		return
	}

	if commandVersion.PublicPreview {
		fmt.Fprintf(os.Stderr, "warning: you've selected a public preview version of the endpoint, this version is subject to breaking changes.\n")
	} else {
		fmt.Fprintf(os.Stderr, "warning: you've selected a private preview version of the endpoint, this version might not be available for your account and is subject to breaking changes.\n")
	}
}

func printSunsetWarning(apiCommand shared_api.Command, versionString *string) {
	// only warn if the command is not fully deprecated, assume that if all versions are deprecated,
	// then the command will be marked as deprecated in Cobra.
	if allVersionsDeprecated(apiCommand) {
		return
	}

	version, err := shared_api.ParseVersion(*versionString)
	if err != nil {
		return
	}

	// Find the version in the command versions
	var commandVersion *shared_api.CommandVersion
	for i := range apiCommand.Versions {
		if apiCommand.Versions[i].Version.Equal(version) {
			commandVersion = &apiCommand.Versions[i]
			break
		}
	}

	// If the version is not found or has no sunset date, return
	if commandVersion == nil || commandVersion.Sunset == nil {
		return
	}

	sunsetDate := commandVersion.Sunset.Format("2006-01-02")

	// if date is in the past, warn the user that it will not work
	if commandVersion.Sunset.Before(time.Now()) {
		fmt.Fprintf(os.Stderr, "warning: version '%s' is deprecated for this command and has already been sunset since %s. Consider upgrading to a newer version if available.\n", *versionString, sunsetDate)
		return
	}
	fmt.Fprintf(os.Stderr, "warning: version '%s' is deprecated for this command and will be sunset on %s. Consider upgrading to a newer version if available.\n", *versionString, sunsetDate)
}

func printDeprecatedWarning(apiCommand shared_api.Command, versionString *string) {
	version, err := shared_api.ParseVersion(*versionString)
	if err != nil {
		return
	}

	// Find the version in the command versions
	var commandVersion *shared_api.CommandVersion
	for i := range apiCommand.Versions {
		if apiCommand.Versions[i].Version.Equal(version) {
			commandVersion = &apiCommand.Versions[i]
			break
		}
	}

	// If the version is not found or is not deprecated, return
	if commandVersion == nil || !commandVersion.Deprecated {
		return
	}

	// If the version has a sunset date, don't show a separate deprecation warning
	// as the sunset warning already indicates deprecation
	if commandVersion.Sunset != nil {
		return
	}

	fmt.Fprintf(os.Stderr, "warning: version '%s' is deprecated. Consider upgrading to a newer version.\n", *versionString)
}

func needsFileFlag(apiCommand shared_api.Command) bool {
	for _, version := range apiCommand.Versions {
		if version.RequestContentType != "" {
			return true
		}
	}

	return false
}

func addWatchFlagIfNeeded(cmd *cobra.Command, apiCommand shared_api.Command, watch *bool, watchTimeout *int64) {
	if apiCommand.Watcher == nil || apiCommand.Watcher.Get.OperationID == "" {
		return
	}

	cmd.Flags().BoolVarP(watch, flag.EnableWatch, flag.EnableWatchShort, false, usage.EnableWatchDefault)
	cmd.Flags().Int64Var(watchTimeout, flag.WatchTimeout, 0, usage.WatchTimeout)
}

func addVersionFlag(cmd *cobra.Command, apiCommand shared_api.Command, version *string) {
	// Create a unique list of all supported versions
	versions := make(map[string]struct{}, 0)
	for _, version := range apiCommand.Versions {
		versions[version.Version.String()] = struct{}{}
	}

	// Convert the keys of the map into a list
	supportedVersionsVersions := make([]string, 0, len(versions))
	for version := range versions {
		supportedVersionsVersions = append(supportedVersionsVersions, `"`+version+`"`)
	}

	// Sort the list
	slices.Sort(supportedVersionsVersions)

	// Convert the list to a string
	supportedVersionsVersionsString := strings.Join(supportedVersionsVersions, ", ")

	cmd.Flags().StringVar(version, flag.Version, *version, fmt.Sprintf("api version to use when calling the api call [options: %s], defaults to the latest version or the profiles api_version config value if set", supportedVersionsVersionsString))
}

func addOutputFlags(cmd *cobra.Command, apiCommand shared_api.Command, format *string, outputFile *string) error {
	// Get the list of supported content types for the apiCommand
	supportedContentTypesList := getContentTypes(&apiCommand)

	// If there's only one content type, set the format to that
	numSupportedContentTypes := len(supportedContentTypesList)
	if numSupportedContentTypes == 1 {
		*format = supportedContentTypesList[0]
	}

	// If the content type has json, also add go-template as an option to the --format help
	containsJSON := slices.Contains(supportedContentTypesList, "json")

	// Place quotes around every supported content type
	for i, value := range supportedContentTypesList {
		supportedContentTypesList[i] = `"` + value + `"`
	}

	// Add go-template, we add it here because we don't want go-template to be between quotes
	if containsJSON {
		supportedContentTypesList = append(supportedContentTypesList, "go-template")
	}

	// Generate a list of supported content types and add it to --help for --format
	// Example ['csv', 'json', go-template]
	supportedContentTypesString := strings.Join(supportedContentTypesList, ", ")

	// Set the flags
	cmd.Flags().StringVarP(format, flag.Output, flag.OutputShort, *format, fmt.Sprintf("preferred api format, can be [%s]", supportedContentTypesString))
	cmd.Flags().StringVar(outputFile, flag.OutputFile, "", "file to write the api output to. This flag is required when the output of an endpoint is binary (ex: gzip) and the command is not piped (ex: atlas command > out.zip)")

	// If there's multiple content types, mark --format as required
	if numSupportedContentTypes > 1 {
		if err := cmd.MarkFlagRequired(flag.Output); err != nil {
			return err
		}
	}

	return nil
}

func getContentTypes(apiCommand *shared_api.Command) []string {
	// Create a unique list of all supported content types
	// First create a map to convert 2 nested lists into a map
	supportedContentTypes := make(map[string]struct{}, 0)
	for _, version := range apiCommand.Versions {
		for _, contentType := range version.ResponseContentTypes {
			supportedContentTypes[contentType] = struct{}{}
		}
	}

	// Convert the keys of the map into a list
	supportedContentTypesList := make([]string, 0, len(supportedContentTypes))
	for contentType := range supportedContentTypes {
		supportedContentTypesList = append(supportedContentTypesList, contentType)
	}

	// Sort the list
	slices.Sort(supportedContentTypesList)

	return supportedContentTypesList
}

func getOutputWriteCloser(outputFile string) (io.WriteCloser, error) {
	// If an output file is specified, create/open the file and return the writer
	if outputFile != "" {
		//nolint: mnd
		file, err := os.OpenFile(outputFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return nil, err
		}
		return file, nil
	}

	// Return stdout by default
	return os.Stdout, nil
}

// addDeprecationMessageIfNeeded adds a deprecation message to the command if the command is deprecated.
func addDeprecationMessageIfNeeded(cmd *cobra.Command, apiCommand shared_api.Command) {
	allVersionsAreDeprecated := allVersionsDeprecated(apiCommand)

	if allVersionsAreDeprecated {
		cmd.Deprecated = "all of the available endpoint versions have been deprecated."

		if len(apiCommand.Versions) == 1 && apiCommand.Versions[0].Sunset != nil {
			version := apiCommand.Versions[0]
			sunsetDate := version.Sunset.Format("2006-01-02")
			cmd.Deprecated += fmt.Sprintf(" The API endpoint version %s will no longer be available after the sunset date of %s.", version.Version.String(), sunsetDate)
			return
		}

		for _, version := range apiCommand.Versions {
			if version.Sunset != nil {
				sunsetDate := version.Sunset.Format("2006-01-02")
				cmd.Deprecated += fmt.Sprintf(" The API endpoint version %s will no longer be available after the sunset date of %s.", version.Version.String(), sunsetDate)
				break
			}
		}

		cmd.Deprecated += "."
	}
}

// allVersionsDeprecated checks if all the versions are deprecated.
// we classify the command as deprecated if all the versions have a sunset date or all the versions are deprecated.
func allVersionsDeprecated(apiCommand shared_api.Command) bool {
	for _, version := range apiCommand.Versions {
		if version.Sunset == nil && !version.Deprecated {
			return false
		}
	}

	return true
}
