package search_test

import (
	"testing"

	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/search"
)

func TestStringInSlice(t *testing.T) {
	s := []string{"a", "b", "c"}
	t.Run("value exists", func(t *testing.T) {
		if !search.StringInSlice(s, "b") {
			t.Error("StringInSlice() should find the value")
		}
	})

	t.Run("value not exists", func(t *testing.T) {
		if search.StringInSlice(s, "d") {
			t.Error("StringInSlice() should not find the value")
		}
	})
}

func TestClusterExists(t *testing.T) {
	t.Run("value exists", func(t *testing.T) {
		if !search.ClusterExists(fixtures.AutomationConfig(), "myReplicaSet") {
			t.Error("ClusterExists() should find the value")
		}
	})

	t.Run("value not exists", func(t *testing.T) {
		if search.ClusterExists(fixtures.AutomationConfig(), "X") {
			t.Error("StringInSlice() should not find the value")
		}
	})
}
