package rosalindchapter3

import (
	"fmt"
	"log"
	"strconv"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA3a: Generate k-mer Composition of a String
func BA3aDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA3a:",
		"Generate k-mer Composition of a String",
		"",
		"Given an input string, generate a list of all kmers that are in the input string.",
		"",
		"URL: http://rosalind.info/problems/ba3a/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA3a(filename string) {

	BA3aDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	// Input file contents
	k_str := lines[0]
	k, err := strconv.Atoi(k_str)
	if err != nil {
		msg := fmt.Sprintf("Error: string to int conversion failed for %s\n",
			k_str)
		log.Fatalf(msg)
	}

	input := lines[1]

	result, _ := rosa.KmerComposition(input, k)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	for _, kmer := range result {
		fmt.Printf("%s\n", kmer)
	}
	fmt.Printf("\n")
}
