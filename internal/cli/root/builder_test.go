// Copyright 2021 MongoDB Inc
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

package root

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mongodb/mongocli/internal/mocks"
	"github.com/mongodb/mongocli/internal/version"
)

func TestBuilder(t *testing.T) {
	type args struct {
		argsWithoutProg []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "atlas",
			want: 5,
			args: args{
				argsWithoutProg: []string{"atlas"},
			},
		},
		{
			name: "ops-manager",
			want: 4,
			args: args{
				argsWithoutProg: []string{"ops-manager"},
			},
		},
		{
			name: "cloud-manager",
			want: 4,
			args: args{
				argsWithoutProg: []string{"cloud-manager"},
			},
		},
		{
			name: "ops-manager alias",
			want: 4,
			args: args{
				argsWithoutProg: []string{"om"},
			},
		},
		{
			name: "cloud-manager alias",
			want: 4,
			args: args{
				argsWithoutProg: []string{"cm"},
			},
		},
		{
			name: "iam",
			want: 4,
			args: args{
				argsWithoutProg: []string{"iam"},
			},
		},
		{
			name: "empty",
			want: 5,
			args: args{
				argsWithoutProg: []string{},
			},
		},
		{
			name: "autocomplete",
			want: 5,
			args: args{
				argsWithoutProg: []string{"__complete"},
			},
		},
		{
			name: "completion",
			want: 5,
			args: args{
				argsWithoutProg: []string{"completion"},
			},
		},
		{
			name: "--version",
			want: 5,
			args: args{
				argsWithoutProg: []string{"completion"},
			},
		},
	}
	var profile string
	for _, tt := range tests {
		name := tt.name
		args := tt.args
		want := tt.want
		t.Run(name, func(t *testing.T) {
			got := Builder(&profile, args.argsWithoutProg)
			if len(got.Commands()) != want {
				t.Fatalf("got=%d, want=%d", len(got.Commands()), want)
			}
		})
	}
}

func TestOutputOpts_printNewVersionAvailable(t *testing.T) {
	prevVersion := version.Version
	version.Version = "v1.0.0"
	defer func() {
		version.Version = prevVersion
	}()

	ctrl := gomock.NewController(t)
	mockStore := mocks.NewMockVersionDescriber(ctrl)
	defer ctrl.Finish()

	mockStore.
		EXPECT().
		LatestVersion().
		Return("v2.0.0", nil).
		Times(1)

	bufOut := new(bytes.Buffer)
	opts := &BuilderOpts{
		store: mockStore,
	}
	err := opts.printNewVersionAvailable(bufOut)
	if err != nil {
		t.Errorf("printNewVersionAvailable() unexpected error: %v", err)
	}

	if got, want := bufOut.String(), `
A new version of mongocli is available 'v2.0.0'!
To upgrade, see: https://dochub.mongodb.org/core/mongocli-install.

To disable this alert, run "mongocli config set skip_update_check true".
`; got != want {
		t.Errorf("printNewVersionAvailable() got = %v, want %v", got, want)
	}
}
