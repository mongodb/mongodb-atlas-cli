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
	"strings"
	"testing"

	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/atlas-cli-core/transport"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
	"github.com/stretchr/testify/require"
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
		RequestParameters: api.RequestParameters{
			URL: "/api/atlas/v2/unauth/ephemeralClusters:create",
		},
	}

	client, err := newDefaultHTTPClient(&testProfileProvider{}, command)

	require.NoError(t, err)
	require.Same(t, transport.Default(), client.Transport)
}

func TestNewDefaultHTTPClient_AuthenticatedAPICommandUsesCredentials(t *testing.T) {
	expectedErr := errors.New("missing token")
	command := api.Command{
		RequestParameters: api.RequestParameters{
			URL: "/api/atlas/v2/groups",
		},
	}

	_, err := newDefaultHTTPClient(&testProfileProvider{
		authType: config.UserAccount,
		tokenErr: expectedErr,
	}, command)

	require.ErrorIs(t, err, expectedErr)
}
