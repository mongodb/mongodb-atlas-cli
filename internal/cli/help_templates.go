// Copyright 2025 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

const (
	// HelpTemplate
	// Modified version of https://github.com/spf13/cobra/blob/01ffff4eca5a08384ef2b85f39ec0dac192a5f7b/command.go#L595 which shows both .Short and .Long help descriptions.
	HelpTemplate = `{{.Short | trimTrailingWhitespaces}} {{.Long | trimTrailingWhitespaces}}

{{if or .Runnable .HasSubCommands}}{{.UsageString}}{{end}}`

	ExperimentalText = "Public Preview: The atlas api sub-command, " +
		"automatically generated from the MongoDB Atlas Admin API, offers full coverage of the Admin API and " +
		"is currently in Public Preview (please provide feedback " +
		"at https://feedback.mongodb.com/forums/930808-atlas-cli).\n" +
		"Admin API capabilities have their own release lifecycle, which you can check via the provided API endpoint " +
		"documentation link.\n\n"

	ExperimentalHelpTemplate = ExperimentalText + HelpTemplate
)
