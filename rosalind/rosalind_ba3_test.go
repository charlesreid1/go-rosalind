package rosalind

import (
	"fmt"
	"sort"
	"testing"
)

/////////////////////////////////
// BA3a Test

func TestKmerComposition(t *testing.T) {
	k := 5
	input := "CAATCCAAC"
	gold := []string{"AATCC", "ATCCA", "CAATC", "CCAAC", "TCCAA"}

	results, err := KmerComposition(input, k)
	if err != nil {
		t.Error(fmt.Sprintf("Error: %v", err))
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(results)

	if !EqualStringSlices(results, gold) {
		msg := fmt.Sprintf("Error testing KmerComposition()\ncomputed = %v\ngold = %v",
			results, gold)
		t.Error(msg)
	}
}
