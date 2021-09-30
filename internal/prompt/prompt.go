// Copyright 2021 MongoDB Inc
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

package prompt

import (
	"fmt"
	"io"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/mattn/go-isatty"
)

// IsTerminalInput returns true if input device is a terminal
// Will return false in the case of input being piped.
func IsTerminalInput() bool {
	return IsTerminalInputStream(os.Stdin)
}

// isTerminalInput returns true if input device is a terminal
// Will return false in the case of the reader not being a terminal.
func IsTerminalInputStream(stream io.Reader) bool {
	if inputStream, ok := stream.(*os.File); ok {
		return isatty.IsTerminal(inputStream.Fd()) || isatty.IsCygwinTerminal(inputStream.Fd())
	}
	return false
}

// NewDeleteConfirm creates a prompt to confirm if the entry should be deleted.
func NewDeleteConfirm(entry string) survey.Prompt {
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete: %s", entry),
	}
	return prompt
}

// NewConfirm creates a prompt to confirm if the entry should be deleted.
func NewConfirm(message string) survey.Prompt {
	prompt := &survey.Confirm{
		Message: message,
	}
	return prompt
}

// NewProfileReplaceConfirm creates a prompt to confirm if an existing profile should be replaced.
func NewProfileReplaceConfirm(entry string) survey.Prompt {
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("There is already a profile called %s.\nDo you want to replace it?", entry),
	}
	return prompt
}
