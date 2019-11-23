package dbserver

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

// Config provides a database server configuration
type Config struct {
	Username string `yaml:"username"`
	Passkey  string `yaml:"passkey"`
	Host     string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

// ReadConfigFile loads a database configuration from a file.
func ReadConfigFile(fname string) (Config, error) {
	f, err := os.Open(fname)

	if err != nil {
		return Config{}, err
	}

	return ReadConfig(f)
}

// ReadConfig loads a database configuration.
func ReadConfig(r io.Reader) (Config, error) {
	var config Config
	d := yaml.NewDecoder(r)
	err := d.Decode(&config)
	return config, err
}
