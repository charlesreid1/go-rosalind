package main

import (
    "fmt"
    "log"
)

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

// Describe the problem,
// print the name of the input file,
// print the output/result
func BA1A(filename string) {

    BA1ADescription()

    // Read the contents of the input file
    // into a single string
    lines, err := readLines(filename)
    if err != nil {
        log.Fatalf("readLines: %v",err)
    }

    // Input file contents
    var input, pattern string
    input   = lines[0]
    pattern = lines[1]

    result := PatternCount(input, pattern)

    fmt.Println("")
    fmt.Printf("Computed result from input file: %s\n",filename)
    fmt.Println(result)
}

