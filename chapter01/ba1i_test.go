package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"
)

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

func TestMatrixMostFrequentKmersMismatches(t *testing.T) {
	var tests = []struct {
		input string   // input string
		k     int      // kmer size
		d     int      // max Hamming distance
		gold  []string // old standard true value
	}{
		{"AAATTTCCC",
			1, 0,
			[]string{"A", "T", "C"},
		},
		{"ACGTTGCATGTCGCATGATGCATGAGAGCT",
			4, 1,
			[]string{"ATGC", "ATGT", "GATG"},
		},
		{"AAAAAAAAAA",
			2, 1,
			[]string{"AA", "AC", "AG", "CA", "AT", "GA", "TA"},
		},
		{"AGTCAGTC",
			4, 2,
			[]string{"TCTC", "CGGC", "AAGC", "TGTG", "GGCC", "AGGT", "ATCC", "ACTG", "ACAC", "AGAG", "ATTA", "TGAC", "AATT", "CGTT", "GTTC", "GGTA", "AGCA", "CATC"},
		},
		{"AATTAATTGGTAGGTAGGTA",
			4, 0,
			[]string{"GGTA"},
		},
		{"ATA",
			3, 1,
			[]string{"GTA", "ACA", "AAA", "ATC", "ATA", "AGA", "ATT", "CTA", "TTA", "ATG"},
		},
		{"AAT",
			3, 0,
			[]string{"AAT"},
		},
		{"TAGCG",
			2, 1,
			[]string{"GG", "TG"},
		},
	}
	for _, test := range tests {

		// Money shot
		result, err := MostFrequentKmersMismatches(test.input, test.k, test.d)

		// Check if there was error
		if err != nil {
			t.Error(err)
		}

		// Sort before comparing
		sort.Strings(test.gold)
		sort.Strings(result)

		if !EqualStringSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing MostFrequentKmersMismatches():\ninput = %s, k = %d, d = %d\ncomputed = %v\ngold     = %v\n",
				test.input, test.k, test.d,
				result, test.gold)
			t.Error(err)
		}
	}
}

func TestMostFrequentKmersMismatchesFile(t *testing.T) {

	filename := "data/frequent_words_mismatch.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	dna := lines[1]
	params := strings.Split(lines[2], " ")
	if len(params) < 1 {
		log.Fatalf("Error splitting second line: only found 0-1 tokens")
	}
	// lines[3]: Output
	gold := strings.Split(lines[4], " ")

	k_str, d_str := params[0], params[1]

	k, err := strconv.Atoi(k_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for parameter k: %v", err)
	}

	d, err := strconv.Atoi(d_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for parameter d: %v", err)
	}

	result, err := MostFrequentKmersMismatches(dna, k, d)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// Check that there _was_ a result
	if len(result) == 0 {
		msg := fmt.Sprintf("Error testing MostFrequentKmersMismatches() using test case from file: length of most frequent kmers found was 0: %q",
			result)
		t.Error(msg)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	// These will only be unequal if something went wrong
	if !EqualStringSlices(gold, result) {
		msg := fmt.Sprintf("Error testing MostFrequentKmersMismatches() using test case from file: most frequent kmers do not match.\ncomputed = %q\ngold     = %q\n",
			result, gold)
		t.Error(msg)
	}
}
