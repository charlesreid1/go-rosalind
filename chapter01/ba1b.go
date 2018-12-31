package rosalindchapter01

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Rosalind: Problem BA1B: Most Frequent k-mers

// Describe the problem
func BA1BDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1B:",
		"Most Frequest k-mers",
		"",
		"Given an input string and a length k,",
		"report the k-mer or k-mers that occur",
		"most frequently.",
		"",
		"URL: http://rosalind.info/problems/ba1b/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1B(filename string) {

	BA1BDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// Input file contents
	input := lines[0]
	k_str := lines[1]

	k, err := strconv.Atoi(k_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion: %v", err)
	}

	mfks, _ := MostFrequentKmers(input, k)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(mfks, " "))
}
