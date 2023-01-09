// Copyright 2022 MongoDB Inc
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

package terminal

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

// IsTerminal returns true is the current file descriptor is TTY kind of terminal.
func IsTerminal(w io.Writer) bool {
	if f, isFile := w.(*os.File); isFile {
		return isatty.IsTerminal(f.Fd()) || IsCygwinTerminal(w)
	}

	return false
}

// IsCygwinTerminal returns true is the current file descriptor is cygwin.
func IsCygwinTerminal(w io.Writer) bool {
	if f, isFile := w.(*os.File); isFile {
		return isatty.IsCygwinTerminal(f.Fd())
	}

	return false
}

// IsTerminalInput returns true is the current file descriptor is TTY kind of terminal.
func IsTerminalInput(r io.Reader) bool {
	if f, isFile := r.(*os.File); isFile {
		return isatty.IsTerminal(f.Fd()) || IsCygwinTerminalInput(r)
	}

	return false
}

// IsCygwinTerminalInput returns true is the current file descriptor is cygwin.
func IsCygwinTerminalInput(r io.Reader) bool {
	if f, isFile := r.(*os.File); isFile {
		return isatty.IsCygwinTerminal(f.Fd())
	}

	return false
}
