package rosalind

import (
	"fmt"
	"sort"
	"testing"
)

/////////////////////////////////
// BA2A Test

func TestKeySetIntersection(t *testing.T) {
	gold := []string{"AAA", "BBB"}
	m1 := map[string]int{
		"AAA": 1,
		"BBB": 2,
		"CCC": 2,
		"DDD": 2,
	}
	m2 := map[string]int{
		"AAA": 2,
		"BBB": 3,
		"EEE": 3,
		"FFF": 3,
	}
	m3 := map[string]int{
		"AAA": 3,
		"BBB": 4,
		"GGG": 4,
		"HHH": 4,
	}
	mslice := make([]map[string]int, 3)
	mslice[0] = m1
	mslice[1] = m2
	mslice[2] = m3
	result, err := KeySetIntersection(mslice)
	if err != nil {
		t.Error(fmt.Sprintf("Error: KeySetIntersection() returned error: %v", err))
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	if !EqualStringSlices(result, gold) {
		msg := fmt.Sprintf("Error testing KeySetIntersection\ncomputed = %v\ngold = %v",
			result, gold)
		t.Error(msg)
	}
}

func TestFindMotifs(t *testing.T) {
	k := 3
	d := 1
	dna := []string{"ATTTGGC", "TGCCTTA", "CGGTATC", "GAAAATT"}

	results, err := FindMotifs(dna, k, d)
	if err != nil {
		t.Error(fmt.Sprintf("Error: FindMotifs() returned error: %v", err))
	}
	gold := []string{"ATA", "ATT", "GTT", "TTT"}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(results)

	if !EqualStringSlices(results, gold) {
		msg := fmt.Sprintf("Error testing FindMotifs():\ncomputed = %v\ngold = %v",
			results, gold)
		t.Error(msg)
	}
}
