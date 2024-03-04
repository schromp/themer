package themer

import (
	"fmt"
	"log"
	"os/exec"
)

func applyBorders(theme Theme) error {
	activeBorders := exec.Command("hyprctl", "keyword", "general:col.active_border", fmt.Sprintf("rgb(%s) rgb(%s)", theme.Base08, theme.Base09))

	activeBordersErr := activeBorders.Err
	if activeBordersErr != nil {
		log.Fatal(activeBordersErr)
		return activeBordersErr
	}

	inactiveBorders := exec.Command("hyprctl", "keyword", "general:col.inactive_border", fmt.Sprintf("rgb(%s)", theme.Base04))

	inactiveBordersErr := inactiveBorders.Err
	if inactiveBordersErr != nil {
		log.Fatal(inactiveBordersErr)
		return inactiveBordersErr
	}

	return nil
}

func ApplyHyprland(theme Theme) error {

	bordersErr := applyBorders(theme)
	if bordersErr != nil {
		return bordersErr
	}

	return nil
}
