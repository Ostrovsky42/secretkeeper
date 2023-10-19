package configparcer

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func GetConfig(filePath string, cfg any) {
	filename, _ := filepath.Abs(filePath)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		panic(err)
	}
}
