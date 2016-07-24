package main_test

import (
	"bytes"
	"testing"

	"github.com/mukulrawat1986/ghi"
	"github.com/stretchr/testify/assert"
)

func Test_Run_ReadIssues(t *testing.T) {

	a := assert.New(t)

	// create input using bytes Buffer
	b := []byte("repo:golang/go is:open json decoder\n")
	r := bytes.NewBuffer(b)

	// use bytes Buffer to create a sink for output
	w := &bytes.Buffer{}

	main.Run_ReadIssues(r, w)

	res := w.String()

	a.Contains(res, "Hello World!")
}
