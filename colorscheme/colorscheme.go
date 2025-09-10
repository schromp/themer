package colorscheme

import (
	"fmt"
	"os"
	"path/filepath"
)

type Colorscheme struct {
	Name     string
	Path     string
	Programs []string
}

func GenerateColorschemes() map[string]*Colorscheme {

	colorschemeList := listColorschemes()
	colorschemes := make(map[string]*Colorscheme)

	for _, name := range colorschemeList {
		cs := &Colorscheme{
			Name: name,
		}
		colorschemes[name] = cs
	}

	return colorschemes
}

func listColorschemes() []string {

	configDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error getting user config directory:", err)
		return nil
	}
	schemesPath := filepath.Join(configDir, "themer", "colorschemes")
	entries, err := os.ReadDir(schemesPath)
	if err != nil {
		fmt.Println("Error reading colorschemes directory:", err)
		return nil
	}
	colorschemes := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			colorschemes = append(colorschemes, entry.Name())
		}
	}
	return colorschemes
}
