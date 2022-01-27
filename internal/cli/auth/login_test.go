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
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/auth"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		1,
		[]string{},
	)
}

func TestLoginBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		LoginBuilder(),
		0,
		[]string{"gov", "cm", "noBrowser"},
	)
}

func Test_loginOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAuthenticator(ctrl)
	defer ctrl.Finish()

	opts := &loginOpts{
		flow:      mockStore,
		OutWriter: os.NewFile(0, os.DevNull),
		noBrowser: true,
	}
	expectedCode := &auth.DeviceCode{
		UserCode:        "12345678",
		VerificationURI: "http://localhost",
		DeviceCode:      "123",
		ExpiresIn:       300,
		Interval:        10,
	}
	mockStore.
		EXPECT().
		RequestCode(gomock.Any()).
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
	mockStore.
		EXPECT().
		PollToken(gomock.Any(), expectedCode).
		Return(expectedToken, nil, nil).
		Times(1)

	require.NoError(t, opts.Run())
}
