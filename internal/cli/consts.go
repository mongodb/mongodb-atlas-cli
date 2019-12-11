package cli

import (
	"runtime"

	"github.com/10gen/mcli/internal/version"
)

const (
	// DefaultUserAgent to be submitted by the client
	DefaultUserAgent = Name + "/" + version.Version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
	// CloudDefaultURL Atlas default URL
	CloudDefaultURL = "https://cloud-qa.mongodb.com/api/atlas/v1.0/"
	Name            = "mcli"
	// CLoudService setting when using Atlas API
	CLoudService = "cloud"
	// CloudManagerService settings when using CLoud Manager API
	CloudManagerService = "cloud-manager"
	// OpsManagerService settings when using Ops Manager API
	OpsManagerService = "ops-manager"
)
