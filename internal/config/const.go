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

package config

import "os"

const (
	ToolName                     = "mongocli"      // ToolName of the CLI
	EnvPrefix                    = "mcli"          // EnvPrefix prefix for ENV variables
	DefaultProfile               = "default"       // DefaultProfile default
	CloudService                 = "cloud"         // CloudService setting when using Atlas API
	CloudGovService              = "cloudgov"      // CloudGovService setting when using Atlas API for Government
	CloudManagerService          = "cloud-manager" // CloudManagerService settings when using CLoud Manager API
	OpsManagerService            = "ops-manager"   // OpsManagerService settings when using Ops Manager API
	JSON                         = "json"          // JSON output format as json
	projectID                    = "project_id"
	orgID                        = "org_id"
	mongoShellPath               = "mongosh_path"
	configType                   = "toml"
	service                      = "service"
	publicAPIKey                 = "public_api_key"
	privateAPIKey                = "private_api_key"
	opsManagerURL                = "ops_manager_url"
	baseURL                      = "base_url"
	opsManagerCACertificate      = "ops_manager_ca_certificate"
	opsManagerSkipVerify         = "ops_manager_skip_verify"
	opsManagerVersionManifestURL = "ops_manager_version_manifest_url"
	output                       = "output"
	CloudGovServiceURL           = "https://cloud.mongodbgov.com/"
	fileFlags                    = os.O_CREATE | os.O_TRUNC | os.O_WRONLY
	configPerm                   = 0600
)
