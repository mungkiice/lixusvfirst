package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App struct {
		Key string `yaml:"key"`
	} `yaml:"app"`
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"database"`
}

func GetObject() Config {
	f, err := os.Open("/usr/share/nginx/html/config.yml")
	if err != nil {
		log.Fatal("Error while opening config file")
	}
	defer f.Close()
	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal("Error while decoding config file")
	}
	return cfg
}
