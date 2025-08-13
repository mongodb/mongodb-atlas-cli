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
	"errors"
	"testing"
)

func TestOrgOpts_ValidateOrgID(t *testing.T) {
	t.Run("empty org ID", func(t *testing.T) {
		o := &OrgOpts{}
		if err := o.ValidateOrgID(); !errors.Is(err, ErrMissingOrgID) {
			t.Errorf("Expected err: %#v, got: %#v\n", ErrMissingOrgID, err)
		}
	})
	t.Run("invalid org ID", func(t *testing.T) {
		o := &OrgOpts{OrgID: "1"}
		if err := o.ValidateOrgID(); err == nil {
			t.Errorf("Expected an error\n")
		}
	})
	t.Run("valid org ID", func(t *testing.T) {
		o := &OrgOpts{OrgID: "5e98249d937cfc52efdc2a9f"}
		if err := o.ValidateOrgID(); err != nil {
			t.Fatalf("PreRunE() unexpected error %v\n", err)
		}
	})
}

func TestOrgOpts_PreRunE(t *testing.T) {
	t.Run("no error", func(t *testing.T) {
		noErrorFunc := func() error {
			return nil
		}

		o := &OrgOpts{}
		if err := o.PreRunE(noErrorFunc); err != nil {
			t.Errorf("Expected err == nil")
		}
	})

	t.Run("error", func(t *testing.T) {
		errorFunc := func() error {
			return errors.New("error")
		}

		o := &OrgOpts{}
		if err := o.PreRunE(errorFunc); err == nil {
			t.Errorf("Expected err != nil")
		}
	})
}
