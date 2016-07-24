package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type stringReader interface {
	ReadString(byte) (string, error)
}

var (
	Query []string
)

// Run_ReadIssue runs the functions to Read Issues from the command line and prints them out
func Run_ReadIssues(in stringReader, out io.Writer) {
	fmt.Fprintf(out, "Hello World!")
}

func main() {
	Run_ReadIssues(bufio.NewReader(os.Stdin), os.Stdout)
}
