package configparcer

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
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
