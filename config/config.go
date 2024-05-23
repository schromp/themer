package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func SetupConfig() {
	viper.SetConfigName("themer")
	viper.SetConfigType("toml")
	viper.AddConfigPath("$HOME/.config/themer")

	viper.SetDefault("hyprland.enable", false)

	viper.SetDefault("kitty.enable", false)
	viper.SetDefault("kitty.path", os.Getenv("HOME") + "/.config/kitty/theme.conf")

	viper.SetDefault("nvim.enable", false)
	viper.SetDefault("nvim.path", os.Getenv("HOME") + "/.config/nvim/config.json")

	viper.SetDefault("tmux.enable", false)
	viper.SetDefault("tmux.path", os.Getenv("HOME") + "/.config/tmux/themer.conf")

	viper.SetDefault("zsh.enable", false)
	viper.SetDefault("zsh.path", os.Getenv("HOME") + "/.config/zsh/prompt.sh")

  viper.SetDefault("walker.enable", false)
	viper.SetDefault("walker.path", os.Getenv("HOME") + "/.config/walker/colors.sh")

	viper.SetDefault("swww.enable", false)
	viper.SetDefault("swww.wallpaper_dir", os.Getenv("HOME") + "/Documents/Wallpaper")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config file: %w", err))
	}
}
