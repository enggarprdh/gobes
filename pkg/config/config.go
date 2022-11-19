package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Addr string `yaml:"addr"`
	Mode string `yaml:"mode"`
}

var config *Config

func Load(path string) error {
	r, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(r, &config)
}

func Get() *Config {
	return config
}
