# Go-Rosalind

Solving problems from Rosalind.info using Go

## Organization

Each chapter has its own directory.

Within the chapter directory, each problem has 
its own driver program, which prints info about
the problem, loads the input file from Rosalind,
and prints the solution. Each problem also has
its own test suite using the examples provided
on Rosalind.info.

For example, the function that loads the 
input file for problem BA1A is in `ba1a.go`
and the code to test the functionality
of the solution to BA1A is in `ba1a_test.go`.

## Quick Start

To run all the tests in a chapter directory:

```
go test -v
```

To run only a particular problem:

1. Edit `main.go` to call the right method
   for the right problem with the right input
   file name.

2. Run `main.go` using `go run`, and point Go
   to all the relevant Go files:

```
go run main.go utils.go rosalind.go <name-of-BA-file>
```

