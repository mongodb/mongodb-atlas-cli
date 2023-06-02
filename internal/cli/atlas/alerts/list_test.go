// Copyright 2023 MongoDB Inc
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

package alerts

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongodb-atlas-cli/internal/cli"
	mocks "github.com/mongodb/mongodb-atlas-cli/internal/mocks/atlas"
	"github.com/mongodb/mongodb-atlas-cli/internal/pointer"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/atlas-sdk/admin"
)

func TestList_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockAlertLister(ctrl)

	expected := &admin.PaginatedAlert{
		Links: nil,
		Results: []admin.AlertViewForNdsGroup{
			{
				Id:            pointer.Get("test"),
				EventTypeName: pointer.Get("NO_PRIMARY"),
				Status:        pointer.Get("test"),
				MetricName:    pointer.Get("test"),
			},
		},
	}

	buf := new(bytes.Buffer)

	listOpts := &ListOpts{
		store:  mockStore,
		status: "OPEN",
		OutputOpts: cli.OutputOpts{
			Template:  listTemplate,
			OutWriter: buf,
		},
	}

	params := &admin.ListAlertsApiParams{
		GroupId:      listOpts.ProjectID,
		ItemsPerPage: &listOpts.ItemsPerPage,
		PageNum:      &listOpts.PageNum,
		Status:       &listOpts.status,
	}

	mockStore.
		EXPECT().
		Alerts(params).
		Return(expected, nil).
		Times(1)

	if err := listOpts.Run(); err != nil {
		t.Fatalf("Run() unexpected error: %v", err)
	}

	assert.Equal(t, `ID     TYPE         STATUS
test   NO_PRIMARY   test
`, buf.String())
	t.Log(buf.String())
}
