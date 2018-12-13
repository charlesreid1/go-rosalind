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

// Count occurrences of a substring pattern 
// in a string input
func PatternCount(input string, pattern string) int {

    // Number of substring overlaps
    var overlap = len(input) - len(pattern) + 1

    // If overlap < 1, we are looking 
    // for a pattern longer than our input
    if overlap<1 {
        return 0
    }

    // Count of occurrences
    count:=0

    // Loop over each substring overlap
    for i:=0; i<overlap; i++ {
        // Grab a slice of the full input
        start:=i
        end:=i+len(pattern)
        var slice = input[start:end]
        if slice==pattern {
            count += 1
        }
    }
    return count
}

// Describe the problem, and call the function
func BA1A() {
    BA1ADescription()
    res := PatternCount("GCGCG","GCG")
    fmt.Println("PatternCount(GCGCG,GCG) yields:",res)
}

