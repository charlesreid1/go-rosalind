package rosalindchapter2

import (
	"fmt"
	"log"
	"strconv"
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
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	// Input file contents
	params := strings.Split(lines[0], " ")
	k, _ := strconv.Atoi(params[0])
	d, _ := strconv.Atoi(params[1])

	// 1 line in the input file is for
	// parameters/gold standard.
	// The rest of the lines are DNA strings.

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

	results, _ := rosa.FindMotifs(dna, k, d)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(results, " "))
}
