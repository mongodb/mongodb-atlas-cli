// Copyright 2020 MongoDB Inc
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

package setup

import (
	"bytes"
	"context"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/cli/auth"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/test"
	"github.com/stretchr/testify/require"

	"testing"
)

func TestBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		Builder(),
		0,
		[]string{flag.Region, flag.ClusterName, flag.Provider, flag.AccessListIP, flag.Username, flag.Password, flag.SkipMongosh, flag.SkipSampleData},
	)
}

func Test_registerOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRegFlow := mocks.NewMockRegisterFlow(ctrl)
	defer ctrl.Finish()
	ctx := context.TODO()
	buf := new(bytes.Buffer)

	opts := &Opts{
		register: mockRegFlow,
		login:    auth.LoginOpts{},
	}

	opts.login.OutWriter = buf

	mockRegFlow.
		EXPECT().
		Run(ctx).
		Return(nil).
		Times(1)

	require.NoError(t, opts.Run(ctx))
}
