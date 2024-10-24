package config

import (
	"context"
	"sync"
)

// ConfigTokenCache uses CLI config for token storage
type ConfigTokenCache struct {
	mu          sync.Mutex
}

// NOTE: We are reusing *existing* AccessToken config. 
// That config in many places in CLI is used to determine if we dealing with legacy based Auth.
// In future implementation we likely want to separate legacy OAuth with new implementation by separate configuration property

func (s *ConfigTokenCache) RetrieveToken(ctx context.Context) (*string, error) {
	// Locking added to ensure no race condition happens with SaveToken
	s.mu.Lock()
	defer s.mu.Unlock()

	token := GetString("oauth_token_cache")
	return &token, nil
}

func (s *ConfigTokenCache) SaveToken(ctx context.Context, tkn string) error {
	// Locking added to ensure no race condition happens with RetrieveToken
	s.mu.Lock()
	defer s.mu.Unlock()
	Set("oauth_token_cache", tkn)
	return nil
}
