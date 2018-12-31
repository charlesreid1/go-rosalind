package rosalindchapter01

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1j: Most Frequent Words with Mismatches and Reverse Complements

// Describe the problem
func BA1jDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1j:",
		"Most Frequent Words with Mismatches and Reverse Complements",
		"",
		"Given an input string and a maximum allowable",
		"Hamming distance d, report the most frequent",
		"kmer that either occurs or whose Hamming neighbors",
		"occur most frequently in the input string and in the",
		"reverse complement of the input string.",
		"",
		"URL: http://rosalind.info/problems/ba1j/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1j(filename string) {

	BA1jDescription()

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

	mfks_mis, _ := rosa.MostFrequentKmersMismatchesRevComp(input, k, d)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(mfks_mis, " "))
}
