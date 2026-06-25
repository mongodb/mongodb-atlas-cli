// Copyright 2026 MongoDB Inc
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

package auth

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/atlas/auth"
	"go.uber.org/mock/gomock"
)

func TestMetadataExpired(t *testing.T) {
	tests := []struct {
		name   string
		cached map[string]any
		want   bool
	}{
		{
			name:   "future expiry is not expired",
			cached: map[string]any{"expiry": time.Now().Add(time.Hour).Format(time.RFC3339)},
			want:   false,
		},
		{
			name:   "past expiry is expired",
			cached: map[string]any{"expiry": time.Now().Add(-time.Hour).Format(time.RFC3339)},
			want:   true,
		},
		{
			name:   "missing expiry key is expired",
			cached: map[string]any{},
			want:   true,
		},
		{
			name:   "invalid expiry format is expired",
			cached: map[string]any{"expiry": "not-a-timestamp"},
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, metadataExpired(tt.cached))
		})
	}
}

func TestDiscoverOrLoadMetadata(t *testing.T) {
	// Stand up a discovery server so we can verify "fetched" vs "cached" by
	// counting hits.
	var hits int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		hits++
		w.Header().Set("Cache-Control", "max-age=3600")
		fmt.Fprintf(w, `{
			"issuer": %q,
			"authorization_endpoint": %q,
			"token_endpoint": %q
		}`, "https://issuer.example.com", "https://issuer.example.com/authorize", "https://issuer.example.com/token")
	}))
	defer server.Close()

	authCfg := auth.NewConfig(nil)
	authCfg.AuthServerURL, _ = url.Parse(server.URL)
	// authCfg.AuthServerURL is what the cache compares the cached issuer against;
	// for cache-hit cases the cached metadata must declare this as its issuer.
	cachedIssuer := authCfg.AuthServerURL.String()

	freshCache := map[string]any{
		"metadata": map[string]any{"issuer": cachedIssuer, "authorization_endpoint": "a", "token_endpoint": "t"},
		"expiry":   time.Now().Add(time.Hour).Format(time.RFC3339),
	}
	staleCache := map[string]any{
		"metadata": map[string]any{"issuer": cachedIssuer},
		"expiry":   time.Now().Add(-time.Hour).Format(time.RFC3339),
	}
	wrongIssuerCache := map[string]any{
		"metadata": map[string]any{"issuer": "https://other.example.com"},
		"expiry":   time.Now().Add(time.Hour).Format(time.RFC3339),
	}

	tests := []struct {
		name         string
		discover     bool
		cached       map[string]any
		expectFetch  bool
		expectIssuer string
	}{
		{name: "fresh cache, matching issuer", cached: freshCache, expectFetch: false, expectIssuer: cachedIssuer},
		{name: "stale cache forces refetch", cached: staleCache, expectFetch: true, expectIssuer: "https://issuer.example.com"},
		{name: "issuer mismatch forces refetch", cached: wrongIssuerCache, expectFetch: true, expectIssuer: "https://issuer.example.com"},
		{name: "no cache forces refetch", cached: nil, expectFetch: true, expectIssuer: "https://issuer.example.com"},
		{name: "discover flag clears cache and refetches", discover: true, cached: freshCache, expectFetch: true, expectIssuer: "https://issuer.example.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockConfig := NewMockUserDelegationConfig(ctrl)

			if tt.discover {
				// Discover flag clears the cache first; subsequent reads see nil.
				mockConfig.EXPECT().SetAuthServerMetadata(nil).Times(1)
				mockConfig.EXPECT().Save().Return(nil).Times(1)
				mockConfig.EXPECT().AuthServerMetadata().Return(nil).Times(1)
			} else {
				mockConfig.EXPECT().AuthServerMetadata().Return(tt.cached).Times(1)
			}
			if tt.expectFetch {
				mockConfig.EXPECT().SetAuthServerMetadata(gomock.Any()).Times(1)
				mockConfig.EXPECT().Save().Return(nil).Times(1)
			}

			opts := &UserDelegationFlow{config: mockConfig, Discover: tt.discover}

			startHits := hits
			metadata, err := opts.discoverOrLoadMetadata(t.Context(), authCfg)
			require.NoError(t, err)
			assert.Equal(t, tt.expectIssuer, metadata["issuer"])

			fetched := hits > startHits
			assert.Equal(t, tt.expectFetch, fetched, "expectFetch=%v fetched=%v", tt.expectFetch, fetched)
		})
	}
}
