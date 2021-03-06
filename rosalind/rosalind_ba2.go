package rosalind

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
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
// This particular method does not
// pay attention to order of occurrence
// of kmers.
func ProfileMostProbableKmers(dna string, k int, profile [][]float32) ([]string, error) {

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
	// in the DNA string.
	// We use the keys of this map for
	// iterating over all kmers in the
	// DNA string.
	// Keys are not ordered!
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
	max_prob := float32(-1.0)
	for kmer := range hist {
		probability := float32(1.0)
		for j := 0; j < len(kmer); j++ {
			ix := indexOfString(nucleotides, string(kmer[j]))
			probability *= profile[ix][j]
		}
		if probability > max_prob {
			max_prob = probability
			max_prob_kmer = []string{kmer}
		} else if probability == max_prob {
			max_prob_kmer = append(max_prob_kmer, kmer)
		}
	}

	return max_prob_kmer, nil
}

// Only return the _most_ probable kmer.
func ProfileMostProbableKmer(dna string, k int, profile [][]float32) (string, error) {
	results, err := ProfileMostProbableKmers(dna, k, profile)
	return results[0], err
}

// Return a list of kmers of length k that occur in a DNA string.
// This list preserves order in which the kmers appear in DNA.
// This list does not include duplicates.
func KmerInOrderList(dna string, k int) ([]string, error) {

	// Ensure we have well-formed inputs
	if len(dna) < k {
		msg := "Error: KmerInOrderList called with a DNA string smaller than specified kmer size"
		return nil, errors.New(msg)
	}

	// kmers is a list of kmers that occur, in order
	// (without duplicates), in dna
	var kmers []string
	overlap := len(dna) - k + 1
	for i := 0; i < overlap; i++ {
		kmer := dna[i : i+k]

		// Check if we already have this kmer
		dupe := false
		for j := 0; j < len(kmers); j++ {
			if kmer == kmers[j] {
				dupe = true
				break
			}
		}
		if !dupe {
			kmers = append(kmers, kmer)
		}
	}
	return kmers, nil
}

// This uses a probility matrix and evaluates
// all possible kmers in a DNA string to determine
// which kmers in the DNA string match the profile
// most closely.
//
// The greedy version maintains the order in which
// kmers occur in the original DNA string, and stops
// as soon as the first match is found.
func ProfileMostProbableKmersGreedy(dna string, k int, profile [][]float32) (string, error) {

	nucleotides := []string{"A", "C", "G", "T"}

	// Make sure we have well-formed inputs
	if k < 1 {
		msg := "Error: specified kmer length k was < 1\n"
		return "", errors.New(msg)
	}
	if !CheckIsDNA(dna) {
		msg := fmt.Sprintf("Error: input was not DNA: %s\n", dna)
		return "", errors.New(msg)
	}
	if len(profile) != len(nucleotides) {
		msg := fmt.Sprintf("Error: incorrect number of rows (%d) in profile, need 4, one for each nucleotide\n", len(profile))
		return "", errors.New(msg)
	}

	kmers_inorder, err := KmerInOrderList(dna, k)
	if err != nil || len(kmers_inorder) == 0 {
		msg := fmt.Sprintf("Error: call to KmerInOrderList() failed: dna = %s, k = %d",
			dna, k)
		return "", errors.New(msg)
	}

	var max_prob_kmer string
	max_prob := float32(-1.0)
	for _, kmer := range kmers_inorder {
		probability := float32(1.0)
		for j := 0; j < len(kmer); j++ {
			ix := indexOfString(nucleotides, string(kmer[j]))
			probability *= profile[ix][j]
		}
		if probability > max_prob {
			max_prob = probability
			max_prob_kmer = kmer
		} else if probability == max_prob {
			// do nothing, be greedy
		}
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
	s.score = -1
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

	if len(s.motifs) == 0 {
		msg := fmt.Sprintf("Error: call to scored matrix motif UpdateScore() method failed: there are no motifs!")
		return errors.New(msg)
	}

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

func (s *ScoredMotifMatrix) MakeProfile(pseudocounts bool) ([][]float32, error) {
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

		if pseudocounts {
			found_zero := false
			for _, nuc := range nucleotides {
				count := counts[nuc]
				if count == 0 {
					found_zero = true
					break
				}
			}
			if found_zero {
				for _, nuc := range nucleotides {
					counts[nuc] += 1
				}
			}
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
// BA2D and BA2E functions
//
// Note: the function below is for
// BA2D and BA2E, depending on the
// value of the pseudocounts boolean.

// Given an integer k (kmer size) and t (len(dna)),
// return a collection of kmer strings
// that have the lowest score (highest similarity).
// If at any step you find more than one
// Profile-most probable k-mer in a given
// DNA string, use the one occurring first.
// Boolean pseudocounts turns on/off pseudocounts.
func GreedyMotifSearch(dna []string, k, t int, pseudocounts bool) ([]string, error) {

	var best_smm ScoredMotifMatrix

	// bestmotifs is initially an empty list with score 0
	best_smm = NewScoredMotifMatrix()

	// One way of getting kmer motifs
	// is to create a hash table of all
	// motifs that occur in the DNA string.
	// This does not maintain the original
	// order in which the motifs occur.
	// To be *greedy* we should pay attention
	// to which one comes first.
	motifs, err := KmerInOrderList(dna[0], k)
	if err != nil {
		msg := fmt.Sprintf("Error: call to KmerInOrderList() failed with params:\n\tdna = %s\n\tk = %d",
			dna[0], k)
		return nil, errors.New(msg)
	}

	for _, kmer_motif := range motifs {

		// Create a new scored motif group
		this_smm := NewScoredMotifMatrix()

		// Add our motif, which we chose from dna[0]
		// This motif kicks off the new motif group
		this_smm.AddMotif(kmer_motif)

		// Loop over all remaining dna strings
		for i := 1; i < len(dna); i++ {

			idna := dna[i]

			// Form a profile matrix from
			// all the motifs from dna strings
			// up to, but not including, this one
			profile, _ := this_smm.MakeProfile(pseudocounts)

			// Use the profile to find the profile-most
			// probable kmer in this string of dna, idna
			result, _ := ProfileMostProbableKmersGreedy(idna, k, profile)

			// Add the profile-most probable kmer
			// to the list of motifs
			if len(result) > 0 {
				this_smm.AddMotif(result)
			}
		}

		err := this_smm.UpdateScore()
		if err != nil {
			return nil, err
		}
		if this_smm.score < best_smm.score || best_smm.score < 0 {
			best_smm = this_smm
		}
	}

	return best_smm.motifs, nil
}

// ----------------------------
// BA2F functions

// Run a greedy motif search using regular counts.
func GreedyMotifSearchNoPseudocounts(dna []string, k, t int) ([]string, error) {
	return GreedyMotifSearch(dna, k, t, false)
}

// ----------------------------
// BA2E functions

// Run a greedy motif search using pseudocounts.
func GreedyMotifSearchPseudocounts(dna []string, k, t int) ([]string, error) {
	return GreedyMotifSearch(dna, k, t, true)
}

// ----------------------------
// BA2F functions

// Run a random motif search with pseudocounts.
func RandomMotifSearchPseudocounts(dna []string, k, t int) ([]string, int, error) {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	pseudocounts := true

	var result_motifs []string
	var result_score int

	// ---------------------------------
	// Fencepost algorithm:
	// Create a set of random motifs and calculate
	// their score, once, before we go into the loop.

	// Create a new scored motif group to create the first profile
	this_smm := NewScoredMotifMatrix()

	// Pick a random kmer motif from each DNA string
	var ri int
	for i := 0; i < len(dna); i++ {
		// Non-inclusive [0,n)
		overlap := len(dna[i]) - k + 1
		ri = r.Intn(overlap)
		result := dna[i][ri : ri+k]
		this_smm.AddMotif(result)
	}

	// Update the first (currently best) score
	this_smm.UpdateScore()
	best_score := this_smm.score

	// ---------------------------------
	// Main loop:
	// Get the profile from our current scored
	// motif matrix, and use it to choose the
	// profile-most probable kmers for the next
	// round.
	stop_loop := false
	for stop_loop == false {

		//fmt.Printf("----------------------------\n")
		//fmt.Printf("Current motifs = %s\n", strings.Join(this_smm.motifs, " "))
		//fmt.Printf("Current best score = %d\n", best_score)

		// Get profile from this_smm first
		profile, _ := this_smm.MakeProfile(pseudocounts)
		//fmt.Printf("Current profile = \n%v\n", profile)

		// Make a new scored motif matrix
		next_smm := NewScoredMotifMatrix()

		// Loop over all dna strings
		for i := 0; i < len(dna); i++ {

			// Use the profile to find the profile-most
			// probable kmer in this string of dna, idna
			result, _ := ProfileMostProbableKmersGreedy(dna[i], k, profile)

			// Add the profile-most probable kmer
			// to the list of motifs
			if len(result) > 0 {
				next_smm.AddMotif(result)
			}

		}
		next_smm.UpdateScore()
		next_score := next_smm.score

		//fmt.Printf("Next motifs = %s\n", strings.Join(next_smm.motifs, " "))
		//fmt.Printf("Next score = %d\n", next_score)

		if next_score < best_score {
			best_score = next_score
			this_smm = next_smm
			//fmt.Printf(" +++ Next motifs are better... continuing... new score = %d\n", best_score)
		} else {
			// This score does not improve the best score,
			// so stop now and return prior result.
			result_motifs = this_smm.motifs
			result_score = this_smm.score
			stop_loop = true
			//fmt.Printf(" --- Next motifs are not better... ending... old score = %d\n", best_score)
		}

	}
	return result_motifs, result_score, nil
}

// Driver function to run multiple random motif searches
// and keep the best of all runs.
func ManyRandomMotifSearches(dna []string, k, t, n int) ([]string, error) {
	// Initial best motifs
	min_bm, min_bm_score, _ := RandomMotifSearchPseudocounts(dna, k, t)

	// Run algorithm n times,
	// look for lowest score
	for i := 0; i < n; i += 1 {
		bm, bm_score, _ := RandomMotifSearchPseudocounts(dna, k, t)
		if bm_score < min_bm_score {
			min_bm_score = bm_score
			min_bm = bm
		}
	}
	return min_bm, nil
}

// ----------------------------
// BA2G functions

// Implement a Gibbs sampler with pseudocounts.
// The Gibbs sampler starts with random kmers,
// and samples kmers randomly generated from a
// Profile matrix. Better sampling makes the
// algorithm faster.
func GibbsSampler(dna []string, k, t, n int) ([]string, int, error) {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	pseudocounts := true

	var best_motifs []string
	var best_score int
	var ri int

	// Create a new scored motif group to create the first profile
	this_smm := NewScoredMotifMatrix()

	// Pick a random kmer motif from each DNA string
	for i := 0; i < len(dna); i++ {
		// Non-inclusive [0,n)
		overlap := len(dna[i]) - k + 1
		ri = r.Intn(overlap)
		result := dna[i][ri : ri+k]
		this_smm.AddMotif(result)
	}

	// Update the first (currently best) score
	this_smm.UpdateScore()
	best_score = this_smm.score

	// ---------------------------------
	// Main loop:
	// Pick out a random motif, and hold it out
	// while we calculate a profile and a resulting
	// profile-most-probable from all remaining
	// motifs. If the motif score improves,
	// keep the new motifs, otherwise toss 'em.
	for j := 0; j < n; j++ {

		// Non-inclusive [0,n)
		ri = r.Intn(t)

		// Now, we hold out a random motif from
		// the current scored motif matrix.

		// Make a new scored motif matrix that
		// holds out that motif, and assemble it
		// from all motifs except the holdout
		holdout_smm := NewScoredMotifMatrix()
		for i := 0; i < len(this_smm.motifs); i++ {
			if i != ri {
				holdout_smm.AddMotif(this_smm.motifs[i])
			}
		}

		// Make profile with holdout motif missing
		profile, err := holdout_smm.MakeProfile(pseudocounts)
		if err != nil {
			return nil, 0, err
		}

		// Use the profile to pick the
		// profile-most-probable kmer
		// (from DNA string ri)
		// as our candidate motif
		var candidate_motif string
		candidate_motif, err = ProfileMostProbableKmer(dna[ri], k, profile)
		if err != nil {
			return nil, 0, err
		}

		// Add candidate motif to holdout motif matrix
		holdout_smm.AddMotif(candidate_motif)

		// Update the score and use it to determine if
		// we keep the candidate motif or the old motif
		err = holdout_smm.UpdateScore()
		if err != nil {
			return nil, 0, err
		}

		// Update current scored motif matrix with
		// the candidate motif
		this_smm.motifs[ri] = candidate_motif
		err = this_smm.UpdateScore()
		if err != nil {
			return nil, 0, err
		}

		// If candidate motif leads to a better motif matrix
		// (if holdout score > this score), replace old random
		// motif with candidate motif.
		// Otherwise, candidate motif did not improve
		// the score, so keep the old motif.
		if this_smm.score < best_score {
			// Clear best_motifs and copy in this_smm.motifs
			best_motifs = []string{}
			for cc := 0; cc < len(this_smm.motifs); cc++ {
				best_motifs = append(best_motifs, this_smm.motifs[cc])
			}
			best_score = this_smm.score
		}
	}
	return best_motifs, best_score, nil
}

// Driver function to run multiple random motif searches
// and keep the best of all runs.
// n is the number of inner loops in one run of the Gibbs Sampler.
// n_starts is the number of times the Gibbs Sampler is run.
func ManyGibbsSamplers(dna []string, k, t, n, n_starts int) ([]string, error) {

	// Initial best motifs
	min_bm, min_bm_score, err := GibbsSampler(dna, k, t, n)
	if err != nil {
		return nil, err
	}

	// Run algorithm n times
	for i := 0; i < n_starts-1; i += 1 {
		// Get a new motifs and score
		bm, bm_score, err := GibbsSampler(dna, k, t, n)
		if err != nil {
			return nil, err
		}
		// If we did better, save it
		if bm_score < min_bm_score {
			min_bm_score = bm_score
			min_bm = bm
		}
	}
	return min_bm, nil
}
