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

func (c *Colorscheme) Set() {
	fmt.Println("Setting colorscheme to:", c.Name)

	// TODO: we probably need the following here:
	// 1. Config file so we know which programs to set
	// 2. Logic of colorscheme for copying the correct files
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
