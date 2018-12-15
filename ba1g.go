package main

import (
    "fmt"
)

// Rosalind: Problem BA1G: Find Hamming distance between two DNA strings

// Describe the problem
func BA1GDescription() {
    description := []string{
        "-----------------------------------------",
        "Rosalind: Problem BA1G:",
        "Find Hamming distance between two DNA strings",
        "",
        "The Hamming distance between two strings HammingDistance(p,q)",
        "is the number of characters different between the two",
        "strands. This program computes the Hamming distance",
        "between two strings.",
        "",
        "URL: http://rosalind.info/problems/ba1g/",
        "",
    }
    for _, line := range description {
        fmt.Println(line)
    }
}

// Describe the problem, and call the function
func BA1G() {
    BA1GDescription()
    p := "GGGCCGTTGGT"
    q := "GGACCGTTGAC"
    hamm,_ := HammingDistance(p,q)
    fmt.Println(hamm)
}

