package config

const (
	// Name of the CLI
	Name           = "mcli"
	DefaultProfile = "default"
	// CloudService setting when using Atlas API
	CloudService = "cloud"
	// CloudManagerService settings when using CLoud Manager API
	CloudManagerService = "cloud-manager"
	// OpsManagerService settings when using Ops Manager API
	OpsManagerService = "ops-manager"
	configType        = "toml"
	service           = "service"
	publicAPIKey      = "public_api_key"
	privateAPIKey     = "private_api_key"
	opsManagerURL     = "ops_manager_url"
	projectID         = "project_id"
	baseURL           = "base_url"
	publicAPIPath     = "/api/public/v1.0/"
	atlasAPIPath      = "/api/atlas/v1.0/"
)
