package rosalind

import (
	"fmt"
	"testing"
)

func TestCountNucleotides(t *testing.T) {
	input := "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	results, err := CountNucleotidesArray(input)
	if err != nil {
		t.Error(err)
	}
	gold := []int{20, 12, 17, 21}

	if !EqualIntSlices(results, gold) {
		err := fmt.Sprintf("Error testing CountNucleotides(): input = %s\ncomputed = %v\ngold     = %v\n",
			input, results, gold)
		t.Error(err)
	}
}
