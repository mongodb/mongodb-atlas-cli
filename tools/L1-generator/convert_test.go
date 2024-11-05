package main

import (
	"testing"
)

func TestExtractVersionAndContentType(t *testing.T) {
	tests := []struct {
		input       string
		version     string
		contentType string
	}{
		{
			"application/vnd.atlas.2025-01-01+json",
			"2025-01-01",
			"json",
		},
		{
			"application/vnd.atlas.2024-08-05+json",
			"2024-08-05",
			"json",
		},
		{
			"application/vnd.atlas.2023-01-01+csv",
			"2023-01-01",
			"csv",
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			version, contentType, err := extractVersionAndContentType(tt.input)
			if err != nil {
				t.Errorf("Error = %v", err)
			}

			if version != tt.version || contentType != tt.contentType {
				t.Errorf("Expected version=%s, contentType=%s. Got: version=%s, contentType=%s", tt.version, tt.contentType, version, contentType)
			}
		})
	}
}
