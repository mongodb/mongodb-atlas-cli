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


package store

import (
	"net/url"
	"strings"
	"errors"
	"go.mongodb.org/atlas-sdk/v20241113004/admin"
)

// isSafeURL checks if the URL is safe to use and not an internal service.
func isSafeURL(rawURL string) (bool, error) {
	// Parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return false, err
	}

	// Check if the URL is empty or uses an unsafe scheme (e.g., file://)
	if parsedURL.Scheme == "" || parsedURL.Scheme == "file" {
		return false, errors.New("invalid URL scheme")
	}

	// Check for internal addresses (localhost, 127.0.0.1, private IP ranges)
	if strings.HasPrefix(parsedURL.Host, "localhost") ||
		strings.HasPrefix(parsedURL.Host, "127.0.0.1") ||
		strings.HasPrefix(parsedURL.Host, "0.0.0.0") ||
		strings.HasPrefix(parsedURL.Host, "::1") ||
		// Allow custom ranges like 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16
		(strings.HasPrefix(parsedURL.Host, "10.") ||
			strings.HasPrefix(parsedURL.Host, "172.") ||
			strings.HasPrefix(parsedURL.Host, "192.168")) {
		return false, errors.New("unsafe internal URL")
	}

	// Optionally: Add a whitelist for allowed domains (e.g., only allow specific domains)
	allowedDomains := []string{"example.com", "mytrustedapi.com"}
	for _, domain := range allowedDomains {
		if strings.HasSuffix(parsedURL.Host, domain) {
			return true, nil
		}
	}

	// If the domain is not in the whitelist, block it
	return false, errors.New("domain not allowed")
}

// CreateIntegration encapsulates the logic to manage different cloud providers.
func (s *Store) CreateIntegration(projectID, integrationType string, integration *atlasv2.ThirdPartyIntegration) (*atlasv2.PaginatedIntegration, error) {
	// Validate the webhook URL before proceeding
	if integration.WebhookURL != "" {
		isValid, err := isSafeURL(integration.WebhookURL)
		if !isValid {
			return nil, err // Return the error if URL is unsafe
		}
	}

	// Proceed with creating the integration if URL is safe
	resp, _, err := s.clientv2.ThirdPartyIntegrationsApi.CreateThirdPartyIntegration(s.ctx,
		integrationType, projectID, integration).Execute()
	return resp, err
}
