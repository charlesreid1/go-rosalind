package main

import (
    "fmt"
    "errors"
    s "strings"
)

// Rosalind: Problem BA1C: 

// Describe the problem
func BA1CDescription() {
    description := []string{
        "-----------------------------------------",
        "Rosalind: Problem BA1C:",
        "asdf",
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

// Given an alleged DNA input string,
// iterate through it character by character
// to ensure that it only contains ATGC.
// Returns true if this is DNA (ATGC only),
// false otherwise.
func CheckIsDNA(input string) bool {

    // Convert input to uppercase
    input = s.ToUpper(input)

    // If any character is not ATCG, fail
    for _, c := range input {
        if c!='A' && c!='T' && c!='C' && c!='G' {
            return false
        }
    }

    // If we made it here, everything's gravy!
    return true
}

// Convert a DNA string into four bitmasks:
// one each for ATGC. That is, for the DNA
// string AATCCGCT, it would become:
//
// bitmask[A] = 11000000
// bitmask[T] = 00100001
// bitmask[C] = 00011010
// bitmask[G] = 00000100
func DNA2Bitmasks(input string) (map[string][]bool,error) {

    // Convert input to uppercase
    input = s.ToUpper(input)

    // Allocate space for the map
    m := make(map[string][]bool)

    // Start by checking whether we have DNA
    if CheckIsDNA(input)==false {
        err := fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s",input)
        return m, errors.New(err)
    }

    // Important: we want to iterate over the
    // DNA string ONCE and only once. That means
    // we need to have the bit vectors initialized
    // already, and as we step through the DNA 
    // string, we access the appropriate index
    // of the appropriate bit vector and set 
    // it to true.
    m["A"] = make([]bool, len(input))
    m["T"] = make([]bool, len(input))
    m["C"] = make([]bool, len(input))
    m["G"] = make([]bool, len(input))

    // To begin with, every bit vector is false.
    for i,c := range input {
        cs := string(c)
        // Get the corresponding bit vector - O(1)
        bitty := m[cs]
        // Flip to true for this position - O(1)
        bitty[i] = true
    }

    return m,nil
}


// Given a DNA input string, find the
// reverse complement (that is, swap
// Gs and Cs, and As and Ts, and reverse
// the result).
func ReverseComplement(input string) (string,error) {

    // Convert input to uppercase
    input = s.ToUpper(input)

    var complement string

    // Start by checking whether we have DNA
    if CheckIsDNA(input)==false {
        return complement, errors.New(fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s",input))
    }

    return complement,nil
}


// Describe the problem, and call the function
func BA1C() {
    BA1CDescription()
}

