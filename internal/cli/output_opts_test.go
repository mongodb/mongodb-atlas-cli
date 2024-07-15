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

package cli

import (
	"io"
	"reflect"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pointer"
	atlasv2 "go.mongodb.org/atlas-sdk/v20240530002/admin"
)

func TestOutputOpts_outputTypeAndValue(t *testing.T) {
	type fields struct {
		Template  string
		OutWriter io.Writer
		Output    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "go-template",
			fields: fields{Output: "go-template=test", Template: ""},
			want:   "test",
		},
		{
			name:   "not-valid",
			fields: fields{Output: "not-valid", Template: "default"},
			want:   "default",
		},
		{
			name:   "json-path",
			fields: fields{Output: "json-path=$.[0].id", Template: ""},
			want:   "$.[0].id",
		},
	}
	for _, tt := range tests {
		opts := &OutputOpts{
			Template:  tt.fields.Template,
			OutWriter: tt.fields.OutWriter,
			Output:    tt.fields.Output,
		}
		want := tt.want
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, got := opts.outputTypeAndValue()
			if got != want {
				t.Errorf("parseTemplate() got = %v, want %v", got, want)
			}
		})
	}
}

func TestOutputOpts_mapReduceResults(t *testing.T) {
	t.Run("when results present", func(t *testing.T) {
		input := *atlasv2.NewPaginatedTeam()
		wantID := "123"
		wantName := "Team A"
		input.Results = &[]atlasv2.TeamResponse{
			{
				Id:   pointer.Get(wantID),
				Name: pointer.Get(wantName),
			},
		}

		compactResults, err := mapReduceResults(input)
		if err != nil {
			t.Fatalf("mapReduceResults() unexpected error: %v", err)
		}

		mapArrayResponse := reflect.ValueOf(compactResults).Interface().([]any)
		mapResponse := mapArrayResponse[0].(map[string]any)
		gotID := mapResponse["id"]
		gotName := mapResponse["name"]
		if gotID != wantID {
			t.Errorf("mapReduceResults() got = %v, want %v", gotID, wantID)
		}
		if gotName != wantName {
			t.Errorf("mapReduceResults() got = %v, want %v", gotName, wantName)
		}
	})
}
