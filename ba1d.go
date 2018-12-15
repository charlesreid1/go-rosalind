package main

import (
    "fmt"
)

// Rosalind: Problem BA1D: Find all occurrences of pattern in string

// Describe the problem
func BA1DDescription() {
    description := []string{
        "-----------------------------------------",
        "Rosalind: Problem BA1D:",
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

// Given a large string (genome) and a substring (pattern),
// find the zero-based indices where pattern occurs in genome.
func FindOccurrences(pattern, genome string) ([]int,error) {
    locations := []int{}
    slots := len(genome)-len(pattern)+1

    if slots<1 {
        // pattern is longer than genome
        return locations,nil
    }

    // Loop over each character,
    // saving the position if it
    // is the start of pattern
    for i:=0; i<slots; i++ {
        start := i
        end := i+len(pattern)
        if genome[start:end]==pattern {
            locations = append(locations,i)
        }
    }
    return locations,nil
}


// Describe the problem, and call the function
func BA1D() {
    BA1DDescription()
    genome := "GATATATGCATATACTT"
    pattern := "ATAT"
    locs,_ := FindOccurrences(pattern,genome)
    fmt.Println(locs)
}

