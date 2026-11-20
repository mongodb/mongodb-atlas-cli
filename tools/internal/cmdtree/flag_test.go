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

package cmdtree

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "list"}
	cmd.Flags().String("clusterName", "", "Cluster name")
	cmd.Flags().String("projectId", "", "Project id")
	cmd.Flags().String("start", "", "Start")
	cmd.Flags().Int("limit", 10, "Limit")
	cmd.Flags().Bool("includeCount", true, "Include count")
	if err := cmd.MarkFlagRequired("projectId"); err != nil {
		t.Fatalf("MarkFlagRequired: %v", err)
	}

	want := []Flag{
		{Name: "clusterName", Type: "string"},
		{Name: "includeCount", Type: "bool", Default: "true"},
		{Name: "limit", Type: "int", Default: "10"},
		{Name: "projectId", Type: "string", Required: true},
		{Name: "start", Type: "string"},
	}
	if got := flags(cmd); !reflect.DeepEqual(got, want) {
		t.Errorf("flags() = %+v, want %+v", got, want)
	}
}

// A flag declared on a parent belongs to every command, so it is not reported
// as if the command itself declared it.
func TestFlagsExcludesInherited(t *testing.T) {
	root := &cobra.Command{Use: "atlas"}
	root.PersistentFlags().String("profile", "", "Profile")
	child := &cobra.Command{Use: "list"}
	child.Flags().String("clusterName", "", "Cluster name")
	root.AddCommand(child)

	want := []Flag{{Name: "clusterName", Type: "string"}}
	if got := flags(child); !reflect.DeepEqual(got, want) {
		t.Errorf("flags(child) = %+v, want %+v", got, want)
	}
}
