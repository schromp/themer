package themer

import (
	"encoding/json"
	"fmt"
	"os"
)

type jsonConfig struct {
	Theme        string `json:"theme"`
	Transparency bool  `json:"transparency"`
}

func ApplyNvim(theme Theme, bg string) error {

	configPath := os.Getenv("HOME") + "/.config/nvim/config.json"

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

	return nil
}
