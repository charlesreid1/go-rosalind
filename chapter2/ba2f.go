package rosalindchapter2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA2f: Implement RandomizedMotifSearch with Pseudocounts
func BA2fDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA2f:",
		"Implement RandomizedMotifSearch with Pseudocounts",
		"",
		"Re-implement problem BA2e (greedy motif search with pseudocounts) but use a random, instead of greedy, algorithm to pick motif kmers from each DNA string.",
		"",
		"URL: http://rosalind.info/problems/ba2f/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA2f(filename string) {

	BA2fDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	// Input file contents
	params := strings.Split(lines[0], " ")
	k, _ := strconv.Atoi(params[0])
	t, _ := strconv.Atoi(params[1])

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

	result, _ := rosa.RandomMotifSearchPseudocounts(dna, k, t)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(result, " "))
}
