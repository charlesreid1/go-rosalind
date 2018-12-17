package main

import (
    "fmt"
    "log"
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
func BA1D(filename string) {

    BA1DDescription()

    // Read the contents of the input file
    // into a single string
    lines, err := readLines(filename)
    if err != nil {
        log.Fatalf("Error: readLines: %v",err)
    }

    // Input file contents
    genome  := lines[0]
    pattern := lines[1]

    locs,_ := FindOccurrences(pattern,genome)

    fmt.Println("")
    fmt.Printf("Computed result from input file: %s\n",filename)
    fmt.Println(locs)
}

