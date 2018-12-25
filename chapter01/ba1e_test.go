package main

import (
	"fmt"
	"testing"
)

func TestMatrixFindClumps(t *testing.T) {
	var tests = []struct {
		genome string
		k      int
		L      int
		t      int
		gold   []string
	}{
		{"CGGACTCGACAGATGTGAAGAACGACAATGTGAAGACTCGACACGACAGAGTGAAGAGAAGAGGAAACATTGTAA",
			5, 50, 4,
			[]string{"CGACA", "GAAGA"}},
		{"AAAACGTCGAAAAA",
			2, 4, 2,
			[]string{"AA"}},
		{"ACGTACGT",
			1, 5, 2,
			[]string{"A", "C", "G", "T"}},
		{"CCACGCGGTGTACGCTGCAAAAAGCCTTGCTGAATCAAATAAGGTTCCAGCACATCCTCAATGGTTTCACGTTCTTCGCCAATGGCTGCCGCCAGGTTATCCAGACCTACAGGTCCACCAAAGAACTTATCGATTACCGCCAGCAACAATTTGCGGTCCATATAATCGAAACCTTCAGCATCGACATTCAACATATCCAGCG",
			3, 25, 3,
			[]string{"AAA", "CAG", "CAT", "CCA", "GCC", "TTC"}},
	}
	for _, test := range tests {
		result, err := FindClumps(test.genome,
			test.k, test.L, test.t)
		if err != nil {
			t.Error(err)
		}
		if !EqualStringSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing FindClumps(): k = %d, L = %d, t = %d", test.k, test.L, test.t)
			t.Error(err)
		}
	}
}
