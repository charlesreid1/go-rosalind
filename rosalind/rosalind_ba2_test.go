package rosalind

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
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
	results, err := KeySetIntersection(mslice)
	if err != nil {
		t.Error(fmt.Sprintf("Error: KeySetIntersection() returned error: %v", err))
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(results)

	if !EqualStringSlices(results, gold) {
		msg := fmt.Sprintf("Error testing KeySetIntersection()\ncomputed = %v\ngold = %v",
			results, gold)
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

func TestFindMotifsFile(t *testing.T) {
	filename := "data/motif_enumeration.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// Input file contents
	// lines[0]: Input
	params := strings.Split(lines[1], " ")
	k, _ := strconv.Atoi(params[0])
	d, _ := strconv.Atoi(params[1])

	// lines[-2]: Output
	// lines[-1]: gold standard
	gold := strings.Split(lines[len(lines)-1], " ")

	// This requires some trickery.

	// 4 lines in the input file are for
	// input/parameters/output/gold standard.
	// The rest of the lines are DNA strings.

	// Make space for DNA strings
	dna := make([]string, len(lines)-4)
	iLstart := 2
	iLend := len(lines) - 2
	// Two counters:
	// one for the line index (iL),
	// one for the array index (iA).
	for iA, iL := 0, iLstart; iL < iLend; iA, iL = iA+1, iL+1 {
		dna[iA] = lines[iL]
	}

	// Money shot
	results, _ := FindMotifs(dna, k, d)

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(results)

	if !EqualStringSlices(results, gold) {
		msg := fmt.Sprintf("Error testing FindMotifs()\ncomputed = %v\ngold = %v",
			results, gold)
		t.Error(msg)
	}
}
