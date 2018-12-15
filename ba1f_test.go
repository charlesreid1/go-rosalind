package main

import (
    "fmt"
    "testing"
)

func TestMatrixMinSkewPosition(t *testing.T) {
    var tests = []struct {
        genome  string
        gold    []int
    }{
        {"CCTATCGGTGGATTAGCATGTCCCTGTACGTTTCGCCGCGAACTAGTTCACACGGCTTGATGGCAAATGGTTTTTCCGGCGACCGTAATCGTCCACCGAG",
                []int{53, 97}},
        {"TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGA",
                []int{11, 24}},
        {"ACCG",
                []int{3}},
        {"ACCC",
                []int{4}},
        {"CCGGGT",
                []int{2}},
        {"CCGGCCGG",
                []int{2,6}},
    }
    for _, test := range tests {
        result,err := FindClumps(test.genome,
                test.k, test.L, test.t)
        if err!=nil {
            t.Error(err)
        }
        if !EqualStringSlices(result,test.gold) {
            err := fmt.Sprintf("Error testing FindClumps(): k = %d, L = %d, t = %d",test.k,test.L,test.t)
            t.Error(err)
        }
    }
}

