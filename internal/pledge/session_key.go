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
	"fmt"
	"regexp"
)

// SessionKey identifies the active pledge session. Kind is one of "sid" (POSIX
// session ID) or "claude" (Claude Code conversation UUID). On disk the pledge
// file is stored as <stateDir>/<kind>-<value>.json.
type SessionKey struct {
	Kind  string // "sid" or "claude"
	Value string // numeric SID string, or UUID
}

var allowedKinds = map[string]bool{"sid": true, "claude": true}

// uuidRE matches a UUID v4 (case-insensitive, standard hyphenated form).
var uuidRE = regexp.MustCompile(`(?i)^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`)

// IsValidUUID reports whether s is a UUID v4.
func IsValidUUID(s string) bool {
	return uuidRE.MatchString(s)
}

// NewSessionKey constructs a SessionKey, returning an error when kind is not
// in the allow-list.
func NewSessionKey(kind, value string) (SessionKey, error) {
	if !allowedKinds[kind] {
		return SessionKey{}, fmt.Errorf("unknown session key kind %q: must be one of: sid, claude", kind)
	}
	return SessionKey{Kind: kind, Value: value}, nil
}

// String serialises the key to the form used in filenames: "<kind>-<value>".
func (k SessionKey) String() string {
	return k.Kind + "-" + k.Value
}
