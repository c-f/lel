package controllers

import (
	"strings"
	"sort"
)

// TagStructure contains all tags in a treeStructure
type TreeStructure struct {
	Root *Node
}

// NewTagStructure returns a new tagstructure starting with the root node
func NewTreeStructure() *TreeStructure {
	return &TreeStructure{
		Root: &Node{
			TagType:  "root",
			Children: make(map[string]*Node),
			Content:  []string{"root"},
		},
	}
}

func (s *TreeStructure) AddItem(origin, item string, additional map[string]interface{}) {
	cleanTag := strings.TrimSpace(item)
	parts := strings.Split(cleanTag, "/")
	// hosts/lel
	end := len(parts)
	var current *Node
	current = s.Root
	for i, part := range parts {
		tagType := "node"
		if i+1 == end {
			tagType = "leaf"
		}

		// add new Tag
		if _, ok := current.Children[part]; !ok {
			// new value should be added

			current.Children[part] = &Node{
				Content:  []string{origin},
				TagType:  tagType,
				Children: make(map[string]*Node),
				Context:  additional,
			}
		} else {
			if tagType == "leaf" {
				current.Content = append(current.Content, origin)
			}
		}
		current = current.Children[part]
	}
}

// AddItems adds items to the structure, items will be seperated based on content
func (s *TreeStructure) AddItems(origin string, items ...string) {
	for _, item := range items {
		s.AddItem(origin, item, map[string]interface{}{})
	}
	return
}

// BuildAntGraph builds a node
func (s *TreeStructure) BuildAntGraph() []*AntNode {
	result := []*AntNode{}
	for name, node := range s.Root.Children {
		antNode := node.ConvertToAnt(name)

		result = SortedInsert(result, antNode)
	}
	return result
}

// AntNode is a node representation for the ant framwork
// https://ant.design/components/tree/
type AntNode struct {
	Title    string     `json:"title"`
	Value    string     `json:"value"`
	Key      string     `json:"key"`
	Children []*AntNode `json:"children,omitempty"`

	//
	Other map[string]interface{} `json:"other,omitempty"`
}



// Tag defines a Tag
type Node struct {
	Content  []string         `json:"content"`
	Children map[string]*Node `json:"nodes"`

	TagType string `json:"type"`

	// additionalData
	Context map[string]interface{} `json:"ctx,omitempty"`
}


// ConvertToAnt converts the node to the format, where trees can be built
func (n *Node) ConvertToAnt(name string) *AntNode {
	core := &AntNode{
		Title:    name,
		Value:    name,
		Key:      "unknown",
		Children: []*AntNode{},

		Other: n.Context,
	}
	if len(n.Content) > 0 {
		core.Key = n.Content[0]

		if n.TagType == "node" {
			parts := strings.Split(core.Key, "/")
			core.Key = strings.Join(parts[:len(parts)-1], "/")
		}
	}
	
	for name, node := range n.Children {
		core.Children = SortedInsert(core.Children, node.ConvertToAnt(name))
	}
	return core
}

// SortInsert creates a stable representation
func SortedInsert (nodes []*AntNode, newNode *AntNode) []*AntNode {
	index := sort.Search(len(nodes), func(i int) bool { return strings.Compare(nodes[i].Title,newNode.Title ) > 0 })
	nodes = append(nodes, &AntNode{})
	copy(nodes[index+1:], nodes[index:])
	nodes[index] = newNode
	return nodes
}


// todo directories have the wrong key - should only be test.

/*
   0: {title: "0-0-0", key: "0-0-0", children: Array(3)}
   1: {title: "0-0-1", key: "0-0-1", children: Array(3)}
   2: {title: "0-0-2", key: "0-0-2"}
   1: {title: "0-1", key: "0-1", children: Array(3)}
   2: {title: "0-2", key: "0-2"}
*/

/*
{
  "content": [
    "root"
  ],
  "nodes": {
    "angela.md": {
      "content": [
        "angela.md"
      ],
      "nodes": {},
      "type": "leaf"
    },
    "lel.md": {
      "content": [
        "lel.md"
      ],
      "nodes": {},
      "type": "leaf"
    },
    "lels.md": {
      "content": [
        "lels.md"
      ],
      "nodes": {},
      "type": "leaf"
    },
    "navigation.md": {
      "content": [
        "navigation.md"
      ],
      "nodes": {},
      "type": "leaf"
    },
    "test": {
      "content": [
        "test/test.md"
      ],
      "nodes": {
        "test.md": {
          "content": [
            "test/test.md"
          ],
          "nodes": {},
          "type": "leaf"
        },
        "test2.md": {
          "content": [
            "test/test2.md"
          ],
          "nodes": {},
          "type": "leaf"
        }
      },
      "type": "node"
    }
  },
  "type": "root"
}

*/

// array - title key children Children needs to be array
