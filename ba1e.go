package main

import (
    "fmt"
)

// Rosalind: Problem BA1E: Find patterns forming clumps in a string

// Describe the problem
func BA1EDescription() {
    description := []string{
        "-----------------------------------------",
        "Rosalind: Problem BA1E:",
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
func BA1E() {
    BA1EDescription()
    genome := "CGGACTCGACAGATGTGAAGAAATGTGAAGACTGAGTGAAGAGAAGAGGAAACACGACACGACATTGCGACATAATGTACGAATGTAATGTGCCTATGGC"
    k := 5
    L := 75
    t := 4
    patterns,_ := FindClumps(genome,k,L,t)
    fmt.Println(patterns)
}

