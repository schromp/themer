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

	viper.SetDefault("hyprland.enable", true)

	viper.SetDefault("kitty.enable", true)
	viper.SetDefault("kitty.path", os.Getenv("HOME") + "/.config/kitty/theme.conf")

	viper.SetDefault("nvim.enable", true)
	viper.SetDefault("nvim.path", os.Getenv("HOME") + "/.config/nvim/config.json")

	viper.SetDefault("tmux.enable", true)
	viper.SetDefault("tmux.path", os.Getenv("HOME") + "/.config/tmux/themer.conf")

	viper.SetDefault("zsh.enable", true)
	viper.SetDefault("zsh.path", os.Getenv("HOME") + "/.config/zsh/prompt.sh")

  viper.SetDefault("walker.enable", true)
	viper.SetDefault("walker.path", os.Getenv("HOME") + "/.config/walker/colors.sh")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config file: %w", err))
	}
}
