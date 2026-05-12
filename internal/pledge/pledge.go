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
// the CLI is permitted to perform. State is anchored to a session key (POSIX SID
// or Claude Code conversation UUID) so all atlas invocations within the same
// session share the same pledge.
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
	Version    int                `json:"version"`
	Profile    Profile            `json:"profile"`
	AllowedOps []string           `json:"allowedOps,omitempty"`
	MaxTier    api.PermissionTier `json:"maxTier"`
	NarrowedAt time.Time          `json:"narrowedAt"`
}

// keyPath returns the path to the pledge file for the given SessionKey.
func keyPath(stateDir string, k SessionKey) string {
	return filepath.Join(stateDir, k.String()+".json")
}

// Load reads the pledge file for the given SessionKey. Returns ErrNoPledge if none exists.
func Load(k SessionKey) (*PledgeFile, error) {
	dir, err := StateDir()
	if err != nil {
		return nil, err
	}
	return loadFrom(keyPath(dir, k))
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

// Save persists the pledge for the given SessionKey, mode 0600.
func Save(k SessionKey, pf *PledgeFile) error {
	dir, err := StateDir()
	if err != nil {
		return err
	}
	if err := ensureDir(dir); err != nil {
		return fmt.Errorf("creating state directory: %w", err)
	}
	gcStalePledgeFiles(dir, k)

	data, err := json.MarshalIndent(pf, "", "  ")
	if err != nil {
		return err
	}
	path := keyPath(dir, k)
	LogAudit(AuditEntry{
		SessionKeyStr: k.String(),
		OperationID:   "<pledge-saved>",
		Outcome:       AuditOutcome(fmt.Sprintf("profile:%s", pf.Profile)),
	})
	return atomicWrite(path, data, 0o600)
}

// Widen saves pf as the new pledge for k without enforcing the narrowing
// constraint. Only call this after obtaining explicit interactive consent.
func Widen(k SessionKey, next *PledgeFile) error {
	next.NarrowedAt = time.Now().UTC()
	if err := Save(k, next); err != nil {
		return err
	}
	LogAudit(AuditEntry{
		SessionKeyStr: k.String(),
		OperationID:   "<pledge-widened>",
		Outcome:       AuditOutcome(fmt.Sprintf("widened-to:%s", next.MaxTier)),
	})
	return nil
}

// Narrow applies next as the new pledge for k, returning ErrWouldWiden if next
// would be more permissive than the current pledge.
func Narrow(k SessionKey, next *PledgeFile) error {
	current, err := Load(k)
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
	LogAudit(AuditEntry{
		SessionKeyStr: k.String(),
		OperationID:   "<pledge-narrowed>",
		Outcome:       AuditOutcome(fmt.Sprintf("profile:%s", next.Profile)),
	})
	return Save(k, next)
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

// gcStalePledgeFiles lazily removes stale pledge files. It handles both
// "sid-<n>.json" files (reaped when POSIX SID has no live process) and
// "claude-<uuid>.json" files (reaped when no breadcrumb exists AND mtime > 7 days).
func gcStalePledgeFiles(dir string, current SessionKey) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if name == current.String()+".json" {
			continue
		}

		switch {
		case isSIDFile(name):
			gcSIDFile(dir, name)
		case isClaudeFile(name):
			gcClaudeFile(dir, name, e)
		}
	}
}

func isSIDFile(name string) bool {
	var n int
	_, err := fmt.Sscanf(name, "sid-%d.json", &n)
	return err == nil
}

func isClaudeFile(name string) bool {
	if len(name) < len("claude-")+len(".json") {
		return false
	}
	prefix := "claude-"
	suffix := ".json"
	if !startsWith(name, prefix) || !endsWith(name, suffix) {
		return false
	}
	uuid := name[len(prefix) : len(name)-len(suffix)]
	return IsValidUUID(uuid)
}

func startsWith(s, prefix string) bool { return len(s) >= len(prefix) && s[:len(prefix)] == prefix }
func endsWith(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func gcSIDFile(dir, name string) {
	var sid int
	if _, err := fmt.Sscanf(name, "sid-%d.json", &sid); err != nil {
		return
	}
	if !sidHasLiveProcess(sid) {
		_ = os.Remove(filepath.Join(dir, name))
	}
}

func gcClaudeFile(dir, name string, e os.DirEntry) {
	prefix := "claude-"
	suffix := ".json"
	uuid := name[len(prefix) : len(name)-len(suffix)]

	// Check for any live breadcrumb pointing at this UUID.
	if claudeUUIDHasLiveBreadcrumb(uuid) {
		return
	}
	info, err := e.Info()
	if err != nil {
		return
	}
	if time.Since(info.ModTime()) > 7*24*time.Hour {
		_ = os.Remove(filepath.Join(dir, name))
	}
}

// claudeUUIDHasLiveBreadcrumb returns true if any breadcrumb file matching uuid
// exists under any project-dir subdirectory.
func claudeUUIDHasLiveBreadcrumb(uuid string) bool {
	// We don't know the project dir at GC time, so scan the entire
	// claude-session directory for any matching breadcrumb filename.
	breadcrumbRoot := os.Getenv("XDG_RUNTIME_DIR")
	if breadcrumbRoot == "" {
		stateDir, err := StateDir()
		if err != nil {
			return false
		}
		breadcrumbRoot = filepath.Dir(stateDir)
	}
	sessionDir := filepath.Join(breadcrumbRoot, "atlascli", "claude-session")
	entries, err := os.ReadDir(sessionDir)
	if err != nil {
		return false
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		candidate := filepath.Join(sessionDir, e.Name(), uuid)
		if _, err := os.Stat(candidate); err == nil {
			return true
		}
	}
	return false
}
