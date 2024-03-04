package themer

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Theme struct {
	name string
	Base00 string `yaml:"base00"`
	Base01 string `yaml:"base01"`
	Base02 string `yaml:"base02"`
	Base03 string `yaml:"base03"`
	Base04 string `yaml:"base04"`
	Base05 string `yaml:"base05"`
	Base06 string `yaml:"base06"`
	Base07 string `yaml:"base07"`
	Base08 string `yaml:"base08"`
	Base09 string `yaml:"base09"`
	Base0A string `yaml:"base0A"`
	Base0B string `yaml:"base0B"`
	Base0C string `yaml:"base0C"`
	Base0D string `yaml:"base0D"`
	Base0E string `yaml:"base0E"`
	Base0F string `yaml:"base0F"`
}

func GetBase16(filename string) (Theme, error) {

	themePath := filepath.Join(os.Getenv("HOME"), ".config/themer/", filename+".yaml")
	yamlFile, err := os.ReadFile(themePath)

	if err != nil {
		log.Fatal(err)
		return Theme{}, err
	}

	theme := Theme{ name: filename}

	err = yaml.Unmarshal(yamlFile, &theme)
	if err != nil {
		log.Fatal(err)
		return Theme{}, err
	}

	return theme, nil
}
