package main

import (
	"github.com/hzlmn/msh/bundle"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "msh"
	app.Version = "1.0.0.beta"
	app.Usage = "Tool for writing modular shell scripts"

	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:  "bundle",
			Usage: "Bundle single entry point to specific output",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "chunks",
					Usage: "Allow chunk imports from module",
				},
				cli.StringFlag{
					Name:  "entry",
					Usage: "Provide entry file",
				},
			},
			Action: func(c *cli.Context) error {
				entryFile := c.String("entry")

				b := bundle.New(entryFile)
				b.Build()

				return nil
			},
		},
	}

	app.Run(os.Args)
}
