package main

import (
    "fmt"
    "sort"
    "strings"
    "strconv"
    "io/ioutil"
    "testing"
)

// Utility function: check if two arrays/array slices
// are equal. This is necessary because of squirrely
// behavior when comparing arrays (of type [1]string)
// and slices (of type []string).
func EqualStringSlices(a, b []string) bool {
    for i:=0; i<len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

// Run a test of the MostFrequentKmers function
func TestMostFrequentKmers(t *testing.T) {
    // Call MostFrequentKmers
    input := "AAAATGCGCTAGTAAAAGTCACTGAAAA"
    k := 4
    result,err := MostFrequentKmers(input,k)
    gold := []string{"AAAA"}

    if err!=nil {
        t.Error(err)
    }

    if !EqualStringSlices(result,gold) {
        err := fmt.Sprintf("Error testing MostFrequentKmers(): input = %s, k = %d, result = %s (should be %s)",
            input, k, result, gold)
        t.Error(err)
    }
}

// Run a test of the PatternCount function
// using inputs/outputs from a file.
func TestMostFrequentKmersFile(t *testing.T) {

    // Extract inputs/outputs from file
    input, k, output := GetMostFrequentKmersFileContents()

    // Delete \r
    output = strings.Replace(output, "\r", "", -1)
    k      = strings.Replace(k,      "\r", "", -1)

    // Convert k to integer
    i64, err := strconv.ParseInt(k,10,64)
    if err!=nil {
        t.Error(err)
    }
    int_k := int(i64)

    // Split the gold standard output
    // by spaces, to get a string array slice.
    slice_output := strings.Split(output," ")

    // Call the function with the given inputs
    result, err := MostFrequentKmers(input,int_k)

    // Check if function threw error
    if err!=nil {
        t.Error(err)
    }

    // Check that there _was_ a result
    if len(result)==0 {
        err := fmt.Sprintf("Error testing MostFrequentKmers using test case from file: length of most frequent kmers found was 0: %q",result)
        t.Error(err)
    }

    // Sort before comparing
    sort.Strings(slice_output)
    sort.Strings(result)

    // These will only be unequal if something went wrong
    if !EqualStringSlices(slice_output,result) {
        err := fmt.Sprintf("Error testing MostFrequentKmers using test case from file: most frequent kmers mismatch.\ncomputed = %q\ngold = %q\n",result,slice_output)
        t.Error(err)
    }
}

// Get input and output information for the MostFrequentKmers
// test from the corresponding file.
func GetMostFrequentKmersFileContents() (string,string,string) {
    // Read the contents of the input file
    // into a single string
    dat, err := ioutil.ReadFile("data/frequent_words.txt")
    check(err)
    contents := string(dat)

    // Buncha index algebra
    ix_input_start  := strings.Index(contents,"Input")
    ix_input_end    := ix_input_start + len("Input")
    ix_output_start := strings.Index(contents,"Output")
    ix_output_end   := ix_output_start + len("Output")
    ix_file_end     := len(contents)

    input_contents  := strings.Split(contents[ix_input_end:ix_output_start],"\n")
    input_contents   = input_contents[1:len(input_contents)-1]

    output_contents := strings.Split(contents[ix_output_end:ix_file_end],"\n")
    output_contents  = output_contents[1:len(output_contents)-1]

    return input_contents[0], input_contents[1], output_contents[0]
}

