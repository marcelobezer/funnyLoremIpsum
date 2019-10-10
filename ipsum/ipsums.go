package ipsum

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// LoremIpsum models a lorem ipsum
type LoremIpsum struct {
	Mussum string `yaml:"mussum"`
}

func getIpsums() (LoremIpsum, error) {
	li := LoremIpsum{}

	path, err := filepath.Abs("../ipsum")
	if err != nil {
		return li, err
	}

	yamlFile, err := ioutil.ReadFile(path + "/loremipsum.yaml")
	if err != nil {
		return li, err
	}

	err = yaml.Unmarshal(yamlFile, &li)
	if err != nil {
		return li, err
	}

	return li, nil
}
