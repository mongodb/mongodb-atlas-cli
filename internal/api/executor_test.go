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
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312027/auth/code"
	"go.mongodb.org/atlas/auth"
	"go.uber.org/mock/gomock"
)

type testProfileProvider struct {
	authType config.AuthMechanism
	tokenErr error
}

func (p *testProfileProvider) AuthType() config.AuthMechanism { return p.authType }
func (*testProfileProvider) PublicAPIKey() string             { panic("PublicAPIKey should not be called") }
func (*testProfileProvider) PrivateAPIKey() string            { panic("PrivateAPIKey should not be called") }
func (p *testProfileProvider) Token() (*auth.Token, error)    { return nil, p.tokenErr }
func (*testProfileProvider) ServiceAccountToken() (*auth.Token, error) {
	panic("ServiceAccountToken should not be called")
}
func (*testProfileProvider) SetAccessToken(string)  { panic("SetAccessToken should not be called") }
func (*testProfileProvider) SetRefreshToken(string) { panic("SetRefreshToken should not be called") }
func (*testProfileProvider) Save() error            { panic("Save should not be called") }
func (*testProfileProvider) ClientID() string       { panic("ClientID should not be called") }
func (*testProfileProvider) ClientSecret() string   { panic("ClientSecret should not be called") }
func (*testProfileProvider) OpsManagerURL() string  { panic("OpsManagerURL should not be called") }
func (*testProfileProvider) AccessToken() string    { panic("AccessToken should not be called") }
func (*testProfileProvider) RefreshToken() string   { panic("RefreshToken should not be called") }
func (*testProfileProvider) SetTokenExpiry(string)  { panic("SetTokenExpiry should not be called") }
func (*testProfileProvider) Service() string        { panic("Service should not be called") }
func (*testProfileProvider) AuthServerIssuer() string {
	panic("AuthServerIssuer should not be called")
}
func (*testProfileProvider) SetAuthServerIssuer(string) {
	panic("SetAuthServerIssuer should not be called")
}
func (*testProfileProvider) FreshRefreshToken() (string, error) {
	panic("FreshRefreshToken should not be called")
}
func (*testProfileProvider) AuthServerMetadata() *code.AuthServerMetadata {
	panic("AuthServerMetadata should not be called")
}
func (*testProfileProvider) SetAuthServerMetadata(*code.AuthServerMetadata) {
	panic("SetAuthServerMetadata should not be called")
}

func TestExecutorHappyPathNoLogging(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)

	configProvider := NewMockConfigProvider(ctrl)
	configProvider.EXPECT().GetBaseURL().Return("https://cloud.mongodb.com", nil).AnyTimes()

	commandConverter, err := NewDefaultCommandConverter(configProvider)
	require.NoError(t, err)
	require.NotNil(t, commandConverter)

	httpClient := NewMockDoer(ctrl)
	httpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"success": true}`)),
	}, nil).AnyTimes()

	logger := NewMockLogger(ctrl)
	logger.EXPECT().IsDebugLevel().Return(false).AnyTimes()

	executor, err := NewExecutor(commandConverter, httpClient, NewFormatter(), logger)
	require.NoError(t, err)
	require.NotNil(t, executor)

	// Act
	commandRequest := CommandRequest{
		Command: api.Command{
			OperationID: "testOperation",
			RequestParameters: api.RequestParameters{
				URL: "/test/url",
			},
			Versions: []api.CommandVersion{{
				Version:              api.NewStableVersion(1991, 5, 17),
				RequestContentType:   "json",
				ResponseContentTypes: []string{"json"},
			}},
		},
		ContentType: "json",
		Format:      "json",
		Parameters:  nil,
		Version:     api.NewStableVersion(1991, 5, 17),
	}
	response, err := executor.ExecuteCommand(t.Context(), commandRequest)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestExecutorHappyPathDebugLogging(t *testing.T) {
	// Setup
	ctrl := gomock.NewController(t)

	configProvider := NewMockConfigProvider(ctrl)
	configProvider.EXPECT().GetBaseURL().Return("https://cloud.mongodb.com", nil).AnyTimes()

	commandConverter, err := NewDefaultCommandConverter(configProvider)
	require.NoError(t, err)
	require.NotNil(t, commandConverter)

	httpClient := NewMockDoer(ctrl)
	httpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"success": true}`)),
	}, nil).AnyTimes()

	logger := NewMockLogger(ctrl)
	logger.EXPECT().IsDebugLevel().Return(true).AnyTimes()

	gomock.InOrder(
		logger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Do(func(_ string, args ...any) {
			require.Len(t, args, 1)
			require.IsType(t, "", args[0])
			require.Equal(t, "GET /test/url HTTP/1.1\r\nHost: cloud.mongodb.com\r\nUser-Agent: Go-http-client/1.1\r\nAccept: application/vnd.atlas.1991-05-17+json\r\nContent-Type: application/vnd.atlas.1991-05-17+json\r\nAccept-Encoding: gzip\r\n\r\n", args[0].(string))
		}).Times(1),
		logger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Do(func(_ string, args ...any) {
			require.Len(t, args, 1)
			require.IsType(t, "", args[0])
			require.Equal(t, "HTTP/0.0 200 OK\r\n\r\n{\"success\": true}", args[0].(string))
		}).Times(1),
	)

	executor, err := NewExecutor(commandConverter, httpClient, NewFormatter(), logger)
	require.NoError(t, err)
	require.NotNil(t, executor)

	// Act
	commandRequest := CommandRequest{
		Command: api.Command{
			OperationID: "testOperation",
			RequestParameters: api.RequestParameters{
				URL: "/test/url",
			},
			Versions: []api.CommandVersion{{
				Version:              api.NewStableVersion(1991, 5, 17),
				RequestContentType:   "json",
				ResponseContentTypes: []string{"json"},
			}},
		},
		ContentType: "json",
		Format:      "json",
		Parameters:  nil,
		Version:     api.NewStableVersion(1991, 5, 17),
	}
	response, err := executor.ExecuteCommand(t.Context(), commandRequest)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, response)
}

func TestNewDefaultHTTPClient_UnauthenticatedAPICommandSkipsCredentials(t *testing.T) {
	command := api.Command{
		Unauthenticated: true,
	}

	client, err := newDefaultHTTPClient(&testProfileProvider{}, command)

	require.NoError(t, err)
	require.Same(t, transport.Default(), client.Transport)
}

func TestNewDefaultHTTPClient_AuthenticatedAPICommandUsesCredentials(t *testing.T) {
	expectedErr := errors.New("missing token")

	_, err := newDefaultHTTPClient(&testProfileProvider{
		authType: config.UserAccount,
		tokenErr: expectedErr,
	}, api.Command{})

	require.ErrorIs(t, err, expectedErr)
}

func TestNewDefaultHTTPClient_UserAccountLegacyRefreshUsesInferredClientID(t *testing.T) {
	var gotClientID string
	var gotRefreshToken string
	var gotAuthorization string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/private/unauth/account/device/token":
			if err := r.ParseForm(); err != nil {
				t.Errorf("ParseForm() error = %v", err)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			gotClientID = r.Form.Get("client_id")
			gotRefreshToken = r.Form.Get("refresh_token")
			w.Header().Set("Content-Type", "application/json")
			_, _ = io.WriteString(w, `{"access_token":"refreshed-access","refresh_token":"refreshed-refresh","token_type":"Bearer","expires_in":3600}`)
		case "/resource":
			gotAuthorization = r.Header.Get("Authorization")
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected request path %q", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer server.Close()

	profile := &legacyTestProfileProvider{
		testProfileProvider: testProfileProvider{authType: config.UserAccountLegacy},
		token: &auth.Token{
			AccessToken:  "expired-access",
			RefreshToken: "old-refresh",
			Expiry:       time.Now().Add(-time.Hour),
		},
		service:       config.CloudService,
		opsManagerURL: "https://cloud-dev.mongodb.com/",
		accountURL:    server.URL + "/",
	}

	client, err := newDefaultHTTPClient(profile, api.Command{})
	require.NoError(t, err)
	request, err := http.NewRequestWithContext(t.Context(), http.MethodGet, server.URL+"/resource", nil)
	require.NoError(t, err)
	response, err := client.Do(request) //nolint:gosec // G704: URL comes from httptest.NewServer, not user input
	require.NoError(t, err)
	assert.NoError(t, response.Body.Close())

	assert.Equal(t, "0oadn4hoajpzxeSEy357", gotClientID)
	assert.Equal(t, "old-refresh", gotRefreshToken)
	assert.Equal(t, "Bearer refreshed-access", gotAuthorization)
	assert.Equal(t, "refreshed-access", profile.savedAccessToken)
	assert.Equal(t, "refreshed-refresh", profile.savedRefreshToken)
	assert.Equal(t, 1, profile.saveCalls)
}

type legacyTestProfileProvider struct {
	testProfileProvider
	token             *auth.Token
	service           string
	opsManagerURL     string
	clientID          string
	accountURL        string
	savedAccessToken  string
	savedRefreshToken string
	saveCalls         int
}

func (p *legacyTestProfileProvider) Token() (*auth.Token, error) { return p.token, nil }
func (p *legacyTestProfileProvider) Service() string             { return p.service }
func (p *legacyTestProfileProvider) OpsManagerURL() string       { return p.opsManagerURL }
func (p *legacyTestProfileProvider) ClientID() string            { return p.clientID }
func (p *legacyTestProfileProvider) AccountURL() string          { return p.accountURL }
func (p *legacyTestProfileProvider) SetAccessToken(token string) { p.savedAccessToken = token }
func (p *legacyTestProfileProvider) SetRefreshToken(token string) {
	p.savedRefreshToken = token
}
func (p *legacyTestProfileProvider) Save() error {
	p.saveCalls++
	return nil
}

func TestExecutorHTTPClientFactoryUsesRequestCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	command := api.Command{
		OperationID:     "testOperation",
		Unauthenticated: true,
		RequestParameters: api.RequestParameters{
			URL: "/test/url",
		},
		Versions: []api.CommandVersion{{
			Version:              api.NewStableVersion(1991, 5, 17),
			RequestContentType:   "json",
			ResponseContentTypes: []string{"json"},
		}},
	}
	commandRequest := CommandRequest{
		Command:     command,
		ContentType: "json",
		Format:      "json",
		Version:     api.NewStableVersion(1991, 5, 17),
	}

	commandConverter := NewMockCommandConverter(ctrl)
	commandConverter.EXPECT().ConvertToHTTPRequest(commandRequest).Return(&http.Request{}, nil)

	httpClient := NewMockDoer(ctrl)
	httpClient.EXPECT().Do(gomock.Any()).Return(&http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(`{"success": true}`)),
	}, nil)

	logger := NewMockLogger(ctrl)
	logger.EXPECT().IsDebugLevel().Return(false).AnyTimes()

	var gotCommand api.Command
	executor, err := newExecutorWithHTTPClientFactory(commandConverter, func(command api.Command) (Doer, error) {
		gotCommand = command
		return httpClient, nil
	}, NewFormatter(), logger)
	require.NoError(t, err)

	response, err := executor.ExecuteCommand(t.Context(), commandRequest)

	require.NoError(t, err)
	require.NotNil(t, response)
	require.True(t, gotCommand.Unauthenticated)
	require.Equal(t, "testOperation", gotCommand.OperationID)
}
