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

package interfaces

import (
	"testing"

	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/flag"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/mocks"
	"github.com/fmenezes/mongodb-atlas-cli/atlascli/internal/test"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/mongodbatlas"
)

func TestDescribeOpts_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockInterfaceEndpointDescriberDeprecated(ctrl)

	opts := &DescribeOpts{
		store:             mockStore,
		privateEndpointID: "privateEndpointID",
		id:                "id",
	}

	expected := &mongodbatlas.InterfaceEndpointConnectionDeprecated{}

	mockStore.
		EXPECT().
		InterfaceEndpointDeprecated(opts.ProjectID, opts.privateEndpointID, opts.id).
		Return(expected, nil).
		Times(1)

	err := opts.Run()
	require.NoError(t, err)
}

func TestDescribeBuilder(t *testing.T) {
	test.CmdValidator(
		t,
		DescribeBuilder(),
		0,
		[]string{flag.Output, flag.ProjectID, flag.PrivateEndpointID},
	)
}
