# Chapter 1

In this chapter we perform basic operations with 
strings and data structures.

## How to run

* Each problem has its own function

* To run the code for a particular problem,
  call the function for that problem in `main.go`

* Edit `main.go` to call the right function,
  and pass in the name of the input file you 
  want to use: for example, `BA1A("input.txt")`
  
* The function you call is implemented in the
  corresponding Go file (for example, `ba1a.go`).
  It loads the inputs from the input file,
  calls the right function with the inputs,
  and prints the results.

* The functions that load data from input files
  are tested along with the functions themselves,
  since each problem has a sample input file 
  in `data/`

## Directory Layout

* Each problem has one Go file and one test

* The `data/` directory contains input files
  for the tests (i.e., files that contain both
  inputs and corresponding outputs)

* The `for_real/` directory contains sample 
  input files from Rosalind.info for each
  problem (i.e., files that contain only the
  inputs)

* The `main.go` file contains the `main()` 
  driver function and is the entrypoint for
  `go run`

* The `rosalind.go` file contains most of the
  computational functionality implemented
  for the problems.

* The `utils.go` file contains utilties unrelated
  to bioinformatics.

## Compiling and Running

To run all tests, `go test`:

```
go test -v
```

To run a specific problem, edit `main.go`
to call the corresponding problem's function
and then `go run`:

```
go run main.go utils.go rosalind.go <name of ba1 file.go> 
```

## To Do

Add a Snakefile



