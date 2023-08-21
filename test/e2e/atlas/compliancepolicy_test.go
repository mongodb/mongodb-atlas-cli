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

//go:build e2e || (atlas && backup && compliancepolicy)

package atlas_test

import (
	"strings"
	"testing"
)

const (
	authorizedEmail = "firstname.lastname@example.com"
)

func TestCompliancePolicy(t *testing.T) {
	g := newAtlasE2ETestGenerator(t)
	g.generateProject("compliancePolicy")

	testCompliancePolicySetup(t, g.projectID)
	testDescribe(t, g)
	testPoliciesDescribe(t, g.projectID)
	testCopyProtection(t, g.projectID)
}

// If we watch a command in a testing environment,
// the output has some dots in the beginning (depending on how long it took to finish) that need to be removed.
func removeDotsFromWatching(consoleOutput []byte) []byte {
	return []byte(strings.TrimLeft(string(consoleOutput), "."))
}
