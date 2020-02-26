package search_test

import (
	"fmt"

	"github.com/mongodb-labs/pcgc/cloudmanager"
	"github.com/mongodb/mongocli/internal/fixtures"
	"github.com/mongodb/mongocli/internal/search"
)

// This example demonstrates searching a list.
func ExampleStringInSlice() {
	a := []string{"a", "b", "c"}
	x := "a"

	if search.StringInSlice(a, x) {
		fmt.Printf("found %s in %v\n", x, a)
	} else {
		fmt.Printf("%s not found in %v\n", x, a)
	}
	// Output:
	// found a in [a b c]
}

// This example demonstrates searching a list of processes by name.
func ExampleProcesses() {
	a := fixtures.AutomationConfig().Processes
	x := "myReplicaSet_2"
	i, found := search.Processes(a, func(p *cloudmanager.Process) bool { return p.Name == x })
	if i < len(a) && found {
		fmt.Printf("found %v at index %d\n", x, i)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found myReplicaSet_2 at index 1
}

// This example demonstrates searching a list of replica sets by ID.
func ExampleReplicaSets() {
	a := fixtures.AutomationConfig().ReplicaSets
	x := "myReplicaSet"
	i, found := search.ReplicaSets(a, func(r *cloudmanager.ReplicaSet) bool { return r.ID == x })
	if i < len(a) && found {
		fmt.Printf("found %v at index %d\n", x, i)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found myReplicaSet at index 0
}

// This example demonstrates searching a list of members by host.
func ExampleMembers() {
	a := fixtures.AutomationConfig().ReplicaSets[0].Members
	x := "myReplicaSet_2"
	i, found := search.Members(a, func(m cloudmanager.Member) bool { return m.Host == x })
	if i < len(a) && found {
		fmt.Printf("found %v at index %d\n", x, i)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found myReplicaSet_2 at index 1
}

// This example demonstrates searching a cluster in an automation config.
func ExampleClusterExists() {
	a := fixtures.AutomationConfig()
	x := "myReplicaSet"
	found := search.ClusterExists(a, x)
	if found {
		fmt.Printf("found %v\n", x)
	} else {
		fmt.Printf("%s not found\n", x)
	}
	// Output:
	// found myReplicaSet
}
