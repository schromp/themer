package internal

func ApplyPrograms(theme Theme, transparency string, wallpaper string) {
	ApplyHyprland(theme, transparency)
	ApplyNvim(theme, transparency)
	ApplyKitty(theme, transparency)
	ApplyTmux(theme)
	ApplyZsh(theme)
	ApplyWalker(theme)
	ApplySwww(wallpaper)
}
