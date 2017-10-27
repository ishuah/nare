package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config struct holds configuration variables
type Config struct {
	DownloadDir string `yaml:"download_dir"`
}

// GetConfig returns a Config object with variables held in config.yaml
func GetConfig() (Config, error) {
	configFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return Config{}, err
	}

	var c Config
	err = yaml.Unmarshal(configFile, &c)

	if err != nil {
		return Config{}, err
	}

	return c, nil
}
