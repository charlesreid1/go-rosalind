package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func tff() {

	filename := "data/frequent_words_mismatch.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	dna := lines[1]
	params := strings.Split(lines[2], " ")
	if len(params) < 1 {
		log.Fatalf("Error splitting second line: only found 0-1 tokens")
	}
	// lines[3]: Output
	gold := strings.Split(lines[4], " ")

	k_str, d_str := params[0], params[1]

	k, err := strconv.Atoi(k_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for parameter k: %v", err)
	}

	d, err := strconv.Atoi(d_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for parameter d: %v", err)
	}

	fmt.Println("dna:", dna)
	fmt.Println("k:", k)
	fmt.Println("d:", d)
	result, err := MostFrequentKmersMismatches(dna, k, d)

	// Check if function threw error
	if err != nil {
		msg := "Error testing MostFrequentKmersMismatch using test case from file: function call threw error"
		log.Fatalf(msg)
	}

	// Check that there _was_ a result
	if len(result) == 0 {
		msg := fmt.Sprintf("Error testing MostFrequentKmersMismatch using test case from file: length of most frequent kmers found was 0: %q",
			result)
		log.Fatalf(msg)
	}

	// Sort first
	sort.Strings(gold)
	sort.Strings(result)

	fmt.Println("result:", result)
	fmt.Println("gold:", gold)

	//hist, _ := KmerHistogramMismatches(dna, k, d)
	//for i, j := range hist {
	//	if j > 70 {
	//		fmt.Println(i, ":", j)
	//	}
	//}
	//fmt.Println(hist)
	//fmt.Println(len(hist))
	//fmt.Println(CountHammingNeighbors(len(dna), d, 4))
}

func main() {
	//BA1A("for_real/rosalind_ba1a.txt")
	//BA1B("for_real/rosalind_ba1b.txt")
	//BA1C("for_real/rosalind_ba1c.txt")
	//BA1D("for_real/rosalind_ba1d.txt")
	//BA1E("for_real/rosalind_ba1e.txt")
	//BA1F("for_real/rosalind_ba1f.txt")
	//BA1G("for_real/rosalind_ba1g.txt")
	//BA1H("for_real/rosalind_ba1h.txt")
	BA1i("for_real/rosalind_ba1i.txt")
}
