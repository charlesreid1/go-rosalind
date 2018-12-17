package main

import (
    "fmt"
    "log"
)

// Rosalind: Problem BA1C: Find the Reverse Complement of a String

// Describe the problem
func BA1CDescription() {
    description := []string{
        "-----------------------------------------",
        "Rosalind: Problem BA1C:",
        "Find the Reverse Complement of a String",
        "",
        "Given a DNA input string,",
        "find the reverse complement",
        "of the DNA string.",
        "",
        "URL: http://rosalind.info/problems/ba1c/",
        "",
    }
    for _, line := range description {
        fmt.Println(line)
    }
}

// Describe the problem, and call the function
func BA1C(filename string) {

    BA1CDescription()

    // Read the contents of the input file
    // into a single string
    lines, err := readLines(filename)
    if err != nil {
        log.Fatalf("Error: readLines: %v",err)
    }

    // Input file contents
    input := lines[0]

    result,_ := ReverseComplement(input)

    fmt.Println("")
    fmt.Printf("Computed result from input file: %s\n",filename)
    fmt.Println(result)
}

