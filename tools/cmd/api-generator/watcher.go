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

package main

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

var (
	ErrWatcherGetPropertiesExtIsNil           = errors.New("extension map is nil")
	ErrWatcherGetPropertiesInvalidOperationID = errors.New("invalid OperationID")
	ErrWatcherMatchPropertiesPathIsMissing    = errors.New("path is empty or missing")
	ErrWatcherMatchPropertiesValuesAreMissing = errors.New("values are empty or missing")
)

/* This is what the YAML looks like: */
// watcher:
//	get:
//		operation-id: getCluster
//		version: 2023-02-01
//		params:
//			- groupId: input:groupId
//			- clusterName: body:$.name
//	expect:
//		http-code: 200
//		match:
//			path: $.stateName
//			values: IDLE,...

func extractWatcherProperties(ext map[string]any) (*api.WatcherProperties, error) {
	cliExt := extractObject(ext, "x-xgen-atlascli")
	if cliExt == nil {
		return nil, nil
	}

	watcherExt := extractObject(cliExt, "watcher")
	if watcherExt == nil {
		return nil, nil
	}

	getExt := extractObject(watcherExt, "get")
	if getExt == nil {
		return nil, nil
	}

	get, err := newWatcherGetProperties(getExt)
	if err != nil {
		return nil, err
	}

	var expect *api.WatcherExpectProperties
	expectExt := extractObject(watcherExt, "expect")
	if expectExt != nil {
		expect, err = newWatcherExpectProperties(expectExt)
		if err != nil {
			return nil, err
		}
	}

	return &api.WatcherProperties{
		Get:    *get,
		Expect: expect,
	}, nil
}

func newWatcherGetProperties(ext map[string]any) (*api.WatcherGetProperties, error) {
	if ext == nil {
		return nil, ErrWatcherGetPropertiesExtIsNil
	}

	operationID, operationIDOk := ext["operation-id"].(string)
	if operationID == "" || !operationIDOk {
		return nil, ErrWatcherGetPropertiesInvalidOperationID
	}

	versionString, _ := ext["version"].(string)
	var version api.Version
	if versionString != "" {
		parsedVersion, err := api.ParseVersion(versionString)
		if err != nil {
			return nil, err
		}

		version = parsedVersion
	}

	params := make(map[string]string)
	for key, value := range extractObject(ext, "params") {
		if stringValue, ok := value.(string); ok {
			params[key] = stringValue
		}
	}

	return &api.WatcherGetProperties{
		OperationID: operationID,
		Version:     version,
		Params:      params,
	}, nil
}

func newWatcherExpectProperties(ext map[string]any) (*api.WatcherExpectProperties, error) {
	httpCode := 0
	if httpCodeValue, ok := ext["http-code"]; ok {
		// ext["http-code"] is passed from the yaml converter
		// the converter is entirely free to decide which integer type it returns, could be any numeric type
		var err error
		if httpCode, err = toInt(httpCodeValue); err != nil {
			return nil, err
		}
	}

	var match *api.WatcherMatchProperties
	matchExt := extractObject(ext, "match")
	if matchExt != nil {
		var err error
		match, err = newWatcherMatchProperties(matchExt)
		if err != nil {
			return nil, err
		}
	}

	return &api.WatcherExpectProperties{
		HTTPCode: httpCode,
		Match:    match,
	}, nil
}

func newWatcherMatchProperties(ext map[string]any) (*api.WatcherMatchProperties, error) {
	if ext == nil {
		return nil, nil
	}

	path, pathOk := ext["path"].(string)
	if path == "" || !pathOk {
		return nil, ErrWatcherMatchPropertiesPathIsMissing
	}

	valuesAny, valuesOk := ext["values"].([]any)
	if !valuesOk || valuesAny == nil {
		return nil, ErrWatcherMatchPropertiesValuesAreMissing
	}

	values := make([]string, 0, len(valuesAny))
	for _, value := range valuesAny {
		if stringValue, ok := value.(string); ok {
			values = append(values, stringValue)
		}
	}

	return &api.WatcherMatchProperties{
		Path:   path,
		Values: values,
	}, nil
}

func extractObject(ext map[string]any, name string) map[string]any {
	if object, ok := ext[name].(map[string]any); ok && object != nil {
		return object
	}

	return nil
}

func toInt(value any) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		//nolint:gosec
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		//nolint:gosec
		return int(v), nil
	case float32:
		if float32(int(v)) != v {
			return 0, fmt.Errorf("value %v has decimal places", v)
		}
		return int(v), nil
	case float64:
		if float64(int(v)) != v {
			return 0, fmt.Errorf("value %v has decimal places", v)
		}
		return int(v), nil
	default:
		return 0, fmt.Errorf("value %v of type %T cannot be converted to int", value, value)
	}
}

func validateAllWatchers(allCommands map[string]*api.Group) error {
	var errs []error

	for _, group := range allCommands {
		for _, command := range group.Commands {
			if err := validateWatchersForCommand(allCommands, command); err != nil {
				errs = append(errs, err)
			}
		}
	}

	// If no errors occurred, return nil
	if len(errs) == 0 {
		return nil
	}

	// Join all errors into a single error
	return errors.Join(errs...)
}

//nolint:gocyclo
func validateWatchersForCommand(allCommands map[string]*api.Group, command api.Command) error {
	watcher := command.Watcher
	if watcher == nil {
		return nil
	}

	baseErr := fmt.Errorf("watcher for operationID='%s' is invalid", command.OperationID)

	// ensure the OperationID is not empty
	operationID := watcher.Get.OperationID
	if operationID == "" {
		return fmt.Errorf("%w: the watcher get operation operationID is empty", baseErr)
	}

	// ensure the command the watcher defines exists
	var getWatcherCommand *api.Command
	for _, group := range allCommands {
		for _, command := range group.Commands {
			if command.OperationID == operationID {
				getWatcherCommand = &command
			}
		}
	}

	if getWatcherCommand == nil {
		return fmt.Errorf("%w: the watcher get operation with operationID '%s' was not found", baseErr, operationID)
	}

	// ensure the version exists
	versionFound := false
	for _, apiVersion := range getWatcherCommand.Versions {
		if apiVersion.Version.Equal(watcher.Get.Version) {
			versionFound = true
		}
	}

	if !versionFound {
		return fmt.Errorf("%w: the watcher get operation with operationID '%s' was found, but the version '%s' was not found", baseErr, operationID, watcher.Get.Version)
	}

	// verify that all the parameters exist, and that required parameters are set
	// make 2 sets we'll fill up with parameters from getWatcherCommand
	parameterNames := make(map[string]struct{})
	requiredParameterNames := make(map[string]struct{})

	addParameters := func(p []api.Parameter) {
		for _, parameter := range p {
			parameterNames[parameter.Name] = struct{}{}

			if parameter.Required {
				requiredParameterNames[parameter.Name] = struct{}{}
			}
		}
	}

	addParameters(getWatcherCommand.RequestParameters.QueryParameters)
	addParameters(getWatcherCommand.RequestParameters.URLParameters)

	// 1. verify that all parameters used exist
	// 2. remove items from requiredParameterNames, set should be empty after this for loop
	for parameterName := range watcher.Get.Params {
		if _, valid := parameterNames[parameterName]; !valid {
			return fmt.Errorf("%w: invalid parameter was provided, parameter does not exist: '%s'", baseErr, parameterName)
		}

		delete(requiredParameterNames, parameterName)
	}

	if len(requiredParameterNames) > 0 {
		missingRequiredParameters := ""
		for name := range requiredParameterNames {
			if missingRequiredParameters != "" {
				missingRequiredParameters += ", "
			}

			missingRequiredParameters += name
		}

		return fmt.Errorf("%w: some required parameter(s) are missing: '%s'", baseErr, missingRequiredParameters)
	}

	return nil
}
