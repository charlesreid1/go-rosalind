package main

import (
	"fmt"
	"sort"
	"testing"
)

// TODO: (done) add a test for count
// TODO: (done) add a test for generate
// TODO: add a test formost frequent neighbors histogram
// TODO: add a test from a file

func TestMatrixCountHammingNeighbors(t *testing.T) {
	var tests = []struct {
		n    int // length of input string
		d    int // maximum Hamming distance
		c    int // number of codons
		gold int // gold standard true value
	}{
		{1, 0, 4, 1},
		{1, 1, 4, 4},
		{3, 1, 4, 10},
		{5, 1, 4, 16},
		{3, 2, 4, 37},
		{5, 2, 4, 106},
	}
	for _, test := range tests {
		// Count em up
		result, err := CountHammingNeighbors(test.n, test.d, test.c)

		// Check if there was error
		if err != nil {
			t.Error(err)
		}
		if result != test.gold {
			err := fmt.Sprintf("Error testing CountHammingNeighbors:\ncomputed = %d\ngold     = %d",
				result, test.gold)
			t.Error(err)
		}
	}
}

func TestMatrixVisitHammingNeighbors(t *testing.T) {
	var tests = []struct {
		input string
		d     int
		gold  []string
	}{
		{"AAA", 1,
			[]string{"AAC", "AAT", "AAG", "AAA", "CAA", "GAA", "TAA", "ATA", "ACA", "AGA"},
		},
	}
	for _, test := range tests {

		// Money shot
		result, err := VisitHammingNeighbors(test.input, test.d)

		// Check if there was error
		if err != nil {
			t.Error(err)
		}

		// Sort before comparing
		sort.Strings(test.gold)
		sort.Strings(result)

		if !EqualStringSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing VisitHammingNeighbors:\ncomputed = %v\ngold     = %v",
				result, test.gold)
			t.Error(err)
		}
	}
}
