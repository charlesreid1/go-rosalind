# Chapter 1

In this chapter we perform basic operations with 
strings and data structures.


## Directory Layout

Each problem has one corresponding Go file.
The Go file contains a function that prints
a summary of the problem, and solves the
problem for a simple case.

The bulk of functionality is in Go functions
in `rosalind.go`, to make it easier to add
and share functions.

The `main.go` file is the driver and can be
used to run individual problems.

The `*_test.go` files run tests for each
problem using the example solutions provided
by Rosalind.

## Compiling and Running

To run all tests, `go test`:

```
go test -v
```

To run a specific problem, edit `main.go`
to call the corresponding problem's function
and then `go run`:

```
go run main.go rosalind.go <name of ba1 file.go> 
```

