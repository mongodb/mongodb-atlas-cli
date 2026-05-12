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

//go:build !windows

package pledge

import (
	"crypto/sha1" //nolint:gosec // SHA1 is fine for non-cryptographic directory naming
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

// Session returns the POSIX session ID. Kept as a thin shim for callers that
// still use the raw int (e.g. CheckSessionLineage).
func Session() (int, error) {
	sid, err := unix.Getsid(0)
	if err != nil {
		return 0, fmt.Errorf("getsid(0): %w", err)
	}
	return sid, nil
}

// ResolveSessionKey determines the pledge session key for the current process.
//
// Resolution order:
//  1. ATLAS_PLEDGE_SESSION_KEY=<kind>:<value> — explicit override (tests, future agents).
//  2. ATLAS_PLEDGE_SESSION_ID — explicit Claude UUID override.
//  3. CLAUDE_CODE_SESSION_ID — Claude Code injects this into hook subprocesses
//     and Bash tool invocations. We trust it whenever it's set and well-formed,
//     even when CLAUDECODE is absent (SessionStart hooks don't always inherit
//     CLAUDECODE but do receive CLAUDE_CODE_SESSION_ID).
//  4. CLAUDECODE=1 in env → breadcrumb lookup under CLAUDE_PROJECT_DIR.
//  5. Fall back to POSIX SID.
func ResolveSessionKey() (SessionKey, error) {
	// 1. Explicit kind:value override.
	if raw := os.Getenv("ATLAS_PLEDGE_SESSION_KEY"); raw != "" {
		parts := strings.SplitN(raw, ":", 2)
		if len(parts) == 2 {
			return NewSessionKey(parts[0], parts[1])
		}
		return SessionKey{}, fmt.Errorf("ATLAS_PLEDGE_SESSION_KEY must be <kind>:<value>, got %q", raw)
	}

	// 2. Explicit Claude UUID override.
	if uuid := os.Getenv("ATLAS_PLEDGE_SESSION_ID"); uuid != "" && IsValidUUID(uuid) {
		return SessionKey{Kind: "claude", Value: uuid}, nil
	}

	// 3. Claude Code session ID — authoritative when present, regardless of CLAUDECODE.
	if uuid := os.Getenv("CLAUDE_CODE_SESSION_ID"); uuid != "" && IsValidUUID(uuid) {
		return SessionKey{Kind: "claude", Value: uuid}, nil
	}

	// 4. Claude Code path with breadcrumb fallback (CLAUDECODE=1, no UUID in env).
	if os.Getenv("CLAUDECODE") == "1" {
		if uuid, found := resolveClaudeBreadcrumb(); found {
			return SessionKey{Kind: "claude", Value: uuid}, nil
		}
		// Log + fall through to SID.
		sid, err := unix.Getsid(0)
		if err != nil {
			return SessionKey{}, fmt.Errorf("getsid(0): %w", err)
		}
		LogAudit(AuditEntry{
			SessionKeyStr: fmt.Sprintf("sid-%d", sid),
			OperationID:   "<claude-key-missing>",
			Outcome:       AuditOutcome(fmt.Sprintf("claude-key-missing: falling-back-to-sid-%d", sid)),
		})
		return SessionKey{Kind: "sid", Value: strconv.Itoa(sid)}, nil
	}

	// 5. POSIX SID.
	sid, err := unix.Getsid(0)
	if err != nil {
		return SessionKey{}, fmt.Errorf("getsid(0): %w", err)
	}
	return SessionKey{Kind: "sid", Value: strconv.Itoa(sid)}, nil
}

// resolveClaudeBreadcrumb scans CLAUDE_PROJECT_DIR's breadcrumb directory for
// the newest UUID file and returns it.
func resolveClaudeBreadcrumb() (string, bool) {
	projectDir := os.Getenv("CLAUDE_PROJECT_DIR")
	if projectDir == "" {
		return "", false
	}
	breadcrumbDir := claudeBreadcrumbDir(projectDir)
	entries, err := os.ReadDir(breadcrumbDir)
	if err != nil {
		return "", false
	}

	var bestName string
	var bestTime int64
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !IsValidUUID(name) {
			continue
		}
		info, err := e.Info()
		if err != nil {
			continue
		}
		if info.ModTime().UnixNano() > bestTime {
			bestTime = info.ModTime().UnixNano()
			bestName = name
		}
	}
	if bestName != "" {
		return bestName, true
	}
	return "", false
}

// claudeBreadcrumbDir returns the directory that holds UUID breadcrumb files
// for the given Claude project directory.
//
// Layout: <breadcrumbRoot>/atlascli/claude-session/<sha1(projectDir)>/
//
// where breadcrumbRoot is $XDG_RUNTIME_DIR when set, otherwise the pledge
// state directory.
func claudeBreadcrumbDir(projectDir string) string {
	root := os.Getenv("XDG_RUNTIME_DIR")
	if root == "" {
		// Fall back to pledge state dir — always writable.
		stateDir, err := StateDir()
		if err != nil {
			return ""
		}
		root = filepath.Dir(stateDir) // parent of .../atlascli/pledge → .../atlascli
	}
	//nolint:gosec // SHA1 used only for directory naming, not cryptographic security
	hash := sha1.Sum([]byte(projectDir))
	return filepath.Join(root, "atlascli", "claude-session", fmt.Sprintf("%x", hash))
}

// ClaudeBreadcrumbDir is the exported variant used by the hook installer to
// construct the shell command that writes breadcrumbs.
func ClaudeBreadcrumbDir(projectDir string) string {
	return claudeBreadcrumbDir(projectDir)
}

// CheckSessionLineage logs if this process is in a different POSIX session
// than its parent (escape-via-setsid detection) or if Claude env is set but
// no UUID was recovered (claude-key-missing).
func CheckSessionLineage() {
	mySID, err := unix.Getsid(0)
	if err != nil {
		return
	}
	ppid := os.Getppid()
	parentSID, err := unix.Getsid(ppid)
	if err != nil {
		return
	}
	if mySID != parentSID {
		LogAudit(AuditEntry{
			SessionKeyStr: fmt.Sprintf("sid-%d", mySID),
			OperationID:   "<sid-lineage-break>",
			Outcome:       AuditOutcome(fmt.Sprintf("session-break: parent-sid=%d", parentSID)),
		})
	}
}

// CheckSIDLineage is kept for backward-compat; delegates to CheckSessionLineage.
func CheckSIDLineage() {
	CheckSessionLineage()
}
