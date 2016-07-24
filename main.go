package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

	IO_ReadIssues(in)
	result, err := ReadIssues(Query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func main() {
	Run_ReadIssues(bufio.NewReader(os.Stdin), os.Stdout)
}
