package rosalindchapter2

import (
	"fmt"
	"log"

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
		"Given a set of DNA strings, find a k-mer pattern that minimizes the magnitude of the distance from it to the minimum Hamming distance (closest Hamming neighbor) kmer in each DNA string",
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
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	//// Input file contents
	//input := lines[0]
	//params := lines[1]
	//result := rosa.PatternCount(input, pattern)
	// 
	//fmt.Println("")
	//fmt.Printf("Computed result from input file: %s\n", filename)
	//fmt.Println(result)
}
