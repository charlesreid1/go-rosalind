package rosalindchapter01

import (
	"fmt"
	"log"
)

// Rosalind: Problem BA1G: Find Hamming distance between two DNA strings

// Describe the problem
func BA1GDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1G:",
		"Find Hamming distance between two DNA strings",
		"",
		"The Hamming distance between two strings HammingDistance(p,q)",
		"is the number of characters different between the two",
		"strands. This program computes the Hamming distance",
		"between two strings.",
		"",
		"URL: http://rosalind.info/problems/ba1g/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1G(filename string) {

	BA1GDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// Input file contents
	p := lines[0]
	q := lines[1]

	hamm, _ := HammingDistance(p, q)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(hamm)
}
