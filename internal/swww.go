package themer

import (
	"os/exec"

	"github.com/spf13/viper"
)

func ApplySwww(wallpaper string) error {

	if viper.GetBool("swww.enable") == false || wallpaper == "" {
		return nil
	}

	folder_path := viper.GetString("swww.wallpaper_dir")
	wallpaper_path := folder_path + "/" + wallpaper

	err := exec.Command("swww", "img", wallpaper_path).Run()
	if err != nil {
		return err
	}

	return nil
}
