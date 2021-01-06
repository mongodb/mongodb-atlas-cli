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

// +build unit

package convert

import (
	"testing"

	"github.com/go-test/deep"
	"go.mongodb.org/atlas/mongodbatlas"
	"go.mongodb.org/ops-manager/opsmngr"
)

func TestBuildAtlasRoles(t *testing.T) {
	t.Run("No database defaults to admin", func(t *testing.T) {
		r := BuildAtlasRoles([]string{"admin"})
		expected := []mongodbatlas.Role{
			{
				RoleName:     "admin",
				DatabaseName: "admin",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})

	t.Run("should split by @", func(t *testing.T) {
		r := BuildAtlasRoles([]string{"admin@test"})
		expected := []mongodbatlas.Role{
			{
				RoleName:     "admin",
				DatabaseName: "test",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})

	t.Run("all", func(t *testing.T) {
		r := BuildAtlasRoles([]string{"admin@test", "something"})
		expected := []mongodbatlas.Role{
			{
				RoleName:     "admin",
				DatabaseName: "test",
			},
			{
				RoleName:     "something",
				DatabaseName: "admin",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})
}

func TestBuildOMRoles(t *testing.T) {
	t.Run("No database defaults to admin", func(t *testing.T) {
		r := BuildOMRoles([]string{"admin"})
		expected := []*opsmngr.Role{
			{
				Role:     "admin",
				Database: "admin",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})

	t.Run("should split by @", func(t *testing.T) {
		r := BuildOMRoles([]string{"admin@test"})
		expected := []*opsmngr.Role{
			{
				Role:     "admin",
				Database: "test",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})

	t.Run("all", func(t *testing.T) {
		r := BuildOMRoles([]string{"admin@test", "something"})
		expected := []*opsmngr.Role{
			{
				Role:     "admin",
				Database: "test",
			},
			{
				Role:     "something",
				Database: "admin",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})
}

func TestBuildAtlasInheritedRoles(t *testing.T) {
	t.Run("No database defaults to admin", func(t *testing.T) {
		r := BuildAtlasInheritedRoles([]string{"admin"})
		expected := []mongodbatlas.InheritedRole{
			{
				Role: "admin",
				Db:   "admin",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})

	t.Run("should split by @", func(t *testing.T) {
		r := BuildAtlasInheritedRoles([]string{"admin@test"})
		expected := []mongodbatlas.InheritedRole{
			{
				Role: "admin",
				Db:   "test",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})

	t.Run("all", func(t *testing.T) {
		r := BuildAtlasInheritedRoles([]string{"admin@test", "something"})
		expected := []mongodbatlas.InheritedRole{
			{
				Role: "admin",
				Db:   "test",
			},
			{
				Role: "something",
				Db:   "admin",
			},
		}
		if err := deep.Equal(r, expected); err != nil {
			t.Error(err)
		}
	})
}
