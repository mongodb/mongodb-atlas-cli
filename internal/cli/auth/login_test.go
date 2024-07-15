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
	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
	"go.mongodb.org/atlas/auth"
)

func TestLoginBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		LoginBuilder(),
		0,
		[]string{"gov", "noBrowser"},
	)
}

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
			mockConfig := mocks.NewMockLoginConfig(ctrl)
			opts := &LoginOpts{
				NoBrowser:    true,
				AccessToken:  "at",
				RefreshToken: "rt",
				IsGov:        tt.isGov,
			}
			opts.OutWriter = new(bytes.Buffer)

			mockConfig.EXPECT().Set("service", tt.expectedService).Times(1)
			mockConfig.EXPECT().Set("access_token", opts.AccessToken).Times(1)
			mockConfig.EXPECT().Set("refresh_token", opts.RefreshToken).Times(1)
			mockConfig.EXPECT().Set("ops_manager_url", gomock.Any()).Times(0)

			require.NoError(t, opts.SyncWithOAuthAccessProfile(mockConfig)())
		})
	}
}

func Test_loginOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := mocks.NewMockLoginConfig(ctrl)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)

	buf := new(bytes.Buffer)

	opts := &LoginOpts{
		config:    mockConfig,
		NoBrowser: true,
	}
	opts.WithFlow(mockFlow)

	opts.OutWriter = buf
	opts.Store = mockStore
	expectedCode := &auth.DeviceCode{
		UserCode:        "12345678",
		VerificationURI: "http://localhost",
		DeviceCode:      "123",
		ExpiresIn:       300,
		Interval:        10,
	}
	ctx := context.TODO()
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

	mockConfig.EXPECT().Set("service", "cloud").Times(1)
	mockConfig.EXPECT().Set("access_token", "asdf").Times(1)
	mockConfig.EXPECT().Set("refresh_token", "querty").Times(1)
	mockConfig.EXPECT().Set("ops_manager_url", gomock.Any()).Times(0)
	mockConfig.EXPECT().OrgID().Return("").AnyTimes()
	mockConfig.EXPECT().ProjectID().Return("").AnyTimes()
	mockConfig.EXPECT().AccessTokenSubject().Return("test@10gen.com", nil).Times(1)
	mockConfig.EXPECT().Save().Return(nil).Times(2)
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
	require.NoError(t, opts.LoginRun(ctx))
	assert.Equal(t, `
To verify your account, copy your one-time verification code:
1234-5678

Paste the code in the browser when prompted to activate your Atlas CLI. Your code will expire after 5 minutes.

To continue, go to http://localhost
Successfully logged in as test@10gen.com.
`, buf.String())
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
	t.Setenv("DO_NOT_TRACK", "1")
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
			gotRetry, err := shouldRetryAuthenticate(tt.args.err, tt.args.p)
			tt.wantErr(t, err, fmt.Sprintf("shouldRetryAuthenticate(%v, %v)", tt.args.err, tt.args.p))
			assert.Equalf(t, tt.wantRetry, gotRetry, "shouldRetryAuthenticate(%v, %v)", tt.args.err, tt.args.p)
		})
	}
}
