package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	Logger LoggerConfig `yaml:"logger"`
	Server ServerConfig `yaml:"server"`
}

func ReadConfig(path string) (*Config, error) {
	f, _ := os.Open(path)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	conf := Config{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}