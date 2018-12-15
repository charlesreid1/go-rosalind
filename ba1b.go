package main

import (
    "fmt"
    "errors"
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

// Return the histogram of kmers of length k 
// found in the given input
func KmerHistogram(input string, k int) (map[string]int,error) {

    result := map[string]int{}

    if len(input)<1 {
        err := fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s",input)
        return result, errors.New(err)
    }

    // Number of substring overlaps
    overlap := len(input) - k + 1

    // If overlap < 1, we are looking
    // for kmers longer than our input
    if overlap<1 {
        return result,nil
    }

    // Iterate over each position,
    // extract the string,
    // increment the count.
    for i:=0; i<overlap; i++ {
        // Get the kmer of interest
        substr := input[i:i+k]

        // If it doesn't exist, the value is 0
        result[substr] += 1
    }

    return result,nil
}

// Find the most frequent kmer(s) in the kmer histogram,
// and return as a string array
func MostFrequentKmers(input string, k int) ([]string,error) {
    max := 0
    mfks := []string{}

    if k<1 {
        err := fmt.Sprintf("Error: MostFrequentKmers received a kmer size that was not a natural number: k = %d",k)
        return mfks, errors.New(err)
    }

    khist,err := KmerHistogram(input,k)

    if err != nil {
        err := fmt.Sprintf("Error: MostFrequentKmers failed when calling KmerHistogram()")
        return mfks, errors.New(err)
    }

    for kmer,freq := range khist {
        if freq > max {
            // We have a new maximum, and a new set of kmers
            max = freq
            mfks = []string{kmer}
        } else if freq==max {
            // We have another maximum
            mfks = append(mfks,kmer)
        }
    }
    return mfks,nil
}

// Describe the problem, and call the function
func BA1B() {
    BA1BDescription()
    mfks,_ := MostFrequentKmers("ACGTTGCATGTCGCATGATGCATGAGAGCT",4)
    fmt.Println("MostFrequentKmer(ACGTTGCATGTCGCATGATGCATGAGAGCT) yields:")
    fmt.Println(mfks)
}

