package main

import (
    "fmt"
    "sort"
    "errors"
    s "strings"
)


/*
rosalind.go:

This file contains core functions that 
are used to solve Rosalind problems.
*/


////////////////////////////////
// BA1A


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


////////////////////////////////
// BA1B


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
// and return as a string array slice
func MostFrequentKmers(input string, k int) ([]string,error) {
    max := 0

    // most frequent kmers
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


// Find the kmer(s) in the kmer histogram
// exceeding a count of N, and return as
// a string array slice
func MoreFrequentThanNKmers(input string, k, N int) ([]string,error) {

    // more frequent than n kmers
    mftnks := []string{}

    if k<1 || N<1 {
        err := fmt.Sprintf("Error: MoreFrequentThanNKmers received a kmer or frequency size that was not a natural number: k = %d, N = %d",k,N)
        return mftnks, errors.New(err)
    }

    khist,err := KmerHistogram(input,k)

    if err != nil {
        err := fmt.Sprintf("Error: MoreFrequentThanNKmers failed when calling KmerHistogram()")
        return mftnks, errors.New(err)
    }

    for kmer,freq := range khist {
        if freq >= N {
            // Add another more frequent than n
            mftnks = append(mftnks,kmer)
        }
    }
    return mftnks,nil
}


////////////////////////////////
// BA1C


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


////////////////////////////////
// BA1D


// Given a large string (genome) and a string (pattern),
// find the zero-based indices where pattern occurs in genome.
func FindOccurrences(pattern, genome string) ([]int,error) {
    locations := []int{}
    slots := len(genome)-len(pattern)+1

    if slots<1 {
        // pattern is longer than genome
        return locations,nil
    }

    // Loop over each character,
    // saving the position if it
    // is the start of pattern
    for i:=0; i<slots; i++ {
        start := i
        end := i+len(pattern)
        if genome[start:end]==pattern {
            locations = append(locations,i)
        }
    }
    return locations,nil
}


////////////////////////////////
// BA1E

// Find k-mers (patterns) of length k occuring at least
// t times over an interval of length L in a genome. 
func FindClumps(genome string, k, L, t int) ([]string,error) {

    // Algorithm:
    // allocate a list of kmers
    // for each possible position of L window,
    //   feed string L to KmerHistogram()
    //   save any kmers with frequency > t
    // return master list of saved kmers

    L_slots := len(genome)-L+1

    // Set kmers
    kmers := map[string]bool{}

    // List kmers
    kmers_list := []string{}

    // Loop over each possible window of length L
    for iL:=0; iL<L_slots; iL++ {

        // Grab this portion of the genome
        winstart := iL
        winend := iL+L
        genome_window := genome[winstart:winend]

        // Get the number of kmers that occur more 
        // frequently than t times
        new_kmers,err := MoreFrequentThanNKmers(genome_window,k,t)
        if err!=nil {
            return kmers_list,err
        }
        // Add these to the set kmers
        for _,new_kmer := range new_kmers {
            kmers[new_kmer] = true
        }
    }

    for k := range kmers {
        kmers_list = append(kmers_list,k)
    }
    sort.Strings(kmers_list)

    return kmers_list,nil
}


////////////////////////////////
// BA1F

// The skew of a genome is the difference between
// the number of G and C codons that have occurred
// cumulatively in a given strand of DNA.
// This function computes the positions in the genome
// at which the cumulative skew is minimized.
func MinSkewPositions(genome string) ([]int,error) {

    n := len(genome)
    cumulative_skew := make([]int,n+1)

    // Get C/G bitmasks
    bitmasks,err := DNA2Bitmasks(genome)
    if err!=nil {
        return cumulative_skew,err
    }
    c := bitmasks["C"]
    g := bitmasks["G"]

    // Init
    cumulative_skew[0] = 0

    // Make space to keep track of the 
    // minima we have encountered so far
    min := 999
    min_skew_ix := []int{}

    // At each position, compute the next skew value.
    // We need two indices b/c for a genome of size N,
    // the cumulative skew array index is of size N+1.
    for i,ibit:=1,0; i<=n; i,ibit=i+1,ibit+1 {

        var next int
        // Next skew value
        if c[ibit] {
            // C -1
            next = -1
        } else if g[ibit] {
            // G +1
            next = 1
        } else {
            next = 0
        }
        cumulative_skew[i] = cumulative_skew[i-1] + next

        if cumulative_skew[i] < min {
            // New min and min_skew
            min = cumulative_skew[i]
            min_skew_ix = []int{i}
        } else if cumulative_skew[i] == min {
            // Additional min and min_skew
            min_skew_ix = append(min_skew_ix,i)
        }

    }
    return min_skew_ix,nil
}


////////////////////////////////
// BA1G

// Compute the Hamming distance between
// two strings. The Hamming distance is
// defined as the number of characters
// different between two strings.
func HammingDistance(p, q string) (int,error) {

    // Technically a Hamming distance when
    // one string is empty would be 0, but
    // we will throw an error instead.
    if len(p)==0 || len(q)==0 {
        err := fmt.Sprintf("Error: HammingDistance: one or more arguments had length 0. len(p) = %d, len(q) = %d",len(p),len(q))
        return -1,errors.New(err)
    }

    // Get longest length common to both
    var m int
    if len(p)>len(q) {
        m = len(q)
    } else {
        m = len(p)
    }

    // Accumulate distance
    dist := 0
    for i:=0; i<m; i++ {
        if p[i]!=q[i] {
            dist += 1
        }
    }
    return dist,nil
}


////////////////////////////////
// BA1H


// Given a large string (text) and a string (pattern),
// find the zero-based indices where we have an occurrence
// of pattern or a string with Hamming distance d or less
// from pattern.
func FindApproximateOccurrences(pattern, text string, d int) ([]int,error) {

    locations := []int{}
    slots := len(text)-len(pattern)+1

    if slots<1 {
        // pattern is longer than genome
        return locations,nil
    }

    // Loop over each character,
    // saving the position if it
    // is the start of pattern
    for i:=0; i<slots; i++ {
        start := i
        end := i+len(pattern)
        poss_approx_pattern := text[start:end]
        hamm,_ := HammingDistance(poss_approx_pattern,pattern)
        if hamm<=d {
            locations = append(locations,i)
        }
    }

    return locations,nil
}


