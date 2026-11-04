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
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/atlas-cli-core/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/api"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/telemetry"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312027/admin"
	authcode "go.mongodb.org/atlas-sdk/v20250312027/auth/code"
	"go.mongodb.org/atlas/auth"
	"go.uber.org/mock/gomock"
)

func Test_newRegisterOpts_setsAsker(t *testing.T) {
	opts := newRegisterOpts()
	require.NotNil(t, opts.Asker)
	assert.IsType(t, &telemetry.Ask{}, opts.Asker)
}

const testRegistrationURL = "https://account.mongodb.com/account/register/cli"

func Test_registerOpts_Run(t *testing.T) {
	t.Setenv("MCLI_CLIENT_ID", "")
	profile := setDefaultProfile(t, "", "")
	profile.SetOpsManagerURL("https://cloud-dev.mongodb.com/")

	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := NewMockLoginConfig(ctrl)
	mockAuthCode := NewMockAuthorizationCodeFlow(ctrl)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)
	mockAsker := NewMockTrackAsker(ctrl)

	buf := new(bytes.Buffer)
	ctx := t.Context()

	opts := &RegisterOpts{}
	opts.NoBrowser = true
	opts.config = mockConfig
	opts.OutWriter = buf
	opts.Store = mockStore
	opts.Asker = mockAsker
	opts.savedProfile = savedProfileWith(nil)
	opts.authCode = mockAuthCode
	opts.WithFlow(mockFlow)

	mockFlow.EXPECT().
		RegistrationConfig(ctx).
		Return(&auth.RegistrationConfig{RegistrationURL: testRegistrationURL}, nil, nil).
		Times(1)

	metadata := &authcode.AuthServerMetadata{
		Issuer:      "https://authorize.mongodb.com",
		CacheExpiry: time.Now().Add(time.Hour),
	}
	mockConfig.EXPECT().AuthServerMetadata().Return(nil).Times(1)
	mockAuthCode.EXPECT().ResolveAuthServerMetadata(ctx, nil).Return(metadata, nil).Times(1)
	mockConfig.EXPECT().SetAuthServerMetadata(metadata).Times(1)

	const authorizationCode = "auth-code-123"
	mockAuthCode.EXPECT().
		AuthorizationURL(metadata, gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(completeAuthCodeCallback(ctx, authorizationCode)).
		Times(1)

	tokenExpiry := time.Now().Add(time.Hour).UTC()
	mockAuthCode.EXPECT().
		ExchangeCode(ctx, metadata, authorizationCode, gomock.Any(), gomock.Any()).
		Return(&authcode.Token{AccessToken: "asdf", RefreshToken: "querty", Expiry: tokenExpiry}, nil).
		Times(1)

	mockConfig.EXPECT().SetAuthType(config.UserAccount).Times(1)
	mockConfig.EXPECT().SetService("cloud").Times(1)
	mockConfig.EXPECT().SetAccessToken("asdf").Times(2)
	mockConfig.EXPECT().SetRefreshToken("querty").Times(2)
	mockConfig.EXPECT().SetTokenExpiry(tokenExpiry.Format(time.RFC3339)).Times(1)
	mockConfig.EXPECT().SetAuthServerIssuer(metadata.Issuer).Times(1)
	mockConfig.EXPECT().SetOpsManagerURL(gomock.Any()).Times(0)
	mockConfig.EXPECT().OrgID().Return("").AnyTimes()
	mockConfig.EXPECT().ProjectID().Return("").AnyTimes()
	mockConfig.EXPECT().AccessTokenSubject().Return("test@10gen.com", nil).Times(1)
	mockConfig.EXPECT().Save().Return(nil).Times(3)

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
	mockStore.
		EXPECT().
		GetOrgProjects("o1", gomock.Any()).
		Return(expectedProjects, nil).
		Times(1)

	mockAsker.EXPECT().
		TrackAsk(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ []*survey.Question, answer any, _ ...survey.AskOpt) error {
			if o, ok := answer.(*LoginOpts); ok {
				o.Output = jsonOutputFormat
			}
			return nil
		})

	require.NoError(t, opts.RegisterRun(ctx))
	assert.Equal(t, `To continue, go to https://account.mongodb.com/account/register?n=https%3A%2F%2Fauthorize.mongodb.com%2Fauthorize
Successfully logged in as test@10gen.com.
`, buf.String())
}

func Test_registerOpts_Run_ExchangeFailsDoesNotSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := NewMockLoginConfig(ctrl)
	mockAuthCode := NewMockAuthorizationCodeFlow(ctrl)

	buf := new(bytes.Buffer)
	ctx := t.Context()

	opts := &RegisterOpts{}
	opts.NoBrowser = true
	opts.config = mockConfig
	opts.OutWriter = buf
	opts.savedProfile = savedProfileWith(nil)
	opts.authCode = mockAuthCode
	opts.WithFlow(mockFlow)

	mockFlow.EXPECT().
		RegistrationConfig(ctx).
		Return(&auth.RegistrationConfig{RegistrationURL: testRegistrationURL}, nil, nil).
		Times(1)

	metadata := &authcode.AuthServerMetadata{
		Issuer:      "https://authorize.mongodb.com",
		CacheExpiry: time.Now().Add(time.Hour),
	}
	mockConfig.EXPECT().AuthServerMetadata().Return(nil).Times(1)
	mockAuthCode.EXPECT().ResolveAuthServerMetadata(ctx, nil).Return(metadata, nil).Times(1)
	mockConfig.EXPECT().SetAuthType(config.UserAccount).Times(1)

	const authorizationCode = "auth-code-123"
	mockAuthCode.EXPECT().
		AuthorizationURL(metadata, gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(completeAuthCodeCallback(ctx, authorizationCode)).
		Times(1)

	errExchange := errors.New("invalid_grant")
	mockAuthCode.EXPECT().
		ExchangeCode(ctx, metadata, authorizationCode, gomock.Any(), gomock.Any()).
		Return(nil, errExchange).
		Times(1)

	mockConfig.EXPECT().Save().Times(0)
	mockConfig.EXPECT().SetAccessToken(gomock.Any()).Times(0)
	mockConfig.EXPECT().SetRefreshToken(gomock.Any()).Times(0)

	require.ErrorIs(t, opts.RegisterRun(ctx), errExchange)
	assert.NotContains(t, buf.String(), "Successfully logged in")
}

func Test_registerOpts_Run_OtherCredentialsBlocksBeforeFlow(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)

	opts := &RegisterOpts{}
	opts.config = NewMockLoginConfig(ctrl)
	opts.savedProfile = savedProfileWith(map[string]string{
		config.ClientIDField:     "mdb_sa_id_1",
		config.ClientSecretField: "secret",
	})
	opts.WithFlow(mockFlow)

	require.ErrorIs(t, opts.RegisterRun(t.Context()), validate.ErrAlreadyAuthenticatedSA)
}

func Test_registerOpts_Run_UnexpectedRegistrationURLFailsBeforeFlow(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := NewMockLoginConfig(ctrl)

	buf := new(bytes.Buffer)
	ctx := t.Context()

	opts := &RegisterOpts{}
	opts.config = mockConfig
	opts.OutWriter = buf
	opts.savedProfile = savedProfileWith(nil)
	opts.authCode = NewMockAuthorizationCodeFlow(ctrl)
	opts.WithFlow(mockFlow)

	mockFlow.EXPECT().
		RegistrationConfig(ctx).
		Return(&auth.RegistrationConfig{RegistrationURL: "https://account.example.com/account/register/cli"}, nil, nil).
		Times(1)
	mockConfig.EXPECT().SetAuthType(config.UserAccount).Times(1)
	mockConfig.EXPECT().Save().Times(0)

	require.ErrorIs(t, opts.RegisterRun(ctx), errUnexpectedRegistrationURL)
	assert.Empty(t, buf.String())
}

func Test_registerOpts_Run_RegistrationConfigFails(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := NewMockLoginConfig(ctrl)

	opts := &RegisterOpts{}
	opts.config = mockConfig
	opts.savedProfile = savedProfileWith(nil)
	opts.WithFlow(mockFlow)

	errRegistration := errors.New("registration config unavailable")
	mockConfig.EXPECT().SetAuthType(config.UserAccount).Times(1)
	mockFlow.EXPECT().RegistrationConfig(gomock.Any()).Return(nil, nil, errRegistration).Times(1)
	mockConfig.EXPECT().Save().Times(0)

	require.ErrorIs(t, opts.RegisterRun(t.Context()), errRegistration)
}

func Test_registerOpts_Run_EmptyRegistrationConfig(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := NewMockLoginConfig(ctrl)

	opts := &RegisterOpts{}
	opts.config = mockConfig
	opts.savedProfile = savedProfileWith(nil)
	opts.WithFlow(mockFlow)

	mockConfig.EXPECT().SetAuthType(config.UserAccount).Times(1)
	mockFlow.EXPECT().RegistrationConfig(gomock.Any()).Return(nil, nil, nil).Times(1)
	mockConfig.EXPECT().Save().Times(0)

	require.ErrorIs(t, opts.RegisterRun(t.Context()), errUnexpectedRegistrationURL)
}

func Test_accountRegisterURL(t *testing.T) {
	tests := []struct {
		name            string
		registrationURL string
		want            string
		wantErr         bool
	}{
		{name: "prod", registrationURL: "https://account.mongodb.com/account/register/cli", want: "https://account.mongodb.com/account/register"},
		{name: "gov", registrationURL: "https://account.mongodbgov.com/account/register/cli", want: "https://account.mongodbgov.com/account/register"},
		{name: "drops query", registrationURL: "https://account-dev.mongodb.com/account/register/cli?n=https%3A%2F%2Fevil.example.com", want: "https://account-dev.mongodb.com/account/register"},
		{name: "http", registrationURL: "http://account.mongodb.com/account/register/cli", wantErr: true},
		{name: "other host", registrationURL: "https://account.example.com/account/register/cli", wantErr: true},
		{name: "suffix without dot", registrationURL: "https://evilmongodb.com/account/register/cli", wantErr: true},
		{name: "mongodb host as subdomain", registrationURL: "https://account.mongodb.com.example.com/register", wantErr: true},
		{name: "userinfo", registrationURL: "https://account.mongodb.com@example.com/register", wantErr: true},
		{name: "empty", registrationURL: "", wantErr: true},
		{name: "unparseable", registrationURL: "https://account.mongodb.com/%zz", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := accountRegisterURL(tt.registrationURL)
			if tt.wantErr {
				require.ErrorIs(t, err, errUnexpectedRegistrationURL)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got.String())
		})
	}
}
