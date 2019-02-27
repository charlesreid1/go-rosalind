package rosalindchapter2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA2e: Implement GreedyMotifSearch with Pseudocounts
func BA2eDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA2e:",
		"Implement GreedyMotifSearch with Pseudocounts",
		"",
		"Re-implement problem BA2d (greedy motif search) using pseudocounts, which avoid setting probabilities to an absolute value of zero.",
		"",
		"URL: http://rosalind.info/problems/ba2e/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA2e(filename string) {

	BA2eDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	//// Input file contents
	params := strings.Split(lines[0], " ")
	k, _ := strconv.Atoi(params[0])
	t, _ := strconv.Atoi(params[1])

	// 1 line in the input file is for
	// parameters.
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

	result, _ := rosa.GreedyMotifSearchPseudocounts(dna, k, t)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(result, " "))
}
