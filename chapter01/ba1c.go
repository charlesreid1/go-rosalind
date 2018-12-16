package main

import (
    "fmt"
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
func BA1C() {
    BA1CDescription()
    input := "AAAACCCGGT"
    result,_ := ReverseComplement(input)
    fmt.Printf("Reverse complement of %s is: %s\n",input,result)
}

