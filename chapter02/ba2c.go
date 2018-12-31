package rosalindchapter2

import (
	"fmt"
	"log"

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
