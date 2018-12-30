package main

import (
	"fmt"
	"testing"
)

// TODO: add a test that loads from a file in data/

func TestMatrixApproximateOccurrences(t *testing.T) {
	var tests = []struct {
		pattern string
		text    string
		d       int
		gold    []int
	}{
		{"ATTCTGGA",
			"CGCCCGAATCCAGAACGCATTCCCATATTTCGGGACCACTGGCCTCCACGGTACGGACGTCAATCAAATGCCTAGCGGCTTGTGGTTTCTCCTACGCTCC",
			3,
			[]int{6, 7, 26, 27, 78}},
		{"AAA",
			"TTTTTTAAATTTTAAATTTTTT",
			2,
			[]int{4, 5, 6, 7, 8, 11, 12, 13, 14, 15}},
		{"GAGCGCTGG",
			"GAGCGCTGGGTTAACTCGCTACTTCCCGACGAGCGCTGTGGCGCAAATTGGCGATGAAACTGCAGAGAGAACTGGTCATCCAACTGAATTCTCCCCGCTATCGCATTTTGATGCGCGCCGCGTCGATT",
			2,
			[]int{0, 30, 66}},
		{"AATCCTTTCA",
			"CCAAATCCCCTCATGGCATGCATTCCCGCAGTATTTAATCCTTTCATTCTGCATATAAGTAGTGAAGGTATAGAAACCCGTTCAAGCCCGCAGCGGTAAAACCGAGAACCATGATGAATGCACGGCGATTGCGCCATAATCCAAACA",
			3,
			[]int{3, 36, 74, 137}},
		{"CCGTCATCC",
			"CCGTCATCCGTCATCCTCGCCACGTTGGCATGCATTCCGTCATCCCGTCAGGCATACTTCTGCATATAAGTACAAACATCCGTCATGTCAAAGGGAGCCCGCAGCGGTAAAACCGAGAACCATGATGAATGCACGGCGATTGC",
			3,
			[]int{0, 7, 36, 44, 48, 72, 79, 112}},
		{"TTT",
			"AAAAAA",
			3,
			[]int{0, 1, 2, 3}},
		{"CCA",
			"CCACCT",
			0,
			[]int{0}},
	}
	for _, test := range tests {
		result, err := FindApproximateOccurrences(test.pattern, test.text, test.d)
		if err != nil {
			t.Error(err)
		}
		if !EqualIntSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing FindApproximateOccurrences:\ncomputed = %v\ngold     = %v",
				result, test.gold)
			t.Error(err)
		}
	}
}
