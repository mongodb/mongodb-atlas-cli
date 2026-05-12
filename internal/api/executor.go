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
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/terminal"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
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
	httpClient       Doer
	formatter        ResponseFormatter
	logger           Logger
}

// We're expecting a http client that's authenticated.
func NewExecutor(commandConverter CommandConverter, httpClient Doer, formatter ResponseFormatter, logger Logger) (*Executor, error) {
	if commandConverter == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("commandConverter is nil"))
	}

	if httpClient == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("httpClient is nil"))
	}

	if formatter == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("formatter is nil"))
	}

	if logger == nil {
		return nil, errors.Join(ErrMissingDependency, errors.New("logger is nil"))
	}

	return &Executor{
		commandConverter: commandConverter,
		httpClient:       httpClient,
		formatter:        formatter,
		logger:           logger,
	}, nil
}

// Executor wired up to use the default profile and static functions on config.
func NewDefaultExecutor(formatter ResponseFormatter) (*Executor, error) {
	profile := config.Default()

	client, err := transport.HTTPClient(version.Version, transport.Default())
	if err != nil {
		return nil, err
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
		log.Default(),
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

	if err := checkPledge(commandRequest); err != nil {
		return nil, err
	}

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
	e.logRequest(httpRequest)

	// Execute the request
	httpResponse, err := e.httpClient.Do(httpRequest)
	if err != nil {
		return nil, errors.Join(ErrFailedToConvertToHTTPRequest, err)
	}

	e.logResponse(httpResponse)

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

// Log the request if the logger is set to debug
// Copied behavior and format used in the SDK: https://github.com/mongodb/atlas-sdk-go/blob/b3fee40e236a8ff2a1f1c160b6984a242136dbe6/admin/client.go#L322
func (e *Executor) logRequest(httpRequest *http.Request) {
	if !e.logger.IsDebugLevel() {
		return
	}

	dump, err := httputil.DumpRequestOut(httpRequest, true)
	if err != nil {
		return
	}

	_, _ = e.logger.Debugf("\n%s\n", string(dump))
}

// checkPledge loads the active pledge for the current session and rejects the
// request if the command's permission tier exceeds what the pledge allows.
// When stdin is a TTY (human at a shell), the user is prompted inline to widen
// the session pledge. Non-TTY callers (agents, scripts) receive an error immediately.
func checkPledge(req CommandRequest) error {
	key, err := pledge.ResolveSessionKey()
	if err != nil {
		// Windows or no session support — skip enforcement.
		return nil
	}
	pf, err := pledge.Load(key)
	if err != nil {
		// No pledge set for this session.
		return nil
	}
	opID := req.Command.OperationID
	tier := req.Command.Permission
	outcome, _ := pledge.Check(pf, tier, opID)
	if outcome != pledge.Block {
		return nil
	}

	pledge.LogAudit(pledge.AuditEntry{
		SessionKeyStr: key.String(),
		OperationID:   opID,
		Outcome:       pledge.AuditBlocked,
	})

	if !terminal.IsTerminalInput(os.Stdin) {
		return requestOutOfBandApproval(key, pf, tier, opID)
	}

	return widenInteractiveExec(os.Stdin, os.Stderr, key, pf, tier, opID)
}

// requestOutOfBandApproval generates an approval token and returns an error
// containing it so the agent or user can run `atlas pledge allow <token>` in
// another terminal and then retry the command.
func requestOutOfBandApproval(key pledge.SessionKey, pf *pledge.PledgeFile, tier shared_api.PermissionTier, opID string) error {
	token, err := pledge.WriteApprovalRequest(key, opID, tier)
	if err != nil {
		// Fall back to the plain blocked error if we can't write the request.
		return &pledge.BlockedError{
			OperationID: opID,
			Required:    tier,
			MaxAllowed:  pf.MaxTier,
			Profile:     pf.Profile,
		}
	}

	return fmt.Errorf(
		"atlas pledge [%s]: operation %q requires %q but session is restricted to %q\n"+
			"Run in another terminal to approve, then retry:\n\n  atlas pledge allow %s",
		pf.Profile, opID, tier, pf.MaxTier, token,
	)
}

// widenInteractiveExec prompts the user inline to widen the session pledge.
// Called only when os.Stdin is a TTY.
func widenInteractiveExec(in io.Reader, errW io.Writer, key pledge.SessionKey, current *pledge.PledgeFile, required shared_api.PermissionTier, opID string) error {
	targetProfile := pledge.ProfileReadWrite
	if required == shared_api.PermissionAdmin {
		targetProfile = pledge.ProfileAdmin
	}

	fmt.Fprintf(errW, "\natlas pledge [%s]: %q requires %q permission.\n", current.Profile, opID, required)
	if targetProfile == pledge.ProfileAdmin {
		fmt.Fprintln(errW, "WARNING: admin allows all operations including destructive org-level actions.")
	}
	fmt.Fprintf(errW, "Widen this session to %s? [y/N] ", targetProfile)

	var answer string
	if _, err := fmt.Fscan(in, &answer); err != nil || answer != "y" && answer != "Y" {
		return &pledge.BlockedError{
			OperationID: opID,
			Required:    required,
			MaxAllowed:  current.MaxTier,
			Profile:     current.Profile,
		}
	}

	pf, err := pledge.NewPledgeFile(targetProfile, nil)
	if err != nil {
		return err
	}
	return pledge.Widen(key, pf)
}

// Log the response if the logger is set to debug
// Copied behavior and format used in the SDK: https://github.com/mongodb/atlas-sdk-go/blob/b3fee40e236a8ff2a1f1c160b6984a242136dbe6/admin/client.go#L335
func (e *Executor) logResponse(httpResponse *http.Response) {
	if !e.logger.IsDebugLevel() {
		return
	}

	dump, err := httputil.DumpResponse(httpResponse, true)
	if err != nil {
		return
	}

	_, _ = e.logger.Debugf("\n%s\n", string(dump))
}
