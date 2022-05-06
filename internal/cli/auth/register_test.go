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
// +build unit

package auth

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/auth"
	atlas "go.mongodb.org/atlas/mongodbatlas"
)

type confirmPromptMock struct {
	message   string
	nbOfCalls int
	responses []bool
	outWriter io.Writer
}

func (c *confirmPromptMock) confirm() (bool, error) {
	c.nbOfCalls++
	_, _ = fmt.Fprintf(c.outWriter, "? "+c.message+" (Y/n)\n")
	return c.responses[c.nbOfCalls-1], nil
}

func TestRegisterBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		RegisterBuilder(),
		0,
		[]string{"gov", "noBrowser"},
	)
}

func Test_registerOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockAuthenticator(ctrl)
	mockConfig := mocks.NewMockLoginConfig(ctrl)
	mockStore := mocks.NewMockProjectOrgsLister(ctrl)
	defer ctrl.Finish()
	buf := new(bytes.Buffer)
	ctx := context.TODO()

	loginOpts := &LoginOpts{
		flow:       mockFlow,
		config:     mockConfig,
		NoBrowser:  true,
		SkipConfig: true,
	}

	opts := &registerOpts{
		login:                loginOpts,
		regenerateCodePrompt: nil,
	}

	opts.OutWriter = buf
	opts.login.OutWriter = buf
	opts.login.Store = mockStore

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
		PollToken(ctx, expectedCode).
		Return(expectedToken, nil, nil).
		Times(1)

	mockConfig.EXPECT().Set("service", "cloud").Times(1)
	mockConfig.EXPECT().Set("access_token", "asdf").Times(1)
	mockConfig.EXPECT().Set("refresh_token", "querty").Times(1)
	mockConfig.EXPECT().Set("ops_manager_url", gomock.Any()).Times(0)
	mockConfig.EXPECT().AccessTokenSubject().Return("test@10gen.com", nil).Times(1)
	mockConfig.EXPECT().Save().Return(nil).Times(1)
	expectedOrgs := &atlas.Organizations{}
	mockStore.EXPECT().Organizations(gomock.Any()).Return(expectedOrgs, nil).Times(0)
	expectedProjects := &atlas.Projects{}
	mockStore.EXPECT().Projects(gomock.Any()).Return(expectedProjects, nil).Times(0)

	require.NoError(t, opts.Run(ctx))
	assert.Equal(t, `Create and verify your MongoDB Atlas account from the web browser and return to Atlas CLI after activation.

First, copy your one-time code: 1234-5678

Next, sign in with your browser and enter the code.

Or go to https://account.mongodb.com/account/register?fromURI=https://account.mongodb.com/account/connect

Your code will expire after 5 minutes.
Successfully logged in as test@10gen.com.
`, buf.String())
}

func Test_registerOpts_registerAndAuthenticate(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockAuthenticator(ctrl)
	mockConfig := mocks.NewMockLoginConfig(ctrl)
	defer ctrl.Finish()
	buf := new(bytes.Buffer)
	ctx := context.TODO()

	loginOpts := &LoginOpts{
		flow:       mockFlow,
		config:     mockConfig,
		NoBrowser:  true,
		SkipConfig: true,
	}

	opts := &registerOpts{
		login:                loginOpts,
		regenerateCodePrompt: nil,
	}

	opts.login.OutWriter = buf

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
		PollToken(ctx, expectedCode).
		Return(expectedToken, nil, nil).
		Times(1)

	require.NoError(t, opts.registerAndAuthenticate(ctx))
	assert.Equal(t, `
First, copy your one-time code: 1234-5678

Next, sign in with your browser and enter the code.

Or go to https://account.mongodb.com/account/register?fromURI=https://account.mongodb.com/account/connect

Your code will expire after 5 minutes.
`, buf.String())
}

func Test_registerOpts_registerAndAuthenticate_pollTimeout(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockFlow := mocks.NewMockAuthenticator(ctrl)
	mockConfig := mocks.NewMockLoginConfig(ctrl)
	defer ctrl.Finish()
	buf := new(bytes.Buffer)
	ctx := context.TODO()
	regenerateCodePromptMock := &confirmPromptMock{
		message:   "Your one-time verification code is expired. Would you like to generate a new one?",
		nbOfCalls: 0,
		responses: []bool{true, false},
		outWriter: buf,
	}

	loginOpts := &LoginOpts{
		flow:       mockFlow,
		config:     mockConfig,
		NoBrowser:  true,
		SkipConfig: true,
	}

	opts := &registerOpts{
		login:                loginOpts,
		regenerateCodePrompt: regenerateCodePromptMock,
	}

	opts.login.OutWriter = buf

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
		Times(2)

	mockFlow.
		EXPECT().
		PollToken(ctx, expectedCode).
		Return(nil, nil, auth.ErrTimeout).
		Times(2)

	err := opts.registerAndAuthenticate(ctx)
	assert.Equal(t, err, auth.ErrTimeout)
	assert.Equal(t, `
First, copy your one-time code: 1234-5678

Next, sign in with your browser and enter the code.

Or go to https://account.mongodb.com/account/register?fromURI=https://account.mongodb.com/account/connect

Your code will expire after 5 minutes.
? Your one-time verification code is expired. Would you like to generate a new one? (Y/n)

First, copy your one-time code: 1234-5678

Next, sign in with your browser and enter the code.

Or go to https://account.mongodb.com/account/register?fromURI=https://account.mongodb.com/account/connect

Your code will expire after 5 minutes.
? Your one-time verification code is expired. Would you like to generate a new one? (Y/n)
`, buf.String())
}

func Test_registerOpts_RegisterPreRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	loginOpts := &LoginOpts{
		flow:       mocks.NewMockAuthenticator(ctrl),
		config:     mocks.NewMockLoginConfig(ctrl),
		NoBrowser:  true,
		SkipConfig: true,
	}
	defer ctrl.Finish()
	buf := new(bytes.Buffer)

	opts := &registerOpts{
		login:                loginOpts,
		regenerateCodePrompt: nil,
	}

	opts.OutWriter = buf
	opts.login.OutWriter = buf

	config.SetPublicAPIKey("public")
	config.SetPrivateAPIKey("private")
	require.ErrorContains(t, opts.registerPreRun(), fmt.Sprintf(AlreadyAuthenticatedMsg, "public"), WithProfileMsg)
}
