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

// Package pledge implements session-scoped permission restriction for atlas-cli.
// A pledge is a voluntary, monotonically-narrowing allowlist of operations that
// the CLI is permitted to perform. State is anchored to the POSIX session ID so
// all atlas invocations within the same shell session share the same pledge.
package pledge

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/tools/shared/api"
)

// Profile names the built-in pledge profiles.
type Profile string

const (
	ProfileReadonly  Profile = "readonly"
	ProfileReadWrite Profile = "read-write"
	ProfileAdmin     Profile = "admin"
)

// tierOrder maps profiles to an integer; lower means more restricted.
// A session can only move to a lower (or equal) number — not higher.
var tierOrder = map[api.PermissionTier]int{
	api.PermissionRead:       0,
	api.PermissionWrite:      1,
	api.PermissionLocalWrite: 1,
	api.PermissionAdmin:      2,
}

// profileMaxTier maps profiles to the highest tier they permit.
var profileMaxTier = map[Profile]api.PermissionTier{
	ProfileReadonly:  api.PermissionRead,
	ProfileReadWrite: api.PermissionWrite,
	ProfileAdmin:     api.PermissionAdmin,
}

// Errors.
var (
	ErrNoPledge   = errors.New("no pledge active for this session")
	ErrWouldWiden = errors.New("pledge can only be narrowed within a session; open a new terminal to reset")
)

// PledgeFile is the on-disk representation of a session pledge.
type PledgeFile struct {
	Version    int                   `json:"version"`
	Profile    Profile               `json:"profile"`
	AllowedOps []string              `json:"allowedOps,omitempty"`
	MaxTier    api.PermissionTier    `json:"maxTier"`
	NarrowedAt time.Time             `json:"narrowedAt"`
}

// sidPath returns the path to the pledge file for the given SID.
func sidPath(stateDir string, sid int) string {
	return filepath.Join(stateDir, fmt.Sprintf("%d.json", sid))
}

// Load reads the pledge file for the given SID. Returns ErrNoPledge if none exists.
func Load(sid int) (*PledgeFile, error) {
	dir, err := StateDir()
	if err != nil {
		return nil, err
	}
	return loadFrom(sidPath(dir, sid))
}

func loadFrom(path string) (*PledgeFile, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, ErrNoPledge
	}
	if err != nil {
		return nil, fmt.Errorf("reading pledge file: %w", err)
	}
	var pf PledgeFile
	if err := json.Unmarshal(data, &pf); err != nil {
		return nil, fmt.Errorf("parsing pledge file: %w", err)
	}
	return &pf, nil
}

// Save persists the pledge for the given SID, mode 0600.
func Save(sid int, pf *PledgeFile) error {
	dir, err := StateDir()
	if err != nil {
		return err
	}
	if err := ensureDir(dir); err != nil {
		return fmt.Errorf("creating state directory: %w", err)
	}
	gcStalePledgeFiles(dir, sid)

	data, err := json.MarshalIndent(pf, "", "  ")
	if err != nil {
		return err
	}
	path := sidPath(dir, sid)
	return atomicWrite(path, data, 0o600)
}

// Widen saves pf as the new pledge for sid without enforcing the narrowing
// constraint. Only call this after obtaining explicit interactive consent.
func Widen(sid int, next *PledgeFile) error {
	next.NarrowedAt = time.Now().UTC()
	if err := Save(sid, next); err != nil {
		return err
	}
	LogAudit(AuditEntry{
		SID:         sid,
		OperationID: "<pledge-widened>",
		Outcome:     AuditOutcome(fmt.Sprintf("widened-to:%s", next.MaxTier)),
	})
	return nil
}

// Narrow applies next as the new pledge for sid, returning ErrWouldWiden if next
// would be more permissive than the current pledge.
func Narrow(sid int, next *PledgeFile) error {
	current, err := Load(sid)
	if err != nil && !errors.Is(err, ErrNoPledge) {
		return err
	}
	if current != nil {
		// Only allow narrowing.
		if tierOrder[next.MaxTier] > tierOrder[current.MaxTier] {
			return ErrWouldWiden
		}
	}
	next.Version = 1
	if current != nil {
		next.Version = current.Version + 1
	}
	next.NarrowedAt = time.Now().UTC()
	return Save(sid, next)
}

// Outcome is the result of a pledge check.
type Outcome int

const (
	Allow Outcome = iota
	Block
)

// BlockedError carries enough context for the error message and audit log.
type BlockedError struct {
	OperationID string
	Required    api.PermissionTier
	MaxAllowed  api.PermissionTier
	Profile     Profile
}

func (e *BlockedError) Error() string {
	return fmt.Sprintf(
		"atlas pledge [%s]: operation %q requires %q but session is restricted to %q\n"+
			"Run 'atlas pledge allow <token>' in another terminal to approve this specific operation.",
		e.Profile, e.OperationID, e.Required, e.MaxAllowed,
	)
}

// Check returns Block+error if pf forbids a command with the given permission tier and operationID.
// operationID is used only for the error message; enforcement is tier-based.
func Check(pf *PledgeFile, required api.PermissionTier, operationID string) (Outcome, error) {
	if pf == nil {
		return Allow, nil
	}
	// If operationID is in the AllowedOps list, permit regardless of tier.
	for _, op := range pf.AllowedOps {
		if op == operationID {
			return Allow, nil
		}
	}
	if tierOrder[required] > tierOrder[pf.MaxTier] {
		return Block, &BlockedError{
			OperationID: operationID,
			Required:    required,
			MaxAllowed:  pf.MaxTier,
			Profile:     pf.Profile,
		}
	}
	return Allow, nil
}

// NewPledgeFile builds a PledgeFile from a named profile plus optional extra allowed ops.
func NewPledgeFile(profile Profile, allowedOps []string) (*PledgeFile, error) {
	maxTier, ok := profileMaxTier[profile]
	if !ok {
		return nil, fmt.Errorf("unknown pledge profile %q; valid values: readonly, read-write, admin", profile)
	}
	return &PledgeFile{
		Profile:    profile,
		MaxTier:    maxTier,
		AllowedOps: allowedOps,
	}, nil
}

// gcStalePledgeFiles lazily removes <sid>.json files whose SID no longer has
// a running process. It is best-effort and ignores errors.
func gcStalePledgeFiles(dir string, currentSID int) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		var sid int
		if _, err := fmt.Sscanf(e.Name(), "%d.json", &sid); err != nil {
			continue
		}
		if sid == currentSID {
			continue
		}
		if !sidHasLiveProcess(sid) {
			_ = os.Remove(filepath.Join(dir, e.Name()))
		}
	}
}
