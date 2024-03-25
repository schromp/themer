package themer

import (
	"fmt"
	"log"
	"os/exec"
)

func applyBorders(theme Theme) error {
	activeBorders := exec.Command("hyprctl", "keyword", "general:col.active_border", fmt.Sprintf("rgb(%s) rgb(%s)", theme.Base0D, theme.Base08))

	activeBordersErr := activeBorders.Run()
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

func applyBackground(bg string) error {
	if bg == "solid" || bg == "transparent" {
		err := exec.Command("hyprctl", "keyword", "decorations:blur:enabled", "false").Run()
		if err != nil {
			return err
		}

	} else if bg == "blur" {
		err := exec.Command("hyprctl", "keyword", "decorations:blur:enabled", "true").Run()
		if err != nil {
			return err
		}

		err = exec.Command("hyprctl", "keyword", "decorations:blur:size", "8").Run()
		if err != nil {
			return err
		}

		err = exec.Command("hyprctl", "keyword", "decorations:blur:passes", "3").Run()
		if err != nil {
			return err
		}

	} else {
		return fmt.Errorf("Invalid background type")
	}
	return nil
}

func ApplyHyprland(theme Theme, bg string) error {

	bordersErr := applyBorders(theme)
	if bordersErr != nil {
		return bordersErr
	}

	bgErr := applyBackground(bg)
	if bgErr != nil {
		return bgErr
	}

	return nil
}
