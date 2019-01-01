package rosalindchapter1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1e: Find patterns forming clumps in a string

// Describe the problem
func BA1eDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1e:",
		"Find patterns forming clumps in a string",
		"",
		"A clump is characterized by integers L and t",
		"if there is an interval in the genome of length L",
		"in which a given pattern occurs t or more times.",
		"",
		"URL: http://rosalind.info/problems/ba1e/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1e(filename string) {

	BA1eDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("Error: readLines: %v", err)
	}

	// Input file contents
	genome := lines[0]
	params_str := lines[1]
	params_slice := strings.Split(params_str, " ")

	k, _ := strconv.Atoi(params_slice[0])
	L, _ := strconv.Atoi(params_slice[1])
	t, _ := strconv.Atoi(params_slice[2])

	patterns, _ := rosa.FindClumps(genome, k, L, t)

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(patterns, " "))
}
