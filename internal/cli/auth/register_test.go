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
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas-sdk/v20240530002/admin"
	"go.mongodb.org/atlas/auth"
)

func TestRegisterBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		RegisterBuilder(),
		0,
		[]string{
			"noBrowser",
		},
	)
}

func Test_registerOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockRefresher(ctrl)
	mockConfig := mocks.NewMockLoginConfig(ctrl)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)

	buf := new(bytes.Buffer)
	ctx := context.TODO()

	opts := &RegisterOpts{}
	opts.NoBrowser = true
	opts.config = mockConfig
	opts.OutWriter = buf
	opts.Store = mockStore
	opts.WithFlow(mockFlow)

	expectedCode := &auth.DeviceCode{
		UserCode:        "12345678",
		VerificationURI: "http://localhost",
		DeviceCode:      "123",
		ExpiresIn:       300,
		Interval:        10,
	}

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
		RegistrationConfig(ctx).
		Return(&auth.RegistrationConfig{RegistrationURL: "https://localhost/account/register/cli"}, nil, nil).
		Times(1)
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
	mockStore.
		EXPECT().
		Organizations(gomock.Any()).
		Return(expectedOrgs, nil).
		Times(1)
	expectedProjects := &admin.PaginatedAtlasGroup{TotalCount: pointer.Get(1),
		Results: &[]admin.Group{
			{Id: pointer.Get("p1"), Name: "Project1"},
		},
	}
	mockStore.
		EXPECT().
		GetOrgProjects("o1", gomock.Any()).
		Return(expectedProjects, nil).
		Times(1)

	require.NoError(t, opts.RegisterRun(ctx))
	assert.Equal(t, `
To verify your account, copy your one-time verification code:
1234-5678

Paste the code in the browser when prompted to activate your Atlas CLI. Your code will expire after 5 minutes.

To continue, go to https://localhost/account/register/cli
Successfully logged in as test@10gen.com.
`, buf.String())
}
