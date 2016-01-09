package amproxy

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

type config struct {
	Apikeys map[string]Creds `yaml:"apikeys"`
}

type Creds struct {
	SecretKey string           `yaml:"secret_key"`
	Metrics   map[string]uint8 `yaml:"metrics"`
}

func LoadUserConfigFile(filename string) (map[string]Creds, time.Time) {
	var f config

	info, err := os.Stat(filename)
	if err != nil {
		panic(err)
	}

	yamlFile, err := ioutil.ReadFile(filename)
	err = yaml.Unmarshal(yamlFile, &f)
	if err != nil {
		panic(err)
	}

	return f.Apikeys, info.ModTime()
}
