# go-rosalind

`rosalind` is a Go (golang) package for solving bioinformatics problems.

[![travis](https://img.shields.io/travis/charlesreid1/go-rosalind.svg)](https://travis-ci.org/charlesreid1/go-rosalind.svg)

[![golang](https://img.shields.io/badge/language-golang-00ADD8.svg)](https://golang.org)

[![license](https://img.shields.io/github/license/charlesreid1/go-rosalind.svg)](https://github.com/charlesreid1/go-rosalind/blob/master/LICENSE)

[![godoc](https://godoc.org/github.com/charlesreid1/go-rosalind?status.svg)](http://godoc.org/github.com/charlesreid1/go-rosalind)

## Summary

This repo contains a Go (golang) library, `rosalind`, that implements
functionality for solving bioinformatics problems. This is mainly
useful for problems on Rosalind.info but is for general use as well.

Rosalind problems are grouped by chapter. Each problem has its own 
function and is implemented in a library called `chapter1`, `chapter2`,
etc.

For example, Chapter 1 question A is implemented in package
`chapter1` as the function `BA1a( <input-file-name> )`. 
This (specific) functionality wraps the (general purpose)
`rosalind` library.

## Quick Start

### Rosalind

The `rosalind` library can be installed using `go get`:

```
go get https://github.com/charlesreid1/go-rosalind/rosalind
```

The library can now be imported and its functions called directly.
Here is a brief example:

```
package main

import (
    "fmt"
    "github.com/charlesreid1/go-rosalind/rosalind"
)

func main() {
    input := "AAAATGCGCTAGTAAAAGTCACTGAAAA"
    k := 4
    result, _ := rosalind.MostFrequentKmers(input, k)
    fmt.Println(result)
}
```

### Problem Sets

Each set of problems is grouped into its own package. These
packages import the `rosalind` package, so it should be
available.

You can install the Chapter 1 problem set, for example, like so:

```
go get https://github.com/charlesreid1/go-rosalind/chapter1
```

This can now be imported and used in any Go program. 

Try creating a `main.go` file in a temporary directory,
and run it with `go run main.go`:

```
package main

import (
    rch1 "github.com/charlesreid1/go-rosalind/chapter1"
)

func main() {
    filename := "rosalind_ba1a.txt"
    rch1.BA1a(filename)
}
```

Assuming an input file `rosalind_ba1a.txt` is available,
you should see a problem description and the output of
the problem, which can be copied and pasted into
Rosalind.info:

```
$ go run main.go

-----------------------------------------
Rosalind: Problem BA1a:
Most Frequest k-mers

Given an input string and a length k,
report the k-mer or k-mers that occur
most frequently.

URL: http://rosalind.info/problems/ba1a/


Computed result from input file: for_real/rosalind_ba1a.txt
39
```

## Command Line Interface

TBA

## Organization

The repo contains the following directories:

* `rosalind/` - code and functions for the Rosalind library

* `chapter1/` - solutions to chapter 1 questions (utilizes `rosalind` library)

* `chapter2/` - solutions to chapter 2 questions

* `chapter3/` - solutions to chapter 3 questions

* `stronghold/` - solutions to questions from the stronghold section of Rosalind.info

See the Readme file in each respective directory for more info.

