// Copyright 2023 MongoDB Inc
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

package auditing

import (
	"bytes"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	atlasv2 "go.mongodb.org/atlas-sdk/v20250312006/admin"
	"go.uber.org/mock/gomock"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := NewMockDescriber(ctrl)
	buf := new(bytes.Buffer)
	opts := &DescribeOpts{
		store: mockStore,
		OutputOpts: cli.OutputOpts{
			Template:  describeTemplate,
			OutWriter: buf,
		},
	}

	expected := &atlasv2.AuditLog{}
	mockStore.
		EXPECT().
		Auditing(opts.ConfigProjectID()).
		Return(expected, nil).
		Times(1)

	require.NoError(t, opts.Run())
	assert.Equal(t, `AUDIT AUTHORIZATION SUCCESS   AUDIT FILTER   CONFIGURATION TYPE   ENABLED
<nil>                         <nil>          <nil>                <nil>
`, buf.String())
	t.Log(buf.String())
}
