package themer

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func updateKittyConfig() {
	cmd := exec.Command("ps", "-e", "-o", "pid,comm")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the output to string
	psOutput := string(output)

	// Split the output by newline to get individual lines
	lines := strings.Split(psOutput, "\n")

	// Iterate over each line
	for _, line := range lines {
		// Split the line by whitespace to get PID and command
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		pid := parts[0]
		command := parts[1]

		if strings.Contains(command, "kitty") {
			killcmd := exec.Command("kill", "-SIGUSR1", pid)
			_, err := killcmd.CombinedOutput()
			if err != nil {
				return
			}
		}
	}
}

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

	updateKittyConfig()

	return nil
}
