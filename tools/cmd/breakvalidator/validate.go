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

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/spf13/cobra"
)

var (
	errFlagDeleted        = errors.New("flag was deleted")
	errFlagTypeChanged    = errors.New("flag type changed")
	errFlagDefaultChanged = errors.New("flag default value changed")
	errCmdDeleted         = errors.New("command deleted")
	errCmdRemovedAlias    = errors.New("command alias removed")
	errFlagShortChanged   = errors.New("flag shorthand changed")
)

func compareFlags(cmdPath string, mainFlags, changedFlags map[string]flagData) []error {
	if mainFlags == nil {
		return nil
	}

	changes := []error{}

	for flagName, flagValue := range mainFlags {
		changedFlagValue, ok := changedFlags[flagName]
		if !ok {
			changes = append(changes, fmt.Errorf("%w: %s --%s", errFlagDeleted, cmdPath, flagName))
			continue
		}

		if flagValue.Type != changedFlagValue.Type {
			changes = append(changes, fmt.Errorf("%w: %s --%s", errFlagTypeChanged, cmdPath, flagName))
		}

		// Ignore default value changes for the --version flag as this is allowed
		if flagValue.Default != changedFlagValue.Default && flagName != "version" {
			changes = append(changes, fmt.Errorf("%w: %s --%s", errFlagDefaultChanged, cmdPath, flagName))
		}

		if flagValue.Short != changedFlagValue.Short {
			changes = append(changes, fmt.Errorf("%w: %s --%s", errFlagShortChanged, cmdPath, flagName))
		}
	}

	return changes
}

func compareCmds(changedData, mainData map[string]cmdData) error {
	changes := []error{}
	for cmdPath, mv := range mainData {
		cv, ok := changedData[cmdPath]
		if !ok {
			changes = append(changes, fmt.Errorf("%w: %s", errCmdDeleted, cmdPath))
			continue
		}

		if mv.Aliases != nil {
			mainAliases := mv.Aliases
			changedAliases := cv.Aliases

			for _, alias := range mainAliases {
				if !slices.Contains(changedAliases, alias) {
					changes = append(changes, fmt.Errorf("%w: %s", errCmdRemovedAlias, cmdPath))
				}
			}
		}

		changes = append(changes, compareFlags(cmdPath, mv.Flags, cv.Flags)...)
	}

	if len(changes) > 0 {
		return errors.Join(changes...)
	}

	return nil
}

func validateCmdRun(output io.Writer, mainFile, changedFile string) error {
	mainBuffer, err := os.ReadFile(mainFile)
	if err != nil {
		return err
	}
	changedBuffer, err := os.ReadFile(changedFile)
	if err != nil {
		return err
	}

	var mainData, changedData map[string]cmdData
	if err := json.Unmarshal(mainBuffer, &mainData); err != nil {
		return err
	}
	if err := json.Unmarshal(changedBuffer, &changedData); err != nil {
		return err
	}

	err = compareCmds(changedData, mainData)
	if err != nil {
		return err
	}

	fmt.Fprintln(output, "no breaking changes detected")
	return nil
}

func buildValidateCmd() *cobra.Command {
	var mainFile, changedFile string
	validateCmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate the CLI command structure.",
		PreRunE: func(_ *cobra.Command, _ []string) error {
			if mainFile == "" {
				return errors.New("--main flag is required")
			}
			if changedFile == "" {
				return errors.New("--changed flag is required")
			}
			if _, err := os.Stat(mainFile); os.IsNotExist(err) {
				return fmt.Errorf("file does not exist: %s", mainFile)
			}
			if _, err := os.Stat(changedFile); os.IsNotExist(err) {
				return fmt.Errorf("file does not exist: %s", changedFile)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return validateCmdRun(cmd.OutOrStdout(), mainFile, changedFile)
		},
	}
	validateCmd.Flags().StringVarP(&mainFile, "main", "m", "", "Main file")
	validateCmd.Flags().StringVarP(&changedFile, "changed", "c", "", "Changed file")
	_ = validateCmd.MarkFlagRequired("main")
	_ = validateCmd.MarkFlagRequired("changed")
	return validateCmd
}
