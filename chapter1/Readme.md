# Rosalind Chapter 1

This folder contains the `chapter1` module, which 
provides functions for each of the problems from
Chapter 1 of Rosalind.info's Bionformatics Textbook
track.

## How to run

* Each problem has its own function (example: `BA1a(...)`)

* Each problem expects an input file
  (example input files in `for_real` directory,
  or provide the input file downloaded
  from Rosalind.info)

* Pass the input file name to the function, like this:
  `BA1a("rosalind_ba1a.txt")`

## Quick Start

To use the functions in this package, start by installing it:

```
go get github.com/charlesreid1/go-rosalind/chapter1
```

Once you have installed the `chapter1` package,
you can import it, then call the function for whichever 
Rosalind.info problem you want to solve from Chapter 1:

```
package main

import (
    rch1 "github.com/charlesreid1/go-rosalind/chapter1"
)

func main() {
    rch1.BA1a("rosalind_ba1a.txt")
}
```

## Examples

See `chapter1_test.go` for examples.

## Tests

To run tests of all Chapter 1 problems, run
`go test` from this directory:

```
go test -v
```

or, from the parent directory, the root of the
go-rosalind repository:

```
go test -v ./chapter1/...
```

Note that this solves every problem in
Chapter 1 and prints the solutions (so there
is a lot of spew). It does not check the 
solutions (for that, see the tests in the
`rosalind` library.)

