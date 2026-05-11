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

import "syscall"

// sidHasLiveProcess returns false when no process in session sid is alive.
// It uses kill(0) on the session leader ID as a proxy; if the session has no
// processes the signal delivery fails with ESRCH.
func sidHasLiveProcess(sid int) bool {
	// getsid on another process would require iterating /proc; instead send
	// signal 0 to negative-PID (process group) as an approximation.
	// Stale sessions typically have no process group alive.
	err := syscall.Kill(-sid, 0)
	return err == nil
}
