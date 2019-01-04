package rosalindchapter3

import (
	"fmt"
	"log"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA3b: Reconstruct string from genome path
func BA3bDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA3b:",
		"Reconstruct string from genome path",
		"",
		"Reconstruct a string from its genome path, i.e., sequential fragments of overlapping DNA.",
		"",
		"URL: http://rosalind.info/problems/ba3b/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA3b(filename string) {

	BA3bDescription()

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

	genome, err := rosa.ReconstructGenomeFromPath(lines)
	if err != nil {
		log.Fatalf("Error when calling ReconstructGenomeFromPath()")
	}

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(genome)
}
