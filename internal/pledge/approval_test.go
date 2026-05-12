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

package pledge_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

var testKey, _ = pledge.NewSessionKey("sid", "42")

// TestApprovalHappyPath: write request, approve it, consume it.
func TestApprovalHappyPath(t *testing.T) {
	withTempStateDir(t)

	token, err := pledge.WriteApprovalRequest(testKey, "deleteCluster", api.PermissionWrite)
	if err != nil {
		t.Fatal(err)
	}

	// Approve with no pledge (approver has full access).
	if err := pledge.Approve(token, nil, 43); err != nil {
		t.Fatalf("Approve failed: %v", err)
	}

	// Wait returns immediately since the grant was written.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grant, err := pledge.WaitForGrant(ctx, testKey, token)
	if err != nil {
		t.Fatalf("WaitForGrant: %v", err)
	}
	if grant.OperationID != "deleteCluster" {
		t.Errorf("grant has wrong operationID: %q", grant.OperationID)
	}

	// Consume.
	if err := pledge.ConsumeGrant(testKey, token, "deleteCluster"); err != nil {
		t.Fatalf("ConsumeGrant: %v", err)
	}

	// Double-consume must fail.
	if err := pledge.ConsumeGrant(testKey, token, "deleteCluster"); !errors.Is(err, pledge.ErrApprovalConsumed) {
		t.Errorf("expected ErrApprovalConsumed on double-consume, got: %v", err)
	}
}

// TestApprovalNotFound: consuming a non-existent token returns ErrApprovalNotFound.
func TestApprovalNotFound(t *testing.T) {
	withTempStateDir(t)
	if err := pledge.ConsumeGrant(testKey, "deadbeef", "someOp"); !errors.Is(err, pledge.ErrApprovalNotFound) {
		t.Errorf("expected ErrApprovalNotFound, got: %v", err)
	}
}

// TestApprovalTokenNotFound: approving a missing token returns ErrApprovalNotFound.
func TestApprovalTokenNotFound(t *testing.T) {
	withTempStateDir(t)
	if err := pledge.Approve("doesnotexist", nil, 0); !errors.Is(err, pledge.ErrApprovalNotFound) {
		t.Errorf("expected ErrApprovalNotFound, got: %v", err)
	}
}

// TestApprovalApproverUnderPledged: approver's pledge does not permit the tier.
func TestApprovalApproverUnderPledged(t *testing.T) {
	withTempStateDir(t)

	token, err := pledge.WriteApprovalRequest(testKey, "deleteOrg", api.PermissionAdmin)
	if err != nil {
		t.Fatal(err)
	}

	// Approver only has readonly.
	roPledge, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Approve(token, roPledge, 43); !errors.Is(err, pledge.ErrApproverUnderPledged) {
		t.Errorf("expected ErrApproverUnderPledged, got: %v", err)
	}
}

// TestApprovalForgedGrant: a grant with a tampered HMAC is rejected.
func TestApprovalForgedGrant(t *testing.T) {
	withTempStateDir(t)

	token, err := pledge.WriteApprovalRequest(testKey, "listClusters", api.PermissionRead)
	if err != nil {
		t.Fatal(err)
	}
	if err := pledge.Approve(token, nil, 43); err != nil {
		t.Fatal(err)
	}

	// Tamper: try to consume with a different operationID.
	if err := pledge.ConsumeGrant(testKey, token, "deleteCluster"); !errors.Is(err, pledge.ErrApprovalForged) {
		t.Errorf("expected ErrApprovalForged on opID mismatch, got: %v", err)
	}
}

// TestApprovalHappyPath_ClaudeSession: same flow using a Claude UUID session key.
func TestApprovalHappyPath_ClaudeSession(t *testing.T) {
	withTempStateDir(t)
	claudeKey, err := pledge.NewSessionKey("claude", "47a415e1-c599-401f-9bac-65e71a7232b2")
	if err != nil {
		t.Fatal(err)
	}

	token, err := pledge.WriteApprovalRequest(claudeKey, "deleteCluster", api.PermissionWrite)
	if err != nil {
		t.Fatal(err)
	}

	if err := pledge.Approve(token, nil, 1); err != nil {
		t.Fatalf("Approve failed: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	grant, err := pledge.WaitForGrant(ctx, claudeKey, token)
	if err != nil {
		t.Fatalf("WaitForGrant: %v", err)
	}
	if grant.OperationID != "deleteCluster" {
		t.Errorf("grant has wrong operationID: %q", grant.OperationID)
	}

	if err := pledge.ConsumeGrant(claudeKey, token, "deleteCluster"); err != nil {
		t.Fatalf("ConsumeGrant: %v", err)
	}
}
