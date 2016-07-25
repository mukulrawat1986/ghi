package main

import (
	"fmt"
	"log"

	"github.com/urfave/cli"
)

func runListIssues(c *cli.Context) error {

	var terms []string

	// if the name of repository is not given, exit
	if c.NArg() < 1 {
		return cli.NewExitError("No Arguments given", 10)
	}

	// get the name of the repository from the argument
	repo := c.Args().Get(0)

	// get the value of the flag state
	state := c.String("state")

	// get the value of the author flag
	author := c.String("author")

	// since we are only searching for issues, we will restrict result to issues
	issue := "issue"

	// append the flags and arguments in terms
	terms = append(terms, "repo:"+repo)
	terms = append(terms, "type:"+issue)

	if state != "" {
		terms = append(terms, "state:"+state)
	}

	if author != "" {
		terms = append(terms, "author:"+author)
	}

	result, err := listIssues(terms)

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
