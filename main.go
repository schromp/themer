package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"

	"schromp/themer/config"
	"schromp/themer/themer"
)

func main() {

	config.SetupConfig()

	app := &cli.App{
		Name:  "themer",
		Usage: "themer <theme>",
		Commands: []*cli.Command{
			{
				Name:  "preset",
				Usage: "",
				Action: func(cCtx *cli.Context) error {
					presets := viper.GetStringMap("presets")
					if len(presets) != 0 {
						preset := cCtx.Args().Get(0)
						themer.CallPreset(preset)
					} else {
						return fmt.Errorf("No presets defined")
					}

					return nil
				},
			},
			{
				Name:  "set",
				Usage: "",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "transparency",
						Aliases: []string{"t"},
						Value:   "solid",
						Usage:   "solid, blur, transparent",
					},
					&cli.StringFlag{
						Name:    "wallpaper",
						Aliases: []string{"w"},
						Value:   "",
						Usage: "wallpaper.png (path is set through config)",
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

					transparency := cCtx.String("transparency")
					wallpaper := cCtx.String("wallpaper")

					themer.ApplyPrograms(theme, transparency, wallpaper)

					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
