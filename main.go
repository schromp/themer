package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"schromp/themer/themer"
)

func main() {
	app := &cli.App{
		Name:  "themer",
		Usage: "themer <theme>",
		Action: func(cCtx *cli.Context) error {
			themeName := cCtx.Args().Get(0)
			if themeName == "" {
				return cli.Exit("Please provide a theme name ", 1)
			}

			theme, err := themer.GetBase16(themeName)
			if err != nil {
				return cli.Exit("Error reading theme file", 2)
			}

			themer.ApplyHyprland(theme)
			themer.ApplyNvim(theme)
			themer.ApplyKitty(theme)
			themer.ApplyTmux(theme)
			themer.ApplyZsh(theme)

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
