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
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mongodb/mongodb-atlas-cli/test/e2e"
	atlasv2 "go.mongodb.org/atlas-sdk/v20230201004/admin"
)

const (
	authorizedEmail = "firstname.lastname@example.com"
)

// If we watch a command in a testing environment,
// the output has some dots in the beginning (depending on how long it took to finish) that need to be removed.
func removeDotsFromWatching(consoleOutput []byte) []byte {
	return []byte(strings.TrimLeft(string(consoleOutput), "."))
}

func enableCompliancePolicy(projectID string) (*atlasv2.DataProtectionSettings, error) {
	cliPath, err := e2e.AtlasCLIBin()
	if err != nil {
		return nil, fmt.Errorf("%w: invalid bin", err)
	}
	cmd := exec.Command(cliPath,
		backupsEntity,
		compliancepolicyEntity,
		"enable",
		"--projectId",
		projectID,
		"--authorizedEmail",
		authorizedEmail,
		"-o=json",
		"--watch",
	)
	cmd.Env = os.Environ()
	resp, outputErr := cmd.CombinedOutput()
	if outputErr != nil {
		return nil, outputErr
	}
	trimmedResponse := removeDotsFromWatching(resp)

	var result atlasv2.DataProtectionSettings
	if err := json.Unmarshal(trimmedResponse, &result); err != nil {
		fmt.Printf("%+v", trimmedResponse)
		return nil, err
	}
	return &result, nil
}
