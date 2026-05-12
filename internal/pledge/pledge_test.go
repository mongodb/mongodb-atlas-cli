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
	"errors"
	"os"
	"testing"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/pledge"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

const testUUID = "47a415e1-c599-401f-9bac-65e71a7232b2"

// withTempStateDir sets ATLAS_PLEDGE_STATE_DIR to a fresh temp dir for the test.
func withTempStateDir(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("ATLAS_PLEDGE_STATE_DIR", dir)
}

func TestLoadSave(t *testing.T) {
	withTempStateDir(t)
	k, _ := pledge.NewSessionKey("sid", "1234")
	pf, err := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := pledge.Save(k, pf); err != nil {
		t.Fatal(err)
	}
	loaded, err := pledge.Load(k)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Profile != pledge.ProfileReadonly {
		t.Errorf("want profile %q, got %q", pledge.ProfileReadonly, loaded.Profile)
	}
	if loaded.MaxTier != api.PermissionRead {
		t.Errorf("want maxTier %q, got %q", api.PermissionRead, loaded.MaxTier)
	}
}

func TestLoadNoPledge(t *testing.T) {
	withTempStateDir(t)
	k, _ := pledge.NewSessionKey("sid", "9999")
	_, err := pledge.Load(k)
	if !errors.Is(err, pledge.ErrNoPledge) {
		t.Errorf("expected ErrNoPledge, got %v", err)
	}
}

func TestNarrowAllowed(t *testing.T) {
	withTempStateDir(t)
	k, _ := pledge.NewSessionKey("sid", "1234")
	rw, _ := pledge.NewPledgeFile(pledge.ProfileReadWrite, nil)
	if err := pledge.Narrow(k, rw); err != nil {
		t.Fatal(err)
	}
	ro, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Narrow(k, ro); err != nil {
		t.Errorf("narrowing from read-write to readonly should succeed, got: %v", err)
	}
}

func TestNarrowWideningRejected(t *testing.T) {
	withTempStateDir(t)
	k, _ := pledge.NewSessionKey("sid", "1234")
	ro, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Narrow(k, ro); err != nil {
		t.Fatal(err)
	}
	rw, _ := pledge.NewPledgeFile(pledge.ProfileReadWrite, nil)
	err := pledge.Narrow(k, rw)
	if !errors.Is(err, pledge.ErrWouldWiden) {
		t.Errorf("widening should return ErrWouldWiden, got: %v", err)
	}
}

func TestCheckAllow(t *testing.T) {
	pf, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	outcome, err := pledge.Check(pf, api.PermissionRead, "getCluster")
	if err != nil || outcome != pledge.Allow {
		t.Errorf("read op under readonly should be allowed, got outcome=%v err=%v", outcome, err)
	}
}

func TestCheckBlock(t *testing.T) {
	pf, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	outcome, err := pledge.Check(pf, api.PermissionWrite, "deleteCluster")
	if err == nil || outcome != pledge.Block {
		t.Errorf("write op under readonly should be blocked, got outcome=%v err=%v", outcome, err)
	}
	var be *pledge.BlockedError
	if !errors.As(err, &be) {
		t.Errorf("expected BlockedError, got %T", err)
	}
}

func TestCheckAllowedOpOverride(t *testing.T) {
	pf, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, []string{"deleteCluster"})
	outcome, err := pledge.Check(pf, api.PermissionWrite, "deleteCluster")
	if err != nil || outcome != pledge.Allow {
		t.Errorf("explicit allowed op should bypass tier check, got outcome=%v err=%v", outcome, err)
	}
}

func TestCheckNilPledge(t *testing.T) {
	outcome, err := pledge.Check(nil, api.PermissionAdmin, "createOrg")
	if err != nil || outcome != pledge.Allow {
		t.Errorf("nil pledge should allow everything, got outcome=%v err=%v", outcome, err)
	}
}

func TestAuditLog(t *testing.T) {
	withTempStateDir(t)
	dir := os.Getenv("ATLAS_PLEDGE_STATE_DIR")
	pledge.LogAudit(pledge.AuditEntry{
		SessionKeyStr: "sid-1234",
		OperationID:   "deleteCluster",
		Outcome:       pledge.AuditBlocked,
	})
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, e := range entries {
		if e.Name() == "audit.log" {
			found = true
		}
	}
	if !found {
		t.Error("audit.log not created")
	}
}

func TestHMACIdempotent(t *testing.T) {
	withTempStateDir(t)
	data := []byte("test payload")
	sig1, err := pledge.SignHMAC(data)
	if err != nil {
		t.Fatal(err)
	}
	sig2, err := pledge.SignHMAC(data)
	if err != nil {
		t.Fatal(err)
	}
	if sig1 != sig2 {
		t.Error("HMAC should be deterministic for the same key and data")
	}
	ok, err := pledge.VerifyHMAC(data, sig1)
	if err != nil || !ok {
		t.Errorf("VerifyHMAC failed: ok=%v err=%v", ok, err)
	}
}

func TestNewSessionKeyInvalidKind(t *testing.T) {
	_, err := pledge.NewSessionKey("unknown", "value")
	if err == nil {
		t.Error("expected error for unknown kind")
	}
}

func TestNewSessionKeyValidKinds(t *testing.T) {
	for _, kind := range []string{"sid", "claude"} {
		k, err := pledge.NewSessionKey(kind, "value")
		if err != nil {
			t.Errorf("kind %q should be valid, got: %v", kind, err)
		}
		if k.Kind != kind {
			t.Errorf("kind mismatch: want %q got %q", kind, k.Kind)
		}
	}
}

func TestSessionKeyString(t *testing.T) {
	k, _ := pledge.NewSessionKey("claude", testUUID)
	want := "claude-" + testUUID
	if got := k.String(); got != want {
		t.Errorf("String() = %q, want %q", got, want)
	}
}

func TestIsValidUUID(t *testing.T) {
	valid := []string{
		"47a415e1-c599-401f-9bac-65e71a7232b2",
		"550e8400-e29b-41d4-a716-446655440000",
	}
	invalid := []string{
		"not-a-uuid",
		"47a415e1-c599-401f-9bac",
		"",
	}
	for _, u := range valid {
		if !pledge.IsValidUUID(u) {
			t.Errorf("IsValidUUID(%q) = false, want true", u)
		}
	}
	for _, u := range invalid {
		if pledge.IsValidUUID(u) {
			t.Errorf("IsValidUUID(%q) = true, want false", u)
		}
	}
}

func TestClaudeKeyedSaveLoad(t *testing.T) {
	withTempStateDir(t)
	k, _ := pledge.NewSessionKey("claude", testUUID)
	pf, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Save(k, pf); err != nil {
		t.Fatal(err)
	}
	loaded, err := pledge.Load(k)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Profile != pledge.ProfileReadonly {
		t.Errorf("want profile %q, got %q", pledge.ProfileReadonly, loaded.Profile)
	}
	dir := os.Getenv("ATLAS_PLEDGE_STATE_DIR")
	if _, err := os.Stat(dir + "/claude-" + testUUID + ".json"); err != nil {
		t.Errorf("expected claude-keyed pledge file: %v", err)
	}
}
