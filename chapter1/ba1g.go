package rosalindchapter1

import (
	"fmt"
	"log"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1g: Find Hamming distance between two DNA strings

// Describe the problem
func BA1gDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1g:",
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
func BA1g(filename string) {

	BA1gDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// Input file contents
	p := lines[0]
	q := lines[1]

	hamm, _ := rosa.HammingDistance(p, q)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(hamm)
}
