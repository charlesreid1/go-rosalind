package rosalindchapter1

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	rosa "github.com/charlesreid1/go-rosalind/rosalind"
)

// Rosalind: Problem BA1h: Find approximate occurrences of pattern in string

// Describe the problem
func BA1hDescription() {
	description := []string{
		"-----------------------------------------",
		"Rosalind: Problem BA1h:",
		"Find approximate occurrences of pattern in string",
		"",
		"Given a string Text and a string Pattern, and a maximum",
		"Hamming distance d, return all locations in Text where",
		"there is an approximate match with Pattern (i.e., a pattern",
		"with a Hamming distance from Pattern of d or less).",
		"",
		"URL: http://rosalind.info/problems/ba1h/",
		"",
	}
	for _, line := range description {
		fmt.Println(line)
	}
}

// Describe the problem, and call the function
func BA1h(filename string) {

	BA1hDescription()

	// Read the contents of the input file
	// into a single string
	lines, err := rosa.ReadLines(filename)
	if err != nil {
		log.Fatalf("Error: rosa.ReadLines: %v", err)
	}

	// Input file contents
	pattern := lines[0]
	text := lines[1]
	d_str := lines[2]

	d, _ := strconv.Atoi(d_str)

	approx, _ := rosa.FindApproximateOccurrences(pattern, text, d)

	approx_str := make([]string, len(approx))
	for i, j := range approx {
		approx_str[i] = strconv.Itoa(j)
		if err != nil {
			log.Fatalf("Error: conversion from int to string: %v", err)
		}
	}

	fmt.Println("")
	fmt.Printf("Computed result from input file: %s\n", filename)
	fmt.Println(strings.Join(approx_str, " "))
}
