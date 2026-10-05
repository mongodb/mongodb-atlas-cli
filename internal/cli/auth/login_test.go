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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/commonerrors"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prompt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312025/admin"
	"go.mongodb.org/atlas/auth"
	"go.uber.org/mock/gomock"
)

func Test_loginOpts_SyncWithOAuthAccessProfile(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := []struct {
		name            string
		isGov           bool
		expectedService string
	}{
		{name: "cloud service run", isGov: false, expectedService: "cloud"},
		{name: "cloudgov service run", isGov: true, expectedService: "cloudgov"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockConfig := NewMockLoginConfig(ctrl)
			opts := &LoginOpts{
				NoBrowser:    true,
				AccessToken:  "at",
				RefreshToken: "rt",
				IsGov:        tt.isGov,
			}
			opts.OutWriter = new(bytes.Buffer)

			mockConfig.EXPECT().SetService(tt.expectedService).Times(1)
			mockConfig.EXPECT().SetAccessToken(opts.AccessToken).Times(1)
			mockConfig.EXPECT().SetRefreshToken(opts.RefreshToken).Times(1)
			mockConfig.EXPECT().SetClientID(gomock.Any()).Times(0)
			mockConfig.EXPECT().SetOpsManagerURL(gomock.Any()).Times(0)

			require.NoError(t, opts.SyncWithOAuthAccessProfile(mockConfig)())
		})
	}
}

func Test_loginOpts_LoginRun_UserAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := NewMockLoginConfig(ctrl)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)
	mockAsker := NewMockTrackAsker(ctrl)

	buf := new(bytes.Buffer)

	opts := &LoginOpts{
		config:    mockConfig,
		Asker:     mockAsker,
		NoBrowser: true,
	}
	opts.WithFlow(mockFlow)
	opts.OutWriter = buf
	opts.Store = mockStore

	mockAsker.EXPECT().
		TrackAskOne(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ survey.Prompt, answer any, _ ...survey.AskOpt) error {
			if s, ok := answer.(*string); ok {
				*s = userAccountAuth
			}
			return nil
		})

	expectedCode := &auth.DeviceCode{
		UserCode:        "12345678",
		VerificationURI: "http://localhost",
		DeviceCode:      "123",
		ExpiresIn:       300,
		Interval:        10,
	}
	ctx := t.Context()
	mockFlow.
		EXPECT().
		RequestCode(ctx).
		Return(expectedCode, nil, nil).
		Times(1)

	expectedToken := &auth.Token{
		AccessToken:  "asdf",
		RefreshToken: "querty",
		Scope:        "openid",
		IDToken:      "1",
		TokenType:    "Bearer",
		ExpiresIn:    3600,
	}
	mockFlow.
		EXPECT().
		PollToken(ctx, expectedCode).
		Return(expectedToken, nil, nil).
		Times(1)

	mockConfig.EXPECT().SetAuthType(config.UserAccount).Times(1)
	mockConfig.EXPECT().SetService("cloud").Times(1)
	mockConfig.EXPECT().SetAccessToken("asdf").Times(1)
	mockConfig.EXPECT().SetRefreshToken("querty").Times(1)
	mockConfig.EXPECT().SetOpsManagerURL(gomock.Any()).Times(0)
	mockConfig.EXPECT().OrgID().Return("").AnyTimes()
	mockConfig.EXPECT().ProjectID().Return("").AnyTimes()
	mockConfig.EXPECT().AccessTokenSubject().Return("test@10gen.com", nil).Times(1)
	mockConfig.EXPECT().Save().Return(nil).Times(1)
	mockConfig.EXPECT().AuthType().Return(config.UserAccount).AnyTimes()

	opts.SkipConfig = true

	err := opts.LoginRun(ctx)
	require.NoError(t, err)
}

func TestLoginRun_APIKeys_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConfig := NewMockLoginConfig(ctrl)
	mockAsker := NewMockTrackAsker(ctrl)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)

	opts := &LoginOpts{
		config: mockConfig,
		Asker:  mockAsker,
	}
	opts.OutWriter = new(bytes.Buffer)
	opts.Store = mockStore

	mockAsker.EXPECT().
		TrackAskOne(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ survey.Prompt, answer any, _ ...survey.AskOpt) error {
			if s, ok := answer.(*string); ok {
				*s = prompt.APIKeysAuth
			}
			return nil
		})

	mockAsker.EXPECT().
		TrackAsk(gomock.Any(), opts).
		DoAndReturn(func(_ []*survey.Question, answer any, _ ...survey.AskOpt) error {
			if o, ok := answer.(*LoginOpts); ok {
				o.PublicAPIKey = "public-key"
				o.PrivateAPIKey = "private-key"
			}
			return nil
		})

	mockConfig.EXPECT().SetAuthType(config.APIKeys).Times(1)
	mockConfig.EXPECT().SetService("cloud").Times(1)
	mockConfig.EXPECT().SetPublicAPIKey("public-key").Times(1)
	mockConfig.EXPECT().SetPrivateAPIKey("private-key").Times(1)
	mockConfig.EXPECT().PublicAPIKey().Return("public-key").AnyTimes()
	mockConfig.EXPECT().PrivateAPIKey().Return("private-key").AnyTimes()
	mockConfig.EXPECT().AuthType().Return(config.APIKeys).AnyTimes()

	opts.SkipConfig = true

	err := opts.LoginRun(t.Context())
	require.NoError(t, err)
}

type confirmMock struct{}

func (confirmMock) Prompt(_ *survey.PromptConfig) (any, error) {
	return true, nil
}

func (confirmMock) Cleanup(_ *survey.PromptConfig, _ any) error {
	return nil
}

func (confirmMock) Error(_ *survey.PromptConfig, err error) error {
	return err
}

func Test_shouldRetryAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockAsker := NewMockTrackAsker(ctrl)
	opts := &LoginOpts{Asker: mockAsker}

	type args struct {
		err error
		p   survey.Prompt
	}
	tests := []struct {
		name      string
		args      args
		wantRetry bool
		wantErr   require.ErrorAssertionFunc
	}{
		{
			name: "timed out error",
			args: args{
				err: auth.ErrTimeout,
				p:   &confirmMock{},
			},
			wantRetry: true,
			wantErr:   require.NoError,
		},
		{
			name: "random error",
			args: args{
				err: errors.New("random"),
				p:   &confirmMock{},
			},
			wantRetry: false,
			wantErr:   require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAsker.EXPECT().TrackAskOne(gomock.Any(), gomock.Any()).DoAndReturn(
				func(_ survey.Prompt, answer any, _ ...survey.AskOpt) error {
					if b, ok := answer.(*bool); ok {
						*b = tt.wantRetry
					}
					return nil
				},
			).AnyTimes()
			gotRetry, err := opts.shouldRetryAuthenticate(tt.args.err, tt.args.p)
			tt.wantErr(t, err, fmt.Sprintf("shouldRetryAuthenticate(%v, %v)", tt.args.err, tt.args.p))
			assert.Equalf(t, tt.wantRetry, gotRetry, "shouldRetryAuthenticate(%v, %v)", tt.args.err, tt.args.p)
		})
	}
}

func TestLoginOpts_setUpProfile_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockConfig := NewMockLoginConfig(ctrl)
	mockAsker := NewMockTrackAsker(ctrl)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)

	buf := new(bytes.Buffer)
	opts := &LoginOpts{
		config: mockConfig,
		Asker:  mockAsker,
	}
	opts.OutWriter = buf
	opts.Store = mockStore

	opts.OrgID = ""
	opts.ProjectID = ""

	mockConfig.EXPECT().OrgID().Return("").Times(1)
	mockConfig.EXPECT().ProjectID().Return("").Times(1)

	orgsBody, err := json.Marshal(&admin.PaginatedOrganization{
		TotalCount: pointer.Get(1),
		Results: []admin.AtlasOrganization{
			{Id: pointer.Get("o1"), Name: "Org1"},
		},
	})
	require.NoError(t, err)

	mockExecutor := api.NewMockCommandExecutor(ctrl)
	mockExecutor.EXPECT().
		ExecuteCommand(gomock.Any(), gomock.Any()).
		Return(&api.CommandResponse{
			IsSuccess: true,
			HTTPCode:  http.StatusOK,
			Output:    io.NopCloser(bytes.NewReader(orgsBody)),
		}, nil).
		Times(1)
	opts.OrgExecutor = mockExecutor

	expectedProjects := &admin.PaginatedAtlasGroup{TotalCount: pointer.Get(1),
		Results: []admin.Group{
			{Id: pointer.Get("p1"), Name: "Project1"},
		},
	}
	mockStore.EXPECT().GetOrgProjects("o1", gomock.Any()).Return(expectedProjects, nil).Times(1)
	mockAsker.EXPECT().
		TrackAsk(gomock.Any(), opts).
		DoAndReturn(func(_ []*survey.Question, answer any, _ ...survey.AskOpt) error {
			if o, ok := answer.(*LoginOpts); ok {
				o.Output = jsonOutputFormat
			}
			return nil
		})

	mockConfig.EXPECT().Save().Return(nil).Times(1)

	ctx := t.Context()
	require.NoError(t, opts.setUpProfile(ctx))
}

func TestLoginRun_ServiceAccount_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockConfig := NewMockLoginConfig(ctrl)
	mockAsker := NewMockTrackAsker(ctrl)

	buf := new(bytes.Buffer)
	opts := &LoginOpts{
		config: mockConfig,
		Asker:  mockAsker,
	}
	opts.OutWriter = buf

	mockAsker.EXPECT().
		TrackAskOne(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ survey.Prompt, answer any, _ ...survey.AskOpt) error {
			if s, ok := answer.(*string); ok {
				*s = prompt.ServiceAccountAuth
			}
			return nil
		})

	mockAsker.EXPECT().
		TrackAsk(gomock.Any(), opts).
		DoAndReturn(func(_ []*survey.Question, answer any, _ ...survey.AskOpt) error {
			if o, ok := answer.(*LoginOpts); ok {
				o.ClientID = "client-id"
				o.ClientSecret = "client-secret"
			}
			return nil
		})

	mockConfig.EXPECT().SetAuthType(config.ServiceAccount).Times(1)
	mockConfig.EXPECT().SetService("cloud").Times(1)
	mockConfig.EXPECT().SetClientID("client-id").Times(1)
	mockConfig.EXPECT().SetClientSecret("client-secret").Times(1)
	mockConfig.EXPECT().ClientID().Return("client-id").AnyTimes()
	mockConfig.EXPECT().ClientSecret().Return("client-secret").AnyTimes()
	mockConfig.EXPECT().AuthType().Return(config.ServiceAccount).AnyTimes()

	opts.SkipConfig = true

	require.NoError(t, opts.LoginRun(t.Context()))
	assert.Contains(t, buf.String(), "environment variables take precedence over the values you enter here")
	assert.Contains(t, buf.String(), commonerrors.EnvVarsDocsURL)
}

const clientIDEnvVar = config.AtlasCLIEnvPrefix + "_CLIENT_ID"

func TestLoginOpts_checkEnvAuthTypeOverride(t *testing.T) {
	tests := []struct {
		name      string
		selected  config.AuthMechanism
		effective config.AuthMechanism
		envVars   map[string]string
		errSubstr string
	}{
		{
			name:      "env credentials hijack the selected mechanism",
			selected:  config.UserAccount,
			effective: config.ServiceAccount,
			envVars:   map[string]string{clientIDEnvVar: "from-env"},
			errSubstr: "selects service_account authentication and overrides the user_account method you chose",
		},
		{
			name:      "stale profile hijacks the mechanism without env vars",
			selected:  config.UserAccount,
			effective: config.ServiceAccount,
		},
		{
			name:      "mechanism matches",
			selected:  config.UserAccount,
			effective: config.UserAccount,
			envVars:   map[string]string{clientIDEnvVar: "from-env"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envVars {
				t.Setenv(k, v)
			}

			mockConfig := NewMockLoginConfig(gomock.NewController(t))
			mockConfig.EXPECT().AuthType().Return(tt.effective).AnyTimes()
			opts := &LoginOpts{config: mockConfig}

			err := opts.checkEnvAuthTypeOverride(tt.selected)
			if tt.errSubstr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errSubstr)
			assert.Contains(t, err.Error(), commonerrors.EnvVarsDocsURL)
		})
	}
}

func TestLoginOpts_checkEnvCredentialOverride(t *testing.T) {
	tests := []struct {
		name      string
		authType  string
		envVars   map[string]string
		opts      LoginOpts
		resolved  func(*MockLoginConfig)
		errSubstr string
	}{
		{
			name:     "entered value survives",
			authType: prompt.ServiceAccountAuth,
			envVars:  map[string]string{clientIDEnvVar: "from-env"},
			opts:     LoginOpts{ClientID: "from-env", ClientSecret: "secret"},
			resolved: func(c *MockLoginConfig) {
				c.EXPECT().ClientID().Return("from-env").AnyTimes()
				c.EXPECT().ClientSecret().Return("secret").AnyTimes()
			},
		},
		{
			name:     "env var discards the entered value",
			authType: prompt.ServiceAccountAuth,
			envVars:  map[string]string{clientIDEnvVar: "from-env"},
			opts:     LoginOpts{ClientID: "typed-by-user", ClientSecret: "secret"},
			resolved: func(c *MockLoginConfig) {
				c.EXPECT().ClientID().Return("from-env").AnyTimes()
				c.EXPECT().ClientSecret().Return("secret").AnyTimes()
			},
			errSubstr: clientIDEnvVar + " takes precedence over the Client ID you entered",
		},
		{
			name:     "mismatch without the env var set is not reported",
			authType: prompt.ServiceAccountAuth,
			opts:     LoginOpts{ClientID: "typed-by-user", ClientSecret: "secret"},
			resolved: func(c *MockLoginConfig) {
				c.EXPECT().ClientID().Return("stale-from-keyring").AnyTimes()
				c.EXPECT().ClientSecret().Return("secret").AnyTimes()
			},
		},
		{
			name:     "nothing supplied",
			authType: userAccountAuth,
			resolved: func(_ *MockLoginConfig) {},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envVars {
				t.Setenv(k, v)
			}

			mockConfig := NewMockLoginConfig(gomock.NewController(t))
			tt.resolved(mockConfig)

			opts := tt.opts
			opts.authType = tt.authType
			opts.config = mockConfig

			err := opts.checkEnvCredentialOverride()
			if tt.errSubstr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errSubstr)
			assert.Contains(t, err.Error(), commonerrors.EnvVarsDocsURL)
		})
	}
}

// The hijack is detectable from the environment alone, so LoginRun must fail before
// setUpCredentials sends the user through the OAuth device flow.
func TestLoginRun_HijackFailsBeforeDeviceFlow(t *testing.T) {
	t.Setenv(clientIDEnvVar, "from-env")

	ctrl := gomock.NewController(t)
	mockConfig := NewMockLoginConfig(ctrl)
	mockAsker := NewMockTrackAsker(ctrl)
	mockFlow := mocks.NewMockRefresher(ctrl)

	opts := &LoginOpts{config: mockConfig, Asker: mockAsker, NoBrowser: true}
	opts.WithFlow(mockFlow)
	opts.OutWriter = new(bytes.Buffer)

	mockAsker.EXPECT().
		TrackAskOne(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ survey.Prompt, answer any, _ ...survey.AskOpt) error {
			if s, ok := answer.(*string); ok {
				*s = userAccountAuth
			}
			return nil
		})

	mockConfig.EXPECT().SetAuthType(config.UserAccount).Times(1)
	mockConfig.EXPECT().AuthType().Return(config.ServiceAccount).AnyTimes()
	// the device flow must never start
	mockFlow.EXPECT().RequestCode(gomock.Any()).Times(0)
	mockConfig.EXPECT().Save().Times(0)

	err := opts.LoginRun(t.Context())
	require.Error(t, err)
	assert.Contains(t, err.Error(), clientIDEnvVar)
}
