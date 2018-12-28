package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestNumberToPattern(t *testing.T) {
	n := 11
	k := 3
	gold := "AGT"

	// Money shot
	pattern, err := NumberToPattern(n, k)

	// Check if there was error
	if err != nil {
		t.Error(err)
	}

	if pattern != gold {
		msg := fmt.Sprintf("Error testing NumberToPattern():\nn = %d, k = %d\ncomputed = %v\ngold     = %v\n",
			n, k, pattern, gold)
		t.Error(msg)
	}
}

func TestNumberToPatternFile(t *testing.T) {

	filename := "data/number_to_pattern.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	n_str := lines[1]
	k_str := lines[2]
	// lines[3]: Output
	gold := lines[4]

	n, err := strconv.Atoi(n_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for integer n representing DNA string: %v", err)
	}

	k, err := strconv.Atoi(k_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for integer k kmer length: %v", err)
	}

	pattern, err := NumberToPattern(n, k)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// These will only be unequal if something went wrong
	if pattern != gold {
		err := fmt.Sprintf("Error testing NumberToPattern():\nn = %d, k = %d\ncomputed = %v\ngold     = %v\n",
			n, k, pattern, gold)
		t.Error(err)
	}
}
