package main_test

import (
	"bytes"
	"testing"

	"github.com/mukulrawat1986/ghi"
	"github.com/stretchr/testify/assert"
)

func Test_IO_ReadIssues(t *testing.T) {

	a := assert.New(t)

	// create a slice of bytes to seed the bytes buffer
	b := []byte("repo:golang/go is:open json decoder\n")

	// create a bytes buffer
	r := bytes.NewBuffer(b)

	// pass it to the IO_ReadIssues function
	main.IO_ReadIssues(r)

	a.Equal(len(main.Query), 4)
	a.Equal(main.Query[0], "repo:golang/go")
	a.Equal(main.Query[1], "is:open")
	a.Equal(main.Query[2], "json")
	a.Equal(main.Query[3], "decoder")
}
