package main

import (
    "fmt"
)

// Rosalind: Problem BA1B: Most Frequent k-mers

// Describe the problem
func BA1BDescription() {
    description := []string{
        "-----------------------------------------",
        "Rosalind: Problem BA1B:",
        "Most Frequest k-mers",
        "",
        "Given an input string and a length k,",
        "report the k-mer or k-mers that occur",
        "most frequently.",
        "",
        "URL: http://rosalind.info/problems/ba1b/",
        "",
    }
    for _, line := range description {
        fmt.Println(line)
    }
}

// Describe the problem, and call the function
func BA1B() {
    BA1BDescription()
    mfks,_ := MostFrequentKmers("ACGTTGCATGTCGCATGATGCATGAGAGCT",4)
    fmt.Println("MostFrequentKmer(ACGTTGCATGTCGCATGATGCATGAGAGCT) yields:")
    fmt.Println(mfks)
}

