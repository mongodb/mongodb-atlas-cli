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

//go:build unit

package auth

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/prompt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20250312005/admin"
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

	err := opts.LoginRun(context.Background())
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

	expectedOrgs := &admin.PaginatedOrganization{
		TotalCount: pointer.Get(1),
		Results: &[]admin.AtlasOrganization{
			{Id: pointer.Get("o1"), Name: "Org1"},
		},
	}
	mockStore.EXPECT().Organizations(gomock.Any()).Return(expectedOrgs, nil).Times(1)
	expectedProjects := &admin.PaginatedAtlasGroup{TotalCount: pointer.Get(1),
		Results: &[]admin.Group{
			{Id: pointer.Get("p1"), Name: "Project1"},
		},
	}
	mockStore.EXPECT().GetOrgProjects("o1", gomock.Any()).Return(expectedProjects, nil).Times(1)
	mockAsker.EXPECT().
		TrackAsk(gomock.Any(), opts).
		DoAndReturn(func(_ []*survey.Question, answer any, _ ...survey.AskOpt) error {
			if o, ok := answer.(*LoginOpts); ok {
				o.Output = "json"
			}
			return nil
		})

	mockConfig.EXPECT().Save().Return(nil).Times(1)

	ctx := context.Background()
	err := opts.setUpProfile(ctx)
	require.NoError(t, err)
}
