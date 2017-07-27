package main

import (
	"fmt"
	"github.com/hzlmn/msh/bundle"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "msh"
	app.Version = "1.0.0"
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
				fmt.Println(b.Graph)

				fmt.Println(c.Bool("chunks"))
				return nil
			},
		},
	}

	app.Run(os.Args)
}
