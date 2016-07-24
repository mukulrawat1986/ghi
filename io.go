package main

import (
	"fmt"
	"strings"
)

func IO_ReadIssues(r stringReader) {
	fmt.Println("Enter your query: ")
	q, _ := r.ReadString('\n')
	q = strings.TrimSpace(q)
	Query = strings.Split(q, " ")
}
