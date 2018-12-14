package main

import (
    "fmt"
    "strings"
    "strconv"
    "io/ioutil"
    "testing"
)

// To run this test:
//
// $ go test -v -run TestPatternCount

// Run a single test of the PatternCount function
func TestPatternCount(t *testing.T) {
    // Call the PatternCount function
    input := "GCGCG"
    pattern := "GCG"
    result := PatternCount(input,pattern)
    gold := 2
    if result != gold {
        err := fmt.Sprintf("Error testing PatternCount(): input = %s, pattern = %s, result = %d (should be %d)",
            input, pattern, result, gold)
        t.Error(err)
    }
}

// Run a test matrix of the PatternCount function
func TestPatternCounts(t *testing.T) {
    // Construct a test matrix
    var tests = []struct {
		input    string
        pattern  string
		gold int
	}{
		{"GCGCG",        "GCG",      2},
		{"GAGGGGGGGAG",  "AGG",      1},
		{"GCACGCACGCAC", "GCAC",     3},
        {"",             "GC",       0},
        {"GCG",          "GTACTCTC", 0},
        {"ACGTACGTACGT", "CG",       3},
        {"AAAGAGTGTCTGATAGCAGCTTCTGAACTGGTTACCTGCCGTGAGTAAATTAAATTTTATTGACTTAGGTCACTAAATACTTTAACCAATATAGGCATAGCGCACAGACAGATAATAATTACAGAGTACACAACATCCA",
                         "AAA",      4},
        {"AGCGTGCCGAAATATGCCGCCAGACCTGCTGCGGTGGCCTCGCCGACTTCACGGATGCCAAGTGCATAGAGGAAGCGAGCAAAGGTGGTTTCTTTCGCTTTATCCAGCGCGTTAACCACGTTCTGTGCCGACTTT",
                         "TTT",      4},
        {"GGACTTACTGACGTACG","ACT",  2},
        {"ATCCGATCCCATGCCCATG","CC", 5},
        {"CTGTTTTTGATCCATGATATGTTATCTCTCCGTCATCAGAAGAACAGTGACGGATCGCCCTCTCTCTTGGTCAGGCGACCGTTTGCCATAATGCCCATGCTTTCCAGCCAGCTCTCAAACTCCGGTGACTCGCGCAGGTTGAGT",
                         "CTC",      9},
	}
    for _, test := range tests {
        result := PatternCount(test.input, test.pattern)
        if result != test.gold {
            err := fmt.Sprintf("Error testing PatternCount(): input = %s, pattern = %s, result = %d (should be %d)",
                        test.input, test.pattern, result, test.gold)
            t.Error(err)
        }
    }
}


// Load a PatternCount test (input and output)
// from a file. Run the test with the input
// and verify the output matches the output
// contained in the file.
func TestPatternCountFile(t *testing.T) {
    input, pattern, output := GetPatternCountFileContents()
    int_output, _ := strconv.Atoi(output)
    result := PatternCount(input, pattern)
    if result != int_output {
        err := fmt.Sprintf("Error testing PatternCount using test case from file: results do not match:\rcomputed result = %d\nexpected output = %d",result,int_output)
        t.Error(err)
    }
}


// Load the contents of the PatternCount file
// and return the inputs and outputs for that
// pattern count test.
func GetPatternCountFileContents() (string,string,string) {
    // Read the contents of the input file
    // into a single string
    dat, err := ioutil.ReadFile("data/pattern_count.txt")
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

    // Two inputs:
    // input_contents[0:1]
    // One output:
    // output_contents[0]
    return input_contents[0], input_contents[1], output_contents[0]
}

