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

// withTempStateDir sets ATLAS_PLEDGE_STATE_DIR to a fresh temp dir for the test.
func withTempStateDir(t *testing.T) {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("ATLAS_PLEDGE_STATE_DIR", dir)
}

func TestLoadSave(t *testing.T) {
	withTempStateDir(t)
	pf, err := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := pledge.Save(1234, pf); err != nil {
		t.Fatal(err)
	}
	loaded, err := pledge.Load(1234)
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
	_, err := pledge.Load(9999)
	if !errors.Is(err, pledge.ErrNoPledge) {
		t.Errorf("expected ErrNoPledge, got %v", err)
	}
}

func TestNarrowAllowed(t *testing.T) {
	withTempStateDir(t)
	// Start with read-write, narrow to readonly.
	rw, _ := pledge.NewPledgeFile(pledge.ProfileReadWrite, nil)
	if err := pledge.Narrow(1234, rw); err != nil {
		t.Fatal(err)
	}
	ro, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Narrow(1234, ro); err != nil {
		t.Errorf("narrowing from read-write to readonly should succeed, got: %v", err)
	}
}

func TestNarrowWideningRejected(t *testing.T) {
	withTempStateDir(t)
	ro, _ := pledge.NewPledgeFile(pledge.ProfileReadonly, nil)
	if err := pledge.Narrow(1234, ro); err != nil {
		t.Fatal(err)
	}
	rw, _ := pledge.NewPledgeFile(pledge.ProfileReadWrite, nil)
	err := pledge.Narrow(1234, rw)
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
		SID:         1234,
		OperationID: "deleteCluster",
		Outcome:     pledge.AuditBlocked,
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
