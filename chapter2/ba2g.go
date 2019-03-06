package rosalindchapter2

import (
	"fmt"
	"log"

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
		"Generate probabilities of each kmer in a DNA string using its profile. Use these to assemble a list of probabilities. Use these probabilities to create a random number generator. (Probability of that pattern given the Profile, Pr(Pattern|Profile), to generate n probabilities.) GibbsSampler uses this random number generator to generate a random k-mer.",
		"",
		"URL: http://rosalind.info/problems/ba2f/",
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

	//// Input file contents
	//input := lines[0]
	//params := lines[1]
	//result := rosa.PatternCount(input, pattern)
	// 
	//fmt.Println("")
	//fmt.Printf("Computed result from input file: %s\n", filename)
	//fmt.Println(result)
}
