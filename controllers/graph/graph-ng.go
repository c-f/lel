package graph

import (
	"errors"
	"log"

	"strings"
	"sort"
)

// Error information
var (
	ErrNoMetaInfo = errors.New("No Metainformation could be found ")
)

// Export defines all information necessary, which can be used by the frontend graph
type Export struct {
	Nodes   []*Node           `json:"nodes"`
	Edges   []*Edge           `json:"edges"`
	Unknown map[string]string `json:"unknown"`
	Errs    map[string]string `json:"-"`
}

// MetaGraph is a collector for all Metainformation and its nameRegister
// MetaGraph can be used to create a Export "object" and mapps stuff
type MetaGraph struct {
	infos map[string]*MetaInfo
	edges map[string]*MetaInfo
	nodes map[string]*MetaInfo

	nameRegister map[string]*Register

	doublicates map[string]string
}

// Register defines which types a enty is (info,edge,node) and what the key(name)
// should be used. This is necessary to gerate multiple map for names
type Register struct {
	Type int
	Key  string
}

// New generates a new MetaGraph object
func New() *MetaGraph {
	g := &MetaGraph{}
	g.Clear()
	return g
}

// IndexFile gets a filepath and a origin(key), which is then added to the Metagraph, based on the facts
func (g *MetaGraph) IndexFile(filePath, origin string) error {
	info, err := MetaInfoFromFile(filePath, origin)
	if err != nil {
		return err
	}
	return g.Add(info)
}

// Get returns the metainfo for a specific id
func (g *MetaGraph) Get(id string) *MetaInfo {
	if val, ok := g.nodes[id]; ok {
		return val
	}
	if val, ok := g.edges[id]; ok {
		return val
	}
	if val, ok := g.infos[id]; ok {
		return val
	}
	return nil
}

// Add adds the MetaInfo to the Graph, based on the type, adds names to the map
func (g *MetaGraph) Add(m *MetaInfo) error {
	if ok := m.Valid(); !ok {
		return ErrNoMetaInfo
	}

	// get Type of file
	switch m.GetType() {

	case NODE_TYPE:
		g.nodes[m.Id] = m

	case EDGE_TYPE:
		g.edges[m.Id] = m

	case INFO_TYPE:
		g.infos[m.Id] = m

	case BULK_TYPE:
		// ignore
	case UNKNOWN_TYPE:
		// ignore
	}

	// register Alias
	for _, alias := range m.Names {
		// detect doublicated key
		if _, ok := g.nameRegister[alias]; ok {
			g.doublicates[alias] = m.Id
		}
		g.nameRegister[alias] = &Register{
			Type: m.GetType(),
			Key:  m.Id,
		}
	}
	return nil
}

// Clear empties all registries for the Graph
func (g *MetaGraph) Clear() {
	g.infos = make(map[string]*MetaInfo)
	g.edges = make(map[string]*MetaInfo)
	g.nodes = make(map[string]*MetaInfo)
	g.nameRegister = make(map[string]*Register)
	g.doublicates = make(map[string]string)
}


// Export converts the MetaGraph in a Export object
func (g *MetaGraph) Export() Export {

	// initial empty state
	nodes := []*Node{}
	edges := []*Edge{}
	unknown := make(map[string]string)
	errs := make(map[string]string)

	// Add Infos
	for origin, info := range g.infos {
		for _, raw := range info.Relations {
			alias := GetRefInfo(raw)
			if alias == "" {
				continue
			}

			// get key for an alias
			if register, ok := g.nameRegister[alias]; ok {
				
				// Add reference
				switch register.Type {
				case NODE_TYPE:
					g.nodes[register.Key].AddRef(origin)

				case EDGE_TYPE:
					g.edges[register.Key].AddRef(origin)
				default:
					log.Println("Could not add info to this type", register.Type)
				}
			} else {

				// currently unknown, therefore saving the origin with the alias
				unknown[alias] = origin
				continue
			}
		}
	}

	// Add nodes
	for origin, info := range g.nodes {
		node := NewOriginNode(info)
		nodes = SortedInsertNodes(nodes, node)

		// First relations
		for _, raw := range info.Relations {
			var edge *Edge
			
			// origin
			subject, relation, object, err := GetRefRelation(raw, origin)
			if err != nil {
				errs[origin] = err.Error()
				continue
			}
			// search for subject
			if subject != origin {
				if register, ok := g.nameRegister[subject]; ok {
					subject = register.Key
				} else {
					unknown[subject] = origin

					emptyNode := NewNode(subject)
					nodes = SortedInsertNodes(nodes, emptyNode)
					subject = emptyNode.Id
					continue
				}
			}
			// search for object
			if object != origin {
				if register, ok := g.nameRegister[object]; ok {
					object = register.Key
				} else {
					unknown[object] = origin

					emptyNode := NewNode(object)
					nodes = SortedInsertNodes(nodes, emptyNode)
					object = emptyNode.Id
					continue
				}
			}

			// search for relation
			if IsRelationAlias(relation) {
				alias := GetAlias(relation)

				if register, ok := g.nameRegister[alias]; ok {
					if register.Type != EDGE_TYPE {
						log.Println("Could not add info to this type", register.Type)
						continue
					}
					alias = register.Key
				} else {
					unknown[alias] = origin
					continue
				}
				// Get Informations
				edgeInfo, _ := g.edges[alias]
				edge = NewOriginEdge(edgeInfo, subject, object)
			} else {
				// create empty node
				edge = NewEdge(relation, subject, object)
			}

			// append relation
			edges = SortedInsertEdges(edges, edge)

		}
	}

	// Add Edges
	for origin, info := range g.edges {
		for _, raw := range info.Relations {
			var edge *Edge

			//origin
			subject, relation, object, err := GetRefRelation(raw, origin)
			if err != nil {
				errs[origin] = err.Error()
				continue
			}
			if relation == origin {
				edge = NewOriginEdge(info, subject, object)
			} else {
				// wether subject nor object should contain this before
				if subject == origin || object == origin {
					errs[origin] = "cannot specify relation as a node"
					continue
				}

				// search for subject
				if register, ok := g.nameRegister[subject]; ok {
					subject = register.Key
				} else {
					unknown[subject] = origin

					emptyNode := NewNode(subject)
					nodes = SortedInsertNodes(nodes, emptyNode)
					subject = emptyNode.Id
					continue
				}

				// search for object
				if register, ok := g.nameRegister[object]; ok {
					object = register.Key
				} else {
					unknown[object] = origin

					emptyNode := NewNode(object)
					nodes = SortedInsertNodes(nodes, emptyNode)
					object = emptyNode.Id
					continue
				}

				// create empty node
				edge = NewEdge(relation, subject, object)
			}

			// append relation
			edges = SortedInsertEdges(edges, edge)
		}
	}
	return Export{
		Nodes:   nodes,
		Edges:   edges,
		Unknown: unknown,
		Errs:    errs,
	}
}


// SortInsert creates a stable representation
func SortedInsertNodes (nodes []*Node, newNode *Node) []*Node {
	index := sort.Search(len(nodes), func(i int) bool { return strings.Compare(nodes[i].Id,newNode.Id ) > 0 })
	nodes = append(nodes, &Node{})
	copy(nodes[index+1:], nodes[index:])
	nodes[index] = newNode
	return nodes
}

// SortInsert creates a stable representation
func SortedInsertEdges (edges []*Edge, edge *Edge) []*Edge {
	index := sort.Search(len(edges), func(i int) bool { return strings.Compare(edges[i].Id,edge.Id ) > 0 })
	edges = append(edges, &Edge{})
	copy(edges[index+1:], edges[index:])
	edges[index] = edge
	return edges
}