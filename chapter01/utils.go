package main

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
    for i:=0; i<len(a); i++ {
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
    for i:=0; i<len(a); i++ {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

