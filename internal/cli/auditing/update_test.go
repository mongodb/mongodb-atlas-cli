// Copyright 2024 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build unit

package auditing

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestUpdateOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAuditingUpdater(ctrl)
	buf := new(bytes.Buffer)
	opts := &UpdateOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			OutWriter: buf,
			Template:  updateTemplate,
		},
		auditFilter:               "test",
		enabled:                   false,
		auditAuthorizationSuccess: false,
	}

	expected := &atlasv2.AuditLog{
		Enabled:                   pointer.Get(false),
		AuditFilter:               pointer.Get("test"),
		AuditAuthorizationSuccess: pointer.Get(false),
	}

	body, err := opts.newAuditLog()
	require.NoError(t, err)
	mockStore.
		EXPECT().
		UpdateAuditingConfig(opts.ConfigProjectID(), body).
		Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run())
	assert.Equal(t, "Auditing configuration successfully updated.\n", buf.String())
	t.Log(buf.String())
}

func TestUpdateBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		UpdateBuilder(),
		0,
		[]string{flag.Output, flag.ProjectID, flag.AuditAuthorizationSuccess, flag.AuditFilter, flag.Enabled, flag.File},
	)
}
