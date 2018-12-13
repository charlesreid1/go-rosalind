package main

import (
    "fmt"
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
