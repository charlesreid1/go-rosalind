package main

import (
    "fmt"
    "testing"
)

// Utility function: check if two arrays/array slices
// are equal. This is necessary because of squirrely
// behavior when comparing arrays (of type [1]string)
// and slices (of type []string).
func EqualArrays(a, b []string) bool {
    for i:=0; i<len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TestMostFrequentKmers(t *testing.T) {
    // Call MostFrequentKmers
    input := "AAAATGCGCTAGTAAAAGTCACTGAAAA"
    k := 4
    result := MostFrequentKmers(input,k)
    gold := []string{"AAAA"}

    if !EqualArrays(result,gold) {
        err := fmt.Sprintf("Error testing MostFrequentKmers(): input = %s, k = %d, result = %s (should be %s)",
            input, k, result, gold)
        t.Error(err)
    }
}

