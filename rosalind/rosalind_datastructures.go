package rosalind

import (
	"fmt"
	"strings"
	"sync"
)

// Directed graph type
type DirGraph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// Graph node
type Node struct {
	name string
}

// Convert a node to a string
func (n *Node) String() string {
	return fmt.Sprintf("%s", n.name)
}

// add a node to the directed graph
func (g *DirGraph) AddNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

// add a directed edge
func (g *DirGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.lock.Unlock()
}

func (g *DirGraph) EdgeCount() int {
	iC := 0
	for _, targets := range g.edges {
		for i := 0; i < len(targets); i++ {
			iC++
		}
	}
	return iC
}

func (g *DirGraph) String() string {
	g.lock.RLock()

	// Keep it simple:
	// iterate through set of edges in random order,
	// and sort a bunch of strings at the end.
	edge_strings := make([]string, g.EdgeCount())
	iS := 0
	iE := 0
	for edge_src, edge_targets := range g.edges {
		for _, edge_target := range edge_targets {
			edge_string := edge_src.name + " -> " + edge_target.name
			edge_strings[iE] = edge_string
			iE += 1
		}
		iS += 1
	}
	result := strings.Join(edge_strings, "\n")

	g.lock.RUnlock()
	return result
}
