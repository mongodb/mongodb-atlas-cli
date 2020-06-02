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

package cli

import "testing"

func TestGlobalOpts_PreRunE(t *testing.T) {
	t.Run("empty project ID", func(t *testing.T) {
		o := &GlobalOpts{}
		if err := o.PreRunE(); err != errMissingProjectID {
			t.Errorf("Expected err: %#v, got: %#v\n", errMissingProjectID, err)
		}
	})
	t.Run("invalid project ID", func(t *testing.T) {
		o := &GlobalOpts{ProjectID: "1"}
		if err := o.PreRunE(); err == nil {
			t.Errorf("Expected an error\n")
		}
	})
	t.Run("valid project ID", func(t *testing.T) {
		o := &GlobalOpts{ProjectID: "5e98249d937cfc52efdc2a9f"}
		if err := o.PreRunE(); err != nil {
			t.Fatalf("PreRunE() unexpected error %v\n", err)
		}
	})
}
