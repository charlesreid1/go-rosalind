package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestPatternToNumber(t *testing.T) {
	input := "AGT"
	gold := 11

	// Money shot
	number, err := PatternToNumber(input)

	// Check if there was error
	if err != nil {
		t.Error(err)
	}

	if number != gold {
		err := fmt.Sprintf("Error testing PatternToNumber():\ninput = %s\ncomputed = %v\ngold     = %v\n",
			input, number, gold)
		t.Error(err)
	}
}

func TestPatternToNumberFile(t *testing.T) {

	filename := "data/pattern_to_number.txt"

	// Read the contents of the input file
	// into a single string
	lines, err := readLines(filename)
	if err != nil {
		log.Fatalf("readLines: %v", err)
	}

	// lines[0]: Input
	input := lines[1]
	// lines[2]: Output
	gold_str := lines[3]

	gold, err := strconv.Atoi(gold_str)
	if err != nil {
		log.Fatalf("Error: string to int conversion for integer representation of input DNA string: %v", err)
	}

	number, err := PatternToNumber(input)

	// Check if function threw error
	if err != nil {
		t.Error(err)
	}

	// These will only be unequal if something went wrong
	if number != gold {
		err := fmt.Sprintf("Error testing PatternToNumber():\ninput = %s\ncomputed = %v\ngold     = %v\n",
			input, number, gold)
		t.Error(err)
	}
}
