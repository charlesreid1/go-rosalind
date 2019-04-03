package rosalindchapter2

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Print problem description for Rosalind.info
// Problem BA2g: Implement GibbsSampler
func BA2gDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA2g:",
		"Implement GibbsSampler",
		"",
		"Generate probabilities of each kmer in a DNA string using its profile. Use these to assemble a list of probabilities. GibbsSampler uses this random number generator to generate a random k-mer.",
		"",
		"URL: http://rosalind.info/problems/ba2g/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Run the problem
func BA2g(filename string) {

	BA2gDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("rosa.ReadLines: %v", err)
	}

	// Input file contents
	params := strings.Split(lines[0], " ")
	k, _ := strconv.Atoi(params[0])
	t, _ := strconv.Atoi(params[1])

	// Make space for DNA strings
	dna := make([]string, len(lines)-1)
	iLstart := 1
	iLend := len(lines)
	// Two counters:
	// one for the line index (iL),
	// one for the array index (iA).
	for iA, iL := 0, iLstart; iL < iLend; iA, iL = iA+1, iL+1 {
		dna[iA] = lines[iL]
	}

	n := 100
	n_starts := 20
	result, _ := rosa.ManyGibbsSamplers(dna, k, t, n, n_starts)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(result, "\n"))
}
