package cmd

import "runtime"

const (
	// Version for client
	Version = "0.0.1"
	// DefaultUserAgent to be submitted by the client
	DefaultUserAgent = toolName + "/" + Version + " (" + runtime.GOOS + "; " + runtime.GOARCH + ")"
	// CloudDefaultURL Atlas default URL
	CloudDefaultURL = "https://cloud-qa.mongodb.com/api/atlas/v1.0/"
	toolName        = "mcli"
	configType      = "toml"
)
