# go-rosalind

_A Go (golang) package for solving bioinformatics problems._

## Summary

This repository contains a Go (golang) library called `rosalind`
that provides a set of functions that are useful for solving
bioinformatics problems from Rosalind.info.

## Quick Start

This library can be installed using `go get`:

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

## Command Line Interface

TBA

## Organization

The repo contains the following directories:

* `chapter01/` - initial working directory; standalone code

* `chapter02/` - first set of solutions utilizing `rosalind` library

* `rosalind/` - contains Rosalind library

See the Readme file in each respective directory for more info.

