// Copyright 2022 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/log"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prerun"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prompt"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/version"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312027/admin"
	authcode "go.mongodb.org/atlas-sdk/v20250312027/auth/code"
	"go.mongodb.org/atlas/auth"
)

//go:generate go tool go.uber.org/mock/mockgen -typed -destination=login_mock_test.go -package=auth . SetSaver,LoginConfig,TrackAsker,AuthorizationCodeFlow

type SetSaver interface {
	SetAuthType(config.AuthMechanism)
	SetAccessToken(string)
	SetRefreshToken(string)
	SetPublicAPIKey(string)
	SetPrivateAPIKey(string)
	SetClientID(string)
	SetClientSecret(string)
	SetOrgID(string)
	SetProjectID(string)
	SetOpsManagerURL(string)
	SetService(string)
	Save() error
	SetGlobal(string, any)
}

type LoginConfig interface {
	SetSaver
	transport.ProfileProvider
	AccessTokenSubject() (string, error)
	OrgID() string
	ProjectID() string
}

type TrackAsker interface {
	TrackAsk([]*survey.Question, any, ...survey.AskOpt) error
	TrackAskOne(survey.Prompt, any, ...survey.AskOpt) error
}

const (
	userAccountAuth        = "UserAccount"
	userAccountLegacyAuth  = "UserAccountLegacy"
	atlasName              = "atlas"
	authCodeCallbackPath   = "/atlas-cli/callback"
	authCodeLoginNotice    = "User Account login now uses Authorization Code + PKCE. To use the deprecated device-code flow, run `atlas auth login --authType=UserAccountLegacy`."
	userAccountDeprecation = "Warning: the device-code User Account flow is deprecated and will be removed in a future release."
)

var (
	ErrProjectIDNotFound = errors.New("project is inaccessible. You either don't have access to this project or the project doesn't exist")
	ErrOrgIDNotFound     = errors.New("organization is inaccessible. You don't have access to this organization or the organization doesn't exist")
	authTypeOptions      = []string{userAccountAuth, userAccountLegacyAuth, prompt.ServiceAccountAuth, prompt.APIKeysAuth}
	outputFormatOptions  = []string{"plaintext", "json"}
	authTypeDescription  = map[string]string{
		userAccountAuth:           "(best for getting started)",
		userAccountLegacyAuth:     "(deprecated device-code flow)",
		prompt.ServiceAccountAuth: "(best for automation)",
		prompt.APIKeysAuth:        "(for existing automations)",
	}
)

// AuthorizationCodeFlow is the Authorization Code + PKCE surface the login flow needs.
// The default implementation wraps *authcode.Config.
type AuthorizationCodeFlow interface {
	ResolveAuthServerMetadata(ctx context.Context, cached *authcode.AuthServerMetadata) (*authcode.AuthServerMetadata, error)
	AuthorizationURL(metadata *authcode.AuthServerMetadata, redirectURI, state, codeVerifier string) (string, error)
	ExchangeCode(ctx context.Context, metadata *authcode.AuthServerMetadata, authorizationCode, redirectURI, codeVerifier string) (*authcode.Token, error)
}

type authCodeConfig struct {
	cfg *authcode.Config
}

func (a authCodeConfig) ResolveAuthServerMetadata(ctx context.Context, cached *authcode.AuthServerMetadata) (*authcode.AuthServerMetadata, error) {
	return transport.ResolveAuthServerMetadata(ctx, a.cfg, cached, false)
}

// loginAuthServer ignores the stored issuer so --gov and MONGODB_ATLAS_AUTH_SERVER_URL apply on re-login.
type loginAuthServer struct {
	service string
}

func (s loginAuthServer) Service() string        { return s.service }
func (loginAuthServer) AuthServerIssuer() string { return "" }

func (a authCodeConfig) AuthorizationURL(metadata *authcode.AuthServerMetadata, redirectURI, state, codeVerifier string) (string, error) {
	return a.cfg.AuthorizationURL(metadata, redirectURI, state, codeVerifier)
}

func (a authCodeConfig) ExchangeCode(ctx context.Context, metadata *authcode.AuthServerMetadata, authorizationCode, redirectURI, codeVerifier string) (*authcode.Token, error) {
	token, _, err := a.cfg.ExchangeCode(ctx, metadata, authorizationCode, redirectURI, codeVerifier)
	return token, err
}

type LoginOpts struct {
	cli.DefaultSetterOpts
	cli.RefresherOpts
	AccessToken      string
	RefreshToken     string
	ClientID         string
	ClientSecret     string
	PublicAPIKey     string
	PrivateAPIKey    string
	IsGov            bool
	NoBrowser        bool
	orgIDFlag        string
	projectIDFlag    string
	authType         string
	authTypeFromFlag bool
	force            bool
	SkipConfig       bool
	config           LoginConfig
	Asker            TrackAsker
	authCode         AuthorizationCodeFlow
}

func (opts *LoginOpts) validateOutput() error {
	if !slices.Contains(outputFormatOptions, opts.Output) {
		return fmt.Errorf("invalid --%s %q: must be one of %s", flag.Output, opts.Output, strings.Join(outputFormatOptions, ", "))
	}
	return nil
}

func (opts *LoginOpts) validateOrgProjectFlags() error {
	if opts.orgIDFlag != "" {
		if err := validate.ObjectID(opts.orgIDFlag); err != nil {
			return fmt.Errorf("invalid --%s: %w", flag.OrgID, err)
		}
	}
	if opts.projectIDFlag != "" {
		if err := validate.ObjectID(opts.projectIDFlag); err != nil {
			return fmt.Errorf("invalid --%s: %w", flag.ProjectID, err)
		}
	}
	return nil
}

// validateLoginFlags ensures that once any programmatic login flag is used, every flag
// needed to complete that flow without a prompt is also set, instead of silently falling
// back to a prompt for whatever was left out.
func (opts *LoginOpts) validateLoginFlags() error {
	if err := opts.validateOrgProjectFlags(); err != nil {
		return err
	}

	anyOther := opts.ClientID != "" || opts.ClientSecret != "" ||
		opts.PublicAPIKey != "" || opts.PrivateAPIKey != "" || opts.Output != ""

	if opts.authType == "" {
		if anyOther {
			return fmt.Errorf("--%s is required when using --%s, --%s, --%s, --%s, or --%s",
				flag.AuthType, flag.ClientID, flag.ClientSecret, flag.PublicAPIKey, flag.PrivateAPIKey, flag.Output)
		}
		return nil
	}
	if opts.Output == "" {
		return fmt.Errorf("--%s is required when --%s is set", flag.Output, flag.AuthType)
	}
	if err := opts.validateOutput(); err != nil {
		return err
	}

	return opts.validateCredentials()
}

func (opts *LoginOpts) validateCredentials() error {
	switch opts.authType {
	case prompt.ServiceAccountAuth:
		if opts.ClientID == "" || opts.ClientSecret == "" {
			return fmt.Errorf("--%s and --%s are required when --%s is %s", flag.ClientID, flag.ClientSecret, flag.AuthType, prompt.ServiceAccountAuth)
		}
	case prompt.APIKeysAuth:
		if opts.PublicAPIKey == "" || opts.PrivateAPIKey == "" {
			return fmt.Errorf("--%s and --%s are required when --%s is %s", flag.PublicAPIKey, flag.PrivateAPIKey, flag.AuthType, prompt.APIKeysAuth)
		}
	}
	return nil
}

func (opts *LoginOpts) promptAuthType() error {
	if opts.authType != "" {
		opts.authTypeFromFlag = true
		if !slices.Contains(authTypeOptions, opts.authType) {
			return fmt.Errorf("invalid --%s %q: must be one of %s", flag.AuthType, opts.authType, strings.Join(authTypeOptions, ", "))
		}
		return nil
	}
	if opts.force {
		opts.authType = userAccountAuth
		return nil
	}
	authTypePrompt := &survey.Select{
		Message: "Select authentication type:",
		Options: authTypeOptions,
		Default: userAccountAuth,
		Description: func(value string, _ int) string {
			return authTypeDescription[value]
		},
	}
	return opts.Asker.TrackAskOne(authTypePrompt, &opts.authType)
}

func (opts *LoginOpts) promptOutput() error {
	if opts.Output != "" {
		return opts.validateOutput()
	}
	return opts.Asker.TrackAsk(opts.DefaultQuestions(), opts)
}

func (opts *LoginOpts) setUserAccountCredentials(ctx context.Context) error {
	switch opts.authType {
	case userAccountAuth:
		if err := opts.authCodeFlow(ctx); err != nil {
			return err
		}
	case userAccountLegacyAuth:
		_, _ = fmt.Fprintln(opts.OutWriter, userAccountDeprecation)
		if err := opts.oauthFlow(ctx); err != nil {
			return err
		}
	}
	// Sync config with OAuth tokens
	if err := opts.SyncWithOAuthAccessProfile(opts.config)(); err != nil {
		return err
	}
	s, err := opts.config.AccessTokenSubject()
	if err != nil {
		return err
	}

	if err := opts.checkProfile(ctx); err != nil {
		return err
	}

	if err := opts.config.Save(); err != nil {
		return err
	}
	_, _ = fmt.Fprintf(opts.OutWriter, "Successfully logged in as %s.\n", s)

	return nil
}

// credentialsProvided reports whether the credential pair for opts.authType was fully
// supplied. validateLoginFlags already rejects a half-supplied pair before this is called,
// so there is nothing left to validate here, only to check.
func (opts *LoginOpts) credentialsProvided() bool {
	switch opts.authType {
	case prompt.ServiceAccountAuth:
		return opts.ClientID != "" && opts.ClientSecret != ""
	case prompt.APIKeysAuth:
		return opts.PublicAPIKey != "" && opts.PrivateAPIKey != ""
	default:
		return false
	}
}

func (opts *LoginOpts) setProgrammaticCredentials() error {
	if opts.credentialsProvided() {
		opts.setUpAccess()
		return nil
	}

	_, _ = fmt.Fprintf(opts.OutWriter, `You are configuring a profile for %s.

All values are optional and you can use environment variables (MONGODB_ATLAS_*) instead.
Note: environment variables take precedence over the values you enter here. To learn more, see %s.

Enter [?] on any option to get help.

`, atlasName, commonerrors.EnvVarsDocsURL)

	q := prompt.AccessQuestions(opts.authType)
	if err := opts.Asker.TrackAsk(q, opts); err != nil {
		return err
	}

	opts.setUpAccess()

	return nil
}

func (opts *LoginOpts) setUpCredentials(ctx context.Context) error {
	switch opts.authType {
	case userAccountAuth, userAccountLegacyAuth:
		return opts.setUserAccountCredentials(ctx)
	case prompt.ServiceAccountAuth, prompt.APIKeysAuth:
		return opts.setProgrammaticCredentials()
	default:
		return errors.New("no authentication type selected")
	}
}

func (opts *LoginOpts) setUpAccess() {
	// Set service
	switch {
	case opts.IsGov:
		opts.Service = config.CloudGovService
	default:
		opts.Service = config.CloudService
	}
	opts.config.SetService(opts.Service)

	// Set authentication credentials
	switch opts.authType {
	case prompt.ServiceAccountAuth:
		if opts.ClientID != "" {
			opts.config.SetClientID(opts.ClientID)
		}
		if opts.ClientSecret != "" {
			opts.config.SetClientSecret(opts.ClientSecret)
		}
	case prompt.APIKeysAuth:
		if opts.PublicAPIKey != "" {
			opts.config.SetPublicAPIKey(opts.PublicAPIKey)
		}
		if opts.PrivateAPIKey != "" {
			opts.config.SetPrivateAPIKey(opts.PrivateAPIKey)
		}
	}
}

// SyncWithOAuthAccessProfile returns a function that is synchronizing the oauth settings
// from a login config profile with the provided command opts.
func (opts *LoginOpts) SyncWithOAuthAccessProfile(c LoginConfig) func() error {
	return func() error {
		opts.config = c

		switch {
		case opts.IsGov:
			opts.Service = config.CloudGovService
		default:
			opts.Service = config.CloudService
		}
		opts.config.SetService(opts.Service)

		if opts.AccessToken != "" {
			opts.config.SetAccessToken(opts.AccessToken)
		}
		if opts.RefreshToken != "" {
			opts.config.SetRefreshToken(opts.RefreshToken)
		}
		if config.ClientID() != "" {
			opts.config.SetClientID(config.ClientID())
		}

		// sync OpsManagerURL from command opts (higher priority)
		// and OpsManagerURL from default profile
		if opts.OpsManagerURL != "" {
			opts.config.SetOpsManagerURL(opts.OpsManagerURL)
		}
		if config.OpsManagerURL() != "" {
			opts.OpsManagerURL = config.OpsManagerURL()
		}

		return nil
	}
}

func (opts *LoginOpts) LoginRun(ctx context.Context) error {
	if err := opts.validateLoginFlags(); err != nil {
		return err
	}

	if err := opts.promptAuthType(); err != nil {
		return fmt.Errorf("failed to select authentication type: %w", err)
	}

	switch opts.authType {
	case userAccountAuth:
		opts.config.SetAuthType(config.UserAccount)
	case userAccountLegacyAuth:
		opts.config.SetAuthType(config.UserAccountLegacy)
	case prompt.ServiceAccountAuth:
		opts.config.SetAuthType(config.ServiceAccount)
	case prompt.APIKeysAuth:
		opts.config.SetAuthType(config.APIKeys)
	default:
		return errors.New("no authentication type selected")
	}

	if err := opts.setUpCredentials(ctx); err != nil {
		return err
	}

	if opts.SkipConfig {
		return nil
	}

	if err := opts.setUpProfile(ctx); err != nil {
		return err
	}

	if config.Name() != config.DefaultProfile {
		_, _ = fmt.Fprintf(opts.OutWriter, "To use this profile, you must set the flag [-%s %s] for every command.\n", flag.ProfileShort, config.Name())
	}

	return nil
}

func (opts *LoginOpts) checkProfile(ctx context.Context) error {
	if err := opts.InitStore(ctx); err != nil {
		return err
	}
	if opts.config.OrgID() != "" && !opts.OrgExists(opts.config.OrgID()) {
		opts.config.SetOrgID("")
	}

	if opts.config.ProjectID() != "" && !opts.ProjectExists(opts.config.ProjectID()) {
		opts.config.SetProjectID("")
	}
	return nil
}

func (opts *LoginOpts) setUpProfile(ctx context.Context) error {
	if err := opts.InitStore(ctx); err != nil {
		return err
	}
	// Initialize the text to be displayed if users are asked to select orgs or projects
	opts.OnMultipleOrgsOrProjects = func() {
		if !opts.AskedOrgsOrProjects {
			_, _ = fmt.Fprintln(opts.OutWriter, `Select one default organization and one default project.`)
		}
	}

	// The org list doubles as the credential check for programmatic login.
	soleOrgID := ""
	if opts.authTypeFromFlag {
		var err error
		soleOrgID, err = opts.OnlyOrgID(ctx)
		if err != nil {
			return err
		}
	}

	if err := opts.resolveOrg(ctx, soleOrgID); err != nil {
		return err
	}
	if err := opts.resolveProject(); err != nil {
		return err
	}
	// After resolveProject, which may derive the org from --projectId.
	opts.SetUpOrg()
	opts.SetUpProject()

	if err := opts.promptOutput(); err != nil {
		return err
	}
	opts.SetUpOutput()

	if err := opts.config.Save(); err != nil {
		return err
	}

	if opts.authTypeFromFlag {
		if opts.config.OrgID() == "" {
			_, _ = fmt.Fprintf(opts.OutWriter, "No default organization set; pass --%s on commands that require one.\n", flag.OrgID)
		}
		if opts.config.ProjectID() == "" {
			_, _ = fmt.Fprintf(opts.OutWriter, "No default project set; pass --%s on commands that require one.\n", flag.ProjectID)
		}
	}

	return opts.validateOrgAndProject()
}

func (opts *LoginOpts) resolveOrg(ctx context.Context, soleOrgID string) error {
	if opts.orgIDFlag != "" {
		if err := opts.verifyOrgAccess(opts.orgIDFlag); err != nil {
			return err
		}
		opts.OrgID = opts.orgIDFlag
		return nil
	}

	if cfgOrgID := opts.config.OrgID(); cfgOrgID != "" && opts.OrgExists(cfgOrgID) {
		opts.OrgID = cfgOrgID
		return nil
	}

	if opts.authTypeFromFlag {
		opts.config.SetOrgID("")
		opts.OrgID = soleOrgID
		return nil
	}

	return opts.AskOrg(ctx)
}

func (opts *LoginOpts) resolveProject() error {
	if opts.projectIDFlag != "" {
		project, err := opts.verifyProjectAccess(opts.projectIDFlag)
		if err != nil {
			return err
		}
		if opts.orgIDFlag != "" {
			if project.GetOrgId() != opts.orgIDFlag {
				return fmt.Errorf("--%s %s belongs to organization %s, not --%s %s",
					flag.ProjectID, opts.projectIDFlag, project.GetOrgId(), flag.OrgID, opts.orgIDFlag)
			}
		} else {
			opts.OrgID = project.GetOrgId()
		}
		opts.ProjectID = opts.projectIDFlag
		return nil
	}

	if cfgProjectID := opts.config.ProjectID(); cfgProjectID != "" {
		project, err := opts.verifyProjectAccess(cfgProjectID)
		switch {
		case err != nil:
			// Inaccessible or server error; fall through and re-resolve.
		case opts.OrgID != "" && project.GetOrgId() != opts.OrgID:
			if opts.orgIDFlag != "" {
				return fmt.Errorf("--%s %s belongs to organization %s, not --%s %s",
					flag.ProjectID, cfgProjectID, project.GetOrgId(), flag.OrgID, opts.orgIDFlag)
			}
			// Stale project from another org; fall through and re-resolve.
		default:
			opts.ProjectID = cfgProjectID
			return nil
		}
	}

	if opts.authTypeFromFlag {
		opts.config.SetProjectID("")
		var err error
		opts.ProjectID, err = opts.OnlyProjectID()
		return err
	}

	return opts.AskProject()
}

// Credentials are verified before this runs, so a 401 means no access to the resource.
func isAccessError(err error) bool {
	apiErr, ok := atlasv2.AsError(err)
	if !ok {
		return false
	}
	switch apiErr.GetError() {
	case http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound:
		return true
	}
	return false
}

func (opts *LoginOpts) verifyOrgAccess(id string) error {
	o, err := opts.Store.Organization(id)
	if err != nil && !isAccessError(err) {
		return fmt.Errorf("--%s %s: %w", flag.OrgID, id, err)
	}
	if err != nil || o == nil {
		return fmt.Errorf("--%s %s: %w", flag.OrgID, id, ErrOrgIDNotFound)
	}
	return nil
}

func (opts *LoginOpts) verifyProjectAccess(id string) (*atlasv2.Group, error) {
	p, err := opts.Store.Project(id)
	if err != nil && !isAccessError(err) {
		return nil, fmt.Errorf("--%s %s: %w", flag.ProjectID, id, err)
	}
	if err != nil || p == nil {
		return nil, fmt.Errorf("--%s %s: %w", flag.ProjectID, id, ErrProjectIDNotFound)
	}
	return p, nil
}

func (opts *LoginOpts) validateOrgAndProject() error {
	// Only make references to profile if user was asked about org or projects
	if opts.AskedOrgsOrProjects && opts.ProjectID != "" && opts.OrgID != "" {
		if !opts.ProjectExists(opts.config.ProjectID()) {
			return ErrProjectIDNotFound
		}

		if !opts.OrgExists(opts.config.OrgID()) {
			return ErrOrgIDNotFound
		}

		_, _ = fmt.Fprint(opts.OutWriter, `
You have successfully configured your profile.
You can use [atlas config set] to change your profile settings later.
`)
	}
	return nil
}

func (opts *LoginOpts) printAuthInstructions(code *auth.DeviceCode) {
	codeDuration := time.Duration(code.ExpiresIn) * time.Second
	_, _ = fmt.Fprintf(opts.OutWriter, `
To verify your account, copy your one-time verification code:
`)

	userCode := fmt.Sprintf("%s-%s", code.UserCode[0:len(code.UserCode)/2], code.UserCode[len(code.UserCode)/2:])
	_, _ = fmt.Fprintln(opts.OutWriter, userCode)

	_, _ = fmt.Fprintf(opts.OutWriter, `
Paste the code in the browser when prompted to activate your Atlas CLI. Your code will expire after %.0f minutes.

To continue, go to `,
		codeDuration.Minutes(),
	)
	_, _ = fmt.Fprintln(opts.OutWriter, code.VerificationURI)
}

func (opts *LoginOpts) browserConfirmationRequired() bool {
	return !opts.force && !opts.authTypeFromFlag
}

func (opts *LoginOpts) handleBrowser(uri string) {
	if opts.NoBrowser {
		return
	}

	if opts.browserConfirmationRequired() {
		_, _ = fmt.Fprintf(opts.OutWriter, "\nPress Enter to open the browser and complete authentication...")
		_, _ = fmt.Scanln()
	}
	if errBrowser := browser.OpenURL(uri); errBrowser != nil {
		_, _ = log.Warningln("There was an issue opening your browser")
	}
}

func (opts *LoginOpts) oauthFlow(ctx context.Context) error {
	askedToOpenBrowser := false
	for {
		code, _, err := opts.RequestCode(ctx)
		if err != nil {
			return err
		}

		opts.printAuthInstructions(code)
		if !askedToOpenBrowser {
			opts.handleBrowser(code.VerificationURI)
			askedToOpenBrowser = true
		}

		accessToken, _, err := opts.PollToken(ctx, code)
		if retry, errRetry := opts.shouldRetryAuthenticate(err, newRegenerationPrompt()); errRetry != nil {
			return errRetry
		} else if retry {
			continue
		}
		if err != nil {
			return err
		}

		opts.AccessToken = accessToken.AccessToken
		opts.RefreshToken = accessToken.RefreshToken
		return nil
	}
}

// authCodeFlow runs the Authorization Code + PKCE login through the SDK loopback
// callback server and persists the token via the shared core write path.
func (opts *LoginOpts) authCodeFlow(ctx context.Context) error {
	flow := opts.authCode
	if flow == nil {
		cfg, err := transport.FlowForAuthServer(loginAuthServer{service: opts.config.Service()}, &http.Client{Transport: transport.Default()}, version.Version)
		if err != nil {
			return err
		}
		flow = authCodeConfig{cfg}
		opts.authCode = flow
	}
	metadata, err := flow.ResolveAuthServerMetadata(ctx, opts.config.AuthServerMetadata())
	if err != nil {
		return err
	}

	state, err := authcode.GenerateState()
	if err != nil {
		return err
	}
	verifier := authcode.GenerateVerifier()

	server, err := authcode.StartCallbackServer(state, authCodeCallbackPath, metadata)
	if err != nil {
		return err
	}
	defer func() { _ = server.Close() }()

	authURL, err := flow.AuthorizationURL(metadata, server.RedirectURI(), state, verifier)
	if err != nil {
		return err
	}

	_, _ = fmt.Fprintf(opts.OutWriter, "%s\n\n", authCodeLoginNotice)
	if opts.NoBrowser {
		_, _ = fmt.Fprintf(opts.OutWriter, "To continue, go to %s\n", authURL)
	} else {
		opts.handleBrowser(authURL)
	}

	authorizationCode, err := server.WaitForCallback(ctx)
	if err != nil {
		return err
	}

	token, err := flow.ExchangeCode(ctx, metadata, authorizationCode, server.RedirectURI(), verifier)
	if err != nil {
		return err
	}
	opts.AccessToken = token.AccessToken
	opts.RefreshToken = token.RefreshToken

	return transport.SaveUserAccountToken(opts.config, token, metadata)
}

func (opts *LoginOpts) shouldRetryAuthenticate(err error, p survey.Prompt) (retry bool, errSurvey error) {
	if err == nil || !auth.IsTimeoutErr(err) {
		return false, nil
	}
	err = opts.Asker.TrackAskOne(p, &retry)
	return retry, err
}

func newRegenerationPrompt() survey.Prompt {
	return &survey.Confirm{
		Message: "Your one-time verification code is expired. Would you like to generate a new one?",
		Default: true,
	}
}

func (opts *LoginOpts) LoginPreRun(ctx context.Context) func() error {
	return func() error {
		// ignore expired tokens since logging in
		if err := opts.RefreshAccessToken(ctx); err != nil {
			// clean up any expired or invalid tokens
			opts.config.SetAccessToken(
				"")

			if !commonerrors.IsInvalidRefreshToken(err) {
				return err
			}
		}

		return nil
	}
}

func LoginBuilder() *cobra.Command {
	opts := &LoginOpts{
		Asker: &telemetry.Ask{},
	}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with MongoDB Atlas.",
		Long: `This command allows you to authenticate with MongoDB Atlas using User Account, Service Account, or API Key authentication methods.

User Account login uses the Authorization Code + PKCE flow with a local callback server. To use the deprecated device-code flow instead, run ` + "`atlas auth login --authType=UserAccountLegacy`" + `.

Note: If you have credentials set in environment variables, they take precedence over the values you provide during authentication. To learn more, see ` + commonerrors.EnvVarsDocsURL + `.

To log in non-interactively, set --authType, --output, and the credential flags that type requires (--clientId and --clientSecret for ServiceAccount, or --publicApiKey and --privateApiKey for APIKeys). Using any one of these flags without the others needed to complete the flow returns an error instead of prompting.

Without --orgId or --projectId, a default is saved only if the credentials can access exactly one organization or project.`,
		Example: `  # Log in to your MongoDB Atlas account in interactive mode:
  atlas auth login

  # Log in non-interactively with a Service Account:
  atlas auth login --authType ServiceAccount --clientId <clientId> --clientSecret <clientSecret> --output plaintext

  # Log in non-interactively with a Service Account and store the default organization and project:
  atlas auth login --authType ServiceAccount --clientId <clientId> --clientSecret <clientSecret> --output plaintext --orgId <orgId> --projectId <projectId>

  # Log in non-interactively with API Keys:
  atlas auth login --authType APIKeys --publicApiKey <publicApiKey> --privateApiKey <privateApiKey> --output plaintext
`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			opts.OutWriter = cmd.OutOrStdout()
			defaultProfile := config.Default()
			return prerun.ExecuteE(
				opts.SyncWithOAuthAccessProfile(defaultProfile),
				opts.InitFlow(defaultProfile),
				opts.LoginPreRun(cmd.Context()),
				validate.NoAccessToken,
			)
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return opts.LoginRun(cmd.Context())
		},
		Args: require.NoArgs,
	}

	cmd.Flags().BoolVar(&opts.IsGov, "gov", false, "Log in to Atlas for Government.")
	cmd.Flags().BoolVar(&opts.NoBrowser, "noBrowser", false, "Don't automatically open a browser session.")
	cmd.Flags().BoolVar(&opts.SkipConfig, "skipConfig", false, "Skip profile configuration.")
	_ = cmd.Flags().MarkDeprecated("skipConfig", "if you configured a profile, the command skips the config step by default.")
	cmd.Flags().BoolVar(&opts.force, flag.Force, false, usage.Force)
	_ = cmd.Flags().MarkHidden(flag.Force)
	cmd.Flags().StringVar(&opts.authType, flag.AuthType, "", usage.LoginAuthType)
	cmd.Flags().StringVar(&opts.ClientID, flag.ClientID, "", usage.LoginClientID)
	cmd.Flags().StringVar(&opts.ClientSecret, flag.ClientSecret, "", usage.LoginClientSecret)
	cmd.Flags().StringVar(&opts.PublicAPIKey, flag.PublicAPIKey, "", usage.LoginPublicAPIKey)
	cmd.Flags().StringVar(&opts.PrivateAPIKey, flag.PrivateAPIKey, "", usage.LoginPrivateAPIKey)
	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.LoginOutput)
	cmd.Flags().StringVar(&opts.orgIDFlag, flag.OrgID, "", usage.LoginOrgID)
	cmd.Flags().StringVar(&opts.projectIDFlag, flag.ProjectID, "", usage.LoginProjectID)
	return cmd
}
