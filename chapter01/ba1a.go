package main

import "fmt"

// Rosalind: Problem BA1A: Most Frequent k-mers

// Describe the problem
func BA1ADescription() {
    description := []string{
        "-----------------------------------------",
        "Rosalind: Problem BA1A:",
        "Most Frequest k-mers",
        "",
        "Given an input string and a length k,",
        "report the k-mer or k-mers that occur",
        "most frequently.",
        "",
        "URL: http://rosalind.info/problems/ba1a/",
        "",
    }
    for _, line := range description {
        fmt.Println(line)
    }
}

// Describe the problem, and call the function
func BA1A() {
    BA1ADescription()
    res := PatternCount("GCGCG","GCG")
    fmt.Println("PatternCount(GCGCG,GCG) yields:",res)
}

