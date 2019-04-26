package rosalindchapter3

import (
	"fmt"
	"log"
	"strconv"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA3d: Construct the DeBruijn graph of a string
func BA3dDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA3d:",
		"Construct the DeBruijn graph of a string",
		"",
		"Given a collection of arbitrary kmers, construct the DeBruijn graph of all overlapping (k-1)mers",
		"",
		"URL: http://rosalind.info/problems/ba3d/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA3d(filename string) {

	BA3dDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	// Input file contents
	k, _ := strconv.Atoi(lines[0])
	dna := lines[1]
	dbg, err := rosa.ConstructDeBruijnGraph(dna, k)
	if err != nil {
		log.Fatalf("Error constructing DeBruijn graph: %v", err)
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
