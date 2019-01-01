# Rosalind Chapter 2

This folder contains the `chapter2` module, which 
provides functions for each of the problems from
Chapter 2 of Rosalind.info's Bionformatics Textbook
track.

## How to run

* Each problem has its own function (example: `BA2a(...)`)

* Each problem expects an input file
  (example input files in `for_real` directory,
  or provide the input file downloaded
  from Rosalind.info)

* Pass the input file name to the function, like this:
  `BA2a("rosalind_ba2a.txt")`

## Quick Start

To use the functions in this package, start by installing it:

```
go get github.com/charlesreid1/go-rosalind/chapter2
```

Once you have installed the `chapter2` package,
you can import it, then call the function for whichever 
Rosalind.info problem you want to solve from Chapter 2:

```
package main

import (
    rch1 "github.com/charlesreid1/go-rosalind/chapter2"
)

func main() {
    rch1.BA2a("rosalind_ba2a.txt")
}
```

## Examples

See `chapter2_test.go` for examples.

## Tests

To run tests of all Chapter 2 problems, run
`go test` from this directory:

```
go test -v
```

or, from the parent directory, the root of the
go-rosalind repository:

```
go test -v ./chapter2/...
```

Note that this solves every problem in
Chapter 2 and prints the solutions (so there
is a lot of spew). It does not check the 
solutions (for that, see the tests in the
`rosalind` library.)

