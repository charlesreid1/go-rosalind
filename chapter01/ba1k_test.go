package main

import (
	"fmt"
	"testing"
)

func TestMatrixFrequencyArray(t *testing.T) {
	var tests = []struct {
		input string // input string
		k     int    // kmer size
		gold  []int  // array of kmer frequencies
	}{
		{"ACGCGGCTCTGAAA", 2,
			[]int{2, 1, 0, 0, 0, 0, 2, 2, 1, 2, 1, 0, 0, 1, 1, 0},
		},
		{"AAAAC", 2,
			[]int{3, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{"TTAAA", 2,
			[]int{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1},
		},
		{"AAA", 2,
			[]int{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, test := range tests {

		// Money shot
		freq_arr, err := FrequencyArray(test.input, test.k)

		// Check if there was error
		if err != nil {
			t.Error(err)
		}

		if !EqualIntSlices(freq_arr, test.gold) {
			err := fmt.Sprintf("Error testing FrequencyArray():\ninput = %s, k = %d\ncomputed = %v\ngold     = %v\n",
				test.input, test.k,
				freq_arr, test.gold)
			t.Error(err)
		}
	}
}

//func TestMostFrequentKmersMismatchesRevCompFile(t *testing.T) {
//
//	filename := "data/frequent_words_mismatch_complements.txt"
//
//	// Read the contents of the input file
//	// into a single string
//	lines, err := readLines(filename)
//	if err != nil {
//		log.Fatalf("readLines: %v", err)
//	}
//
//	// lines[0]: Input
//	dna := lines[1]
//	params := strings.Split(lines[2], " ")
//	if len(params) < 1 {
//		log.Fatalf("Error splitting second line: only found 0-1 tokens")
//	}
//	// lines[3]: Output
//	gold := strings.Split(lines[4], " ")
//
//	k_str, d_str := params[0], params[1]
//
//	k, err := strconv.Atoi(k_str)
//	if err != nil {
//		log.Fatalf("Error: string to int conversion for parameter k: %v", err)
//	}
//
//	d, err := strconv.Atoi(d_str)
//	if err != nil {
//		log.Fatalf("Error: string to int conversion for parameter d: %v", err)
//	}
//
//	result, err := MostFrequentKmersMismatchesRevComp(dna, k, d)
//
//	// Check if function threw error
//	if err != nil {
//		t.Error(err)
//	}
//
//	// Check that there _was_ a result
//	if len(result) == 0 {
//		msg := fmt.Sprintf("Error testing MostFrequentKmersMismatchesRevComp() using test case from file: length of most frequent kmers found was 0: %q",
//			result)
//		t.Error(msg)
//	}
//
//	// Sort before comparing
//	sort.Strings(gold)
//	sort.Strings(result)
//
//	// These will only be unequal if something went wrong
//	if !EqualStringSlices(gold, result) {
//		msg := fmt.Sprintf("Error testing MostFrequentKmersMismatchesRevComp() using test case from file: most frequent kmers do not match.\ncomputed = %q\ngold     = %q\n",
//			result, gold)
//		t.Error(msg)
//	}
//}
