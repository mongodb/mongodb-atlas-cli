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

	testCompliancePolicySetup(t, g)
	testDescribe(t, g)
	testPoliciesDescribe(t, g)
	testCopyProtection(t, g)
}

// For tests that update BCP, we must --watch to avoid HTTP 400 Bad Request "CANNOT_UPDATE_BACKUP_COMPLIANCE_POLICY_SETTINGS_WITH_PENDING_ACTION".
// Because we watch the command and this is a testing environment,
// the resp output has some dots in the beginning (depending on how long it took to finish) that need to be removed.
// It looks something like this:
//
// "...{"projectId": "string", ...}".
func removeDotsFromWatching(consoleOutput []byte) []byte {
	return []byte(strings.TrimLeft(string(consoleOutput), "."))
}
