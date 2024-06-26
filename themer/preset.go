package themer

import (
	"fmt"

	"github.com/spf13/viper"
)

func CallPreset(preset string) error {

  transparency := viper.GetString("presets." + preset + ".transparency")
  colorscheme := viper.GetString("presets." + preset + ".colorscheme")
  background := viper.GetString("presets." + preset + ".wallpaper")

  theme, err := GetBase16(colorscheme)
  
  if err != nil {
    fmt.Println("Bad colorscheme name")
    return err
  }

  ApplyPrograms(theme, transparency, background)

  return nil
}
