// Copyright 2022 MongoDB Inc
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

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidNamespace(t *testing.T) {
	tests := []struct {
		name            string
		targetNamespace string
		expectedErr     string
	}{
		{
			name:            "Valid Namespace",
			targetNamespace: "valid-namespace",
			expectedErr:     "",
		},
		{
			name:            "Valid Empty Namespace",
			targetNamespace: "",
			expectedErr:     "",
		},
		{
			name:            "Invalid Namespace with special characters",
			targetNamespace: "invalid_namespace!",
			expectedErr:     "targetNamespace parameter is invalid: [a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')]",
		},
		{
			name:            "Invalid Namespace starting with non-alphanumeric",
			targetNamespace: "-invalidnamespace",
			expectedErr:     "targetNamespace parameter is invalid: [a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')]",
		},
		{
			name:            "Invalid Namespace ending with non-alphanumeric",
			targetNamespace: "invalidnamespace-",
			expectedErr:     "targetNamespace parameter is invalid: [a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')]",
		},
		{
			name:            "Invalid Namespace with uppercase letters",
			targetNamespace: "InvalidNamespace",
			expectedErr:     "targetNamespace parameter is invalid: [a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := &GenerateOpts{
				targetNamespace: tt.targetNamespace,
			}
			err := opts.ValidateTargetNamespace()

			if tt.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.expectedErr)
			}
		})
	}
}
