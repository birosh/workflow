package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func simpleApplication() {
	(&cli.App{}).Run(os.Args)
}

func boomApplication() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!")
			fmt.Println("boom boom ", c.Args().Slice())
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func subcommandApplication() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version: %s\n", c.App.Version)
	}
	app := &cli.App{
		Name:    "boom",
		Version: "0.0.1",
		Usage:   "make an explosive entrance",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "lang",
				Value:   "english",
				Aliases: []string{"l"},
				Usage:   "language for the greeting",
			},
			&cli.BoolFlag{
				Name:  "toggle",
				Usage: "Help toggle something",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("boom! I say!", c.String("lang"))
			fmt.Println("boom boom ", c.Args().Slice())
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "add",
				Usage: "add a task to the list",
				Action: func(c *cli.Context) error {
					fmt.Println("added task: ", c.Args().First(), " in ", c.String("lang"))
					return nil
				},
			},
			{
				Name:  "complete",
				Usage: "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
