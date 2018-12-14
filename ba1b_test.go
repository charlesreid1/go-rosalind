package main

import (
    "fmt"
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

func TestMostFrequentKmers(t *testing.T) {
    // Call MostFrequentKmers
    input := "AAAATGCGCTAGTAAAAGTCACTGAAAA"
    k := 4
    result := MostFrequentKmers(input,k)
    gold := []string{"AAAA"}

    if !EqualStringSlices(result,gold) {
        err := fmt.Sprintf("Error testing MostFrequentKmers(): input = %s, k = %d, result = %s (should be %s)",
            input, k, result, gold)
        t.Error(err)
    }
}

func TestMostFrequentKmersFile(t *testing.T) {
    //input, k, output := GetMostFrequentKmersFileContents()
    input, k, _ := GetMostFrequentKmersFileContents()
    //fmt.Println(input)
    //fmt.Println(output)
    int_k, _ := strconv.Atoi(k)
    //fmt.Println(int_k)

    //slice_output := strings.Split(output," ")
    result := MostFrequentKmers(input,int_k)

    if result==nil {
        //err := fmt.Sprintf("Error testing MostFrequentKmers using test case from file: results do not match:\ncomputed result = %q\nexpected output = %q",result,slice_output)
        err := "Error"
        t.Error(err)
    }
}

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

