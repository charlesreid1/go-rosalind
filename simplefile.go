package main

import (
    //"bufio"
    "fmt"
    "strings"
    //"io"
    "io/ioutil"
    //"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func simple() {

    dat, err := ioutil.ReadFile("PatternCount.txt")
    check(err)
    contents := string(dat)

    ix_input_start  := strings.Index(contents,"Input")
    ix_input_end    := ix_input_start + len("Input")
    ix_output_start := strings.Index(contents,"Output")
    ix_output_end   := ix_output_start + len("Output")
    ix_file_end     := len(contents)

    input_contents  := strings.Split(contents[ix_input_end:ix_output_start],"\n")
    input_contents   = input_contents[1:len(input_contents)-1]

    output_contents := strings.Split(contents[ix_output_end:ix_file_end],"\n")
    output_contents  = output_contents[1:len(output_contents)-1]

    fmt.Println("---------------")
    fmt.Println(input_contents[1])
    fmt.Println("---------------")
    fmt.Println(output_contents[0])
    fmt.Println("---------------")

    // input_contents[0] and input_contents[3] are empty
    // input_contents[1] is input 1
    // input_contents[2] is input 2
}

