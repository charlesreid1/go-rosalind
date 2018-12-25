package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"
)

// Run a test of the MostFrequentKmers function
func TestMostFrequentKmers(t *testing.T) {
	// Call MostFrequentKmers
	input := "AAAATGCGCTAGTAAAAGTCACTGAAAA"
	k := 4
	result, err := MostFrequentKmers(input, k)
	gold := []string{"AAAA"}

	if err != nil {
		t.Error(err)
	}

	if !EqualStringSlices(result, gold) {
		err := fmt.Sprintf("Error testing MostFrequentKmers(): input = %s, k = %d, result = %s (should be %s)",
			input, k, result, gold)
		t.Error(err)
	}
}

// Run a test of the PatternCount function
// using inputs/outputs from a file.
func TestMostFrequentKmersFile(t *testing.T) {

	filename := "data/frequent_words.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	dna := lines[1]
	k_str := lines[2]
	// lines[3]: Output
	gold := strings.Split(lines[4], " ")

	// Convert k to integer
	k, err := strconv.Atoi(k_str)
	if err != nil {
		t.Error(err)
	}

	// Call the function with the given inputs
	result, err := MostFrequentKmers(dna, k)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// Check that there _was_ a result
	if len(result) == 0 {
		err := fmt.Sprintf("Error testing MostFrequentKmers using test case from file: length of most frequent kmers found was 0: %q",
			result)
		t.Error(err)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	// These will only be unequal if something went wrong
	if !EqualStringSlices(gold, result) {
		err := fmt.Sprintf("Error testing MostFrequentKmers using test case from file: most frequent kmers mismatch.\ncomputed = %q\ngold     = %q\n",
			result, gold)
		t.Error(err)
	}
}
