package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"slices"
	"time"

	"github.com/PaesslerAG/jsonpath"
)

var (
	ErrWatcherFailedToBuildRequest                   = errors.New("failed to build watcher request")
	ErrWatcherFailedToWatch                          = errors.New("failed to watch")
	ErrWatcherFailedToExecuteWatchRequest            = errors.New("failed to execute watch request")
	ErrWatcherGetCommandNotFound                     = errors.New("watcher get command not found")
	ErrWatcherFailedToApplyJSONPathToWatcherResponse = errors.New("failed to apply json path to watcher response")
)

const (
	WatcherWatchInterval = 1 * time.Second
)

type Watcher struct {
	executor CommandExecutor
}

func (w *Watcher) Watch(ctx context.Context, props WatcherProperties) error {
	// TODO: add timeout support, separate ticket (also add flag)
	// ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// defer cancel()
	request, err := buildRequest(Commands, props)
	if err != nil || request == nil {
		return errors.Join(ErrWatcherFailedToBuildRequest, err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			done, err := watchInner(ctx, w.executor, props.Expect, *request)
			if err != nil {
				return errors.Join(ErrWatcherFailedToWatch, err)
			}

			if done {
				return nil
			}

			time.Sleep(WatcherWatchInterval)
		}
	}
}

func buildRequest(allCommands GroupedAndSortedCommands, props WatcherProperties) (*CommandRequest, error) {
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

	requestParams := make(map[string][]string, len(props.Get.Params))
	for key, value := range props.Get.Params {
		requestParams[key] = []string{value}
	}

	return &CommandRequest{
		Command:    *command,
		Format:     "json",
		Parameters: requestParams,
		Version:    props.Get.Version,
	}, nil
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

func applyJSONPath(readerCloser io.ReadCloser, jsonPath string) (string, error) {
	var data any
	if err := json.NewDecoder(readerCloser).Decode(&data); err != nil {
		return "", err
	}

	output, err := jsonpath.Get(jsonPath, data)
	if err != nil {
		return "", err
	}

	return anyToString(output), nil
}

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
