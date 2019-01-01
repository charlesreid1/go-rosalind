package rosalindchapter1

import (
	"bufio"
	"fmt"
	"os"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 2)

	// This is awkward.
	// Scanners aren't good for big files,
	// just simple stuff.
	BIGNUMBER := 90000
	scanner.Buffer(buf, BIGNUMBER)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// Utility function: check if two string arrays/array slices
// are equal. This is necessary because of squirrely
// behavior when comparing arrays (of type [1]string)
// and slices (of type []string).
func EqualStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Utility function: check if two boolean arrays/array slices
// are equal. This is necessary because of squirrely
// behavior when comparing arrays (of type [1]bool)
// and slices (of type []bool).
func EqualBoolSlices(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Check if two int arrays/array slices are equal.
func EqualIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Compute the factorial of an integer.
func Factorial(n int) int {
	if n < 2 {
		// base case
		return 1
	} else {
		// recursive case
		return n * Factorial(n-1)
	}
}

// Returns value of Binomial Coefficient Binom(n, k).
func Binomial(n, k int) int {

	result := 1

	// Since C(n, k) = C(n, n-k)
	if k > (n - k) {
		k = n - k
	}

	// Calculate value of:
	// ( n * (n-1) * ... * (n-k+1) )
	// -----------------------------
	//   ( k * (k-1) * ... * 1 )
	for i := 0; i < k; i++ {
		result *= n - i
		result /= i + 1
	}

	return result
}
