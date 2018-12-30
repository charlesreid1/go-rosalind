package main

import (
	"fmt"
	"log"
)

// Rosalind: Problem BA1L: Pattern to Number

// Describe the problem
func BA1LDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1L:",
		"Pattern to Number",
		"",
		"Given an input kmer of length k, convert it to",
		"an integer corresponding to its lexicographic",
		"order among kmers of length k.",
		"",
		"URL: http://rosalind.info/problems/ba1l/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1L(filename string) {

	BA1LDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// Input file contents
	input := lines[0]

	number, _ := PatternToNumber(input)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(number)
}