package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Rosalind: Problem BA1i: Most Frequent Words with Mismatches

// Describe the problem
func BA1iDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1i:",
		"Most Frequent Words with Mismatches",
		"",
		"Given an input string and a maximum allowable",
		"Hamming distance d, report the most frequent",
		"kmer that either occurs or whose Hamming neighbors",
		"occur most frequently.",
		"",
		"URL: http://rosalind.info/problems/ba1i/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1i(filename string) {

	BA1iDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// Input file contents
	input := lines[0]
	params := strings.Split(lines[1], " ")
	if len(params) < 1 {
		log.Fatalf("Error splitting second line: only found 0-1 tokens")
	}

	k_str, d_str := params[0], params[1]

	k, err := strconv.Atoi(k_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for parameter k: %v", err)
	}

	d, err := strconv.Atoi(d_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for parameter d: %v", err)
	}

	mfks_mis, _ := MostFrequentKmersMismatches(input, k, d)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(mfks_mis, " "))
}
