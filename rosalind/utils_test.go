package rosalind

import "testing"

func TestEqualStringSlices(t *testing.T) {
	a := []string{"peanut", "butter", "jelly", "time"}
	b := make([]string, 4)
	b[0] = "peanut"
	b[1] = "butter"
	b[2] = "jelly"
	b[3] = "time"
	if !EqualStringSlices(a, b) {
		msg := "Error: EqualStringSlices() is broken!"
		t.Fatal(msg)
	}
}

func TestEqualBoolSlices(t *testing.T) {
	a := []bool{true, true, true, false, false, false, true, true, true}
	b := make([]bool, 9)
	b[0], b[1], b[2] = true, true, true
	b[3], b[4], b[5] = false, false, false
	b[6], b[7], b[8] = true, true, true
	if !EqualBoolSlices(a, b) {
		msg := "Error: EqualBoolSlices() is broken!"
		t.Fatal(msg)
	}
}

func TestEqualIntSlices(t *testing.T) {
	a := []int{3, 1, 4, 1, 5, 9}
	b := make([]int, 6)
	b[0], b[1], b[2] = 3, 1, 4
	b[3], b[4], b[5] = 1, 5, 9
	if !EqualIntSlices(a, b) {
		msg := "Error: EqualIntSlices() is broken!"
		t.Fatal(msg)
	}
}
