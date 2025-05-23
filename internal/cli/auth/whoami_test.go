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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_whoOpts_Run(t *testing.T) {
	buf := new(bytes.Buffer)
	opts := &whoOpts{
		OutWriter:   buf,
		authSubject: "test@test.com",
		authType:    "account",
	}
	require.NoError(t, opts.Run())
	assert.Equal(t, "Logged in as test@test.com account\n", buf.String())
}
