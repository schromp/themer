package themer

import "os"

func ApplyWalker(theme Theme) error {
  content := GenerateGTKCSS(theme)

	configPath := os.Getenv("HOME") + "/.config/walker/colors.css"

	writeErr := os.WriteFile(configPath, []byte(content), 0644)
	if writeErr != nil {
		return writeErr
	}

  return nil
}
