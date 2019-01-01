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
// BA1A Test

// Run a single test of the PatternCount function
func TestPatternCount(t *testing.T) {
	// Call the PatternCount function
	input := "GCGCG"
	pattern := "GCG"
	result := PatternCount(input, pattern)
	gold := 2
	if result != gold {
		err := fmt.Sprintf("Error testing PatternCount(): input = %s, pattern = %s, result = %d (should be %d)",
			input, pattern, result, gold)
		t.Error(err)
	}
}

// Run a test matrix of the PatternCount function
func TestMatrixPatternCount(t *testing.T) {
	// Construct a test matrix
	var tests = []struct {
		input   string
		pattern string
		gold    int
	}{
		{"GCGCG", "GCG", 2},
		{"GAGGGGGGGAG", "AGG", 1},
		{"GCACGCACGCAC", "GCAC", 3},
		{"", "GC", 0},
		{"GCG", "GTACTCTC", 0},
		{"ACGTACGTACGT", "CG", 3},
		{"AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAGACAGATAATAATTACAGAGTACACAACATCCA",
			"AAA", 4},
		{"AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCCAGCGCGTTAACCACGTTCTGTGCCGACTTT",
			"TTT", 4},
		{"GGACTTACTGACGTACG", "ACT", 2},
		{"ATCCGATCCCATGCCCATG", "CC", 5},
		{"CTGTTTTTGATCCATGATATGTTATCTCTCCGTCATCAGAAGAACAGTGACGGATCGCCCTCTCTCTTGGTCAGGCGACCGTTTGCCATAATGCCCATGCTTTCCAGCCAGCTCTCAAACTCCGGTGACTCGCGCAGGTTGAGT",
			"CTC", 9},
	}
	for _, test := range tests {
		result := PatternCount(test.input, test.pattern)
		if result != test.gold {
			err := fmt.Sprintf("Error testing PatternCount(): input = %s, pattern = %s, result = %d (should be %d)",
				test.input, test.pattern, result, test.gold)
			t.Error(err)
		}
	}
}

// Load a PatternCount test (input and output)
// from a file. Run the test with the input
// and verify the output matches the output
// contained in the file.
func TestPatternCountFile(t *testing.T) {

	filename := "data/pattern_count.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	input := lines[1]
	pattern := lines[2]

	// lines[3]: Output
	output_str := lines[4]

	// Convert output to inteter
	output, err := strconv.Atoi(output_str)
	if err != nil {
		t.Error(err)
	}

	// Call the function with the given inputs
	result := PatternCount(input, pattern)

	// Verify answer
	if result != output {
		err := fmt.Sprintf("Error testing PatternCount using test case from file: results do not match:\rcomputed result = %d\nexpected output = %d", result, output)
		t.Error(err)
	}
}

/////////////////////////////////
// BA1B Test

// Run a test of the MostFrequentKmers function
func TestMostFrequentKmers(t *testing.T) {
	// Call MostFrequentKmers
	input := "AAAATGCGCTAGTAAAAGTCACTGAAAA"
	k := 4
	result, err := MostFrequentKmers(input, k)
	gold := []string{"AAAA"}

	if err != nil {
		t.Error(err)
	}

	if !EqualStringSlices(result, gold) {
		err := fmt.Sprintf("Error testing MostFrequentKmers(): input = %s, k = %d, result = %s (should be %s)",
			input, k, result, gold)
		t.Error(err)
	}
}

// Run a test of the PatternCount function
// using inputs/outputs from a file.
func TestMostFrequentKmersFile(t *testing.T) {

	filename := "data/frequent_words.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	dna := lines[1]
	k_str := lines[2]
	// lines[3]: Output
	gold := strings.Split(lines[4], " ")

	// Convert k to integer
	k, err := strconv.Atoi(k_str)
	if err != nil {
		t.Error(err)
	}

	// Call the function with the given inputs
	result, err := MostFrequentKmers(dna, k)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// Check that there _was_ a result
	if len(result) == 0 {
		err := fmt.Sprintf("Error testing MostFrequentKmers using test case from file: length of most frequent kmers found was 0: %q",
			result)
		t.Error(err)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	// These will only be unequal if something went wrong
	if !EqualStringSlices(gold, result) {
		err := fmt.Sprintf("Error testing MostFrequentKmers using test case from file: most frequent kmers do not match.\ncomputed = %q\ngold     = %q\n",
			result, gold)
		t.Error(err)
	}
}

/////////////////////////////////
// BA1C Test

// Check that the DNA2Bitmasks utility
// extracts the correct bitmasks from
// a DNA input string.
func TestDNA2Bitmasks(t *testing.T) {

	input := "AATCCGCT"

	result, func_err := DNA2Bitmasks(input)

	// Handle errors from in the DNA2Bitmasks function
	if func_err != nil {
		err := fmt.Sprintf("Error in function DNA2Bitmasks(): input = %s", input)
		t.Error(err)
	}

	// Assemble gold standard answer (bitvectors)
	tt := true
	ff := false
	gold := make(map[string][]bool)
	gold["A"] = []bool{tt, tt, ff, ff, ff, ff, ff, ff}
	gold["T"] = []bool{ff, ff, tt, ff, ff, ff, ff, tt}
	gold["C"] = []bool{ff, ff, ff, tt, tt, ff, tt, ff}
	gold["G"] = []bool{ff, ff, ff, ff, ff, tt, ff, ff}

	// Verify result from DNA2Bitmasks is same as
	// our gold standard
	for _, cod := range "ATCG" {
		cods := string(cod)
		if !EqualBoolSlices(result[cods], gold[cods]) {
			err := fmt.Sprintf("Error testing DNA2Bitmasks(): input = %s, codon = %s, extracted = %v, gold = %v",
				input, cods, result[cods], gold[cods])
			t.Error(err)
		}
	}
}

// Check that the Bitmasks2DNA utility
// constructs the correct DNA string
// from bitmasks.
func TestBitmasks2DNA(t *testing.T) {
	// Assemble input bitmasks
	tt := true
	ff := false
	input := make(map[string][]bool)
	input["A"] = []bool{tt, tt, ff, ff, ff, ff, ff, ff}
	input["T"] = []bool{ff, ff, tt, ff, ff, ff, ff, tt}
	input["C"] = []bool{ff, ff, ff, tt, tt, ff, tt, ff}
	input["G"] = []bool{ff, ff, ff, ff, ff, tt, ff, ff}

	gold := "AATCCGCT"

	result, func_err := Bitmasks2DNA(input)

	// Handle errors from in the DNA2Bitmasks function
	if func_err != nil {
		err := fmt.Sprintf("Error in function Bitmasks2DNA(): function returned error")
		t.Error(err)
	}

	// Verify result from DNA2Bitmasks is same as
	// our gold standard
	if result != gold {
		err := fmt.Sprintf("Error testing Bitmasks2DNA(): result = %s, gold = %s", result, gold)
		t.Error(err)
	}
}

// Run a test of the function that computes
// the ReverseComplement of a DNA string.
func TestReverseComplement(t *testing.T) {
	input := "AAAACCCGGT"
	result, _ := ReverseComplement(input)
	gold := "ACCGGGTTTT"
	if result != gold {
		err := fmt.Sprintf("Error testing ReverseComplement(): input = %s, result = %s (should be %s)",
			input, result, gold)
		t.Error(err)
	}
}

// Run a test of the ReverseComplement function
// using inputs/outputs from a file.
func TestReverseComplementFile(t *testing.T) {

	filename := "data/reverse_complement.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		t.Error(err)
	}

	// lines[0]: Input
	input := lines[1]
	// lines[2]: Output
	gold := lines[3]

	// Call the function with the given inputs
	result, err := ReverseComplement(input)

	// Check that there _was_ a result
	if len(result) == 0 || err != nil {
		msg := fmt.Sprintf("Error testing ReverseComplement using test case from file")
		t.Error(msg)
	}

	if result != gold {
		msg := fmt.Sprintf("Error testing ReverseComplement(): input = %s, result = %s (should be %s)",
			input, result, gold)
		t.Error(msg)
	}
}

/////////////////////////////////
// BA1D Test

func TestFindOccurrences(t *testing.T) {
	// Call FindOccurrences
	pattern := "ATAT"
	genome := "GATATATGCATATACTT"

	result, err := FindOccurrences(pattern, genome)
	gold := []int{1, 3, 9}

	if !EqualIntSlices(result, gold) || err != nil {
		err := fmt.Sprintf("Error testing FindOccurrences(): result = %q, should be %q",
			result, gold)
		t.Error(err)
	}
}

func TestFindOccurrencesDebug(t *testing.T) {
	// Construct a test matrix
	var tests = []struct {
		pattern string
		genome  string
		gold    []int
	}{
		{"ACAC", "TTTTACACTTTTTTGTGTAAAAA",
			[]int{4}},
		{"AAA", "AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAGACAGATAATAATTACAGAGTACACAACATCCAT",
			[]int{0, 46, 51, 74}},
		{"TTT", "AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCCAGCGCGTTAACCACGTTCTGTGCCGACTTT",
			[]int{88, 92, 98, 132}},
		{"ATA", "ATATATA",
			[]int{0, 2, 4}},
	}
	for _, test := range tests {

		result, err := FindOccurrences(test.pattern, test.genome)

		if err != nil {
			t.Error(err)
		}

		if !EqualIntSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing FindOccurrences(): result = %q, should be %q",
				result, test.gold)
			t.Error(err)
		}
	}
}

func TestFindOccurrencesFiles(t *testing.T) {

	filename := "data/pattern_matching.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// lines[0]: Input
	pattern := lines[1]
	genome := lines[2]

	// lines[3]: Output
	gold_str := lines[4]
	gold_slice := strings.Split(gold_str, " ")

	gold := make([]int, len(gold_slice))
	for i, g := range gold_slice {
		gold[i], err = strconv.Atoi(g)
		if err != nil {
			t.Error(err)
		}
	}

	result, err := FindOccurrences(pattern, genome)

	if err != nil {
		t.Error(err)
	}

	if !EqualIntSlices(result, gold) {
		err := fmt.Sprintf("Error testing FindOccurrences():\nresult = %v\ngold   = %v\n",
			result, gold)
		t.Error(err)
	}
}

/////////////////////////////////
// BA1E Test

// TODO: add a test that loads a file in data/

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

/////////////////////////////////
// BA1F Test

// TODO: add a test loaded from a file in data/

func TestMatrixMinSkewPosition(t *testing.T) {
	var tests = []struct {
		genome string
		gold   []int
	}{
		{"CCTATCGGTGGATTAGCATGTCCCTGTACGTTTCGCCGCGAACTAGTTCACACGGCTTGATGGCAAATGGTTTTTCCGGCGACCGTAATCGTCCACCGAG",
			[]int{53, 97}},
		{"TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGA",
			[]int{11, 24}},
		{"ACCG",
			[]int{3}},
		{"ACCC",
			[]int{4}},
		{"CCGGGT",
			[]int{2}},
		{"CCGGCCGG",
			[]int{2, 6}},
	}
	for _, test := range tests {

		// Do it - find the positions that minimize skew
		result, err := MinSkewPositions(test.genome)
		if err != nil {
			t.Error(err)
		}

		// Check length of result
		if len(result) != len(test.gold) {
			err := fmt.Sprintf("Error testing MinSkewPositions():\nfor genome: %s\nlength of result (%d) did not match length of gold standard (%d).\nFound: %v\nShould be: %v",
				test.genome, len(result), len(test.gold),
				result, test.gold)
			t.Error(err)
		}

		// Sort before comparing
		sort.Ints(result)
		sort.Ints(test.gold)
		if !EqualIntSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing MinSkewPositions():\nfor genome: %s\nfound: %v\nshould be: %v",
				test.genome, result, test.gold)
			t.Error(err)
		}
	}
}

/////////////////////////////////
// BA1G Test

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

/////////////////////////////////
// BA1H Test

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

/////////////////////////////////
// BA1i Test

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

/////////////////////////////////
// BA1J Test

func TestMatrixMostFrequentKmersMismatchesRevComp(t *testing.T) {
	var tests = []struct {
		input string   // input string
		k     int      // kmer size
		d     int      // max Hamming distance
		gold  []string // old standard true value
	}{
		{"ACGTTGCATGTCGCATGATGCATGAGAGCT",
			4, 1,
			[]string{"ATGT", "ACAT"},
		},
	}
	for _, test := range tests {

		// Money shot
		result, err := MostFrequentKmersMismatchesRevComp(test.input, test.k, test.d)

		// Check if there was error
		if err != nil {
			t.Error(err)
		}

		// Sort before comparing
		sort.Strings(test.gold)
		sort.Strings(result)

		if !EqualStringSlices(result, test.gold) {
			err := fmt.Sprintf("Error testing MostFrequentKmersMismatchesRevComp():\ninput = %s, k = %d, d = %d\ncomputed = %v\ngold     = %v\n",
				test.input, test.k, test.d,
				result, test.gold)
			t.Error(err)
		}
	}
}

/*
func TestMostFrequentKmersMismatchesRevCompFile(t *testing.T) {

	filename := "data/frequent_words_mismatch_complements.txt"

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

	result, err := MostFrequentKmersMismatchesRevComp(dna, k, d)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// Check that there _was_ a result
	if len(result) == 0 {
		msg := fmt.Sprintf("Error testing MostFrequentKmersMismatchesRevComp() using test case from file: length of most frequent kmers found was 0: %q",
			result)
		t.Error(msg)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	// These will only be unequal if something went wrong
	if !EqualStringSlices(gold, result) {
		msg := fmt.Sprintf("Error testing MostFrequentKmersMismatchesRevComp() using test case from file: most frequent kmers do not match.\ncomputed = %q\ngold     = %q\n",
			result, gold)
		t.Error(msg)
	}
}
*/

/////////////////////////////////
// BA1K Test

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

/////////////////////////////////
// BA1Lima Test

func TestPatternToNumber(t *testing.T) {
	input := "AGT"
	gold := 11

	// Money shot
	number, err := PatternToNumber(input)

	// Check if there was error
	if err != nil {
		t.Error(err)
	}

	if number != gold {
		err := fmt.Sprintf("Error testing PatternToNumber():\ninput = %s\ncomputed = %v\ngold     = %v\n",
			input, number, gold)
		t.Error(err)
	}
}

func TestPatternToNumberFile(t *testing.T) {

	filename := "data/pattern_to_number.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	input := lines[1]
	// lines[2]: Output
	gold_str := lines[3]

	gold, err := strconv.Atoi(gold_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for integer representation of input DNA string: %v", err)
	}

	number, err := PatternToNumber(input)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// These will only be unequal if something went wrong
	if number != gold {
		err := fmt.Sprintf("Error testing PatternToNumber():\ninput = %s\ncomputed = %v\ngold     = %v\n",
			input, number, gold)
		t.Error(err)
	}
}

/////////////////////////////////
// BA1M Test

func TestNumberToPattern(t *testing.T) {
	n := 11
	k := 3
	gold := "AGT"

	// Money shot
	pattern, err := NumberToPattern(n, k)

	// Check if there was error
	if err != nil {
		t.Error(err)
	}

	if pattern != gold {
		msg := fmt.Sprintf("Error testing NumberToPattern():\nn = %d, k = %d\ncomputed = %v\ngold     = %v\n",
			n, k, pattern, gold)
		t.Error(msg)
	}
}

func TestNumberToPatternFile(t *testing.T) {

	filename := "data/number_to_pattern.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	n_str := lines[1]
	k_str := lines[2]
	// lines[3]: Output
	gold := lines[4]

	n, err := strconv.Atoi(n_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for integer n representing DNA string: %v", err)
	}

	k, err := strconv.Atoi(k_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for integer k kmer length: %v", err)
	}

	pattern, err := NumberToPattern(n, k)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// These will only be unequal if something went wrong
	if pattern != gold {
		err := fmt.Sprintf("Error testing NumberToPattern():\nn = %d, k = %d\ncomputed = %v\ngold     = %v\n",
			n, k, pattern, gold)
		t.Error(err)
	}
}

/////////////////////////////////
// BA1N Test

func TestVisitNeighborhood(t *testing.T) {
	input := "ACG"
	d := 1
	gold := []string{"CCG", "TCG", "GCG", "AAG", "ATG", "AGG", "ACA", "ACC", "ACT", "ACG"}

	// Money shot
	result, err := VisitHammingNeighbors(input, d)

	// Check if there was error
	if err != nil {
		t.Error(err)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	if !EqualStringSlices(gold, result) {
		msg := fmt.Sprintf("Error testing VisitHammingNeighbors():\ninput = %s, = %d\ncomputed = %v\ngold     = %v\n",
			input, d,
			result, gold)
		t.Error(msg)
	}
}

func TestVisitNeighborhoodFile(t *testing.T) {
	filename := "data/neighbors.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	input := strings.TrimSpace(lines[1])
	d_str := strings.TrimSpace(lines[2])
	// lines[3]: Output

	d, err := strconv.Atoi(d_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for parameter d: %v", err)
	}

	gold := make([]string, len(lines)-4)
	for i := 4; i < len(lines); i++ {
		j := i - 4
		gold[j] = lines[i]
	}

	result, err := VisitHammingNeighbors(input, d)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// Check that there _was_ a result
	if len(result) == 0 {
		msg := fmt.Sprintf("Error testing VisitHammingNeighbors() using test case from file: no results")
		t.Error(msg)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(result)

	// These will only be unequal if something went wrong
	if !EqualStringSlices(gold, result) {
		msg := fmt.Sprintf("Error testing VisitHammingNeighbors() using test case from file: final results do not match.\ncomputed = %q\ngold     = %q\n",
			result, gold)
		t.Error(msg)
	}
}
