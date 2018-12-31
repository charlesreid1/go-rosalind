package rosalindchapter2

import (
	"fmt"
	"log"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA2a: Implement Motif Enumeration
func BA2aDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA2a:",
		"Implement Motif Enumeration",
		"",
		"Given a collection of strings of DNA, find all motifs (kmers of length k and Hamming distance d from all DNA strings).",
		"",
		"URL: http://rosalind.info/problems/ba2a/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA2a(filename string) {

	BA2aDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// Input file contents
	params := lines[0]
	dna := make([]string, len(lines)-1)
	for ip1 := 1; ip1 < len(lines); ip1++ {
		i := ip1 - 1
		dna[i] = lines[ip1]
	}

	results := rosa.MotifEnumeration(input, dna)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(results, " "))
}
