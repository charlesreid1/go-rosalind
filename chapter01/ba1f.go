package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Rosalind: Problem BA1F: Find positions in a gene that minimizing skew

// Describe the problem
func BA1FDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1F:",
		"Find positions in a gene that minimize skew",
		"",
		"The skew of a genome is defined as the difference",
		"between the number of C codons and the number of G",
		"codons. Given a DNA string, this function should",
		"compute the cumulative skew for each position in",
		"the genome, and report the indices where the skew",
		"value is minimzed.",
		"",
		"URL: http://rosalind.info/problems/ba1f/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1F(filename string) {

	BA1FDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// Input file contents
	genome := lines[0]

	minskew, _ := MinSkewPositions(genome)

	minskew_str := make([]string, len(minskew))
	for i, j := range minskew {
		minskew_str[i] = strconv.Itoa(j)
	}

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(minskew_str, " "))
}
