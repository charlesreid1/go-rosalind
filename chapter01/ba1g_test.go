package main

import (
	"fmt"
	"testing"
)

// TODO: add a test that loads from a file in data/

func TestMatrixHammingDistance(t *testing.T) {
	var tests = []struct {
		p    string
		q    string
		dist int
	}{
		{"GGGCCGTTGGT",
			"GGACCGTTGAC",
			3},
		{"AAAA",
			"TTTT",
			4},
		{"ACGTACGT",
			"TACGTACG",
			8},
		{"ACGTACGT",
			"CCCCCCCC",
			6},
		{"ACGTACGT",
			"TGCATGCA",
			8},
		{"GATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAAATTTTATTGACTTAGGTCACTAAATAC",
			"AATAGCAGCTTCTCAACTGGTTACCTCGTATGAGTAAATTAGGTCATTATTGACTCAGGTCACTAACGTC",
			15},
		{"AGAAACAGACCGCTATGTTCAACGATTTGTTTTATCTCGTCACCGGGATATTGCGGCCACTCATCGGTCAGTTGATTACGCAGGGCGTAAATCGCCAGAATCAGGCTG",
			"AGAAACCCACCGCTAAAAACAACGATTTGCGTAGTCAGGTCACCGGGATATTGCGGCCACTAAGGCCTTGGATGATTACGCAGAACGTATTGACCCAGAATCAGGCTC",
			28},
	}
	for _, test := range tests {
		result, err := HammingDistance(test.p, test.q)
		if err != nil {
			t.Error(err)
		}
		if result != test.dist {
			err := fmt.Sprintf("Error testing HammingDistance(): computed dist = %d (should be %d)\np = %s\nq = %s\n",
				result, test.dist,
				test.p, test.q)
			t.Error(err)
		}
	}
}
