package rosalindchapter1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1b: Most Frequent k-mers

// Describe the problem
func BA1bDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1b:",
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
func BA1b(filename string) {

	BA1bDescription()

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

	mfks, _ := rosa.MostFrequentKmers(input, k)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(mfks, " "))
}
