package main

import (
	"github.com/prongbang/localizegen/cmd"
	"github.com/prongbang/localizegen/pkg/arguments"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "platform",
				Value: "",
				Usage: "-platform android, ios",
			},
			&cli.StringFlag{
				Name:  "target",
				Value: "",
				Usage: "-target path",
			},
			&cli.StringFlag{
				Name:  "filename",
				Value: "",
				Usage: "-filename name",
			},
			&cli.StringFlag{
				Name:  "locale",
				Value: "",
				Usage: "-locale en, th",
			},
			&cli.StringFlag{
				Name:  "document",
				Value: "",
				Usage: "-document document-id",
			},
			&cli.StringFlag{
				Name:  "sheet",
				Value: "",
				Usage: "-sheet sheet-id",
			},
		},
		Action: func(c *cli.Context) error {
			return cmd.Run(arguments.Get(c))
		},
	}

	_ = app.Run(os.Args)
}
