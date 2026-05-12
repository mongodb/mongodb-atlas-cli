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

package pledge

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	shared_api "github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

const (
	approvalTTL     = 5 * time.Minute
	pollInterval    = 500 * time.Millisecond
	requestsSubdir  = "requests"
	grantsSubdir    = "grants"
)

// ApprovalRequest is written to <state>/<sessionKey>/requests/<token>.json by
// the blocked process. It is read by the approver running atlas pledge allow.
type ApprovalRequest struct {
	Token         string                   `json:"token"`
	SessionKeyStr string                   `json:"sessionKey"`
	OperationID   string                   `json:"operationID"`
	Tier          shared_api.PermissionTier `json:"tier"`
	ParamsHash    string                   `json:"paramsHash,omitempty"`
	CreatedAt     time.Time                `json:"createdAt"`
}

// ApprovalGrant is written to <state>/<sessionKey>/grants/<token>.json by the approver.
type ApprovalGrant struct {
	Token         string    `json:"token"`
	OperationID   string    `json:"operationID"`
	ApproverSID   int       `json:"approverSID"`
	GrantedAt     time.Time `json:"grantedAt"`
	ExpiresAt     time.Time `json:"expiresAt"`
	HMAC          string    `json:"hmac"`
}

var (
	ErrApprovalExpired   = errors.New("approval token has expired")
	ErrApprovalConsumed  = errors.New("approval token already consumed")
	ErrApprovalNotFound  = errors.New("approval token not found")
	ErrApprovalForged    = errors.New("approval grant signature invalid")
	ErrApproverUnderPledged = errors.New("approver's pledge does not permit this operation")
)

func sessionKeyDir(k SessionKey) (string, error) {
	dir, err := StateDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, k.String()), nil
}

// WriteApprovalRequest creates a request file and returns the token.
// The request is keyed on the SessionKey so both SID-based and Claude-based
// sessions can use the approval flow.
func WriteApprovalRequest(k SessionKey, opID string, tier shared_api.PermissionTier) (string, error) {
	raw := make([]byte, 16)
	if _, err := rand.Read(raw); err != nil {
		return "", fmt.Errorf("generating token: %w", err)
	}
	token := hex.EncodeToString(raw)

	base, err := sessionKeyDir(k)
	if err != nil {
		return "", err
	}
	reqDir := filepath.Join(base, requestsSubdir)
	if err := ensureDir(reqDir); err != nil {
		return "", err
	}

	req := ApprovalRequest{
		Token:         token,
		SessionKeyStr: k.String(),
		OperationID:   opID,
		Tier:          tier,
		CreatedAt:     time.Now().UTC(),
	}
	data, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	path := filepath.Join(reqDir, token+".json")
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return "", fmt.Errorf("writing request file: %w", err)
	}
	return token, nil
}

// Approve reads a request file and, if the approver's pledge permits it, writes
// a signed grant file.
func Approve(token string, approverPledge *PledgeFile, approverSID int) error {
	// Find the request by scanning all SID dirs (token is globally unique by hex entropy).
	stateDir, err := StateDir()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(stateDir)
	if err != nil {
		return ErrApprovalNotFound
	}

	var req *ApprovalRequest
	var reqPath string
	var grantDir string

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		candidate := filepath.Join(stateDir, e.Name(), requestsSubdir, token+".json")
		data, readErr := os.ReadFile(candidate)
		if readErr != nil {
			continue
		}
		var r ApprovalRequest
		if err := json.Unmarshal(data, &r); err != nil {
			continue
		}
		req = &r
		reqPath = candidate
		grantDir = filepath.Join(stateDir, e.Name(), grantsSubdir)
		break
	}

	if req == nil {
		return ErrApprovalNotFound
	}

	// Check expiry.
	if time.Since(req.CreatedAt) > approvalTTL {
		return ErrApprovalExpired
	}

	// Check that approver's pledge is at least as permissive as requested tier.
	if approverPledge != nil {
		outcome, _ := Check(approverPledge, req.Tier, req.OperationID)
		if outcome == Block {
			return ErrApproverUnderPledged
		}
	}

	// Check for already-consumed grant.
	consumedPath := filepath.Join(grantDir, token+".consumed")
	if _, err := os.Stat(consumedPath); err == nil {
		return ErrApprovalConsumed
	}

	if err := ensureDir(grantDir); err != nil {
		return err
	}

	grant := ApprovalGrant{
		Token:       token,
		OperationID: req.OperationID,
		ApproverSID: approverSID,
		GrantedAt:   time.Now().UTC(),
		ExpiresAt:   time.Now().UTC().Add(approvalTTL),
	}

	// Sign the grant.
	grantData, err := json.Marshal(grant)
	if err != nil {
		return err
	}
	sig, err := SignHMAC(grantData)
	if err != nil {
		return err
	}
	grant.HMAC = sig

	finalData, err := json.Marshal(grant)
	if err != nil {
		return err
	}

	grantPath := filepath.Join(grantDir, token+".json")
	if err := atomicWrite(grantPath, finalData, 0o600); err != nil {
		return fmt.Errorf("writing grant file: %w", err)
	}

	// Clean up the request file after grant is written.
	_ = os.Remove(reqPath)
	return nil
}

// WaitForGrant polls for a grant file for the given session key and token.
// Returns the grant if one appears within the context deadline.
func WaitForGrant(ctx context.Context, k SessionKey, token string) (*ApprovalGrant, error) {
	base, err := sessionKeyDir(k)
	if err != nil {
		return nil, err
	}
	grantPath := filepath.Join(base, grantsSubdir, token+".json")
	consumedPath := filepath.Join(base, grantsSubdir, token+".consumed")

	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			if _, err := os.Stat(consumedPath); err == nil {
				return nil, ErrApprovalConsumed
			}
			data, err := os.ReadFile(grantPath)
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			if err != nil {
				return nil, err
			}
			var grant ApprovalGrant
			if err := json.Unmarshal(data, &grant); err != nil {
				continue
			}
			return &grant, nil
		}
	}
}

// ConsumeGrant verifies the HMAC of a grant and renames it to .consumed.
// Returns ErrApprovalExpired, ErrApprovalForged, or ErrApprovalConsumed on failure.
func ConsumeGrant(k SessionKey, token string, expectedOpID string) error {
	base, err := sessionKeyDir(k)
	if err != nil {
		return err
	}
	grantPath := filepath.Join(base, grantsSubdir, token+".json")
	consumedPath := filepath.Join(base, grantsSubdir, token+".consumed")

	if _, err := os.Stat(consumedPath); err == nil {
		return ErrApprovalConsumed
	}

	data, err := os.ReadFile(grantPath)
	if errors.Is(err, os.ErrNotExist) {
		return ErrApprovalNotFound
	}
	if err != nil {
		return err
	}

	var grant ApprovalGrant
	if err := json.Unmarshal(data, &grant); err != nil {
		return fmt.Errorf("malformed grant: %w", err)
	}

	if time.Now().After(grant.ExpiresAt) {
		return ErrApprovalExpired
	}

	// Verify HMAC: re-sign the grant without the HMAC field.
	savedHMAC := grant.HMAC
	grant.HMAC = ""
	checkData, err := json.Marshal(grant)
	if err != nil {
		return err
	}
	ok, err := VerifyHMAC(checkData, savedHMAC)
	if err != nil || !ok {
		return ErrApprovalForged
	}

	if grant.OperationID != expectedOpID {
		return ErrApprovalForged
	}

	// Atomically consume: rename to .consumed.
	if err := os.Rename(grantPath, consumedPath); err != nil {
		return fmt.Errorf("consuming grant: %w", err)
	}
	return nil
}
