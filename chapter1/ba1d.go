package rosalindchapter1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1d: Find all occurrences of pattern in string

// Describe the problem
func BA1dDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1d:",
		"Find all occurrences of pattern in string",
		"",
		"Given a string input (genome) and a substring (pattern),",
		"return all starting positions in the genome where the",
		"pattern occurs in the genome.",
		"",
		"URL: http://rosalind.info/problems/ba1d/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1d(filename string) {

	BA1dDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("Error: rosa.ReadLines: %v", err)
	}

	// Input file contents
	pattern := lines[0]
	genome := lines[1]

	// Result is a slice of ints
	locs, _ := rosa.FindOccurrences(pattern, genome)

	// Convert to a slice of strings for easier printing
	locs_str := make([]string, len(locs))
	for i, j := range locs {
		locs_str[i] = strconv.Itoa(j)
	}

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(locs_str, " "))
}
