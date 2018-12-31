# go-rosalind

`rosalind` is a Go (golang) package for solving bioinformatics problems.

![https://travis-ci.org/charlesreid1/go-rosalind/](https://img.shields.io/travis/charlesreid1/go-rosalind.svg)
![https://golang.org/](https://img.shields.io/badge/language-golang-00ADD8.svg)
![https://github.com/charlesreid1/go-rosalind](https://img.shields.io/github/license/charlesreid1/go-rosalind.svg)

## Summary

This repo contains a Go (golang) library, `rosalind`, that implements
functionality for solving bioinformatics problems. This is mainly
useful for problems on Rosalind.info but is for general use as well.

Rosalind problems are grouped by chapter. Each problem has its own 
function and is implemented in a library called `chapter01`, `chapter02`,
etc.

For example, Chapter 1 question A is implemented in package
`chapter01` as the function `BA1A( <input-file-name> )`. 
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
go get https://github.com/charlesreid1/go-rosalind/chapter01
```

This can now be imported and used in any Go program. Try creating
a new Go program in a temporary directory and running it with
`go run`:

```
package main

import (
    rch1 "github.com/charlesreid1/go-rosalind/chapter01"
)

func main() {
    filename := "rosalind_ba1a.txt"
    rch1.BA1A(filename)
}
```

Assuming an input file is available, you should see a problem description
and the output of the problem, which can be copied and pasted into 
Rosalind.info.

## Command Line Interface

TBA

## Organization

The repo contains the following directories:

* `chapter01/` - initial working directory; standalone code

* `chapter02/` - first set of solutions utilizing `rosalind` library

* `rosalind/` - contains Rosalind library

See the Readme file in each respective directory for more info.

