package rosalind

import (
	"fmt"
	"testing"
)

func fillGraph() DirGraph {

	var g DirGraph

	n1a := Node{"AGGCA"}
	n1b := Node{"GGCAT"}

	g.AddNode(&n1a)
	g.AddNode(&n1b)
	g.AddEdge(&n1a, &n1b)

	n2a := Node{"CATGC"}
	n2b := Node{"ATGCG"}

	g.AddNode(&n2a)
	g.AddNode(&n2b)
	g.AddEdge(&n2a, &n2b)

	n3a := Node{"GCATG"}
	n3b := Node{"CATGC"}

	g.AddNode(&n3a)
	g.AddNode(&n3b)
	g.AddEdge(&n3a, &n3b)

	n4a := Node{"GGCAT"}
	n4b := Node{"GCATG"}

	g.AddNode(&n4a)
	g.AddNode(&n4b)
	g.AddEdge(&n4a, &n4b)

	return g
}

func TestDatastructureDirGraph(t *testing.T) {
	g := fillGraph()
	s := g.String()
	gold := "AGGCA -> GGCAT\nCATGC -> ATGCG\nGCATG -> CATGC\nGGCAT -> GCATG"
	if s != gold {
		msg := "Error: DirGraph data structure did not print properly\n"
		msg += fmt.Sprintf("computed:\n%v\n\n", s)
		msg += fmt.Sprintf("gold:\n%v\n\n", gold)
		t.Error(msg)
	}
}
