package rosalindchapter1

import (
	"fmt"
	"log"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1c: Find the Reverse Complement of a String

// Describe the problem
func BA1cDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1c:",
		"Find the Reverse Complement of a String",
		"",
		"Given a DNA input string,",
		"find the reverse complement",
		"of the DNA string.",
		"",
		"URL: http://rosalind.info/problems/ba1c/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1c(filename string) {

	BA1cDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("Error: rosa.ReadLines: %v", err)
	}

	// Input file contents
	input := lines[0]

	result, _ := rosa.ReverseComplement(input)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(result)
}
