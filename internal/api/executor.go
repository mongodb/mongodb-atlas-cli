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
	"context"
	"errors"
	"net/http"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	storeTransport "github.com/mongodb/mongodb-atlas-cli/atlascli/internal/transport"
)

var (
	ErrFailedToAccessToken          = errors.New("failed to get access token")
	ErrFailedToConvertToHTTPRequest = errors.New("failed to convert to HTTP request")
	ErrFailedToExecuteHTTPRequest   = errors.New("failed to execute HTTP request")
	ErrFailedToGetBaseURL           = errors.New("failed to get base url")
	ErrFailedToHandleFormat         = errors.New("failed to handle format")
	ErrMissingDependency            = errors.New("missing executor dependency")
)

type Executor struct {
	commandConverter CommandConverter
	httpClient       *http.Client
	formatter        ResponseFormatter
}

// We're expecting a http client that's authenticated.
func NewExecutor(commandConverter CommandConverter, httpClient *http.Client, formatter ResponseFormatter) (*Executor, error) {
	if commandConverter == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("commandConverter is nil"))
	}

	if httpClient == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("httpClient is nil"))
	}

	if formatter == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("formatter is nil"))
	}

	return &Executor{
		commandConverter: commandConverter,
		httpClient:       httpClient,
		formatter:        formatter,
	}, nil
}

// Executor wired up to use the default profile and static functions on config.
func NewDefaultExecutor(formatter ResponseFormatter) (*Executor, error) {
	profile := config.Default()

	client := &http.Client{
		Transport: authenticatedTransport(profile, storeTransport.Default()),
	}

	configWrapper := NewAuthenticatedConfigWrapper(profile)
	commandConverter, err := NewDefaultCommandConverter(configWrapper)
	if err != nil {
		return nil, err
	}

	return NewExecutor(
		commandConverter,
		client,
		formatter,
	)
}

func (e *Executor) ensureInitialized() {
	if e.commandConverter == nil || e.httpClient == nil {
		// panic because this is developer error, not user error
		// should never happen
		panic("the executor was not properly initialized, use the NewExecutor method to initialize this struct")
	}
}

func (e *Executor) ExecuteCommand(ctx context.Context, commandRequest CommandRequest) (*CommandResponse, error) {
	e.ensureInitialized()

	// Set the content type
	if err := e.SetContentType(&commandRequest); err != nil {
		return nil, err
	}

	// Convert the request (api command definition + execution context) into a http request
	httpRequest, err := e.commandConverter.ConvertToHTTPRequest(commandRequest)
	if err != nil {
		return nil, errors.Join(ErrFailedToBuildHTTPRequest, err)
	}

	// Set the context, so we can cancel the request
	httpRequest = httpRequest.WithContext(ctx)

	// Execute the request
	httpResponse, err := e.httpClient.Do(httpRequest)
	if err != nil {
		return nil, errors.Join(ErrFailedToConvertToHTTPRequest, err)
	}

	//nolint: mnd // httpResponse.StatusCode >= StatusOK && httpResponse.StatusCode < StatusMultipleChoices makes this code harder to read
	isSuccess := httpResponse.StatusCode >= 200 && httpResponse.StatusCode < 300
	httpCode := httpResponse.StatusCode
	output := httpResponse.Body

	response := CommandResponse{
		IsSuccess: isSuccess,
		HTTPCode:  httpCode,
		Output:    output,
	}

	return &response, nil
}

func (e *Executor) SetContentType(commandRequest *CommandRequest) error {
	e.ensureInitialized()

	// Update the format if needed
	// For example if the requested format is a go template, change the request format to json
	contentType, err := e.formatter.ContentType(commandRequest.Format)
	if err != nil {
		return errors.Join(ErrFailedToHandleFormat, err)
	}
	commandRequest.ContentType = contentType

	return nil
}
