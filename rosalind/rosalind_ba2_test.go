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
// BA2a Test

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

// Test the FindMotifs function using a single problem.
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

// Test the FindMotifs function using a test matrix
// of debug cases.
func TestMatrixFindMotifs(t *testing.T) {
	var tests = []struct {
		k    int
		d    int
		dna  []string
		gold []string
	}{
		{3, 1,
			[]string{"ATTTGGC", "TGCCTTA", "CGGTATC", "GAAAATT"},
			[]string{"ATA", "ATT", "GTT", "TTT"},
		},
		{3, 0,
			[]string{"ACGT", "ACGT", "ACGT"},
			[]string{"ACG", "CGT"},
		},
		{3, 1,
			[]string{"AAAAA", "AAAAA", "AAAAA"},
			[]string{"AAA", "AAC", "AAG", "AAT", "ACA", "AGA", "ATA", "CAA", "GAA", "TAA"},
		},
		{3, 3,
			[]string{"AAAAA", "AAAAA", "AAAAA"},
			[]string{"AAA", "AAC", "AAG", "AAT", "ACA", "ACC", "ACG", "ACT", "AGA", "AGC", "AGG", "AGT", "ATA", "ATC", "ATG", "ATT", "CAA", "CAC", "CAG", "CAT", "CCA", "CCC", "CCG", "CCT", "CGA", "CGC", "CGG", "CGT", "CTA", "CTC", "CTG", "CTT", "GAA", "GAC", "GAG", "GAT", "GCA", "GCC", "GCG", "GCT", "GGA", "GGC", "GGG", "GGT", "GTA", "GTC", "GTG", "GTT", "TAA", "TAC", "TAG", "TAT", "TCA", "TCC", "TCG", "TCT", "TGA", "TGC", "TGG", "TGT", "TTA", "TTC", "TTG", "TTT"},
		},
		{3, 0,
			[]string{"AAAAA", "AAAAA", "AACAA"},
			[]string{},
		},
		{3, 0,
			[]string{"AACAA", "AAAAA", "AAAAA"},
			[]string{},
		},
	}
	for _, test := range tests {

		// Money shot
		results, err := FindMotifs(test.dna, test.k, test.d)

		if err != nil {
			t.Error(err)
		}

		// Sort before comparing
		sort.Strings(test.gold)
		sort.Strings(results)

		if !EqualStringSlices(results, test.gold) {
			msg := fmt.Sprintf("Error testing FindMotifs()\nk = %d, d = %d, len(dna) = %d\ncomputed = %v\ngold = %v",
				test.k, test.d, len(test.dna),
				results, test.gold)
			t.Error(msg)
		}
	}
}

// Test the FindMotifs function using a large
// test case loaded from a file.
func TestFindMotifsFile(t *testing.T) {

	filename := "data/motif_enumeration.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := ReadLines(filename)
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
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
	results, err := FindMotifs(dna, k, d)

	if err != nil {
		t.Error(err)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(results)

	if !EqualStringSlices(results, gold) {
		msg := fmt.Sprintf("Error testing FindMotifs()\ncomputed = %v\ngold = %v",
			results, gold)
		t.Error(msg)
	}
}

/////////////////////////////////
// BA2b Test

// Test the MinKmerDistance function.
func TestMatrixMinKmerDistance(t *testing.T) {
	var tests = []struct {
		pattern string
		text    string
		d       int
	}{
		{"ATA", "AAATTGACGCAT", 1},
		{"AAA", "AAAAAAAAAAA", 0},
		{"AAA", "CCCCCCCCC", 3},
		{"AAA", "GAAGAAGAAGAA", 1},
		{"AAAA", "GAAG", 2},
		{"AAAA", "GAAGAA", 1},
	}
	for _, test := range tests {

		// Money shot
		c, err := MinKmerDistance(test.pattern, test.text)
		if err != nil {
			t.Error(err)
		}
		if c != test.d {
			msg := fmt.Sprintf("Error testing MinKmerDistance()\npattern = %s, text = %s\ncomputed = %d\ngold = %d",
				test.pattern, test.text,
				c, test.d)
			t.Error(msg)
		}

	}
}

// Test the MinKmerDistances function.
func TestMatrixMinKmerDistances(t *testing.T) {
	var tests = []struct {
		pattern string
		inputs  []string
		d       int
	}{
		{
			"AAA",
			[]string{"AAAA", "CCCC", "GGGG", "TTTT"},
			9},
		{
			"AAA",
			[]string{"GAAG", "CAAC", "TAAG", "TAAC"},
			4},
	}
	for _, test := range tests {

		// Money shot
		c, err := MinKmerDistances(test.pattern, test.inputs)
		if err != nil {
			t.Error(err)
		}
		if c != test.d {
			msg := fmt.Sprintf("Error testing MinKmerDistance()\npattern = %s, inputs = %v\ncomputed = %d\ngold = %d",
				test.pattern, test.inputs,
				c, test.d)
			t.Error(msg)
		}

	}
}

// Test MedianString
func TestMedianString(t *testing.T) {
	k := 3
	dna := []string{
		"AAATTGACGCAT",
		"GACGACCACGTT",
		"CGTCAGCGCCTG",
		"GCTGAGCACCGG",
		"AGTACGGGACAG",
	}
	result, _ := MedianString(dna, k)

	gold := "GAC"

	// Since they only report one kmer, and we report all,
	// we should check if their kmer is in our slice.
	var passed_test bool
	for _, r := range result {
		if r == gold {
			passed_test = true
			break
		}
	}
	if !passed_test {
		// Uh oh, their kmer is not in our slice.
		msg := fmt.Sprintf("Error testing MostFrequentKmers using test case from file: most frequent kmers in gold not in results.\ncomputed = %q\ngold     = %q\n",
			result, gold)
		t.Error(msg)
	}
}

/////////////////////////////////
// BA2c Test

func TestProfileMostProbableKmers(t *testing.T) {
	gold := "CCGAG"
	dna := "ACCTGTTTATTGCCTAAGTTCCGAACAAACCCAATATAGCCCGAGGGCCT"
	k := 5
	prof := [][]float32{
		[]float32{0.2, 0.2, 0.3, 0.2, 0.3},
		[]float32{0.4, 0.3, 0.1, 0.5, 0.1},
		[]float32{0.3, 0.3, 0.5, 0.2, 0.4},
		[]float32{0.1, 0.2, 0.1, 0.1, 0.2},
	}
	result, _ := ProfileMostProbableKmers(dna, k, prof)

	// Check if gold answer is in our results slice
	var passed_test bool
	for _, r := range result {
		if r == gold {
			passed_test = true
			break
		}
	}

	if !passed_test {
		// The correct kmer was not found in our result
		msg := fmt.Sprintf("Error testing ProfileMostProbableKmer(): found incorrect most probable kmer:\n    Gold: %s\n    Computed: %s\n",
			gold, strings.Join(result, " "))
		t.Error(msg)
	}
}

func TestProfileMostProbableKmers2(t *testing.T) {
	gold := "TGTCGC"
	dna := "TGCCCGAGCTATCTTATGCGCATCGCATGCGGACCCTTCCCTAGGCTTGTCGCAAGCCATTATCCTGGGCGCTAGTTGCGCGAGTATTGTCAGACCTGATGACGCTGTAAGCTAGCGTGTTCAGCGGCGCGCAATGAGCGGTTTAGATCACAGAATCCTTTGGCGTATTCCTATCCGTTACATCACCTTCCTCACCCCTA"
	k := 6
	prof := [][]float32{
		[]float32{0.364, 0.333, 0.303, 0.212, 0.121, 0.242},
		[]float32{0.182, 0.182, 0.212, 0.303, 0.182, 0.303},
		[]float32{0.121, 0.303, 0.182, 0.273, 0.333, 0.303},
		[]float32{0.333, 0.182, 0.303, 0.212, 0.364, 0.152},
	}
	result, _ := ProfileMostProbableKmers(dna, k, prof)

	// Check if gold answer is in our results slice
	var passed_test bool
	for _, r := range result {
		if r == gold {
			passed_test = true
			break
		}
	}

	if !passed_test {
		// The correct kmer was not found in our result
		msg := fmt.Sprintf("Error testing ProfileMostProbableKmer(): found incorrect most probable kmer:\n    Gold: %s\n    Computed: %s\n",
			gold, strings.Join(result, " "))
		t.Error(msg)
	}
}
