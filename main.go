package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := initApp()
	app.Run(os.Args)
}

func initApp() *cli.App {

	// Initiaize the new app
	app := cli.NewApp()

	app.Name = "ghi"
	app.Version = "0.0.1"
	app.Usage = "Github Issues on the command line"
	app.Author = "Mukul Rawat"
	app.Email = "mukulsrawat@gmail.com"

	app.Commands = []cli.Command{
		listIssuesCommand(),
	}
	return app
}

func listIssuesCommand() cli.Command {
	command := cli.Command{
		Name: "list",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "s, state",
				Usage: "enter the state of the repository, either 'open' or 'closed'.",
			},
		},
		Usage:  "List the issues of the mentioned repository",
		Action: runListIssues,
	}
	return command
}
