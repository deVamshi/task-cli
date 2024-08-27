package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {

	var app *application

	appCLI := &cli.App{

		Before: func(ctx *cli.Context) error {

			app = &application{}

			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action:  app.Add,
			},
			{
				Name:   "list",
				Usage:  "list all tasks",
				Action: app.List,
			},
		},
	}

	if err := appCLI.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
