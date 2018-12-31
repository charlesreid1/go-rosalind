package rosalindchapter01

import (
	"fmt"
	"log"
	"strconv"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1K: Generate Frequency Array

// Describe the problem
func BA1KDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1K:",
		"Generate Frequency Array",
		"",
		"Given an integer k, generate the frequency array of",
		"an input string. The frequency array is an array of",
		"counts with one count per index, and integers mapped",
		"to kmers.",
		"",
		"URL: http://rosalind.info/problems/ba1k/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1K(filename string) {

	BA1KDescription()

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
		log.Fatalf("Error: string to int conversion for parameter k: %v", err)
	}

	arr, _ := rosa.FrequencyArray(input, k)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	for _, e := range arr {
		fmt.Print(e, " ")
	}
	//fmt.Println(strings.Join(arr, " "))
}
