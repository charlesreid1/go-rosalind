package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestVisitNeighborhood(t *testing.T) {
	input := "ACG"
	d := 1
	gold := []string{"CCG", "TCG", "GCG", "AAG", "ATG", "AGG", "ACA", "ACC", "ACT", "ACG"}

	// Money shot
	result, err := VisitHammingNeighbors(input, d)

	// Check if there was error
	if err != nil {
		t.Error(err)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	if !EqualStringSlices(gold, result) {
		msg := fmt.Sprintf("Error testing VisitHammingNeighbors():\ninput = %s, = %d\ncomputed = %v\ngold     = %v\n",
			input, d,
			result, gold)
		t.Error(msg)
	}
}

func TestVisitNeighborhoodFile(t *testing.T) {
	filename := "data/neighbors.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	input := strings.TrimSpace(lines[1])
	d_str := strings.TrimSpace(lines[2])
	// lines[3]: Output

	d, err := strconv.Atoi(d_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for parameter d: %v", err)
	}

	gold := make([]string, len(lines)-4)
	for i := 4; i < len(lines); i++ {
		j := i - 4
		gold[j] = lines[i]
	}

	result, err := VisitHammingNeighbors(input, d)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// Check that there _was_ a result
	if len(result) == 0 {
		msg := fmt.Sprintf("Error testing VisitHammingNeighbors() using test case from file: no results")
		t.Error(msg)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	// These will only be unequal if something went wrong
	if !EqualStringSlices(gold, result) {
		msg := fmt.Sprintf("Error testing VisitHammingNeighbors() using test case from file: final results do not match.\ncomputed = %q\ngold     = %q\n",
			result, gold)
		t.Error(msg)
	}
}
