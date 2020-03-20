package convert

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/go-client-mongodb-ops-manager/opsmngr"
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
