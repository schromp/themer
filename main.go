package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"schromp/themer/themer"
	"schromp/themer/config"
)

func main() {

	config.SetupConfig()

	app := &cli.App{
		Name:  "themer",
		Usage: "themer <theme>",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "background",
				Aliases: []string{"b"},
				Value:   "solid",
				Usage:   "solid, blur, transparent",
			},
		},
		Action: func(cCtx *cli.Context) error {
			themeName := cCtx.Args().Get(0)
			if themeName == "" {
				return cli.Exit("Please provide a theme name ", 1)
			}

			theme, err := themer.GetBase16(themeName)
			if err != nil {
				return cli.Exit("Error reading theme file", 2)
			}

			background := cCtx.String("background")

			themer.ApplyHyprland(theme, background)
			themer.ApplyNvim(theme, background)
			themer.ApplyKitty(theme, background)
			themer.ApplyTmux(theme)
			themer.ApplyZsh(theme)
			themer.ApplyWalker(theme)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
