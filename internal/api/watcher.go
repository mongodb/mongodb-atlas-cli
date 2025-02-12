// Copyright 2025 MongoDB Inc
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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"slices"
	"strings"
	"time"

	"github.com/PaesslerAG/jsonpath"
)

var (
	ErrWatcherFailedToBuildRequest                           = errors.New("failed to build watcher request")
	ErrWatcherFailedToBuildRequestParameters                 = errors.New("failed to build watcher request parameters")
	ErrWatcherFailedToBuildRequestParametersInvalidParameter = errors.New("invalid parameter")
	ErrWatcherFailedToBuildRequestFailedToApplyJSONPath      = errors.New("request parameter failed to apply jsonpath")
	ErrWatcherFailedToBuildRequestInputParameterNotFound     = errors.New("input parameter not found")
	ErrWatcherFailedToBuildRequestInvalidParameterOperation  = errors.New("parameter operation not found")
	ErrWatcherFailedToExecuteWatchRequest                    = errors.New("failed to execute watch request")
	ErrWatcherGetCommandNotFound                             = errors.New("watcher get command not found")
	ErrWatcherFailedToApplyJSONPathToWatcherResponse         = errors.New("failed to apply json path to watcher response")
)

const (
	WatcherWatchInterval = 1 * time.Second
)

// Watcher provides the functionality to watch a certain operation
//
// The main goal of the watcher is to execute one operation over and verify that the output matches `expect`
// This is the API layer watcher, so this watcher struct is only responsible for execting the operation and verifying operation output.
type Watcher struct {
	executor CommandExecutor
	request  *CommandRequest
	expect   *WatcherExpectProperties
}

// Create a new watcher
//
//   - executor: executor used to make the API calls
//   - requestParams: parameters used for the preceding call
//     for example we're watching until a cluster is deleted then requestParams would contain the parameters which
//     were used to execute the deleteCluster operation.
//   - responseBody: the body from the preceding call
//   - props: the watcher configuration, this is generated using the api-generator tool
func NewWatcher(executor CommandExecutor, requestParams map[string][]string, responseBody []byte, props WatcherProperties) (*Watcher, error) {
	// We're making the same request over and over again, so we only have to build it once
	request, err := buildRequest(Commands, requestParams, responseBody, props)
	if err != nil || request == nil {
		return nil, errors.Join(ErrWatcherFailedToBuildRequest, err)
	}

	return &Watcher{
		executor: executor,
		request:  request,
		expect:   props.Expect,
	}, nil
}

// Builds the request we're going to send over and over again until it returns what we expect
//
// - `allCommands`, array with all command definitions, realisticly this will always be the static `Commands` variable, but this allows easier unit testing
// See `NewWatcher` for other parameter descriptions.
func buildRequest(allCommands GroupedAndSortedCommands, requestParams map[string][]string, responseBody []byte, props WatcherProperties) (*CommandRequest, error) {
	// Search for the command definition that we're expected to execute
	var command *Command
	for _, commandGroup := range allCommands {
		for _, loopCommand := range commandGroup.Commands {
			if loopCommand.OperationID == props.Get.OperationID {
				command = &loopCommand
			}
		}
	}

	if command == nil {
		return nil, fmt.Errorf("%w (OperationID=%s)", ErrWatcherGetCommandNotFound, props.Get.OperationID)
	}

	// Transform the previous request context into parameters for our watch request
	watcherRequestParams, err := buildRequestParameters(requestParams, responseBody, props.Get.Params)
	if err != nil {
		return nil, errors.Join(ErrWatcherFailedToBuildRequestParameters, err)
	}

	return &CommandRequest{
		Command:    *command,
		Format:     "json",
		Parameters: watcherRequestParams,
		Version:    props.Get.Version,
	}, nil
}

// This function takes the preceding commands context (requestParams, responseBody) and turns it into parameters for the watcher request
//
// Right now watchers support 3 types of functions/transformations:
// - `input:[variable_name]`: copy the parameter from requestParams
// - `body:[json_path]`: execute a jsonpath query on `responseBody` and use that value
// - `const:[value]`: use a const value.
func buildRequestParameters(requestParams map[string][]string, responseBody []byte, watcherParams map[string]string) (map[string][]string, error) {
	watcherRequestParams := make(map[string][]string, len(watcherParams))

	// we don't care if the serialization fails or not
	// the response doesn't have to be json if the watcherParameters don't use the "body:" function
	// if the watcherParameters do use the "body:" function then we'll return an error in that place
	var response any
	_ = json.Unmarshal(responseBody, &response)

	// Loop through all the watch parameters we need to create
	for watcherParameterKey, watcherParameterValue := range watcherParams {
		// Split the parameter value so we can use it as follows: `[function]:[args]`
		parameterFunction, parameterArgs, found := strings.Cut(watcherParameterValue, ":")
		if !found {
			return nil, fmt.Errorf("%w: %s", ErrWatcherFailedToBuildRequestParametersInvalidParameter, watcherParameterValue)
		}

		// Apply the function
		switch parameterFunction {
		case "body":
			value, err := applyJSONPathToObject(response, parameterArgs)
			if err != nil {
				return nil, errors.Join(ErrWatcherFailedToBuildRequestFailedToApplyJSONPath, err)
			}
			watcherRequestParams[watcherParameterKey] = []string{value}
		case "input":
			values, ok := requestParams[parameterArgs]
			if !ok {
				return nil, fmt.Errorf("%w: %s", ErrWatcherFailedToBuildRequestInputParameterNotFound, watcherParameterKey)
			}
			watcherRequestParams[watcherParameterKey] = values
		case "const":
			watcherRequestParams[watcherParameterKey] = []string{parameterArgs}
		default:
			return nil, fmt.Errorf("%w: '%s'", ErrWatcherFailedToBuildRequestInvalidParameterOperation, parameterFunction)
		}
	}

	return watcherRequestParams, nil
}

// This function will make the API request once and return true if the expected output was returned.
// - This function does not sleep
// - This function only fails when the HTTP call fails, not when the endpoint returns a non-OK response.
func (w *Watcher) WatchOne(ctx context.Context) (bool, error) {
	return watchInner(ctx, w.executor, w.expect, *w.request)
}

// Actual watcher logic without handling.
func watchInner(ctx context.Context, executor CommandExecutor, expect *WatcherExpectProperties, commandRequest CommandRequest) (bool, error) {
	// execute the command using the executor
	response, err := executor.ExecuteCommand(ctx, commandRequest)
	if err != nil {
		return false, errors.Join(ErrWatcherFailedToExecuteWatchRequest, err)
	}

	// verify the HTTPCode, 200 is the default when nothing is provided
	expectedHTTPCode := 200
	if expect != nil && expect.HTTPCode != 0 {
		expectedHTTPCode = expect.HTTPCode
	}

	if expectedHTTPCode != response.HTTPCode {
		return false, nil
	}

	// check if we need to match the contents or not
	if expect != nil && expect.Match != nil && expect.Match.Values != nil {
		// get the value from the response.output
		value, err := applyJSONPath(response.Output, expect.Match.Path)
		if err != nil {
			return false, errors.Join(ErrWatcherFailedToApplyJSONPathToWatcherResponse, err)
		}

		// check if the value is one of the expected values
		return slices.Contains(expect.Match.Values, value), nil
	}

	// if we get here, everything is ok!
	return true, nil
}

// Apply a JSON path to a reader which contains a JSON stream.
func applyJSONPath(readerCloser io.ReadCloser, jsonPath string) (string, error) {
	var data any
	if err := json.NewDecoder(readerCloser).Decode(&data); err != nil {
		return "", err
	}

	return applyJSONPathToObject(data, jsonPath)
}

func applyJSONPathToObject(object any, jsonPath string) (string, error) {
	output, err := jsonpath.Get(jsonPath, object)
	if err != nil {
		return "", err
	}

	return anyToString(output), nil
}

// Convert any type to string, for slices we take the first value.
func anyToString(v any) string {
	if v == nil {
		return ""
	}

	// Handle slice types
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		slice := reflect.ValueOf(v)
		if slice.Len() > 0 {
			return fmt.Sprint(slice.Index(0).Interface())
		}
		return ""
	}

	return fmt.Sprint(v)
}
