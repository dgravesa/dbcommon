package dbcommon

import (
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

// ReadConfig loads a database configuration from a file.
func ReadConfig(fname string) (*Config, error) {
	var config Config
	f, err := os.Open(fname)

	if err == nil {
		d := yaml.NewDecoder(f)
		err = d.Decode(&config)
	}

	return &config, err
}
