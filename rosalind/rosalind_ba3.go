package rosalind

import (
	"errors"
	"fmt"
	"sort"
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
// BA3c (new and clean)

// Construct the overlap graph of a collection of kmers.
// Given: arbitrary collection of kmers.
// Create: graph having 1 node for each kmer in kmer patterns
// Connect: kmers Pattern and Pattern' by directed edge
// if Suffix(Pattern) is equal to Prefix(Pattern')
// The resulting graph is called the overlap graph
// on these k-mers, denoted Overlap(Patterns).
//
// Return the overlap graph Overlap(Patterns),
// in the form of an adjacency list.
func OverlapGraph(patterns []string) (map[string][]string, error) {

	// Initialize prefix map
	prefix_map := make(map[string][]string)
	suffix_map := make(map[string][]string)

	// Loop over each kmer
	for _, pattern := range patterns {

		// Get the prefix and suffix
		n := len(pattern)
		prefix := pattern[0 : n-2]
		suffix := pattern[1 : n-1]

		// Insert prefix into prefix map (prefix -> kmer)
		prefix_map[prefix] = append(prefix_map[prefix], pattern)

		// Insert suffix into suffix map (suffix -> kmer)
		suffix_map[suffix] = append(suffix_map[suffix], pattern)
	}

	// Form the overlap graph:
	// - check each kmer prefix and kmer suffix
	// - create one directed edge for each pair
	// - cartesian product of all prefix-suffix matches
	// - edge is the p/s match
	// - kmers with the given suffix form the edge sources
	// - kmers with the given prefix form the edge destinations

	overlap_graph := make(map[string][]string)

	for prefix_key, prefix_kmers := range prefix_map {

		suffix_kmers, k_in_suffix_map := suffix_map[prefix_key]

		if k_in_suffix_map {

			// Create one edge per pair
			// in the Cartesian product of
			// prefix_kmers and suffix_kmers
			for _, source_kmer := range suffix_kmers {

				// Append each new destination to the source's adjacency list
				for _, destination_kmer := range prefix_kmers {

					overlap_graph[source_kmer] = append(overlap_graph[source_kmer], destination_kmer)

				}
			}
		}

	} // end for each prefix/suffix

	// map: string -> []string
	// (source to destination kmer)
	return overlap_graph, nil
}

// Utility method: given a map of string to []string,
// extract a list of all string keys, sort them, and
// return the sorted list.
func GetSortedKeys(m map[string][]string) ([]string, error) {

	// Make a list of sorted keys
	sorted_keys := make([]string, len(m))

	// Populate the list of sorted keys
	for source, _ := range m {
		sorted_keys = append(sorted_keys, source)
	}
	sort.Strings(sorted_keys)

	return sorted_keys, nil
}

// Print string representation of an overlap graph
// (map of string to []string) with the form:
// "SRC -> DEST" (no double quotes, one edge per line)
// and return the resulting string.
// The edges are ordered.
func SPrintOverlapGraph(overlap_graph map[string][]string, one_edge_per_line bool) (string, error) {

	// Get a list of sorted keys from the adjacency list
	sorted_keys, err := GetSortedKeys(overlap_graph)
	if err != nil {
		return "", err
	}

	// Initialize the string array that
	// we will use to build the final result
	var s []string

	// Now iterate over each edge,
	// ordered by source (key)
	for _, source := range sorted_keys {

		destinations := overlap_graph[source]

		// Sort the destinations too
		sort.Strings(destinations)

		if len(destinations) >= 1 {
			if one_edge_per_line {
				for _, destination := range destinations {
					s = append(s, source+" -> "+destination)
				}
			} else {
				s = append(s, source+" -> "+strings.Join(destinations, ","))
			}
		}
	}

	return strings.Join(s, "\n"), nil
}

// ////////////////////////////////
// // BA3c (old and wrong graph)
//
// // Given a set of k-mers, construct an overlap graph
// // where each k-mer is represented by a node, and each
// // directed edge represents a pair of k-mers such that
// // the suffix (k-1 chars) of the k-mer at the source of
// // the edge overlaps with the prefix (k-1 chars) of the
// // k-mer at the head of the edge.
// func OverlapGraph_old(patterns []string) (DirGraph, error) {
//
// 	var g DirGraph
//
// 	// Add every k-mer as a node to the overlap graph
// 	k := len(patterns[0])
// 	for _, pattern := range patterns {
// 		n := Node{pattern}
// 		g.AddNode(&n)
//
// 		// Verify k-mers are all same length
// 		if len(pattern) != k {
// 			msg := fmt.Sprintf("Error: kmer lengths do not match, k = %d but len(\"%s\") = %d\n",
// 				k, pattern, len(pattern))
// 			return g, errors.New(msg)
// 		}
// 	}
//
// 	// Iterate pairwise through the input patterns
// 	// to determine which pairs should have edges
// 	// and in which direction
// 	for i, pattern1 := range patterns {
// 		for j, pattern2 := range patterns {
// 			if j > i {
// 				prefix1 := pattern1[:k-1]
// 				suffix1 := pattern1[1:]
//
// 				prefix2 := pattern2[:k-1]
// 				suffix2 := pattern2[1:]
//
// 				if suffix1 == prefix2 {
// 					// 1 -> 2
// 					n1 := g.GetNode(pattern1)
// 					n2 := g.GetNode(pattern2)
// 					g.AddEdge(n1, n2)
// 				} else if suffix2 == prefix1 {
// 					// 2 -> 1
// 					n2 := g.GetNode(pattern2)
// 					n1 := g.GetNode(pattern1)
// 					g.AddEdge(n2, n1)
// 				}
// 			}
// 		}
// 	}
//
// 	return g, nil
// }

////////////////////////////////
// BA3d

// Construct the DeBruijn graph of a string.
// Given: integer k (kmer length) and DNA string text
// Return: DeBruijn graph, in the form of
// an adjacency list of (k-1)mers, which maps
// source prefixes (length k-1) to
// destination suffixes (length k-1)
// map string -> []string

func ConstructDeBruijnGraph(text string, k int) (map[string][]string, error) {

	// Create the adjacency list
	adj_map := make(map[string][]string)

	// Sliding window over each kmer
	overlap := len(text) - k + 1
	for i := 0; i < overlap; i++ {
		kmer := text[i : i+k]
		prefix := kmer[0 : k-1]
		suffix := kmer[1:k]

		// Create an adjacency map edge
		// prefix -> suffix
		adj_map[prefix] = append(adj_map[prefix], suffix)
	}

	return adj_map, nil
}
