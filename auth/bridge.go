package auth

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Ip string `yaml:"ip"`
	User string `yaml:"user"`
}

func LoadConfig(file string) Config {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
