package dbcommon

import (
	"os"

	"gopkg.in/yaml.v2"
)

// DBConfig provides a database server configuration
type DBConfig struct {
	Username string `yaml:"username"`
	Passkey  string `yaml:"passkey"`
	Host     string `yaml:"hostname"`
	Port     int    `yaml:"port"`
}

// ReadDBConfig loads a database configuration from a file.
func ReadDBConfig(fname string) (*DBConfig, error) {
	var config DBConfig
	f, err := os.Open(fname)

	if err == nil {
		d := yaml.NewDecoder(f)
		err = d.Decode(&config)
	}

	return &config, err
}
