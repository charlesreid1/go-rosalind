package main

import (
    "fmt"
    "testing"
)

func EqualIntSlices(a, b []int) bool {
    for i:=0; i<len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TestFindOccurrences(t *testing.T) {
    // Call FindOccurrences
    pattern := "ATAT"
    genome := "GATATATGCATATACTT"

    result,err := FindOccurrences(pattern,genome)
    gold := []int{1,3,9}

    if !EqualIntSlices(result,gold) || err!=nil {
        err := fmt.Sprintf("Error testing FindOccurrences(): result = %q, should be %q",
            result, gold)
        t.Error(err)
    }
}

func TestFindOccurrencesDebug(t *testing.T) {
    // Construct a test matrix
    var tests = []struct {
        pattern  string
		genome   string
		gold     []int
	}{
        {"ACAC",    "TTTTACACTTTTTTGTGTAAAAA",
                []int{4}},
        {"AAA",     "AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAGACAGATAATAATTACAGAGTACACAACATCCAT",
                []int{0,46,51,74}},
        {"TTT",     "AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCCAGCGCGTTAACCACGTTCTGTGCCGACTTT",
                []int{88,92,98,132}},
        {"ATA",     "ATATATA",
                []int{0,2,4}},
	}
    for _, test := range tests {
        fmt.Sprintf("%q",test.pattern)
        result,err := FindOccurrences(test.pattern, test.genome)
        if !EqualIntSlices(result,test.gold) || err!=nil {
            err := fmt.Sprintf("Error testing FindOccurrences(): result = %q, should be %q",
                result, test.gold)
            t.Error(err)
        }
    }
}

