package themer

import (
	"os"

	"github.com/tidwall/sjson"
)

func ApplyNvim(theme Theme) error {
	configPath := os.Getenv("HOME") + "/.config/nvim/config.json"
	themeName := theme.name
	if theme.NvimName != "" {
		themeName = theme.NvimName
	}

	nvimConfig, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	val, jsonErr := sjson.Set(string(nvimConfig), "theme", themeName)
	if jsonErr != nil {
		return jsonErr
	}

	writeErr := os.WriteFile(configPath, []byte(val), 0644)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
