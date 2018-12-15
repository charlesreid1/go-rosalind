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


// Describe the problem, and call the function
func BA1D() {
    BA1DDescription()
    genome := "GATATATGCATATACTT"
    pattern := "ATAT"
    locs,_ := FindOccurrences(pattern,genome)
    fmt.Println(locs)
}

