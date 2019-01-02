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
// BA3b

func ReconstructGenomeFromPath(contigs []string) (string, error) {

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
