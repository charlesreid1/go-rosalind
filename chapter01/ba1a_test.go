package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

// To run this test:
//
// $ go test -v -run TestPatternCount

// Run a single test of the PatternCount function
func TestPatternCount(t *testing.T) {
	// Call the PatternCount function
	input := "GCGCG"
	pattern := "GCG"
	result := PatternCount(input, pattern)
	gold := 2
	if result != gold {
		err := fmt.Sprintf("Error testing PatternCount(): input = %s, pattern = %s, result = %d (should be %d)",
			input, pattern, result, gold)
		t.Error(err)
	}
}

// Run a test matrix of the PatternCount function
func TestMatrixPatternCount(t *testing.T) {
	// Construct a test matrix
	var tests = []struct {
		input   string
		pattern string
		gold    int
	}{
		{"GCGCG", "GCG", 2},
		{"GAGGGGGGGAG", "AGG", 1},
		{"GCACGCACGCAC", "GCAC", 3},
		{"", "GC", 0},
		{"GCG", "GTACTCTC", 0},
		{"ACGTACGTACGT", "CG", 3},
		{"AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAGACAGATAATAATTACAGAGTACACAACATCCA",
			"AAA", 4},
		{"AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCCAGCGCGTTAACCACGTTCTGTGCCGACTTT",
			"TTT", 4},
		{"GGACTTACTGACGTACG", "ACT", 2},
		{"ATCCGATCCCATGCCCATG", "CC", 5},
		{"CTGTTTTTGATCCATGATATGTTATCTCTCCGTCATCAGAAGAACAGTGACGGATCGCCCTCTCTCTTGGTCAGGCGACCGTTTGCCATAATGCCCATGCTTTCCAGCCAGCTCTCAAACTCCGGTGACTCGCGCAGGTTGAGT",
			"CTC", 9},
	}
	for _, test := range tests {
		result := PatternCount(test.input, test.pattern)
		if result != test.gold {
			err := fmt.Sprintf("Error testing PatternCount(): input = %s, pattern = %s, result = %d (should be %d)",
				test.input, test.pattern, result, test.gold)
			t.Error(err)
		}
	}
}

// Load a PatternCount test (input and output)
// from a file. Run the test with the input
// and verify the output matches the output
// contained in the file.
func TestPatternCountFile(t *testing.T) {

	filename := "data/pattern_count.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	input := lines[1]
	pattern := lines[2]

	// lines[3]: Output
	output_str := lines[4]

	// Convert output to inteter
	output, err := strconv.Atoi(output_str)
	if err != nil {
		t.Error(err)
	}

	// Call the function with the given inputs
	result := PatternCount(input, pattern)

	// Verify answer
	if result != output {
		err := fmt.Sprintf("Error testing PatternCount using test case from file: results do not match:\rcomputed result = %d\nexpected output = %d", result, output)
		t.Error(err)
	}
}
