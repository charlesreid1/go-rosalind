package rosalind

import (
	"errors"
	"fmt"
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
