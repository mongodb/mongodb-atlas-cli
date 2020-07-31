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

// +build unit

package output

import "testing"

type testConfig string

func (t testConfig) Output() string {
	return string(t)
}

func Test_templateValue(t *testing.T) {
	type args struct {
		c               Config
		defaultTemplate string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "go-template",
			args:    args{c: testConfig("go-template=test"), defaultTemplate: ""},
			want:    "test",
			wantErr: false,
		},
		{
			name:    "not-valid",
			args:    args{c: testConfig("not-valid"), defaultTemplate: "default"},
			want:    "default",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		args := tt.args
		wantErr := tt.wantErr
		want := tt.want
		t.Run(tt.name, func(t *testing.T) {
			got, err := templateValue(args.c, args.defaultTemplate)
			if (err != nil) != wantErr {
				t.Errorf("templateValue() error = %v, wantErr %v", err, wantErr)
				return
			}
			if got != want {
				t.Errorf("templateValue() got = %v, want %v", got, want)
			}
		})
	}
}
