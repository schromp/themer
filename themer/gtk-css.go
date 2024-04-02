package themer

func GenerateGTKCSS(theme Theme) string {

	var content string

	content += "/*# This file is generated by themer. Changes may be overwritten.*/\n"
	content += "@define-color base00 #" + theme.Base00 + ";\n"
	content += "@define-color base01 #" + theme.Base01 + ";\n"
	content += "@define-color base02 #" + theme.Base02 + ";\n"
	content += "@define-color base03 #" + theme.Base03 + ";\n"
	content += "@define-color base04 #" + theme.Base04 + ";\n"
	content += "@define-color base05 #" + theme.Base05 + ";\n"
	content += "@define-color base06 #" + theme.Base06 + ";\n"
	content += "@define-color base07 #" + theme.Base07 + ";\n"
	content += "@define-color base08 #" + theme.Base08 + ";\n"
	content += "@define-color base09 #" + theme.Base09 + ";\n"
	content += "@define-color base0A #" + theme.Base0A + ";\n"
	content += "@define-color base0B #" + theme.Base0B + ";\n"
	content += "@define-color base0C #" + theme.Base0C + ";\n"
	content += "@define-color base0D #" + theme.Base0D + ";\n"
	content += "@define-color base0E #" + theme.Base0E + ";\n"
	content += "@define-color base0F #" + theme.Base0F + ";\n"

	return content
}
