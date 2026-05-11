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
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

// Session returns the POSIX session ID of the current process.
// Every child of the same shell session shares this SID, so pledge state stored
// under this key is inherited by all atlas invocations within the session.
func Session() (int, error) {
	sid, err := unix.Getsid(0)
	if err != nil {
		return 0, fmt.Errorf("getsid(0): %w", err)
	}
	return sid, nil
}

// CheckSIDLineage logs a one-line audit note if this process is in a different
// session than its parent. This happens when a user runs `setsid atlas ...` to
// escape a pledge — the note makes the break detectable in audit logs.
func CheckSIDLineage() {
	mySID, err := unix.Getsid(0)
	if err != nil {
		return
	}
	ppid := os.Getppid()
	parentSID, err := unix.Getsid(ppid)
	if err != nil {
		return
	}
	if mySID == parentSID {
		return
	}
	LogAudit(AuditEntry{
		SID:         mySID,
		OperationID: "<sid-lineage-break>",
		Outcome:     AuditOutcome(fmt.Sprintf("session-break: parent-sid=%d", parentSID)),
	})
}
