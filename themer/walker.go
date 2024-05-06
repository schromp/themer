package themer

import (
	"os"

	"github.com/spf13/viper"
)

func ApplyWalker(theme Theme) error {
	if viper.GetBool("walker.enable") == false {
		return nil
	}

	configPath := viper.GetString("walker.path")

  content := GenerateGTKCSS(theme)

	writeErr := os.WriteFile(configPath, []byte(content), 0644)
	if writeErr != nil {
		return writeErr
	}

  return nil
}
