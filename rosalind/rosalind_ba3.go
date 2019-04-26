package rosalind

import (
	"errors"
	"fmt"
	"strings"
)

////////////////////////////////
// BA3a

// Given an input DNA string, generate a set of all
// k-mers of length k in the input string.
func KmerComposition(input string, k int) ([]string, error) {
	// Get a histogram of all kmers in this string
	hist, err := KmerHistogram(input, k)
	if err != nil {
		msg := fmt.Sprintf("Error: Function KmerHistogram(%s,%d) returned an error\n",
			input, k)
		return nil, errors.New(msg)
	}

	// Populate the string slice of kmers
	result := make([]string, len(hist))
	i := 0
	for k, _ := range hist {
		result[i] = k
		i++
	}

	// Return the string slice
	return result, nil
}

////////////////////////////////
// BA3b (new and clean)

// Given a set of kmers that overlap such that
// the last k-1 symbols of pattern i
// equal the first k-1 symbols of pattern i+1
// for all i = 1 to n-1,
// return a string of length k + n - 1
// where the ith kmer is equal to pattern i
func ReconstructGenomeFromPath(contigs []string) (string, error) {

	n := len(contigs)
	assembled := []string{}
	for i := 0; i < n; i++ {
		if i == 0 {
			// append entire contig
			assembled = append(assembled, contigs[i])
		} else {
			// convert last char to string, and append
			assembled = append(assembled, string(contigs[i][len(contigs[i])-1]))
		}
	}
	return strings.Join(assembled, ""), nil

}

////////////////////////////////
// BA3b (old and over-complicated)

// Given a genome path, i.e., a set of k-mers that
// overlap by some unknown number (up to k-1) of
// characters each, assemble the paths into a
// single string containing the genome.
//
// Note: This solved a problem that is slightly more
// general than the problem actually given - here we
// assume the number of characters overlapping is unknown,
// but the problem on Rosalind.info says it's always 1.
func ReconstructGenomeFromPath_old(contigs []string) (string, error) {

	pieces := []string{}
	for i := 0; i < len(contigs)-1; i++ {
		pattern1 := contigs[i]
		pattern2 := contigs[i+1]

		// Stride left-hand string and find where
		// it lines up to right-hand string
		overlap_index1 := -1
		overlap_index2 := -1
		for i := 0; i < len(pattern1); i++ {

			// Left-hand string: backwards-sliding window
			start1 := i           // sliding
			end1 := len(pattern1) // fixed
			slice1 := pattern1[start1:end1]

			// Right-hand string: fixed shrinking window
			start2 := 0                 // fixed
			end2 := (end1 - start2) - i // sliding
			slice2 := pattern2[start2:end2]

			if slice1 == slice2 {
				// Many Bothans died to discover this algorithm.
				overlap_index1 = start1
				overlap_index2 = end2
				break
			}
		}

		if overlap_index1 < 0 {
			msg := fmt.Sprintf("Error: ReconstructGenomeFromPath(): No overlap detected between %s and %s\n",
				pattern1, pattern2)
			return "", errors.New(msg)
		}

		// Add on the prefix of the left-hand piece - that's the part
		// that doesn't overlap with the next right-hand piece.
		pieces = append(pieces, pattern1[:overlap_index1])

		// Once we're on the last pair of pieces,
		// include the suffix of the left-hand piece -
		// the part that overlaps with the next right-hand piece -
		// then add the suffix of the right-hand piece
		// (the part that doesn't overlap with the previous
		// left-hand piece).
		if i == len(contigs)-2 {
			pieces = append(pieces, pattern1[overlap_index1:]+pattern2[overlap_index2:])
		}
	}
	return strings.Join(pieces, ""), nil
}

////////////////////////////////
// BA3c

// Given a set of k-mers, construct an overlap graph
// where each k-mer is represented by a node, and each
// directed edge represents a pair of k-mers such that
// the suffix (k-1 chars) of the k-mer at the source of
// the edge overlaps with the prefix (k-1 chars) of the
// k-mer at the head of the edge.
func OverlapGraph(patterns []string) (DirGraph, error) {

	var g DirGraph

	// Add every k-mer as a node to the overlap graph
	k := len(patterns[0])
	for _, pattern := range patterns {
		n := Node{pattern}
		g.AddNode(&n)

		// Verify k-mers are all same length
		if len(pattern) != k {
			msg := fmt.Sprintf("Error: kmer lengths do not match, k = %d but len(\"%s\") = %d\n",
				k, pattern, len(pattern))
			return g, errors.New(msg)
		}
	}

	// Iterate pairwise through the input patterns
	// to determine which pairs should have edges
	// and in which direction
	for i, pattern1 := range patterns {
		for j, pattern2 := range patterns {
			if j > i {
				prefix1 := pattern1[:k-1]
				suffix1 := pattern1[1:]

				prefix2 := pattern2[:k-1]
				suffix2 := pattern2[1:]

				if suffix1 == prefix2 {
					// 1 -> 2
					n1 := g.GetNode(pattern1)
					n2 := g.GetNode(pattern2)
					g.AddEdge(n1, n2)
				} else if suffix2 == prefix1 {
					// 2 -> 1
					n2 := g.GetNode(pattern2)
					n1 := g.GetNode(pattern1)
					g.AddEdge(n2, n1)
				}
			}
		}
	}

	return g, nil
}
