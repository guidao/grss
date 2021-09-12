package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Github *Github `yaml:"github"`
}

type Github struct {
	Token string   `yaml:"token"`
	Repos []string `yaml:"repos"`
}

var (
	conf Config
)

func Init(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	fmt.Println("init config with:", conf.Github)
	return err
}

func GetConf() Config {
	return conf
}
