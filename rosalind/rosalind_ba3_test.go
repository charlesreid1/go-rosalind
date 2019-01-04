package rosalind

import (
	"fmt"
	"log"
	"sort"
	"strconv"
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

func TestKmerCompositionFile(t *testing.T) {

	filename := "data/string_composition.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := ReadLines(filename)
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	// Input file contents
	// lines[0]: Input
	k_str := lines[1]
	k, _ := strconv.Atoi(k_str)

	input := lines[2]

	// lines[3]: Output
	// lines[4+]: gold standard answers

	// Make space for DNA strings
	iLstart := 4
	iLend := len(lines)
	gold := make([]string, len(lines)-iLstart)

	// Two counters:
	// one for the line index (iL),
	// one for the array index (iA).
	for iA, iL := 0, iLstart; iL < iLend; iA, iL = iA+1, iL+1 {
		gold[iA] = lines[iL]
	}

	results, err := KmerComposition(input, k)
	if err != nil {
		t.Error(fmt.Sprintf("Error: %v", err))
	}

	// Check that lengths are equal
	if len(results) != len(gold) {
		msg := "Error testing KmerComposition(): length of computed kmer composition does not match gold standard:"
		msg += fmt.Sprintf("len(computed) = %d, len(gold) = %d\n", len(results), len(gold))
		t.Error(msg)
	}

	// Sort before comparing
	sort.Strings(gold)
	sort.Strings(results)

	if !EqualStringSlices(results, gold) {
		msg := fmt.Sprintf("Error testing KmerComposition() from file %s:\ncomputed = %d\ngold = %d",
			filename,
			len(results), len(gold))
		t.Error(msg)
	}
}

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
	gold = strings.Trim(gold, " ")

	results, err := ReconstructGenomeFromPath(contigs)
	if err != nil {
		msg := "Error: ReconstructGenomeFromPath(): function returned an error"
		t.Error(msg)
	}

	if len(results) != len(gold) {
		msg := "Error testing ReconstructGenomeFromPath(): length of reconstructed genome does not match length of correct result\n"
		msg += fmt.Sprintf("len(computed) = %d, len(gold) = %d\n", len(results), len(gold))
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

/////////////////////////////////
// BA3c Test

func TestOverlapGraph(t *testing.T) {
	patterns := []string{"ATGCG", "GCATG", "CATGC", "AGGCA", "GGCAT"}

	g, err := OverlapGraph(patterns)
	if err != nil {
		t.Error(err)
	}

	s := g.String()
	gold := "AGGCA -> GGCAT\nCATGC -> ATGCG\nGCATG -> CATGC\nGGCAT -> GCATG"

	if s != gold {
		msg := "Error testing OverlapGraph(): string representation of graphs don't match"
		t.Error(msg)
	}
}

func TestOverlapGraphFile(t *testing.T) {

	filename := "data/overlap_graph.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := ReadLines(filename)
	if err != nil {
		log.Fatalf("ReadLines: %v", err)
	}

	// Input file contents
	// lines[0]: Input

	// We have an unknown number of fragments
	// and an unknown number of edges,
	// but they are split by a line with
	// "Output:"

	contigs := []string{}
	gold_edges := []string{}
	var stop bool

	// Loop over the first section of the file,
	// containing overlapping kmers
	stop = false
	iL := 1
	for stop == false {

		// Abort if we prematurely reach the
		// end of the file
		if iL >= len(lines) {
			msg := "Error: could not properly parse file, no line with 'Output:' found."
			t.Error(msg)
		}

		// Get the line
		line := lines[iL]

		// Break if we reached "Output:"
		if "Output:" == strings.Trim(line, " ") {
			// step over this line
			iL++
			break
		}

		// Add line to list of contigs
		contigs = append(contigs, strings.Trim(line, " "))

		iL++
	}

	// Loop over the second section of the file,
	// containing overlapping kmer edges
	stop = false
	for stop == false {

		// Break if we reach the end of the file
		if iL == len(lines) {
			break
		}

		// Get the line
		line := lines[iL]

		// Add line to list of edges
		gold_edges = append(gold_edges, strings.Trim(line, " "))

		iL++
	}

	// Construct the graph
	g, err := OverlapGraph(contigs)
	if err != nil {
		t.Error(err)
	}

	// Get the edge list representation of the graph
	computed_edges := strings.Split(g.String(), "\n")

	if !EqualStringSlices(computed_edges, gold_edges) {
		msg := fmt.Sprintf("Error testing OverlapGraph() with file %s: edge lists do not match\n", filename)
		msg += fmt.Sprintf("len(gold_edges) = %d\nlen(computed_edges) = %d\n", len(gold_edges), len(computed_edges))
		t.Error(msg)
	}

}
