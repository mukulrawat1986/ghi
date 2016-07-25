package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

func runListIssues(c *cli.Context) error {

	if c.NArg() < 1 {
		return cli.NewExitError("No arguments given", 10)
	}

	repo := c.Args().Get(0)

	result, err := listIssues(repo)

	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	return nil
}
