package rosalind

import (
	"errors"
	"fmt"
)

////////////////////////////////
// BA2A

// Given a collection of strings Dna and an integer d,
// a k-mer is a (k,d)-motif if it appears in every
// string from Dna with at most d mismatches.
func FindMotifs(dna []string, k, d int) ([]string, error) {

	for _, input := range dna {
		if !CheckIsDNA(input) {
			msg := fmt.Sprintf("Error: input was not DNA: %s\n", input)
			return nil, errors.New(msg)
		}
	}

	// Pseudocode:
	// for each dna string:
	//   get hamming neighbor histogram k,d
	//   KmerHistogramMismatches(input,k,d)
	// find intersection of all hamming neighbor histogram key sets

	// start using GoDS -
	// efficient data structures.
	// learn from them and use them.

	// For each dna string:
	sets := make([]map[string]int, len(dna))
	for i, input := range dna {

		// Get hamming neighbor histogram
		hist, _ := KmerHistogramMismatches(input, k, d)

		// Add each Hamming neighbor to a hash set
		sets[i] = hist
	}

	// Now we want the intersection of
	// all of the key sets
	intersect, err := KeySetIntersection(sets)
	if err != nil {
		return nil, err
	}

	return intersect, nil
}

// Find the intersection of the key sets
// for a slice of string to integer maps.
func KeySetIntersection(input []map[string]int) ([]string, error) {

	saves := []string{}
	for key := range input[0] {
		// Assume this kmer is in each histogram
		in_everyone := true

		// Iterate over each histogram and
		// make note if it is missing
		for i := 1; i < len(input); i++ {
			hist := input[i]
			if hist[key] == 0 {
				in_everyone = false
				break
			}
		}

		// If this kmer is in everyone's
		// Hamming neighbor histogram,
		// save it
		if in_everyone {
			saves = append(saves, key)
		}
	}
	return saves, nil
}

////////////////////////////////
// BA2b

// Given a k-mer pattern
// and a longer string text,
// find the minimum distance
// from k-mer pattern to
// any possible k-mer in text.
func MinKmerDistance(pattern, text string) (int, error) {

	// Algorithm 1 (faster):
	//
	// Run a sliding window over the input string,
	// and extract all k-mers of width window and
	// add them to a window set.
	//
	// Once the set is assembled, compute the
	// distance from k-mer pattern to k-mers
	// in the window set, and add to distance map.

	hist, err := KmerHistogram(text, len(pattern))
	if err != nil {
		msg := fmt.Sprintf("Error: KmerHistogram(%s,%d) returned error",
			text, len(pattern))
		return -1, errors.New(msg)
	}

	min_dist := len(pattern) // max possible value
	for kmer := range hist {
		d, err := HammingDistance(pattern, kmer)
		if err != nil {
			msg := "Error: HammingDistance() returned error"
			return -1, errors.New(msg)
		}
		if d < min_dist {
			min_dist = d
		}
	}

	// // Algorithm 2 (slower):
	// //
	// // Run a sliding window over the input string,
	// // and compute the distance between the k-mer
	// // pattern and the k-mer in the window.
	// //
	// // This is slow if we have small k and large
	// // input string length, or many duplicate
	// // distance calculations (e.g., 1M ATGATGATG).
	// k := len(pattern)
	// overlap := len(text) - k + 1
	// min_dist := k // max possible value
	// for i := 0; i < overlap; i++ {
	// 	this_kmer := text[i : i+k]
	// 	dist, err := HammingDistance(this_kmer, pattern)
	// 	if err != nil {
	// 		msg := "Error: HammingDistance() returned error"
	// 		return -1, errors.New(msg)
	// 	}
	// 	if dist < min_dist {
	// 		min_dist = dist
	// 	}
	// }

	return min_dist, nil
}

// Given a k-mer pattern
// and a set of strings,
// find the sum (L1 norm)
// of the shortest distances
// from k-mer pattern to
// each input string.
func MinKmerDistances(pattern string, inputs []string) (int, error) {
	s := 0
	for _, text := range inputs {
		d, err := MinKmerDistance(pattern, text)
		s += d
		if err != nil {
			msg := fmt.Sprintf("Error: MinKmerDistance(%s,%s) returned error",
				pattern, text)
			return -1, errors.New(msg)
		}
	}
	return s, nil
}

// k
// Given a k-mer Pattern and a longer string Text, we use d(Pattern, Text) to denote the minimum Hamming distance between Pattern and any k-mer in Text,
//
//d(Pattern,Text)=minall k-mers Pattern' in TextHammingDistance(Pattern,Pattern′).
//
//Given a k-mer Pattern and a set of strings Dna = {Dna1, … , Dnat}, we define d(Pattern, Dna) as the sum of distances between Pattern and all strings in Dna,
//
//d(Pattern,Dna)=∑i=1td(Pattern,Dnai).
//
//Our goal is to find a k-mer Pattern that minimizes d(Pattern, Dna) over all k-mers Pattern, the same task that the Equivalent Motif Finding Problem is trying to achieve. We call such a k-mer a median string for Dna.
//Median String Problem
//
//Find a median string.
//
//Given: An integer k and a collection of strings Dna.
//
//Return: A k-mer Pattern that minimizes d(Pattern, Dna) over all k-mers Pattern. (If multiple answers exist, you may return any one.)
//
//
