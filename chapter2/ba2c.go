package rosalindchapter2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA2c: Find a Profile-most Probable k-mer in a String
func BA2cDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA2c:",
		"Find a Profile-most Probable k-mer in a String",
		"",
		"Given a profile matrix, find the most probable k-mer to generate the given DNA string.",
		"",
		"URL: http://rosalind.info/problems/ba2c/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA2c(filename string) {

	BA2cDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	// Input file contents
	dna := lines[0]
	k_str := lines[1]
	k, _ := strconv.Atoi(k_str)

	// To make multidimensional slice,
	// make a slice, then loop and make more slices
	profile, _ := rosa.ReadMatrix32(lines[2:6], k)

	// Find the most probable kmer
	result, _ := rosa.ProfileMostProbableKmers(dna, k, profile)
	fmt.Println(strings.Join(result, " "))
}
