package themer

import (
	"os"
)

func ApplyKitty(theme Theme) error {

	var content string

	content += "background #" + theme.Base00 + "\n"
	content += "foreground #" + theme.Base08 + "\n"
	content += "selection_background #" + theme.Base05 + "\n"
	content += "selection_foreground #" + theme.Base00 + "\n"
  content += "url_color #" + theme.Base04 + "\n"
  content += "cursor #" + theme.Base05 + "\n"
  content += "active_border_color #" + theme.Base03 + "\n"
  content += "inactive_border_color #" + theme.Base01 + "\n"
  content += "active_tab_background #" + theme.Base00 + "\n"
  content += "active_tab_foreground #" + theme.Base05 + "\n"
  content += "inactive_tab_background #" + theme.Base01 + "\n"
  content += "inactive_tab_foreground #" + theme.Base04 + "\n"
  content += "tab_bar_background #" + theme.Base01 + "\n"

  content += "color0 #" + theme.Base00 + "\n"
  content += "color1 #" + theme.Base08 + "\n"
  content += "color2 #" + theme.Base0B + "\n"
  content += "color3 #" + theme.Base0A + "\n"
  content += "color4 #" + theme.Base0D + "\n"
  content += "color5 #" + theme.Base0E + "\n"
  content += "color6 #" + theme.Base0C + "\n"
  content += "color7 #" + theme.Base05 + "\n"

  content += "color8 #" + theme.Base03 + "\n"
  content += "color9 #" + theme.Base09 + "\n"
  content += "color10 #" + theme.Base01 + "\n"
  content += "color11 #" + theme.Base02 + "\n"
  content += "color12 #" + theme.Base04 + "\n"
  content += "color13 #" + theme.Base06 + "\n"
  content += "color14 #" + theme.Base0F + "\n"
  content += "color15 #" + theme.Base07 + "\n"

	configPath := os.Getenv("HOME") + "/.config/kitty/theme.conf"

	writeErr := os.WriteFile(configPath, []byte(content), 0644)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
