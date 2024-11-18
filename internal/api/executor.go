package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
)

var (
	ErrMissingDependency            = errors.New("missing executor dependency")
	ErrFailedToGetBaseURL           = errors.New("failed to get base url")
	ErrFailedToConvertToHTTPRequest = errors.New("failed to convert to HTTP request")
	ErrFailedToAccessToken          = errors.New("failed to get access token")
	ErrFailedToExecuteHTTPRequest   = errors.New("failed to execute HTTP request")
)

type Executor struct {
	commandConverter CommandConverter
	httpClient       *http.Client
}

// We're expecting a http client that's authenticated.
func NewExecutor(commandConverter CommandConverter, httpClient *http.Client) (*Executor, error) {
	if commandConverter == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("commandConverter is nil"))
	}

	if httpClient == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("httpClient is nil"))
	}

	return &Executor{
		commandConverter: commandConverter,
		httpClient:       httpClient,
	}, nil
}

// Executor wired up to use the default profile and static functions on config.
func NewDefaultExecutor() (*Executor, error) {
	profile := config.Default()

	client := &http.Client{
		Transport: authenticatedTransport(profile, store.DefaultTransport),
	}

	configWrapper := NewAuthenticatedConfigWrapper(profile)
	commandConverter, err := NewDefaultCommandConverter(configWrapper)
	if err != nil {
		return nil, err
	}

	return NewExecutor(
		commandConverter,
		client,
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

	// TODO: CLOUDP-280747, formatting if isSuccess == true
	response := CommandResponse{
		IsSuccess: isSuccess,
		Output:    httpResponse.Body,
	}

	return &response, nil
}
