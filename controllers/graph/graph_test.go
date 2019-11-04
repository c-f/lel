package graph

import (
	"encoding/json"
	"testing"
)

// contains is a helper function to test if a key is in the slice
func contains(slice []string, key string) bool {

	for i := range slice {
		if slice[i] == key {
			return true
		}
	}
	return false
}

// TestEmptyRef creates two nodes
// empty node -> edge -> this
func TestEmptyRef(t *testing.T) {
	t.Log("Testing", "TestEmptyRef")
	g := New()

	info := &MetaInfo{
		Names:  []string{"reference"},
		Entity: "ref",
		Id:     "non-existing/ref.md",
	}

	// initialize
	err := g.Add(info)
	if err != nil {
		t.Fatal(err)
	}

	// Build and export
	export := g.Export()
	// further Export
	if len(export.Nodes) != 0 || len(export.Nodes) != 0 {
		t.Fatal("no Nodes should be displayed ")
	}
}

func TestBadInput(t *testing.T) {
	g := New()
	badInfos := []*MetaInfo{
		&MetaInfo{
			Names:  []string{"bob", "bob.dns.local"},
			Entity: "human",
			Id:     "non-existing/bob-origin.md",
			Relations: []string{
				"(this)->[(pointsto)]->(alice)",
			},
		},
		&MetaInfo{
			Names:  []string{"alice", "alice.dns.local"},
			Entity: "human",
			Id:     "non-existing/alice-origin.md",
			Relations: []string{
				"(thiss)->[pointsto]->(unknown)",
				"(this)->[:this]->(this)",
			},
		},
		&MetaInfo{
			Names:  []string{"alice", "alice.dns.local"},
			Entity: "human",
			Id:     "non-existing/alice-origin.md",
			Relations: []string{
				"(thiss)->[pointsto]->(unknown)",
				"(this)->[:this]->(this)",
			},
		},
	}

	// initialize
	for _, info := range badInfos {
		err := g.Add(info)
		if err != nil {
			t.Fatal(err)
		}

	}

	// Build and export
	export := g.Export()

	bts, err := json.Marshal(export)
	if err != nil {
		t.Fatal("Convertion to json failed")
	}
	t.Log(string(bts))
}

// TestEmptyEdgeLoop creates two nodes, which both
// reference each other and an empty node
// This should create three nodes, including one empty
func TestEmptyEdgeLoop(t *testing.T) {
	t.Log("Testing", "TestNodes")
	g := New()

	// test variables
	var names []string
	infos := []*MetaInfo{
		&MetaInfo{
			Names:  []string{"bob", "bob.dns.local"},
			Entity: "human",
			Id:     "non-existing/bob-origin.md",
			Relations: []string{
				"(this)->[(pointsto)]->(alice)",
			},
		},
		&MetaInfo{
			Names:  []string{"alice", "alice.dns.local"},
			Entity: "human",
			Id:     "non-existing/alice-origin.md",
			Relations: []string{
				"(this)->[pointsto]->(unknown)",
				"(this)->[pointsto]->(bob)",
			},
		},
	}

	// initialize
	for _, info := range infos {
		err := g.Add(info)
		if err != nil {
			t.Fatal(err)
		}
		names = append(names, info.Id)
	}

	// initialize
	for _, info := range infos {
		err := g.Add(info)
		if err != nil {
			t.Fatal(err)
		}
		names = append(names, info.Id)
	}

	// Build and export
	export := g.Export()

	bts, err := json.Marshal(export)
	if err != nil {
		t.Fatal("Convertion to json failed")
	}
	t.Log(string(bts))

	// validate
	// TODO
}

// TestNodes testing the creation of nodes
func TestNodesWithEmptyEdge(t *testing.T) {
	t.Log("Testing", "TestNodes")
	g := New()

	// test variables
	var names []string
	infos := []*MetaInfo{
		&MetaInfo{
			Names:  []string{"bob", "bob.dns.local"},
			Entity: "human",
			Id:     "non-existing/bob-origin.md",
			Relations: []string{
				"(this)->[pointsto]->(alice)",
			},
		},
		&MetaInfo{
			Names:     []string{"alice", "alice.dns.local"},
			Entity:    "human",
			Id:        "non-existing/alice-origin.md",
			Relations: []string{},
		},
	}

	// initialize
	for _, info := range infos {
		err := g.Add(info)
		if err != nil {
			t.Fatal(err)
		}
		names = append(names, info.Id)
	}

	// Build and export
	export := g.Export()

	bts, err := json.Marshal(export)
	if err != nil {
		t.Fatal("Convertion to json failed")
	}
	t.Log(string(bts))

	// tests
	/*
		if len(export.Nodes) != len(infos)  {
			t.Fatal("More nodes are created")
		}
		if len(export.Edges) != 0{
			t.Fatal("More eges were created")
		}
		for _, node := range export.Nodes {
			if node.IsEmpty {
				t.Fatal("Node should not be empty")
			}
			if node.Entity != "human"{
				t.Fatal("Node should have the entity human")
			}
			if len(node.Tags) != 0 {
				t.Fatal("Tags should be zero")
			}
			if ok := contains(names, node.Id ); !ok {
				t.Fatal("Unknown ID")
			}
		}
	*/
}

func TestTwoNodes(t *testing.T) {
	t.Log("Testing", "TestTwoNodes")
	g := New()
	infos := []*MetaInfo{
		&MetaInfo{
			Names:  []string{"bob"},
			Id:     "non-existing/node.md",
			Entity: "node",
			Relations: []string{
				"(this)->[likes]->(alice)",
			},
		},
		&MetaInfo{
			Names:  []string{"alice"},
			Id:     "non-existing/alice-node.md",
			Entity: "node",
		},
	}
	// initialize
	for _, info := range infos {
		err := g.Add(info)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestEmptyRelation(t *testing.T) {
	t.Log("Testing", "TestEmptyRelation")
	g := New()
	info := &MetaInfo{
		Names:  []string{"bob"},
		Id:     "non-existing/node.md",
		Entity: "node",
		Relations: []string{
			"(this)->[likes]->(alice)",
		},
	}

	// initialize
	err := g.Add(info)
	if err != nil {
		t.Fatal(err)
	}

	// Build and export
	export := g.Export()

	bts, err := json.MarshalIndent(export, "", "  ")
	if err != nil {
		t.Fatal("Convertion to json failed")
	}
	t.Log(string(bts))

	// evaluate
}
