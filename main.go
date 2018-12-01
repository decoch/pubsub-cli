package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/decoch/pubsub-cli/action"
	"github.com/urfave/cli"
)

var (
	Version = "No Version Provided"
)

func main() {
	app := cli.NewApp()

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, V",
		Usage: "Show version number and quit",
	}

	app.Name = "pubsub-cli"
	app.Usage = "A command line tool for Google Cloud Pub/Sub."
	app.Version = Version

	app.Commands = []cli.Command{
		{
			Name:  "list",
			Usage: "Get all subscription.",
			Action: func(c *cli.Context) error {
				args := c.Args()
				if l := len(args); l == 0 {
					return errors.New("project id is not specified")

				} else if l > 2 {
					return errors.New("too many args")
				}
				projectID := args[0]
				filename := ""
				if len(args) > 1 {
					filename = args[1]
				}
				ctx := context.Background()
				err := action.List(ctx, projectID, filename)
				if err != nil {
					return err
				}
				return nil
			},
		},
		{
			Name:  "create",
			Usage: "Create subscription of Pub/Sub.",
			Action: func(c *cli.Context) error {
				args := c.Args()
				if l := len(args); l < 2 {
					return errors.New("too less args")

				} else if l > 3 {
					return errors.New("too many args")
				}
				projectID := args[0]
				filename := args[1]
				ctx := context.Background()
				err := action.Create(ctx, projectID, filename)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
