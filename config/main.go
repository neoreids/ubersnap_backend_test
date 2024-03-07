package config

import (
	"github.com/rotisserie/eris"
	"gopkg.in/yaml.v2"
	"os"
	"ubersnap/static"
	"ubersnap/utilities"
)

var c *Config

type (
	Server struct {
		Port int `yaml:"port"`
	}

	Config struct {
		Production bool     `yaml:"production"`
		SecretKey  string   `yaml:"secret_key"`
		Server     Server   `yaml:"server"`
		MimesImage []string `yaml:"mimes_image"`
	}
)

func Set() *Config {
	file, err := os.Open("./config.yml")
	if err != nil {
		panic(static.FILE_CONFIG_NOT_FOUND)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(eris.Wrap(err, "failed close file config"))
		}
	}(file)

	decoder := yaml.NewDecoder(file)
	var config Config
	if err := decoder.Decode(&config); err != nil {
		panic(eris.Wrap(err, "error decode"))
	}

	return &config
}

func Get() *Config {
	return c
}

func (c *Config) IsValidImage(mime string) (bool, error) {
	if len(c.MimesImage) < 1 {
		return false, static.CONFIG_NOT_FOUND
	}

	return utilities.ContainsInArray(c.MimesImage, mime), nil
}