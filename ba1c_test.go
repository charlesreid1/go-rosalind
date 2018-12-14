package main

import (
    "fmt"
    "testing"
)

func EqualBoolSlices(a, b []bool) bool {
    for i:=0; i<len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

// Check that the DNA2Bitmasks utility
// extracts the correct bitmasks from
// a DNA input string.
func TestDNA2Bitmasks(t *testing.T) {
    input := "AATCCGCT"

    result, func_err := DNA2Bitmasks(input)
    // Handle errors from in the DNA2Bitmasks function
    if func_err != nil {
        err := fmt.Sprintf("Error in function DNA2Bitmasks(): input = %s", input)
        t.Error(err)
    }

    // Assemble gold standard answer (bitvectors)
    tt := true
    ff := false
    gold := make(map[string][]bool)
    gold["A"] = []bool{tt,tt,ff,ff,ff,ff,ff,ff}
    gold["T"] = []bool{ff,ff,tt,ff,ff,ff,ff,tt}
    gold["C"] = []bool{ff,ff,ff,tt,tt,ff,tt,ff}
    gold["G"] = []bool{ff,ff,ff,ff,ff,tt,ff,ff}

    // Verify result from DNA2Bitmasks is same as 
    // our gold standard
    for _,cod := range "ATCG" {
        cods := string(cod)
        if !EqualBoolSlices(result[cods],gold[cods]) {
            err := fmt.Sprintf("Error testing DNA2Bitmasks(): input = %s, codon = %s, extracted = %v, gold = %v",
                        input, cods, result[cods], gold[cods])
            t.Error(err)
        }
    }
}

