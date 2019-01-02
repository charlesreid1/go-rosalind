package rosalindchapter2

import (
	"fmt"
	"log"
	"strconv"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA2b: Find a Median String
func BA2bDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA2b:",
		"Find a Median String",
		"",
		"Given a kmer length k and a set of strings of DNA, find the kmer(s) that minimize the L1 norm of the distance from it to all other DNA strings.",
		"",
		"URL: http://rosalind.info/problems/ba2b/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA2b(filename string) {

	BA2bDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	// Input file contents
	k_str := lines[0]
	k, _ := strconv.Atoi(k_str)

	// Make space for DNA strings
	dna := make([]string, len(lines)-1)
	iLstart := 1
	iLend := len(lines)
	// Two counters:
	// one for the line index (iL),
	// one for the array index (iA).
	for iA, iL := 0, iLstart; iL < iLend; iA, iL = iA+1, iL+1 {
		dna[iA] = lines[iL]
	}

	results, _ := rosa.MedianString(dna, k)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(results)
}
