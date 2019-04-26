package rosalindchapter3

import (
	"fmt"
	"log"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA3c: Construct the overlap graph of a set of k-mers
func BA3cDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA3c:",
		"Construct the overlap graph of a set of k-mers",
		"",
		"Given a set of overlapping k-mers, construct the overlap graph and print a sorted adjacency matrix",
		"",
		"URL: http://rosalind.info/problems/ba3c/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA3c(filename string) {

	BA3cDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	// Trim each line and there are your contigs
	for i, line := range lines {
		lines[i] = strings.Trim(line, " ")
	}

	og, err := rosa.OverlapGraph(lines)
	if err != nil {
		log.Fatalf("Error when calling OverlapGraph()")
	}

	one_edge_per_line := true
	ogs, err := rosa.SPrintOverlapGraph(og, one_edge_per_line)
	if err != nil {
		log.Fatalf("Error when calling SPrintOverlapGraph()")
	}

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(ogs)
}
