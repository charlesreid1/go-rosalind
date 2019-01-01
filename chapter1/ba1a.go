package rosalindchapter1

import (
	"fmt"
	"log"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1a: Most Frequent k-mers

// Describe the problem
func BA1aDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1a:",
		"Most Frequest k-mers",
		"",
		"Given an input string and a length k,",
		"report the k-mer or k-mers that occur",
		"most frequently.",
		"",
		"URL: http://rosalind.info/problems/ba1a/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem,
// print the name of the input file,
// print the output/result
func BA1a(filename string) {

	BA1aDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// Input file contents
	var input, pattern string
	input = lines[0]
	pattern = lines[1]

	result := rosa.PatternCount(input, pattern)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(result)
}
