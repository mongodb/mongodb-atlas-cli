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

package podman

import (
	"reflect"
	"testing"
)

func TestRunContainerArgsExtraOptions(t *testing.T) {
	got := runContainerArgs(RunContainerOpts{
		ExtraOptions: []string{"--ulimit", "nofile=64000:64000"},
		Image:        "mongodb/mongodb-atlas-local:8",
	})
	want := []string{
		"run",
		"--name", "",
		"--hostname", "",
		"--ulimit", "nofile=64000:64000",
		"mongodb/mongodb-atlas-local:8",
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("runContainerArgs() = %#v; want %#v", got, want)
	}
}
