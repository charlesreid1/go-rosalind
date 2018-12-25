package main

import (
	"errors"
	"fmt"
	"sort"
	s "strings"
)

/*
rosalind.go:

This file contains core functions that
are used to solve Rosalind problems.
*/

////////////////////////////////
// BA1A

// Count occurrences of a substring pattern
// in a string input
func PatternCount(input string, pattern string) int {

	// Number of substring overlaps
	var overlap = len(input) - len(pattern) + 1

	// If overlap < 1, we are looking
	// for a pattern longer than our input
	if overlap < 1 {
		return 0
	}

	// Count of occurrences
	count := 0

	// Loop over each substring overlap
	for i := 0; i < overlap; i++ {
		// Grab a slice of the full input
		start := i
		end := i + len(pattern)
		var slice = input[start:end]
		if slice == pattern {
			count += 1
		}
	}
	return count
}

////////////////////////////////
// BA1B

// Return the histogram of kmers of length k
// found in the given input
func KmerHistogram(input string, k int) (map[string]int, error) {

	if len(input) < 1 {
		err := fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s", input)
		return nil, errors.New(err)
	}

	result := map[string]int{}

	// Number of substring overlaps
	overlap := len(input) - k + 1

	// If overlap < 1, we are looking
	// for kmers longer than our input
	if overlap < 1 {
		return result, nil
	}

	// Iterate over each position,
	// extract the string,
	// increment the count.
	for i := 0; i < overlap; i++ {
		// Get the kmer of interest
		substr := input[i : i+k]

		// If it doesn't exist, the value is 0
		result[substr] += 1
	}

	return result, nil
}

// Find the most frequent kmer(s) in the kmer histogram,
// and return as a string array slice
func MostFrequentKmers(input string, k int) ([]string, error) {
	max := 0

	// most frequent kmers
	mfks := []string{}

	if k < 1 {
		err := fmt.Sprintf("Error: MostFrequentKmers received a kmer size that was not a natural number: k = %d", k)
		return mfks, errors.New(err)
	}

	khist, err := KmerHistogram(input, k)

	if err != nil {
		err := fmt.Sprintf("Error: MostFrequentKmers failed when calling KmerHistogram()")
		return mfks, errors.New(err)
	}

	for kmer, freq := range khist {
		if freq > max {
			// We have a new maximum, and a new set of kmers
			max = freq
			mfks = []string{kmer}
		} else if freq == max {
			// We have another maximum
			mfks = append(mfks, kmer)
		}
	}
	return mfks, nil
}

// Find the kmer(s) in the kmer histogram
// exceeding a count of N, and return as
// a string array slice
func MoreFrequentThanNKmers(input string, k, N int) ([]string, error) {

	// more frequent than n kmers
	mftnks := []string{}

	if k < 1 || N < 1 {
		err := fmt.Sprintf("Error: MoreFrequentThanNKmers received a kmer or frequency size that was not a natural number: k = %d, N = %d", k, N)
		return mftnks, errors.New(err)
	}

	khist, err := KmerHistogram(input, k)

	if err != nil {
		err := fmt.Sprintf("Error: MoreFrequentThanNKmers failed when calling KmerHistogram()")
		return mftnks, errors.New(err)
	}

	for kmer, freq := range khist {
		if freq >= N {
			// Add another more frequent than n
			mftnks = append(mftnks, kmer)
		}
	}
	return mftnks, nil
}

////////////////////////////////
// BA1C

// Reverse returns its argument string reversed
// rune-wise left to right.
// https://github.com/golang/example/blob/master/stringutil/reverse.go
func ReverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Given an alleged DNA input string,
// iterate through it character by character
// to ensure that it only contains ATGC.
// Returns true if this is DNA (ATGC only),
// false otherwise.
func CheckIsDNA(input string) bool {

	// Convert input to uppercase
	input = s.ToUpper(input)

	// If any character is not ATCG, fail
	for _, c := range input {
		if c != 'A' && c != 'T' && c != 'C' && c != 'G' {
			return false
		}
	}

	// If we made it here, everything's gravy!
	return true
}

// Convert a DNA string into four bitmasks:
// one each for ATGC. That is, for the DNA
// string AATCCGCT, it would become:
//
// bitmask[A] = 11000000
// bitmask[T] = 00100001
// bitmask[C] = 00011010
// bitmask[G] = 00000100
func DNA2Bitmasks(input string) (map[string][]bool, error) {

	// Convert input to uppercase
	input = s.ToUpper(input)

	// Allocate space for the map
	m := make(map[string][]bool)

	// Start by checking whether we have DNA
	if CheckIsDNA(input) == false {
		err := fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s", input)
		return m, errors.New(err)
	}

	// Important: we want to iterate over the
	// DNA string ONCE and only once. That means
	// we need to have the bit vectors initialized
	// already, and as we step through the DNA
	// string, we access the appropriate index
	// of the appropriate bit vector and set
	// it to true.
	m["A"] = make([]bool, len(input))
	m["T"] = make([]bool, len(input))
	m["C"] = make([]bool, len(input))
	m["G"] = make([]bool, len(input))

	// To begin with, every bit vector is false.
	for i, c := range input {
		cs := string(c)
		// Get the corresponding bit vector - O(1)
		bitty := m[cs]
		// Flip to true for this position - O(1)
		bitty[i] = true
	}

	return m, nil
}

// Convert four bitmasks (one each for ATGC)
// into a DNA string.
func Bitmasks2DNA(bitmasks map[string][]bool) (string, error) {

	// Verify ATGC keys are all present
	_, Aok := bitmasks["A"]
	_, Tok := bitmasks["T"]
	_, Gok := bitmasks["G"]
	_, Cok := bitmasks["C"]
	if !(Aok && Tok && Gok && Cok) {
		err := fmt.Sprintf("Error: input bitmask was missing one of: ATGC (Keys present? A: %t, T: %t, G: %t, C: %t", Aok, Tok, Gok, Cok)
		return "", errors.New(err)
	}

	// Hope that all bitmasks are the same size
	size := len(bitmasks["A"])

	// Make a rune array that we'll turn into
	// a string for our final return value
	dna := make([]rune, size)

	// Iterate over the bitmask, using only
	// the index and not the mask value itself
	for i, _ := range bitmasks["A"] {
		if bitmasks["A"][i] == true {
			dna[i] = 'A'
		} else if bitmasks["T"][i] == true {
			dna[i] = 'T'
		} else if bitmasks["G"][i] == true {
			dna[i] = 'G'
		} else if bitmasks["C"][i] == true {
			dna[i] = 'C'
		}
	}

	return string(dna), nil
}

// Given a DNA input string, find the
// complement. The complement swaps
// Gs and Cs, and As and Ts.
func Complement(input string) (string, error) {

	// Convert input to uppercase
	input = s.ToUpper(input)

	// Start by checking whether we have DNA
	if CheckIsDNA(input) == false {
		return "", errors.New(fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s", input))
	}

	m, _ := DNA2Bitmasks(input)

	// Swap As and Ts
	newT := m["A"]
	newA := m["T"]
	m["T"] = newT
	m["A"] = newA

	// Swap Cs and Gs
	newG := m["C"]
	newC := m["G"]
	m["G"] = newG
	m["C"] = newC

	output, _ := Bitmasks2DNA(m)

	return output, nil
}

// Given a DNA input string, find the
// reverse complement. The complement
// swaps Gs and Cs, and As and Ts.
// The reverse complement reverses that.
func ReverseComplement(input string) (string, error) {

	// Convert input to uppercase
	input = s.ToUpper(input)

	// Start by checking whether we have DNA
	if CheckIsDNA(input) == false {
		err := fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s", input)
		return "", errors.New(err)
	}

	comp, _ := Complement(input)

	revcomp := ReverseString(comp)

	return revcomp, nil
}

////////////////////////////////
// BA1D

// Given a large string (genome) and a string (pattern),
// find the zero-based indices where pattern occurs in genome.
func FindOccurrences(pattern, genome string) ([]int, error) {
	locations := []int{}
	slots := len(genome) - len(pattern) + 1

	if slots < 1 {
		// pattern is longer than genome
		return locations, nil
	}

	// Loop over each character,
	// saving the position if it
	// is the start of pattern
	for i := 0; i < slots; i++ {
		start := i
		end := i + len(pattern)
		if genome[start:end] == pattern {
			locations = append(locations, i)
		}
	}
	return locations, nil
}

////////////////////////////////
// BA1E

// Find k-mers (patterns) of length k occuring at least
// t times over an interval of length L in a genome.
func FindClumps(genome string, k, L, t int) ([]string, error) {

	// Algorithm:
	// allocate a list of kmers
	// for each possible position of L window,
	//   feed string L to KmerHistogram()
	//   save any kmers with frequency > t
	// return master list of saved kmers

	L_slots := len(genome) - L + 1

	// Set kmers
	kmers := map[string]bool{}

	// List kmers
	kmers_list := []string{}

	// Loop over each possible window of length L
	for iL := 0; iL < L_slots; iL++ {

		// Grab this portion of the genome
		winstart := iL
		winend := iL + L
		genome_window := genome[winstart:winend]

		// Get the number of kmers that occur more
		// frequently than t times
		new_kmers, err := MoreFrequentThanNKmers(genome_window, k, t)
		if err != nil {
			return kmers_list, err
		}
		// Add these to the set kmers
		for _, new_kmer := range new_kmers {
			kmers[new_kmer] = true
		}
	}

	for k := range kmers {
		kmers_list = append(kmers_list, k)
	}
	sort.Strings(kmers_list)

	return kmers_list, nil
}

////////////////////////////////
// BA1F

// The skew of a genome is the difference between
// the number of G and C codons that have occurred
// cumulatively in a given strand of DNA.
// This function computes the positions in the genome
// at which the cumulative skew is minimized.
func MinSkewPositions(genome string) ([]int, error) {

	n := len(genome)
	cumulative_skew := make([]int, n+1)

	// Get C/G bitmasks
	bitmasks, err := DNA2Bitmasks(genome)
	if err != nil {
		return cumulative_skew, err
	}
	c := bitmasks["C"]
	g := bitmasks["G"]

	// Init
	cumulative_skew[0] = 0

	// Make space to keep track of the
	// minima we have encountered so far
	min := 999
	min_skew_ix := []int{}

	// At each position, compute the next skew value.
	// We need two indices b/c for a genome of size N,
	// the cumulative skew array index is of size N+1.
	for i, ibit := 1, 0; i <= n; i, ibit = i+1, ibit+1 {

		var next int
		// Next skew value
		if c[ibit] {
			// C -1
			next = -1
		} else if g[ibit] {
			// G +1
			next = 1
		} else {
			next = 0
		}
		cumulative_skew[i] = cumulative_skew[i-1] + next

		if cumulative_skew[i] < min {
			// New min and min_skew
			min = cumulative_skew[i]
			min_skew_ix = []int{i}
		} else if cumulative_skew[i] == min {
			// Additional min and min_skew
			min_skew_ix = append(min_skew_ix, i)
		}

	}
	return min_skew_ix, nil
}

////////////////////////////////
// BA1G

// Compute the Hamming distance between
// two strings. The Hamming distance is
// defined as the number of characters
// different between two strings.
func HammingDistance(p, q string) (int, error) {

	// Technically a Hamming distance when
	// one string is empty would be 0, but
	// we will throw an error instead.
	if len(p) == 0 || len(q) == 0 {
		err := fmt.Sprintf("Error: HammingDistance: one or more arguments had length 0. len(p) = %d, len(q) = %d", len(p), len(q))
		return -1, errors.New(err)
	}

	// Get longest length common to both
	var m int
	if len(p) > len(q) {
		m = len(q)
	} else {
		m = len(p)
	}

	// Accumulate distance
	dist := 0
	for i := 0; i < m; i++ {
		if p[i] != q[i] {
			dist += 1
		}
	}
	return dist, nil
}

////////////////////////////////
// BA1H

// Given a large string (text) and a string (pattern),
// find the zero-based indices where we have an occurrence
// of pattern or a string with Hamming distance d or less
// from pattern.
func FindApproximateOccurrences(pattern, text string, d int) ([]int, error) {

	locations := []int{}
	slots := len(text) - len(pattern) + 1

	if slots < 1 {
		// pattern is longer than genome
		return locations, nil
	}

	// Loop over each character,
	// saving the position if it
	// is the start of pattern
	for i := 0; i < slots; i++ {
		start := i
		end := i + len(pattern)
		poss_approx_pattern := text[start:end]
		hamm, _ := HammingDistance(poss_approx_pattern, pattern)
		if hamm <= d {
			locations = append(locations, i)
		}
	}

	return locations, nil
}

////////////////////////////////
// BA1i

// Count the number of times a given kmer
// and any Hamming neighbors (distance d
// or less) occur in the input string.
func CountKmersMismatches(input string, k, d int) (int, error) {
	// Note this skips step 1 of most frequent
	// (extract all kmers and get a list of all variants)
	// and goes straight to step 2
	// (given a kmer, increment count of all its variants)
	return 0, nil
}

// Find the most frequent kmer(s) of length k
// in the given input string. Include mismatches
// of Hamming distance <= d.
func MostFrequentKmersMismatches(input string, k, d int) ([]string, error) {

	max := 0

	// most frequent kmers
	mfks := []string{}

	if k < 1 {
		err := fmt.Sprintf("Error: MostFrequentKmers received a kmer size that was not a natural number: k = %d", k)
		return mfks, errors.New(err)
	}

	khist, err := KmerHistogramMismatches(input, k, d)

	if err != nil {
		err := fmt.Sprintf("Error: MostFrequentKmers failed when calling KmerHistogram()")
		return mfks, errors.New(err)
	}

	for kmer, freq := range khist {
		if freq > max {
			// We have a new maximum, and a new set of kmers
			max = freq
			mfks = []string{kmer}
		} else if freq == max {
			// We have another maximum
			mfks = append(mfks, kmer)
		}
	}
	return mfks, nil
}

// Return the histogram of all kmers of length k
// that are in the input, or whose Hamming neighbors
// within distance d are in the input.
func KmerHistogramMismatches(input string, k, d int) (map[string]int, error) {

	// Make sure our input string is well-formed
	if !CheckIsDNA(input) {
		err := fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s", input)
		return nil, errors.New(err)
	}

	// Number of substring overlaps
	overlap := len(input) - k + 1

	// If overlap < 1, we are looking
	// for kmers longer than our input
	if overlap < 1 {
		err := fmt.Sprintf("Error: looking for kmer longer than input string (len(kmer) = %d, len(input) = %d).", k, len(input))
		return nil, errors.New(err)
	}

	// Algorithm:
	// -----------
	//
	// Make two passes over the input string.
	//
	// Pass 1:
	// - Assemble a mapping of each kmer to its Hamming neighbors
	//   so we know which kmers to increment when we see one
	//   (generating all permutations is tiresome and expensive,
	//   so we only want to do it once)
	//
	// Pass 2:
	// - Extract each kmer, get its Hamming neighbors,
	//   increment all of them

	/////////////////////////////////////
	// Pass 1:
	//
	// Assemble a mapping of each kmer to its Hamming neighbors.
	//
	// For each window in overlap:
	//     extract the kmer at window
	//     find all permutations of given string
	//     add kmer as key, variants as value

	hamm_neighbors := map[string][]string{}

	// Iterate over each position
	for i := 0; i < overlap; i++ {

		// TODO goroutines 1:
		// Spawn goroutines, and block
		// until all return results

		// Get the kmer of interest
		kmer := input[i : i+k]

		// Find Hamming neighbors
		neighbors, err := VisitHammingNeighbors(kmer, d)
		if err != nil {
			err := fmt.Sprintf("Error: failed to visit Hamming neighbors for kmer %s (d = %d)", kmer, d)
			return nil, errors.New(err)
		}

		// Increment the count
		hamm_neighbors[kmer] = neighbors
	}

	/////////////////////////////////////
	// Pass 2:
	//
	// Extract each kmer and increment all of its
	// neighbor kmers.

	result := map[string]int{}

	// Iterate over each position
	for i := 0; i < overlap; i++ {

		// TODO goroutines 2:
		// Spawn goroutines, and block
		// until all return results

		// Get the kmer of interest
		kmer := input[i : i+k]

		// Get the hamming neighbors
		neighbors := hamm_neighbors[kmer]

		// Increment count for kmer and neighbors
		result[kmer] += 1
		for _, neighbor := range neighbors {
			result[neighbor] += 1
		}

	}

	return result, nil
}

// Given an input string of DNA, generate variants
// of said string that are a Hamming distance of
// less than or equal to d.
func VisitHammingNeighbors(input string, d int) ([]string, error) {
	// a.k.a. visit_kmer_neighbors

	// number of codons
	n_codons := 4

	// Use combinatorics to calculate number
	// of permutations
	buffsize := 0
	for dd := 0; dd <= d; dd++ {

		next_term := Binomial(len(input), dd)
		// old fashioned Pow() function
		for j := 0; j < dd; j++ {
			next_term *= (n_codons - 1)
		}
		buffsize += next_term

	}

	// We need to store all permutations,
	// but the number of permutations will
	// blow up fast, so cut off at some point
	MAX := 100000
	if buffsize > MAX {
		err := fmt.Sprintf("Error: you are generating over MAX = %d permutations, you probably don't want to do this.", d)
		return nil, errors.New(err)
	}

	// Make a channel
	results := make(chan string, buffsize+10)

	// Algorithm:
	// ------------
	//
	// For each depth up to the maximum,
	// begin a recursive function call stack
	// that progressively picks an index to
	// change, then calls itself with a depth
	// parameter decreased by 1.

	// Careful of your index here:
	// - include dd=0 (original kmer)
	// - include dd=d (max depth)
	for dd := 0; dd <= d; dd++ {

		// The choices array will change with each recursive call.
		// Go passes all arguments by copy, which is good for us.
		choices := []int{}

		// Populate list of neighbors
		visitHammingNeighbors_recursive(input, dd, choices, results)

	}

	// Add them to the resulting list of hamming neighbors
	permutations := make([]string, buffsize)
	for k := 0; k < buffsize; k++ {
		permutations[k] = <-results
	}

	return permutations, nil
}

// Recursive function: given an input string of DNA,
// generate Hamming neighbors that are a Hamming
// distance of exactly d. Populate the neighbors
// array with the resulting neighbors.
func visitHammingNeighbors_recursive(base_kmer string, depth int, choices []int, results chan<- string) error {

	if depth == 0 {

		// Base case
		go visit(base_kmer, choices, results)
		return nil

	} else {

		// Recursive case
		for c := 0; c <= len(base_kmer); c++ {

			// This will make a new copy of choices
			// for each recursive function call
			choices2 := append(choices, c)
			err := visitHammingNeighbors_recursive(base_kmer, depth-1, choices2, results)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

// Given a base kmer and a choice of indices where
// the kmer should be changed, generate all possible
// variations on this base_kmer.
func visit(base_kmer string, choices []int, results chan<- string) {

	// We have already made choices,
	// so we don't need to make new choices,
	// we just need to pop them and apply
	// them to the base_kmer input string.
	if len(choices) > 0 {

		all_codons := []string{"A", "T", "G", "C"}

		// Pop the next choice
		// https://github.com/golang/go/wiki/SliceTricks
		ch_ix, choices := choices[0], choices[1:]

		// Get the value of the codon at that location
		// (need string here b/c otherwise this is a byte)
		if ch_ix < len(base_kmer) {
			// slice of string is bytes,
			// so convert back to string
			this_codon := string(base_kmer[ch_ix])
			for _, codon := range all_codons {

				if codon != this_codon {
					// Swap out the old codon with the new codon
					new_kmer := base_kmer[0:ch_ix] + codon + base_kmer[ch_ix+1:]
					go visit(new_kmer, choices, results)
				}
			}
		}

	} else {
		results <- base_kmer
	}
	//return nil
}
