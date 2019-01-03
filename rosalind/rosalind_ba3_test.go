package rosalind

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"testing"
)

/////////////////////////////////
// BA3a Test

func TestKmerComposition(t *testing.T) {
	k := 5
	input := "CAATCCAAC"
	gold := []string{"AATCC", "ATCCA", "CAATC", "CCAAC", "TCCAA"}

	results, err := KmerComposition(input, k)
	if err != nil {
		t.Error(fmt.Sprintf("Error: %v", err))
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(results)

	if !EqualStringSlices(results, gold) {
		msg := fmt.Sprintf("Error testing KmerComposition()\ncomputed = %v\ngold = %v",
			results, gold)
		t.Error(msg)
	}
}

/// func TestKmerCompositionFile(t *testing.T) {
///
/// 	filename := "data/string_composition.txt"
///
/// 	// Read the contents of the input file
/// 	// into a single string
/// 	lines, err := ReadLines(filename)
/// 	if err != nil {
/// 		log.Fatalf("ReadLines: %v", err)
/// 	}
///
/// 	// Input file contents
/// 	// lines[0]: Input
/// 	k_str := lines[1]
/// 	k, _ := strconv.Atoi(k_str)
///
/// 	input := lines[2]
///
/// 	// lines[3]: Output
///
/// 	// Make space for DNA strings
/// 	gold := make([]string, len(lines)-4)
/// 	iLstart := 2
/// 	iLend := len(lines) - 2
/// 	// Two counters:
/// 	// one for the line index (iL),
/// 	// one for the array index (iA).
/// 	for iA, iL := 0, iLstart; iL < iLend; iA, iL = iA+1, iL+1 {
/// 		gold[iA] = lines[iL]
/// 	}
///
/// 	results, err := KmerComposition(input, k)
/// 	if err != nil {
/// 		t.Error(fmt.Sprintf("Error: %v", err))
/// 	}
///
/// 	// Sort before comparing
/// 	sort.Strings(gold)
/// 	sort.Strings(results)
///
/// 	if !EqualStringSlices(results, gold) {
/// 		for i, j := 0, 0; i < len(results); i, j = i+1, j+1 {
/// 			if results[i] != gold[j] {
/// 				fmt.Printf("rslt[i] = rslt[%d] = %s\ngold[j] = gold[%d] = %s\n\n",
/// 					i, results[i], j, gold[j])
/// 			}
/// 		}
/// 		msg := fmt.Sprintf("Error testing KmerComposition() from file %s:\ncomputed = %d\ngold = %d",
/// 			filename,
/// 			len(results), len(gold))
/// 		t.Error(msg)
/// 	}
/// }

/////////////////////////////////
// BA3b Test

func TestReconstructGenome(t *testing.T) {
	contigs := []string{"ACCGA", "CCGAA", "CGAAG", "GAAGC", "AAGCT"}
	gold := "ACCGAAGCT"

	results, err := ReconstructGenomeFromPath(contigs)
	if err != nil {
		t.Error(err)
	}
	if results != gold {
		msg := fmt.Sprintf("Error testing ReconstructGenomeFromPath():\ninputs = %s\ncomputed = %s\ngold     = %s",
			strings.Join(contigs, " "), results, gold)
		t.Error(msg)
	}
}

func TestReconstructGenomeFile(t *testing.T) {

	filename := "data/genome_path_string.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := ReadLines(filename)
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	// Input file contents
	// lines[0]: Input

	// Make space for DNA fragments
	contigs := make([]string, len(lines)-3)
	iLstart := 1
	iLend := len(lines) - 2
	// Two counters:
	// one for the line index (iL),
	// one for the array index (iA).
	for iA, iL := 0, iLstart; iL < iLend; iA, iL = iA+1, iL+1 {
		contigs[iA] = lines[iL]
	}

	// lines[-2]: Output
	gold := lines[len(lines)-1]

	results, err := ReconstructGenomeFromPath(contigs)
	if err != nil {
		msg := "Error: ReconstructGenomeFromPath(): function returned an error"
		t.Error(msg)
	}

	if len(results) != len(gold) {
		msg := "Error testing ReconstructGenomeFromPath(): length of reconstructed genome does not match length of correct result\n"
		msg += fmt.Sprintf("len(computed) = %d, len(gold) = %d\n", len(results), len(gold))
		fmt.Println(gold[len(gold)-10:])
		t.Error(msg)

	} else if results != gold {
		msg := "Error testing ReconstructGenomeFromPath(): computed genome and correct genome do not match\n"
		for i := 0; i < len(results); i++ {
			if results[i] != gold[i] {
				msg += fmt.Sprintf("Difference at index i = %d: computed[%d] = %s, gold[%d] = %s\n", i, i, string(results[i]), i, string(gold[i]))
			}
		}
		t.Error(msg)
	}
}
