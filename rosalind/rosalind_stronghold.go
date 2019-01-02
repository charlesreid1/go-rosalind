package rosalind

import (
	"errors"
	"fmt"
)

// Count the number of each type of nucleotide ACGT.
func CountNucleotides(dna string) (map[string]int, error) {

	if !CheckIsDNA(dna) {
		msg := fmt.Sprintf("Error: input string was not DNA: %s", dna)
		return nil, errors.New(msg)
	}

	// Map to store counts for each nucleotide
	result := make(map[string]int)

	// Get bitmask representations
	bms, err := DNA2Bitmasks(dna)

	if err != nil {
		msg := fmt.Sprintf("Error: DNA2Bitmasks() threw an error for input %s",
			dna)
		return nil, errors.New(msg)
	}

	// Iterate over every possible nucleotide
	bases := []string{"A", "C", "G", "T"}
	for _, base := range bases {

		// Bitmap for this nucleotide
		bm := bms[base]

		// Frequency for this nucleotide
		sum := 0
		for j := 0; j < len(bm); j++ {
			if bm[j] {
				sum++
			}
		}

		// Store the result
		result[base] = sum
	}

	return result, nil
}

// Count the number of each type of nucleotide ACGT
// and return as an array in order A, C, G, T.
func CountNucleotidesArray(dna string) ([]int, error) {
	result := make([]int, 4)
	mresult, err := CountNucleotides(dna)
	if err != nil {
		msg := fmt.Sprintf("Error: CountNucleotides() returned an error: %v", err)
		return nil, errors.New(msg)
	}
	result[0] = mresult["A"]
	result[1] = mresult["C"]
	result[2] = mresult["G"]
	result[3] = mresult["T"]
	return result, nil
}
