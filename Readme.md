# go-rosalind

`rosalind` is a Go (golang) package for solving bioinformatics problems.

![](https://img.shields.io/travis/charlesreid1/go-rosalind.svg)

## Summary

This repo contains a Go (golang) library, `rosalind`, that implements
functionality for solving bioinformatics problems. This is mainly
useful for problems on Rosalind.info but is for general use as well.

Each set of questions (grouped by chapter) also has its own set of
functions in a package called `rosalindchapterXX`.
For example, Chapter 1 question A is implemented in package
`rosalindchapter01` as the function `BA1A()`. This is mainly
useful if you are answering questions on the Rosalind website.

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

