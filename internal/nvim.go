package themer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/viper"
)

type jsonConfig struct {
	Theme        string `json:"theme"`
	Transparency bool   `json:"transparency"`
}

func hotReloadNvim(theme string) {

	runtime_dir := os.Getenv("XDG_RUNTIME_DIR")

	files, err := os.ReadDir(runtime_dir)
	if err != nil {
		fmt.Println("Error:", err)
	}

	for _, value := range files {
		match, _ := regexp.MatchString(`^nvim.*\.0$`, value.Name())
		if match {
			path := runtime_dir + "/" + value.Name()
			reload_err := exec.Command("nvim", "--server", path, "--remote-expr", fmt.Sprintf("execute('colorscheme %s')", theme)).Run()
			if reload_err != nil {
				fmt.Println("Error:", reload_err)
			}
		}
	}

}

func ApplyNvim(theme Theme, bg string) error {

	if viper.GetBool("nvim.enable") == false {
		return nil
	}

	configPath := viper.GetString("nvim.path")
	// configPath := os.Getenv("HOME") + "/.config/nvim/config.json"

	themeName := theme.name
	if theme.NvimName != "" {
		themeName = theme.NvimName
	}

	var transparency bool

	if bg == "solid" {
		transparency = false
	} else if bg == "transparent" || bg == "blur" {
		transparency = true
	} else {
		return fmt.Errorf("Invalid background type")
	}

	config := jsonConfig{
		Theme:        themeName,
		Transparency: transparency,
	}

	encoded, err := json.Marshal(config)
	if err != nil {
		return err
	}

	writeErr := os.WriteFile(configPath, encoded, 0644)
	if writeErr != nil {
		return writeErr
	}

	hotReloadNvim(themeName)

	return nil
}
