package rosalindstronghold

import (
	"fmt"
	"log"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem DNA: Counting DNA Nucleotides
func DNADescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem DNA:",
		"Counting DNA Nucleotides",
		"",
		"Given a DNA string, return a count of each base pair as an array, in the order A, C, G, T",
		"",
		"URL: http://rosalind.info/problems/dna/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func DNA(filename string) {

	DNADescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// Input file contents
	input := lines[0]
	result, _ := rosa.CountNucleotidesArray(input)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	for _, r := range result {
		fmt.Printf("%d ", r)
	}
	fmt.Printf("\n\n")
}
