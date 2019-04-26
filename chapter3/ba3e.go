package rosalindchapter3

import (
	"fmt"
	"log"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA3e: Construct the DeBruijn graph of a collection of kmers
func BA3eDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA3e:",
		"Construct the DeBruijn graph of a collection of kmers",
		"",
		"Given a collection of arbitrary kmers, construct the DeBruijn graph of all overlapping (k-1)mers",
		"",
		"URL: http://rosalind.info/problems/ba3e/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA3e(filename string) {

	BA3eDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	// Trim each line and there are your kmers
	for i, line := range lines {
		lines[i] = strings.Trim(line, " ")
	}

	// Make the De Bruijn graph
	dbg, err := rosa.ConstructDeBruijnGraphKmers(lines)
	if err != nil {
		log.Fatalf("Error constructing DeBruijn graph from kmers: %v", err)
	}

	one_edge_per_line := false
	dbgs, err := rosa.SPrintOverlapGraph(dbg, one_edge_per_line)
	if err != nil {
		log.Fatalf("Error printing DeBruijn graph: %v", err)
	}

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(dbgs)
}
