package rosalindchapter01

import (
	"fmt"
	"log"
	"strconv"
)

// Rosalind: Problem BA1M: Pattern to Number

// Describe the problem
func BA1MDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1M:",
		"Number to Pattern",
		"",
		"Given an integer and a kmer length k, convert",
		"the integer to its corresponding kmer.",
		"",
		"URL: http://rosalind.info/problems/ba1m/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1M(filename string) {

	BA1MDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// Input file contents
	number_str := lines[0]
	k_str := lines[1]

	number, err := strconv.Atoi(number_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for number: %v", err)
	}

	k, err := strconv.Atoi(k_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for k: %v", err)
	}

	result, _ := NumberToPattern(number, k)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(result)
}
