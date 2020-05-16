package main

import (
	"os"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Level string `yaml:"level"`
	File string `yaml:"file"`
}

func ConfigureLogger(conf *LoggerConfig) (*logrus.Logger, error) {
	lg := logrus.New()

	level, err := logrus.ParseLevel(conf.Level)
	if err != nil {
		return nil, err
	}
	lg.SetLevel(level)
		if conf.File != "" {
			f, err := os.Create(conf.File)
			if err != nil {
				return nil, err
			}
			lg.SetOutput(f)
		}
	return lg, nil
}