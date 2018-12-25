package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestFindOccurrences(t *testing.T) {
	// Call FindOccurrences
	pattern := "ATAT"
	genome := "GATATATGCATATACTT"

	result, err := FindOccurrences(pattern, genome)
	gold := []int{1, 3, 9}

	if !EqualIntSlices(result, gold) || err != nil {
		err := fmt.Sprintf("Error testing FindOccurrences(): result = %q, should be %q",
			result, gold)
		t.Error(err)
	}
}

func TestFindOccurrencesDebug(t *testing.T) {
	// Construct a test matrix
	var tests = []struct {
		pattern string
		genome  string
		gold    []int
	}{
		{"ACAC", "TTTTACACTTTTTTGTGTAAAAA",
			[]int{4}},
		{"AAA", "AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAGACAGATAATAATTACAGAGTACACAACATCCAT",
			[]int{0, 46, 51, 74}},
		{"TTT", "AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCCAGCGCGTTAACCACGTTCTGTGCCGACTTT",
			[]int{88, 92, 98, 132}},
		{"ATA", "ATATATA",
			[]int{0, 2, 4}},
	}
	for _, test := range tests {

		result, err := FindOccurrences(test.pattern, test.genome)

		if err != nil {
			t.Error(err)
		}

		if !EqualIntSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing FindOccurrences(): result = %q, should be %q",
				result, test.gold)
			t.Error(err)
		}
	}
}

func TestFindOccurrencesFiles(t *testing.T) {

	filename := "data/pattern_matching.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// lines[0]: Input
	pattern := lines[1]
	genome := lines[2]

	// lines[3]: Output
	gold_str := lines[4]
	gold_slice := strings.Split(gold_str, " ")

	gold := make([]int, len(gold_slice))
	for i, g := range gold_slice {
		gold[i], err = strconv.Atoi(g)
		if err != nil {
			t.Error(err)
		}
	}

	result, err := FindOccurrences(pattern, genome)

	if err != nil {
		t.Error(err)
	}

	if !EqualIntSlices(result, gold) {
		err := fmt.Sprintf("Error testing FindOccurrences():\nresult = %v\ngold   = %v\n",
			result, gold)
		t.Error(err)
	}
}
