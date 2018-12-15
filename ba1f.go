package main

import (
    "fmt"
)

// Rosalind: Problem BA1F: Find positions in a gene that minimizing skew

// Describe the problem
func BA1FDescription() {
    description := []string{
        "-----------------------------------------",
        "Rosalind: Problem BA1F:",
        "Find positions in a gene that minimize skew",
        "",
        "The skew of a genome is defined as the difference",
        "between the number of C codons and the number of G",
        "codons. Given a DNA string, this function should",
        "compute the cumulative skew for each position in",
        "the genome, and report the indices where the skew",
        "value is minimzed.",
        "",
        "URL: http://rosalind.info/problems/ba1f/",
        "",
    }
    for _, line := range description {
        fmt.Println(line)
    }
}

// Describe the problem, and call the function
func BA1F() {
    BA1FDescription()
    //genome := "CATGGGCATCGGCCATACGCC"
    genome := "CCTATCGGTGGATTAGCATGTCCCTGTACGTTTCGCCGCGAACTAGTTCACACGGCTTGATGGCAAATGGTTTTTCCGGCGACCGTAATCGTCCACCGAG"
    minskew,_ := MinSkewPositions(genome)
    fmt.Println(minskew)
}


