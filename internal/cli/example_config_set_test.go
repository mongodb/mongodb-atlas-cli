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

import (
	"log"
)

func ExampleConfigSetBuilder_projectID() {
	cmd := ConfigSetBuilder()
	cmd.SetArgs([]string{"project_id", "1"})
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// Updated prop 'project_id'
}

func ExampleConfigSetBuilder_orgID() {
	cmd := ConfigSetBuilder()
	cmd.SetArgs([]string{"org_id", "1"})
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// Updated prop 'org_id'
}

func ExampleConfigSetBuilder_privateAPIKey() {
	cmd := ConfigSetBuilder()
	cmd.SetArgs([]string{"private_api_key", "1"})
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// Updated prop 'private_api_key'
}

func ExampleConfigSetBuilder_publicAPIKey() {
	cmd := ConfigSetBuilder()
	cmd.SetArgs([]string{"public_api_key", "1"})
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// Updated prop 'public_api_key'
}

func ExampleConfigSetBuilder_service() {
	cmd := ConfigSetBuilder()
	cmd.SetArgs([]string{"service", "1"})
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// Updated prop 'service'
}

func ExampleConfigSetBuilder_opsManagerURL() {
	cmd := ConfigSetBuilder()
	cmd.SetArgs([]string{"ops_manager_url", "1"})
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	// Updated prop 'ops_manager_url'
}
