package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env  string     `yaml:"env" env-default:"dev"`
	GRPC GRPCConfig `yaml:"grpc"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

// Must - without error processing
func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("Config file path is empty.")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("Config file doesn't exist " + path)
	}

	var config Config

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("Failed while reading config file " + err.Error())
	}

	return &config
}

// Priority - flag -> env -> default
func fetchConfigPath() string {
	path := ""

	// --config="path/to/config.yaml"
	flag.StringVar(&path, "config", "", "Config file path")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path

}
