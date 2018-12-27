package main

import (
	"fmt"
	"sort"
	"testing"
)

// TODO: add a test loaded from a file in data/

func TestMatrixMinSkewPosition(t *testing.T) {
	var tests = []struct {
		genome string
		gold   []int
	}{
		{"CCTATCGGTGGATTAGCATGTCCCTGTACGTTTCGCCGCGAACTAGTTCACACGGCTTGATGGCAAATGGTTTTTCCGGCGACCGTAATCGTCCACCGAG",
			[]int{53, 97}},
		{"TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGA",
			[]int{11, 24}},
		{"ACCG",
			[]int{3}},
		{"ACCC",
			[]int{4}},
		{"CCGGGT",
			[]int{2}},
		{"CCGGCCGG",
			[]int{2, 6}},
	}
	for _, test := range tests {

		// Do it - find the positions that minimize skew
		result, err := MinSkewPositions(test.genome)
		if err != nil {
			t.Error(err)
		}

		// Check length of result
		if len(result) != len(test.gold) {
			err := fmt.Sprintf("Error testing MinSkewPositions():\nfor genome: %s\nlength of result (%d) did not match length of gold standard (%d).\nFound: %v\nShould be: %v",
				test.genome, len(result), len(test.gold),
				result, test.gold)
			t.Error(err)
		}

		// Sort before comparing
		sort.Ints(result)
		sort.Ints(test.gold)
		if !EqualIntSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing MinSkewPositions():\nfor genome: %s\nfound: %v\nshould be: %v",
				test.genome, result, test.gold)
			t.Error(err)
		}
	}
}
