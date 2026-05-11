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
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// AuditOutcome describes the outcome of an enforced operation.
type AuditOutcome string

const (
	AuditBlocked        AuditOutcome = "blocked"
	AuditAllowedByToken AuditOutcome = "allowed-by-token"
)

// AuditEntry is a single JSON line in the audit log.
type AuditEntry struct {
	Timestamp    time.Time    `json:"timestamp"`
	SID          int          `json:"sid"`
	OperationID  string       `json:"operationID"`
	ParamsHash   string       `json:"paramsHash,omitempty"`
	Outcome      AuditOutcome `json:"outcome"`
	ApproverSID  int          `json:"approverSID,omitempty"`
}

// LogAudit appends an entry to the audit log, creating it if necessary.
// Errors are ignored — audit failures must not block the CLI.
func LogAudit(entry AuditEntry) {
	dir, err := StateDir()
	if err != nil {
		return
	}
	if err := ensureDir(dir); err != nil {
		return
	}
	logPath := filepath.Join(dir, "audit.log")

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		return
	}
	defer f.Close()

	data, err := json.Marshal(entry)
	if err != nil {
		return
	}
	_, _ = fmt.Fprintf(f, "%s\n", data)
}
