package rosalindchapter1

import (
	"fmt"
	"log"
	"strconv"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1n: Calculating d-Neighborhood of String

// Describe the problem
func BA1nDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1n:",
		"Calculating d-Neighborhood of String",
		"",
		"Given an input string of DNA and a Hamming",
		"distance d, compute all DNA strings that",
		"are a Hamming distance of up to d away.",
		"",
		"URL: http://rosalind.info/problems/ba1n/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1n(filename string) {

	BA1nDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("Error: rosa.ReadLines: %v", err)
	}

	// Input file contents
	input := lines[0]
	d_str := lines[1]

	d, err := strconv.Atoi(d_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for d: %v", err)
	}

	result, _ := rosa.VisitHammingNeighbors(input, d)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	for _, j := range result {
		fmt.Println(j)
	}
}
