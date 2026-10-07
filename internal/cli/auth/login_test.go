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
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prompt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312026/admin"
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

func Test_loginOpts_promptAuthType(t *testing.T) {
	t.Run("valid authType flag skips the prompt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		opts := &LoginOpts{Asker: mockAsker, authType: prompt.ServiceAccountAuth}

		require.NoError(t, opts.promptAuthType())
		assert.Equal(t, prompt.ServiceAccountAuth, opts.authType)
	})

	t.Run("invalid authType flag errors without prompting", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		opts := &LoginOpts{Asker: mockAsker, authType: "NotARealType"}

		require.Error(t, opts.promptAuthType())
	})

	t.Run("authType flag takes precedence over force", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		opts := &LoginOpts{Asker: mockAsker, authType: prompt.APIKeysAuth, force: true}

		require.NoError(t, opts.promptAuthType())
		assert.Equal(t, prompt.APIKeysAuth, opts.authType)
	})

	t.Run("force defaults to UserAccount without prompting when authType is unset", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		opts := &LoginOpts{Asker: mockAsker, force: true}

		require.NoError(t, opts.promptAuthType())
		assert.Equal(t, userAccountAuth, opts.authType)
	})

	t.Run("no flags falls back to the interactive prompt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		mockAsker.EXPECT().TrackAskOne(gomock.Any(), gomock.Any()).
			DoAndReturn(func(_ survey.Prompt, response any, _ ...survey.AskOpt) error {
				*response.(*string) = prompt.APIKeysAuth
				return nil
			}).Times(1)
		opts := &LoginOpts{Asker: mockAsker}

		require.NoError(t, opts.promptAuthType())
		assert.Equal(t, prompt.APIKeysAuth, opts.authType)
	})
}

func Test_loginOpts_promptOutput(t *testing.T) {
	t.Run("valid output flag skips the prompt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		opts := &LoginOpts{Asker: mockAsker}
		opts.Output = "json"

		require.NoError(t, opts.promptOutput())
		assert.Equal(t, "json", opts.Output)
	})

	t.Run("invalid output flag errors without prompting", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		opts := &LoginOpts{Asker: mockAsker}
		opts.Output = "xml"

		require.Error(t, opts.promptOutput())
	})

	t.Run("no flag falls back to the interactive prompt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		mockAsker.EXPECT().TrackAsk(gomock.Any(), gomock.Any()).Times(1)
		opts := &LoginOpts{Asker: mockAsker}

		require.NoError(t, opts.promptOutput())
	})
}

func Test_loginOpts_validateLoginFlags(t *testing.T) {
	tests := []struct {
		name                                                            string
		authType, output, clientID, clientSecret, publicKey, privateKey string
		wantErr                                                         bool
	}{
		{name: "no flags set is valid (fully interactive)"},
		{name: "credential flag set without authType errors", clientID: "id", wantErr: true},
		{name: "output set without authType errors", output: "json", wantErr: true},
		{name: "authType set without output errors", authType: userAccountAuth, wantErr: true},
		{name: "authType UserAccount with output is valid", authType: userAccountAuth, output: "json"},
		{name: "authType ServiceAccount without credentials errors", authType: prompt.ServiceAccountAuth, output: "json", wantErr: true},
		{name: "authType ServiceAccount with only clientId errors", authType: prompt.ServiceAccountAuth, output: "json", clientID: "id", wantErr: true},
		{name: "authType ServiceAccount with full credentials is valid", authType: prompt.ServiceAccountAuth, output: "json", clientID: "id", clientSecret: "secret"},
		{name: "authType APIKeys without credentials errors", authType: prompt.APIKeysAuth, output: "json", wantErr: true},
		{name: "authType APIKeys with full credentials is valid", authType: prompt.APIKeysAuth, output: "json", publicKey: "pub", privateKey: "priv"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &LoginOpts{
				authType:      tt.authType,
				ClientID:      tt.clientID,
				ClientSecret:  tt.clientSecret,
				PublicAPIKey:  tt.publicKey,
				PrivateAPIKey: tt.privateKey,
			}
			opts.Output = tt.output

			err := opts.validateLoginFlags()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func Test_credentialsProvided(t *testing.T) {
	tests := []struct {
		name string
		opts LoginOpts
		want bool
	}{
		{
			name: "service account both flags set",
			opts: LoginOpts{authType: prompt.ServiceAccountAuth, ClientID: "id", ClientSecret: "secret"},
			want: true,
		},
		{
			name: "service account only clientId set",
			opts: LoginOpts{authType: prompt.ServiceAccountAuth, ClientID: "id"},
			want: false,
		},
		{
			name: "service account neither flag set",
			opts: LoginOpts{authType: prompt.ServiceAccountAuth},
			want: false,
		},
		{
			name: "api keys both flags set",
			opts: LoginOpts{authType: prompt.APIKeysAuth, PublicAPIKey: "pub", PrivateAPIKey: "priv"},
			want: true,
		},
		{
			name: "api keys only publicApiKey set",
			opts: LoginOpts{authType: prompt.APIKeysAuth, PublicAPIKey: "pub"},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.opts.credentialsProvided())
		})
	}
}

func Test_loginOpts_setProgrammaticCredentials(t *testing.T) {
	t.Run("both credential flags set skips the prompt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		mockConfig := NewMockLoginConfig(ctrl)
		opts := &LoginOpts{
			Asker:        mockAsker,
			config:       mockConfig,
			authType:     prompt.ServiceAccountAuth,
			ClientID:     "id",
			ClientSecret: "secret",
		}
		opts.OutWriter = new(bytes.Buffer)

		mockConfig.EXPECT().SetService(gomock.Any()).Times(1)
		mockConfig.EXPECT().SetClientID("id").Times(1)
		mockConfig.EXPECT().SetClientSecret("secret").Times(1)

		require.NoError(t, opts.setProgrammaticCredentials())
		assert.Empty(t, opts.OutWriter.(*bytes.Buffer).String())
	})

	t.Run("partial credential flags fall back to the interactive prompt", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		mockAsker := NewMockTrackAsker(ctrl)
		mockConfig := NewMockLoginConfig(ctrl)
		opts := &LoginOpts{Asker: mockAsker, config: mockConfig, authType: prompt.ServiceAccountAuth, ClientID: "id"}
		opts.OutWriter = new(bytes.Buffer)

		mockAsker.EXPECT().TrackAsk(gomock.Any(), opts).Return(nil).Times(1)
		mockConfig.EXPECT().SetService(gomock.Any()).Times(1)
		mockConfig.EXPECT().SetClientID("id").Times(1)

		require.NoError(t, opts.setProgrammaticCredentials())
	})
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
