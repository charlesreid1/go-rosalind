package main

import (
    "fmt"
    "errors"
    s "strings"
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

// Reverse returns its argument string reversed 
// rune-wise left to right.
// https://github.com/golang/example/blob/master/stringutil/reverse.go
func ReverseString(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
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


// Convert four bitmasks (one each for ATGC) 
// into a DNA string.
func Bitmasks2DNA(bitmasks map[string][]bool) (string,error) {

    // Verify ATGC keys are all present
    _,Aok := bitmasks["A"]
    _,Tok := bitmasks["T"]
    _,Gok := bitmasks["G"]
    _,Cok := bitmasks["C"]
    if !(Aok && Tok && Gok && Cok) {
        err := fmt.Sprintf("Error: input bitmask was missing one of: ATGC (Keys present? A: %t, T: %t, G: %t, C: %t",Aok,Tok,Gok,Cok)
        return "", errors.New(err)
    }

    // Hope that all bitmasks are the same size
    size := len(bitmasks["A"])

    // Make a rune array that we'll turn into 
    // a string for our final return value
    dna := make([]rune,size)

    // Iterate over the bitmask, using only 
    // the index and not the mask value itself
    for i, _ := range bitmasks["A"] {
        if bitmasks["A"][i] == true {
            dna[i] = 'A'
        } else if bitmasks["T"][i] == true {
            dna[i] = 'T'
        } else if bitmasks["G"][i] == true {
            dna[i] = 'G'
        } else if bitmasks["C"][i] == true {
            dna[i] = 'C'
        }
    }

    return string(dna),nil
}


// Given a DNA input string, find the
// complement. The complement swaps
// Gs and Cs, and As and Ts.
func Complement(input string) (string,error) {

    // Convert input to uppercase
    input = s.ToUpper(input)

    // Start by checking whether we have DNA
    if CheckIsDNA(input)==false {
        return "", errors.New(fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s",input))
    }

    m,_ := DNA2Bitmasks(input)

    // Swap As and Ts
    newT := m["A"]
    newA := m["T"]
    m["T"] = newT
    m["A"] = newA

    // Swap Cs and Gs
    newG := m["C"]
    newC := m["G"]
    m["G"] = newG
    m["C"] = newC

    output,_ := Bitmasks2DNA(m)

    return output,nil
}


// Given a DNA input string, find the
// reverse complement. The complement
// swaps Gs and Cs, and As and Ts.
// The reverse complement reverses that.
func ReverseComplement(input string) (string,error) {

    // Convert input to uppercase
    input = s.ToUpper(input)

    // Start by checking whether we have DNA
    if CheckIsDNA(input)==false {
        err := fmt.Sprintf("Error: input string was not DNA. Only characters ATCG are allowed, you had %s",input)
        return "", errors.New(err)
    }

    comp,_ := Complement(input)

    revcomp := ReverseString(comp)

    return revcomp,nil
}


// Describe the problem, and call the function
func BA1C() {
    BA1CDescription()
    input := "AAAACCCGGT"
    result,_ := ReverseComplement(input)
    fmt.Printf("Reverse complement of %s is: %s\n",input,result)
}

