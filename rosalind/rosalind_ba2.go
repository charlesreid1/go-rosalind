package rosalind

import (
	"errors"
	"fmt"
	"strings"
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

func MedianString(dna []string, k int) ([]string, error) {

	// Algorithm:

	// start with set of DNA strings dna_i

	// turn each string into set of k-mers
	// set_dna_i is set of k-mers from string dna_i

	// for this_kmer in all_kmers:
	//
	//     for set in set_dna_i:
	//
	//         min_dist = k
	//         for that_kmer in set:
	//             dist = dist(this_kmer,that_kmer)
	//			   min_dist = min(min_dist,dist)

	// Turn each DNA string into a set of kmers
	histograms := make([]map[string]int, len(dna))
	for i, dna_i := range dna {
		h, err := KmerHistogram(dna_i, k)
		if err != nil {
			msg := fmt.Sprintf("Error: KmerHistogram(%s, %d) returned an error",
				dna_i, k)
			return nil, errors.New(msg)
		}
		histograms[i] = h
	}

	// Total number of possible kmer
	maxx := 1
	for i := 0; i < k; i++ {
		maxx *= 4
	}

	// Track min distance sum d(pattern,dna)
	// for all possible kmer patterns
	distances := make([]int, maxx)

	// Iterate over every possible kmer
	for iK := 0; iK < maxx; iK++ {

		// Turn integer iK into kmer pattern
		pattern, err := NumberToPattern(iK, k)
		if err != nil {
			msg := fmt.Sprintf("Error: NumberToPattern(%d,%d) raised an error",
				iK, k)
			return nil, errors.New(msg)
		}

		// Accumulate a min distance sum \sigma d(pattern,dna)
		sigma_min_d := 0

		// Iterate over every possible DNA string('s histogram)
		for _, histogram := range histograms {

			// Accumulate a min distance d(pattern,dna)
			// for this kmer pattern
			// and this DNA string
			min_d := k

			// Iterate over kmers in this DNA string('s histogram)
			// (k,v - map)
			for this_kmer, _ := range histogram {
				d, err := HammingDistance(this_kmer, pattern)
				if err != nil {
					msg := fmt.Sprintf("Error: HammingDistance(%s,%s) returned error",
						this_kmer, pattern)
					return nil, errors.New(msg)
				}
				if d < min_d {
					// New minimum
					min_d = d
				}
			}

			// Accumulate
			sigma_min_d += min_d
		}

		distances[iK] = sigma_min_d

	}

	// Find the kmer corresponding to the minimum distance
	running_min := distances[0]
	var results_str []string
	for i, d := range distances {
		if d < running_min {
			p, err := NumberToPattern(i, k)
			if err != nil {
				msg := fmt.Sprintf("Error: NumberToPattern(%d,%d) returned error",
					i, k)
				return nil, errors.New(msg)
			}
			// New running min, new min kmer
			running_min = d
			results_str = []string{p}

		} else if d == running_min {
			p, err := NumberToPattern(i, k)
			if err != nil {
				msg := fmt.Sprintf("Error: NumberToPattern(%d,%d) returned error",
					i, k)
				return nil, errors.New(msg)
			}
			// Another running min, another min kmer
			results_str = append(results_str, p)
		}
	}

	return results_str, nil
}

////////////////////////////////
// BA2c

// Given a slice of strings, determine
// the index of the given string.
func indexOfString(list []string, item string) int {
	for i := 0; i < len(list); i++ {
		if list[i] == item {
			return i
		}
	}
	return -1
}

// Given a profile matrix,
// and given a DNA input string,
// evaluate the probability of
// every kmer in the DNA string
// and find the most probable
// kmer in the text - the kmer that
// was most likely to have been
// generated by profile among all
// kmers in text.
//
// This is basically a one-at-a-time
// walk down the kmer string, and
// counting up the probability of
// which nucleotide occurred.

func ProfileMostProbableKmers(dna string, k int, profile [][]float32, greedy bool) ([]string, error) {

	nucleotides := []string{"A", "C", "G", "T"}

	// Make sure we have well-formed inputs
	if k < 1 {
		msg := "Error: specified kmer length k was < 1\n"
		return nil, errors.New(msg)
	}
	if !CheckIsDNA(dna) {
		msg := fmt.Sprintf("Error: input was not DNA: %s\n", dna)
		return nil, errors.New(msg)
	}
	if len(profile) != len(nucleotides) {
		msg := fmt.Sprintf("Error: incorrect number of rows (%d) in profile, need 4, one for each nucleotide\n", len(profile))
		return nil, errors.New(msg)
	}

	// Extract all k-mers occurring
	// in the DNA string
	hist, err := KmerHistogram(dna, k)
	if err != nil {
		return nil, err
	}

	// Compute the probability of each kmer
	// by doing pairwise multiplication of
	// probability of the nucleotide that
	// occurs at the corresponding position.
	//
	// Keep track of the running maximum
	// and the corresponding kmer(s).
	var max_prob_kmer []string
	max_prob := float32(0.0)
	for kmer := range hist {
		probability := float32(1.0)
		for j := 0; j < len(kmer); j++ {
			ix := indexOfString(nucleotides, string(kmer[j]))
			probability *= profile[ix][j]
		}
		if probability > max_prob {
			max_prob = probability
			max_prob_kmer = []string{kmer}
		} else if probability == max_prob && !greedy {
			max_prob_kmer = append(max_prob_kmer, kmer)
		}
	}

	if greedy && len(max_prob_kmer) > 1 {
		msg := fmt.Sprintf("Error: incorrect number of profile-most probable kmers. Greedy mode should return 1, but found %d",
			len(max_prob_kmer))
		return nil, errors.New(msg)
	}

	return max_prob_kmer, nil
}

////////////////////////////////
// BA2D
//
// This problem makes about as much sense
// as a camel in a jacuzzi.
//
// After much searching, and re-reading,
// found this great explanation:
//
// http://www.mrgraeme.co.uk/greedy-motif-search/

// ----------------------------
// Scored Motif Matrix struct

// Create a struct to hold a set of motifs (kmers)
// and their associated score. We continually assemble
// many of these possible sets of motifs, checking to
// find a set of motifs with a minimum score.
// The score is not updated dyanmically, see UpdateScore().
type ScoredMotifMatrix struct {
	motifs []string
	score  int
}

// Constructor
func NewScoredMotifMatrix() ScoredMotifMatrix {
	var s ScoredMotifMatrix
	s.motifs = []string{}
	s.score = 0
	return s
}

// Add a motif to the motif matrix
func (s *ScoredMotifMatrix) AddMotif(motif string) error {
	if len(s.motifs) > 0 {
		if len(motif) != len(s.motifs[0]) {
			msg := fmt.Sprintf("Error: could not add motif %s: length %d does not match existing motif length %d",
				motif, len(motif), len(s.motifs[0]))
			return errors.New(msg)
		}
	}
	s.motifs = append(s.motifs, motif)
	return nil
}

// Update the value of the score of a ScoredMotifMatrix.
// This assembles a kmer composed of the most common
// nucleotide per position, then computes the sum of
// the Hamming distances from that kmer for all motifs.
func (s *ScoredMotifMatrix) UpdateScore() error {

	// Params
	t := len(s.motifs)
	k := len(s.motifs[0])

	// Start by assembling a "most common"
	// mer - the kmer containing the most
	// probable nucleotide at each position.
	most_common_kmer := make([]string, k)

	// Loop over every nucleotide
	for ik := 0; ik < k; ik++ {

		// Determine most common nucleotide
		// using a map to count frequencies
		frequency := make(map[string]int)

		// Loop over every DNA string,
		// count nucleotide frequencies
		for it := 0; it < t; it++ {
			bp := string(s.motifs[it][ik])
			frequency[bp] += 1
		}

		// Determine most frequent nucleotide
		var max_bp string
		var max_freq int
		max_freq = 0
		for ibp, ibp_freq := range frequency {
			if ibp_freq > max_freq {
				// Set new maximum occurring base pair
				max_freq = ibp_freq
				max_bp = ibp
			}
		}
		most_common_kmer[ik] = max_bp
	}

	commonkmer := strings.Join(most_common_kmer, "")

	// Now that we have the common kmer,
	// we can compute the score of each motif,
	// and sum their scores to get the total score.
	s.score = 0
	for it := 0; it < t; it++ {
		d, _ := HammingDistance(commonkmer, s.motifs[it])
		s.score += d
	}

	// Done
	return nil
}

func (s *ScoredMotifMatrix) MakeProfile() ([][]float32, error) {
	// Params
	t := len(s.motifs)
	k := len(s.motifs[0])
	nucleotides := []string{"A", "C", "G", "T"}

	// Profile is a 4 x k matrix of float32s
	profile := make([][]float32, 4)
	for jj := 0; jj < 4; jj++ {
		profile[jj] = make([]float32, k)
	}

	// For each column, i.e. kmer nucleotide location,
	// compute the probability
	// of each of the four nucleotides
	//
	// P_i = N_i / sum_j N_j
	//
	for ik := 0; ik < k; ik++ {
		counts := map[string]int{
			"A": 0,
			"C": 0,
			"G": 0,
			"T": 0,
		}

		// Populate counts
		for it := 0; it < t; it++ {
			nucleotide := string(s.motifs[it][ik])
			counts[nucleotide] += 1
		}

		// Sum all values
		summ := 0
		for _, nuc := range nucleotides {
			summ += counts[nuc]
		}

		// Populate p_i
		for inuc, nuc := range nucleotides {
			val := float32(counts[nuc])
			val /= float32(summ)
			profile[inuc][ik] = val
		}
	}

	return profile, nil
}

// ----------------------------
// BA2D functions

// Given an integer k (kmer size) and t (len(dna)),
// return a collection of kmer strings
// that have the lowest score (highest similarity).
//
// If at any step you find more than one
// Profile-most probable k-mer in a given
// DNA string, use the one occurring first.
//
func GreedyMotifSearch(dna []string, k, t int) ([]string, error) {

	var best_smg ScoredMotifMatrix

	// bestmotifs is initially an empty list with score 0
	best_smg = NewScoredMotifMatrix()

	// Loop over all motifs for d = 0,
	// which is equivalent to the
	// keys in the kmer frequencies map
	hist, _ := KmerHistogram(dna[0], k)
	for kmer_motif, _ := range hist {

		// Create a new scored motif group
		this_smg := NewScoredMotifMatrix()

		// Add our motif, which we chose from dna[0]
		// This motif kicks off the new motif group
		this_smg.AddMotif(kmer_motif)

		// Loop over all remaining dna strings
		for _, idna := range dna {

			// Form a profile matrix from
			// all the motifs from dna strings
			// up to, but not including, this one
			profile, _ := this_smg.MakeProfile()
			fmt.Println(profile)

			// Use the profile to find the profile-most
			// probable kmer in this string of dna, idna
			result, _ := ProfileMostProbableKmers(idna, k, profile, true)
			//if len(result) == 0 {
			//	fmt.Println("----------")
			//	fmt.Println(idna)
			//	fmt.Println(profile)
			//}

			//this_smg.AddMotif("") //result[0])
			if false {
				fmt.Println(result)
			}
		}

		this_smg.UpdateScore()
		if this_smg.score > best_smg.score {
			best_smg = this_smg
		}
	}

	return best_smg.motifs, nil

	/*
		for each kmer in the first dna string:

			// examining this kmer
			for each remaining dna string:
				form profile from all dna strings up to this one
				find profile-most probable kmer

			// the motifs you found are each
			// the (first) most probable kmers

			// create a score for that motif:
			//  - find most common nucleotide, per position
			//  - compute number of differences from that nucleotide


	*/
	/*
		If the motifs are the following:

		GTTCAGGCA
		AATCAGTCA
		CGAGTTCGC
		GTCAATCAC
		TAATATTCG
		Score = 7

		The consensus string (most common) is AAG.
		The score is the number of differences
		from that string.

		You get AAG from checking each character
		from the 5 kmers.

		Position 0 has G, A, A, C, C [A most common]
		Position 1 has G, A, A, A, A [A most common]
		Position 2 has C, G, G, C, A [G most common]
		.: AAG

		GGC - AAG: 3 differences
		AAG - AAG: 0 differences
		AAG - AAG: 0 differences
		CAC - AAG: 2 differences
		CAA - AAG: 2 differences

		7 differences total

	*/

	/*
		GREEDYMOTIFSEARCH(Dna, k, t):
		   	BestMotifs ← motif matrix formed by first
						  k-mers in each string
		   	              from Dna
		   	for each k-mer Motif in the first string from Dna
		   	    Motif1 ← Motif
		   	    for i = 2 to t
		   	        form Profile from motifs Motif1, …, Motifi - 1
		   	        Motifi ← Profile-most probable
								k-mer in the i-th i
								string in Dna
		   	    Motifs ← (Motif1, …, Motift)
		   	    if Score(Motifs) < Score(BestMotifs)
		   	        BestMotifs ← Motifs
		   	return BestMotifs
	*/
	/*
		GreedyMotifSearch(k,t,Dna)
			bestMotifs ← empty list (score of 0)
			for i from 0 to |Dna[0]| - k
				motifs ← list with only Dna[0](i,k)
				for j from 1 to |Dna| - 1
					Add ProfileMostProbableKmer( Dna[j], k, Profile(motifs) )
					to motifs
				if score(motifs) < score(bestMotifs)
					bestMotifs = motifs
			return bestMotifs
	*/
}
