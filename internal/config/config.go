package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Log struct {
		Level string `yaml:"level"`
	} `yaml: "log"`

	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`
}

func New(name string) (*Config, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var conf Config

	return &conf, yaml.Unmarshal(data, &conf)
}
