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

const (
	ToolName            = "mongocli"      // ToolName of the CLI
	EnvPrefix           = "mcli"          // Prefix for ENV variables
	DefaultProfile      = "default"       // DefaultProfile default
	CloudService        = "cloud"         // CloudService setting when using Atlas API
	CloudManagerService = "cloud-manager" // CloudManagerService settings when using CLoud Manager API
	OpsManagerService   = "ops-manager"   // OpsManagerService settings when using Ops Manager API
	projectID           = "project_id"
	orgID               = "org_id"
	configType          = "toml"
	service             = "service"
	publicAPIKey        = "public_api_key"
	privateAPIKey       = "private_api_key"
	opsManagerURL       = "ops_manager_url"
	baseURL             = "base_url"
	LabelKey            = "Infrastructure Tool"
	LabelValue          = "mongoCLI"
)
