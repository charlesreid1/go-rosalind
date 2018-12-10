package main

import "fmt"

// Rosalind: Problem BA1A
//
// To run:
//
// $ go run ba1a.go

func pattern_count(input string, pattern string) int {

    // Number of substring overlaps
    var overlap = len(input) - len(pattern) + 1

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

func main() {
    // Call the pattern_count function
    fmt.Println("Number of occurrences of GCG in GCGCG:")
    res := pattern_count("GCGCG","GCG")
    fmt.Println(res)
}

