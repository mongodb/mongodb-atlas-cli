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
	"errors"
	"testing"
)

func TestCompareCmds(t *testing.T) {
	testCases := []struct {
		name        string
		mainData    map[string]cmdData
		changedData map[string]cmdData
		expectedErr error
	}{
		{
			name: "no changes",
			mainData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
						},
					},
				},
				"cmd2": {},
			},
			changedData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
						},
					},
				},
				"cmd2": {
					Aliases: []string{"cmdb"},
				},
			},
		},
		{
			name: "alias removed",
			mainData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
						},
					},
				},
				"cmd2": {
					Aliases: []string{"cmdb"},
				},
			},
			changedData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
						},
					},
				},
				"cmd2": {},
			},
			expectedErr: errCmdRemovedAlias,
		},
		{
			name: "deleted command",
			mainData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
						},
					},
				},
				"cmd2": {},
			},
			changedData: map[string]cmdData{
				"cmd2": {},
			},
			expectedErr: errCmdDeleted,
		},
		{
			name: "deleted flag",
			mainData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
						},
					},
				},
				"cmd2": {},
			},
			changedData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
				},
				"cmd2": {},
			},
			expectedErr: errFlagDeleted,
		},
		{
			name: "changed flag type",
			mainData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
						},
					},
				},
				"cmd2": {},
			},
			changedData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "int",
							Default: "default1",
						},
					},
				},
				"cmd2": {},
			},
			expectedErr: errFlagTypeChanged,
		},
		{
			name: "changed flag default value",
			mainData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
						},
					},
				},
				"cmd2": {},
			},
			changedData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default2",
						},
					},
				},
				"cmd2": {},
			},
			expectedErr: errFlagDefaultChanged,
		},
		{
			name: "changed flag shorthand",
			mainData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
							Short:   "f1",
						},
					},
				},
				"cmd2": {},
			},
			changedData: map[string]cmdData{
				"cmd1": {
					Aliases: []string{"cmda"},
					Flags: map[string]flagData{
						"flag1": {
							Type:    "string",
							Default: "default1",
							Short:   "f2",
						},
					},
				},
				"cmd2": {},
			},
			expectedErr: errFlagShortChanged,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			err := compareCmds(testCase.changedData, testCase.mainData)
			if err != nil && !errors.Is(err, testCase.expectedErr) {
				t.Fatalf("compareCmds failed: %v", err)
			} else if err == nil && testCase.expectedErr != nil {
				t.Fatalf("compareCmds should have failed: %v", err)
			}
		})
	}
}
